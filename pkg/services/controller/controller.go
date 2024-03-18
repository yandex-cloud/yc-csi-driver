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

package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"

	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"
	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/util"
)

type controller struct {
	diskAPI      diskapi.DiskAPI
	inFlight     inflight.InFlight
	capabilities []*csi.ControllerServiceCapability
}

func New(diskapi diskapi.DiskAPI, inFlight inflight.InFlight, caps []*csi.ControllerServiceCapability) csi.ControllerServer {
	return &controller{
		diskAPI:      diskapi,
		inFlight:     inFlight,
		capabilities: caps,
	}
}

func (c *controller) ControllerGetVolume(ctx context.Context, request *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	klog.Info("ControllerGetVolume is unimplemented")
	return nil, status.Error(codes.Unimplemented, "ControllerGetVolume is unimplemented")
}

func (c *controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.Infof("CreateVolume(%+v)", req)

	volName := req.GetName()
	if len(volName) == 0 {
		klog.Errorln("No volume name in request")
		return nil, status.Error(codes.InvalidArgument, "No volume name in request")
	}

	volCaps := req.GetVolumeCapabilities()
	if len(volCaps) == 0 {
		klog.Errorln("Volume capabilities not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities not provided")
	}

	err := services.VerifyVolumeCapabilities(volCaps, services.VolumeCaps)
	if err != nil {
		klog.Errorln(err)
		return nil, err
	}
	// zoneID may be empty; we'll check requirements in diskapi.
	zoneID := getAZ(req.GetAccessibilityRequirements())

	snapshotID := ""
	volumeSource := req.GetVolumeContentSource()
	if volumeSource != nil {
		if _, ok := volumeSource.GetType().(*csi.VolumeContentSource_Snapshot); !ok {
			return nil, status.Error(codes.InvalidArgument, "volumeContentSource is unsupported")
		}
		sourceSnapshot := volumeSource.GetSnapshot()
		if sourceSnapshot == nil {
			return nil, status.Error(codes.InvalidArgument, "Error retrieving snapshot from the volumeContentSource")
		}
		snapshotID = sourceSnapshot.GetSnapshotId()
	}

	disk, err := c.diskAPI.GetDiskByName(ctx, &diskapi.GetDiskByNameRequest{Name: services.CSIToYCName(volName)})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			klog.Infoln(err.Error())
			return c.createNewVolume(ctx, req, zoneID, snapshotID)
		case codes.Internal:
			klog.Errorf("Original error: %+v", err)
			return nil, err
		default:
			return nil, err
		}
	}

	if zoneID != "" && disk.ZoneID != zoneID {
		return nil, status.Errorf(codes.AlreadyExists, "Volume already exists, "+
			"but in different zone: (disk name: %s, disk zoneId: %s, topologyZoneId: %s)", req.GetName(), disk.ZoneID, zoneID)
	}

	if (req.GetCapacityRange().GetRequiredBytes() != 0 && disk.Size < req.GetCapacityRange().GetRequiredBytes()) ||
		(req.GetCapacityRange().GetLimitBytes() != 0 && disk.Size > req.GetCapacityRange().GetLimitBytes()) {
		klog.Errorf("volume %s already exists with size %d "+
			"incompatible with the new capacity requirements %+v", req.Name, disk.Size, req.GetCapacityRange())
		return nil, status.Errorf(codes.AlreadyExists, "volume %s already exists with size %d "+
			"incompatible with the new capacity requirements %+v", req.Name, disk.Size, req.GetCapacityRange())
	}

	if disk.SourceSnapshotID != snapshotID {
		return nil, status.Errorf(codes.AlreadyExists, "Volume already exists, "+
			"but created with another snapshot: (disk snapshotID: %s)", snapshotID)
	}

	if !disk.Ready {
		return nil, status.Errorf(codes.Aborted, "volume creation in progress; disk: %+v", disk)
	}

	return createVolumeResponse(disk)
}

func (c *controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.Infof("DeleteVolume(VolumeId=%s)", req.VolumeId)

	if req.GetVolumeId() == "" {
		return nil, status.Error(codes.InvalidArgument, "DeleteVolume requires VolumeId")
	}

	err := c.diskAPI.DeleteDisk(
		ctx,
		&diskapi.DeleteDiskRequest{
			DiskID: req.GetVolumeId(),
		})

	if err == nil || status.Code(err) == codes.NotFound {
		// https://github.com/container-storage-interface/spec/blob/master/spec.md#deletevolume
		return &csi.DeleteVolumeResponse{}, nil
	}

	klog.Errorf("Disk removal failed: %v", err)
	return nil, status.Error(status.Code(err), "Disk removal failed")
}

func (c *controller) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Infof("ControllerPublishVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	nodeID := req.GetNodeId()
	if len(nodeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Node ID not provided")
	}

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		klog.Errorln("Volume capability not provided")
		return nil, status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	volCaps := []*csi.VolumeCapability{volCap}
	err := services.VerifyVolumeCapabilities(volCaps, services.VolumeCaps)
	if err != nil {
		klog.Errorln(err)
		return nil, err
	}

	if c.inFlight.Get(util.StringToStringer(volumeID)) {
		klog.Infof("Volume=%s locked for resize", volumeID)
		return nil, status.Errorf(codes.FailedPrecondition, "Volume=%s locked for resize", volumeID)
	}

	readOnlyPublish := req.GetReadonly()
	attachReq := &diskapi.AttachDiskRequest{
		InstanceID: nodeID,
		DiskID:     volumeID,
		Readonly:   readOnlyPublish,
	}

	attached, err := c.getDiskAttachmentInfo(ctx, &getDiskAttachmentInfoRequest{
		InstanceID: attachReq.InstanceID,
		DiskID:     attachReq.DiskID,
	})

	if err != nil {
		return nil, err
	}

	err = IsCloudStateValidForDiskAttachRequest(attached, attachReq)
	if err != nil {
		return nil, err
	}

	// This check could be above IsCloudStateValidForDiskAttachRequest check, but had to be moved here, because of sanity tests
	// only allow volumes with VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY access mode to be attached as RO
	// also, forbid to attach volumes that have VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY access mode as RW
	// currently, we only support TWO volume access modes - see services.VolumeCaps for exact list.
	volumeAccessMode := req.GetVolumeCapability().GetAccessMode().GetMode()
	roVolumeAccessMode := volumeAccessMode == csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY

	// Request read only flag doesn't match that of volume access mode
	if readOnlyPublish != roVolumeAccessMode {
		return nil, status.Errorf(codes.FailedPrecondition,
			"Volume=%s access mode %q does not match public request ReadOnly flag: %t", volumeID,
			volumeAccessMode, readOnlyPublish)
	}

	err = c.diskAPI.EnsureObjectLabels(ctx, &diskapi.EnsureObjectLabelsRequest{
		Disk:     attached.Disk,
		Instance: attached.Instance,
	})

	if err != nil {
		return nil, err
	}

	if attached.Attached {
		return &csi.ControllerPublishVolumeResponse{
			PublishContext: services.ROPublishContext(readOnlyPublish),
		}, nil
	}

	err = c.diskAPI.AttachDisk(ctx, attachReq)
	if err != nil {
		return nil, err
	}

	return &csi.ControllerPublishVolumeResponse{
		PublishContext: services.ROPublishContext(readOnlyPublish),
	}, nil
}

func (c *controller) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.Infof("ControllerUnpublishVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	nodeID := req.GetNodeId()
	if len(nodeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Node ID not provided")
	}

	attached, err := c.getDiskAttachmentInfo(ctx, &getDiskAttachmentInfoRequest{
		InstanceID: nodeID,
		DiskID:     volumeID,
	})

	if err != nil {
		klog.Errorf("Disk (%s) unpublish operation failed on GetDiskAttachmentInfo (instanceID: %s ): %+v", volumeID, nodeID, err)
		return nil, err
	}

	// We don't really need to bother with labels here. keeping this code just in case.
	err = c.diskAPI.EnsureObjectLabels(ctx, &diskapi.EnsureObjectLabelsRequest{
		Disk:     attached.Disk,
		Instance: attached.Instance,
	})

	if err != nil {
		klog.Errorf("Disk (%s) unpublish operation failed on EnsureObjectLabels (instanceID: %s ): %+v", volumeID, nodeID, err)
		return nil, err
	}

	if !attached.Attached {
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	err = c.diskAPI.DetachDisk(
		ctx,
		&diskapi.DetachDiskRequest{
			InstanceID: nodeID,
			DiskID:     volumeID,
		},
	)

	if err != nil {
		klog.Errorf("Disk (%s) unpublish operation failed (instanceID: %s ): %+v", volumeID, nodeID, err)
		return nil, err
	}

	klog.V(5).Infof("ControllerUnpublishVolume: volume %s detached from instance %s", volumeID, nodeID)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (c *controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	klog.Infof("ValidateVolumeCapabilities(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities requires volumeId")
	}

	volCaps := req.GetVolumeCapabilities()
	if len(volCaps) == 0 {
		return nil, status.Error(codes.InvalidArgument, "ValidateVolumeCapabilities requires volumeCapabilities")
	}

	_, err := c.diskAPI.GetDisk(
		ctx,
		&diskapi.GetDiskRequest{
			ID: volumeID,
		},
	)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ValidateVolumeCapabilities failed to find volume %s", volumeID)
		}

		klog.Errorf("GetDiskRequest failed: %v", err)
		return nil, status.Errorf(codes.Internal, "ValidateVolumeCapabilities failed to get volume %s: %+v", volumeID, err)
	}

	// K8s doesn't seem to be using ValidateVolumeCapabilities.
	return &csi.ValidateVolumeCapabilitiesResponse{
		Message: "ValidateVolumeCapabilities is unimplemented",
	}, nil
}

func (*controller) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	klog.Info("ListVolumes is unimplemented")
	return nil, status.Error(codes.Unimplemented, "ListVolumes is unimplemented")
}

func (*controller) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	klog.Info("GetCapacity is unimplemented")
	return nil, status.Error(codes.Unimplemented, "GetCapacity is unimplemented")
}

func (c *controller) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	klog.Info("ControllerGetCapabilities")

	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: c.capabilities,
	}, nil
}

func (c *controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	klog.Infof("CreateSnapshot (%+v)", req)

	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Name is required for CreateSnapshotRequest")
	}

	if req.GetSourceVolumeId() == "" {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId is required for CreateSnapshotRequest")
	}

	snapshotYCName := services.CSIToYCName(req.GetName())

	snapshot, err := c.diskAPI.GetSnapshotByName(ctx,
		&diskapi.GetSnapshotByNameRequest{Name: snapshotYCName},
	)
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}

		klog.Infoln(err.Error())
		return c.createNewSnapshot(ctx, req)
	}

	if snapshot.SourceDiskID != req.GetSourceVolumeId() {
		msg := fmt.Sprintf(
			"snapshot with the requested name (%s) already exists but its sourceVolumeId (%s) "+
				"is different from the requested volumeId (%s)",
			snapshotYCName, snapshot.SourceDiskID, req.GetSourceVolumeId())
		klog.Errorf(msg)
		return nil, status.Errorf(codes.AlreadyExists, msg)
	}

	return createSnapshotResponse(snapshot)
}

func (c *controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	klog.Infof("DeleteSnapshot (%+v)", req)

	snapshotID := req.GetSnapshotId()
	if snapshotID == "" {
		return nil, status.Error(codes.InvalidArgument, "DeleteSnapshot requires SnapshotId")
	}

	err := c.diskAPI.DeleteSnapshot(
		ctx,
		&diskapi.DeleteSnapshotRequest{ID: snapshotID},
	)

	if err != nil {
		klog.Errorf("Error during snapshot ( %s ) removal: (%+v)", snapshotID, err)
		return nil, err
	}

	return &csi.DeleteSnapshotResponse{}, nil
}

func (c *controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	klog.Infof("ListSnapshots (%+v)", req)

	var (
		snapshot *diskapi.Snapshot
		err      error
	)

	if req.GetSnapshotId() != "" {
		snapshot, err = c.diskAPI.GetSnapshot(ctx, &diskapi.GetSnapshotRequest{ID: req.GetSnapshotId()})
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				return nil, fmt.Errorf("ListSnapshots: error getting status from request error (%+v)", err)
			}

			if st.Code() == codes.NotFound {
				klog.Infof("ListSnapshots: snapshot by id not found, returning empty result")
				return &csi.ListSnapshotsResponse{}, nil
			}

			klog.Errorf("ListSnapshots: error getting snapshot by id: %+v", err)
			return nil, err
		}

		klog.Infof("ListSnapshots: snapshot (%+v) by id (%s) found", snapshot, req.GetSnapshotId())

		return &csi.ListSnapshotsResponse{
			Entries: []*csi.ListSnapshotsResponse_Entry{
				{
					Snapshot: &csi.Snapshot{
						SizeBytes:      snapshot.Size,
						SnapshotId:     snapshot.ID,
						SourceVolumeId: snapshot.SourceDiskID,
						CreationTime:   snapshot.CreationTime,
						ReadyToUse:     snapshot.Ready,
					},
				},
			},
		}, nil
	}

	diskAPIRequest := &diskapi.ListSnapshotsRequest{}

	if req.GetMaxEntries() > 0 {
		klog.Infof("ListSnapshots: max_entries %d found", req.GetMaxEntries())
		diskAPIRequest.MaxEntries = req.GetMaxEntries()
	}

	if req.GetSourceVolumeId() != "" {
		klog.Infof("ListSnapshots: source_volume_id %s found", req.GetSourceVolumeId())

		_, err := c.diskAPI.GetDisk(ctx, &diskapi.GetDiskRequest{ID: req.GetSourceVolumeId()})
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				return nil, fmt.Errorf("ListSnapshots: error getting status from request error (%+v)", err)
			}

			if st.Code() == codes.NotFound {
				klog.Infof("ListSnapshots: source_volume_id %s not found, returning empty response", req.GetSourceVolumeId())
				return &csi.ListSnapshotsResponse{}, nil
			}

			klog.Errorf("ListSnapshots: error getting disk by source_volume_id: %s; error: %+v", req.GetSourceVolumeId(), err)
			return nil, err
		}

		diskAPIRequest.SourceVolumeID = req.GetSourceVolumeId()
	}

	diskAPIRequest.PageToken = req.GetStartingToken()

	resp, err := c.diskAPI.ListSnapshots(
		ctx,
		diskAPIRequest,
	)

	if err != nil {
		klog.Errorf("ListSnapshots: error listing snapshots with diskApiRequest (%+v); error: (%+v)", diskAPIRequest, err)
		return nil, err
	}

	var entries []*csi.ListSnapshotsResponse_Entry

	for _, s := range resp.Snapshots {
		entries = append(
			entries,
			&csi.ListSnapshotsResponse_Entry{
				Snapshot: &csi.Snapshot{
					SizeBytes:      s.Size,
					SnapshotId:     s.ID,
					SourceVolumeId: s.SourceDiskID,
					CreationTime:   s.CreationTime,
					ReadyToUse:     s.Ready,
				},
			},
		)
	}

	return &csi.ListSnapshotsResponse{
		Entries:   entries,
		NextToken: resp.NextPageToken,
	}, nil
}

func (c *controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	klog.Infof("ControllerExpandVolume(%+v)", req)

	volumeID := req.GetVolumeId()
	if len(volumeID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID is not provided")
	}

	limitBytes := req.GetCapacityRange().GetLimitBytes()
	requiredBytes := req.GetCapacityRange().GetRequiredBytes()

	validRequiredBytes := (0 <= requiredBytes) &&
		((requiredBytes <= limitBytes) || limitBytes == 0)
	validLimitBytes := 0 <= limitBytes

	if !validLimitBytes || !validRequiredBytes || (requiredBytes == 0 && limitBytes == 0) {
		return nil, status.Errorf(codes.OutOfRange, "Invalid capacity range ( limitBytes: %d ) , ( requiredBytes: %d )", limitBytes, requiredBytes)
	}

	disk, err := c.diskAPI.GetDisk(
		ctx,
		&diskapi.GetDiskRequest{ID: volumeID},
	)

	if err != nil {
		klog.Errorf("Cannot get disk with id ( %s ) : %+v", volumeID, err)
		return nil, status.Errorf(codes.NotFound, "Cannot get disk with id ( %s ) : %+v", volumeID, err)
	}

	if err = c.diskAPI.EnsureObjectLabels(ctx, &diskapi.EnsureObjectLabelsRequest{Disk: disk}); err != nil {
		return nil, err
	}

	if requiredBytes <= disk.Size {
		klog.Infof("Current disk size already meets the requirements; "+
			"diskID: %s, diskSize: %d, requestedSize: %d", volumeID, disk.Size, requiredBytes)
		c.inFlight.Delete(util.StringToStringer(volumeID))
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: disk.Size, NodeExpansionRequired: true}, nil
	}

	err = c.diskAPI.ExpandDisk(ctx, &diskapi.ExpandDiskRequest{ID: volumeID, Size: requiredBytes})
	if err != nil {
		if status.Code(err) == codes.FailedPrecondition {
			c.inFlight.Insert(util.StringToStringer(volumeID))
			klog.Infof("Set lock on Volume=%s", volumeID)
			return nil, status.Errorf(status.Code(err), "Error updating volume=%s, %+v", volumeID, err)
		}
		return nil, status.Errorf(codes.Internal, "Error updating volume=%s, %+v", volumeID, err)
	}
	klog.Infof("Removed lock from Volume=%s", volumeID)
	c.inFlight.Delete(util.StringToStringer(volumeID))

	return &csi.ControllerExpandVolumeResponse{CapacityBytes: requiredBytes, NodeExpansionRequired: true}, nil
}

func (c *controller) createNewVolume(
	ctx context.Context,
	req *csi.CreateVolumeRequest,
	zoneID, snapshotID string,
) (*csi.CreateVolumeResponse, error) {
	diskAPIRequest := &diskapi.CreateDiskRequest{
		CSIVolumeName: req.GetName(),
		Name:          services.CSIToYCName(req.GetName()),
		ZoneID:        zoneID,
	}

	if snapshotID != "" {
		_, err := c.diskAPI.GetSnapshot(ctx, &diskapi.GetSnapshotRequest{ID: snapshotID})
		if err != nil {
			klog.Errorf("Error (%+v) getting snapshot by snapshot id: %s", err, snapshotID)
			return nil, err
		}

		diskAPIRequest.SnapshotID = snapshotID
	}

	klog.Infof("Creating new volume %+v", req)

	parameters := req.GetParameters()

	volumeTypeID := util.GetOrDefaultStringFromMap(parameters, services.VolumeTypeKey, services.DefaultVolumeType)
	if volumeTypeID == "network-nvme" {
		volumeTypeID = "network-ssd"
	}

	diskAPIRequest.TypeID = volumeTypeID

	size, err := getVolumeSize(req.GetCapacityRange())
	if err != nil {
		return nil, err
	}

	diskAPIRequest.Size = size

	diskPlacementPolicy, err := getDiskPlacementPolicy(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to get DiskPlacementPolicy from parameters: %w", err)
	}

	diskAPIRequest.DiskPlacementPolicy = diskPlacementPolicy

	disk, err := c.diskAPI.CreateDisk(
		ctx,
		diskAPIRequest,
	)

	if err != nil {
		klog.Errorf("Disk creation (%+v) failed: %v", req, err)
		return nil, status.Error(status.Code(err), fmt.Sprintf("Disk creation (%+v) failed: %s", req, err.Error()))
	}

	return createVolumeResponse(disk)
}

func getDiskPlacementPolicy(parameters map[string]string) (*diskapi.DiskPlacementPolicy, error) {
	if parameters == nil {
		return nil, nil
	}

	var result *diskapi.DiskPlacementPolicy

	if groupID, ok := parameters[services.DiskPlacementGroupIDKey]; ok {
		result = &diskapi.DiskPlacementPolicy{
			PlacementGroupID: groupID,
		}
	}

	if groupPartitionStr, ok := parameters[services.DiskPlacementGroupPartitionKey]; ok {
		groupPartition, err := strconv.Atoi(groupPartitionStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q parameter to integer value: %w", services.DiskPlacementGroupPartitionKey, err)
		}

		if result == nil {
			result = &diskapi.DiskPlacementPolicy{}
		}

		result.PlacementGroupPartition = int64(groupPartition)
	}

	return result, nil
}

func (c *controller) createNewSnapshot(
	ctx context.Context,
	req *csi.CreateSnapshotRequest,
) (*csi.CreateSnapshotResponse, error) {
	snapshotYCName := services.CSIToYCName(req.GetName())

	snapshot, err := c.diskAPI.CreateSnapshot(ctx, &diskapi.CreateSnapshotRequest{
		Name:            snapshotYCName,
		CSISnapshotName: req.GetName(),
		SourceDiskID:    req.SourceVolumeId,
	})

	if err != nil {
		msg := fmt.Sprintf("CreateSnapshot ( name: %s ) returned error ( %+v )", req.GetName(), err)
		klog.Errorf(msg)
		return nil, status.Error(status.Code(err), msg)
	}

	return createSnapshotResponse(snapshot)
}

type getDiskAttachmentInfoRequest struct {
	InstanceID string
	DiskID     string
}

type diskAttachmentInfo struct {
	Attached bool
	Readonly bool
	Disk     *diskapi.Disk
	Instance *diskapi.Instance
}

func (c *controller) getDiskAttachmentInfo(
	ctx context.Context,
	req *getDiskAttachmentInfoRequest,
) (*diskAttachmentInfo, error) {
	disk, err := c.diskAPI.GetDisk(ctx, &diskapi.GetDiskRequest{ID: req.DiskID})
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	instance, err := c.diskAPI.GetInstance(ctx, &diskapi.GetInstanceRequest{InstanceID: req.InstanceID})
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	info := &diskAttachmentInfo{
		Disk:     disk,
		Instance: instance,
	}

	// Either disk or instance do no longer exist.
	if disk == nil || instance == nil {
		return info, nil
	}

	// Both disk and instance exists. Let's check, their attachment to each other.
	diskHasInstance := arrayOfStringsContains(disk.InstanceIDs, req.InstanceID)

	instanceHasDisk := false
	var readOnly bool

	for _, d := range instance.DiskAttachments {
		if d.DiskID == req.DiskID {
			instanceHasDisk = true
			readOnly = d.Readonly
			break
		}
	}

	// Here, we have some sort of invalid state, between disk and instance.
	// let's hope, the attach (or detach) operation is in progress, and we shall retry it.
	if diskHasInstance != instanceHasDisk {
		// instanceHasDisk == false, using logical not, for better code reading here.
		if !instanceHasDisk {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf(
				"disk %q is in invalid attachment state to instance %q, instance does not have disk id as its boot or secondary disk",
				req.DiskID, req.InstanceID))
		}

		// Here, diskHasInstance == false
		return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf(
			"disk %q is in invalid attachment state to instance %q, disk does not have instance id in its InstanceIds array",
			req.DiskID, req.InstanceID))
	}

	// Either both flags are true => disk is attached - ok, or both are false, => disk isn't attached - ok.
	info.Attached = diskHasInstance
	info.Readonly = readOnly
	return info, nil
}

func IsCloudStateValidForDiskAttachRequest(attached *diskAttachmentInfo, req *diskapi.AttachDiskRequest) error {
	// Trying to attach a disk, that doesn't exist.
	if attached.Disk == nil {
		return status.Error(codes.NotFound, fmt.Sprintf(
			"disk: %q does not exist", req.DiskID))
	}

	// Trying to attach disk to an instance, that doesn't exist.
	if attached.Instance == nil {
		return status.Error(codes.NotFound, fmt.Sprintf(
			"instance %q does not exist", req.InstanceID))
	}

	// Check, that disk is attached in the same mode (readonly, or not), to instance, as in this request.
	if attached.Attached && attached.Readonly != req.Readonly {
		return status.Errorf(codes.AlreadyExists, "cannot attach disk %q to instance %q in different mode. "+
			"disk is attached in RO mode: %t, attach requested in RO mode: %t",
			req.DiskID, req.InstanceID, attached.Readonly, req.Readonly)
	}

	return nil
}

func arrayOfStringsContains(arrayOfStings []string, s string) bool {
	for _, el := range arrayOfStings {
		if el == s {
			return true
		}
	}
	return false
}

func createSnapshotResponse(snapshot *diskapi.Snapshot) (*csi.CreateSnapshotResponse, error) {
	return &csi.CreateSnapshotResponse{
		Snapshot: &csi.Snapshot{
			SizeBytes:      snapshot.Size,
			SnapshotId:     snapshot.ID,
			SourceVolumeId: snapshot.SourceDiskID,
			CreationTime:   snapshot.CreationTime,
			ReadyToUse:     snapshot.Ready,
		},
	}, nil
}

func getVolumeSize(capacityRange *csi.CapacityRange) (int64, error) {
	if capacityRange == nil {
		// csi-sanity fails if CreateVolume with unspecified CapacityRange fails; also see comment for the parameter at
		// https://github.com/container-storage-interface/spec/blob/master/spec.md#createvolume
		return services.MinDiskSizeBytes, nil
	}

	size := capacityRange.GetRequiredBytes()

	if capacityRange.GetRequiredBytes() < services.MinDiskSizeBytes {
		if services.MinDiskSizeBytes > capacityRange.GetLimitBytes() {
			// https://github.com/container-storage-interface/spec/blob/master/spec.md#createvolume-errors
			return 0, status.Error(codes.OutOfRange,
				fmt.Sprintf("disk minimum size is %d bytes", services.MinDiskSizeBytes))
		}

		size = services.MinDiskSizeBytes
	}

	return size, nil
}

func createVolumeResponse(disk *diskapi.Disk) (*csi.CreateVolumeResponse, error) {
	klog.Infof("Got disk; volumeID: %+s; Name: %+s", disk.ID, disk.Name)

	var src *csi.VolumeContentSource
	if disk.SourceSnapshotID != "" {
		src = &csi.VolumeContentSource{
			Type: &csi.VolumeContentSource_Snapshot{
				Snapshot: &csi.VolumeContentSource_SnapshotSource{
					SnapshotId: disk.SourceSnapshotID,
				},
			},
		}
	}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      disk.ID,
			CapacityBytes: disk.Size,
			AccessibleTopology: []*csi.Topology{
				{Segments: map[string]string{services.ZoneKey: disk.ZoneID}},
			},
			ContentSource: src,
			VolumeContext: map[string]string{services.VolumeTypeKey: disk.TypeID},
		}}, nil
}

func getAZ(requirement *csi.TopologyRequirement) string {
	if requirement == nil {
		klog.Errorf("can't provision yc disks without zone information")
		return ""
	}

	for _, topology := range requirement.GetPreferred() {
		klog.Infof("preferred topology requirement: %+v", topology)

		zone, exists := topology.GetSegments()[services.ZoneKey]
		if exists {
			return zone
		}
	}

	for _, topology := range requirement.GetRequisite() {
		klog.Infof("requisite topology requirement: %+v", topology)

		zone, exists := topology.GetSegments()[services.ZoneKey]
		if exists {
			return zone
		}
	}

	return ""
}
