// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/mysql/v1/backup.proto

package mysql

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Backup_BackupCreationType int32

const (
	Backup_BACKUP_CREATION_TYPE_UNSPECIFIED Backup_BackupCreationType = 0
	// Backup created by automated daily schedule
	Backup_AUTOMATED Backup_BackupCreationType = 1
	// Backup created by user request
	Backup_MANUAL Backup_BackupCreationType = 2
)

// Enum value maps for Backup_BackupCreationType.
var (
	Backup_BackupCreationType_name = map[int32]string{
		0: "BACKUP_CREATION_TYPE_UNSPECIFIED",
		1: "AUTOMATED",
		2: "MANUAL",
	}
	Backup_BackupCreationType_value = map[string]int32{
		"BACKUP_CREATION_TYPE_UNSPECIFIED": 0,
		"AUTOMATED":                        1,
		"MANUAL":                           2,
	}
)

func (x Backup_BackupCreationType) Enum() *Backup_BackupCreationType {
	p := new(Backup_BackupCreationType)
	*p = x
	return p
}

func (x Backup_BackupCreationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Backup_BackupCreationType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes[0].Descriptor()
}

func (Backup_BackupCreationType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes[0]
}

func (x Backup_BackupCreationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Backup_BackupCreationType.Descriptor instead.
func (Backup_BackupCreationType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescGZIP(), []int{0, 0}
}

type Backup_BackupStatus int32

const (
	Backup_BACKUP_STATUS_UNSPECIFIED Backup_BackupStatus = 0
	// Backup is done
	Backup_DONE Backup_BackupStatus = 1
	// Backup is creating
	Backup_CREATING Backup_BackupStatus = 2
)

// Enum value maps for Backup_BackupStatus.
var (
	Backup_BackupStatus_name = map[int32]string{
		0: "BACKUP_STATUS_UNSPECIFIED",
		1: "DONE",
		2: "CREATING",
	}
	Backup_BackupStatus_value = map[string]int32{
		"BACKUP_STATUS_UNSPECIFIED": 0,
		"DONE":                      1,
		"CREATING":                  2,
	}
)

func (x Backup_BackupStatus) Enum() *Backup_BackupStatus {
	p := new(Backup_BackupStatus)
	*p = x
	return p
}

func (x Backup_BackupStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Backup_BackupStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes[1].Descriptor()
}

func (Backup_BackupStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes[1]
}

func (x Backup_BackupStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Backup_BackupStatus.Descriptor instead.
func (Backup_BackupStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescGZIP(), []int{0, 1}
}

// An object that represents MySQL backup.
//
// See [the documentation](/docs/managed-mysql/concepts/backup) for details.
type Backup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp (the time when the backup operation was completed).
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Start timestamp (the time when the backup operation was started).
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Size of backup, in bytes
	Size int64 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	// How this backup was created (manual/automatic/etc...)
	Type Backup_BackupCreationType `protobuf:"varint,7,opt,name=type,proto3,enum=yandex.cloud.mdb.mysql.v1.Backup_BackupCreationType" json:"type,omitempty"`
	// Status of backup
	Status Backup_BackupStatus `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.mdb.mysql.v1.Backup_BackupStatus" json:"status,omitempty"`
}

func (x *Backup) Reset() {
	*x = Backup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backup) ProtoMessage() {}

func (x *Backup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backup.ProtoReflect.Descriptor instead.
func (*Backup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescGZIP(), []int{0}
}

func (x *Backup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Backup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Backup) GetSourceClusterId() string {
	if x != nil {
		return x.SourceClusterId
	}
	return ""
}

func (x *Backup) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Backup) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Backup) GetType() Backup_BackupCreationType {
	if x != nil {
		return x.Type
	}
	return Backup_BACKUP_CREATION_TYPE_UNSPECIFIED
}

func (x *Backup) GetStatus() Backup_BackupStatus {
	if x != nil {
		return x.Status
	}
	return Backup_BACKUP_STATUS_UNSPECIFIED
}

var File_yandex_cloud_mdb_mysql_v1_backup_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x14,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x12, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x20, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x02,
	0x22, 0x45, 0x0a, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x19, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x64, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescData = file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDesc
)

func file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDescData
}

var file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_mdb_mysql_v1_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_mdb_mysql_v1_backup_proto_goTypes = []any{
	(Backup_BackupCreationType)(0), // 0: yandex.cloud.mdb.mysql.v1.Backup.BackupCreationType
	(Backup_BackupStatus)(0),       // 1: yandex.cloud.mdb.mysql.v1.Backup.BackupStatus
	(*Backup)(nil),                 // 2: yandex.cloud.mdb.mysql.v1.Backup
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_mdb_mysql_v1_backup_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.mdb.mysql.v1.Backup.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: yandex.cloud.mdb.mysql.v1.Backup.started_at:type_name -> google.protobuf.Timestamp
	0, // 2: yandex.cloud.mdb.mysql.v1.Backup.type:type_name -> yandex.cloud.mdb.mysql.v1.Backup.BackupCreationType
	1, // 3: yandex.cloud.mdb.mysql.v1.Backup.status:type_name -> yandex.cloud.mdb.mysql.v1.Backup.BackupStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mysql_v1_backup_proto_init() }
func file_yandex_cloud_mdb_mysql_v1_backup_proto_init() {
	if File_yandex_cloud_mdb_mysql_v1_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mysql_v1_backup_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Backup); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_mysql_v1_backup_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mysql_v1_backup_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_mysql_v1_backup_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_mysql_v1_backup_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mysql_v1_backup_proto = out.File
	file_yandex_cloud_mdb_mysql_v1_backup_proto_rawDesc = nil
	file_yandex_cloud_mdb_mysql_v1_backup_proto_goTypes = nil
	file_yandex_cloud_mdb_mysql_v1_backup_proto_depIdxs = nil
}
