// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/loadtesting/agent/v1/agent_registration_service.proto

package agent

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputeInstanceId string `protobuf:"bytes,1,opt,name=compute_instance_id,json=computeInstanceId,proto3" json:"compute_instance_id,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetComputeInstanceId() string {
	if x != nil {
		return x.ComputeInstanceId
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentInstanceId string `protobuf:"bytes,1,opt,name=agent_instance_id,json=agentInstanceId,proto3" json:"agent_instance_id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetAgentInstanceId() string {
	if x != nil {
		return x.AgentInstanceId
	}
	return ""
}

type ExternalAgentRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FolderId          string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	ComputeInstanceId string `protobuf:"bytes,2,opt,name=compute_instance_id,json=computeInstanceId,proto3" json:"compute_instance_id,omitempty"`
	Name              string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AgentVersion      string `protobuf:"bytes,4,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
}

func (x *ExternalAgentRegisterRequest) Reset() {
	*x = ExternalAgentRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAgentRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAgentRegisterRequest) ProtoMessage() {}

func (x *ExternalAgentRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAgentRegisterRequest.ProtoReflect.Descriptor instead.
func (*ExternalAgentRegisterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExternalAgentRegisterRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ExternalAgentRegisterRequest) GetComputeInstanceId() string {
	if x != nil {
		return x.ComputeInstanceId
	}
	return ""
}

func (x *ExternalAgentRegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalAgentRegisterRequest) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

type ExternalAgentRegisterMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentInstanceId string `protobuf:"bytes,1,opt,name=agent_instance_id,json=agentInstanceId,proto3" json:"agent_instance_id,omitempty"`
}

func (x *ExternalAgentRegisterMetadata) Reset() {
	*x = ExternalAgentRegisterMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalAgentRegisterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalAgentRegisterMetadata) ProtoMessage() {}

func (x *ExternalAgentRegisterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalAgentRegisterMetadata.ProtoReflect.Descriptor instead.
func (*ExternalAgentRegisterMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalAgentRegisterMetadata) GetAgentInstanceId() string {
	if x != nil {
		return x.AgentInstanceId
	}
	return ""
}

var File_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDesc = []byte{
	0x0a, 0x42, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x41, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x3e, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x1c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x1d, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x32, 0xb2, 0x03, 0x0a, 0x18, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0xed, 0x01, 0x0a, 0x15,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0xb2, 0xd2, 0x2a, 0x2e, 0x0a,
	0x1d, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x22, 0x33, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x74, 0x0a, 0x25, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescData = file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDesc
)

func file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDescData
}

var file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),               // 0: yandex.cloud.loadtesting.agent.v1.RegisterRequest
	(*RegisterResponse)(nil),              // 1: yandex.cloud.loadtesting.agent.v1.RegisterResponse
	(*ExternalAgentRegisterRequest)(nil),  // 2: yandex.cloud.loadtesting.agent.v1.ExternalAgentRegisterRequest
	(*ExternalAgentRegisterMetadata)(nil), // 3: yandex.cloud.loadtesting.agent.v1.ExternalAgentRegisterMetadata
	(*operation.Operation)(nil),           // 4: yandex.cloud.operation.Operation
}
var file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.loadtesting.agent.v1.AgentRegistrationService.Register:input_type -> yandex.cloud.loadtesting.agent.v1.RegisterRequest
	2, // 1: yandex.cloud.loadtesting.agent.v1.AgentRegistrationService.ExternalAgentRegister:input_type -> yandex.cloud.loadtesting.agent.v1.ExternalAgentRegisterRequest
	1, // 2: yandex.cloud.loadtesting.agent.v1.AgentRegistrationService.Register:output_type -> yandex.cloud.loadtesting.agent.v1.RegisterResponse
	4, // 3: yandex.cloud.loadtesting.agent.v1.AgentRegistrationService.ExternalAgentRegister:output_type -> yandex.cloud.operation.Operation
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_init() }
func file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_init() {
	if File_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAgentRegisterRequest); i {
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
		file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalAgentRegisterMetadata); i {
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
			RawDescriptor: file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto = out.File
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_goTypes = nil
	file_yandex_cloud_loadtesting_agent_v1_agent_registration_service_proto_depIdxs = nil
}