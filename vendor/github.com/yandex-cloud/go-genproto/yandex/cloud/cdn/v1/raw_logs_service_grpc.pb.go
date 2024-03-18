// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/cdn/v1/raw_logs_service.proto

package cdn

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RawLogsService_Activate_FullMethodName   = "/yandex.cloud.cdn.v1.RawLogsService/Activate"
	RawLogsService_Deactivate_FullMethodName = "/yandex.cloud.cdn.v1.RawLogsService/Deactivate"
	RawLogsService_Get_FullMethodName        = "/yandex.cloud.cdn.v1.RawLogsService/Get"
	RawLogsService_Update_FullMethodName     = "/yandex.cloud.cdn.v1.RawLogsService/Update"
)

// RawLogsServiceClient is the client API for RawLogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RawLogsServiceClient interface {
	Activate(ctx context.Context, in *ActivateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Deactivate(ctx context.Context, in *DeactivateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Get(ctx context.Context, in *GetRawLogsRequest, opts ...grpc.CallOption) (*GetRawLogsResponse, error)
	Update(ctx context.Context, in *UpdateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type rawLogsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRawLogsServiceClient(cc grpc.ClientConnInterface) RawLogsServiceClient {
	return &rawLogsServiceClient{cc}
}

func (c *rawLogsServiceClient) Activate(ctx context.Context, in *ActivateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RawLogsService_Activate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawLogsServiceClient) Deactivate(ctx context.Context, in *DeactivateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RawLogsService_Deactivate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawLogsServiceClient) Get(ctx context.Context, in *GetRawLogsRequest, opts ...grpc.CallOption) (*GetRawLogsResponse, error) {
	out := new(GetRawLogsResponse)
	err := c.cc.Invoke(ctx, RawLogsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawLogsServiceClient) Update(ctx context.Context, in *UpdateRawLogsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RawLogsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawLogsServiceServer is the server API for RawLogsService service.
// All implementations should embed UnimplementedRawLogsServiceServer
// for forward compatibility
type RawLogsServiceServer interface {
	Activate(context.Context, *ActivateRawLogsRequest) (*operation.Operation, error)
	Deactivate(context.Context, *DeactivateRawLogsRequest) (*operation.Operation, error)
	Get(context.Context, *GetRawLogsRequest) (*GetRawLogsResponse, error)
	Update(context.Context, *UpdateRawLogsRequest) (*operation.Operation, error)
}

// UnimplementedRawLogsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRawLogsServiceServer struct {
}

func (UnimplementedRawLogsServiceServer) Activate(context.Context, *ActivateRawLogsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedRawLogsServiceServer) Deactivate(context.Context, *DeactivateRawLogsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedRawLogsServiceServer) Get(context.Context, *GetRawLogsRequest) (*GetRawLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRawLogsServiceServer) Update(context.Context, *UpdateRawLogsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeRawLogsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RawLogsServiceServer will
// result in compilation errors.
type UnsafeRawLogsServiceServer interface {
	mustEmbedUnimplementedRawLogsServiceServer()
}

func RegisterRawLogsServiceServer(s grpc.ServiceRegistrar, srv RawLogsServiceServer) {
	s.RegisterService(&RawLogsService_ServiceDesc, srv)
}

func _RawLogsService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRawLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawLogsServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RawLogsService_Activate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawLogsServiceServer).Activate(ctx, req.(*ActivateRawLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RawLogsService_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateRawLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawLogsServiceServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RawLogsService_Deactivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawLogsServiceServer).Deactivate(ctx, req.(*DeactivateRawLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RawLogsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRawLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawLogsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RawLogsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawLogsServiceServer).Get(ctx, req.(*GetRawLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RawLogsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRawLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawLogsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RawLogsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawLogsServiceServer).Update(ctx, req.(*UpdateRawLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RawLogsService_ServiceDesc is the grpc.ServiceDesc for RawLogsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RawLogsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.cdn.v1.RawLogsService",
	HandlerType: (*RawLogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _RawLogsService_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _RawLogsService_Deactivate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RawLogsService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RawLogsService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/cdn/v1/raw_logs_service.proto",
}