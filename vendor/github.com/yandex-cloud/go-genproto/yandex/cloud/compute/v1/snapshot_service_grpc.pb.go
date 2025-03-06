// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/compute/v1/snapshot_service.proto

package compute

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
	SnapshotService_Get_FullMethodName                  = "/yandex.cloud.compute.v1.SnapshotService/Get"
	SnapshotService_List_FullMethodName                 = "/yandex.cloud.compute.v1.SnapshotService/List"
	SnapshotService_Create_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotService/Create"
	SnapshotService_Update_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotService/Update"
	SnapshotService_Delete_FullMethodName               = "/yandex.cloud.compute.v1.SnapshotService/Delete"
	SnapshotService_ListOperations_FullMethodName       = "/yandex.cloud.compute.v1.SnapshotService/ListOperations"
	SnapshotService_ListAccessBindings_FullMethodName   = "/yandex.cloud.compute.v1.SnapshotService/ListAccessBindings"
	SnapshotService_SetAccessBindings_FullMethodName    = "/yandex.cloud.compute.v1.SnapshotService/SetAccessBindings"
	SnapshotService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.compute.v1.SnapshotService/UpdateAccessBindings"
)

// SnapshotServiceClient is the client API for SnapshotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Snapshot resources.
type SnapshotServiceClient interface {
	// Returns the specified Snapshot resource.
	//
	// To get the list of available Snapshot resources, make a [List] request.
	Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error)
	// Retrieves the list of Snapshot resources in the specified folder.
	List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
	// Creates a snapshot of the specified disk.
	Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified snapshot.
	//
	// Values of omitted parameters are not changed.
	Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified snapshot.
	//
	// Deleting a snapshot removes its data permanently and is irreversible.
	Delete(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified snapshot.
	ListOperations(ctx context.Context, in *ListSnapshotOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotOperationsResponse, error)
	// Lists access bindings for the snapshot.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the snapshot.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the snapshot.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type snapshotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotServiceClient(cc grpc.ClientConnInterface) SnapshotServiceClient {
	return &snapshotServiceClient{cc}
}

func (c *snapshotServiceClient) Get(ctx context.Context, in *GetSnapshotRequest, opts ...grpc.CallOption) (*Snapshot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Snapshot)
	err := c.cc.Invoke(ctx, SnapshotService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, SnapshotService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Update(ctx context.Context, in *UpdateSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) Delete(ctx context.Context, in *DeleteSnapshotRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) ListOperations(ctx context.Context, in *ListSnapshotOperationsRequest, opts ...grpc.CallOption) (*ListSnapshotOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSnapshotOperationsResponse)
	err := c.cc.Invoke(ctx, SnapshotService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, SnapshotService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SnapshotService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotServiceServer is the server API for SnapshotService service.
// All implementations should embed UnimplementedSnapshotServiceServer
// for forward compatibility.
//
// A set of methods for managing Snapshot resources.
type SnapshotServiceServer interface {
	// Returns the specified Snapshot resource.
	//
	// To get the list of available Snapshot resources, make a [List] request.
	Get(context.Context, *GetSnapshotRequest) (*Snapshot, error)
	// Retrieves the list of Snapshot resources in the specified folder.
	List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	// Creates a snapshot of the specified disk.
	Create(context.Context, *CreateSnapshotRequest) (*operation.Operation, error)
	// Updates the specified snapshot.
	//
	// Values of omitted parameters are not changed.
	Update(context.Context, *UpdateSnapshotRequest) (*operation.Operation, error)
	// Deletes the specified snapshot.
	//
	// Deleting a snapshot removes its data permanently and is irreversible.
	Delete(context.Context, *DeleteSnapshotRequest) (*operation.Operation, error)
	// Lists operations for the specified snapshot.
	ListOperations(context.Context, *ListSnapshotOperationsRequest) (*ListSnapshotOperationsResponse, error)
	// Lists access bindings for the snapshot.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the snapshot.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the snapshot.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedSnapshotServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSnapshotServiceServer struct{}

func (UnimplementedSnapshotServiceServer) Get(context.Context, *GetSnapshotRequest) (*Snapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSnapshotServiceServer) List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnapshotServiceServer) Create(context.Context, *CreateSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnapshotServiceServer) Update(context.Context, *UpdateSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnapshotServiceServer) Delete(context.Context, *DeleteSnapshotRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSnapshotServiceServer) ListOperations(context.Context, *ListSnapshotOperationsRequest) (*ListSnapshotOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedSnapshotServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedSnapshotServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedSnapshotServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedSnapshotServiceServer) testEmbeddedByValue() {}

// UnsafeSnapshotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotServiceServer will
// result in compilation errors.
type UnsafeSnapshotServiceServer interface {
	mustEmbedUnimplementedSnapshotServiceServer()
}

func RegisterSnapshotServiceServer(s grpc.ServiceRegistrar, srv SnapshotServiceServer) {
	// If the following call pancis, it indicates UnimplementedSnapshotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SnapshotService_ServiceDesc, srv)
}

func _SnapshotService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Get(ctx, req.(*GetSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).List(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Create(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Update(ctx, req.(*UpdateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).Delete(ctx, req.(*DeleteSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).ListOperations(ctx, req.(*ListSnapshotOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnapshotService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnapshotService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnapshotService_ServiceDesc is the grpc.ServiceDesc for SnapshotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnapshotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.SnapshotService",
	HandlerType: (*SnapshotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SnapshotService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SnapshotService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SnapshotService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SnapshotService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SnapshotService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _SnapshotService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _SnapshotService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _SnapshotService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _SnapshotService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/snapshot_service.proto",
}
