// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/apploadbalancer/v1/virtual_host_service.proto

package apploadbalancer

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
	VirtualHostService_Get_FullMethodName         = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/Get"
	VirtualHostService_List_FullMethodName        = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/List"
	VirtualHostService_Create_FullMethodName      = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/Create"
	VirtualHostService_Update_FullMethodName      = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/Update"
	VirtualHostService_Delete_FullMethodName      = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/Delete"
	VirtualHostService_RemoveRoute_FullMethodName = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/RemoveRoute"
	VirtualHostService_UpdateRoute_FullMethodName = "/yandex.cloud.apploadbalancer.v1.VirtualHostService/UpdateRoute"
)

// VirtualHostServiceClient is the client API for VirtualHostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtualHostServiceClient interface {
	// Returns the specified virtual host.
	//
	// To get the list of all virtual hosts of an HTTP router, make a [List] request.
	Get(ctx context.Context, in *GetVirtualHostRequest, opts ...grpc.CallOption) (*VirtualHost, error)
	// Lists virtual hosts of the specified HTTP router.
	List(ctx context.Context, in *ListVirtualHostsRequest, opts ...grpc.CallOption) (*ListVirtualHostsResponse, error)
	// Creates a virtual host in the specified HTTP router.
	Create(ctx context.Context, in *CreateVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified virtual host of the specified HTTP router.
	Update(ctx context.Context, in *UpdateVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified virtual host.
	Delete(ctx context.Context, in *DeleteVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified route from the specified virtual host.
	RemoveRoute(ctx context.Context, in *RemoveRouteRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified route of the specified virtual host.
	UpdateRoute(ctx context.Context, in *UpdateRouteRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type virtualHostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualHostServiceClient(cc grpc.ClientConnInterface) VirtualHostServiceClient {
	return &virtualHostServiceClient{cc}
}

func (c *virtualHostServiceClient) Get(ctx context.Context, in *GetVirtualHostRequest, opts ...grpc.CallOption) (*VirtualHost, error) {
	out := new(VirtualHost)
	err := c.cc.Invoke(ctx, VirtualHostService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) List(ctx context.Context, in *ListVirtualHostsRequest, opts ...grpc.CallOption) (*ListVirtualHostsResponse, error) {
	out := new(ListVirtualHostsResponse)
	err := c.cc.Invoke(ctx, VirtualHostService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) Create(ctx context.Context, in *CreateVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, VirtualHostService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) Update(ctx context.Context, in *UpdateVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, VirtualHostService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) Delete(ctx context.Context, in *DeleteVirtualHostRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, VirtualHostService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) RemoveRoute(ctx context.Context, in *RemoveRouteRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, VirtualHostService_RemoveRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHostServiceClient) UpdateRoute(ctx context.Context, in *UpdateRouteRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, VirtualHostService_UpdateRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHostServiceServer is the server API for VirtualHostService service.
// All implementations should embed UnimplementedVirtualHostServiceServer
// for forward compatibility
type VirtualHostServiceServer interface {
	// Returns the specified virtual host.
	//
	// To get the list of all virtual hosts of an HTTP router, make a [List] request.
	Get(context.Context, *GetVirtualHostRequest) (*VirtualHost, error)
	// Lists virtual hosts of the specified HTTP router.
	List(context.Context, *ListVirtualHostsRequest) (*ListVirtualHostsResponse, error)
	// Creates a virtual host in the specified HTTP router.
	Create(context.Context, *CreateVirtualHostRequest) (*operation.Operation, error)
	// Updates the specified virtual host of the specified HTTP router.
	Update(context.Context, *UpdateVirtualHostRequest) (*operation.Operation, error)
	// Deletes the specified virtual host.
	Delete(context.Context, *DeleteVirtualHostRequest) (*operation.Operation, error)
	// Deletes the specified route from the specified virtual host.
	RemoveRoute(context.Context, *RemoveRouteRequest) (*operation.Operation, error)
	// Updates the specified route of the specified virtual host.
	UpdateRoute(context.Context, *UpdateRouteRequest) (*operation.Operation, error)
}

// UnimplementedVirtualHostServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVirtualHostServiceServer struct {
}

func (UnimplementedVirtualHostServiceServer) Get(context.Context, *GetVirtualHostRequest) (*VirtualHost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedVirtualHostServiceServer) List(context.Context, *ListVirtualHostsRequest) (*ListVirtualHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedVirtualHostServiceServer) Create(context.Context, *CreateVirtualHostRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVirtualHostServiceServer) Update(context.Context, *UpdateVirtualHostRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVirtualHostServiceServer) Delete(context.Context, *DeleteVirtualHostRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVirtualHostServiceServer) RemoveRoute(context.Context, *RemoveRouteRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoute not implemented")
}
func (UnimplementedVirtualHostServiceServer) UpdateRoute(context.Context, *UpdateRouteRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoute not implemented")
}

// UnsafeVirtualHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtualHostServiceServer will
// result in compilation errors.
type UnsafeVirtualHostServiceServer interface {
	mustEmbedUnimplementedVirtualHostServiceServer()
}

func RegisterVirtualHostServiceServer(s grpc.ServiceRegistrar, srv VirtualHostServiceServer) {
	s.RegisterService(&VirtualHostService_ServiceDesc, srv)
}

func _VirtualHostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVirtualHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).Get(ctx, req.(*GetVirtualHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVirtualHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).List(ctx, req.(*ListVirtualHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVirtualHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).Create(ctx, req.(*CreateVirtualHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVirtualHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).Update(ctx, req.(*UpdateVirtualHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVirtualHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).Delete(ctx, req.(*DeleteVirtualHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_RemoveRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).RemoveRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_RemoveRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).RemoveRoute(ctx, req.(*RemoveRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHostService_UpdateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHostServiceServer).UpdateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VirtualHostService_UpdateRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHostServiceServer).UpdateRoute(ctx, req.(*UpdateRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VirtualHostService_ServiceDesc is the grpc.ServiceDesc for VirtualHostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtualHostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.apploadbalancer.v1.VirtualHostService",
	HandlerType: (*VirtualHostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _VirtualHostService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _VirtualHostService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _VirtualHostService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VirtualHostService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VirtualHostService_Delete_Handler,
		},
		{
			MethodName: "RemoveRoute",
			Handler:    _VirtualHostService_RemoveRoute_Handler,
		},
		{
			MethodName: "UpdateRoute",
			Handler:    _VirtualHostService_UpdateRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/apploadbalancer/v1/virtual_host_service.proto",
}
