// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/datatransfer/v1/endpoint/mongo.proto

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

type OnPremiseMongo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts      []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Port       int64    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	TlsMode    *TLSMode `protobuf:"bytes,6,opt,name=tls_mode,json=tlsMode,proto3" json:"tls_mode,omitempty"`
	ReplicaSet string   `protobuf:"bytes,5,opt,name=replica_set,json=replicaSet,proto3" json:"replica_set,omitempty"`
}

func (x *OnPremiseMongo) Reset() {
	*x = OnPremiseMongo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnPremiseMongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPremiseMongo) ProtoMessage() {}

func (x *OnPremiseMongo) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPremiseMongo.ProtoReflect.Descriptor instead.
func (*OnPremiseMongo) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{0}
}

func (x *OnPremiseMongo) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *OnPremiseMongo) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *OnPremiseMongo) GetTlsMode() *TLSMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

func (x *OnPremiseMongo) GetReplicaSet() string {
	if x != nil {
		return x.ReplicaSet
	}
	return ""
}

type MongoConnectionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//
	//	*MongoConnectionOptions_MdbClusterId
	//	*MongoConnectionOptions_OnPremise
	Address isMongoConnectionOptions_Address `protobuf_oneof:"address"`
	// User name
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for user
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Database name associated with the credentials
	AuthSource string `protobuf:"bytes,5,opt,name=auth_source,json=authSource,proto3" json:"auth_source,omitempty"`
}

func (x *MongoConnectionOptions) Reset() {
	*x = MongoConnectionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoConnectionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoConnectionOptions) ProtoMessage() {}

func (x *MongoConnectionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoConnectionOptions.ProtoReflect.Descriptor instead.
func (*MongoConnectionOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{1}
}

func (m *MongoConnectionOptions) GetAddress() isMongoConnectionOptions_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *MongoConnectionOptions) GetMdbClusterId() string {
	if x, ok := x.GetAddress().(*MongoConnectionOptions_MdbClusterId); ok {
		return x.MdbClusterId
	}
	return ""
}

func (x *MongoConnectionOptions) GetOnPremise() *OnPremiseMongo {
	if x, ok := x.GetAddress().(*MongoConnectionOptions_OnPremise); ok {
		return x.OnPremise
	}
	return nil
}

func (x *MongoConnectionOptions) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MongoConnectionOptions) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *MongoConnectionOptions) GetAuthSource() string {
	if x != nil {
		return x.AuthSource
	}
	return ""
}

type isMongoConnectionOptions_Address interface {
	isMongoConnectionOptions_Address()
}

type MongoConnectionOptions_MdbClusterId struct {
	MdbClusterId string `protobuf:"bytes,1,opt,name=mdb_cluster_id,json=mdbClusterId,proto3,oneof"`
}

type MongoConnectionOptions_OnPremise struct {
	OnPremise *OnPremiseMongo `protobuf:"bytes,2,opt,name=on_premise,json=onPremise,proto3,oneof"`
}

func (*MongoConnectionOptions_MdbClusterId) isMongoConnectionOptions_Address() {}

func (*MongoConnectionOptions_OnPremise) isMongoConnectionOptions_Address() {}

type MongoConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Connection:
	//
	//	*MongoConnection_ConnectionOptions
	Connection isMongoConnection_Connection `protobuf_oneof:"connection"`
}

func (x *MongoConnection) Reset() {
	*x = MongoConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoConnection) ProtoMessage() {}

func (x *MongoConnection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoConnection.ProtoReflect.Descriptor instead.
func (*MongoConnection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{2}
}

func (m *MongoConnection) GetConnection() isMongoConnection_Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (x *MongoConnection) GetConnectionOptions() *MongoConnectionOptions {
	if x, ok := x.GetConnection().(*MongoConnection_ConnectionOptions); ok {
		return x.ConnectionOptions
	}
	return nil
}

type isMongoConnection_Connection interface {
	isMongoConnection_Connection()
}

type MongoConnection_ConnectionOptions struct {
	ConnectionOptions *MongoConnectionOptions `protobuf:"bytes,3,opt,name=connection_options,json=connectionOptions,proto3,oneof"`
}

func (*MongoConnection_ConnectionOptions) isMongoConnection_Connection() {}

type MongoCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseName   string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	CollectionName string `protobuf:"bytes,2,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
}

func (x *MongoCollection) Reset() {
	*x = MongoCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoCollection) ProtoMessage() {}

func (x *MongoCollection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoCollection.ProtoReflect.Descriptor instead.
func (*MongoCollection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{3}
}

func (x *MongoCollection) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *MongoCollection) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

type MongoSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *MongoConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	SubnetId   string           `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,11,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	// List of collections for replication. Empty list implies replication of all
	// tables on the deployment. Allowed to use * as collection name.
	Collections []*MongoCollection `protobuf:"bytes,6,rep,name=collections,proto3" json:"collections,omitempty"`
	// List of forbidden collections for replication. Allowed to use * as collection
	// name for forbid all collections of concrete schema.
	ExcludedCollections []*MongoCollection `protobuf:"bytes,7,rep,name=excluded_collections,json=excludedCollections,proto3" json:"excluded_collections,omitempty"`
	// Read mode for mongo client
	SecondaryPreferredMode bool `protobuf:"varint,8,opt,name=secondary_preferred_mode,json=secondaryPreferredMode,proto3" json:"secondary_preferred_mode,omitempty"`
}

func (x *MongoSource) Reset() {
	*x = MongoSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoSource) ProtoMessage() {}

func (x *MongoSource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoSource.ProtoReflect.Descriptor instead.
func (*MongoSource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{4}
}

func (x *MongoSource) GetConnection() *MongoConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MongoSource) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *MongoSource) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *MongoSource) GetCollections() []*MongoCollection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *MongoSource) GetExcludedCollections() []*MongoCollection {
	if x != nil {
		return x.ExcludedCollections
	}
	return nil
}

func (x *MongoSource) GetSecondaryPreferredMode() bool {
	if x != nil {
		return x.SecondaryPreferredMode
	}
	return false
}

type MongoTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *MongoConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	SubnetId   string           `protobuf:"bytes,7,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,8,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	// Database name
	Database      string        `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	CleanupPolicy CleanupPolicy `protobuf:"varint,6,opt,name=cleanup_policy,json=cleanupPolicy,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy" json:"cleanup_policy,omitempty"`
}

func (x *MongoTarget) Reset() {
	*x = MongoTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoTarget) ProtoMessage() {}

func (x *MongoTarget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoTarget.ProtoReflect.Descriptor instead.
func (*MongoTarget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP(), []int{5}
}

func (x *MongoTarget) GetConnection() *MongoConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *MongoTarget) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *MongoTarget) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *MongoTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *MongoTarget) GetCleanupPolicy() CleanupPolicy {
	if x != nil {
		return x.CleanupPolicy
	}
	return CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED
}

var File_yandex_cloud_datatransfer_v1_endpoint_mongo_proto protoreflect.FileDescriptor

var file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDesc = []byte{
	0x0a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x32, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6,
	0x01, 0x0a, 0x0e, 0x4f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x74,
	0x6c, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x4c, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x74,
	0x6c, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x53, 0x65, 0x74, 0x22, 0xa3, 0x02, 0x0a, 0x16, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x64, 0x62, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x64,
	0x62, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0a, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x73, 0x65,
	0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x6d, 0x69,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8f, 0x01,
	0x0a, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x6e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x11,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5f, 0x0a, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xaa, 0x03, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x58,
	0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67,
	0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x69, 0x0a, 0x14, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xa4, 0x02,
	0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x56, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x75, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x42, 0xa7, 0x01, 0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0xaa, 0x02, 0x25, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescData = file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDesc
)

func file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescData)
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_goTypes = []interface{}{
	(*OnPremiseMongo)(nil),         // 0: yandex.cloud.datatransfer.v1.endpoint.OnPremiseMongo
	(*MongoConnectionOptions)(nil), // 1: yandex.cloud.datatransfer.v1.endpoint.MongoConnectionOptions
	(*MongoConnection)(nil),        // 2: yandex.cloud.datatransfer.v1.endpoint.MongoConnection
	(*MongoCollection)(nil),        // 3: yandex.cloud.datatransfer.v1.endpoint.MongoCollection
	(*MongoSource)(nil),            // 4: yandex.cloud.datatransfer.v1.endpoint.MongoSource
	(*MongoTarget)(nil),            // 5: yandex.cloud.datatransfer.v1.endpoint.MongoTarget
	(*TLSMode)(nil),                // 6: yandex.cloud.datatransfer.v1.endpoint.TLSMode
	(*Secret)(nil),                 // 7: yandex.cloud.datatransfer.v1.endpoint.Secret
	(CleanupPolicy)(0),             // 8: yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy
}
var file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.datatransfer.v1.endpoint.OnPremiseMongo.tls_mode:type_name -> yandex.cloud.datatransfer.v1.endpoint.TLSMode
	0, // 1: yandex.cloud.datatransfer.v1.endpoint.MongoConnectionOptions.on_premise:type_name -> yandex.cloud.datatransfer.v1.endpoint.OnPremiseMongo
	7, // 2: yandex.cloud.datatransfer.v1.endpoint.MongoConnectionOptions.password:type_name -> yandex.cloud.datatransfer.v1.endpoint.Secret
	1, // 3: yandex.cloud.datatransfer.v1.endpoint.MongoConnection.connection_options:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoConnectionOptions
	2, // 4: yandex.cloud.datatransfer.v1.endpoint.MongoSource.connection:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoConnection
	3, // 5: yandex.cloud.datatransfer.v1.endpoint.MongoSource.collections:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoCollection
	3, // 6: yandex.cloud.datatransfer.v1.endpoint.MongoSource.excluded_collections:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoCollection
	2, // 7: yandex.cloud.datatransfer.v1.endpoint.MongoTarget.connection:type_name -> yandex.cloud.datatransfer.v1.endpoint.MongoConnection
	8, // 8: yandex.cloud.datatransfer.v1.endpoint.MongoTarget.cleanup_policy:type_name -> yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_mongo_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnPremiseMongo); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoConnectionOptions); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoConnection); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoCollection); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoSource); i {
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
		file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoTarget); i {
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
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MongoConnectionOptions_MdbClusterId)(nil),
		(*MongoConnectionOptions_OnPremise)(nil),
	}
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MongoConnection_ConnectionOptions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_mongo_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_rawDesc = nil
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_mongo_proto_depIdxs = nil
}
