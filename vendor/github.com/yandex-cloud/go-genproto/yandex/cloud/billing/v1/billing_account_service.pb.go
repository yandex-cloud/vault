// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/billing/v1/billing_account_service.proto

package billing

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBillingAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the billing account to return.
	// To get the billing account ID, use [BillingAccountService.List] request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBillingAccountRequest) Reset() {
	*x = GetBillingAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBillingAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBillingAccountRequest) ProtoMessage() {}

func (x *GetBillingAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBillingAccountRequest.ProtoReflect.Descriptor instead.
func (*GetBillingAccountRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBillingAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBillingAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListBillingAccountsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListBillingAccountsResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBillingAccountsRequest) Reset() {
	*x = ListBillingAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillingAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillingAccountsRequest) ProtoMessage() {}

func (x *ListBillingAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillingAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListBillingAccountsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListBillingAccountsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBillingAccountsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBillingAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of billing accounts.
	BillingAccounts []*BillingAccount `protobuf:"bytes,1,rep,name=billing_accounts,json=billingAccounts,proto3" json:"billing_accounts,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListBillingAccountsRequest.page_size], use
	// [next_page_token] as the value
	// for the [ListBillingAccountsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBillingAccountsResponse) Reset() {
	*x = ListBillingAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillingAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillingAccountsResponse) ProtoMessage() {}

func (x *ListBillingAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillingAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListBillingAccountsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBillingAccountsResponse) GetBillingAccounts() []*BillingAccount {
	if x != nil {
		return x.BillingAccounts
	}
	return nil
}

func (x *ListBillingAccountsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListBillableObjectBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the billing account to list associated billable object bindings.
	// To get the billing account ID, use [BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,1,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListBillableObjectBindingsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListBillableObjectBindingsResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListBillableObjectBindingsRequest) Reset() {
	*x = ListBillableObjectBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillableObjectBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillableObjectBindingsRequest) ProtoMessage() {}

func (x *ListBillableObjectBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillableObjectBindingsRequest.ProtoReflect.Descriptor instead.
func (*ListBillableObjectBindingsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListBillableObjectBindingsRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

func (x *ListBillableObjectBindingsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBillableObjectBindingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBillableObjectBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of billable object bindings.
	BillableObjectBindings []*BillableObjectBinding `protobuf:"bytes,1,rep,name=billable_object_bindings,json=billableObjectBindings,proto3" json:"billable_object_bindings,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListBillableObjectBindingsRequest.page_size], use
	// [next_page_token] as the value
	// for the [ListBillableObjectBindingsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListBillableObjectBindingsResponse) Reset() {
	*x = ListBillableObjectBindingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBillableObjectBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBillableObjectBindingsResponse) ProtoMessage() {}

func (x *ListBillableObjectBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBillableObjectBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListBillableObjectBindingsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListBillableObjectBindingsResponse) GetBillableObjectBindings() []*BillableObjectBinding {
	if x != nil {
		return x.BillableObjectBindings
	}
	return nil
}

func (x *ListBillableObjectBindingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type BindBillableObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the billing account to bind billable object.
	// To get the billing account ID, use [BillingAccountService.List] request.
	BillingAccountId string `protobuf:"bytes,1,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
	// [yandex.cloud.billing.v1.BillableObject] to bind with billing account.
	BillableObject *BillableObject `protobuf:"bytes,2,opt,name=billable_object,json=billableObject,proto3" json:"billable_object,omitempty"`
}

func (x *BindBillableObjectRequest) Reset() {
	*x = BindBillableObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindBillableObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindBillableObjectRequest) ProtoMessage() {}

func (x *BindBillableObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindBillableObjectRequest.ProtoReflect.Descriptor instead.
func (*BindBillableObjectRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{5}
}

func (x *BindBillableObjectRequest) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

func (x *BindBillableObjectRequest) GetBillableObject() *BillableObject {
	if x != nil {
		return x.BillableObject
	}
	return nil
}

type BindBillableObjectMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the [yandex.cloud.billing.v1.BillableObject] that was bound to billing account.
	BillableObjectId string `protobuf:"bytes,1,opt,name=billable_object_id,json=billableObjectId,proto3" json:"billable_object_id,omitempty"`
}

func (x *BindBillableObjectMetadata) Reset() {
	*x = BindBillableObjectMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindBillableObjectMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindBillableObjectMetadata) ProtoMessage() {}

func (x *BindBillableObjectMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindBillableObjectMetadata.ProtoReflect.Descriptor instead.
func (*BindBillableObjectMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP(), []int{6}
}

func (x *BindBillableObjectMetadata) GetBillableObjectId() string {
	if x != nil {
		return x.BillableObjectId
	}
	return ""
}

var File_yandex_cloud_billing_v1_billing_account_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_billing_v1_billing_account_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x6f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8,
	0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2,
	0x01, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a,
	0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x18, 0x62, 0x69,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x16, 0x62, 0x69,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa9, 0x01, 0x0a,
	0x19, 0x42, 0x69, 0x6e, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4a, 0x0a, 0x1a, 0x42, 0x69, 0x6e, 0x64,
	0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x32, 0xd9, 0x09, 0x0a, 0x15, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b,
	0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0xe6, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x69,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x49, 0x12, 0x47, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xf7,
	0x01, 0x0a, 0x12, 0x42, 0x69, 0x6e, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0xb2,
	0xd2, 0x2a, 0x33, 0x0a, 0x1a, 0x42, 0x69, 0x6e, 0x64, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x15, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c, 0x3a, 0x01, 0x2a, 0x22,
	0x47, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xbb, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xf7, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0xb2, 0xd2, 0x2a, 0x3c, 0x0a, 0x23, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x3a, 0x01, 0x2a,
	0x32, 0x3e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescData = file_yandex_cloud_billing_v1_billing_account_service_proto_rawDesc
)

func file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescData)
	})
	return file_yandex_cloud_billing_v1_billing_account_service_proto_rawDescData
}

var file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_billing_v1_billing_account_service_proto_goTypes = []interface{}{
	(*GetBillingAccountRequest)(nil),           // 0: yandex.cloud.billing.v1.GetBillingAccountRequest
	(*ListBillingAccountsRequest)(nil),         // 1: yandex.cloud.billing.v1.ListBillingAccountsRequest
	(*ListBillingAccountsResponse)(nil),        // 2: yandex.cloud.billing.v1.ListBillingAccountsResponse
	(*ListBillableObjectBindingsRequest)(nil),  // 3: yandex.cloud.billing.v1.ListBillableObjectBindingsRequest
	(*ListBillableObjectBindingsResponse)(nil), // 4: yandex.cloud.billing.v1.ListBillableObjectBindingsResponse
	(*BindBillableObjectRequest)(nil),          // 5: yandex.cloud.billing.v1.BindBillableObjectRequest
	(*BindBillableObjectMetadata)(nil),         // 6: yandex.cloud.billing.v1.BindBillableObjectMetadata
	(*BillingAccount)(nil),                     // 7: yandex.cloud.billing.v1.BillingAccount
	(*BillableObjectBinding)(nil),              // 8: yandex.cloud.billing.v1.BillableObjectBinding
	(*BillableObject)(nil),                     // 9: yandex.cloud.billing.v1.BillableObject
	(*access.ListAccessBindingsRequest)(nil),   // 10: yandex.cloud.access.ListAccessBindingsRequest
	(*access.UpdateAccessBindingsRequest)(nil), // 11: yandex.cloud.access.UpdateAccessBindingsRequest
	(*operation.Operation)(nil),                // 12: yandex.cloud.operation.Operation
	(*access.ListAccessBindingsResponse)(nil),  // 13: yandex.cloud.access.ListAccessBindingsResponse
}
var file_yandex_cloud_billing_v1_billing_account_service_proto_depIdxs = []int32{
	7,  // 0: yandex.cloud.billing.v1.ListBillingAccountsResponse.billing_accounts:type_name -> yandex.cloud.billing.v1.BillingAccount
	8,  // 1: yandex.cloud.billing.v1.ListBillableObjectBindingsResponse.billable_object_bindings:type_name -> yandex.cloud.billing.v1.BillableObjectBinding
	9,  // 2: yandex.cloud.billing.v1.BindBillableObjectRequest.billable_object:type_name -> yandex.cloud.billing.v1.BillableObject
	0,  // 3: yandex.cloud.billing.v1.BillingAccountService.Get:input_type -> yandex.cloud.billing.v1.GetBillingAccountRequest
	1,  // 4: yandex.cloud.billing.v1.BillingAccountService.List:input_type -> yandex.cloud.billing.v1.ListBillingAccountsRequest
	3,  // 5: yandex.cloud.billing.v1.BillingAccountService.ListBillableObjectBindings:input_type -> yandex.cloud.billing.v1.ListBillableObjectBindingsRequest
	5,  // 6: yandex.cloud.billing.v1.BillingAccountService.BindBillableObject:input_type -> yandex.cloud.billing.v1.BindBillableObjectRequest
	10, // 7: yandex.cloud.billing.v1.BillingAccountService.ListAccessBindings:input_type -> yandex.cloud.access.ListAccessBindingsRequest
	11, // 8: yandex.cloud.billing.v1.BillingAccountService.UpdateAccessBindings:input_type -> yandex.cloud.access.UpdateAccessBindingsRequest
	7,  // 9: yandex.cloud.billing.v1.BillingAccountService.Get:output_type -> yandex.cloud.billing.v1.BillingAccount
	2,  // 10: yandex.cloud.billing.v1.BillingAccountService.List:output_type -> yandex.cloud.billing.v1.ListBillingAccountsResponse
	4,  // 11: yandex.cloud.billing.v1.BillingAccountService.ListBillableObjectBindings:output_type -> yandex.cloud.billing.v1.ListBillableObjectBindingsResponse
	12, // 12: yandex.cloud.billing.v1.BillingAccountService.BindBillableObject:output_type -> yandex.cloud.operation.Operation
	13, // 13: yandex.cloud.billing.v1.BillingAccountService.ListAccessBindings:output_type -> yandex.cloud.access.ListAccessBindingsResponse
	12, // 14: yandex.cloud.billing.v1.BillingAccountService.UpdateAccessBindings:output_type -> yandex.cloud.operation.Operation
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_billing_account_service_proto_init() }
func file_yandex_cloud_billing_v1_billing_account_service_proto_init() {
	if File_yandex_cloud_billing_v1_billing_account_service_proto != nil {
		return
	}
	file_yandex_cloud_billing_v1_billing_account_proto_init()
	file_yandex_cloud_billing_v1_billable_object_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBillingAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillingAccountsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillingAccountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillableObjectBindingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBillableObjectBindingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindBillableObjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindBillableObjectMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_billing_v1_billing_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_billing_v1_billing_account_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_billing_account_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_billing_account_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_billing_account_service_proto = out.File
	file_yandex_cloud_billing_v1_billing_account_service_proto_rawDesc = nil
	file_yandex_cloud_billing_v1_billing_account_service_proto_goTypes = nil
	file_yandex_cloud_billing_v1_billing_account_service_proto_depIdxs = nil
}
