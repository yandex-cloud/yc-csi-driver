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

package public

import (
	"strings"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	pubOperation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"

	"github.com/yandex-cloud/yc-csi-driver/pkg/testutil/pubapimock"
)

type testData struct {
	instance []*compute.Instance
	disk     []*compute.Disk
}

func diskCreateMock(api *pubapimock.API, folderID string, testData *testData) {
	api.Compute.Disk.On("Create", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.CreateDiskRequest) (*pubOperation.Operation, error) {
			disk := &compute.Disk{
				Id:          uuid.New().String(),
				FolderId:    req.GetFolderId(),
				Name:        req.GetName(),
				Description: req.GetDescription(),
				TypeId:      req.GetTypeId(),
				ZoneId:      req.GetZoneId(),
				Size:        req.GetSize(),
				KmsKey: &compute.KMSKey{
					KeyId: req.GetKmsKeyId(),
				},
			}
			testData.disk = append(testData.disk, disk)
			return successOperation(disk)
		},
	)
}

func diskUpdateMock(api *pubapimock.API, folderID string, testData *testData) {
	api.Compute.Disk.On("Update", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.UpdateDiskRequest) (*pubOperation.Operation, error) {
			for _, d := range testData.disk {
				if d.GetId() == req.DiskId {
					d.Size = req.GetSize()
				}
			}
			return successOperation(&anypb.Any{})
		},
	)
}

func diskDeleteMock(api *pubapimock.API, testData *testData) {
	api.Compute.Disk.On("Delete", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.DeleteDiskRequest) (*pubOperation.Operation, error) {
			removeDisk := func(slice []*compute.Disk, s int) []*compute.Disk {
				return append(slice[:s], slice[s+1:]...)
			}
			for i, d := range testData.disk {
				if d.GetId() == req.DiskId {
					testData.disk = removeDisk(testData.disk, i)
				}
			}
			return successOperation(&anypb.Any{})
		},
	)
}

func diskGetMock(api *pubapimock.API, testData *testData) {
	api.Compute.Disk.On("Get", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.GetDiskRequest) (*compute.Disk, error) {
			for _, d := range testData.disk {
				if d.GetId() == req.DiskId {
					return d, nil
				}
			}
			return nil, status.Error(codes.NotFound, "Disk not found")
		},
	)
}

func diskListMock(api *pubapimock.API, testData *testData) {
	api.Compute.Disk.On("List", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.ListDisksRequest) (*compute.ListDisksResponse, error) {
			name := strings.Trim(strings.TrimSpace(strings.Split(req.Filter, "=")[1]), `"`)
			for _, d := range testData.disk {
				if strings.Contains(d.Name, name) {
					return &compute.ListDisksResponse{Disks: []*compute.Disk{d}}, nil
				}
			}
			return &compute.ListDisksResponse{}, nil
		})
}

func instanceGetMock(api *pubapimock.API, testData *testData) {
	api.Compute.Instance.On("Get", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.GetInstanceRequest) (*compute.Instance, error) {
			for _, inst := range testData.instance {
				if inst.GetId() == req.InstanceId {
					return inst, nil
				}
			}
			return nil, status.Error(codes.NotFound, "Instance not found")
		},
	)
}

func instanceUpdateLabelsMock(api *pubapimock.API, testData *testData) {
	api.Compute.Instance.On("Update", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.UpdateInstanceRequest) (*pubOperation.Operation, error) {
			for _, inst := range testData.instance {
				if inst.GetId() == req.GetInstanceId() {
					inst.Labels = req.GetLabels()
					break
				}
			}

			return successOperation(&anypb.Any{})
		},
	)
}

func instanceAttachDiskMock(api *pubapimock.API, testData *testData) {
	api.Compute.Instance.On("AttachDisk", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.AttachInstanceDiskRequest) (*pubOperation.Operation, error) {
			diskID := req.GetAttachedDiskSpec().GetDiskId()
			instanceID := req.GetInstanceId()
			if checkDiskWithinInstances(diskID, testData) && checkInstanceWithinDisks(instanceID, testData) {
				return successOperation(&anypb.Any{})
			}
			if !checkDiskWithinInstances(diskID, testData) {
				addDiskToInstance(testData, req)
			}
			if !checkInstanceWithinDisks(instanceID, testData) {
				addInstanceToDisk(testData, diskID, instanceID)
			}
			return successOperation(&anypb.Any{})
		},
	)
}

func instanceDetachDiskMock(api *pubapimock.API, testData *testData) {
	api.Compute.Instance.On("DetachDisk", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, req *compute.DetachInstanceDiskRequest) (*pubOperation.Operation, error) {
			if !checkInstanceWithinDisks(req.GetInstanceId(), testData) &&
				!checkDiskWithinInstances(req.GetDiskId(), testData) {
				return successOperation(&anypb.Any{})
			}
			if checkDiskWithinInstances(req.GetDiskId(), testData) {
				removeDiskFromInstance(testData, req)
			}
			if checkInstanceWithinDisks(req.GetInstanceId(), testData) {
				removeInstanceFromDisks(testData, req.GetInstanceId())
			}
			return successOperation(&anypb.Any{})
		},
	)
}

func checkDiskWithinInstances(diskID string, testData *testData) bool {
	for _, i := range testData.instance {
		for _, d := range i.SecondaryDisks {
			if d.GetDiskId() == diskID {
				return true
			}
		}
	}
	return false
}

func checkInstanceWithinDisks(instanceID string, testData *testData) bool {
	for _, d := range testData.disk {
		for _, instID := range d.GetInstanceIds() {
			if instID == instanceID {
				return true
			}
		}
	}
	return false
}
func removeDiskFromInstance(testData *testData, req *compute.DetachInstanceDiskRequest) {
	removeSecondaryDisk := func(slice []*compute.AttachedDisk, s int) []*compute.AttachedDisk {
		return append(slice[:s], slice[s+1:]...)
	}
	for index, inst := range testData.instance {
		if inst.GetId() == req.GetInstanceId() {
			for i, d := range inst.SecondaryDisks {
				if d.GetDiskId() == req.GetDiskId() {
					testData.instance[index].SecondaryDisks = removeSecondaryDisk(inst.SecondaryDisks, i)
				}
			}
		}
	}
}

func addDiskToInstance(testData *testData, req *compute.AttachInstanceDiskRequest) {
	for _, inst := range testData.instance {
		if inst.GetId() == req.GetInstanceId() {
			inst.SecondaryDisks = append(inst.SecondaryDisks,
				&compute.AttachedDisk{
					Mode:       compute.AttachedDisk_READ_WRITE,
					DeviceName: req.GetAttachedDiskSpec().GetDeviceName(),
					AutoDelete: false,
					DiskId:     req.GetAttachedDiskSpec().GetDiskId(),
				})
		}
	}
}

func addInstanceToDisk(testData *testData, diskID string, instanceID string) {
	for _, d := range testData.disk {
		if d.GetId() == diskID {
			d.InstanceIds = append(d.InstanceIds, instanceID)
		}
	}
}
func removeInstanceFromDisks(testData *testData, instanceID string) {
	removeInstance := func(slice []string, s int) []string {
		return append(slice[:s], slice[s+1:]...)
	}
	for index, d := range testData.disk {
		for i, instID := range d.GetInstanceIds() {
			if instID == instanceID {
				testData.disk[index].InstanceIds = removeInstance(d.InstanceIds, i)
			}

		}
	}
}

func computeInstance(folderID string) *compute.Instance {
	return &compute.Instance{
		Id:        instanceID,
		FolderId:  folderID,
		CreatedAt: nil,
		BootDisk: &compute.AttachedDisk{
			Mode:       compute.AttachedDisk_READ_WRITE,
			DeviceName: uuid.New().String(),
			AutoDelete: false,
			DiskId:     uuid.New().String(),
		},
		Name:       uuid.New().String(),
		ZoneId:     zoneID,
		PlatformId: "platform",
		Fqdn:       uuid.New().String(),
	}
}

func computeDisk(folderID string) *compute.Disk {
	return &compute.Disk{
		Id:       uuid.New().String(),
		FolderId: folderID,
		Name:     uuid.New().String(),
		Labels:   nil,
		TypeId:   diskTypeID,
		ZoneId:   zoneID,
		Size:     defaultDiskSizeBytes,
	}
}

func successOperation(pb proto.Message) (*pubOperation.Operation, error) {
	prt, err := anypb.New(pb)
	if err != nil {
		return nil, err
	}
	return &pubOperation.Operation{
		Done: true,
		Result: &pubOperation.Operation_Response{
			Response: prt,
		},
	}, nil
}
