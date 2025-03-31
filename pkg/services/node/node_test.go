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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/kubernetes/pkg/volume/util/hostutil"
	"k8s.io/mount-utils"

	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
)

// This tests should run on linux platform inside docker container.

const (
	volumeID        = "vol-test"
	fakeDiskPath    = "/dev/fake"
	fakeStagingPath = "/var/tmp/path"
	fakeDiskSymlink = services.DevicePathPrefix + "-" + volumeID

	fakeInstID = "fake-instance"
	fakeZoneID = "fake-zone-id"
)

type fakeHostUtil struct {
	*hostutil.FakeHostUtil
	FakeDeviceSymlinks map[string]string
}

func (f *fakeHostUtil) GetMode(pathname string) (os.FileMode, error) {
	_, err := f.EvalHostSymlinks(pathname)
	if err != nil {
		return 0, err
	}
	return os.ModeDevice, nil
}

func (f *fakeHostUtil) EvalHostSymlinks(pathname string) (string, error) {
	target, ok := f.FakeDeviceSymlinks[pathname]
	if !ok {
		return "", status.Errorf(codes.NotFound, "%s is not a device symlink", pathname)
	}
	return target, nil
}

func (f *fakeHostUtil) PathIsDevice(pathname string) (bool, error) {
	t, err := f.GetFileType(pathname)
	if err != nil {
		return false, err
	}
	switch t {
	case hostutil.FileTypeBlockDev:
		return true, nil
	case hostutil.FileTypeFile:
		_, ok := f.FakeDeviceSymlinks[pathname]
		return ok, nil
	default:
		return false, nil
	}
}

func newFakeHostUtil() *fakeHostUtil {
	return &fakeHostUtil{
		FakeHostUtil: hostutil.NewFakeHostUtil(nil),
	}
}

func TestNodeGetVolumeStats(t *testing.T) {
	err := os.MkdirAll(fakeStagingPath, 0755)
	require.NoError(t, err)
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)
	testCases := []struct {
		name string
		req  *csi.NodeGetVolumeStatsRequest

		// fakeMounter mocks mounter behaviour.
		fakeMounter *mount.FakeMounter

		// fakeHostUtil mocks hostutil behaviour.
		fakeHostUtil *fakeHostUtil

		// fake filexattrs mock setchattr.
		fakeFileXattrs *fakeFileXattrs

		// expected test error code.
		expErrCode codes.Code
	}{
		{
			name: "success",
			fakeMounter: &mount.FakeMounter{
				MountPoints: []mount.MountPoint{{Device: fakeDiskPath, Path: fakeStagingPath}},
			},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
			req: &csi.NodeGetVolumeStatsRequest{
				VolumeId:   volumeID,
				VolumePath: fakeStagingPath,
			},
		},
		{
			name: "No VolumeID",
			fakeMounter: &mount.FakeMounter{
				MountPoints: []mount.MountPoint{{Device: fakeDiskPath, Path: fakeStagingPath}},
			},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskSymlink: hostutil.FileTypeFile,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			req: &csi.NodeGetVolumeStatsRequest{
				VolumeId:   "",
				VolumePath: fakeStagingPath,
			},
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "No VolumePath",
			fakeMounter: &mount.FakeMounter{
				MountPoints: []mount.MountPoint{{Device: fakeDiskPath, Path: fakeStagingPath}},
			},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskSymlink: hostutil.FileTypeFile,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			req: &csi.NodeGetVolumeStatsRequest{
				VolumeId:   volumeID,
				VolumePath: "",
			},
			expErrCode: codes.InvalidArgument,
		},
		{
			name:        "Device",
			fakeMounter: &mount.FakeMounter{},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath: hostutil.FileTypeBlockDev}),
			},
			req: &csi.NodeGetVolumeStatsRequest{
				VolumeId:   volumeID,
				VolumePath: fakeDiskPath,
			},
			expErrCode: codes.OK,
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			driver := NewTestNodeService(tc.fakeMounter, tc.fakeHostUtil, tc.fakeFileXattrs, mdGetter, services.ProductionNodeCaps)
			r, err := driver.NodeGetVolumeStats(context.TODO(), tc.req)
			fmt.Println("Response: ", r)
			if err != nil {
				srvErr, ok := status.FromError(err)
				if !ok {
					t.Fatalf("Could not get error status code from error: %v", srvErr)
				}
				if srvErr.Code() != tc.expErrCode {
					t.Fatalf("Expected error code %d, got %d message %s", tc.expErrCode, srvErr.Code(), srvErr.Message())
				}
			} else if tc.expErrCode != codes.OK {
				t.Fatalf("Expected error %v, got no error", tc.expErrCode)
			}
		})
	}
}

func TestNodeStageVolume(t *testing.T) {
	var expectedOpts []string
	mountTargetPathPrefix := ""
	expectedOpts = []string{"defaults"}
	if runtime.GOOS != "linux" {
		expectedOpts = []string{}
		mountTargetPathPrefix = "/private"
	}
	stagingPath := mountTargetPathPrefix + fakeStagingPath
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)

	//prepareDevices(t)
	stdVolCap := setVolCap(services.DefaultFsType)
	ext3VolCap := setVolCap(services.FsTypeExt3)
	testCases := []struct {
		name string
		req  *csi.NodeStageVolumeRequest
		// fakeMounter mocks mounter behaviour
		fakeMounter *mount.FakeMounter
		// fakeHostUtil mocks hostutil behaviour
		fakeHostUtil *fakeHostUtil
		// fake filexattrs mock setchattr
		fakeFileXattrs *fakeFileXattrs
		// expected fake mount actions the test will make
		expActions []mount.FakeAction
		// expected test error code
		expErrCode codes.Code
		// expected mount points when test finishes
		expMountPoints []mount.MountPoint
	}{
		{
			name: "success normal",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeCapability:  stdVolCap,
				VolumeId:          volumeID,
			},
			fakeMounter: &mount.FakeMounter{},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
			expActions: []mount.FakeAction{
				{
					Action: "mount",
					Target: stagingPath,
					Source: fakeDiskPath,
					FSType: services.DefaultFsType,
				},
			},
			expMountPoints: []mount.MountPoint{
				{
					Device: fakeDiskPath,
					Opts:   expectedOpts,
					Path:   stagingPath,
					Type:   services.DefaultFsType,
				},
			},
		},
		{
			name: "success RO normal",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeCapability:  stdVolCap,
				VolumeId:          volumeID,
				PublishContext: map[string]string{
					"MULTI_NODE_READER_ONLY": "true",
				},
			},
			fakeMounter: &mount.FakeMounter{},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
			expActions: []mount.FakeAction{
				{
					Action: "mount",
					Target: stagingPath,
					Source: fakeDiskPath,
					FSType: services.DefaultFsType,
				},
			},
			expMountPoints: []mount.MountPoint{
				{
					Device: fakeDiskPath,
					Opts:   []string{"ro", "noatime", "noload", "defaults"},
					Path:   stagingPath,
					Type:   services.DefaultFsType,
				},
			},
		},
		{
			name: "success mount options fsType ext3",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: fakeStagingPath,
				VolumeCapability:  ext3VolCap,
				VolumeId:          volumeID,
			},
			fakeMounter: &mount.FakeMounter{},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
			expActions: []mount.FakeAction{
				{
					Action: "mount",
					Target: stagingPath,
					Source: fakeDiskPath,
					FSType: ext3VolCap.GetMount().GetFsType(),
				},
			},
			expMountPoints: []mount.MountPoint{
				{
					Device: fakeDiskPath,
					Opts:   expectedOpts,
					Path:   stagingPath,
					Type:   ext3VolCap.GetMount().GetFsType(),
				},
			},
		},
		{
			// To test idempotency we need to test the
			// volume corresponding to the volume_id is
			// already staged to the staging_target_path
			// and the Plugin replied with OK. To achieve
			// this we setup the fake mounter to return
			// that /dev/fake is mounted at /test/path.
			name: "success device already mounted at target",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeCapability:  stdVolCap,
				VolumeId:          volumeID,
			},
			fakeMounter: &mount.FakeMounter{
				MountPoints: []mount.MountPoint{{Device: fakeDiskPath, Path: stagingPath}},
			},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
			// no actions means mount isn't called because
			// device is already mounted
			expActions: []mount.FakeAction{},
			// expMountPoints should contain only the
			// fakeMountPoint
			expMountPoints: []mount.MountPoint{
				{
					Device: fakeDiskPath,
					Path:   stagingPath,
				},
			},
		},
		{
			name: "fail no VolumeId",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeCapability:  stdVolCap,
			},
			fakeMounter:  newFakeMounter(),
			fakeHostUtil: newFakeHostUtil(),
			expErrCode:   codes.InvalidArgument,
		},
		{
			name: "fail no StagingTargetPath",
			req: &csi.NodeStageVolumeRequest{
				VolumeCapability: stdVolCap,
				VolumeId:         volumeID,
			},
			fakeMounter:  newFakeMounter(),
			fakeHostUtil: newFakeHostUtil(),
			expErrCode:   codes.InvalidArgument,
		},
		{
			name: "fail no VolumeCapability",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeId:          volumeID,
			},
			fakeMounter:  newFakeMounter(),
			fakeHostUtil: newFakeHostUtil(),
			expErrCode:   codes.InvalidArgument,
		},
		{
			name: "fail invalid VolumeCapabilityAccessMode",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: fakeStagingPath,
				VolumeCapability: &csi.VolumeCapability{
					AccessMode: &csi.VolumeCapability_AccessMode{
						Mode: csi.VolumeCapability_AccessMode_UNKNOWN,
					},
				},
				VolumeId: volumeID,
			},
			fakeMounter:  newFakeMounter(),
			fakeHostUtil: newFakeHostUtil(),
			expErrCode:   codes.InvalidArgument,
		},
		{
			name: "fail invalid VolumeCapabilityAccessType",
			req: &csi.NodeStageVolumeRequest{
				StagingTargetPath: fakeStagingPath,
				VolumeCapability:  VolAccessTypeNotSet(),

				VolumeId: volumeID,
			},
			fakeMounter:  newFakeMounter(),
			fakeHostUtil: newFakeHostUtil(),
			expErrCode:   codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		fmt.Println("TC: ", tc)
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			driver := NewTestNodeService(tc.fakeMounter, tc.fakeHostUtil, tc.fakeFileXattrs, mdGetter, services.ProductionNodeCaps)
			r, err := driver.NodeStageVolume(context.TODO(), tc.req)
			fmt.Println("Response: ", r)
			if err != nil {
				srvErr, ok := status.FromError(err)
				if !ok {
					t.Fatalf("Could not get error status code from error: %v", srvErr)
				}
				if srvErr.Code() != tc.expErrCode {
					t.Fatalf("Expected error code %d, got %d message %s", tc.expErrCode, srvErr.Code(), srvErr.Message())
				}
			} else if tc.expErrCode != codes.OK {
				t.Fatalf("Expected error %v, got no error", tc.expErrCode)
			}

			if !equalActions(tc.fakeMounter.GetLog(), tc.expActions) {
				t.Fatalf("Expected actions {%+v}, got {%+v}", tc.expActions, tc.fakeMounter.GetLog())
			}
			if !equalMountPoints(tc.fakeMounter.MountPoints, tc.expMountPoints) {
				t.Fatalf("Expected mount points {%+v}, got {%+v}", tc.expMountPoints, tc.fakeMounter.MountPoints)
			}
		})
	}
}

func TestNodeUnstageVolume(t *testing.T) {
	mountTargetPathPrefix := ""
	stagingPath := mountTargetPathPrefix + fakeStagingPath
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)

	testCases := []struct {
		name string
		req  *csi.NodeUnstageVolumeRequest
		// fakeMounter mocks mounter behaviour
		fakeMounter *mount.FakeMounter
		// fakeHostUtil mocks hostutil behaviour
		fakeHostUtil *fakeHostUtil
		// fake filexattrs mock setchattr
		fakeFileXattrs *fakeFileXattrs
		// expected fake mount actions the test will make
		expActions []mount.FakeAction
		// expected test error code
		expErrCode codes.Code
		// expected mount points when test finishes
		expMountPoints []mount.MountPoint
	}{
		{
			name: "success normal",
			req: &csi.NodeUnstageVolumeRequest{
				StagingTargetPath: stagingPath,
				VolumeId:          volumeID,
			},
			fakeMounter: &mount.FakeMounter{
				MountPoints: []mount.MountPoint{{Device: fakeDiskPath, Path: stagingPath}},
			},
			fakeHostUtil: &fakeHostUtil{
				FakeHostUtil: hostutil.NewFakeHostUtil(map[string]hostutil.FileType{
					fakeDiskPath:    hostutil.FileTypeBlockDev,
					fakeDiskSymlink: hostutil.FileTypeFile,
					fakeStagingPath: hostutil.FileTypeDirectory,
				}),
				FakeDeviceSymlinks: map[string]string{fakeDiskSymlink: fakeDiskPath},
			},
			fakeFileXattrs: &fakeFileXattrs{},
		},
	}

	for _, tc := range testCases {
		fmt.Println("TC: ", tc)
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			driver := NewTestNodeService(tc.fakeMounter, tc.fakeHostUtil, tc.fakeFileXattrs, mdGetter, services.ProductionNodeCaps)
			r, err := driver.NodeUnstageVolume(context.TODO(), tc.req)
			require.NoError(t, err)
			fmt.Println("Response: ", r)
		})
	}
}
func TestNodePublishVolume(t *testing.T) {
	mountTargetPathPrefix := ""
	tmpDir, err := os.MkdirTemp("", "fakepublish")
	require.NoError(t, err)
	if runtime.GOOS == "darwin" {
		mountTargetPathPrefix = "/private"
	}
	defer os.RemoveAll(tmpDir)
	targetPath := filepath.Join(mountTargetPathPrefix, tmpDir)
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)

	testCases := []struct {
		name string
		req  *csi.NodePublishVolumeRequest
		// fakeMounter mocks mounter behaviour
		fakeMounter *mount.FakeMounter
		// fakeHostUtil mocks hostutil behaviour
		fakeHostUtil *fakeHostUtil
		// fake filexattrs mock setchattr
		fakeFileXattrs *fakeFileXattrs
		// expected fake mount actions the test will make
		expActions []mount.FakeAction
		// expected test error code
		expErrCode codes.Code
		// expected mount points when test finishes
		expMountPoints []mount.MountPoint
	}{
		{
			name: "success normal",
			req: &csi.NodePublishVolumeRequest{
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{
							FsType:     services.FsTypeExt4,
							MountFlags: nil,
						},
					},
					AccessMode: &csi.VolumeCapability_AccessMode{
						Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
					},
				},
				StagingTargetPath: fakeStagingPath,
				TargetPath:        targetPath,
				VolumeId:          volumeID,
			},
			fakeMounter:    mount.NewFakeMounter([]mount.MountPoint{}),
			fakeHostUtil:   &fakeHostUtil{},
			fakeFileXattrs: &fakeFileXattrs{},
			expActions: []mount.FakeAction{
				{
					Action: "mount",
					Target: targetPath,
					Source: fakeStagingPath,
					FSType: "",
				},
			},
			expErrCode:     codes.OK,
			expMountPoints: []mount.MountPoint{},
		},
		{
			name: "success already mounted",
			req: &csi.NodePublishVolumeRequest{
				VolumeCapability: &csi.VolumeCapability{
					AccessType: &csi.VolumeCapability_Mount{
						Mount: &csi.VolumeCapability_MountVolume{
							FsType:     services.FsTypeExt4,
							MountFlags: nil,
						},
					},
					AccessMode: &csi.VolumeCapability_AccessMode{
						Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
					},
				},
				StagingTargetPath: fakeStagingPath,
				TargetPath:        targetPath,
				VolumeId:          volumeID,
			},
			fakeMounter:    mount.NewFakeMounter([]mount.MountPoint{{Device: fakeDiskPath, Path: targetPath}}),
			fakeHostUtil:   &fakeHostUtil{},
			fakeFileXattrs: &fakeFileXattrs{},
			expActions:     []mount.FakeAction{},
			expErrCode:     codes.OK,
			expMountPoints: []mount.MountPoint{},
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			driver := NewTestNodeService(tc.fakeMounter, tc.fakeHostUtil, tc.fakeFileXattrs, mdGetter, services.ProductionNodeCaps)
			r, err := driver.NodePublishVolume(context.TODO(), tc.req)
			require.NoError(t, err)
			if !equalActions(tc.fakeMounter.GetLog(), tc.expActions) {
				t.Fatalf("Expected actions {%+v}, got {%+v}", tc.expActions, tc.fakeMounter.GetLog())
			}
			fmt.Println("Response: ", r)
		})
	}
}

func TestNodeUnpublishVolumeSuccess(t *testing.T) {
	mountTargetPathPrefix := ""
	tmpDir, err := os.MkdirTemp("", "fakepublish")
	require.NoError(t, err)
	if runtime.GOOS == "darwin" {
		mountTargetPathPrefix = "/private"
	}
	defer os.RemoveAll(tmpDir)
	targetPath := filepath.Join(mountTargetPathPrefix, tmpDir)
	fakeMounter := mount.NewFakeMounter([]mount.MountPoint{{Device: fakeDiskPath, Path: targetPath}})
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)
	driver := NewTestNodeService(fakeMounter, &fakeHostUtil{}, &fakeFileXattrs{}, mdGetter, services.ProductionNodeCaps)
	req := &csi.NodeUnpublishVolumeRequest{
		TargetPath: targetPath,
		VolumeId:   volumeID,
	}
	expActions := []mount.FakeAction{
		{
			Action: "unmount",
			Target: targetPath,
			Source: "",
			FSType: "",
		},
	}
	r, err := driver.NodeUnpublishVolume(context.TODO(), req)
	require.NoError(t, err)
	if !equalActions(fakeMounter.GetLog(), expActions) {
		t.Fatalf("Expected actions {%+v}, got {%+v}", expActions, fakeMounter.GetLog())
	}
	fmt.Println("Response: ", r)
}

func TestNodeUnpublishVolumeErrorNotEmptyDir(t *testing.T) {
	mountTargetPathPrefix := ""
	tmpDir, err := os.MkdirTemp("", "fakepublish")
	defer os.RemoveAll(tmpDir)
	require.NoError(t, err)
	if runtime.GOOS == "darwin" {
		mountTargetPathPrefix = "/private"
	}
	_, err = os.CreateTemp(tmpDir, "fakepublish")
	require.NoError(t, err)
	targetPath := filepath.Join(mountTargetPathPrefix, tmpDir)
	fakeMounter := mount.NewFakeMounter([]mount.MountPoint{{Device: fakeDiskPath, Path: targetPath}})
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)
	driver := NewTestNodeService(fakeMounter, &fakeHostUtil{}, &fakeFileXattrs{}, mdGetter, services.ProductionNodeCaps)
	req := &csi.NodeUnpublishVolumeRequest{
		TargetPath: targetPath,
		VolumeId:   volumeID,
	}
	expActions := []mount.FakeAction{
		{
			Action: "unmount",
			Target: targetPath,
			Source: "",
			FSType: "",
		},
	}
	r, err := driver.NodeUnpublishVolume(context.TODO(), req)
	require.Error(t, err)
	require.Condition(t, func() (success bool) {
		return errors.Is(err, syscall.ENOTEMPTY)
	})
	if !equalActions(fakeMounter.GetLog(), expActions) {
		t.Fatalf("Expected actions {%+v}, got {%+v}", expActions, fakeMounter.GetLog())
	}
	fmt.Println("Response: ", r)
}

func TestNodeGetInfo(t *testing.T) {
	mdGetter := NewFakeMetadataGetter(fakeInstID, fakeZoneID)
	driver := NewTestNodeService(&mount.FakeMounter{}, &fakeHostUtil{}, &fakeFileXattrs{}, mdGetter, services.ProductionNodeCaps)
	res, err := driver.NodeGetInfo(context.TODO(), &csi.NodeGetInfoRequest{})
	expected := &csi.NodeGetInfoResponse{
		NodeId:            fakeInstID,
		MaxVolumesPerNode: int64(DefaultOptions().MaxVolumesPerNode),
		AccessibleTopology: &csi.Topology{
			Segments: map[string]string{
				services.ZoneKey: fakeZoneID,
			},
		},
	}
	require.NoError(t, err)
	require.EqualExportedValues(t, expected, res)
	fmt.Println("resp: ", res)
}

func equalActions(tc, exp []mount.FakeAction) bool {
	if len(tc) == 0 && len(exp) == 0 {
		return true
	}
	return reflect.DeepEqual(tc, exp)
}

func equalMountPoints(tc, exp []mount.MountPoint) bool {
	if len(tc) == 0 && len(exp) == 0 {
		return true
	}
	return reflect.DeepEqual(tc, exp)
}

func setVolCap(fsType string) *csi.VolumeCapability {
	return &csi.VolumeCapability{
		AccessType: &csi.VolumeCapability_Mount{
			Mount: &csi.VolumeCapability_MountVolume{
				FsType: fsType,
			},
		},
		AccessMode: &csi.VolumeCapability_AccessMode{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		},
	}
}

func VolAccessTypeNotSet() *csi.VolumeCapability {
	return &csi.VolumeCapability{
		AccessMode: &csi.VolumeCapability_AccessMode{
			Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		},
	}
}

type fakeFileXattrs struct {
}

func (f fakeFileXattrs) SetImmutable(_ string) error {
	return nil
}

func (f fakeFileXattrs) UnsetImmutable(target string) error {
	return status.Errorf(codes.Internal, "chattr -i on %q failed", target)
}
