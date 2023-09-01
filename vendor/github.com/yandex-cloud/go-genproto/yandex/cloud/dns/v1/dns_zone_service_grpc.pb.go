// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/dns/v1/dns_zone_service.proto

package dns

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
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
	DnsZoneService_Get_FullMethodName                  = "/yandex.cloud.dns.v1.DnsZoneService/Get"
	DnsZoneService_List_FullMethodName                 = "/yandex.cloud.dns.v1.DnsZoneService/List"
	DnsZoneService_Create_FullMethodName               = "/yandex.cloud.dns.v1.DnsZoneService/Create"
	DnsZoneService_Update_FullMethodName               = "/yandex.cloud.dns.v1.DnsZoneService/Update"
	DnsZoneService_Delete_FullMethodName               = "/yandex.cloud.dns.v1.DnsZoneService/Delete"
	DnsZoneService_GetRecordSet_FullMethodName         = "/yandex.cloud.dns.v1.DnsZoneService/GetRecordSet"
	DnsZoneService_ListRecordSets_FullMethodName       = "/yandex.cloud.dns.v1.DnsZoneService/ListRecordSets"
	DnsZoneService_UpdateRecordSets_FullMethodName     = "/yandex.cloud.dns.v1.DnsZoneService/UpdateRecordSets"
	DnsZoneService_UpsertRecordSets_FullMethodName     = "/yandex.cloud.dns.v1.DnsZoneService/UpsertRecordSets"
	DnsZoneService_ListOperations_FullMethodName       = "/yandex.cloud.dns.v1.DnsZoneService/ListOperations"
	DnsZoneService_ListAccessBindings_FullMethodName   = "/yandex.cloud.dns.v1.DnsZoneService/ListAccessBindings"
	DnsZoneService_SetAccessBindings_FullMethodName    = "/yandex.cloud.dns.v1.DnsZoneService/SetAccessBindings"
	DnsZoneService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.dns.v1.DnsZoneService/UpdateAccessBindings"
)

// DnsZoneServiceClient is the client API for DnsZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DnsZoneServiceClient interface {
	// Returns the specified DNS zone.
	//
	// To get the list of all available DNS zones, make a [List] request.
	Get(ctx context.Context, in *GetDnsZoneRequest, opts ...grpc.CallOption) (*DnsZone, error)
	// Retrieves the list of DNS zones in the specified folder.
	List(ctx context.Context, in *ListDnsZonesRequest, opts ...grpc.CallOption) (*ListDnsZonesResponse, error)
	// Creates a DNS zone in the specified folder.
	Create(ctx context.Context, in *CreateDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified DNS zone.
	Update(ctx context.Context, in *UpdateDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified DNS zone.
	Delete(ctx context.Context, in *DeleteDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified record set.
	GetRecordSet(ctx context.Context, in *GetDnsZoneRecordSetRequest, opts ...grpc.CallOption) (*RecordSet, error)
	// Retrieves the list of record sets in the specified folder.
	ListRecordSets(ctx context.Context, in *ListDnsZoneRecordSetsRequest, opts ...grpc.CallOption) (*ListDnsZoneRecordSetsResponse, error)
	// Method with strict control for changing zone state. Returns error when:
	// 1. Deleted record is not found.
	// 2. Found record with matched type and name but different TTL or value.
	// 3. Attempted to add record with existing name and type.
	// Deletions happen first. If a record with the same name and type exists in both lists,
	// then the existing record will be deleted, and a new one added.
	UpdateRecordSets(ctx context.Context, in *UpdateRecordSetsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Method without strict control for changing zone state. Nothing happens if deleted record doesn't exist.
	// Deletes records that match all specified fields which allows to delete only specified records from a record set.
	UpsertRecordSets(ctx context.Context, in *UpsertRecordSetsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified DNS zone.
	ListOperations(ctx context.Context, in *ListDnsZoneOperationsRequest, opts ...grpc.CallOption) (*ListDnsZoneOperationsResponse, error)
	// Lists existing access bindings for the specified DNS zone.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified DNS zone.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified DNS zone.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type dnsZoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsZoneServiceClient(cc grpc.ClientConnInterface) DnsZoneServiceClient {
	return &dnsZoneServiceClient{cc}
}

func (c *dnsZoneServiceClient) Get(ctx context.Context, in *GetDnsZoneRequest, opts ...grpc.CallOption) (*DnsZone, error) {
	out := new(DnsZone)
	err := c.cc.Invoke(ctx, DnsZoneService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) List(ctx context.Context, in *ListDnsZonesRequest, opts ...grpc.CallOption) (*ListDnsZonesResponse, error) {
	out := new(ListDnsZonesResponse)
	err := c.cc.Invoke(ctx, DnsZoneService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) Create(ctx context.Context, in *CreateDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) Update(ctx context.Context, in *UpdateDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) Delete(ctx context.Context, in *DeleteDnsZoneRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) GetRecordSet(ctx context.Context, in *GetDnsZoneRecordSetRequest, opts ...grpc.CallOption) (*RecordSet, error) {
	out := new(RecordSet)
	err := c.cc.Invoke(ctx, DnsZoneService_GetRecordSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) ListRecordSets(ctx context.Context, in *ListDnsZoneRecordSetsRequest, opts ...grpc.CallOption) (*ListDnsZoneRecordSetsResponse, error) {
	out := new(ListDnsZoneRecordSetsResponse)
	err := c.cc.Invoke(ctx, DnsZoneService_ListRecordSets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) UpdateRecordSets(ctx context.Context, in *UpdateRecordSetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_UpdateRecordSets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) UpsertRecordSets(ctx context.Context, in *UpsertRecordSetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_UpsertRecordSets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) ListOperations(ctx context.Context, in *ListDnsZoneOperationsRequest, opts ...grpc.CallOption) (*ListDnsZoneOperationsResponse, error) {
	out := new(ListDnsZoneOperationsResponse)
	err := c.cc.Invoke(ctx, DnsZoneService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, DnsZoneService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsZoneServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DnsZoneService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsZoneServiceServer is the server API for DnsZoneService service.
// All implementations should embed UnimplementedDnsZoneServiceServer
// for forward compatibility
type DnsZoneServiceServer interface {
	// Returns the specified DNS zone.
	//
	// To get the list of all available DNS zones, make a [List] request.
	Get(context.Context, *GetDnsZoneRequest) (*DnsZone, error)
	// Retrieves the list of DNS zones in the specified folder.
	List(context.Context, *ListDnsZonesRequest) (*ListDnsZonesResponse, error)
	// Creates a DNS zone in the specified folder.
	Create(context.Context, *CreateDnsZoneRequest) (*operation.Operation, error)
	// Updates the specified DNS zone.
	Update(context.Context, *UpdateDnsZoneRequest) (*operation.Operation, error)
	// Deletes the specified DNS zone.
	Delete(context.Context, *DeleteDnsZoneRequest) (*operation.Operation, error)
	// Returns the specified record set.
	GetRecordSet(context.Context, *GetDnsZoneRecordSetRequest) (*RecordSet, error)
	// Retrieves the list of record sets in the specified folder.
	ListRecordSets(context.Context, *ListDnsZoneRecordSetsRequest) (*ListDnsZoneRecordSetsResponse, error)
	// Method with strict control for changing zone state. Returns error when:
	// 1. Deleted record is not found.
	// 2. Found record with matched type and name but different TTL or value.
	// 3. Attempted to add record with existing name and type.
	// Deletions happen first. If a record with the same name and type exists in both lists,
	// then the existing record will be deleted, and a new one added.
	UpdateRecordSets(context.Context, *UpdateRecordSetsRequest) (*operation.Operation, error)
	// Method without strict control for changing zone state. Nothing happens if deleted record doesn't exist.
	// Deletes records that match all specified fields which allows to delete only specified records from a record set.
	UpsertRecordSets(context.Context, *UpsertRecordSetsRequest) (*operation.Operation, error)
	// Lists operations for the specified DNS zone.
	ListOperations(context.Context, *ListDnsZoneOperationsRequest) (*ListDnsZoneOperationsResponse, error)
	// Lists existing access bindings for the specified DNS zone.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified DNS zone.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified DNS zone.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedDnsZoneServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDnsZoneServiceServer struct {
}

func (UnimplementedDnsZoneServiceServer) Get(context.Context, *GetDnsZoneRequest) (*DnsZone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDnsZoneServiceServer) List(context.Context, *ListDnsZonesRequest) (*ListDnsZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDnsZoneServiceServer) Create(context.Context, *CreateDnsZoneRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDnsZoneServiceServer) Update(context.Context, *UpdateDnsZoneRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDnsZoneServiceServer) Delete(context.Context, *DeleteDnsZoneRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDnsZoneServiceServer) GetRecordSet(context.Context, *GetDnsZoneRecordSetRequest) (*RecordSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecordSet not implemented")
}
func (UnimplementedDnsZoneServiceServer) ListRecordSets(context.Context, *ListDnsZoneRecordSetsRequest) (*ListDnsZoneRecordSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecordSets not implemented")
}
func (UnimplementedDnsZoneServiceServer) UpdateRecordSets(context.Context, *UpdateRecordSetsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecordSets not implemented")
}
func (UnimplementedDnsZoneServiceServer) UpsertRecordSets(context.Context, *UpsertRecordSetsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertRecordSets not implemented")
}
func (UnimplementedDnsZoneServiceServer) ListOperations(context.Context, *ListDnsZoneOperationsRequest) (*ListDnsZoneOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedDnsZoneServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedDnsZoneServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedDnsZoneServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeDnsZoneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DnsZoneServiceServer will
// result in compilation errors.
type UnsafeDnsZoneServiceServer interface {
	mustEmbedUnimplementedDnsZoneServiceServer()
}

func RegisterDnsZoneServiceServer(s grpc.ServiceRegistrar, srv DnsZoneServiceServer) {
	s.RegisterService(&DnsZoneService_ServiceDesc, srv)
}

func _DnsZoneService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDnsZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).Get(ctx, req.(*GetDnsZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDnsZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).List(ctx, req.(*ListDnsZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDnsZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).Create(ctx, req.(*CreateDnsZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDnsZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).Update(ctx, req.(*UpdateDnsZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDnsZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).Delete(ctx, req.(*DeleteDnsZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_GetRecordSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDnsZoneRecordSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).GetRecordSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_GetRecordSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).GetRecordSet(ctx, req.(*GetDnsZoneRecordSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_ListRecordSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDnsZoneRecordSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).ListRecordSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_ListRecordSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).ListRecordSets(ctx, req.(*ListDnsZoneRecordSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_UpdateRecordSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).UpdateRecordSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_UpdateRecordSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).UpdateRecordSets(ctx, req.(*UpdateRecordSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_UpsertRecordSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRecordSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).UpsertRecordSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_UpsertRecordSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).UpsertRecordSets(ctx, req.(*UpsertRecordSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDnsZoneOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).ListOperations(ctx, req.(*ListDnsZoneOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsZoneService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsZoneServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DnsZoneService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsZoneServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DnsZoneService_ServiceDesc is the grpc.ServiceDesc for DnsZoneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DnsZoneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dns.v1.DnsZoneService",
	HandlerType: (*DnsZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DnsZoneService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DnsZoneService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DnsZoneService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DnsZoneService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DnsZoneService_Delete_Handler,
		},
		{
			MethodName: "GetRecordSet",
			Handler:    _DnsZoneService_GetRecordSet_Handler,
		},
		{
			MethodName: "ListRecordSets",
			Handler:    _DnsZoneService_ListRecordSets_Handler,
		},
		{
			MethodName: "UpdateRecordSets",
			Handler:    _DnsZoneService_UpdateRecordSets_Handler,
		},
		{
			MethodName: "UpsertRecordSets",
			Handler:    _DnsZoneService_UpsertRecordSets_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _DnsZoneService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _DnsZoneService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _DnsZoneService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _DnsZoneService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dns/v1/dns_zone_service.proto",
}
