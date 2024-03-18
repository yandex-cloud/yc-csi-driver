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
	"testing"

	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"
	diskapimocks "github.com/yandex-cloud/yc-csi-driver/pkg/diskapi/mocks"
	"github.com/yandex-cloud/yc-csi-driver/pkg/inflight"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
	"github.com/yandex-cloud/yc-csi-driver/pkg/util"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	diskName     = "NewVolume"
	diskID       = "test-disk-id"
	instanceID   = "test-instance-id"
	instanceID2  = "test-instance-id-2"
	zoneID       = "ru-central1-a"
	diskTypeID   = "mytype"
	snapshotName = "NewSnapshotName"
	snapshotID   = "NewSnapshotID"
)

var (
	accessMode = &csi.VolumeCapability_AccessMode{
		Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
	}
	wrongAccessMode = &csi.VolumeCapability_AccessMode{
		Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	}
	volumeCapabilitiesWithBlockType = &csi.VolumeCapability{
		AccessType: &csi.VolumeCapability_Block{
			Block: &csi.VolumeCapability_BlockVolume{},
		},
		AccessMode: accessMode,
	}
	volumeCapabilitiesWithFileSystemType = &csi.VolumeCapability{
		AccessType: &csi.VolumeCapability_Mount{
			Mount: &csi.VolumeCapability_MountVolume{
				FsType: services.DefaultFsType,
			},
		},
		AccessMode: accessMode,
	}
	volumeCapabilitiesWithWrongAccessMode = &csi.VolumeCapability{
		AccessType: &csi.VolumeCapability_Mount{
			Mount: &csi.VolumeCapability_MountVolume{
				FsType: services.DefaultFsType,
			},
		},
		AccessMode: wrongAccessMode,
	}

	newDisk = &diskapi.Disk{
		ID:          diskID,
		Size:        services.DefaultDiskSizeBytes,
		ZoneID:      zoneID,
		Name:        diskName,
		TypeID:      diskTypeID,
		InstanceIDs: []string{},
		Ready:       true,
	}

	newSSDDisk = &diskapi.Disk{
		ID:          "test-disk-id-2",
		Size:        services.DefaultDiskSizeBytes,
		ZoneID:      zoneID,
		Name:        "network-nvme",
		TypeID:      "network-nvme",
		InstanceIDs: []string{},
		Ready:       true,
	}

	newNVMEWithSSDTypeDisk = &diskapi.Disk{
		ID:          "test-disk-id-2",
		Size:        services.DefaultDiskSizeBytes,
		ZoneID:      zoneID,
		Name:        "network-nvme",
		TypeID:      "network-ssd",
		InstanceIDs: []string{},
		Ready:       true,
	}
)

type Mocks struct {
	method     string
	returnData interface{}
	err        error
}

func TestController_CreateVolume(t *testing.T) {
	testCases := []struct {
		name       string
		req        *csi.CreateVolumeRequest
		resp       *csi.CreateVolumeResponse
		expErrCode codes.Code
		mocks      []*Mocks
	}{
		{
			name:       "NoVolumeName",
			req:        &csi.CreateVolumeRequest{},
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "NoVolumeCapabilities",
			req: &csi.CreateVolumeRequest{
				Name: diskName,
			},
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "WrongAccessMode",
			req: &csi.CreateVolumeRequest{
				Name:               diskName,
				VolumeCapabilities: []*csi.VolumeCapability{volumeCapabilitiesWithWrongAccessMode},
			},
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "NoAZ",
			req: &csi.CreateVolumeRequest{
				Name:                      diskName,
				VolumeCapabilities:        []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType},
				AccessibilityRequirements: &csi.TopologyRequirement{},
			},
			expErrCode: codes.InvalidArgument,
			mocks: []*Mocks{
				{
					method:     "CreateDisk",
					returnData: nil,
					err:        status.Error(codes.InvalidArgument, "AZ required"),
				},
				{
					method:     "GetDiskByName",
					returnData: nil,
					err:        status.Error(codes.NotFound, "NotFound"),
				},
			},
		},
		{
			name: "Success",
			req:  newVolumeRequest(newDisk, newDisk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
			resp: &csi.CreateVolumeResponse{
				Volume: &csi.Volume{
					CapacityBytes:      newDisk.Size,
					VolumeId:           newDisk.ID,
					AccessibleTopology: []*csi.Topology{{Segments: map[string]string{services.ZoneKey: newDisk.ZoneID}}},
					VolumeContext:      map[string]string{services.VolumeTypeKey: newDisk.TypeID},
				},
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method:     "CreateDisk",
					returnData: newDisk,
					err:        nil,
				},
				{
					method:     "GetDiskByName",
					returnData: nil,
					err:        status.Error(codes.NotFound, "NotFound"),
				},
			},
		},
		{
			name: "ExistingVolume",
			req:  newVolumeRequest(newDisk, newDisk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
			resp: &csi.CreateVolumeResponse{
				Volume: &csi.Volume{
					VolumeId:      newDisk.ID,
					CapacityBytes: newDisk.Size,
					AccessibleTopology: []*csi.Topology{
						{Segments: map[string]string{services.ZoneKey: newDisk.ZoneID}},
					},
					VolumeContext: map[string]string{services.VolumeTypeKey: newDisk.TypeID},
				}},
			mocks: []*Mocks{
				{
					method:     "CreateDisk",
					returnData: newDisk,
					err:        status.Error(codes.Internal, "Create error"),
				},
				{
					method:     "GetDiskByName",
					returnData: newDisk,
					err:        nil,
				},
			},
		},
		{
			name:       "ExistingVolumeNotReadyStatus",
			req:        newVolumeRequest(newDisk, newDisk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
			resp:       nil,
			expErrCode: codes.Aborted,
			mocks: []*Mocks{
				{
					method: "GetDiskByName",
					returnData: func() *diskapi.Disk {
						disk := &diskapi.Disk{}
						*disk = *newDisk
						disk.Ready = false
						return disk
					}(),
					err: nil,
				},
			},
		},
		{
			name: "network-nvme",
			req:  newVolumeRequest(newSSDDisk, newSSDDisk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
			resp: &csi.CreateVolumeResponse{
				Volume: &csi.Volume{
					CapacityBytes:      newSSDDisk.Size,
					VolumeId:           newSSDDisk.ID,
					AccessibleTopology: []*csi.Topology{{Segments: map[string]string{services.ZoneKey: newSSDDisk.ZoneID}}},
					VolumeContext:      map[string]string{services.VolumeTypeKey: "network-ssd"},
				},
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method:     "CreateDisk",
					returnData: newNVMEWithSSDTypeDisk,
					err:        nil,
				},
				{
					method:     "GetDiskByName",
					returnData: nil,
					err:        status.Error(codes.NotFound, "NotFound"),
				},
			},
		},
		{
			name: "SuccessBlock",
			req:  newVolumeRequest(newDisk, newDisk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithBlockType}),
			resp: &csi.CreateVolumeResponse{
				Volume: &csi.Volume{
					CapacityBytes:      newDisk.Size,
					VolumeId:           newDisk.ID,
					AccessibleTopology: []*csi.Topology{{Segments: map[string]string{services.ZoneKey: newDisk.ZoneID}}},
					VolumeContext:      map[string]string{services.VolumeTypeKey: newDisk.TypeID},
				},
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method:     "CreateDisk",
					returnData: newDisk,
					err:        nil,
				},
				{
					method:     "GetDiskByName",
					returnData: nil,
					err:        status.Error(codes.NotFound, "NotFound"),
				},
			},
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)

		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			d := &diskapimocks.DiskAPI{}
			i := inflight.NewWithTTL()

			c := New(d, i, services.ProductionCapsForPublicAPIController)
			for _, m := range tc.mocks {
				//if m.returnData != nil {
				d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.returnData, m.err)
				//} else {
				//	d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.err)
				//}
			}
			resp, err := c.CreateVolume(context.TODO(), tc.req)
			checkError(t, err, tc.expErrCode)
			require.Equal(t, tc.resp, resp)
		})
	}
}

func TestDeleteVolume(t *testing.T) {
	d := &diskapimocks.DiskAPI{}
	i := inflight.NewWithTTL()
	c := New(d, i, services.ProductionCapsForPublicAPIController)

	d.Mock.On("DeleteDisk", mock.Anything, mock.Anything).Return(status.Error(codes.NotFound, "disk not found"))
	// delete nonexistent disk should return no error
	_, err := c.DeleteVolume(
		context.Background(),
		&csi.DeleteVolumeRequest{VolumeId: newDisk.ID},
	)
	require.NoError(t, err)

	d.Mock.On("DeleteDisk", mock.Anything, mock.Anything).Return(nil).Once()
	disk := newDisk
	_, err = c.DeleteVolume(
		context.Background(),
		&csi.DeleteVolumeRequest{VolumeId: disk.ID},
	)
	require.NoError(t, err)
}

func TestPublishVolume(t *testing.T) {
	d := &diskapimocks.DiskAPI{}
	i := inflight.NewWithTTL()
	c := New(d, i, services.ProductionCapsForPublicAPIController)
	d.Mock.On("GetDiskByName", mock.Anything, mock.Anything).Return(
		nil, status.Error(codes.NotFound, "NotFound"))

	disk := &diskapi.Disk{}
	*disk = *newDisk
	instance := &diskapi.Instance{ID: instanceID}

	d.Mock.On("CreateDisk", mock.Anything, mock.Anything).Return(disk, nil).Once()
	d.Mock.On("GetDisk", mock.Anything, mock.Anything).Return(disk, nil)
	d.Mock.On("GetInstance", mock.Anything, mock.Anything).Return(instance, nil)
	d.Mock.On("EnsureObjectLabels", mock.Anything, mock.Anything).Return(nil)

	volume, err := c.CreateVolume(
		context.Background(),
		newVolumeRequest(disk, disk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
	)
	require.NoError(t, err)
	assert.Equal(t, volume.GetVolume().GetVolumeId(), disk.ID)

	d.Mock.On("AttachDisk", mock.Anything, mock.Anything).Return(nil).Once()
	_, err = c.ControllerPublishVolume(
		context.Background(),
		publishVolumeRequest(disk, instanceID, volumeCapabilitiesWithFileSystemType),
	)
	require.NoError(t, err)
	disk.InstanceIDs = append(disk.InstanceIDs, instanceID)
	instance.DiskAttachments = append(instance.DiskAttachments, diskapi.DiskAttachment{
		DiskID:   disk.ID,
		Readonly: false,
	})

	_, err = c.ControllerPublishVolume(
		context.Background(),
		publishVolumeRequest(disk, instanceID2, volumeCapabilitiesWithFileSystemType),
	)

	require.Error(t, err)
	assert.Equal(t, codes.FailedPrecondition, status.Code(err))
}

func TestPublishUpdatingVolume(t *testing.T) {
	d := &diskapimocks.DiskAPI{}
	i := inflight.NewWithTTL()
	c := New(d, i, services.ProductionCapsForPublicAPIController)

	disk := &diskapi.Disk{}
	*disk = *newDisk
	i.Insert(util.StringToStringer(disk.ID))
	_, err := c.ControllerPublishVolume(
		context.Background(),
		publishVolumeRequest(disk, instanceID, volumeCapabilitiesWithFileSystemType),
	)

	require.Error(t, err)
	require.Equal(t, codes.FailedPrecondition, status.Code(err))
}

func TestPublishBlockVolume(t *testing.T) {
	d := &diskapimocks.DiskAPI{}
	i := inflight.NewWithTTL()
	c := New(d, i, services.ProductionCapsForPublicAPIController)

	d.Mock.On("GetDiskByName", mock.Anything, mock.Anything).Return(
		nil, status.Error(codes.NotFound, "NotFound"))

	disk := newDisk
	d.Mock.On("CreateDisk", mock.Anything, mock.Anything).Return(disk, nil)

	instance := &diskapi.Instance{ID: instanceID}
	d.Mock.On("GetDisk", mock.Anything, mock.Anything).Return(disk, nil)
	d.Mock.On("GetInstance", mock.Anything, mock.Anything).Return(instance, nil)
	d.Mock.On("EnsureObjectLabels", mock.Anything, mock.Anything).Return(nil)

	volume, err := c.CreateVolume(
		context.Background(),
		newVolumeRequest(disk, disk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
	)

	require.NoError(t, err)
	assert.Equal(t, volume.GetVolume().GetVolumeId(), disk.ID)

	d.Mock.On("AttachDisk", mock.Anything, mock.Anything).Return(
		status.Error(codes.InvalidArgument, "InvalidArgument"),
	).Once()
	_, err = c.ControllerPublishVolume(
		context.Background(),
		publishVolumeRequest(disk, instanceID2, volumeCapabilitiesWithBlockType),
	)

	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func TestUnpublishVolume(t *testing.T) {
	d := &diskapimocks.DiskAPI{}
	i := inflight.NewWithTTL()
	c := New(d, i, services.ProductionCapsForPublicAPIController)

	disk := &diskapi.Disk{}
	*disk = *newDisk

	d.Mock.On("CreateDisk", mock.Anything, mock.Anything).Return(disk, nil)
	d.Mock.On("GetDiskByName", mock.Anything, mock.Anything).Return(
		nil, status.Error(codes.NotFound, "NotFound"))

	instance := &diskapi.Instance{ID: instanceID}
	d.Mock.On("GetDisk", mock.Anything, mock.Anything).Return(disk, nil)
	d.Mock.On("GetInstance", mock.Anything, mock.Anything).Return(instance, nil)
	d.Mock.On("EnsureObjectLabels", mock.Anything, mock.Anything).Return(nil)

	volume, err := c.CreateVolume(
		context.Background(),
		newVolumeRequest(disk, disk.Name, []*csi.VolumeCapability{volumeCapabilitiesWithFileSystemType}),
	)
	require.NoError(t, err)
	assert.Equal(t, volume.GetVolume().GetVolumeId(), disk.ID)

	d.Mock.On("AttachDisk", mock.Anything, mock.Anything).Return(nil).Once()
	_, err = c.ControllerPublishVolume(
		context.Background(),
		publishVolumeRequest(disk, instanceID, volumeCapabilitiesWithFileSystemType),
	)
	require.NoError(t, err)
	disk.InstanceIDs = append(disk.InstanceIDs, instanceID)
	instance.DiskAttachments = append(instance.DiskAttachments, diskapi.DiskAttachment{
		DiskID:   disk.ID,
		Readonly: false,
	})

	d.Mock.On("DetachDisk", mock.Anything, mock.Anything).Return(nil)
	_, err = c.ControllerUnpublishVolume(
		context.Background(),
		&csi.ControllerUnpublishVolumeRequest{
			VolumeId: volume.GetVolume().GetVolumeId(),
			NodeId:   disk.InstanceIDs[0],
		},
	)
	require.NoError(t, err)

	//unpublish on non-attached disk should return no error
	_, err = c.ControllerUnpublishVolume(
		context.Background(),
		&csi.ControllerUnpublishVolumeRequest{
			VolumeId: volume.GetVolume().GetVolumeId(),
			NodeId:   disk.InstanceIDs[0],
		},
	)
	require.NoError(t, err)

	d.Mock.On("GetDisk", mock.Anything, mock.Anything).Return(nil, status.Error(codes.NotFound, ""))
	//unpublish on non-existent disk should return no error
	_, err = c.ControllerUnpublishVolume(
		context.Background(),
		&csi.ControllerUnpublishVolumeRequest{
			VolumeId: volume.GetVolume().GetVolumeId(),
			NodeId:   disk.InstanceIDs[0],
		},
	)
	require.NoError(t, err)
}

func TestController_ControllerExpandVolume(t *testing.T) {
	testCases := []struct {
		name       string
		diskSize   int64
		req        *csi.ControllerExpandVolumeRequest
		resp       *csi.ControllerExpandVolumeResponse
		expErrCode codes.Code
		mocks      []*Mocks
	}{
		{
			name:       "Empty request",
			diskSize:   services.DefaultDiskSizeBytes,
			req:        &csi.ControllerExpandVolumeRequest{},
			resp:       nil,
			expErrCode: codes.InvalidArgument,
		},
		{
			name:       "Unspecified *csi.CapacityRange",
			diskSize:   services.DefaultDiskSizeBytes,
			req:        &csi.ControllerExpandVolumeRequest{VolumeId: diskID},
			resp:       nil,
			expErrCode: codes.OutOfRange,
		},
		{
			name:     "Unspecified LimitBytes/RequiredBytes",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					LimitBytes:    0,
					RequiredBytes: 0,
				}},
			resp:       nil,
			expErrCode: codes.OutOfRange,
		},
		{
			name:     "Negative LimitBytes property",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId:      diskID,
				CapacityRange: &csi.CapacityRange{LimitBytes: -1}},
			resp:       nil,
			expErrCode: codes.OutOfRange,
		},
		{
			name:     "RequiredBytes greater than LimitBytes",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					LimitBytes:    100000,
					RequiredBytes: 100001,
				},
			},
			resp:       nil,
			expErrCode: codes.OutOfRange,
		},
		{
			name:     "LimitBytes greater than MaxDiskSize",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: services.DefaultDiskSizeBytes,
				}},
			resp: &csi.ControllerExpandVolumeResponse{
				CapacityBytes:         services.DefaultDiskSizeBytes,
				NodeExpansionRequired: true,
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method: "GetDisk",
					returnData: &diskapi.Disk{
						ID:          diskID,
						Size:        services.DefaultDiskSizeBytes,
						InstanceIDs: nil,
					},
					err: nil,
				}},
		},
		{
			name:     "RequiredBytes Set But LimitBytes is 0",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId:      diskID,
				CapacityRange: &csi.CapacityRange{RequiredBytes: services.DefaultDiskSizeBytes, LimitBytes: 0}},
			resp: &csi.ControllerExpandVolumeResponse{
				CapacityBytes:         services.DefaultDiskSizeBytes,
				NodeExpansionRequired: true,
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method: "GetDisk",
					returnData: &diskapi.Disk{
						ID:          diskID,
						Size:        services.DefaultDiskSizeBytes,
						InstanceIDs: nil,
					},
					err: nil,
				}},
		},
		{
			name:     "DiskSize meets requirements",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: services.DefaultDiskSizeBytes,
					LimitBytes:    services.DefaultDiskSizeBytes,
				}},
			resp: &csi.ControllerExpandVolumeResponse{
				CapacityBytes:         services.DefaultDiskSizeBytes,
				NodeExpansionRequired: true,
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method: "GetDisk",
					returnData: &diskapi.Disk{
						ID:          diskID,
						Size:        services.DefaultDiskSizeBytes,
						InstanceIDs: nil,
					},
					err: nil,
				}},
		},
		{
			name:     "Failed Precondition",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 2 * services.DefaultDiskSizeBytes,
					LimitBytes:    2 * services.DefaultDiskSizeBytes,
				}},
			resp:       nil,
			expErrCode: codes.FailedPrecondition,
			mocks: []*Mocks{
				{
					method: "GetDisk",
					returnData: &diskapi.Disk{
						ID:          diskID,
						Size:        services.DefaultDiskSizeBytes,
						InstanceIDs: nil,
					},
					err: nil,
				},
				{
					method:     "ExpandDisk",
					returnData: nil,
					err:        status.Errorf(codes.FailedPrecondition, "Disk %s sattached", diskID),
				},
			},
		},
		{
			name:     "Updated successfully",
			diskSize: services.DefaultDiskSizeBytes,
			req: &csi.ControllerExpandVolumeRequest{
				VolumeId: diskID,
				CapacityRange: &csi.CapacityRange{
					RequiredBytes: 2 * services.DefaultDiskSizeBytes,
					LimitBytes:    2 * services.DefaultDiskSizeBytes,
				}},
			resp: &csi.ControllerExpandVolumeResponse{
				CapacityBytes:         2 * services.DefaultDiskSizeBytes,
				NodeExpansionRequired: true,
			},
			expErrCode: codes.OK,
			mocks: []*Mocks{
				{
					method: "GetDisk",
					returnData: &diskapi.Disk{
						ID:          diskID,
						Size:        services.DefaultDiskSizeBytes,
						InstanceIDs: nil,
					},
					err: nil,
				},
				{
					method:     "ExpandDisk",
					returnData: nil,
					err:        nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		fmt.Println("TC: ", tc)

		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			d := &diskapimocks.DiskAPI{}
			i := inflight.NewWithTTL()
			for _, m := range tc.mocks {
				if m.returnData != nil {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.returnData, m.err)
				} else {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.err)
				}
			}

			d.Mock.On("EnsureObjectLabels", mock.Anything, mock.Anything).Return(nil)
			c := New(d, i, services.ProductionCapsForPublicAPIController)
			resp, err := c.ControllerExpandVolume(context.TODO(), tc.req)
			checkError(t, err, tc.expErrCode)
			require.Equal(t, tc.resp, resp)
		})
	}
}

func TestController_CreateSnapshot(t *testing.T) {
	creationTime := timestamppb.Now()
	testCases := []struct {
		name          string
		req           *csi.CreateSnapshotRequest
		resp          *csi.CreateSnapshotResponse
		expErrCode    codes.Code
		mocks         []*Mocks
		inFlightValue inflight.Idempotent
	}{
		{
			name:       "No snapshotName",
			req:        &csi.CreateSnapshotRequest{},
			resp:       nil,
			expErrCode: codes.InvalidArgument,
		},
		{
			name:       "No volumeID",
			req:        &csi.CreateSnapshotRequest{Name: snapshotName},
			resp:       nil,
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "ExistingSnapshot not matching volume id",
			req: &csi.CreateSnapshotRequest{
				SourceVolumeId: diskID,
				Name:           snapshotName,
			},
			resp:       nil,
			expErrCode: codes.AlreadyExists,
			mocks: []*Mocks{
				{
					method: "GetSnapshotByName",
					returnData: &diskapi.Snapshot{
						Name:         snapshotName,
						ID:           snapshotID,
						Size:         services.DefaultDiskSizeBytes,
						SourceDiskID: "test",
						Ready:        true,
						CreationTime: creationTime,
					},
					err: nil,
				},
			},
		},
		{
			name: "ExistingSnapshot not ready",
			req: &csi.CreateSnapshotRequest{
				SourceVolumeId: diskID,
				Name:           snapshotName,
			},
			resp: &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SizeBytes:      services.DefaultDiskSizeBytes,
					SnapshotId:     snapshotID,
					SourceVolumeId: diskID,
					CreationTime:   creationTime,
					ReadyToUse:     false,
				},
			},
			mocks: []*Mocks{
				{
					method: "GetSnapshotByName",
					returnData: &diskapi.Snapshot{
						Name:         services.CSIToYCName(snapshotName),
						ID:           snapshotID,
						Size:         services.DefaultDiskSizeBytes,
						SourceDiskID: diskID,
						Ready:        false,
						CreationTime: creationTime,
					},
					err: nil,
				},
			},
		},
		{
			name: "Existing Snapshot success",
			req: &csi.CreateSnapshotRequest{
				SourceVolumeId: diskID,
				Name:           snapshotName,
			},
			resp: &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SizeBytes:      services.DefaultDiskSizeBytes,
					SnapshotId:     snapshotID,
					SourceVolumeId: diskID,
					CreationTime:   creationTime,
					ReadyToUse:     true,
				},
			},
			mocks: []*Mocks{
				{
					method: "GetSnapshotByName",
					returnData: &diskapi.Snapshot{
						Name:         services.CSIToYCName(snapshotName),
						ID:           snapshotID,
						Size:         services.DefaultDiskSizeBytes,
						SourceDiskID: diskID,
						Ready:        true,
						CreationTime: creationTime,
					},
					err: nil,
				},
			},
		},
		{
			name: "NewSnapshot",
			req: &csi.CreateSnapshotRequest{
				SourceVolumeId: diskID,
				Name:           snapshotName,
			},
			resp: &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SizeBytes:      services.DefaultDiskSizeBytes,
					SnapshotId:     snapshotID,
					SourceVolumeId: diskID,
					CreationTime:   creationTime,
					ReadyToUse:     true,
				},
			},
			mocks: []*Mocks{
				{
					method:     "GetSnapshotByName",
					returnData: &diskapi.Snapshot{},
					err:        status.Error(codes.NotFound, "not found"),
				},
				{
					method: "CreateSnapshot",
					returnData: &diskapi.Snapshot{
						Name:         services.CSIToYCName(snapshotName),
						ID:           snapshotID,
						Size:         services.DefaultDiskSizeBytes,
						SourceDiskID: diskID,
						Ready:        true,
						CreationTime: creationTime,
					},
					err: nil,
				},
			},
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)

		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			d := &diskapimocks.DiskAPI{}
			i := inflight.NewWithTTL()
			if tc.inFlightValue != nil {
				require.True(t, i.Insert(tc.inFlightValue))
			}
			for _, m := range tc.mocks {
				if m.returnData != nil {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.returnData, m.err)
				} else {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.err)
				}
			}

			c := New(d, i, services.ProductionCapsForPublicAPIController)
			resp, err := c.CreateSnapshot(context.TODO(), tc.req)
			checkError(t, err, tc.expErrCode)
			require.Equal(t, tc.resp, resp)
			if tc.inFlightValue != nil {
				i.Delete(tc.inFlightValue)
				require.False(t, i.Get(tc.inFlightValue))
			}
		})
	}
}

func TestController_DeleteSnapshot(t *testing.T) {
	testCases := []struct {
		name          string
		req           *csi.DeleteSnapshotRequest
		resp          *csi.DeleteSnapshotResponse
		expErrCode    codes.Code
		mocks         []*Mocks
		inFlightValue inflight.Idempotent
	}{
		{
			name:       "Delete withoud id",
			req:        &csi.DeleteSnapshotRequest{},
			resp:       nil,
			expErrCode: codes.InvalidArgument,
		},
		{
			name: "Delete success",
			req: &csi.DeleteSnapshotRequest{
				SnapshotId: snapshotID,
			},
			resp: &csi.DeleteSnapshotResponse{},
			mocks: []*Mocks{
				{
					method: "DeleteSnapshot",
					err:    nil,
				},
			},
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)

		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			d := &diskapimocks.DiskAPI{}
			i := inflight.NewWithTTL()
			if tc.inFlightValue != nil {
				require.True(t, i.Insert(tc.inFlightValue))
			}
			for _, m := range tc.mocks {
				if m.returnData != nil {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.returnData, m.err)
				} else {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.err)
				}
			}

			c := New(d, i, services.ProductionCapsForPublicAPIController)
			resp, err := c.DeleteSnapshot(context.TODO(), tc.req)
			checkError(t, err, tc.expErrCode)
			require.Equal(t, tc.resp, resp)
			if tc.inFlightValue != nil {
				i.Delete(tc.inFlightValue)
				require.False(t, i.Get(tc.inFlightValue))
			}
		})
	}
}

func TestController_ListSnapshots(t *testing.T) {
	creationTime := timestamppb.Now()
	testCases := []struct {
		name       string
		req        *csi.ListSnapshotsRequest
		resp       *csi.ListSnapshotsResponse
		expErrCode codes.Code
		mocks      []*Mocks
	}{
		{
			name: "ListSnapshots",
			req:  &csi.ListSnapshotsRequest{},
			resp: &csi.ListSnapshotsResponse{
				Entries: []*csi.ListSnapshotsResponse_Entry{
					{
						Snapshot: &csi.Snapshot{
							SizeBytes:      services.DefaultDiskSizeBytes,
							SnapshotId:     snapshotID,
							SourceVolumeId: diskID,
							CreationTime:   creationTime,
							ReadyToUse:     false,
						},
					},
					{
						Snapshot: &csi.Snapshot{
							SizeBytes:      services.DefaultDiskSizeBytes,
							SnapshotId:     snapshotID + "-2",
							SourceVolumeId: diskID + "-2",
							CreationTime:   creationTime,
							ReadyToUse:     true,
						},
					},
				},
			},
			mocks: []*Mocks{
				{
					method: "ListSnapshots",
					returnData: &diskapi.ListSnapshotsResponse{
						Snapshots: []diskapi.Snapshot{
							{
								ID:           snapshotID,
								SourceDiskID: diskID,
								Name:         snapshotName,
								Size:         services.DefaultDiskSizeBytes,
								CreationTime: creationTime,
								Ready:        false,
							},
							{
								ID:           snapshotID + "-2",
								SourceDiskID: diskID + "-2",
								Name:         snapshotName + "-2",
								Size:         services.DefaultDiskSizeBytes,
								CreationTime: creationTime,
								Ready:        true,
							},
						},
						NextPageToken: "",
					},
					err: nil,
				},
				{
					method: "GetDisk",
					returnData: diskapi.Disk{
						ID:    diskID,
						Size:  services.DefaultDiskSizeBytes,
						Ready: true,
					},
					err: nil,
				}, {
					method: "GetDisk",
					returnData: diskapi.Disk{
						ID:    diskID + "-2",
						Size:  services.DefaultDiskSizeBytes,
						Ready: true,
					},
					err: nil,
				},
			},
		},
	}
	for _, tc := range testCases {
		fmt.Println("TC: ", tc)

		t.Run(tc.name, func(t *testing.T) {
			fmt.Println("Name: ", tc.name)
			d := &diskapimocks.DiskAPI{}
			i := inflight.NewWithTTL()
			for _, m := range tc.mocks {
				if m.returnData != nil {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.returnData, m.err)
				} else {
					d.Mock.On(m.method, mock.Anything, mock.Anything).Return(m.err)
				}
			}

			c := New(d, i, services.ProductionCapsForPublicAPIController)
			resp, err := c.ListSnapshots(context.TODO(), tc.req)
			checkError(t, err, tc.expErrCode)
			require.Equal(t, tc.resp, resp)
		})
	}
}

func newVolumeRequest(disk *diskapi.Disk, diskname string, volcaps []*csi.VolumeCapability) *csi.CreateVolumeRequest {
	return &csi.CreateVolumeRequest{
		Name:               diskname,
		CapacityRange:      &csi.CapacityRange{RequiredBytes: disk.Size - 4*1024},
		VolumeCapabilities: volcaps,
		AccessibilityRequirements: &csi.TopologyRequirement{
			Requisite: []*csi.Topology{
				{Segments: map[string]string{services.ZoneKey: disk.ZoneID}},
			},
		},
		Parameters: map[string]string{services.VolumeTypeKey: disk.TypeID},
	}
}

func publishVolumeRequest(disk *diskapi.Disk, instance string, volcap *csi.VolumeCapability) *csi.ControllerPublishVolumeRequest {
	return &csi.ControllerPublishVolumeRequest{
		VolumeId:         disk.ID,
		NodeId:           instance,
		VolumeCapability: volcap,
	}
}

func checkError(t *testing.T, err error, expErrCode codes.Code) {
	if err != nil {
		srvErr, ok := status.FromError(err)
		if !ok {
			t.Fatalf("Could not get error status code from error: %v", srvErr)
		}
		if srvErr.Code() != expErrCode {
			t.Fatalf("Expected error code %d, got %d message %s", expErrCode, srvErr.Code(), srvErr.Message())
		}
	} else if expErrCode != codes.OK {
		t.Fatalf("Expected error %v, got no error", expErrCode)
	}
}
