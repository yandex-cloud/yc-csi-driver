// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/kms/v1/asymmetricsignature/asymmetric_signature_crypto_service.proto

package kms

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
	AsymmetricSignatureCryptoService_Sign_FullMethodName         = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureCryptoService/Sign"
	AsymmetricSignatureCryptoService_SignHash_FullMethodName     = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureCryptoService/SignHash"
	AsymmetricSignatureCryptoService_GetPublicKey_FullMethodName = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureCryptoService/GetPublicKey"
)

// AsymmetricSignatureCryptoServiceClient is the client API for AsymmetricSignatureCryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Set of methods that perform asymmetric signature.
type AsymmetricSignatureCryptoServiceClient interface {
	// Signs data specified KMS key.
	Sign(ctx context.Context, in *AsymmetricSignRequest, opts ...grpc.CallOption) (*AsymmetricSignResponse, error)
	// Signs hash value specified KMS key.
	SignHash(ctx context.Context, in *AsymmetricSignHashRequest, opts ...grpc.CallOption) (*AsymmetricSignHashResponse, error)
	// Gets value of public key.
	GetPublicKey(ctx context.Context, in *AsymmetricGetPublicKeyRequest, opts ...grpc.CallOption) (*AsymmetricGetPublicKeyResponse, error)
}

type asymmetricSignatureCryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAsymmetricSignatureCryptoServiceClient(cc grpc.ClientConnInterface) AsymmetricSignatureCryptoServiceClient {
	return &asymmetricSignatureCryptoServiceClient{cc}
}

func (c *asymmetricSignatureCryptoServiceClient) Sign(ctx context.Context, in *AsymmetricSignRequest, opts ...grpc.CallOption) (*AsymmetricSignResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AsymmetricSignResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureCryptoService_Sign_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureCryptoServiceClient) SignHash(ctx context.Context, in *AsymmetricSignHashRequest, opts ...grpc.CallOption) (*AsymmetricSignHashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AsymmetricSignHashResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureCryptoService_SignHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureCryptoServiceClient) GetPublicKey(ctx context.Context, in *AsymmetricGetPublicKeyRequest, opts ...grpc.CallOption) (*AsymmetricGetPublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AsymmetricGetPublicKeyResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureCryptoService_GetPublicKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsymmetricSignatureCryptoServiceServer is the server API for AsymmetricSignatureCryptoService service.
// All implementations should embed UnimplementedAsymmetricSignatureCryptoServiceServer
// for forward compatibility.
//
// Set of methods that perform asymmetric signature.
type AsymmetricSignatureCryptoServiceServer interface {
	// Signs data specified KMS key.
	Sign(context.Context, *AsymmetricSignRequest) (*AsymmetricSignResponse, error)
	// Signs hash value specified KMS key.
	SignHash(context.Context, *AsymmetricSignHashRequest) (*AsymmetricSignHashResponse, error)
	// Gets value of public key.
	GetPublicKey(context.Context, *AsymmetricGetPublicKeyRequest) (*AsymmetricGetPublicKeyResponse, error)
}

// UnimplementedAsymmetricSignatureCryptoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAsymmetricSignatureCryptoServiceServer struct{}

func (UnimplementedAsymmetricSignatureCryptoServiceServer) Sign(context.Context, *AsymmetricSignRequest) (*AsymmetricSignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedAsymmetricSignatureCryptoServiceServer) SignHash(context.Context, *AsymmetricSignHashRequest) (*AsymmetricSignHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignHash not implemented")
}
func (UnimplementedAsymmetricSignatureCryptoServiceServer) GetPublicKey(context.Context, *AsymmetricGetPublicKeyRequest) (*AsymmetricGetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedAsymmetricSignatureCryptoServiceServer) testEmbeddedByValue() {}

// UnsafeAsymmetricSignatureCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsymmetricSignatureCryptoServiceServer will
// result in compilation errors.
type UnsafeAsymmetricSignatureCryptoServiceServer interface {
	mustEmbedUnimplementedAsymmetricSignatureCryptoServiceServer()
}

func RegisterAsymmetricSignatureCryptoServiceServer(s grpc.ServiceRegistrar, srv AsymmetricSignatureCryptoServiceServer) {
	// If the following call pancis, it indicates UnimplementedAsymmetricSignatureCryptoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AsymmetricSignatureCryptoService_ServiceDesc, srv)
}

func _AsymmetricSignatureCryptoService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsymmetricSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureCryptoServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureCryptoService_Sign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureCryptoServiceServer).Sign(ctx, req.(*AsymmetricSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureCryptoService_SignHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsymmetricSignHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureCryptoServiceServer).SignHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureCryptoService_SignHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureCryptoServiceServer).SignHash(ctx, req.(*AsymmetricSignHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureCryptoService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AsymmetricGetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureCryptoServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureCryptoService_GetPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureCryptoServiceServer).GetPublicKey(ctx, req.(*AsymmetricGetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AsymmetricSignatureCryptoService_ServiceDesc is the grpc.ServiceDesc for AsymmetricSignatureCryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsymmetricSignatureCryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureCryptoService",
	HandlerType: (*AsymmetricSignatureCryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _AsymmetricSignatureCryptoService_Sign_Handler,
		},
		{
			MethodName: "SignHash",
			Handler:    _AsymmetricSignatureCryptoService_SignHash_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _AsymmetricSignatureCryptoService_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/kms/v1/asymmetricsignature/asymmetric_signature_crypto_service.proto",
}
