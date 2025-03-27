// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *ProductIDs) SetProductIds(v []string) {
	m.ProductIds = v
}

type BootDiskSpec_BootSource = isBootDiskSpec_BootSource

func (m *BootDiskSpec) SetBootSource(v BootDiskSpec_BootSource) {
	m.BootSource = v
}

func (m *BootDiskSpec) SetDiskId(v string) {
	m.BootSource = &BootDiskSpec_DiskId{
		DiskId: v,
	}
}

func (m *BootDiskSpec) SetImageId(v string) {
	m.BootSource = &BootDiskSpec_ImageId{
		ImageId: v,
	}
}

func (m *BootDiskSpec) SetSnapshotId(v string) {
	m.BootSource = &BootDiskSpec_SnapshotId{
		SnapshotId: v,
	}
}

func (m *BootDiskSpec) SetProductIds(v *ProductIDs) {
	m.BootSource = &BootDiskSpec_ProductIds{
		ProductIds: v,
	}
}

func (m *ReservedInstancePool) SetId(v string) {
	m.Id = v
}

func (m *ReservedInstancePool) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ReservedInstancePool) SetCloudId(v string) {
	m.CloudId = v
}

func (m *ReservedInstancePool) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ReservedInstancePool) SetName(v string) {
	m.Name = v
}

func (m *ReservedInstancePool) SetDescription(v string) {
	m.Description = v
}

func (m *ReservedInstancePool) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *ReservedInstancePool) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *ReservedInstancePool) SetPlatformId(v string) {
	m.PlatformId = v
}

func (m *ReservedInstancePool) SetResourcesSpec(v *ResourcesSpec) {
	m.ResourcesSpec = v
}

func (m *ReservedInstancePool) SetGpuSettings(v *GpuSettings) {
	m.GpuSettings = v
}

func (m *ReservedInstancePool) SetProductIds(v []string) {
	m.ProductIds = v
}

func (m *ReservedInstancePool) SetNetworkSettings(v *NetworkSettings) {
	m.NetworkSettings = v
}

func (m *ReservedInstancePool) SetSize(v int64) {
	m.Size = v
}
