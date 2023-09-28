// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/certificatemanager/v1/certificate_content_service.proto

package certificatemanager

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
	CertificateContentService_Get_FullMethodName = "/yandex.cloud.certificatemanager.v1.CertificateContentService/Get"
)

// CertificateContentServiceClient is the client API for CertificateContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateContentServiceClient interface {
	// Returns chain and private key of the specified certificate.
	Get(ctx context.Context, in *GetCertificateContentRequest, opts ...grpc.CallOption) (*GetCertificateContentResponse, error)
}

type certificateContentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateContentServiceClient(cc grpc.ClientConnInterface) CertificateContentServiceClient {
	return &certificateContentServiceClient{cc}
}

func (c *certificateContentServiceClient) Get(ctx context.Context, in *GetCertificateContentRequest, opts ...grpc.CallOption) (*GetCertificateContentResponse, error) {
	out := new(GetCertificateContentResponse)
	err := c.cc.Invoke(ctx, CertificateContentService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateContentServiceServer is the server API for CertificateContentService service.
// All implementations should embed UnimplementedCertificateContentServiceServer
// for forward compatibility
type CertificateContentServiceServer interface {
	// Returns chain and private key of the specified certificate.
	Get(context.Context, *GetCertificateContentRequest) (*GetCertificateContentResponse, error)
}

// UnimplementedCertificateContentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCertificateContentServiceServer struct {
}

func (UnimplementedCertificateContentServiceServer) Get(context.Context, *GetCertificateContentRequest) (*GetCertificateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

// UnsafeCertificateContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateContentServiceServer will
// result in compilation errors.
type UnsafeCertificateContentServiceServer interface {
	mustEmbedUnimplementedCertificateContentServiceServer()
}

func RegisterCertificateContentServiceServer(s grpc.ServiceRegistrar, srv CertificateContentServiceServer) {
	s.RegisterService(&CertificateContentService_ServiceDesc, srv)
}

func _CertificateContentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateContentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateContentService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateContentServiceServer).Get(ctx, req.(*GetCertificateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateContentService_ServiceDesc is the grpc.ServiceDesc for CertificateContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.certificatemanager.v1.CertificateContentService",
	HandlerType: (*CertificateContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CertificateContentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/certificatemanager/v1/certificate_content_service.proto",
}
