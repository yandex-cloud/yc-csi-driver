// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/loadtesting/api/v1/test_service.proto

package loadtesting

import (
	context "context"
	test "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/test"
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
	TestService_Create_FullMethodName = "/yandex.cloud.loadtesting.api.v1.TestService/Create"
	TestService_Get_FullMethodName    = "/yandex.cloud.loadtesting.api.v1.TestService/Get"
	TestService_Stop_FullMethodName   = "/yandex.cloud.loadtesting.api.v1.TestService/Stop"
	TestService_Delete_FullMethodName = "/yandex.cloud.loadtesting.api.v1.TestService/Delete"
	TestService_List_FullMethodName   = "/yandex.cloud.loadtesting.api.v1.TestService/List"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing tests.
type TestServiceClient interface {
	// Creates (runs) a test in the specified folder.
	Create(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified test.
	//
	// To get the list of all available tests, make a [List] request.
	Get(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*test.Test, error)
	// Stops the specified test.
	Stop(ctx context.Context, in *StopTestRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes specified tests.
	Delete(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of test in the specified folder.
	List(ctx context.Context, in *ListTestsRequest, opts ...grpc.CallOption) (*ListTestsResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Create(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TestService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Get(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*test.Test, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(test.Test)
	err := c.cc.Invoke(ctx, TestService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Stop(ctx context.Context, in *StopTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TestService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Delete(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TestService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) List(ctx context.Context, in *ListTestsRequest, opts ...grpc.CallOption) (*ListTestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTestsResponse)
	err := c.cc.Invoke(ctx, TestService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations should embed UnimplementedTestServiceServer
// for forward compatibility.
//
// A set of methods for managing tests.
type TestServiceServer interface {
	// Creates (runs) a test in the specified folder.
	Create(context.Context, *CreateTestRequest) (*operation.Operation, error)
	// Returns the specified test.
	//
	// To get the list of all available tests, make a [List] request.
	Get(context.Context, *GetTestRequest) (*test.Test, error)
	// Stops the specified test.
	Stop(context.Context, *StopTestRequest) (*operation.Operation, error)
	// Deletes specified tests.
	Delete(context.Context, *DeleteTestRequest) (*operation.Operation, error)
	// Retrieves the list of test in the specified folder.
	List(context.Context, *ListTestsRequest) (*ListTestsResponse, error)
}

// UnimplementedTestServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) Create(context.Context, *CreateTestRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTestServiceServer) Get(context.Context, *GetTestRequest) (*test.Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTestServiceServer) Stop(context.Context, *StopTestRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTestServiceServer) Delete(context.Context, *DeleteTestRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTestServiceServer) List(context.Context, *ListTestsRequest) (*ListTestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTestServiceServer) testEmbeddedByValue() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Create(ctx, req.(*CreateTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Get(ctx, req.(*GetTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Stop(ctx, req.(*StopTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Delete(ctx, req.(*DeleteTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).List(ctx, req.(*ListTestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.loadtesting.api.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TestService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TestService_Get_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TestService_Stop_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TestService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TestService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/loadtesting/api/v1/test_service.proto",
}
