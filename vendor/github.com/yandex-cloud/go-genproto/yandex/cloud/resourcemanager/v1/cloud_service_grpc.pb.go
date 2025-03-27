// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/resourcemanager/v1/cloud_service.proto

package resourcemanager

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
	CloudService_Get_FullMethodName                  = "/yandex.cloud.resourcemanager.v1.CloudService/Get"
	CloudService_List_FullMethodName                 = "/yandex.cloud.resourcemanager.v1.CloudService/List"
	CloudService_Create_FullMethodName               = "/yandex.cloud.resourcemanager.v1.CloudService/Create"
	CloudService_Update_FullMethodName               = "/yandex.cloud.resourcemanager.v1.CloudService/Update"
	CloudService_Delete_FullMethodName               = "/yandex.cloud.resourcemanager.v1.CloudService/Delete"
	CloudService_ListOperations_FullMethodName       = "/yandex.cloud.resourcemanager.v1.CloudService/ListOperations"
	CloudService_ListAccessBindings_FullMethodName   = "/yandex.cloud.resourcemanager.v1.CloudService/ListAccessBindings"
	CloudService_SetAccessBindings_FullMethodName    = "/yandex.cloud.resourcemanager.v1.CloudService/SetAccessBindings"
	CloudService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.resourcemanager.v1.CloudService/UpdateAccessBindings"
)

// CloudServiceClient is the client API for CloudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Cloud resources.
type CloudServiceClient interface {
	// Returns the specified Cloud resource.
	//
	// To get the list of available Cloud resources, make a [List] request.
	Get(ctx context.Context, in *GetCloudRequest, opts ...grpc.CallOption) (*Cloud, error)
	// Retrieves the list of Cloud resources.
	List(ctx context.Context, in *ListCloudsRequest, opts ...grpc.CallOption) (*ListCloudsResponse, error)
	// Creates a cloud in the specified organization.
	Create(ctx context.Context, in *CreateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified cloud.
	Update(ctx context.Context, in *UpdateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified cloud.
	Delete(ctx context.Context, in *DeleteCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified cloud.
	ListOperations(ctx context.Context, in *ListCloudOperationsRequest, opts ...grpc.CallOption) (*ListCloudOperationsResponse, error)
	// Lists access bindings for the specified cloud.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified cloud.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified cloud.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type cloudServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudServiceClient(cc grpc.ClientConnInterface) CloudServiceClient {
	return &cloudServiceClient{cc}
}

func (c *cloudServiceClient) Get(ctx context.Context, in *GetCloudRequest, opts ...grpc.CallOption) (*Cloud, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Cloud)
	err := c.cc.Invoke(ctx, CloudService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) List(ctx context.Context, in *ListCloudsRequest, opts ...grpc.CallOption) (*ListCloudsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCloudsResponse)
	err := c.cc.Invoke(ctx, CloudService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Create(ctx context.Context, in *CreateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CloudService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Update(ctx context.Context, in *UpdateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CloudService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Delete(ctx context.Context, in *DeleteCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CloudService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) ListOperations(ctx context.Context, in *ListCloudOperationsRequest, opts ...grpc.CallOption) (*ListCloudOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCloudOperationsResponse)
	err := c.cc.Invoke(ctx, CloudService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, CloudService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CloudService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CloudService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServiceServer is the server API for CloudService service.
// All implementations should embed UnimplementedCloudServiceServer
// for forward compatibility.
//
// A set of methods for managing Cloud resources.
type CloudServiceServer interface {
	// Returns the specified Cloud resource.
	//
	// To get the list of available Cloud resources, make a [List] request.
	Get(context.Context, *GetCloudRequest) (*Cloud, error)
	// Retrieves the list of Cloud resources.
	List(context.Context, *ListCloudsRequest) (*ListCloudsResponse, error)
	// Creates a cloud in the specified organization.
	Create(context.Context, *CreateCloudRequest) (*operation.Operation, error)
	// Updates the specified cloud.
	Update(context.Context, *UpdateCloudRequest) (*operation.Operation, error)
	// Deletes the specified cloud.
	Delete(context.Context, *DeleteCloudRequest) (*operation.Operation, error)
	// Lists operations for the specified cloud.
	ListOperations(context.Context, *ListCloudOperationsRequest) (*ListCloudOperationsResponse, error)
	// Lists access bindings for the specified cloud.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified cloud.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified cloud.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedCloudServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCloudServiceServer struct{}

func (UnimplementedCloudServiceServer) Get(context.Context, *GetCloudRequest) (*Cloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCloudServiceServer) List(context.Context, *ListCloudsRequest) (*ListCloudsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCloudServiceServer) Create(context.Context, *CreateCloudRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCloudServiceServer) Update(context.Context, *UpdateCloudRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCloudServiceServer) Delete(context.Context, *DeleteCloudRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCloudServiceServer) ListOperations(context.Context, *ListCloudOperationsRequest) (*ListCloudOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedCloudServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedCloudServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedCloudServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedCloudServiceServer) testEmbeddedByValue() {}

// UnsafeCloudServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudServiceServer will
// result in compilation errors.
type UnsafeCloudServiceServer interface {
	mustEmbedUnimplementedCloudServiceServer()
}

func RegisterCloudServiceServer(s grpc.ServiceRegistrar, srv CloudServiceServer) {
	// If the following call pancis, it indicates UnimplementedCloudServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CloudService_ServiceDesc, srv)
}

func _CloudService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Get(ctx, req.(*GetCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).List(ctx, req.(*ListCloudsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Create(ctx, req.(*CreateCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Update(ctx, req.(*UpdateCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Delete(ctx, req.(*DeleteCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).ListOperations(ctx, req.(*ListCloudOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CloudService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudService_ServiceDesc is the grpc.ServiceDesc for CloudService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.resourcemanager.v1.CloudService",
	HandlerType: (*CloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CloudService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CloudService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _CloudService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _CloudService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _CloudService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _CloudService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/resourcemanager/v1/cloud_service.proto",
}
