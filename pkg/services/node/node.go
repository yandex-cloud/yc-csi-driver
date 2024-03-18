/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/mount-utils"
	"k8s.io/utils/exec"

	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/instancemeta"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
)

const (
	extensiveMountPointCheck = true
)

type Options struct {
	MaxVolumesPerNode int
}

func DefaultOptions() Options {
	return Options{MaxVolumesPerNode: 5}
}

type FileXattrs interface {
	SetImmutable(target string) error
	UnsetImmutable(target string) error
}

type Service struct {
	metadataGetter   instancemeta.MetadataGetter
	mounter          *mount.SafeFormatAndMount
	inFlight         inflight.InFlight
	hostUtil         hostutil.HostUtils
	volumeStats      VolumeStats
	fileXattrs       FileXattrs
	nodeCapabilities []*csi.NodeServiceCapability
	opts             Options
}

var _ csi.NodeServer = &Service{}

func New(mounter *mount.SafeFormatAndMount, hostUtil hostutil.HostUtils, volStat VolumeStats,
	fileXattrs FileXattrs, metadataGetter instancemeta.MetadataGetter, caps []*csi.NodeServiceCapability, opts Options) *Service {
	return &Service{
		metadataGetter:   metadataGetter,
		mounter:          mounter,
		inFlight:         inflight.NewWithTTL(),
		hostUtil:         hostUtil,
		volumeStats:      volStat,
		fileXattrs:       fileXattrs,
		nodeCapabilities: caps,
		opts:             opts,
	}
}

func (n *Service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.Infof("StageVolume(%+v)", req)

	volumeID := req.GetVolumeId()

	if len(volumeID) == 0 {
		klog.Errorln("Volume ID not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	target := req.GetStagingTargetPath()
	if len(target) == 0 {
		klog.Errorln("Staging target not provided")
		return nil, status.Error(codes.InvalidArgument, "Staging target not provided")
	}

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		klog.Errorln("Volume capability not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	if err := services.CheckVolumeAccessType(volCap); err != nil {
		klog.Errorln(err)
		return nil, err
	}

	err := services.VerifyVolumeCapabilities([]*csi.VolumeCapability{volCap}, services.VolumeCaps)
	if err != nil {
		klog.Errorln(err)
		return nil, err
	}

	switch volCap.GetAccessType().(type) {
	case *csi.VolumeCapability_Block:
		klog.Infof("BlockVolume: %+s is successfully staged", volumeID)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	if ok := n.inFlight.Insert(req); !ok {
		msg := fmt.Sprintf("request to stage volume=%q is already in progress", volumeID)
		klog.Infoln(msg)
		return nil, status.Error(codes.Aborted, msg)
	}

	defer func() {
		klog.Infof("StageVolume: volume=%q operation finished", req.GetVolumeId())
		n.inFlight.Delete(req)
	}()

	deviceSymlink := fmt.Sprintf("%s-%s", services.DevicePathPrefix, volumeID)

	device, err := n.followDeviceSymlink(deviceSymlink)
	if err != nil {
		klog.Infof("Failed to find device path %s. %v", deviceSymlink, err)
		return nil, status.Errorf(codes.Internal, "Failed to find device path %s. %v", deviceSymlink, err)
	}

	klog.Errorf("StageVolume: found device path %s -> %s", deviceSymlink, device)

	exists, err := n.hostUtil.PathExists(target)
	if err != nil {
		klog.Infof("failed to check if target %q exists: %v", target, err)
		msg := fmt.Sprintf("failed to check if target %q exists: %v", target, err)
		return nil, status.Error(codes.Internal, msg)
	}

	// When exists is true it means target path was created but device isn't mounted.
	// We don't want to do anything in that case and let the operation proceed.
	// Otherwise, we need to create the target directory.
	if !exists {
		// If target path does not exist we need to create the directory where volume will be staged.
		klog.Infof("StageVolume: creating target dir %q", target)
		if err = makeDir(target); err != nil {
			klog.Infof("could not create target dir %q: %v", target, err)
			msg := fmt.Sprintf("could not create target dir %q: %v", target, err)
			return nil, status.Error(codes.Internal, msg)
		}
	}

	// Check if a device is mounted in target directory.
	deviceFromMount, _, err := mount.GetDeviceNameFromMount(n.mounter, target)
	if err != nil {
		klog.Infof("failed to check if volume is already mounted: %v", err)
		msg := fmt.Sprintf("failed to check if volume is already mounted: %v", err)
		return nil, status.Error(codes.Internal, msg)
	}

	// This operation (NodeStageVolume) MUST be idempotent.
	// If the volume corresponding to the volume_id is already staged to the staging_target_path,
	if deviceFromMount == device {
		klog.Infof("StageVolume: volume ( %s ) already staged", volumeID)
		return &csi.NodeStageVolumeResponse{}, nil
	}

	// If something else mounted to our target.
	if deviceFromMount != "" {
		msg := fmt.Sprintf("Device ( %q ) mounted to stagingPath ( %q ) differs from our device ( %q )", deviceFromMount, target, device)
		klog.Error(msg)
		return nil, status.Error(codes.FailedPrecondition, msg)
	}

	// Setting chattr +i on staging target as it helps to prevent rootfs usage.
	klog.Infof("StageVolume: setting chattr +i on %s", target)
	if chattrErr := n.fileXattrs.SetImmutable(target); chattrErr != nil {
		return nil, status.Error(codes.Internal,
			fmt.Sprintf("chattr +i on %q failed: %v", target, chattrErr))
	}

	fsType := req.GetVolumeCapability().GetMount().GetFsType()
	// FormatAndMount will format only if needed.
	if len(fsType) == 0 {
		fsType = services.DefaultFsType
	}

	mountOptions := req.GetVolumeCapability().GetMount().GetMountFlags()

	// When volume is published as RO, stage it as RO.
	if services.NodeStageRO(req.GetPublishContext()) {
		// "ro" && "noatime" are general (fs independent) options. We shall always use them.
		// "noload" is a fs specific option.
		mountOptions = append(mountOptions, "ro", "noatime")
		switch fsType {
		case services.FsTypeExt3, services.FsTypeExt4:
			mountOptions = append(mountOptions, "noload")
		}
	}

	klog.Infof("StageVolume: formatting %s and mounting at %s", device, target)
	err = n.mounter.FormatAndMount(device, target, fsType, mountOptions)
	if err != nil {
		klog.Errorf("could not format %q and mount it at %q (%+v)", device, target, err)
		msg := fmt.Sprintf("could not format %q and mount it at %q (%+v)", device, target, err)
		return nil, status.Error(codes.Internal, msg)
	}

	return &csi.NodeStageVolumeResponse{}, nil
}

func (n *Service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.Infof("UnstageVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		klog.Errorln("VolumeID not provided")
		return nil, status.Error(codes.InvalidArgument, "VolumeID not provided")
	}

	target := req.GetStagingTargetPath()
	if len(target) == 0 {
		klog.Errorln("StagingTargetPath not provided")
		return nil, status.Error(codes.InvalidArgument, "StagingTargetPath not provided")
	}

	if ok := n.inFlight.Insert(req); !ok {
		msg := fmt.Sprintf("request to unstage volume=%q is already in progress", volumeID)
		klog.Infoln(msg)
		return nil, status.Error(codes.Aborted, msg)
	}

	defer func() {
		klog.Infof("UnstageVolume: volume=%q operation finished", req.GetVolumeId())
		n.inFlight.Delete(req)
	}()

	// Check if target directory is a mount point. GetDeviceNameFromMount
	// given a mnt point, finds the device from /proc/mounts
	// returns the device name, reference count, and error code
	deviceFromMount, refCount, err := mount.GetDeviceNameFromMount(n.mounter, target)
	if err != nil {
		msg := fmt.Sprintf("failed to check if volume is mounted: %v", err)
		klog.Errorln(msg)
		return nil, status.Error(codes.Internal, msg)
	}

	// From the spec: If the volume corresponding to the volume_id
	// is not staged to the staging_target_path, the Plugin MUST
	// reply 0 OK.
	if refCount == 0 {
		klog.Infof("UnstageVolume: %s target not mounted", target)

		// Running chattr -i to be sure that we haven't left +i attr on globalmount
		// during previous UnstageVolume.
		klog.Infof("UnstageVolume: running chattr -i on %s", target)
		if chattrErr := n.fileXattrs.UnsetImmutable(target); chattrErr != nil {
			klog.Errorf("chattr -i on %q failed: %v", target, chattrErr)
		}

		return &csi.NodeUnstageVolumeResponse{}, nil
	}

	if refCount > 1 {
		klog.Warningf("UnstageVolume: found %d references to device %s mounted at target path %s",
			refCount, deviceFromMount, target)
	}

	klog.Infof("UnstageVolume: unmounting %s", target)

	err = n.mounter.Unmount(target)
	if err != nil {
		klog.Errorf("Could not unmount target %q: %+v", target, err)
		return nil, status.Errorf(codes.Internal, "Could not unmount target %q: %+v", target, err)
	}

	klog.Infof("UnstageVolume: completed for device: %s and staging_path %s", deviceFromMount, target)

	// Checking context canceled to decide if we need to leave immutable attr on staging path.
	select {
	case <-ctx.Done():
		klog.Errorf("UnstageVolume canceled for: %s and staging_path %s", deviceFromMount, target)
		return nil, status.Errorf(codes.Canceled, "UnstageVolume canceled for: %s and staging_path %s", deviceFromMount, target)
	default:
		klog.Infof("UnstageVolume: running chattr -i on %s", target)
		if chattrErr := n.fileXattrs.UnsetImmutable(target); chattrErr != nil {
			klog.Errorf("chattr -i on %q failed: %v", target, chattrErr)
		}
		return &csi.NodeUnstageVolumeResponse{}, nil
	}
}

func (n *Service) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("PublishVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		klog.Errorln("Volume ID not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	stagingPath := req.GetStagingTargetPath()
	if len(stagingPath) == 0 {
		klog.Errorln("Staging target not provided")
		return nil, status.Error(codes.InvalidArgument, "Staging target not provided")
	}

	targetPath := req.GetTargetPath()
	if len(targetPath) == 0 {
		klog.Errorln("Target path not provided")
		return nil, status.Error(codes.InvalidArgument, "Target path not provided")
	}

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		klog.Errorln("Volume capability not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	if err := services.CheckVolumeAccessType(volCap); err != nil {
		klog.Errorln(err)
		return nil, err
	}

	volCaps := []*csi.VolumeCapability{volCap}

	err := services.VerifyVolumeCapabilities(volCaps, services.VolumeCaps)
	if err != nil {
		klog.Errorln(err)
		return nil, err
	}

	if ok := n.inFlight.Insert(req); !ok {
		msg := fmt.Sprintf("request to publish volume=%q is already in progress", volumeID)
		klog.Infoln(msg)
		return nil, status.Error(codes.Aborted, msg)
	}

	defer func() {
		klog.Infof("PublishVolume: volume=%q operation finished", req.GetVolumeId())
		n.inFlight.Delete(req)
	}()

	mountOptions := []string{"bind"}

	if req.GetReadonly() {
		mountOptions = append(mountOptions, "ro")
	}

	switch volCap.GetAccessType().(type) {
	case *csi.VolumeCapability_Block:
		if err := n.nodePublishVolumeForBlock(req, mountOptions); err != nil {
			return nil, err
		}
	case *csi.VolumeCapability_Mount:
		if err := n.nodePublishVolumeForFileSystem(req, mountOptions, volCap.GetMount()); err != nil {
			return nil, err
		}
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (n *Service) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.Infof("UnpublishVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		klog.Errorln("VolumeID not provided")
		return nil, status.Error(codes.InvalidArgument, "VolumeID not provided")
	}

	target := req.GetTargetPath()
	if len(target) == 0 {
		klog.Errorln("TargetPath not provided")
		return nil, status.Error(codes.InvalidArgument, "TargetPath not provided")
	}

	if ok := n.inFlight.Insert(req); !ok {
		msg := fmt.Sprintf("request to unpublish volume=%q is already in progress", volumeID)
		klog.Infoln(msg)
		return nil, status.Error(codes.Aborted, msg)
	}

	defer func() {
		klog.Infof("UnpublishVolume: volume=%q operation finished", req.GetVolumeId())
		n.inFlight.Delete(req)
	}()

	err := mount.CleanupMountPoint(target, n.mounter, extensiveMountPointCheck)
	if err != nil {
		klog.Errorf("Error calling mount.CleanupMountPoint on %q with error: %+v", target, err)
		return nil, err
	}

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (n *Service) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	klog.Infof("NodeGetVolumeStats(%+v)", req)

	volumeID := req.GetVolumeId()

	if len(volumeID) == 0 {
		klog.Errorln("Volume ID not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	target := req.GetVolumePath()
	if len(target) == 0 {
		klog.Errorln("VolumePath not provided")
		return nil, status.Error(codes.InvalidArgument, "VolumePath not provided")
	}

	isTargetExists, err := n.hostUtil.PathExists(target)
	if err != nil {
		return nil, err
	}
	if !isTargetExists {
		return nil, status.Errorf(codes.NotFound, "target (%s) not found", target)
	}

	isDevice, err := n.hostUtil.PathIsDevice(target)
	if err != nil {
		return nil, err
	}

	if isDevice {
		return &csi.NodeGetVolumeStatsResponse{}, nil
	}

	devicePathSymlink := fmt.Sprintf("%s-%s", services.DevicePathPrefix, volumeID)

	_, err = n.followDeviceSymlink(devicePathSymlink)
	if err != nil {
		klog.Infof("Failed to find device path %s. %v", devicePathSymlink, err)
		return nil, status.Errorf(codes.NotFound, "Failed to find device path %s. %v", devicePathSymlink, err)
	}

	bytes, inodes, err := n.volumeStats.GetVolumeStats(target)
	if err != nil {
		klog.Infof("getVolumeStats error on target: %s: %+v", target, err)
		return nil, status.Errorf(codes.Internal, "getVolumeStats error on target: %s, %+v", target, err)
	}

	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{Available: bytes.Available, Total: bytes.Total, Used: bytes.Used, Unit: csi.VolumeUsage_BYTES},
			{Available: inodes.Available, Total: inodes.Total, Used: inodes.Used, Unit: csi.VolumeUsage_INODES},
		},
	}, nil
}

func (n *Service) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	klog.Infof("NodeExpandVolume (%+v)", req)
	volumeID := req.GetVolumeId()

	if len(volumeID) == 0 {
		klog.Errorln("Volume ID not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	target := req.GetVolumePath()
	if len(target) == 0 {
		klog.Errorln("VolumePath not provided")
		return nil, status.Error(codes.InvalidArgument, "VolumePath not provided")
	}

	isDevice, err := n.hostUtil.PathIsDevice(target)
	if err != nil {
		return nil, err
	}

	if isDevice {
		return &csi.NodeExpandVolumeResponse{}, nil
	}

	devicePath, _, err := mount.GetDeviceNameFromMount(n.mounter, target)
	if err != nil || len(devicePath) == 0 {
		klog.Errorf("No devicePath for stagingPath ( %q ) available ( %+v )", target, err)
		return nil, status.Errorf(codes.Internal, "No devicePath for stagingPath ( %q ) available ( %+v )", target, err)
	}

	_, resizeErr := mount.NewResizeFs(n.mounter.Exec).Resize(devicePath, target)
	if resizeErr != nil {
		klog.Errorf("Error resizing volume: ( %q ); stagingPath: ( %q ) : ( %+v )", volumeID, target, resizeErr)
		return nil, status.Errorf(codes.Internal, "Error resizing volume: ( %q ); stagingPath: ( %q ) : ( %+v )",
			volumeID, target, resizeErr)
	}

	return &csi.NodeExpandVolumeResponse{}, nil
}

func (n *Service) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.Info("NodeGetCapabilities")
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: n.nodeCapabilities,
	}, nil
}

func (n *Service) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	instanceID, err := n.metadataGetter.InstanceID()
	if err != nil {
		klog.Errorf("NodeGetInfo failed: %s", err.Error())
		return nil, err
	}

	zoneID, err := n.metadataGetter.ZoneID()
	if err != nil {
		klog.Errorf("NodeGetInfo failed: %s", err.Error())
		return nil, err
	}

	resp := &csi.NodeGetInfoResponse{
		NodeId: instanceID,
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				services.ZoneKey: zoneID,
			},
		},
		MaxVolumesPerNode: int64(n.opts.MaxVolumesPerNode),
	}

	klog.Infof("NodeGetInfo, returning %+v", resp)

	return resp, nil
}

func (n *Service) followDeviceSymlink(devicePath string) (string, error) {
	//PathIsDevice also checks existence in case of hostutil_linux.
	isDevice, err := n.hostUtil.PathIsDevice(devicePath)
	if err != nil {
		return "", fmt.Errorf("error checking if target ( %s ) is device: %+v", devicePath, err)
	}

	if !isDevice {
		return "", fmt.Errorf("file %s found, but was not a device", devicePath)
	}

	// Find the target, resolving to an absolute path.
	// For example, /dev/disk/by-id/virtio-epd8f2nvqehnl37fqfr0 -> ../../vdb
	resolved, err := n.hostUtil.EvalHostSymlinks(devicePath)
	if err != nil {
		return "", fmt.Errorf("error reading target of symlink %q: %+v", devicePath, err)
	}

	return resolved, nil
}

func (n *Service) nodePublishVolumeForFileSystem(req *csi.NodePublishVolumeRequest,
	mountOptions []string, volumeMountCap *csi.VolumeCapability_MountVolume) error {

	targetPath := req.GetTargetPath()
	stagingPath := req.GetStagingTargetPath()

	if m := volumeMountCap; m != nil {
		hasOption := func(options []string, opt string) bool {
			for _, o := range options {
				if o == opt {
					return true
				}
			}
			return false
		}

		for _, f := range m.MountFlags {
			if !hasOption(mountOptions, f) {
				mountOptions = append(mountOptions, f)
			}
		}
	}

	klog.Infof("PublishMountVolume: creating dir %s", targetPath)
	if err := makeDir(targetPath); err != nil {
		return status.Errorf(codes.Internal, "Could not create dir %q: %+v", targetPath, err)
	}

	isMountPoint, err := n.mounter.IsMountPoint(targetPath)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not check if %q is a mountpoint: %+v", targetPath, err)
	}

	if isMountPoint {
		klog.Infof("targetPath (%q) is already mounted", targetPath)
		return nil
	}

	klog.Infof("PublishMountVolume: mounting %s at %s", stagingPath, targetPath)

	err = n.mounter.Mount(stagingPath, targetPath, "", mountOptions)
	if err != nil {
		if removeErr := os.Remove(targetPath); removeErr != nil {
			return status.Errorf(codes.Internal,
				"Could not remove mount target on error %s: %+v ( error: %+v )", targetPath, removeErr, err)
		}

		return status.Errorf(codes.Internal, "Could not mount %s at %s: %+v", stagingPath, targetPath, err)
	}

	return err
}

func (n *Service) nodePublishVolumeForBlock(req *csi.NodePublishVolumeRequest,
	mountOptions []string) error {

	volumeID := req.GetVolumeId()
	targetPath := req.GetTargetPath()

	deviceSymlink := fmt.Sprintf("%s-%s", services.DevicePathPrefix, volumeID)

	device, err := n.followDeviceSymlink(deviceSymlink)
	if err != nil {
		klog.Infof("Failed to find device path %s. %v", deviceSymlink, err)
		return status.Errorf(codes.Internal, "Failed to find device path %s. %v", deviceSymlink, err)
	}

	klog.Infof("PublishBlockVolume: found device path %s -> %s", deviceSymlink, device)
	klog.Infof("PublishBlockVolume: creating parent directory for device file %s", filepath.Dir(targetPath))
	if err := makeDir(filepath.Dir(targetPath)); err != nil {
		return status.Errorf(codes.Internal, "Could not create dir %s: %+v", filepath.Dir(targetPath), err)
	}

	klog.V(5).Infof("PublishBlockVolume creating device file %s", targetPath)

	err = makeFile(targetPath)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not create file %s: ( %+v )", targetPath, err)
	}

	isMountPoint, err := n.mounter.IsMountPoint(targetPath)
	if err != nil {
		return status.Errorf(codes.Internal, "Could not check if %q is a mountpoint: %+v", targetPath, err)
	}

	if isMountPoint {
		klog.Infof("targetPath (%q) is already mounted", targetPath)
		return nil
	}

	klog.V(5).Infof("PublishBlockVolume mounting %s at %s", device, targetPath)

	err = n.mounter.Mount(device, targetPath, "", mountOptions)
	if err != nil {
		if removeErr := os.Remove(targetPath); removeErr != nil {
			return status.Errorf(codes.Internal,
				"Could not remove mount target on error %s: %v ( error: %+v )", targetPath, removeErr, err)
		}

		return status.Errorf(codes.Internal, "Could not mount %s at %s: %+v", device, targetPath, err)
	}

	return err
}

func NewSafeMounter() *mount.SafeFormatAndMount {
	return &mount.SafeFormatAndMount{
		Interface: mount.New(""),
		Exec:      exec.New(),
	}
}

// makeDir creates a new directory.
// If pathname already exists as a directory, no error is returned.
// If pathname already exists as a file, an error is returned.
func makeDir(pathname string) error {
	err := os.MkdirAll(pathname, os.FileMode(0755))
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

func makeFile(pathname string) error {
	f, err := os.OpenFile(pathname, os.O_CREATE, os.FileMode(0644))
	defer func() {
		_ = f.Close()
	}()

	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}
