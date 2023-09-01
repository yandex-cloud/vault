// Code generated by protoc-gen-goext. DO NOT EDIT.

package lockbox

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Secret) SetId(v string) {
	m.Id = v
}

func (m *Secret) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Secret) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Secret) SetName(v string) {
	m.Name = v
}

func (m *Secret) SetDescription(v string) {
	m.Description = v
}

func (m *Secret) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Secret) SetKmsKeyId(v string) {
	m.KmsKeyId = v
}

func (m *Secret) SetStatus(v Secret_Status) {
	m.Status = v
}

func (m *Secret) SetCurrentVersion(v *Version) {
	m.CurrentVersion = v
}

func (m *Secret) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Version) SetId(v string) {
	m.Id = v
}

func (m *Version) SetSecretId(v string) {
	m.SecretId = v
}

func (m *Version) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Version) SetDestroyAt(v *timestamppb.Timestamp) {
	m.DestroyAt = v
}

func (m *Version) SetDescription(v string) {
	m.Description = v
}

func (m *Version) SetStatus(v Version_Status) {
	m.Status = v
}

func (m *Version) SetPayloadEntryKeys(v []string) {
	m.PayloadEntryKeys = v
}
