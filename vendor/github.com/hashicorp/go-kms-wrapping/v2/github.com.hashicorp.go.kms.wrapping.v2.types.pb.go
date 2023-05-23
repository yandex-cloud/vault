// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com.hashicorp.go.kms.wrapping.v2.types.proto

package wrapping

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// WrapperConfig is the result of a call to SetConfig on a wrapper, returning
// relevant information about the wrapper and its updated configuration
type WrapperConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string `protobuf:"bytes,10,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WrapperConfig) Reset() {
	*x = WrapperConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrapperConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrapperConfig) ProtoMessage() {}

func (x *WrapperConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrapperConfig.ProtoReflect.Descriptor instead.
func (*WrapperConfig) Descriptor() ([]byte, []int) {
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP(), []int{0}
}

func (x *WrapperConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// EnvelopeInfo contains the information necessary to perfom encryption or
// decryption in an envelope fashion
type EnvelopeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ciphertext is the ciphertext from the envelope
	Ciphertext []byte `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// Key is the key used in the envelope
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// IV is the initialization value used during encryption in the envelope
	Iv []byte `protobuf:"bytes,3,opt,name=iv,proto3" json:"iv,omitempty"`
}

func (x *EnvelopeInfo) Reset() {
	*x = EnvelopeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvelopeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvelopeInfo) ProtoMessage() {}

func (x *EnvelopeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvelopeInfo.ProtoReflect.Descriptor instead.
func (*EnvelopeInfo) Descriptor() ([]byte, []int) {
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP(), []int{1}
}

func (x *EnvelopeInfo) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *EnvelopeInfo) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *EnvelopeInfo) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

// BlobInfo contains information about the encrypted value along with
// information about the key used to encrypt it
type BlobInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ciphertext is the encrypted bytes
	Ciphertext []byte `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// IV is the initialization value used during encryption
	Iv []byte `protobuf:"bytes,2,opt,name=iv,proto3" json:"iv,omitempty"`
	// HMAC is the bytes of the HMAC, if any
	Hmac []byte `protobuf:"bytes,3,opt,name=hmac,proto3" json:"hmac,omitempty"`
	// Wrapped can be used by the client to indicate whether Ciphertext actually
	// contains wrapped data or not. This can be useful if you want to reuse the
	// same struct to pass data along before and after wrapping. Deprecated in
	// favor of plaintext.
	//
	// Deprecated: Do not use.
	Wrapped bool `protobuf:"varint,4,opt,name=wrapped,proto3" json:"wrapped,omitempty"`
	// Plaintext can be used to allow the same struct to be used to pass data
	// along before and after (un)wrapping.
	Plaintext []byte `protobuf:"bytes,7,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	// KeyInfo contains information about the key that was used to create this value
	KeyInfo *KeyInfo `protobuf:"bytes,5,opt,name=key_info,json=keyInfo,proto3" json:"key_info,omitempty"`
	// ValuePath can be used by the client to store information about where the
	// value came from. Deprecated in favor of client_data.
	//
	// Deprecated: Do not use.
	ValuePath string `protobuf:"bytes,6,opt,name=value_path,json=valuePath,proto3" json:"value_path,omitempty"`
	// ClientData can be used by the client to store extra information, for
	// instance, the location/provenance of where an encrypted value came from
	// (useful for associating AAD to the encrypted value).
	ClientData *_struct.Struct `protobuf:"bytes,8,opt,name=client_data,json=clientData,proto3" json:"client_data,omitempty"`
}

func (x *BlobInfo) Reset() {
	*x = BlobInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobInfo) ProtoMessage() {}

func (x *BlobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobInfo.ProtoReflect.Descriptor instead.
func (*BlobInfo) Descriptor() ([]byte, []int) {
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP(), []int{2}
}

func (x *BlobInfo) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *BlobInfo) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

func (x *BlobInfo) GetHmac() []byte {
	if x != nil {
		return x.Hmac
	}
	return nil
}

// Deprecated: Do not use.
func (x *BlobInfo) GetWrapped() bool {
	if x != nil {
		return x.Wrapped
	}
	return false
}

func (x *BlobInfo) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *BlobInfo) GetKeyInfo() *KeyInfo {
	if x != nil {
		return x.KeyInfo
	}
	return nil
}

// Deprecated: Do not use.
func (x *BlobInfo) GetValuePath() string {
	if x != nil {
		return x.ValuePath
	}
	return ""
}

func (x *BlobInfo) GetClientData() *_struct.Struct {
	if x != nil {
		return x.ClientData
	}
	return nil
}

// KeyInfo contains information regarding which Wrapper key was used to
// encrypt the entry
type KeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mechanism is the method used by the wrapper to encrypt and sign the
	// data as defined by the wrapper.
	Mechanism     uint64 `protobuf:"varint,1,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	HmacMechanism uint64 `protobuf:"varint,2,opt,name=hmac_mechanism,json=hmacMechanism,proto3" json:"hmac_mechanism,omitempty"`
	// This is an opaque ID used by the wrapper to identify the specific key to
	// use as defined by the wrapper. This could be a version, key label, or
	// something else.
	KeyId     string `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	HmacKeyId string `protobuf:"bytes,4,opt,name=hmac_key_id,json=hmacKeyId,proto3" json:"hmac_key_id,omitempty"`
	// These value are used when generating our own data encryption keys
	// and encrypting them using the wrapper
	WrappedKey []byte `protobuf:"bytes,5,opt,name=wrapped_key,json=wrappedKey,proto3" json:"wrapped_key,omitempty"`
	// Mechanism specific flags
	Flags uint64 `protobuf:"varint,6,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *KeyInfo) Reset() {
	*x = KeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyInfo) ProtoMessage() {}

func (x *KeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyInfo.ProtoReflect.Descriptor instead.
func (*KeyInfo) Descriptor() ([]byte, []int) {
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP(), []int{3}
}

func (x *KeyInfo) GetMechanism() uint64 {
	if x != nil {
		return x.Mechanism
	}
	return 0
}

func (x *KeyInfo) GetHmacMechanism() uint64 {
	if x != nil {
		return x.HmacMechanism
	}
	return 0
}

func (x *KeyInfo) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *KeyInfo) GetHmacKeyId() string {
	if x != nil {
		return x.HmacKeyId
	}
	return ""
}

func (x *KeyInfo) GetWrappedKey() []byte {
	if x != nil {
		return x.WrappedKey
	}
	return nil
}

func (x *KeyInfo) GetFlags() uint64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

// Options holds options common to all wrappers
type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key ID being specified
	WithKeyId string `protobuf:"bytes,10,opt,name=with_key_id,json=withKeyId,proto3" json:"with_key_id,omitempty"`
	// The AAD bytes, if any
	WithAad []byte `protobuf:"bytes,20,opt,name=with_aad,json=withAad,proto3" json:"with_aad,omitempty"`
	WithIv  []byte `protobuf:"bytes,12,opt,name=with_iv,json=withIv,proto3" json:"with_iv,omitempty"`
	// Wrapper-specific configuration to pass along
	WithConfigMap map[string]string `protobuf:"bytes,30,rep,name=with_config_map,json=withConfigMap,proto3" json:"with_config_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP(), []int{4}
}

func (x *Options) GetWithKeyId() string {
	if x != nil {
		return x.WithKeyId
	}
	return ""
}

func (x *Options) GetWithAad() []byte {
	if x != nil {
		return x.WithAad
	}
	return nil
}

func (x *Options) GetWithIv() []byte {
	if x != nil {
		return x.WithIv
	}
	return nil
}

func (x *Options) GetWithConfigMap() map[string]string {
	if x != nil {
		return x.WithConfigMap
	}
	return nil
}

var File_github_com_hashicorp_go_kms_wrapping_v2_types_proto protoreflect.FileDescriptor

var file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x67, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x67, 0x6f, 0x2e, 0x6b,
	0x6d, 0x73, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x67, 0x6f,
	0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x0c, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x76, 0x22, 0xba, 0x02, 0x0a, 0x08,
	0x42, 0x6c, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6d, 0x61, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x6d, 0x61, 0x63, 0x12, 0x1c, 0x0a, 0x07,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x51, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x67, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x38,
	0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x4b, 0x65, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x73, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x73, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x68, 0x6d, 0x61, 0x63,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0b, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6d, 0x61, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x61, 0x64, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x77, 0x69, 0x74, 0x68, 0x41, 0x61, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x76, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x77, 0x69, 0x74, 0x68, 0x49, 0x76, 0x12, 0x71, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x49, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x67, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x77, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x1a, 0x40, 0x0a, 0x12, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x3b, 0x77, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescOnce sync.Once
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescData = file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDesc
)

func file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescGZIP() []byte {
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescOnce.Do(func() {
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescData)
	})
	return file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDescData
}

var file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_goTypes = []interface{}{
	(*WrapperConfig)(nil),  // 0: github.com.hashicorp.go.kms.wrapping.v2.types.WrapperConfig
	(*EnvelopeInfo)(nil),   // 1: github.com.hashicorp.go.kms.wrapping.v2.types.EnvelopeInfo
	(*BlobInfo)(nil),       // 2: github.com.hashicorp.go.kms.wrapping.v2.types.BlobInfo
	(*KeyInfo)(nil),        // 3: github.com.hashicorp.go.kms.wrapping.v2.types.KeyInfo
	(*Options)(nil),        // 4: github.com.hashicorp.go.kms.wrapping.v2.types.Options
	nil,                    // 5: github.com.hashicorp.go.kms.wrapping.v2.types.WrapperConfig.MetadataEntry
	nil,                    // 6: github.com.hashicorp.go.kms.wrapping.v2.types.Options.WithConfigMapEntry
	(*_struct.Struct)(nil), // 7: google.protobuf.Struct
}
var file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_depIdxs = []int32{
	5, // 0: github.com.hashicorp.go.kms.wrapping.v2.types.WrapperConfig.metadata:type_name -> github.com.hashicorp.go.kms.wrapping.v2.types.WrapperConfig.MetadataEntry
	3, // 1: github.com.hashicorp.go.kms.wrapping.v2.types.BlobInfo.key_info:type_name -> github.com.hashicorp.go.kms.wrapping.v2.types.KeyInfo
	7, // 2: github.com.hashicorp.go.kms.wrapping.v2.types.BlobInfo.client_data:type_name -> google.protobuf.Struct
	6, // 3: github.com.hashicorp.go.kms.wrapping.v2.types.Options.with_config_map:type_name -> github.com.hashicorp.go.kms.wrapping.v2.types.Options.WithConfigMapEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_init() }
func file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_init() {
	if File_github_com_hashicorp_go_kms_wrapping_v2_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrapperConfig); i {
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
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvelopeInfo); i {
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
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobInfo); i {
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
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyInfo); i {
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
		file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_goTypes,
		DependencyIndexes: file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_depIdxs,
		MessageInfos:      file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_msgTypes,
	}.Build()
	File_github_com_hashicorp_go_kms_wrapping_v2_types_proto = out.File
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_rawDesc = nil
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_goTypes = nil
	file_github_com_hashicorp_go_kms_wrapping_v2_types_proto_depIdxs = nil
}
