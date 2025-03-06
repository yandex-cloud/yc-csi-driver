// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/kms/v1/asymmetricencryption/asymmetric_encryption_key_service.proto

package kms

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AsymmetricEncryptionKeyService_Create_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/Create"
	AsymmetricEncryptionKeyService_Get_FullMethodName                  = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/Get"
	AsymmetricEncryptionKeyService_List_FullMethodName                 = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/List"
	AsymmetricEncryptionKeyService_Update_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/Update"
	AsymmetricEncryptionKeyService_Delete_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/Delete"
	AsymmetricEncryptionKeyService_ListOperations_FullMethodName       = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/ListOperations"
	AsymmetricEncryptionKeyService_ListAccessBindings_FullMethodName   = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/ListAccessBindings"
	AsymmetricEncryptionKeyService_SetAccessBindings_FullMethodName    = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/SetAccessBindings"
	AsymmetricEncryptionKeyService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService/UpdateAccessBindings"
)

// AsymmetricEncryptionKeyServiceClient is the client API for AsymmetricEncryptionKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Set of methods for managing asymmetric KMS keys.
type AsymmetricEncryptionKeyServiceClient interface {
	// control plane
	// Creates an asymmetric KMS key in the specified folder.
	Create(ctx context.Context, in *CreateAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified asymmetric KMS key.
	//
	//	To get the list of available asymmetric KMS keys, make a [SymmetricKeyService.List] request.
	Get(ctx context.Context, in *GetAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*AsymmetricEncryptionKey, error)
	// Returns the list of asymmetric KMS keys in the specified folder.
	List(ctx context.Context, in *ListAsymmetricEncryptionKeysRequest, opts ...grpc.CallOption) (*ListAsymmetricEncryptionKeysResponse, error)
	// Updates the specified asymmetric KMS key.
	Update(ctx context.Context, in *UpdateAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified asymmetric KMS key. This action also automatically schedules
	// the destruction of all of the key's versions in 72 hours.
	//
	// The key and its versions appear absent in [AsymmetricEncryptionKeyService.Get] and [AsymmetricEncryptionKeyService.List]
	// requests, but can be restored within 72 hours with a request to tech support.
	Delete(ctx context.Context, in *DeleteAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified asymmetric KMS key.
	ListOperations(ctx context.Context, in *ListAsymmetricEncryptionKeyOperationsRequest, opts ...grpc.CallOption) (*ListAsymmetricEncryptionKeyOperationsResponse, error)
	// Lists existing access bindings for the specified key.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the key.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified key.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type asymmetricEncryptionKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAsymmetricEncryptionKeyServiceClient(cc grpc.ClientConnInterface) AsymmetricEncryptionKeyServiceClient {
	return &asymmetricEncryptionKeyServiceClient{cc}
}

func (c *asymmetricEncryptionKeyServiceClient) Create(ctx context.Context, in *CreateAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) Get(ctx context.Context, in *GetAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*AsymmetricEncryptionKey, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AsymmetricEncryptionKey)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) List(ctx context.Context, in *ListAsymmetricEncryptionKeysRequest, opts ...grpc.CallOption) (*ListAsymmetricEncryptionKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAsymmetricEncryptionKeysResponse)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) Update(ctx context.Context, in *UpdateAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) Delete(ctx context.Context, in *DeleteAsymmetricEncryptionKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) ListOperations(ctx context.Context, in *ListAsymmetricEncryptionKeyOperationsRequest, opts ...grpc.CallOption) (*ListAsymmetricEncryptionKeyOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAsymmetricEncryptionKeyOperationsResponse)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricEncryptionKeyServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricEncryptionKeyService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsymmetricEncryptionKeyServiceServer is the server API for AsymmetricEncryptionKeyService service.
// All implementations should embed UnimplementedAsymmetricEncryptionKeyServiceServer
// for forward compatibility.
//
// Set of methods for managing asymmetric KMS keys.
type AsymmetricEncryptionKeyServiceServer interface {
	// control plane
	// Creates an asymmetric KMS key in the specified folder.
	Create(context.Context, *CreateAsymmetricEncryptionKeyRequest) (*operation.Operation, error)
	// Returns the specified asymmetric KMS key.
	//
	//	To get the list of available asymmetric KMS keys, make a [SymmetricKeyService.List] request.
	Get(context.Context, *GetAsymmetricEncryptionKeyRequest) (*AsymmetricEncryptionKey, error)
	// Returns the list of asymmetric KMS keys in the specified folder.
	List(context.Context, *ListAsymmetricEncryptionKeysRequest) (*ListAsymmetricEncryptionKeysResponse, error)
	// Updates the specified asymmetric KMS key.
	Update(context.Context, *UpdateAsymmetricEncryptionKeyRequest) (*operation.Operation, error)
	// Deletes the specified asymmetric KMS key. This action also automatically schedules
	// the destruction of all of the key's versions in 72 hours.
	//
	// The key and its versions appear absent in [AsymmetricEncryptionKeyService.Get] and [AsymmetricEncryptionKeyService.List]
	// requests, but can be restored within 72 hours with a request to tech support.
	Delete(context.Context, *DeleteAsymmetricEncryptionKeyRequest) (*operation.Operation, error)
	// Lists operations for the specified asymmetric KMS key.
	ListOperations(context.Context, *ListAsymmetricEncryptionKeyOperationsRequest) (*ListAsymmetricEncryptionKeyOperationsResponse, error)
	// Lists existing access bindings for the specified key.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the key.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified key.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedAsymmetricEncryptionKeyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAsymmetricEncryptionKeyServiceServer struct{}

func (UnimplementedAsymmetricEncryptionKeyServiceServer) Create(context.Context, *CreateAsymmetricEncryptionKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) Get(context.Context, *GetAsymmetricEncryptionKeyRequest) (*AsymmetricEncryptionKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) List(context.Context, *ListAsymmetricEncryptionKeysRequest) (*ListAsymmetricEncryptionKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) Update(context.Context, *UpdateAsymmetricEncryptionKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) Delete(context.Context, *DeleteAsymmetricEncryptionKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) ListOperations(context.Context, *ListAsymmetricEncryptionKeyOperationsRequest) (*ListAsymmetricEncryptionKeyOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedAsymmetricEncryptionKeyServiceServer) testEmbeddedByValue() {}

// UnsafeAsymmetricEncryptionKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsymmetricEncryptionKeyServiceServer will
// result in compilation errors.
type UnsafeAsymmetricEncryptionKeyServiceServer interface {
	mustEmbedUnimplementedAsymmetricEncryptionKeyServiceServer()
}

func RegisterAsymmetricEncryptionKeyServiceServer(s grpc.ServiceRegistrar, srv AsymmetricEncryptionKeyServiceServer) {
	// If the following call pancis, it indicates UnimplementedAsymmetricEncryptionKeyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AsymmetricEncryptionKeyService_ServiceDesc, srv)
}

func _AsymmetricEncryptionKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAsymmetricEncryptionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).Create(ctx, req.(*CreateAsymmetricEncryptionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAsymmetricEncryptionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).Get(ctx, req.(*GetAsymmetricEncryptionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAsymmetricEncryptionKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).List(ctx, req.(*ListAsymmetricEncryptionKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAsymmetricEncryptionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).Update(ctx, req.(*UpdateAsymmetricEncryptionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAsymmetricEncryptionKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).Delete(ctx, req.(*DeleteAsymmetricEncryptionKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAsymmetricEncryptionKeyOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).ListOperations(ctx, req.(*ListAsymmetricEncryptionKeyOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricEncryptionKeyService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricEncryptionKeyServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricEncryptionKeyService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricEncryptionKeyServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AsymmetricEncryptionKeyService_ServiceDesc is the grpc.ServiceDesc for AsymmetricEncryptionKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsymmetricEncryptionKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.kms.v1.asymmetricencryption.AsymmetricEncryptionKeyService",
	HandlerType: (*AsymmetricEncryptionKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AsymmetricEncryptionKeyService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AsymmetricEncryptionKeyService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AsymmetricEncryptionKeyService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AsymmetricEncryptionKeyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AsymmetricEncryptionKeyService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _AsymmetricEncryptionKeyService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _AsymmetricEncryptionKeyService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _AsymmetricEncryptionKeyService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _AsymmetricEncryptionKeyService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/kms/v1/asymmetricencryption/asymmetric_encryption_key_service.proto",
}
