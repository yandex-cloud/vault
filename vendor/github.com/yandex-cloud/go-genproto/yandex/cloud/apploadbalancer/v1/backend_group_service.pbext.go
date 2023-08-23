// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetBackendGroupRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *ListBackendGroupsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListBackendGroupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBackendGroupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBackendGroupsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListBackendGroupsResponse) SetBackendGroups(v []*BackendGroup) {
	m.BackendGroups = v
}

func (m *ListBackendGroupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DeleteBackendGroupRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *DeleteBackendGroupMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

type UpdateBackendGroupRequest_Backend = isUpdateBackendGroupRequest_Backend

func (m *UpdateBackendGroupRequest) SetBackend(v UpdateBackendGroupRequest_Backend) {
	m.Backend = v
}

func (m *UpdateBackendGroupRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *UpdateBackendGroupRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateBackendGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateBackendGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateBackendGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateBackendGroupRequest) SetHttp(v *HttpBackendGroup) {
	m.Backend = &UpdateBackendGroupRequest_Http{
		Http: v,
	}
}

func (m *UpdateBackendGroupRequest) SetGrpc(v *GrpcBackendGroup) {
	m.Backend = &UpdateBackendGroupRequest_Grpc{
		Grpc: v,
	}
}

func (m *UpdateBackendGroupRequest) SetStream(v *StreamBackendGroup) {
	m.Backend = &UpdateBackendGroupRequest_Stream{
		Stream: v,
	}
}

func (m *UpdateBackendGroupMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

type CreateBackendGroupRequest_Backend = isCreateBackendGroupRequest_Backend

func (m *CreateBackendGroupRequest) SetBackend(v CreateBackendGroupRequest_Backend) {
	m.Backend = v
}

func (m *CreateBackendGroupRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateBackendGroupRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateBackendGroupRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateBackendGroupRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateBackendGroupRequest) SetHttp(v *HttpBackendGroup) {
	m.Backend = &CreateBackendGroupRequest_Http{
		Http: v,
	}
}

func (m *CreateBackendGroupRequest) SetGrpc(v *GrpcBackendGroup) {
	m.Backend = &CreateBackendGroupRequest_Grpc{
		Grpc: v,
	}
}

func (m *CreateBackendGroupRequest) SetStream(v *StreamBackendGroup) {
	m.Backend = &CreateBackendGroupRequest_Stream{
		Stream: v,
	}
}

func (m *CreateBackendGroupMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

type AddBackendRequest_Backend = isAddBackendRequest_Backend

func (m *AddBackendRequest) SetBackend(v AddBackendRequest_Backend) {
	m.Backend = v
}

func (m *AddBackendRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *AddBackendRequest) SetHttp(v *HttpBackend) {
	m.Backend = &AddBackendRequest_Http{
		Http: v,
	}
}

func (m *AddBackendRequest) SetGrpc(v *GrpcBackend) {
	m.Backend = &AddBackendRequest_Grpc{
		Grpc: v,
	}
}

func (m *AddBackendRequest) SetStream(v *StreamBackend) {
	m.Backend = &AddBackendRequest_Stream{
		Stream: v,
	}
}

func (m *AddBackendMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *AddBackendMetadata) SetBackendName(v string) {
	m.BackendName = v
}

type UpdateBackendRequest_Backend = isUpdateBackendRequest_Backend

func (m *UpdateBackendRequest) SetBackend(v UpdateBackendRequest_Backend) {
	m.Backend = v
}

func (m *UpdateBackendRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *UpdateBackendRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateBackendRequest) SetHttp(v *HttpBackend) {
	m.Backend = &UpdateBackendRequest_Http{
		Http: v,
	}
}

func (m *UpdateBackendRequest) SetGrpc(v *GrpcBackend) {
	m.Backend = &UpdateBackendRequest_Grpc{
		Grpc: v,
	}
}

func (m *UpdateBackendRequest) SetStream(v *StreamBackend) {
	m.Backend = &UpdateBackendRequest_Stream{
		Stream: v,
	}
}

func (m *UpdateBackendMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *UpdateBackendMetadata) SetBackendName(v string) {
	m.BackendName = v
}

func (m *RemoveBackendRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *RemoveBackendRequest) SetBackendName(v string) {
	m.BackendName = v
}

func (m *RemoveBackendMetadata) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *RemoveBackendMetadata) SetBackendName(v string) {
	m.BackendName = v
}

func (m *ListBackendGroupOperationsRequest) SetBackendGroupId(v string) {
	m.BackendGroupId = v
}

func (m *ListBackendGroupOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBackendGroupOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBackendGroupOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListBackendGroupOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
