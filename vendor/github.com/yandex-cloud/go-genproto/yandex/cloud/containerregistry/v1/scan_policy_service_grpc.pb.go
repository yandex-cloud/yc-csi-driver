// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/containerregistry/v1/scan_policy_service.proto

package containerregistry

import (
	context "context"
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
	ScanPolicyService_Get_FullMethodName           = "/yandex.cloud.containerregistry.v1.ScanPolicyService/Get"
	ScanPolicyService_GetByRegistry_FullMethodName = "/yandex.cloud.containerregistry.v1.ScanPolicyService/GetByRegistry"
	ScanPolicyService_Create_FullMethodName        = "/yandex.cloud.containerregistry.v1.ScanPolicyService/Create"
	ScanPolicyService_Update_FullMethodName        = "/yandex.cloud.containerregistry.v1.ScanPolicyService/Update"
	ScanPolicyService_Delete_FullMethodName        = "/yandex.cloud.containerregistry.v1.ScanPolicyService/Delete"
)

// ScanPolicyServiceClient is the client API for ScanPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing scan policy resources.
type ScanPolicyServiceClient interface {
	// Returns the specified scan policy.
	Get(ctx context.Context, in *GetScanPolicyRequest, opts ...grpc.CallOption) (*ScanPolicy, error)
	// Returns scan policy for the registry if any exists.
	GetByRegistry(ctx context.Context, in *GetScanPolicyByRegistryRequest, opts ...grpc.CallOption) (*ScanPolicy, error)
	// Creates a scan policy for the specified registry.
	Create(ctx context.Context, in *CreateScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified scan policy.
	Update(ctx context.Context, in *UpdateScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified scan policy.
	Delete(ctx context.Context, in *DeleteScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type scanPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScanPolicyServiceClient(cc grpc.ClientConnInterface) ScanPolicyServiceClient {
	return &scanPolicyServiceClient{cc}
}

func (c *scanPolicyServiceClient) Get(ctx context.Context, in *GetScanPolicyRequest, opts ...grpc.CallOption) (*ScanPolicy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScanPolicy)
	err := c.cc.Invoke(ctx, ScanPolicyService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanPolicyServiceClient) GetByRegistry(ctx context.Context, in *GetScanPolicyByRegistryRequest, opts ...grpc.CallOption) (*ScanPolicy, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScanPolicy)
	err := c.cc.Invoke(ctx, ScanPolicyService_GetByRegistry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanPolicyServiceClient) Create(ctx context.Context, in *CreateScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ScanPolicyService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanPolicyServiceClient) Update(ctx context.Context, in *UpdateScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ScanPolicyService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scanPolicyServiceClient) Delete(ctx context.Context, in *DeleteScanPolicyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ScanPolicyService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScanPolicyServiceServer is the server API for ScanPolicyService service.
// All implementations should embed UnimplementedScanPolicyServiceServer
// for forward compatibility.
//
// A set of methods for managing scan policy resources.
type ScanPolicyServiceServer interface {
	// Returns the specified scan policy.
	Get(context.Context, *GetScanPolicyRequest) (*ScanPolicy, error)
	// Returns scan policy for the registry if any exists.
	GetByRegistry(context.Context, *GetScanPolicyByRegistryRequest) (*ScanPolicy, error)
	// Creates a scan policy for the specified registry.
	Create(context.Context, *CreateScanPolicyRequest) (*operation.Operation, error)
	// Updates the specified scan policy.
	Update(context.Context, *UpdateScanPolicyRequest) (*operation.Operation, error)
	// Deletes the specified scan policy.
	Delete(context.Context, *DeleteScanPolicyRequest) (*operation.Operation, error)
}

// UnimplementedScanPolicyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScanPolicyServiceServer struct{}

func (UnimplementedScanPolicyServiceServer) Get(context.Context, *GetScanPolicyRequest) (*ScanPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedScanPolicyServiceServer) GetByRegistry(context.Context, *GetScanPolicyByRegistryRequest) (*ScanPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByRegistry not implemented")
}
func (UnimplementedScanPolicyServiceServer) Create(context.Context, *CreateScanPolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedScanPolicyServiceServer) Update(context.Context, *UpdateScanPolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedScanPolicyServiceServer) Delete(context.Context, *DeleteScanPolicyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedScanPolicyServiceServer) testEmbeddedByValue() {}

// UnsafeScanPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScanPolicyServiceServer will
// result in compilation errors.
type UnsafeScanPolicyServiceServer interface {
	mustEmbedUnimplementedScanPolicyServiceServer()
}

func RegisterScanPolicyServiceServer(s grpc.ServiceRegistrar, srv ScanPolicyServiceServer) {
	// If the following call pancis, it indicates UnimplementedScanPolicyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScanPolicyService_ServiceDesc, srv)
}

func _ScanPolicyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScanPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanPolicyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanPolicyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanPolicyServiceServer).Get(ctx, req.(*GetScanPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanPolicyService_GetByRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScanPolicyByRegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanPolicyServiceServer).GetByRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanPolicyService_GetByRegistry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanPolicyServiceServer).GetByRegistry(ctx, req.(*GetScanPolicyByRegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanPolicyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScanPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanPolicyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanPolicyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanPolicyServiceServer).Create(ctx, req.(*CreateScanPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanPolicyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScanPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanPolicyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanPolicyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanPolicyServiceServer).Update(ctx, req.(*UpdateScanPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScanPolicyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScanPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScanPolicyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScanPolicyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScanPolicyServiceServer).Delete(ctx, req.(*DeleteScanPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScanPolicyService_ServiceDesc is the grpc.ServiceDesc for ScanPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScanPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.containerregistry.v1.ScanPolicyService",
	HandlerType: (*ScanPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ScanPolicyService_Get_Handler,
		},
		{
			MethodName: "GetByRegistry",
			Handler:    _ScanPolicyService_GetByRegistry_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ScanPolicyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ScanPolicyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ScanPolicyService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/containerregistry/v1/scan_policy_service.proto",
}
