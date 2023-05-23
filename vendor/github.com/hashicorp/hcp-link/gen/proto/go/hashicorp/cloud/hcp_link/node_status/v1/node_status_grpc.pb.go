// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: hashicorp/cloud/hcp_link/node_status/v1/node_status.proto

package node_statusv1

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

// NodeStatusServiceClient is the client API for NodeStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeStatusServiceClient interface {
	// GetNodeStatus will be used to regularly fetch the node’s current status.
	GetNodeStatus(ctx context.Context, in *GetNodeStatusRequest, opts ...grpc.CallOption) (*GetNodeStatusResponse, error)
}

type nodeStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeStatusServiceClient(cc grpc.ClientConnInterface) NodeStatusServiceClient {
	return &nodeStatusServiceClient{cc}
}

func (c *nodeStatusServiceClient) GetNodeStatus(ctx context.Context, in *GetNodeStatusRequest, opts ...grpc.CallOption) (*GetNodeStatusResponse, error) {
	out := new(GetNodeStatusResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.cloud.hcp_link.node_status.v1.NodeStatusService/GetNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeStatusServiceServer is the server API for NodeStatusService service.
// All implementations must embed UnimplementedNodeStatusServiceServer
// for forward compatibility
type NodeStatusServiceServer interface {
	// GetNodeStatus will be used to regularly fetch the node’s current status.
	GetNodeStatus(context.Context, *GetNodeStatusRequest) (*GetNodeStatusResponse, error)
	mustEmbedUnimplementedNodeStatusServiceServer()
}

// UnimplementedNodeStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeStatusServiceServer struct {
}

func (UnimplementedNodeStatusServiceServer) GetNodeStatus(context.Context, *GetNodeStatusRequest) (*GetNodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeStatus not implemented")
}
func (UnimplementedNodeStatusServiceServer) mustEmbedUnimplementedNodeStatusServiceServer() {}

// UnsafeNodeStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeStatusServiceServer will
// result in compilation errors.
type UnsafeNodeStatusServiceServer interface {
	mustEmbedUnimplementedNodeStatusServiceServer()
}

func RegisterNodeStatusServiceServer(s grpc.ServiceRegistrar, srv NodeStatusServiceServer) {
	s.RegisterService(&NodeStatusService_ServiceDesc, srv)
}

func _NodeStatusService_GetNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatusServiceServer).GetNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.cloud.hcp_link.node_status.v1.NodeStatusService/GetNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatusServiceServer).GetNodeStatus(ctx, req.(*GetNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeStatusService_ServiceDesc is the grpc.ServiceDesc for NodeStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.cloud.hcp_link.node_status.v1.NodeStatusService",
	HandlerType: (*NodeStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeStatus",
			Handler:    _NodeStatusService_GetNodeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hashicorp/cloud/hcp_link/node_status/v1/node_status.proto",
}
