// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/marketplace/licensemanager/v1/instance_service.proto

package licensemanager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InstanceService_Get_FullMethodName  = "/yandex.cloud.marketplace.licensemanager.v1.InstanceService/Get"
	InstanceService_List_FullMethodName = "/yandex.cloud.marketplace.licensemanager.v1.InstanceService/List"
)

// InstanceServiceClient is the client API for InstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing subscription instances.
type InstanceServiceClient interface {
	// Returns the specified subscription instance.
	//
	// To get the list of all available subscription instances, make a [List] request.
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
	// Retrieves the list of subscription instances in the specified folder.
	List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
}

type instanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceServiceClient(cc grpc.ClientConnInterface) InstanceServiceClient {
	return &instanceServiceClient{cc}
}

func (c *instanceServiceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Instance)
	err := c.cc.Invoke(ctx, InstanceService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, InstanceService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServiceServer is the server API for InstanceService service.
// All implementations should embed UnimplementedInstanceServiceServer
// for forward compatibility.
//
// A set of methods for managing subscription instances.
type InstanceServiceServer interface {
	// Returns the specified subscription instance.
	//
	// To get the list of all available subscription instances, make a [List] request.
	Get(context.Context, *GetInstanceRequest) (*Instance, error)
	// Retrieves the list of subscription instances in the specified folder.
	List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
}

// UnimplementedInstanceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInstanceServiceServer struct{}

func (UnimplementedInstanceServiceServer) Get(context.Context, *GetInstanceRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInstanceServiceServer) List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedInstanceServiceServer) testEmbeddedByValue() {}

// UnsafeInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServiceServer will
// result in compilation errors.
type UnsafeInstanceServiceServer interface {
	mustEmbedUnimplementedInstanceServiceServer()
}

func RegisterInstanceServiceServer(s grpc.ServiceRegistrar, srv InstanceServiceServer) {
	// If the following call pancis, it indicates UnimplementedInstanceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InstanceService_ServiceDesc, srv)
}

func _InstanceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InstanceService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).List(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstanceService_ServiceDesc is the grpc.ServiceDesc for InstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.marketplace.licensemanager.v1.InstanceService",
	HandlerType: (*InstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InstanceService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _InstanceService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/marketplace/licensemanager/v1/instance_service.proto",
}
