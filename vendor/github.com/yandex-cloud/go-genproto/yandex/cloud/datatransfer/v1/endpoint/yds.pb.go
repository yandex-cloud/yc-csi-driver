// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/datatransfer/v1/endpoint/yds.proto

package endpoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type YdsCompressionCodec int32

const (
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_UNSPECIFIED YdsCompressionCodec = 0
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_RAW         YdsCompressionCodec = 1
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_GZIP        YdsCompressionCodec = 2
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_ZSTD        YdsCompressionCodec = 4
)

// Enum value maps for YdsCompressionCodec.
var (
	YdsCompressionCodec_name = map[int32]string{
		0: "YDS_COMPRESSION_CODEC_UNSPECIFIED",
		1: "YDS_COMPRESSION_CODEC_RAW",
		2: "YDS_COMPRESSION_CODEC_GZIP",
		4: "YDS_COMPRESSION_CODEC_ZSTD",
	}
	YdsCompressionCodec_value = map[string]int32{
		"YDS_COMPRESSION_CODEC_UNSPECIFIED": 0,
		"YDS_COMPRESSION_CODEC_RAW":         1,
		"YDS_COMPRESSION_CODEC_GZIP":        2,
		"YDS_COMPRESSION_CODEC_ZSTD":        4,
	}
)

func (x YdsCompressionCodec) Enum() *YdsCompressionCodec {
	p := new(YdsCompressionCodec)
	*p = x
	return p
}

func (x YdsCompressionCodec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (YdsCompressionCodec) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes[0].Descriptor()
}

func (YdsCompressionCodec) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes[0]
}

func (x YdsCompressionCodec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use YdsCompressionCodec.Descriptor instead.
func (YdsCompressionCodec) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{0}
}

type YDSSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Database
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Stream
	Stream string `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	// SA which has read access to the stream.
	ServiceAccountId string `protobuf:"bytes,8,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Compression codec
	SupportedCodecs []YdsCompressionCodec `protobuf:"varint,9,rep,packed,name=supported_codecs,json=supportedCodecs,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec" json:"supported_codecs,omitempty"`
	// Data parsing rules
	Parser *Parser `protobuf:"bytes,10,opt,name=parser,proto3" json:"parser,omitempty"`
	// Should continue working, if consumer read lag exceed TTL of topic
	// False: stop the transfer in error state, if detected lost data. True: continue
	// working with losing part of data
	AllowTtlRewind bool `protobuf:"varint,11,opt,name=allow_ttl_rewind,json=allowTtlRewind,proto3" json:"allow_ttl_rewind,omitempty"`
	// for dedicated db
	Endpoint string `protobuf:"bytes,20,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Network interface for endpoint. If none will assume public ipv4
	SubnetId string `protobuf:"bytes,30,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,34,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	// for important streams
	Consumer string `protobuf:"bytes,35,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *YDSSource) Reset() {
	*x = YDSSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YDSSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YDSSource) ProtoMessage() {}

func (x *YDSSource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YDSSource.ProtoReflect.Descriptor instead.
func (*YDSSource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{0}
}

func (x *YDSSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *YDSSource) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *YDSSource) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *YDSSource) GetSupportedCodecs() []YdsCompressionCodec {
	if x != nil {
		return x.SupportedCodecs
	}
	return nil
}

func (x *YDSSource) GetParser() *Parser {
	if x != nil {
		return x.Parser
	}
	return nil
}

func (x *YDSSource) GetAllowTtlRewind() bool {
	if x != nil {
		return x.AllowTtlRewind
	}
	return false
}

func (x *YDSSource) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *YDSSource) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *YDSSource) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *YDSSource) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

type YDSTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Database
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Stream
	Stream string `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	// SA which has read access to the stream.
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Save transaction order
	// Not to split events queue into separate per-table queues.
	// Incompatible with setting Topic prefix, only with Topic full name.
	SaveTxOrder      bool                `protobuf:"varint,4,opt,name=save_tx_order,json=saveTxOrder,proto3" json:"save_tx_order,omitempty"`
	CompressionCodec YdsCompressionCodec `protobuf:"varint,5,opt,name=compression_codec,json=compressionCodec,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec" json:"compression_codec,omitempty"`
	// Data serialization format
	Serializer *Serializer `protobuf:"bytes,8,opt,name=serializer,proto3" json:"serializer,omitempty"`
	// for dedicated db
	Endpoint string `protobuf:"bytes,20,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Network interface for endpoint. If none will assume public ipv4
	SubnetId string `protobuf:"bytes,30,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,34,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
}

func (x *YDSTarget) Reset() {
	*x = YDSTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YDSTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YDSTarget) ProtoMessage() {}

func (x *YDSTarget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YDSTarget.ProtoReflect.Descriptor instead.
func (*YDSTarget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{1}
}

func (x *YDSTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *YDSTarget) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *YDSTarget) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *YDSTarget) GetSaveTxOrder() bool {
	if x != nil {
		return x.SaveTxOrder
	}
	return false
}

func (x *YDSTarget) GetCompressionCodec() YdsCompressionCodec {
	if x != nil {
		return x.CompressionCodec
	}
	return YdsCompressionCodec_YDS_COMPRESSION_CODEC_UNSPECIFIED
}

func (x *YDSTarget) GetSerializer() *Serializer {
	if x != nil {
		return x.Serializer
	}
	return nil
}

func (x *YDSTarget) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *YDSTarget) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *YDSTarget) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

var File_yandex_cloud_datatransfer_v1_endpoint_yds_proto protoreflect.FileDescriptor

var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x79, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x09, 0x59, 0x44, 0x53, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x59, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x52, 0x0f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x12, 0x45, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x74,
	0x6c, 0x5f, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x74, 0x6c, 0x52, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x22, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x23, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x08, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x14, 0x4a, 0x04, 0x08, 0x15, 0x10, 0x1e, 0x4a, 0x04,
	0x08, 0x1f, 0x10, 0x22, 0x22, 0xc7, 0x03, 0x0a, 0x09, 0x59, 0x44, 0x53, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x78, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x61, 0x76,
	0x65, 0x54, 0x78, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x67, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x59, 0x64, 0x73, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x63, 0x12, 0x51, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x22, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x09,
	0x10, 0x14, 0x4a, 0x04, 0x08, 0x15, 0x10, 0x1e, 0x4a, 0x04, 0x08, 0x1f, 0x10, 0x22, 0x2a, 0x9b,
	0x01, 0x0a, 0x13, 0x59, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x21, 0x59, 0x44, 0x53, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x43, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x19, 0x59, 0x44, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x43, 0x5f, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a,
	0x59, 0x44, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x43, 0x5f, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a,
	0x59, 0x44, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x43, 0x5f, 0x5a, 0x53, 0x54, 0x44, 0x10, 0x04, 0x42, 0xa7, 0x01, 0x0a,
	0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0xaa, 0x02,
	0x25, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData = file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc
)

func file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData)
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes = []any{
	(YdsCompressionCodec)(0), // 0: yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	(*YDSSource)(nil),        // 1: yandex.cloud.datatransfer.v1.endpoint.YDSSource
	(*YDSTarget)(nil),        // 2: yandex.cloud.datatransfer.v1.endpoint.YDSTarget
	(*Parser)(nil),           // 3: yandex.cloud.datatransfer.v1.endpoint.Parser
	(*Serializer)(nil),       // 4: yandex.cloud.datatransfer.v1.endpoint.Serializer
}
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.datatransfer.v1.endpoint.YDSSource.supported_codecs:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	3, // 1: yandex.cloud.datatransfer.v1.endpoint.YDSSource.parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.Parser
	0, // 2: yandex.cloud.datatransfer.v1.endpoint.YDSTarget.compression_codec:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	4, // 3: yandex.cloud.datatransfer.v1.endpoint.YDSTarget.serializer:type_name -> yandex.cloud.datatransfer.v1.endpoint.Serializer
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_yds_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_init()
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*YDSSource); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*YDSTarget); i {
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
			RawDescriptor: file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_yds_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc = nil
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs = nil
}
