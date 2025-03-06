// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/compute/v1/disk_placement_group.proto

package compute

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiskPlacementGroup_Status int32

const (
	DiskPlacementGroup_STATUS_UNSPECIFIED DiskPlacementGroup_Status = 0
	DiskPlacementGroup_CREATING           DiskPlacementGroup_Status = 1
	DiskPlacementGroup_READY              DiskPlacementGroup_Status = 2
	DiskPlacementGroup_DELETING           DiskPlacementGroup_Status = 4
)

// Enum value maps for DiskPlacementGroup_Status.
var (
	DiskPlacementGroup_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		4: "DELETING",
	}
	DiskPlacementGroup_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"READY":              2,
		"DELETING":           4,
	}
)

func (x DiskPlacementGroup_Status) Enum() *DiskPlacementGroup_Status {
	p := new(DiskPlacementGroup_Status)
	*p = x
	return p
}

func (x DiskPlacementGroup_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiskPlacementGroup_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_enumTypes[0].Descriptor()
}

func (DiskPlacementGroup_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_disk_placement_group_proto_enumTypes[0]
}

func (x DiskPlacementGroup_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskPlacementGroup_Status.Descriptor instead.
func (DiskPlacementGroup_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescGZIP(), []int{0, 0}
}

type DiskPlacementGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the placement group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the placement group belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the placement group.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the placement group.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the availability zone where the placement group resides.
	ZoneId string `protobuf:"bytes,7,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Current status of the placement group
	Status DiskPlacementGroup_Status `protobuf:"varint,11,opt,name=status,proto3,enum=yandex.cloud.compute.v1.DiskPlacementGroup_Status" json:"status,omitempty"`
	// Placement strategy.
	//
	// Types that are assignable to PlacementStrategy:
	//
	//	*DiskPlacementGroup_SpreadPlacementStrategy
	//	*DiskPlacementGroup_PartitionPlacementStrategy
	PlacementStrategy isDiskPlacementGroup_PlacementStrategy `protobuf_oneof:"placement_strategy"`
}

func (x *DiskPlacementGroup) Reset() {
	*x = DiskPlacementGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskPlacementGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskPlacementGroup) ProtoMessage() {}

func (x *DiskPlacementGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskPlacementGroup.ProtoReflect.Descriptor instead.
func (*DiskPlacementGroup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescGZIP(), []int{0}
}

func (x *DiskPlacementGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DiskPlacementGroup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *DiskPlacementGroup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DiskPlacementGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiskPlacementGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DiskPlacementGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DiskPlacementGroup) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *DiskPlacementGroup) GetStatus() DiskPlacementGroup_Status {
	if x != nil {
		return x.Status
	}
	return DiskPlacementGroup_STATUS_UNSPECIFIED
}

func (m *DiskPlacementGroup) GetPlacementStrategy() isDiskPlacementGroup_PlacementStrategy {
	if m != nil {
		return m.PlacementStrategy
	}
	return nil
}

func (x *DiskPlacementGroup) GetSpreadPlacementStrategy() *DiskSpreadPlacementStrategy {
	if x, ok := x.GetPlacementStrategy().(*DiskPlacementGroup_SpreadPlacementStrategy); ok {
		return x.SpreadPlacementStrategy
	}
	return nil
}

func (x *DiskPlacementGroup) GetPartitionPlacementStrategy() *DiskPartitionPlacementStrategy {
	if x, ok := x.GetPlacementStrategy().(*DiskPlacementGroup_PartitionPlacementStrategy); ok {
		return x.PartitionPlacementStrategy
	}
	return nil
}

type isDiskPlacementGroup_PlacementStrategy interface {
	isDiskPlacementGroup_PlacementStrategy()
}

type DiskPlacementGroup_SpreadPlacementStrategy struct {
	// Distribute disks over distinct failure domains.
	SpreadPlacementStrategy *DiskSpreadPlacementStrategy `protobuf:"bytes,8,opt,name=spread_placement_strategy,json=spreadPlacementStrategy,proto3,oneof"`
}

type DiskPlacementGroup_PartitionPlacementStrategy struct {
	// Distribute disks over partitions.
	PartitionPlacementStrategy *DiskPartitionPlacementStrategy `protobuf:"bytes,9,opt,name=partition_placement_strategy,json=partitionPlacementStrategy,proto3,oneof"`
}

func (*DiskPlacementGroup_SpreadPlacementStrategy) isDiskPlacementGroup_PlacementStrategy() {}

func (*DiskPlacementGroup_PartitionPlacementStrategy) isDiskPlacementGroup_PlacementStrategy() {}

type DiskSpreadPlacementStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiskSpreadPlacementStrategy) Reset() {
	*x = DiskSpreadPlacementStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskSpreadPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskSpreadPlacementStrategy) ProtoMessage() {}

func (x *DiskSpreadPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskSpreadPlacementStrategy.ProtoReflect.Descriptor instead.
func (*DiskSpreadPlacementStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescGZIP(), []int{1}
}

type DiskPartitionPlacementStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partitions int64 `protobuf:"varint,1,opt,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *DiskPartitionPlacementStrategy) Reset() {
	*x = DiskPartitionPlacementStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskPartitionPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskPartitionPlacementStrategy) ProtoMessage() {}

func (x *DiskPartitionPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskPartitionPlacementStrategy.ProtoReflect.Descriptor instead.
func (*DiskPartitionPlacementStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescGZIP(), []int{2}
}

func (x *DiskPartitionPlacementStrategy) GetPartitions() int64 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

var File_yandex_cloud_compute_v1_disk_placement_group_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9,
	0x05, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x72, 0x0a, 0x19, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x72, 0x65, 0x61,
	0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x48, 0x00, 0x52, 0x17, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x7b, 0x0a,
	0x1c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x6b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x1a,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x14,
	0x0a, 0x12, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x69,
	0x73, 0x6b, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x40, 0x0a, 0x1e, 0x44, 0x69, 0x73,
	0x6b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x62, 0x0a, 0x1b, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescData = file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDesc
)

func file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDescData
}

var file_yandex_cloud_compute_v1_disk_placement_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_compute_v1_disk_placement_group_proto_goTypes = []any{
	(DiskPlacementGroup_Status)(0),         // 0: yandex.cloud.compute.v1.DiskPlacementGroup.Status
	(*DiskPlacementGroup)(nil),             // 1: yandex.cloud.compute.v1.DiskPlacementGroup
	(*DiskSpreadPlacementStrategy)(nil),    // 2: yandex.cloud.compute.v1.DiskSpreadPlacementStrategy
	(*DiskPartitionPlacementStrategy)(nil), // 3: yandex.cloud.compute.v1.DiskPartitionPlacementStrategy
	nil,                                    // 4: yandex.cloud.compute.v1.DiskPlacementGroup.LabelsEntry
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_disk_placement_group_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.compute.v1.DiskPlacementGroup.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.compute.v1.DiskPlacementGroup.labels:type_name -> yandex.cloud.compute.v1.DiskPlacementGroup.LabelsEntry
	0, // 2: yandex.cloud.compute.v1.DiskPlacementGroup.status:type_name -> yandex.cloud.compute.v1.DiskPlacementGroup.Status
	2, // 3: yandex.cloud.compute.v1.DiskPlacementGroup.spread_placement_strategy:type_name -> yandex.cloud.compute.v1.DiskSpreadPlacementStrategy
	3, // 4: yandex.cloud.compute.v1.DiskPlacementGroup.partition_placement_strategy:type_name -> yandex.cloud.compute.v1.DiskPartitionPlacementStrategy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_disk_placement_group_proto_init() }
func file_yandex_cloud_compute_v1_disk_placement_group_proto_init() {
	if File_yandex_cloud_compute_v1_disk_placement_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DiskPlacementGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DiskSpreadPlacementStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DiskPartitionPlacementStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes[0].OneofWrappers = []any{
		(*DiskPlacementGroup_SpreadPlacementStrategy)(nil),
		(*DiskPlacementGroup_PartitionPlacementStrategy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_disk_placement_group_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_disk_placement_group_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_compute_v1_disk_placement_group_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_compute_v1_disk_placement_group_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_disk_placement_group_proto = out.File
	file_yandex_cloud_compute_v1_disk_placement_group_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_disk_placement_group_proto_goTypes = nil
	file_yandex_cloud_compute_v1_disk_placement_group_proto_depIdxs = nil
}
