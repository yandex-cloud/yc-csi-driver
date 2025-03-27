// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Snapshot) SetId(v string) {
	m.Id = v
}

func (m *Snapshot) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Snapshot) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Snapshot) SetName(v string) {
	m.Name = v
}

func (m *Snapshot) SetDescription(v string) {
	m.Description = v
}

func (m *Snapshot) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Snapshot) SetStorageSize(v int64) {
	m.StorageSize = v
}

func (m *Snapshot) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Snapshot) SetProductIds(v []string) {
	m.ProductIds = v
}

func (m *Snapshot) SetStatus(v Snapshot_Status) {
	m.Status = v
}

func (m *Snapshot) SetSourceDiskId(v string) {
	m.SourceDiskId = v
}

func (m *Snapshot) SetHardwareGeneration(v *HardwareGeneration) {
	m.HardwareGeneration = v
}

func (m *Snapshot) SetKmsKey(v *KMSKey) {
	m.KmsKey = v
}
