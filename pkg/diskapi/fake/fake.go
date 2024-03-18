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

package diskapi

import (
	"context"
	"sync"

	"github.com/yandex-cloud/yc-csi-driver/pkg/diskapi"
	"github.com/yandex-cloud/yc-csi-driver/pkg/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	diskID     = "fake-disk-id"
	FakeInstID = "fake-instance"
	FakeZoneID = "fake-zone-id"
)

type tokenData struct {
	processedEntriesNum int32
	dataLen             int32
}

type DiskExt struct {
	diskapi.Disk
	readOnlyAttachmentCount int
	readWriteAttachment     string
}

type folderState struct {
	snapshots map[string]diskapi.Snapshot
	disks     map[string]DiskExt
	instance  map[string]diskapi.Instance
	tokens    map[string]tokenData
}

type fakeDiskAPI struct {
	state       *folderState
	folderID    string
	clusterName string
	lock        sync.Mutex
}

func NewFakeDiskAPI(folderID, clusterName string) diskapi.DiskAPI {
	return &fakeDiskAPI{
		state: &folderState{
			snapshots: map[string]diskapi.Snapshot{},
			disks:     map[string]DiskExt{},
			tokens:    map[string]tokenData{},
			instance: map[string]diskapi.Instance{
				FakeInstID: {
					ID: FakeInstID,
				},
			},
		},
		folderID:    folderID,
		clusterName: clusterName,
	}
}

func (f *fakeDiskAPI) EnsureObjectLabels(ctx context.Context, req *diskapi.EnsureObjectLabelsRequest) error {
	return nil
}

func (f *fakeDiskAPI) GetState() (*folderState, func()) {
	f.lock.Lock()
	return f.state, f.lock.Unlock
}

func (f *fakeDiskAPI) CreateDisk(ctx context.Context, request *diskapi.CreateDiskRequest) (*diskapi.Disk, error) {
	zone := FakeZoneID
	if request.ZoneID != "" {
		zone = request.ZoneID
	}

	disk := diskapi.Disk{
		ID:                  idFromName(request.Name),
		Size:                request.Size,
		ZoneID:              zone,
		Name:                request.Name,
		TypeID:              services.DefaultVolumeType,
		InstanceIDs:         nil,
		SourceSnapshotID:    "",
		Ready:               true,
		Labels:              nil,
		DiskPlacementPolicy: request.DiskPlacementPolicy,
	}

	diskExt := DiskExt{
		Disk: disk,
		// Explicitly fill with default values, for better reading.
		readOnlyAttachmentCount: 0,
		readWriteAttachment:     "",
	}

	state, unlock := f.GetState()
	defer unlock()

	state.disks[disk.ID] = diskExt
	return &disk, nil
}

func (f *fakeDiskAPI) DeleteDisk(ctx context.Context, request *diskapi.DeleteDiskRequest) error {
	state, unlock := f.GetState()
	defer unlock()

	disk, ok := state.disks[request.DiskID]
	// Disk already deleted.
	if !ok {
		return nil
	}

	if len(disk.InstanceIDs) > 0 {
		return status.Errorf(codes.FailedPrecondition, "Trying to delete disk that is still attached")
	}

	delete(state.disks, request.DiskID)

	return nil
}

func (f *fakeDiskAPI) ExpandDisk(ctx context.Context, request *diskapi.ExpandDiskRequest) error {
	//TODO implement me
	panic("implement me")
}

// AttachDisk simulates disk api implementation as close as possible.
//   - When disk is attached to instance, RW or RO, on second attach as RW or RO - we get grpc FAILED_PRECONDITION error.
//   - When trying to attach disk, that does not exist, compute returns NOT_FOUND grpc error.
//   - Compute ALLOWS!(really?!) to attach disk as RW, to another instance,
//     when it is already attached as RO somewhere still, lets forbid such attachments on OUR side.
func (f *fakeDiskAPI) AttachDisk(ctx context.Context, request *diskapi.AttachDiskRequest) error {
	// E.M. below, simulate an actual 'compute attach' operation
	state, unlock := f.GetState()
	defer unlock()

	disk, ok := state.disks[request.DiskID]
	if !ok {
		return status.Errorf(codes.NotFound, "Disk not found for attach disk")
	}

	attachedToRequiredInstance := stringInSlice(request.InstanceID, disk.InstanceIDs)

	// Disk is already attached to this instance, return failed precondition error.
	if attachedToRequiredInstance {
		return status.Errorf(codes.FailedPrecondition, "Trying to attach already attached disk")
	}

	if request.Readonly {
		disk.readOnlyAttachmentCount++
	} else {
		// Trying to attach disk in write mode to multiple instances.
		if disk.readWriteAttachment != "" {
			return status.Errorf(codes.FailedPrecondition, "Cannot attach disk in write mode to multiple instances")
		}

		disk.readWriteAttachment = request.InstanceID
	}

	disk.InstanceIDs = append(disk.InstanceIDs, request.InstanceID)
	state.disks[request.DiskID] = disk

	instance := state.instance[request.InstanceID]
	instance.DiskAttachments = append(instance.DiskAttachments, diskapi.DiskAttachment{
		DiskID:   request.DiskID,
		Readonly: request.Readonly,
	})

	state.instance[request.InstanceID] = instance

	return nil
}

func stringInSlice(s string, sl []string) bool {
	for _, ss := range sl {
		if ss == s {
			return true
		}
	}

	return false
}

func removeStringFromSlice(s string, sl []string) []string {
	for i, v := range sl {
		if v == s {
			return append(sl[:i], sl[i+1:]...)
		}
	}
	return sl
}

func removeDiskFromSlice(d string, sl []diskapi.DiskAttachment) []diskapi.DiskAttachment {
	for i, v := range sl {
		if v.DiskID == d {
			return append(sl[:i], sl[i+1:]...)
		}
	}
	return sl
}

// DetachDisk simulates disk api implementation as close as possible.
// - When we try to detach disk, that is already detached OR DELETED, compute returns FailedPrecondition error.
func (f *fakeDiskAPI) DetachDisk(ctx context.Context, request *diskapi.DetachDiskRequest) error {
	// E.M. below, simulate an actual 'compute detach' operation
	state, unlock := f.GetState()
	defer unlock()
	disk, ok := state.disks[request.DiskID]

	// Disk is already deleted, return error.
	if !ok {
		return status.Errorf(codes.FailedPrecondition, "Trying to detach deleted disk")
	}

	attachedToRequiredInstance := stringInSlice(request.InstanceID, disk.InstanceIDs)

	// Disk is already detached from this instance.
	if !attachedToRequiredInstance {
		return status.Errorf(codes.FailedPrecondition, "Trying to detach already detached disk")
	}

	disk.InstanceIDs = removeStringFromSlice(request.InstanceID, disk.InstanceIDs)
	if disk.readWriteAttachment == request.InstanceID {
		disk.readWriteAttachment = ""
	} else {
		disk.readOnlyAttachmentCount--
	}

	state.disks[request.DiskID] = disk

	instance := state.instance[request.InstanceID]
	instance.DiskAttachments = removeDiskFromSlice(request.DiskID, instance.DiskAttachments)
	state.instance[request.InstanceID] = instance

	return nil
}

func (f *fakeDiskAPI) ListDisks(ctx context.Context, request *diskapi.ListDisksRequest) ([]*diskapi.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (f *fakeDiskAPI) GetDisk(ctx context.Context, request *diskapi.GetDiskRequest) (*diskapi.Disk, error) {
	state, unlock := f.GetState()
	defer unlock()

	if disk, ok := state.disks[request.ID]; ok {
		return &disk.Disk, nil
	}

	return nil, status.Errorf(codes.NotFound, "NotFound")
}

func (f *fakeDiskAPI) GetDiskByName(ctx context.Context, request *diskapi.GetDiskByNameRequest) (*diskapi.Disk, error) {
	state, unlock := f.GetState()
	defer unlock()

	var disk *diskapi.Disk

	for _, d := range state.disks {
		if d.Name == request.Name {
			disk = &d.Disk
			break
		}
	}

	if disk != nil {
		return disk, nil
	}

	return nil, status.Errorf(codes.NotFound, "NotFound")
}

func (f *fakeDiskAPI) GetInstance(ctx context.Context, request *diskapi.GetInstanceRequest) (*diskapi.Instance, error) {
	state, unlock := f.GetState()
	defer unlock()

	instance, ok := state.instance[request.InstanceID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "NotFound")
	}

	return &instance, nil
}

func (f *fakeDiskAPI) CreateSnapshot(ctx context.Context, request *diskapi.CreateSnapshotRequest) (*diskapi.Snapshot, error) {
	disk, err := f.GetDisk(ctx, &diskapi.GetDiskRequest{ID: request.SourceDiskID})
	if err != nil {
		return nil, err
	}

	s := diskapi.Snapshot{
		ID:           rand.String(20),
		SourceDiskID: request.SourceDiskID,
		Name:         request.Name,
		Size:         disk.Size,
		CreationTime: timestamppb.Now(),
		Ready:        rand.IntnRange(0, 100) <= 50,
		Labels:       nil,
	}

	state, unlock := f.GetState()
	defer unlock()

	state.snapshots[s.ID] = s
	return &s, nil
}

func (f *fakeDiskAPI) DeleteSnapshot(ctx context.Context, request *diskapi.DeleteSnapshotRequest) error {
	state, unlock := f.GetState()
	defer unlock()

	delete(state.snapshots, request.ID)

	return nil
}

func (f *fakeDiskAPI) GetSnapshotByName(ctx context.Context, request *diskapi.GetSnapshotByNameRequest) (*diskapi.Snapshot, error) {
	state, unlock := f.GetState()
	defer unlock()

	var snap *diskapi.Snapshot

	for _, s := range state.snapshots {
		if s.Name == request.Name {
			snap = &s
			break
		}
	}

	if snap != nil {
		return snap, nil
	}

	return nil, status.Errorf(codes.NotFound, "NotFound")
}

func (f *fakeDiskAPI) ListSnapshots(ctx context.Context, req *diskapi.ListSnapshotsRequest) (*diskapi.ListSnapshotsResponse, error) {
	state, unlock := f.GetState()
	defer unlock()

	var nextPageToken string

	if req.SourceVolumeID != "" {
		// The only filter can be used is source_volume_id.
		var snapshots []diskapi.Snapshot
		for _, s := range state.snapshots {
			if s.SourceDiskID == req.SourceVolumeID {
				snapshots = append(snapshots, s)
			}
		}

		return &diskapi.ListSnapshotsResponse{
			Snapshots:     snapshots,
			NextPageToken: "",
		}, nil
	}

	var snapshots []diskapi.Snapshot

	for _, s := range state.snapshots {
		snapshots = append(snapshots, s)
	}

	ss := snapshots

	if req.MaxEntries > 0 {
		start, finish := f.getRangeWithToken(req.MaxEntries, req.PageToken)
		nextPageToken = f.nextPageToken(req.MaxEntries, req.PageToken)
		ss = snapshots[start:finish]
	}

	if req.PageToken != "" {
		start, finish := f.getRangeWithToken(req.MaxEntries, req.PageToken)
		nextPageToken = f.nextPageToken(req.MaxEntries, req.PageToken)
		ss = snapshots[start:finish]
	}

	return &diskapi.ListSnapshotsResponse{Snapshots: ss, NextPageToken: nextPageToken}, nil
}

func (f *fakeDiskAPI) GetSnapshot(ctx context.Context, request *diskapi.GetSnapshotRequest) (*diskapi.Snapshot, error) {
	state, unlock := f.GetState()
	defer unlock()

	s, ok := state.snapshots[request.ID]
	if !ok {
		return nil, status.Error(codes.NotFound, "NotFound")
	}

	return &s, nil
}

func (f *fakeDiskAPI) nextPageToken(maxEntries int32, token string) string {
	// E.M. f.lock - is not a recursive mutex =( We should be careful about using it.
	state := f.state

	var oldFinish int32

	nextToken := rand.String(20)
	if token != "" {
		if tkn, ok := state.tokens[token]; ok {
			oldFinish = tkn.processedEntriesNum
		}
		delete(state.tokens, token)
	}

	processedEntries := oldFinish + maxEntries
	dataLen := int32(len(state.snapshots))

	if processedEntries > dataLen-1 {
		processedEntries = dataLen - 1
	}

	state.tokens[nextToken] = tokenData{
		processedEntriesNum: processedEntries,
		dataLen:             dataLen,
	}

	return nextToken
}

func (f *fakeDiskAPI) getRangeWithToken(maxEntries int32, token string) (int32, int32) {
	// E.M. f.lock - is not a recursive mutex =( We should be careful about using it.
	state := f.state

	tkn, ok := state.tokens[token]
	if !ok {
		return 0, maxEntries
	}

	start := tkn.processedEntriesNum
	finish := tkn.dataLen

	if maxEntries > 0 && start+maxEntries < finish {
		return start, start + maxEntries
	}

	return start, finish
}
