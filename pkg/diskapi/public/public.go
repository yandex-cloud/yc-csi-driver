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
	"fmt"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	publicCSILabel = "yandex.cloud/public-csi"
)

type publicDiskAPI struct {
	sdk      *ycsdk.SDK
	folderID string
}

func NewPublicDiskAPI(sdk *ycsdk.SDK, folderID string) diskapi.DiskAPI {
	return &publicDiskAPI{
		sdk:      sdk,
		folderID: folderID,
	}
}

func (d *publicDiskAPI) CreateSnapshot(context.Context, *diskapi.CreateSnapshotRequest) (*diskapi.Snapshot, error) {
	return nil, status.Error(codes.Unimplemented, "CreateSnapshot is unimplemented for csi over public sdk")
}

func (d *publicDiskAPI) DeleteSnapshot(context.Context, *diskapi.DeleteSnapshotRequest) error {
	return status.Error(codes.Unimplemented, "DeleteSnapshot is unimplemented for csi over public sdk")
}

func (d *publicDiskAPI) ListSnapshots(context.Context, *diskapi.ListSnapshotsRequest) (*diskapi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ListSnapshots is unimplemented for csi over public sdk")
}

func (d *publicDiskAPI) GetSnapshot(context.Context, *diskapi.GetSnapshotRequest) (*diskapi.Snapshot, error) {
	return nil, status.Error(codes.Unimplemented, "GetSnapshot is unimplemented for csi over public sdk")
}
func (d *publicDiskAPI) GetSnapshotByName(context.Context, *diskapi.GetSnapshotByNameRequest) (*diskapi.Snapshot, error) {
	return nil, status.Error(codes.Unimplemented, "GetSnapshotByName is unimplemented for csi over public sdk")
}

func (d *publicDiskAPI) CreateDisk(ctx context.Context, req *diskapi.CreateDiskRequest) (*diskapi.Disk, error) {
	labels := csiLabels()
	labels["csi-volume-name"] = req.CSIVolumeName

	if req.ZoneID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "zoneID required by diskService (req: %+v)", req)
	}

	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Disk().Create(
			ctx,
			&compute.CreateDiskRequest{
				FolderId:            d.folderID,
				ZoneId:              req.ZoneID,
				TypeId:              req.TypeID,
				Name:                req.Name,
				Labels:              labels,
				Size:                req.Size,
				DiskPlacementPolicy: convertDiskPlacementPolicyToComputeAPI(req.DiskPlacementPolicy),
			},
		),
	)

	if err != nil {
		return nil, err
	}

	if err := op.Wait(ctx); err != nil {
		return nil, err
	}

	resp, err := op.Response()
	if err != nil {
		return nil, err
	}

	disk, ok := resp.(*compute.Disk)
	if !ok {
		return nil, status.Error(codes.Internal, "Error CreateDisk(): no *compute.Disk in op.Response()")
	}

	return convertDiskFromComputeAPI(disk), nil
}

func (d *publicDiskAPI) ExpandDisk(ctx context.Context, req *diskapi.ExpandDiskRequest) error {
	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Disk().Update(
			ctx,
			&compute.UpdateDiskRequest{
				DiskId:     req.ID,
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"size"}},
				Size:       req.Size,
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

func (d *publicDiskAPI) DeleteDisk(ctx context.Context, req *diskapi.DeleteDiskRequest) error {
	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Disk().Delete(
			ctx,
			&compute.DeleteDiskRequest{
				DiskId: req.DiskID,
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

func (d *publicDiskAPI) EnsureObjectLabels(ctx context.Context, req *diskapi.EnsureObjectLabelsRequest) error {
	// Label this compute instance as used for k8s purpose.
	if req.Instance != nil {
		if err := d.labelInstance(ctx, req.Instance); err != nil {
			return err
		}
	}

	// Also, make sure disk is labeled too.
	if req.Disk != nil {
		return d.labelDisk(ctx, req.Disk)
	}

	return nil
}

func (d *publicDiskAPI) AttachDisk(ctx context.Context, req *diskapi.AttachDiskRequest) error {
	attachMode := compute.AttachedDiskSpec_READ_WRITE
	if req.Readonly {
		attachMode = compute.AttachedDiskSpec_READ_ONLY
	}

	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Instance().AttachDisk(
			ctx,
			&compute.AttachInstanceDiskRequest{
				InstanceId: req.InstanceID,
				AttachedDiskSpec: &compute.AttachedDiskSpec{
					Mode: attachMode,
					Disk: &compute.AttachedDiskSpec_DiskId{
						DiskId: req.DiskID,
					},
				},
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

func (d *publicDiskAPI) DetachDisk(ctx context.Context, req *diskapi.DetachDiskRequest) error {
	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Instance().DetachDisk(
			ctx,
			&compute.DetachInstanceDiskRequest{
				InstanceId: req.InstanceID,
				Disk: &compute.DetachInstanceDiskRequest_DiskId{
					DiskId: req.DiskID,
				},
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

func (d *publicDiskAPI) ListDisks(ctx context.Context, req *diskapi.ListDisksRequest) ([]*diskapi.Disk, error) {
	lst, err := d.sdk.Compute().Disk().List(
		ctx,
		&compute.ListDisksRequest{
			FolderId: d.folderID,
			Filter:   req.Filter,
		},
	)

	if err != nil {
		return nil, err
	}

	var dd []*diskapi.Disk
	for _, d := range lst.Disks {
		dd = append(dd, convertDiskFromComputeAPI(d))
	}

	return dd, nil
}

func (d *publicDiskAPI) GetDisk(ctx context.Context, req *diskapi.GetDiskRequest) (*diskapi.Disk, error) {
	disk, err := d.sdk.Compute().Disk().Get(
		ctx,
		&compute.GetDiskRequest{
			DiskId: req.ID,
		},
	)

	if err != nil {
		return nil, err
	}

	return convertDiskFromComputeAPI(disk), nil
}

func (d *publicDiskAPI) GetDiskByName(ctx context.Context, req *diskapi.GetDiskByNameRequest) (*diskapi.Disk, error) {
	disks, err := d.ListDisks(
		ctx,
		&diskapi.ListDisksRequest{
			Filter: fmt.Sprintf("name = \"%s\"", req.Name),
		},
	)

	if err != nil {
		return nil, err
	}

	l := len(disks)
	switch {
	case l < 0:
		return nil, status.Error(codes.Internal, "length of array of disks less than 0")
	case l > 1:
		return nil, status.Errorf(codes.Internal, "multiple disks matches the name: %s", req.Name)
	case l == 1:
		return disks[0], nil
	default:
		return nil, status.Errorf(codes.NotFound, "disk for the name ( %s ) not found", req.Name)
	}
}

func (d *publicDiskAPI) GetInstance(ctx context.Context, req *diskapi.GetInstanceRequest) (*diskapi.Instance, error) {
	instance, err := d.sdk.Compute().Instance().Get(
		ctx, &compute.GetInstanceRequest{
			InstanceId: req.InstanceID,
		},
	)

	if err != nil {
		return nil, err
	}

	return &diskapi.Instance{
		ID:              instance.GetId(),
		DiskAttachments: collectComputeInstanceDisks(instance),
		Labels:          instance.GetLabels(),
	}, nil
}

func (d *publicDiskAPI) labelInstance(ctx context.Context, instance *diskapi.Instance) error {
	labels, updateRequired := csiObjectUpdatedLabels(instance)
	if !updateRequired {
		// Nothing to do, label is already there.
		return nil
	}

	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Instance().Update(
			ctx,
			&compute.UpdateInstanceRequest{
				InstanceId: instance.ID,
				Labels:     labels,
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"labels"}},
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

func (d *publicDiskAPI) labelDisk(ctx context.Context, disk *diskapi.Disk) error {
	labels, updateRequired := csiObjectUpdatedLabels(disk)
	if !updateRequired {
		// Nothing to do, label is already there.
		return nil
	}

	op, err := d.sdk.WrapOperation(
		d.sdk.Compute().Disk().Update(
			ctx,
			&compute.UpdateDiskRequest{
				DiskId:     disk.ID,
				Labels:     labels,
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"labels"}},
			},
		),
	)

	if err != nil {
		return err
	}

	return op.Wait(ctx)
}

type objectWithLabels interface {
	GetLabels() map[string]string
}

func csiObjectUpdatedLabels(obj objectWithLabels) (map[string]string, bool) {
	labels := obj.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}

	requiredLabels := csiLabels()
	updateRequired := false

	for k, rv := range requiredLabels {
		if v, ok := labels[k]; ok && v == rv {
			continue
		}

		labels[k] = rv
		updateRequired = true
	}

	return labels, updateRequired
}

func csiLabels() map[string]string {
	return map[string]string{
		publicCSILabel: "true",
	}
}

func collectComputeInstanceDisks(instance *compute.Instance) []diskapi.DiskAttachment {
	instanceDisks := []diskapi.DiskAttachment{
		{
			DiskID:   instance.GetBootDisk().GetDiskId(),
			Readonly: instance.GetBootDisk().GetMode() == compute.AttachedDisk_READ_ONLY,
		},
	}

	for _, disk := range instance.GetSecondaryDisks() {
		instanceDisks = append(instanceDisks, diskapi.DiskAttachment{
			DiskID:   disk.GetDiskId(),
			Readonly: disk.GetMode() == compute.AttachedDisk_READ_ONLY,
		})
	}

	return instanceDisks
}

func convertDiskFromComputeAPI(disk *compute.Disk) *diskapi.Disk {
	return &diskapi.Disk{
		ID:                  disk.GetId(),
		Size:                disk.GetSize(),
		ZoneID:              disk.GetZoneId(),
		Name:                disk.GetName(),
		TypeID:              disk.GetTypeId(),
		InstanceIDs:         disk.GetInstanceIds(),
		Ready:               disk.GetStatus() == compute.Disk_READY,
		Labels:              disk.GetLabels(),
		SourceSnapshotID:    disk.GetSourceSnapshotId(),
		DiskPlacementPolicy: convertDiskPlacementPolicyFromComputeAPI(disk.GetDiskPlacementPolicy()),
	}
}

func convertDiskPlacementPolicyToComputeAPI(dpp *diskapi.DiskPlacementPolicy) *compute.DiskPlacementPolicy {
	if dpp == nil {
		return nil
	}

	return &compute.DiskPlacementPolicy{
		PlacementGroupId:        dpp.PlacementGroupID,
		PlacementGroupPartition: dpp.PlacementGroupPartition,
	}
}

func convertDiskPlacementPolicyFromComputeAPI(dpp *compute.DiskPlacementPolicy) *diskapi.DiskPlacementPolicy {
	if dpp == nil {
		return nil
	}

	return &diskapi.DiskPlacementPolicy{
		PlacementGroupID:        dpp.PlacementGroupId,
		PlacementGroupPartition: dpp.PlacementGroupPartition,
	}
}
