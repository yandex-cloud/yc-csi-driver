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

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Disk struct {
	ID                  string
	Size                int64
	ZoneID              string
	Name                string
	TypeID              string
	InstanceIDs         []string
	SourceSnapshotID    string
	Ready               bool
	Labels              map[string]string
	DiskPlacementPolicy *DiskPlacementPolicy
}

func (d *Disk) GetLabels() map[string]string {
	return d.Labels
}

type DiskAttachment struct {
	DiskID   string
	Readonly bool
}

type Instance struct {
	ID              string
	DiskAttachments []DiskAttachment
	Labels          map[string]string
}

func (i *Instance) GetLabels() map[string]string {
	return i.Labels
}

type Snapshot struct {
	ID           string
	SourceDiskID string
	Name         string
	Size         int64
	CreationTime *timestamppb.Timestamp
	Ready        bool
	Labels       map[string]string
}

type CreateDiskRequest struct {
	Name                string
	CSIVolumeName       string
	ZoneID              string
	TypeID              string
	Size                int64
	BlockSize           int64
	SnapshotID          string
	DiskPlacementPolicy *DiskPlacementPolicy
	KMSKeyID            string
}

type DiskPlacementPolicy struct {
	PlacementGroupID        string
	PlacementGroupPartition int64
}

type CreateSnapshotRequest struct {
	Name            string
	CSISnapshotName string
	SourceDiskID    string
}

type DeleteSnapshotRequest struct {
	ID string
}

type ExpandDiskRequest struct {
	ID   string
	Size int64
}

type ListDisksRequest struct {
	Filter string
}

type ListSnapshotsRequest struct {
	Name           string
	Filter         string
	SourceVolumeID string
	MaxEntries     int32
	PageToken      string
}

type ListSnapshotsResponse struct {
	Snapshots     []Snapshot
	NextPageToken string
}

type DeleteDiskRequest struct {
	DiskID string
}

type AttachDiskRequest struct {
	InstanceID string
	DiskID     string
	Readonly   bool
}

type DetachDiskRequest struct {
	DiskID     string
	InstanceID string
}
type GetDiskRequest struct {
	ID string
}

type GetDiskByNameRequest struct {
	Name string
}

type GetSnapshotByNameRequest struct {
	Name string
}

type GetSnapshotRequest struct {
	ID string
}

type GetInstanceRequest struct {
	InstanceID string
}

type EnsureObjectLabelsRequest struct {
	Disk     *Disk
	Instance *Instance
}

//go:generate mockery --name=DiskAPI --outpkg=diskapimocks --case=underscore
type DiskAPI interface {
	CreateDisk(context.Context, *CreateDiskRequest) (*Disk, error)
	// Attach,Detach,Delete return success when operation has already been done earlier.
	DeleteDisk(context.Context, *DeleteDiskRequest) error
	ExpandDisk(context.Context, *ExpandDiskRequest) error
	AttachDisk(context.Context, *AttachDiskRequest) error
	DetachDisk(context.Context, *DetachDiskRequest) error
	ListDisks(context.Context, *ListDisksRequest) ([]*Disk, error)
	GetDisk(context.Context, *GetDiskRequest) (*Disk, error)
	GetDiskByName(context.Context, *GetDiskByNameRequest) (*Disk, error)
	GetInstance(context.Context, *GetInstanceRequest) (*Instance, error)
	CreateSnapshot(context.Context, *CreateSnapshotRequest) (*Snapshot, error)
	DeleteSnapshot(context.Context, *DeleteSnapshotRequest) error
	GetSnapshotByName(context.Context, *GetSnapshotByNameRequest) (*Snapshot, error)
	GetSnapshot(context.Context, *GetSnapshotRequest) (*Snapshot, error)
	ListSnapshots(ctx context.Context, req *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	EnsureObjectLabels(ctx context.Context, req *EnsureObjectLabelsRequest) error
}
