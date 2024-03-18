// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/ai/translate/v2/translation_service.proto

package translate

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
	TranslationService_Translate_FullMethodName      = "/yandex.cloud.ai.translate.v2.TranslationService/Translate"
	TranslationService_DetectLanguage_FullMethodName = "/yandex.cloud.ai.translate.v2.TranslationService/DetectLanguage"
	TranslationService_ListLanguages_FullMethodName  = "/yandex.cloud.ai.translate.v2.TranslationService/ListLanguages"
)

// TranslationServiceClient is the client API for TranslationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslationServiceClient interface {
	// Translates the text to the specified language.
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
	// Detects the language of the text.
	DetectLanguage(ctx context.Context, in *DetectLanguageRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error)
	// Retrieves the list of supported languages.
	ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error)
}

type translationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslationServiceClient(cc grpc.ClientConnInterface) TranslationServiceClient {
	return &translationServiceClient{cc}
}

func (c *translationServiceClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, TranslationService_Translate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) DetectLanguage(ctx context.Context, in *DetectLanguageRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error) {
	out := new(DetectLanguageResponse)
	err := c.cc.Invoke(ctx, TranslationService_DetectLanguage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error) {
	out := new(ListLanguagesResponse)
	err := c.cc.Invoke(ctx, TranslationService_ListLanguages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslationServiceServer is the server API for TranslationService service.
// All implementations should embed UnimplementedTranslationServiceServer
// for forward compatibility
type TranslationServiceServer interface {
	// Translates the text to the specified language.
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
	// Detects the language of the text.
	DetectLanguage(context.Context, *DetectLanguageRequest) (*DetectLanguageResponse, error)
	// Retrieves the list of supported languages.
	ListLanguages(context.Context, *ListLanguagesRequest) (*ListLanguagesResponse, error)
}

// UnimplementedTranslationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTranslationServiceServer struct {
}

func (UnimplementedTranslationServiceServer) Translate(context.Context, *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslationServiceServer) DetectLanguage(context.Context, *DetectLanguageRequest) (*DetectLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectLanguage not implemented")
}
func (UnimplementedTranslationServiceServer) ListLanguages(context.Context, *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguages not implemented")
}

// UnsafeTranslationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslationServiceServer will
// result in compilation errors.
type UnsafeTranslationServiceServer interface {
	mustEmbedUnimplementedTranslationServiceServer()
}

func RegisterTranslationServiceServer(s grpc.ServiceRegistrar, srv TranslationServiceServer) {
	s.RegisterService(&TranslationService_ServiceDesc, srv)
}

func _TranslationService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationService_Translate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_DetectLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).DetectLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationService_DetectLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).DetectLanguage(ctx, req.(*DetectLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_ListLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).ListLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationService_ListLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).ListLanguages(ctx, req.(*ListLanguagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslationService_ServiceDesc is the grpc.ServiceDesc for TranslationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.translate.v2.TranslationService",
	HandlerType: (*TranslationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslationService_Translate_Handler,
		},
		{
			MethodName: "DetectLanguage",
			Handler:    _TranslationService_DetectLanguage_Handler,
		},
		{
			MethodName: "ListLanguages",
			Handler:    _TranslationService_ListLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/translate/v2/translation_service.proto",
}
