// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/mdb/postgresql/v1/perf_diag_service.proto

package postgresql

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PerformanceDiagnosticsService_ListRawSessionStates_FullMethodName = "/yandex.cloud.mdb.postgresql.v1.PerformanceDiagnosticsService/ListRawSessionStates"
	PerformanceDiagnosticsService_ListRawStatements_FullMethodName    = "/yandex.cloud.mdb.postgresql.v1.PerformanceDiagnosticsService/ListRawStatements"
)

// PerformanceDiagnosticsServiceClient is the client API for PerformanceDiagnosticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PerformanceDiagnosticsServiceClient interface {
	// Retrieves raw statistics on sessions. Corresponds to the [pg_stat_activity view](https://www.postgresql.org/docs/current/monitoring-stats.html#MONITORING-PG-STAT-ACTIVITY-VIEW).
	ListRawSessionStates(ctx context.Context, in *ListRawSessionStatesRequest, opts ...grpc.CallOption) (*ListRawSessionStatesResponse, error)
	// Retrieves statistics on planning and execution of SQL statements (queries).
	ListRawStatements(ctx context.Context, in *ListRawStatementsRequest, opts ...grpc.CallOption) (*ListRawStatementsResponse, error)
}

type performanceDiagnosticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPerformanceDiagnosticsServiceClient(cc grpc.ClientConnInterface) PerformanceDiagnosticsServiceClient {
	return &performanceDiagnosticsServiceClient{cc}
}

func (c *performanceDiagnosticsServiceClient) ListRawSessionStates(ctx context.Context, in *ListRawSessionStatesRequest, opts ...grpc.CallOption) (*ListRawSessionStatesResponse, error) {
	out := new(ListRawSessionStatesResponse)
	err := c.cc.Invoke(ctx, PerformanceDiagnosticsService_ListRawSessionStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *performanceDiagnosticsServiceClient) ListRawStatements(ctx context.Context, in *ListRawStatementsRequest, opts ...grpc.CallOption) (*ListRawStatementsResponse, error) {
	out := new(ListRawStatementsResponse)
	err := c.cc.Invoke(ctx, PerformanceDiagnosticsService_ListRawStatements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerformanceDiagnosticsServiceServer is the server API for PerformanceDiagnosticsService service.
// All implementations should embed UnimplementedPerformanceDiagnosticsServiceServer
// for forward compatibility
type PerformanceDiagnosticsServiceServer interface {
	// Retrieves raw statistics on sessions. Corresponds to the [pg_stat_activity view](https://www.postgresql.org/docs/current/monitoring-stats.html#MONITORING-PG-STAT-ACTIVITY-VIEW).
	ListRawSessionStates(context.Context, *ListRawSessionStatesRequest) (*ListRawSessionStatesResponse, error)
	// Retrieves statistics on planning and execution of SQL statements (queries).
	ListRawStatements(context.Context, *ListRawStatementsRequest) (*ListRawStatementsResponse, error)
}

// UnimplementedPerformanceDiagnosticsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPerformanceDiagnosticsServiceServer struct {
}

func (UnimplementedPerformanceDiagnosticsServiceServer) ListRawSessionStates(context.Context, *ListRawSessionStatesRequest) (*ListRawSessionStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRawSessionStates not implemented")
}
func (UnimplementedPerformanceDiagnosticsServiceServer) ListRawStatements(context.Context, *ListRawStatementsRequest) (*ListRawStatementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRawStatements not implemented")
}

// UnsafePerformanceDiagnosticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PerformanceDiagnosticsServiceServer will
// result in compilation errors.
type UnsafePerformanceDiagnosticsServiceServer interface {
	mustEmbedUnimplementedPerformanceDiagnosticsServiceServer()
}

func RegisterPerformanceDiagnosticsServiceServer(s grpc.ServiceRegistrar, srv PerformanceDiagnosticsServiceServer) {
	s.RegisterService(&PerformanceDiagnosticsService_ServiceDesc, srv)
}

func _PerformanceDiagnosticsService_ListRawSessionStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRawSessionStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceDiagnosticsServiceServer).ListRawSessionStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceDiagnosticsService_ListRawSessionStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceDiagnosticsServiceServer).ListRawSessionStates(ctx, req.(*ListRawSessionStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerformanceDiagnosticsService_ListRawStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRawStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerformanceDiagnosticsServiceServer).ListRawStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerformanceDiagnosticsService_ListRawStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerformanceDiagnosticsServiceServer).ListRawStatements(ctx, req.(*ListRawStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PerformanceDiagnosticsService_ServiceDesc is the grpc.ServiceDesc for PerformanceDiagnosticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PerformanceDiagnosticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.postgresql.v1.PerformanceDiagnosticsService",
	HandlerType: (*PerformanceDiagnosticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRawSessionStates",
			Handler:    _PerformanceDiagnosticsService_ListRawSessionStates_Handler,
		},
		{
			MethodName: "ListRawStatements",
			Handler:    _PerformanceDiagnosticsService_ListRawStatements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/postgresql/v1/perf_diag_service.proto",
}
