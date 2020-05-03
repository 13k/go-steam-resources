// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: dota2/steam/network/steamnetworkingsockets_messages_certs.proto

package network

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CMsgSteamDatagramCertificate_EKeyType int32

const (
	CMsgSteamDatagramCertificate_INVALID CMsgSteamDatagramCertificate_EKeyType = 0
	CMsgSteamDatagramCertificate_ED25519 CMsgSteamDatagramCertificate_EKeyType = 1
)

// Enum value maps for CMsgSteamDatagramCertificate_EKeyType.
var (
	CMsgSteamDatagramCertificate_EKeyType_name = map[int32]string{
		0: "INVALID",
		1: "ED25519",
	}
	CMsgSteamDatagramCertificate_EKeyType_value = map[string]int32{
		"INVALID": 0,
		"ED25519": 1,
	}
)

func (x CMsgSteamDatagramCertificate_EKeyType) Enum() *CMsgSteamDatagramCertificate_EKeyType {
	p := new(CMsgSteamDatagramCertificate_EKeyType)
	*p = x
	return p
}

func (x CMsgSteamDatagramCertificate_EKeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMsgSteamDatagramCertificate_EKeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_enumTypes[0].Descriptor()
}

func (CMsgSteamDatagramCertificate_EKeyType) Type() protoreflect.EnumType {
	return &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_enumTypes[0]
}

func (x CMsgSteamDatagramCertificate_EKeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMsgSteamDatagramCertificate_EKeyType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMsgSteamDatagramCertificate_EKeyType(num)
	return nil
}

// Deprecated: Use CMsgSteamDatagramCertificate_EKeyType.Descriptor instead.
func (CMsgSteamDatagramCertificate_EKeyType) EnumDescriptor() ([]byte, []int) {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP(), []int{1, 0}
}

type CMsgSteamNetworkingIdentityLegacyBinary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SteamId        *uint64 `protobuf:"fixed64,16,opt,name=steam_id,json=steamId" json:"steam_id,omitempty"`
	XboxPairwiseId *string `protobuf:"bytes,17,opt,name=xbox_pairwise_id,json=xboxPairwiseId" json:"xbox_pairwise_id,omitempty"`
	GenericBytes   []byte  `protobuf:"bytes,2,opt,name=generic_bytes,json=genericBytes" json:"generic_bytes,omitempty"`
	GenericString  *string `protobuf:"bytes,3,opt,name=generic_string,json=genericString" json:"generic_string,omitempty"`
	Ipv6AndPort    []byte  `protobuf:"bytes,4,opt,name=ipv6_and_port,json=ipv6AndPort" json:"ipv6_and_port,omitempty"`
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) Reset() {
	*x = CMsgSteamNetworkingIdentityLegacyBinary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamNetworkingIdentityLegacyBinary) ProtoMessage() {}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamNetworkingIdentityLegacyBinary.ProtoReflect.Descriptor instead.
func (*CMsgSteamNetworkingIdentityLegacyBinary) Descriptor() ([]byte, []int) {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) GetSteamId() uint64 {
	if x != nil && x.SteamId != nil {
		return *x.SteamId
	}
	return 0
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) GetXboxPairwiseId() string {
	if x != nil && x.XboxPairwiseId != nil {
		return *x.XboxPairwiseId
	}
	return ""
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) GetGenericBytes() []byte {
	if x != nil {
		return x.GenericBytes
	}
	return nil
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) GetGenericString() string {
	if x != nil && x.GenericString != nil {
		return *x.GenericString
	}
	return ""
}

func (x *CMsgSteamNetworkingIdentityLegacyBinary) GetIpv6AndPort() []byte {
	if x != nil {
		return x.Ipv6AndPort
	}
	return nil
}

type CMsgSteamDatagramCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyType                 *CMsgSteamDatagramCertificate_EKeyType   `protobuf:"varint,1,opt,name=key_type,json=keyType,enum=dota2.steam.network.CMsgSteamDatagramCertificate_EKeyType,def=0" json:"key_type,omitempty"`
	KeyData                 []byte                                   `protobuf:"bytes,2,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	LegacySteamId           *uint64                                  `protobuf:"fixed64,4,opt,name=legacy_steam_id,json=legacySteamId" json:"legacy_steam_id,omitempty"`
	LegacyIdentityBinary    *CMsgSteamNetworkingIdentityLegacyBinary `protobuf:"bytes,11,opt,name=legacy_identity_binary,json=legacyIdentityBinary" json:"legacy_identity_binary,omitempty"`
	IdentityString          *string                                  `protobuf:"bytes,12,opt,name=identity_string,json=identityString" json:"identity_string,omitempty"`
	GameserverDatacenterIds []uint32                                 `protobuf:"fixed32,5,rep,name=gameserver_datacenter_ids,json=gameserverDatacenterIds" json:"gameserver_datacenter_ids,omitempty"`
	TimeCreated             *uint32                                  `protobuf:"fixed32,8,opt,name=time_created,json=timeCreated" json:"time_created,omitempty"`
	TimeExpiry              *uint32                                  `protobuf:"fixed32,9,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AppIds                  []uint32                                 `protobuf:"varint,10,rep,name=app_ids,json=appIds" json:"app_ids,omitempty"`
}

// Default values for CMsgSteamDatagramCertificate fields.
const (
	Default_CMsgSteamDatagramCertificate_KeyType = CMsgSteamDatagramCertificate_INVALID
)

func (x *CMsgSteamDatagramCertificate) Reset() {
	*x = CMsgSteamDatagramCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramCertificate) ProtoMessage() {}

func (x *CMsgSteamDatagramCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramCertificate.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramCertificate) Descriptor() ([]byte, []int) {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgSteamDatagramCertificate) GetKeyType() CMsgSteamDatagramCertificate_EKeyType {
	if x != nil && x.KeyType != nil {
		return *x.KeyType
	}
	return Default_CMsgSteamDatagramCertificate_KeyType
}

func (x *CMsgSteamDatagramCertificate) GetKeyData() []byte {
	if x != nil {
		return x.KeyData
	}
	return nil
}

func (x *CMsgSteamDatagramCertificate) GetLegacySteamId() uint64 {
	if x != nil && x.LegacySteamId != nil {
		return *x.LegacySteamId
	}
	return 0
}

func (x *CMsgSteamDatagramCertificate) GetLegacyIdentityBinary() *CMsgSteamNetworkingIdentityLegacyBinary {
	if x != nil {
		return x.LegacyIdentityBinary
	}
	return nil
}

func (x *CMsgSteamDatagramCertificate) GetIdentityString() string {
	if x != nil && x.IdentityString != nil {
		return *x.IdentityString
	}
	return ""
}

func (x *CMsgSteamDatagramCertificate) GetGameserverDatacenterIds() []uint32 {
	if x != nil {
		return x.GameserverDatacenterIds
	}
	return nil
}

func (x *CMsgSteamDatagramCertificate) GetTimeCreated() uint32 {
	if x != nil && x.TimeCreated != nil {
		return *x.TimeCreated
	}
	return 0
}

func (x *CMsgSteamDatagramCertificate) GetTimeExpiry() uint32 {
	if x != nil && x.TimeExpiry != nil {
		return *x.TimeExpiry
	}
	return 0
}

func (x *CMsgSteamDatagramCertificate) GetAppIds() []uint32 {
	if x != nil {
		return x.AppIds
	}
	return nil
}

type CMsgSteamDatagramCertificateSigned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert           []byte  `protobuf:"bytes,4,opt,name=cert" json:"cert,omitempty"`
	CaKeyId        *uint64 `protobuf:"fixed64,5,opt,name=ca_key_id,json=caKeyId" json:"ca_key_id,omitempty"`
	CaSignature    []byte  `protobuf:"bytes,6,opt,name=ca_signature,json=caSignature" json:"ca_signature,omitempty"`
	PrivateKeyData []byte  `protobuf:"bytes,1,opt,name=private_key_data,json=privateKeyData" json:"private_key_data,omitempty"`
}

func (x *CMsgSteamDatagramCertificateSigned) Reset() {
	*x = CMsgSteamDatagramCertificateSigned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramCertificateSigned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramCertificateSigned) ProtoMessage() {}

func (x *CMsgSteamDatagramCertificateSigned) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramCertificateSigned.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramCertificateSigned) Descriptor() ([]byte, []int) {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgSteamDatagramCertificateSigned) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *CMsgSteamDatagramCertificateSigned) GetCaKeyId() uint64 {
	if x != nil && x.CaKeyId != nil {
		return *x.CaKeyId
	}
	return 0
}

func (x *CMsgSteamDatagramCertificateSigned) GetCaSignature() []byte {
	if x != nil {
		return x.CaSignature
	}
	return nil
}

func (x *CMsgSteamDatagramCertificateSigned) GetPrivateKeyData() []byte {
	if x != nil {
		return x.PrivateKeyData
	}
	return nil
}

type CMsgSteamDatagramCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert *CMsgSteamDatagramCertificate `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (x *CMsgSteamDatagramCertificateRequest) Reset() {
	*x = CMsgSteamDatagramCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramCertificateRequest) ProtoMessage() {}

func (x *CMsgSteamDatagramCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramCertificateRequest.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramCertificateRequest) Descriptor() ([]byte, []int) {
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgSteamDatagramCertificateRequest) GetCert() *CMsgSteamDatagramCertificate {
	if x != nil {
		return x.Cert
	}
	return nil
}

var File_dota2_steam_network_steamnetworkingsockets_messages_certs_proto protoreflect.FileDescriptor

var file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0xde, 0x01, 0x0a, 0x27, 0x43, 0x4d, 0x73, 0x67, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x78, 0x62, 0x6f, 0x78, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x77, 0x69, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x78, 0x62, 0x6f, 0x78, 0x50, 0x61, 0x69,
	0x72, 0x77, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x61, 0x6e, 0x64, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x36,
	0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x9d, 0x04, 0x0a, 0x1c, 0x43, 0x4d, 0x73, 0x67,
	0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x64, 0x6f, 0x74,
	0x61, 0x32, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x4b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0d, 0x6c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x72, 0x0a, 0x16, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x6f,
	0x74, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x14, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x19, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x07, 0x52, 0x17, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x73, 0x22, 0x24, 0x0a, 0x08, 0x45, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x44,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x01, 0x22, 0xa1, 0x01, 0x0a, 0x22, 0x43, 0x4d, 0x73, 0x67,
	0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65,
	0x72, 0x74, 0x12, 0x1a, 0x0a, 0x09, 0x63, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x63, 0x61, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x23, 0x43,
	0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x42, 0x47, 0x48, 0x01, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x74, 0x61,
	0x32, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x80,
	0x01, 0x00,
}

var (
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescOnce sync.Once
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescData = file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDesc
)

func file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescGZIP() []byte {
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescOnce.Do(func() {
		file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescData)
	})
	return file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDescData
}

var file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_goTypes = []interface{}{
	(CMsgSteamDatagramCertificate_EKeyType)(0),      // 0: dota2.steam.network.CMsgSteamDatagramCertificate.EKeyType
	(*CMsgSteamNetworkingIdentityLegacyBinary)(nil), // 1: dota2.steam.network.CMsgSteamNetworkingIdentityLegacyBinary
	(*CMsgSteamDatagramCertificate)(nil),            // 2: dota2.steam.network.CMsgSteamDatagramCertificate
	(*CMsgSteamDatagramCertificateSigned)(nil),      // 3: dota2.steam.network.CMsgSteamDatagramCertificateSigned
	(*CMsgSteamDatagramCertificateRequest)(nil),     // 4: dota2.steam.network.CMsgSteamDatagramCertificateRequest
}
var file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_depIdxs = []int32{
	0, // 0: dota2.steam.network.CMsgSteamDatagramCertificate.key_type:type_name -> dota2.steam.network.CMsgSteamDatagramCertificate.EKeyType
	1, // 1: dota2.steam.network.CMsgSteamDatagramCertificate.legacy_identity_binary:type_name -> dota2.steam.network.CMsgSteamNetworkingIdentityLegacyBinary
	2, // 2: dota2.steam.network.CMsgSteamDatagramCertificateRequest.cert:type_name -> dota2.steam.network.CMsgSteamDatagramCertificate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_init() }
func file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_init() {
	if File_dota2_steam_network_steamnetworkingsockets_messages_certs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSteamNetworkingIdentityLegacyBinary); i {
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
		file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSteamDatagramCertificate); i {
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
		file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSteamDatagramCertificateSigned); i {
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
		file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSteamDatagramCertificateRequest); i {
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
			RawDescriptor: file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_goTypes,
		DependencyIndexes: file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_depIdxs,
		EnumInfos:         file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_enumTypes,
		MessageInfos:      file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_msgTypes,
	}.Build()
	File_dota2_steam_network_steamnetworkingsockets_messages_certs_proto = out.File
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_rawDesc = nil
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_goTypes = nil
	file_dota2_steam_network_steamnetworkingsockets_messages_certs_proto_depIdxs = nil
}
