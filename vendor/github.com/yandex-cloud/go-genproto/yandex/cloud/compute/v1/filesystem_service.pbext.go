// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetFilesystemRequest) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *ListFilesystemsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListFilesystemsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFilesystemsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFilesystemsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFilesystemsResponse) SetFilesystems(v []*Filesystem) {
	m.Filesystems = v
}

func (m *ListFilesystemsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFilesystemRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateFilesystemRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateFilesystemRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateFilesystemRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateFilesystemRequest) SetTypeId(v string) {
	m.TypeId = v
}

func (m *CreateFilesystemRequest) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *CreateFilesystemRequest) SetSize(v int64) {
	m.Size = v
}

func (m *CreateFilesystemRequest) SetBlockSize(v int64) {
	m.BlockSize = v
}

func (m *CreateFilesystemMetadata) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *UpdateFilesystemRequest) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *UpdateFilesystemRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateFilesystemRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateFilesystemRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateFilesystemRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateFilesystemRequest) SetSize(v int64) {
	m.Size = v
}

func (m *UpdateFilesystemMetadata) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *DeleteFilesystemRequest) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *DeleteFilesystemMetadata) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *ListFilesystemOperationsRequest) SetFilesystemId(v string) {
	m.FilesystemId = v
}

func (m *ListFilesystemOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFilesystemOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFilesystemOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListFilesystemOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
