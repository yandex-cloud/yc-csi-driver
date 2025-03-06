// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/k8s/marketplace/v1/helm_release_service.proto

package k8s_marketplace

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
	HelmReleaseService_List_FullMethodName      = "/yandex.cloud.k8s.marketplace.v1.HelmReleaseService/List"
	HelmReleaseService_Get_FullMethodName       = "/yandex.cloud.k8s.marketplace.v1.HelmReleaseService/Get"
	HelmReleaseService_Install_FullMethodName   = "/yandex.cloud.k8s.marketplace.v1.HelmReleaseService/Install"
	HelmReleaseService_Update_FullMethodName    = "/yandex.cloud.k8s.marketplace.v1.HelmReleaseService/Update"
	HelmReleaseService_Uninstall_FullMethodName = "/yandex.cloud.k8s.marketplace.v1.HelmReleaseService/Uninstall"
)

// HelmReleaseServiceClient is the client API for HelmReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Helm release.
type HelmReleaseServiceClient interface {
	// Retrieves the list of Helm releases in the specified Kubernetes Cluster.
	List(ctx context.Context, in *ListHelmReleasesRequest, opts ...grpc.CallOption) (*ListHelmReleasesResponse, error)
	// Returns the specified Helm release.
	Get(ctx context.Context, in *GetHelmReleaseRequest, opts ...grpc.CallOption) (*HelmRelease, error)
	// Installs helm release into specified Kubernetes Cluster.
	Install(ctx context.Context, in *InstallHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates helm release.
	Update(ctx context.Context, in *UpdateHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Uninstalls helm release.
	Uninstall(ctx context.Context, in *UninstallHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type helmReleaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelmReleaseServiceClient(cc grpc.ClientConnInterface) HelmReleaseServiceClient {
	return &helmReleaseServiceClient{cc}
}

func (c *helmReleaseServiceClient) List(ctx context.Context, in *ListHelmReleasesRequest, opts ...grpc.CallOption) (*ListHelmReleasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListHelmReleasesResponse)
	err := c.cc.Invoke(ctx, HelmReleaseService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseServiceClient) Get(ctx context.Context, in *GetHelmReleaseRequest, opts ...grpc.CallOption) (*HelmRelease, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelmRelease)
	err := c.cc.Invoke(ctx, HelmReleaseService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseServiceClient) Install(ctx context.Context, in *InstallHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, HelmReleaseService_Install_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseServiceClient) Update(ctx context.Context, in *UpdateHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, HelmReleaseService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helmReleaseServiceClient) Uninstall(ctx context.Context, in *UninstallHelmReleaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, HelmReleaseService_Uninstall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelmReleaseServiceServer is the server API for HelmReleaseService service.
// All implementations should embed UnimplementedHelmReleaseServiceServer
// for forward compatibility.
//
// A set of methods for managing Helm release.
type HelmReleaseServiceServer interface {
	// Retrieves the list of Helm releases in the specified Kubernetes Cluster.
	List(context.Context, *ListHelmReleasesRequest) (*ListHelmReleasesResponse, error)
	// Returns the specified Helm release.
	Get(context.Context, *GetHelmReleaseRequest) (*HelmRelease, error)
	// Installs helm release into specified Kubernetes Cluster.
	Install(context.Context, *InstallHelmReleaseRequest) (*operation.Operation, error)
	// Updates helm release.
	Update(context.Context, *UpdateHelmReleaseRequest) (*operation.Operation, error)
	// Uninstalls helm release.
	Uninstall(context.Context, *UninstallHelmReleaseRequest) (*operation.Operation, error)
}

// UnimplementedHelmReleaseServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelmReleaseServiceServer struct{}

func (UnimplementedHelmReleaseServiceServer) List(context.Context, *ListHelmReleasesRequest) (*ListHelmReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHelmReleaseServiceServer) Get(context.Context, *GetHelmReleaseRequest) (*HelmRelease, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHelmReleaseServiceServer) Install(context.Context, *InstallHelmReleaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedHelmReleaseServiceServer) Update(context.Context, *UpdateHelmReleaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedHelmReleaseServiceServer) Uninstall(context.Context, *UninstallHelmReleaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedHelmReleaseServiceServer) testEmbeddedByValue() {}

// UnsafeHelmReleaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelmReleaseServiceServer will
// result in compilation errors.
type UnsafeHelmReleaseServiceServer interface {
	mustEmbedUnimplementedHelmReleaseServiceServer()
}

func RegisterHelmReleaseServiceServer(s grpc.ServiceRegistrar, srv HelmReleaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedHelmReleaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HelmReleaseService_ServiceDesc, srv)
}

func _HelmReleaseService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHelmReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseServiceServer).List(ctx, req.(*ListHelmReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelmReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseServiceServer).Get(ctx, req.(*GetHelmReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseService_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallHelmReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseServiceServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseService_Install_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseServiceServer).Install(ctx, req.(*InstallHelmReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHelmReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseServiceServer).Update(ctx, req.(*UpdateHelmReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelmReleaseService_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallHelmReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelmReleaseServiceServer).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelmReleaseService_Uninstall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelmReleaseServiceServer).Uninstall(ctx, req.(*UninstallHelmReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelmReleaseService_ServiceDesc is the grpc.ServiceDesc for HelmReleaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelmReleaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.k8s.marketplace.v1.HelmReleaseService",
	HandlerType: (*HelmReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _HelmReleaseService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HelmReleaseService_Get_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _HelmReleaseService_Install_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _HelmReleaseService_Update_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _HelmReleaseService_Uninstall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/k8s/marketplace/v1/helm_release_service.proto",
}
