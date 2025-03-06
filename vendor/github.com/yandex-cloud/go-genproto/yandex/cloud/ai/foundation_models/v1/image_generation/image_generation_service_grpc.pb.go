// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/ai/foundation_models/v1/image_generation/image_generation_service.proto

package image_generation

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
	ImageGenerationAsyncService_Generate_FullMethodName = "/yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationAsyncService/Generate"
)

// ImageGenerationAsyncServiceClient is the client API for ImageGenerationAsyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for creating images based on a text description.
type ImageGenerationAsyncServiceClient interface {
	// A method for generating an image based on a textual description.
	Generate(ctx context.Context, in *ImageGenerationRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type imageGenerationAsyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageGenerationAsyncServiceClient(cc grpc.ClientConnInterface) ImageGenerationAsyncServiceClient {
	return &imageGenerationAsyncServiceClient{cc}
}

func (c *imageGenerationAsyncServiceClient) Generate(ctx context.Context, in *ImageGenerationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ImageGenerationAsyncService_Generate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageGenerationAsyncServiceServer is the server API for ImageGenerationAsyncService service.
// All implementations should embed UnimplementedImageGenerationAsyncServiceServer
// for forward compatibility.
//
// Service for creating images based on a text description.
type ImageGenerationAsyncServiceServer interface {
	// A method for generating an image based on a textual description.
	Generate(context.Context, *ImageGenerationRequest) (*operation.Operation, error)
}

// UnimplementedImageGenerationAsyncServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageGenerationAsyncServiceServer struct{}

func (UnimplementedImageGenerationAsyncServiceServer) Generate(context.Context, *ImageGenerationRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedImageGenerationAsyncServiceServer) testEmbeddedByValue() {}

// UnsafeImageGenerationAsyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageGenerationAsyncServiceServer will
// result in compilation errors.
type UnsafeImageGenerationAsyncServiceServer interface {
	mustEmbedUnimplementedImageGenerationAsyncServiceServer()
}

func RegisterImageGenerationAsyncServiceServer(s grpc.ServiceRegistrar, srv ImageGenerationAsyncServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageGenerationAsyncServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageGenerationAsyncService_ServiceDesc, srv)
}

func _ImageGenerationAsyncService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageGenerationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageGenerationAsyncServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageGenerationAsyncService_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageGenerationAsyncServiceServer).Generate(ctx, req.(*ImageGenerationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageGenerationAsyncService_ServiceDesc is the grpc.ServiceDesc for ImageGenerationAsyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageGenerationAsyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationAsyncService",
	HandlerType: (*ImageGenerationAsyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _ImageGenerationAsyncService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/foundation_models/v1/image_generation/image_generation_service.proto",
}
