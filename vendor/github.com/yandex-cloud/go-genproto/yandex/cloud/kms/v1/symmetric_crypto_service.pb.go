// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/kms/v1/symmetric_crypto_service.proto

package kms

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SymmetricEncryptRequest struct {
	// ID of the symmetric KMS key to use for encryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the key version to encrypt plaintext with.
	// Defaults to the primary version if not specified.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Additional authenticated data (AAD context), optional.
	// If specified, this data will be required for decryption with the [SymmetricDecryptRequest].
	// Should be encoded with base64.
	AadContext []byte `protobuf:"bytes,3,opt,name=aad_context,json=aadContext,proto3" json:"aad_context,omitempty"`
	// Plaintext to be encrypted.
	// Should be encoded with base64.
	Plaintext            []byte   `protobuf:"bytes,4,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricEncryptRequest) Reset()         { *m = SymmetricEncryptRequest{} }
func (m *SymmetricEncryptRequest) String() string { return proto.CompactTextString(m) }
func (*SymmetricEncryptRequest) ProtoMessage()    {}
func (*SymmetricEncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{0}
}

func (m *SymmetricEncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricEncryptRequest.Unmarshal(m, b)
}
func (m *SymmetricEncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricEncryptRequest.Marshal(b, m, deterministic)
}
func (m *SymmetricEncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricEncryptRequest.Merge(m, src)
}
func (m *SymmetricEncryptRequest) XXX_Size() int {
	return xxx_messageInfo_SymmetricEncryptRequest.Size(m)
}
func (m *SymmetricEncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricEncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricEncryptRequest proto.InternalMessageInfo

func (m *SymmetricEncryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricEncryptRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *SymmetricEncryptRequest) GetAadContext() []byte {
	if m != nil {
		return m.AadContext
	}
	return nil
}

func (m *SymmetricEncryptRequest) GetPlaintext() []byte {
	if m != nil {
		return m.Plaintext
	}
	return nil
}

type SymmetricEncryptResponse struct {
	// ID of the symmetric KMS key that was used for encryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the key version that was used for encryption.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Resulting ciphertext.
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricEncryptResponse) Reset()         { *m = SymmetricEncryptResponse{} }
func (m *SymmetricEncryptResponse) String() string { return proto.CompactTextString(m) }
func (*SymmetricEncryptResponse) ProtoMessage()    {}
func (*SymmetricEncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{1}
}

func (m *SymmetricEncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricEncryptResponse.Unmarshal(m, b)
}
func (m *SymmetricEncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricEncryptResponse.Marshal(b, m, deterministic)
}
func (m *SymmetricEncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricEncryptResponse.Merge(m, src)
}
func (m *SymmetricEncryptResponse) XXX_Size() int {
	return xxx_messageInfo_SymmetricEncryptResponse.Size(m)
}
func (m *SymmetricEncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricEncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricEncryptResponse proto.InternalMessageInfo

func (m *SymmetricEncryptResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricEncryptResponse) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *SymmetricEncryptResponse) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type SymmetricDecryptRequest struct {
	// ID of the symmetric KMS key to use for decryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Additional authenticated data, must be the same as was provided
	// in the corresponding [SymmetricEncryptRequest].
	// Should be encoded with base64.
	AadContext []byte `protobuf:"bytes,2,opt,name=aad_context,json=aadContext,proto3" json:"aad_context,omitempty"`
	// Ciphertext to be decrypted.
	// Should be encoded with base64.
	Ciphertext           []byte   `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricDecryptRequest) Reset()         { *m = SymmetricDecryptRequest{} }
func (m *SymmetricDecryptRequest) String() string { return proto.CompactTextString(m) }
func (*SymmetricDecryptRequest) ProtoMessage()    {}
func (*SymmetricDecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{2}
}

func (m *SymmetricDecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricDecryptRequest.Unmarshal(m, b)
}
func (m *SymmetricDecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricDecryptRequest.Marshal(b, m, deterministic)
}
func (m *SymmetricDecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricDecryptRequest.Merge(m, src)
}
func (m *SymmetricDecryptRequest) XXX_Size() int {
	return xxx_messageInfo_SymmetricDecryptRequest.Size(m)
}
func (m *SymmetricDecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricDecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricDecryptRequest proto.InternalMessageInfo

func (m *SymmetricDecryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricDecryptRequest) GetAadContext() []byte {
	if m != nil {
		return m.AadContext
	}
	return nil
}

func (m *SymmetricDecryptRequest) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type SymmetricDecryptResponse struct {
	// ID of the symmetric KMS key that was used for decryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the key version that was used for decryption.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Decrypted plaintext.
	Plaintext            []byte   `protobuf:"bytes,3,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricDecryptResponse) Reset()         { *m = SymmetricDecryptResponse{} }
func (m *SymmetricDecryptResponse) String() string { return proto.CompactTextString(m) }
func (*SymmetricDecryptResponse) ProtoMessage()    {}
func (*SymmetricDecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{3}
}

func (m *SymmetricDecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricDecryptResponse.Unmarshal(m, b)
}
func (m *SymmetricDecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricDecryptResponse.Marshal(b, m, deterministic)
}
func (m *SymmetricDecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricDecryptResponse.Merge(m, src)
}
func (m *SymmetricDecryptResponse) XXX_Size() int {
	return xxx_messageInfo_SymmetricDecryptResponse.Size(m)
}
func (m *SymmetricDecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricDecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricDecryptResponse proto.InternalMessageInfo

func (m *SymmetricDecryptResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricDecryptResponse) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *SymmetricDecryptResponse) GetPlaintext() []byte {
	if m != nil {
		return m.Plaintext
	}
	return nil
}

type GenerateDataKeyRequest struct {
	// ID of the symmetric KMS key that the generated data key should be encrypted with.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the key version to encrypt the generated data key with.
	// Defaults to the primary version if not specified.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Additional authenticated data (AAD context), optional.
	// If specified, this data will be required for decryption with the [SymmetricDecryptRequest].
	// Should be encoded with base64.
	AadContext []byte `protobuf:"bytes,3,opt,name=aad_context,json=aadContext,proto3" json:"aad_context,omitempty"`
	// Encryption algorithm and key length for the generated data key.
	DataKeySpec SymmetricAlgorithm `protobuf:"varint,4,opt,name=data_key_spec,json=dataKeySpec,proto3,enum=yandex.cloud.kms.v1.SymmetricAlgorithm" json:"data_key_spec,omitempty"`
	// If `true`, the method won't return the data key as plaintext.
	// Default value is `false`.
	SkipPlaintext        bool     `protobuf:"varint,5,opt,name=skip_plaintext,json=skipPlaintext,proto3" json:"skip_plaintext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateDataKeyRequest) Reset()         { *m = GenerateDataKeyRequest{} }
func (m *GenerateDataKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateDataKeyRequest) ProtoMessage()    {}
func (*GenerateDataKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{4}
}

func (m *GenerateDataKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateDataKeyRequest.Unmarshal(m, b)
}
func (m *GenerateDataKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateDataKeyRequest.Marshal(b, m, deterministic)
}
func (m *GenerateDataKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateDataKeyRequest.Merge(m, src)
}
func (m *GenerateDataKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateDataKeyRequest.Size(m)
}
func (m *GenerateDataKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateDataKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateDataKeyRequest proto.InternalMessageInfo

func (m *GenerateDataKeyRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *GenerateDataKeyRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *GenerateDataKeyRequest) GetAadContext() []byte {
	if m != nil {
		return m.AadContext
	}
	return nil
}

func (m *GenerateDataKeyRequest) GetDataKeySpec() SymmetricAlgorithm {
	if m != nil {
		return m.DataKeySpec
	}
	return SymmetricAlgorithm_SYMMETRIC_ALGORITHM_UNSPECIFIED
}

func (m *GenerateDataKeyRequest) GetSkipPlaintext() bool {
	if m != nil {
		return m.SkipPlaintext
	}
	return false
}

type GenerateDataKeyResponse struct {
	// ID of the symmetric KMS key that was used to encrypt the generated data key.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the key version that was used for encryption.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Generated data key as plaintext.
	// The field is empty, if the [GenerateDataKeyRequest.skip_plaintext] parameter
	// was set to `true`.
	DataKeyPlaintext []byte `protobuf:"bytes,3,opt,name=data_key_plaintext,json=dataKeyPlaintext,proto3" json:"data_key_plaintext,omitempty"`
	// The encrypted data key.
	DataKeyCiphertext    []byte   `protobuf:"bytes,4,opt,name=data_key_ciphertext,json=dataKeyCiphertext,proto3" json:"data_key_ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateDataKeyResponse) Reset()         { *m = GenerateDataKeyResponse{} }
func (m *GenerateDataKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateDataKeyResponse) ProtoMessage()    {}
func (*GenerateDataKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{5}
}

func (m *GenerateDataKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateDataKeyResponse.Unmarshal(m, b)
}
func (m *GenerateDataKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateDataKeyResponse.Marshal(b, m, deterministic)
}
func (m *GenerateDataKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateDataKeyResponse.Merge(m, src)
}
func (m *GenerateDataKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateDataKeyResponse.Size(m)
}
func (m *GenerateDataKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateDataKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateDataKeyResponse proto.InternalMessageInfo

func (m *GenerateDataKeyResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *GenerateDataKeyResponse) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *GenerateDataKeyResponse) GetDataKeyPlaintext() []byte {
	if m != nil {
		return m.DataKeyPlaintext
	}
	return nil
}

func (m *GenerateDataKeyResponse) GetDataKeyCiphertext() []byte {
	if m != nil {
		return m.DataKeyCiphertext
	}
	return nil
}

type SymmetricReEncryptRequest struct {
	// ID of the new key to be used for encryption.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of the version of the new key to be used for encryption.
	// Defaults to the primary version if not specified.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Additional authenticated data to be required for decryption.
	// Should be encoded with base64.
	AadContext []byte `protobuf:"bytes,3,opt,name=aad_context,json=aadContext,proto3" json:"aad_context,omitempty"`
	// ID of the key that the ciphertext is currently encrypted with. May be the same as for the new key.
	SourceKeyId string `protobuf:"bytes,4,opt,name=source_key_id,json=sourceKeyId,proto3" json:"source_key_id,omitempty"`
	// Additional authenticated data provided with the initial encryption request.
	// Should be encoded with base64.
	SourceAadContext []byte `protobuf:"bytes,5,opt,name=source_aad_context,json=sourceAadContext,proto3" json:"source_aad_context,omitempty"`
	// Ciphertext to re-encrypt.
	// Should be encoded with base64.
	Ciphertext           []byte   `protobuf:"bytes,6,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricReEncryptRequest) Reset()         { *m = SymmetricReEncryptRequest{} }
func (m *SymmetricReEncryptRequest) String() string { return proto.CompactTextString(m) }
func (*SymmetricReEncryptRequest) ProtoMessage()    {}
func (*SymmetricReEncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{6}
}

func (m *SymmetricReEncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricReEncryptRequest.Unmarshal(m, b)
}
func (m *SymmetricReEncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricReEncryptRequest.Marshal(b, m, deterministic)
}
func (m *SymmetricReEncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricReEncryptRequest.Merge(m, src)
}
func (m *SymmetricReEncryptRequest) XXX_Size() int {
	return xxx_messageInfo_SymmetricReEncryptRequest.Size(m)
}
func (m *SymmetricReEncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricReEncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricReEncryptRequest proto.InternalMessageInfo

func (m *SymmetricReEncryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricReEncryptRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *SymmetricReEncryptRequest) GetAadContext() []byte {
	if m != nil {
		return m.AadContext
	}
	return nil
}

func (m *SymmetricReEncryptRequest) GetSourceKeyId() string {
	if m != nil {
		return m.SourceKeyId
	}
	return ""
}

func (m *SymmetricReEncryptRequest) GetSourceAadContext() []byte {
	if m != nil {
		return m.SourceAadContext
	}
	return nil
}

func (m *SymmetricReEncryptRequest) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

type SymmetricReEncryptResponse struct {
	// ID of the key that the ciphertext is encrypted with now.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// ID of key version that was used for encryption.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// ID of the key that the ciphertext was encrypted with previously.
	SourceKeyId string `protobuf:"bytes,3,opt,name=source_key_id,json=sourceKeyId,proto3" json:"source_key_id,omitempty"`
	// ID of the key version that was used to decrypt the re-encrypted ciphertext.
	SourceVersionId string `protobuf:"bytes,4,opt,name=source_version_id,json=sourceVersionId,proto3" json:"source_version_id,omitempty"`
	// Resulting re-encrypted ciphertext.
	Ciphertext           []byte   `protobuf:"bytes,5,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SymmetricReEncryptResponse) Reset()         { *m = SymmetricReEncryptResponse{} }
func (m *SymmetricReEncryptResponse) String() string { return proto.CompactTextString(m) }
func (*SymmetricReEncryptResponse) ProtoMessage()    {}
func (*SymmetricReEncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a384b1425a1aa84e, []int{7}
}

func (m *SymmetricReEncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymmetricReEncryptResponse.Unmarshal(m, b)
}
func (m *SymmetricReEncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymmetricReEncryptResponse.Marshal(b, m, deterministic)
}
func (m *SymmetricReEncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymmetricReEncryptResponse.Merge(m, src)
}
func (m *SymmetricReEncryptResponse) XXX_Size() int {
	return xxx_messageInfo_SymmetricReEncryptResponse.Size(m)
}
func (m *SymmetricReEncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SymmetricReEncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SymmetricReEncryptResponse proto.InternalMessageInfo

func (m *SymmetricReEncryptResponse) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *SymmetricReEncryptResponse) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *SymmetricReEncryptResponse) GetSourceKeyId() string {
	if m != nil {
		return m.SourceKeyId
	}
	return ""
}

func (m *SymmetricReEncryptResponse) GetSourceVersionId() string {
	if m != nil {
		return m.SourceVersionId
	}
	return ""
}

func (m *SymmetricReEncryptResponse) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func init() {
	proto.RegisterType((*SymmetricEncryptRequest)(nil), "yandex.cloud.kms.v1.SymmetricEncryptRequest")
	proto.RegisterType((*SymmetricEncryptResponse)(nil), "yandex.cloud.kms.v1.SymmetricEncryptResponse")
	proto.RegisterType((*SymmetricDecryptRequest)(nil), "yandex.cloud.kms.v1.SymmetricDecryptRequest")
	proto.RegisterType((*SymmetricDecryptResponse)(nil), "yandex.cloud.kms.v1.SymmetricDecryptResponse")
	proto.RegisterType((*GenerateDataKeyRequest)(nil), "yandex.cloud.kms.v1.GenerateDataKeyRequest")
	proto.RegisterType((*GenerateDataKeyResponse)(nil), "yandex.cloud.kms.v1.GenerateDataKeyResponse")
	proto.RegisterType((*SymmetricReEncryptRequest)(nil), "yandex.cloud.kms.v1.SymmetricReEncryptRequest")
	proto.RegisterType((*SymmetricReEncryptResponse)(nil), "yandex.cloud.kms.v1.SymmetricReEncryptResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/kms/v1/symmetric_crypto_service.proto", fileDescriptor_a384b1425a1aa84e)
}

var fileDescriptor_a384b1425a1aa84e = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xbf, 0x4f, 0xdb, 0x40,
	0x14, 0xc7, 0xe5, 0x90, 0x04, 0xf2, 0xf8, 0x6d, 0x54, 0xa0, 0x11, 0xb4, 0xc8, 0x05, 0x11, 0x05,
	0x62, 0x93, 0xa0, 0xb6, 0x94, 0x1f, 0x03, 0x81, 0xaa, 0x42, 0x2c, 0x95, 0x91, 0x18, 0xba, 0x44,
	0x87, 0x7d, 0x0a, 0x96, 0x13, 0x9f, 0x6b, 0x3b, 0x11, 0x56, 0xd5, 0xa5, 0x4b, 0x2b, 0xa6, 0x56,
	0x5d, 0x3a, 0x75, 0xee, 0xd2, 0x7f, 0xa2, 0x1d, 0xca, 0xde, 0x7f, 0xa1, 0x43, 0xff, 0x86, 0x4e,
	0x55, 0xee, 0x2e, 0x8e, 0x9d, 0x98, 0x40, 0x45, 0x07, 0x56, 0xdf, 0xf7, 0xf9, 0x7d, 0xde, 0xf7,
	0xee, 0xbd, 0x3b, 0x28, 0xf9, 0xc8, 0xd2, 0xf1, 0x99, 0xa2, 0xd5, 0x48, 0x43, 0x57, 0xcc, 0xba,
	0xab, 0x34, 0x8b, 0x8a, 0xeb, 0xd7, 0xeb, 0xd8, 0x73, 0x0c, 0xad, 0xa2, 0x39, 0xbe, 0xed, 0x91,
	0x8a, 0x8b, 0x9d, 0xa6, 0xa1, 0x61, 0xd9, 0x76, 0x88, 0x47, 0xc4, 0x29, 0x16, 0x23, 0xd3, 0x18,
	0xd9, 0xac, 0xbb, 0x72, 0xb3, 0x98, 0x9d, 0xab, 0x12, 0x52, 0xad, 0x61, 0x05, 0xd9, 0x86, 0x82,
	0x2c, 0x8b, 0x78, 0xc8, 0x33, 0x88, 0xe5, 0xb2, 0x90, 0xec, 0x72, 0xff, 0x34, 0x26, 0xf6, 0xb9,
	0x70, 0x3e, 0x22, 0x6c, 0xa2, 0x9a, 0xa1, 0xd3, 0x1f, 0xb1, 0x65, 0xe9, 0xbb, 0x00, 0x33, 0x47,
	0xed, 0xb0, 0xa7, 0x16, 0xc5, 0x53, 0xf1, 0xcb, 0x06, 0x76, 0x3d, 0xf1, 0x01, 0xa4, 0x4d, 0xec,
	0x57, 0x0c, 0x7d, 0x56, 0x58, 0x10, 0x72, 0x99, 0xf2, 0xc8, 0xef, 0x1f, 0x45, 0xe1, 0xfc, 0xa2,
	0x98, 0xdc, 0xde, 0x79, 0xb8, 0xa6, 0xa6, 0x4c, 0xec, 0x1f, 0xe8, 0xe2, 0x32, 0x40, 0x13, 0x3b,
	0xae, 0x41, 0xac, 0x96, 0x30, 0x41, 0x85, 0x43, 0x81, 0x28, 0xc3, 0xd7, 0x0e, 0x74, 0x71, 0x05,
	0x86, 0x11, 0xd2, 0x2b, 0x1a, 0xb1, 0x3c, 0x7c, 0xe6, 0xcd, 0x0e, 0x2c, 0x08, 0xb9, 0x91, 0x32,
	0x9c, 0x5f, 0x14, 0xd3, 0xdb, 0x3b, 0x1b, 0xc5, 0x27, 0x25, 0x15, 0x10, 0xd2, 0xf7, 0xd8, 0xaa,
	0x58, 0x80, 0x8c, 0x5d, 0x43, 0x06, 0x93, 0x26, 0xa9, 0x74, 0x9c, 0x67, 0x1f, 0xdc, 0xde, 0x59,
	0x2f, 0x3d, 0x7e, 0xb4, 0xa1, 0x76, 0x14, 0xd2, 0x3b, 0x01, 0x66, 0x7b, 0xab, 0x70, 0x6d, 0x62,
	0xb9, 0xf8, 0x3f, 0x97, 0x71, 0x0f, 0x40, 0x33, 0xec, 0x53, 0xec, 0x74, 0xaa, 0x50, 0x43, 0x5f,
	0xa4, 0x0f, 0x61, 0x43, 0xf7, 0xf1, 0xbf, 0x1b, 0xda, 0xe5, 0x53, 0xa2, 0xaf, 0x4f, 0x8b, 0xbd,
	0x34, 0xe5, 0x64, 0xeb, 0xaf, 0x11, 0x26, 0x2b, 0xe4, 0x4e, 0x80, 0xc4, 0xdd, 0xb9, 0x13, 0x65,
	0x6a, 0x53, 0xcc, 0xf7, 0xfa, 0x11, 0x76, 0x61, 0x2e, 0xbc, 0x3f, 0xcc, 0x84, 0xd0, 0x76, 0xbc,
	0x4d, 0xc0, 0xf4, 0x33, 0x6c, 0x61, 0x07, 0x79, 0x78, 0x1f, 0x79, 0xe8, 0x10, 0xfb, 0xb7, 0xe0,
	0x4c, 0x1d, 0xc2, 0xa8, 0x8e, 0x3c, 0xd4, 0xea, 0x8d, 0x8a, 0x6b, 0x63, 0x8d, 0x9e, 0xab, 0xb1,
	0xd2, 0xb2, 0x1c, 0xd3, 0x7d, 0x72, 0xe0, 0xd7, 0x6e, 0xad, 0x4a, 0x1c, 0xc3, 0x3b, 0xad, 0xab,
	0xc3, 0x3a, 0x2b, 0xe5, 0xc8, 0xc6, 0x9a, 0xb8, 0x04, 0x63, 0xae, 0x69, 0xd8, 0x95, 0x8e, 0x0b,
	0xa9, 0x05, 0x21, 0x37, 0xa4, 0x8e, 0xb6, 0xbe, 0x3e, 0x0f, 0x9c, 0xf8, 0x2a, 0xc0, 0x4c, 0x8f,
	0x13, 0x37, 0x72, 0x7e, 0x15, 0xc4, 0xa0, 0x8a, 0xee, 0x2d, 0x98, 0xe0, 0x84, 0x41, 0x7e, 0x51,
	0x86, 0xa9, 0x40, 0x1d, 0x3a, 0x28, 0xb4, 0xa3, 0xd4, 0x49, 0x2e, 0xdf, 0xeb, 0x9c, 0x94, 0x2f,
	0x09, 0xb8, 0x1b, 0x94, 0xae, 0xe2, 0xdb, 0x33, 0x10, 0xd6, 0x60, 0xd4, 0x25, 0x0d, 0x47, 0xc3,
	0x15, 0x4e, 0x90, 0x8c, 0x21, 0x18, 0x66, 0x92, 0x43, 0xca, 0xb1, 0x01, 0x22, 0x8f, 0x08, 0x67,
	0x49, 0xf5, 0x64, 0x99, 0x60, 0xaa, 0xdd, 0xcb, 0x9a, 0x2a, 0x7d, 0x49, 0x53, 0x7d, 0x13, 0x20,
	0x1b, 0x67, 0xd5, 0x8d, 0x76, 0x57, 0xea, 0x2e, 0x73, 0x80, 0x2a, 0x22, 0x85, 0xe5, 0x61, 0x92,
	0x6b, 0x42, 0x7f, 0xa2, 0x76, 0xa8, 0xe3, 0x6c, 0xe1, 0xf8, 0x92, 0x69, 0x95, 0xea, 0x9e, 0x56,
	0xa5, 0x3f, 0x49, 0x98, 0x0e, 0x8a, 0xd8, 0xa3, 0x77, 0xd3, 0x11, 0xbb, 0x9a, 0xc4, 0xf7, 0x02,
	0x0c, 0xf2, 0xa2, 0xc4, 0xd5, 0xfe, 0x3d, 0x12, 0x3d, 0x26, 0xd9, 0xc2, 0x35, 0xd5, 0xcc, 0x29,
	0x29, 0xf7, 0xe6, 0xe7, 0xaf, 0x8f, 0x09, 0x49, 0x9a, 0x6f, 0x5f, 0x63, 0x26, 0xf6, 0x5d, 0xe5,
	0x15, 0x33, 0xe0, 0xf5, 0x26, 0x66, 0xf2, 0x4d, 0x21, 0x4f, 0x91, 0xf8, 0xfc, 0xba, 0x0a, 0x29,
	0x3a, 0x79, 0xaf, 0x42, 0xea, 0x1a, 0x8a, 0x57, 0x21, 0xe9, 0x38, 0x40, 0xfa, 0x24, 0x40, 0x26,
	0xd8, 0x7c, 0x51, 0xee, 0x9f, 0xa6, 0xbb, 0xa1, 0xb2, 0xca, 0xb5, 0xf5, 0x1c, 0x2c, 0x4f, 0xc1,
	0x16, 0xa5, 0xfb, 0xf1, 0x60, 0x4e, 0x3b, 0xa0, 0x85, 0xf6, 0x59, 0x80, 0xf1, 0xae, 0xd9, 0x23,
	0xae, 0xc4, 0x26, 0x8c, 0x9f, 0xd5, 0xd9, 0xd5, 0xeb, 0x89, 0x39, 0xda, 0x1a, 0x45, 0xcb, 0x4b,
	0x4b, 0xf1, 0x68, 0xd5, 0x68, 0xd8, 0xa6, 0x90, 0x2f, 0x1f, 0xc3, 0x4c, 0x24, 0x01, 0xb2, 0x0d,
	0x9e, 0xe4, 0xc5, 0x56, 0xd5, 0xf0, 0x4e, 0x1b, 0x27, 0xb2, 0x46, 0xea, 0x0a, 0xd3, 0x14, 0xd8,
	0x03, 0xa6, 0x4a, 0x0a, 0x55, 0x6c, 0xd1, 0xb7, 0x8b, 0x12, 0xf3, 0x04, 0xda, 0x32, 0xeb, 0xee,
	0x49, 0x9a, 0x2e, 0xaf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xc2, 0x9c, 0x0b, 0x8b, 0x09,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SymmetricCryptoServiceClient is the client API for SymmetricCryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SymmetricCryptoServiceClient interface {
	// Encrypts given plaintext with the specified key.
	Encrypt(ctx context.Context, in *SymmetricEncryptRequest, opts ...grpc.CallOption) (*SymmetricEncryptResponse, error)
	// Decrypts the given ciphertext with the specified key.
	Decrypt(ctx context.Context, in *SymmetricDecryptRequest, opts ...grpc.CallOption) (*SymmetricDecryptResponse, error)
	// Re-encrypts a ciphertext with the specified KMS key.
	ReEncrypt(ctx context.Context, in *SymmetricReEncryptRequest, opts ...grpc.CallOption) (*SymmetricReEncryptResponse, error)
	// Generates a new symmetric data encryption key (not a KMS key) and returns
	// the generated key as plaintext and as ciphertext encrypted with the specified symmetric KMS key.
	GenerateDataKey(ctx context.Context, in *GenerateDataKeyRequest, opts ...grpc.CallOption) (*GenerateDataKeyResponse, error)
}

type symmetricCryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSymmetricCryptoServiceClient(cc grpc.ClientConnInterface) SymmetricCryptoServiceClient {
	return &symmetricCryptoServiceClient{cc}
}

func (c *symmetricCryptoServiceClient) Encrypt(ctx context.Context, in *SymmetricEncryptRequest, opts ...grpc.CallOption) (*SymmetricEncryptResponse, error) {
	out := new(SymmetricEncryptResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.kms.v1.SymmetricCryptoService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) Decrypt(ctx context.Context, in *SymmetricDecryptRequest, opts ...grpc.CallOption) (*SymmetricDecryptResponse, error) {
	out := new(SymmetricDecryptResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.kms.v1.SymmetricCryptoService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) ReEncrypt(ctx context.Context, in *SymmetricReEncryptRequest, opts ...grpc.CallOption) (*SymmetricReEncryptResponse, error) {
	out := new(SymmetricReEncryptResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.kms.v1.SymmetricCryptoService/ReEncrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) GenerateDataKey(ctx context.Context, in *GenerateDataKeyRequest, opts ...grpc.CallOption) (*GenerateDataKeyResponse, error) {
	out := new(GenerateDataKeyResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.kms.v1.SymmetricCryptoService/GenerateDataKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SymmetricCryptoServiceServer is the server API for SymmetricCryptoService service.
type SymmetricCryptoServiceServer interface {
	// Encrypts given plaintext with the specified key.
	Encrypt(context.Context, *SymmetricEncryptRequest) (*SymmetricEncryptResponse, error)
	// Decrypts the given ciphertext with the specified key.
	Decrypt(context.Context, *SymmetricDecryptRequest) (*SymmetricDecryptResponse, error)
	// Re-encrypts a ciphertext with the specified KMS key.
	ReEncrypt(context.Context, *SymmetricReEncryptRequest) (*SymmetricReEncryptResponse, error)
	// Generates a new symmetric data encryption key (not a KMS key) and returns
	// the generated key as plaintext and as ciphertext encrypted with the specified symmetric KMS key.
	GenerateDataKey(context.Context, *GenerateDataKeyRequest) (*GenerateDataKeyResponse, error)
}

// UnimplementedSymmetricCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSymmetricCryptoServiceServer struct {
}

func (*UnimplementedSymmetricCryptoServiceServer) Encrypt(ctx context.Context, req *SymmetricEncryptRequest) (*SymmetricEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (*UnimplementedSymmetricCryptoServiceServer) Decrypt(ctx context.Context, req *SymmetricDecryptRequest) (*SymmetricDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (*UnimplementedSymmetricCryptoServiceServer) ReEncrypt(ctx context.Context, req *SymmetricReEncryptRequest) (*SymmetricReEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReEncrypt not implemented")
}
func (*UnimplementedSymmetricCryptoServiceServer) GenerateDataKey(ctx context.Context, req *GenerateDataKeyRequest) (*GenerateDataKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDataKey not implemented")
}

func RegisterSymmetricCryptoServiceServer(s *grpc.Server, srv SymmetricCryptoServiceServer) {
	s.RegisterService(&_SymmetricCryptoService_serviceDesc, srv)
}

func _SymmetricCryptoService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.kms.v1.SymmetricCryptoService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).Encrypt(ctx, req.(*SymmetricEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.kms.v1.SymmetricCryptoService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).Decrypt(ctx, req.(*SymmetricDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_ReEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricReEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).ReEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.kms.v1.SymmetricCryptoService/ReEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).ReEncrypt(ctx, req.(*SymmetricReEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_GenerateDataKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDataKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).GenerateDataKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.kms.v1.SymmetricCryptoService/GenerateDataKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).GenerateDataKey(ctx, req.(*GenerateDataKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SymmetricCryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.kms.v1.SymmetricCryptoService",
	HandlerType: (*SymmetricCryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _SymmetricCryptoService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _SymmetricCryptoService_Decrypt_Handler,
		},
		{
			MethodName: "ReEncrypt",
			Handler:    _SymmetricCryptoService_ReEncrypt_Handler,
		},
		{
			MethodName: "GenerateDataKey",
			Handler:    _SymmetricCryptoService_GenerateDataKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/kms/v1/symmetric_crypto_service.proto",
}
