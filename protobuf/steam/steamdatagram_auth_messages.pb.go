// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steamdatagram_auth_messages.proto

package steam // import "github.com/13k/go-steam-resources/protobuf/steam"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CMsgSteamDatagramCertificate_EKeyType int32

const (
	CMsgSteamDatagramCertificate_INVALID CMsgSteamDatagramCertificate_EKeyType = 0
	CMsgSteamDatagramCertificate_ED25519 CMsgSteamDatagramCertificate_EKeyType = 1
)

var CMsgSteamDatagramCertificate_EKeyType_name = map[int32]string{
	0: "INVALID",
	1: "ED25519",
}
var CMsgSteamDatagramCertificate_EKeyType_value = map[string]int32{
	"INVALID": 0,
	"ED25519": 1,
}

func (x CMsgSteamDatagramCertificate_EKeyType) Enum() *CMsgSteamDatagramCertificate_EKeyType {
	p := new(CMsgSteamDatagramCertificate_EKeyType)
	*p = x
	return p
}
func (x CMsgSteamDatagramCertificate_EKeyType) String() string {
	return proto.EnumName(CMsgSteamDatagramCertificate_EKeyType_name, int32(x))
}
func (x *CMsgSteamDatagramCertificate_EKeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMsgSteamDatagramCertificate_EKeyType_value, data, "CMsgSteamDatagramCertificate_EKeyType")
	if err != nil {
		return err
	}
	*x = CMsgSteamDatagramCertificate_EKeyType(value)
	return nil
}
func (CMsgSteamDatagramCertificate_EKeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{2, 0}
}

type CMsgSteamDatagramRelayAuthTicket struct {
	TimeExpiry           *uint32                                        `protobuf:"fixed32,1,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AuthorizedSteamId    *uint64                                        `protobuf:"fixed64,2,opt,name=authorized_steam_id,json=authorizedSteamId" json:"authorized_steam_id,omitempty"`
	AuthorizedPublicIp   *uint32                                        `protobuf:"fixed32,3,opt,name=authorized_public_ip,json=authorizedPublicIp" json:"authorized_public_ip,omitempty"`
	GameserverSteamId    *uint64                                        `protobuf:"fixed64,4,opt,name=gameserver_steam_id,json=gameserverSteamId" json:"gameserver_steam_id,omitempty"`
	GameserverNetId      *uint64                                        `protobuf:"fixed64,5,opt,name=gameserver_net_id,json=gameserverNetId" json:"gameserver_net_id,omitempty"`
	LegacySignature      []byte                                         `protobuf:"bytes,6,opt,name=legacy_signature,json=legacySignature" json:"legacy_signature,omitempty"`
	AppId                *uint32                                        `protobuf:"varint,7,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	GameserverPopId      *uint32                                        `protobuf:"fixed32,9,opt,name=gameserver_pop_id,json=gameserverPopId" json:"gameserver_pop_id,omitempty"`
	VirtualPort          *uint32                                        `protobuf:"varint,10,opt,name=virtual_port,json=virtualPort" json:"virtual_port,omitempty"`
	ExtraFields          []*CMsgSteamDatagramRelayAuthTicket_ExtraField `protobuf:"bytes,8,rep,name=extra_fields,json=extraFields" json:"extra_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *CMsgSteamDatagramRelayAuthTicket) Reset()         { *m = CMsgSteamDatagramRelayAuthTicket{} }
func (m *CMsgSteamDatagramRelayAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramRelayAuthTicket) ProtoMessage()    {}
func (*CMsgSteamDatagramRelayAuthTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{0}
}
func (m *CMsgSteamDatagramRelayAuthTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramRelayAuthTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramRelayAuthTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket.Merge(dst, src)
}
func (m *CMsgSteamDatagramRelayAuthTicket) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket.Size(m)
}
func (m *CMsgSteamDatagramRelayAuthTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket proto.InternalMessageInfo

func (m *CMsgSteamDatagramRelayAuthTicket) GetTimeExpiry() uint32 {
	if m != nil && m.TimeExpiry != nil {
		return *m.TimeExpiry
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedSteamId() uint64 {
	if m != nil && m.AuthorizedSteamId != nil {
		return *m.AuthorizedSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedPublicIp() uint32 {
	if m != nil && m.AuthorizedPublicIp != nil {
		return *m.AuthorizedPublicIp
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverSteamId() uint64 {
	if m != nil && m.GameserverSteamId != nil {
		return *m.GameserverSteamId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverNetId() uint64 {
	if m != nil && m.GameserverNetId != nil {
		return *m.GameserverNetId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetLegacySignature() []byte {
	if m != nil {
		return m.LegacySignature
	}
	return nil
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetGameserverPopId() uint32 {
	if m != nil && m.GameserverPopId != nil {
		return *m.GameserverPopId
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetVirtualPort() uint32 {
	if m != nil && m.VirtualPort != nil {
		return *m.VirtualPort
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket) GetExtraFields() []*CMsgSteamDatagramRelayAuthTicket_ExtraField {
	if m != nil {
		return m.ExtraFields
	}
	return nil
}

type CMsgSteamDatagramRelayAuthTicket_ExtraField struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StringValue          *string  `protobuf:"bytes,2,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	Int64Value           *int64   `protobuf:"zigzag64,3,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Fixed64Value         *uint64  `protobuf:"fixed64,5,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) Reset() {
	*m = CMsgSteamDatagramRelayAuthTicket_ExtraField{}
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) ProtoMessage() {}
func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{0, 0}
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramRelayAuthTicket_ExtraField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField.Merge(dst, src)
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField.Size(m)
}
func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramRelayAuthTicket_ExtraField proto.InternalMessageInfo

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetInt64Value() int64 {
	if m != nil && m.Int64Value != nil {
		return *m.Int64Value
	}
	return 0
}

func (m *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetFixed64Value() uint64 {
	if m != nil && m.Fixed64Value != nil {
		return *m.Fixed64Value
	}
	return 0
}

type CMsgSteamDatagramSignedRelayAuthTicket struct {
	ReservedDoNotUse     *uint64  `protobuf:"fixed64,1,opt,name=reserved_do_not_use,json=reservedDoNotUse" json:"reserved_do_not_use,omitempty"`
	KeyId                *uint64  `protobuf:"fixed64,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Ticket               []byte   `protobuf:"bytes,3,opt,name=ticket" json:"ticket,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) Reset() {
	*m = CMsgSteamDatagramSignedRelayAuthTicket{}
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramSignedRelayAuthTicket) ProtoMessage()    {}
func (*CMsgSteamDatagramSignedRelayAuthTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{1}
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramSignedRelayAuthTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket.Merge(dst, src)
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket.Size(m)
}
func (m *CMsgSteamDatagramSignedRelayAuthTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramSignedRelayAuthTicket proto.InternalMessageInfo

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetReservedDoNotUse() uint64 {
	if m != nil && m.ReservedDoNotUse != nil {
		return *m.ReservedDoNotUse
	}
	return 0
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetKeyId() uint64 {
	if m != nil && m.KeyId != nil {
		return *m.KeyId
	}
	return 0
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func (m *CMsgSteamDatagramSignedRelayAuthTicket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type CMsgSteamDatagramCertificate struct {
	KeyType                 *CMsgSteamDatagramCertificate_EKeyType `protobuf:"varint,1,opt,name=key_type,json=keyType,enum=CMsgSteamDatagramCertificate_EKeyType,def=0" json:"key_type,omitempty"`
	KeyData                 []byte                                 `protobuf:"bytes,2,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	SteamId                 *uint64                                `protobuf:"fixed64,4,opt,name=steam_id,json=steamId" json:"steam_id,omitempty"`
	GameserverDatacenterIds []uint32                               `protobuf:"fixed32,5,rep,name=gameserver_datacenter_ids,json=gameserverDatacenterIds" json:"gameserver_datacenter_ids,omitempty"`
	TimeCreated             *uint32                                `protobuf:"fixed32,8,opt,name=time_created,json=timeCreated" json:"time_created,omitempty"`
	TimeExpiry              *uint32                                `protobuf:"fixed32,9,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AppId                   *uint32                                `protobuf:"varint,10,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                               `json:"-"`
	XXX_unrecognized        []byte                                 `json:"-"`
	XXX_sizecache           int32                                  `json:"-"`
}

func (m *CMsgSteamDatagramCertificate) Reset()         { *m = CMsgSteamDatagramCertificate{} }
func (m *CMsgSteamDatagramCertificate) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramCertificate) ProtoMessage()    {}
func (*CMsgSteamDatagramCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{2}
}
func (m *CMsgSteamDatagramCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramCertificate.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramCertificate.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramCertificate.Merge(dst, src)
}
func (m *CMsgSteamDatagramCertificate) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramCertificate.Size(m)
}
func (m *CMsgSteamDatagramCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramCertificate proto.InternalMessageInfo

const Default_CMsgSteamDatagramCertificate_KeyType CMsgSteamDatagramCertificate_EKeyType = CMsgSteamDatagramCertificate_INVALID

func (m *CMsgSteamDatagramCertificate) GetKeyType() CMsgSteamDatagramCertificate_EKeyType {
	if m != nil && m.KeyType != nil {
		return *m.KeyType
	}
	return Default_CMsgSteamDatagramCertificate_KeyType
}

func (m *CMsgSteamDatagramCertificate) GetKeyData() []byte {
	if m != nil {
		return m.KeyData
	}
	return nil
}

func (m *CMsgSteamDatagramCertificate) GetSteamId() uint64 {
	if m != nil && m.SteamId != nil {
		return *m.SteamId
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetGameserverDatacenterIds() []uint32 {
	if m != nil {
		return m.GameserverDatacenterIds
	}
	return nil
}

func (m *CMsgSteamDatagramCertificate) GetTimeCreated() uint32 {
	if m != nil && m.TimeCreated != nil {
		return *m.TimeCreated
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetTimeExpiry() uint32 {
	if m != nil && m.TimeExpiry != nil {
		return *m.TimeExpiry
	}
	return 0
}

func (m *CMsgSteamDatagramCertificate) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CMsgSteamDatagramCertificateSigned struct {
	Cert                 []byte   `protobuf:"bytes,4,opt,name=cert" json:"cert,omitempty"`
	CaKeyId              *uint64  `protobuf:"fixed64,5,opt,name=ca_key_id,json=caKeyId" json:"ca_key_id,omitempty"`
	CaSignature          []byte   `protobuf:"bytes,6,opt,name=ca_signature,json=caSignature" json:"ca_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamDatagramCertificateSigned) Reset()         { *m = CMsgSteamDatagramCertificateSigned{} }
func (m *CMsgSteamDatagramCertificateSigned) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramCertificateSigned) ProtoMessage()    {}
func (*CMsgSteamDatagramCertificateSigned) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{3}
}
func (m *CMsgSteamDatagramCertificateSigned) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramCertificateSigned.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramCertificateSigned) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramCertificateSigned.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramCertificateSigned) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramCertificateSigned.Merge(dst, src)
}
func (m *CMsgSteamDatagramCertificateSigned) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramCertificateSigned.Size(m)
}
func (m *CMsgSteamDatagramCertificateSigned) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramCertificateSigned.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramCertificateSigned proto.InternalMessageInfo

func (m *CMsgSteamDatagramCertificateSigned) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CMsgSteamDatagramCertificateSigned) GetCaKeyId() uint64 {
	if m != nil && m.CaKeyId != nil {
		return *m.CaKeyId
	}
	return 0
}

func (m *CMsgSteamDatagramCertificateSigned) GetCaSignature() []byte {
	if m != nil {
		return m.CaSignature
	}
	return nil
}

type CMsgSteamDatagramCachedCredentialsForApp struct {
	PrivateKey           []byte   `protobuf:"bytes,1,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	Cert                 []byte   `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	RelayTickets         [][]byte `protobuf:"bytes,3,rep,name=relay_tickets,json=relayTickets" json:"relay_tickets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgSteamDatagramCachedCredentialsForApp) Reset() {
	*m = CMsgSteamDatagramCachedCredentialsForApp{}
}
func (m *CMsgSteamDatagramCachedCredentialsForApp) String() string { return proto.CompactTextString(m) }
func (*CMsgSteamDatagramCachedCredentialsForApp) ProtoMessage()    {}
func (*CMsgSteamDatagramCachedCredentialsForApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93, []int{4}
}
func (m *CMsgSteamDatagramCachedCredentialsForApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp.Unmarshal(m, b)
}
func (m *CMsgSteamDatagramCachedCredentialsForApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp.Marshal(b, m, deterministic)
}
func (dst *CMsgSteamDatagramCachedCredentialsForApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp.Merge(dst, src)
}
func (m *CMsgSteamDatagramCachedCredentialsForApp) XXX_Size() int {
	return xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp.Size(m)
}
func (m *CMsgSteamDatagramCachedCredentialsForApp) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgSteamDatagramCachedCredentialsForApp proto.InternalMessageInfo

func (m *CMsgSteamDatagramCachedCredentialsForApp) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *CMsgSteamDatagramCachedCredentialsForApp) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *CMsgSteamDatagramCachedCredentialsForApp) GetRelayTickets() [][]byte {
	if m != nil {
		return m.RelayTickets
	}
	return nil
}

func init() {
	proto.RegisterType((*CMsgSteamDatagramRelayAuthTicket)(nil), "CMsgSteamDatagramRelayAuthTicket")
	proto.RegisterType((*CMsgSteamDatagramRelayAuthTicket_ExtraField)(nil), "CMsgSteamDatagramRelayAuthTicket.ExtraField")
	proto.RegisterType((*CMsgSteamDatagramSignedRelayAuthTicket)(nil), "CMsgSteamDatagramSignedRelayAuthTicket")
	proto.RegisterType((*CMsgSteamDatagramCertificate)(nil), "CMsgSteamDatagramCertificate")
	proto.RegisterType((*CMsgSteamDatagramCertificateSigned)(nil), "CMsgSteamDatagramCertificateSigned")
	proto.RegisterType((*CMsgSteamDatagramCachedCredentialsForApp)(nil), "CMsgSteamDatagramCachedCredentialsForApp")
	proto.RegisterEnum("CMsgSteamDatagramCertificate_EKeyType", CMsgSteamDatagramCertificate_EKeyType_name, CMsgSteamDatagramCertificate_EKeyType_value)
}

func init() {
	proto.RegisterFile("steamdatagram_auth_messages.proto", fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93)
}

var fileDescriptor_steamdatagram_auth_messages_63a0e334c2494b93 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0xad, 0xe2, 0xc4, 0x8e, 0xaf, 0x95, 0x35, 0x63, 0xf7, 0xa1, 0x16, 0x05, 0xa6, 0xb8, 0x43,
	0xa1, 0x0d, 0x8b, 0xd3, 0x66, 0xcb, 0x80, 0xf5, 0x2d, 0xb3, 0x5d, 0xc0, 0xcb, 0x96, 0x05, 0x6a,
	0xd7, 0x87, 0xbd, 0x10, 0x8c, 0x78, 0x2d, 0x13, 0xb6, 0x25, 0x82, 0xa4, 0xbc, 0x68, 0x4f, 0x7b,
	0xd8, 0xcb, 0xfe, 0xc7, 0xfe, 0xd6, 0x7e, 0xc3, 0xfe, 0xc2, 0x40, 0x4a, 0xb6, 0x8c, 0x18, 0xe8,
	0xde, 0xe4, 0x73, 0x0e, 0xef, 0xbd, 0xe4, 0x39, 0xd7, 0x70, 0xa2, 0x0d, 0xb2, 0x25, 0x67, 0x86,
	0xa5, 0x8a, 0x2d, 0x29, 0x2b, 0xcc, 0x8c, 0x2e, 0x51, 0x6b, 0x96, 0xa2, 0x1e, 0x48, 0x95, 0x9b,
	0xbc, 0xff, 0xef, 0x3e, 0x84, 0xc3, 0x9f, 0x74, 0xfa, 0xc6, 0x2a, 0x47, 0xb5, 0x32, 0xc6, 0x05,
	0x2b, 0x2f, 0x0b, 0x33, 0x7b, 0x2b, 0x92, 0x39, 0x1a, 0xf2, 0x19, 0xf4, 0x8c, 0x58, 0x22, 0xc5,
	0x3b, 0x29, 0x54, 0x19, 0x78, 0xa1, 0x17, 0x75, 0x62, 0xb0, 0xd0, 0xd8, 0x21, 0x64, 0x00, 0x8f,
	0x6c, 0xf1, 0x5c, 0x89, 0xdf, 0x91, 0x53, 0xd7, 0x95, 0x0a, 0x1e, 0xec, 0x85, 0x5e, 0xd4, 0x8e,
	0x3f, 0x6c, 0x28, 0xd7, 0x65, 0xc2, 0xc9, 0x0b, 0xf8, 0x68, 0x4b, 0x2f, 0x8b, 0xdb, 0x85, 0x48,
	0xa8, 0x90, 0x41, 0xcb, 0x55, 0x26, 0x0d, 0x77, 0xe3, 0xa8, 0x89, 0xb4, 0x1d, 0x52, 0xb6, 0x44,
	0x8d, 0x6a, 0x85, 0xaa, 0xe9, 0xb0, 0x5f, 0x75, 0x68, 0xa8, 0x75, 0x87, 0x2f, 0x61, 0x0b, 0xa4,
	0x19, 0x1a, 0xab, 0x3e, 0x70, 0xea, 0x87, 0x0d, 0x71, 0x8d, 0x66, 0xc2, 0xc9, 0x17, 0x70, 0xbc,
	0xc0, 0x94, 0x25, 0x25, 0xd5, 0x22, 0xcd, 0x98, 0x29, 0x14, 0x06, 0xed, 0xd0, 0x8b, 0xfc, 0xf8,
	0x61, 0x85, 0xbf, 0x59, 0xc3, 0xe4, 0x63, 0x68, 0x33, 0x29, 0x6d, 0xad, 0x4e, 0xe8, 0x45, 0x47,
	0xf1, 0x01, 0x93, 0x72, 0xa7, 0x9b, 0xcc, 0x9d, 0xa2, 0xeb, 0x2e, 0xb3, 0xd5, 0xed, 0x26, 0xb7,
	0xda, 0x13, 0xf0, 0x57, 0x42, 0x99, 0x82, 0x2d, 0xa8, 0xcc, 0x95, 0x09, 0xc0, 0x15, 0xea, 0xd5,
	0xd8, 0x4d, 0xae, 0x0c, 0xf9, 0x19, 0x7c, 0xbc, 0x33, 0x8a, 0xd1, 0xa9, 0xc0, 0x05, 0xd7, 0xc1,
	0x61, 0xd8, 0x8a, 0x7a, 0xe7, 0x5f, 0x0d, 0xfe, 0xcf, 0xa8, 0xc1, 0xd8, 0x9e, 0x7a, 0x6d, 0x0f,
	0xc5, 0x3d, 0xdc, 0x7c, 0xeb, 0x27, 0x7f, 0x79, 0x00, 0x0d, 0x47, 0x08, 0xec, 0x67, 0x6c, 0x89,
	0xce, 0xc8, 0x6e, 0xec, 0xbe, 0xed, 0x58, 0xda, 0x28, 0x91, 0xa5, 0x74, 0xc5, 0x16, 0x05, 0x3a,
	0xef, 0xba, 0x71, 0xaf, 0xc2, 0xde, 0x59, 0xc8, 0xc6, 0x40, 0x64, 0xe6, 0xdb, 0x6f, 0x6a, 0x85,
	0x35, 0x8b, 0xc4, 0xe0, 0xa0, 0x4a, 0xf0, 0x0c, 0x8e, 0xa6, 0xe2, 0x0e, 0xf9, 0x46, 0x52, 0x3d,
	0xb8, 0x5f, 0x83, 0x4e, 0xd4, 0xff, 0xdb, 0x83, 0xe7, 0x3b, 0x17, 0xb1, 0x2f, 0x8c, 0xfc, 0x7e,
	0xee, 0x4e, 0xe1, 0x91, 0xaa, 0xde, 0x8e, 0x53, 0x9e, 0xd3, 0x2c, 0x37, 0xb4, 0xd0, 0xd5, 0xd8,
	0xed, 0xf8, 0x78, 0x4d, 0x8d, 0xf2, 0xeb, 0xdc, 0xfc, 0xa2, 0x9d, 0x39, 0x73, 0x2c, 0x9b, 0xe0,
	0x1d, 0xcc, 0xb1, 0x9c, 0x70, 0xf2, 0x09, 0xb4, 0x8d, 0xab, 0xe7, 0x26, 0xf6, 0xe3, 0xfa, 0x17,
	0x79, 0x0a, 0xdd, 0xc6, 0xef, 0x7d, 0x47, 0x35, 0x40, 0xff, 0x9f, 0x3d, 0x78, 0xba, 0x33, 0xe6,
	0x10, 0x95, 0x11, 0x53, 0x91, 0x30, 0x83, 0xe4, 0x07, 0x38, 0xb4, 0xdd, 0x4c, 0x29, 0xab, 0x89,
	0x3e, 0x38, 0x7f, 0x3e, 0x78, 0xdf, 0x81, 0xc1, 0xf8, 0x0a, 0xcb, 0xb7, 0xa5, 0xc4, 0x57, 0x9d,
	0xc9, 0xf5, 0xbb, 0xcb, 0x1f, 0x27, 0xa3, 0xb8, 0x33, 0xaf, 0x10, 0xf2, 0xb8, 0xaa, 0x65, 0x37,
	0xd5, 0xcd, 0xee, 0x3b, 0xca, 0x16, 0xb1, 0xd4, 0xbd, 0xb4, 0x77, 0x74, 0x9d, 0xf1, 0x57, 0xf0,
	0x78, 0x2b, 0x75, 0xf6, 0x70, 0x82, 0x99, 0x41, 0x45, 0x05, 0xd7, 0xc1, 0x41, 0xd8, 0x8a, 0x3a,
	0xf1, 0xa7, 0x8d, 0x60, 0xb4, 0xe1, 0x27, 0x5c, 0x5b, 0xbb, 0xdd, 0x4a, 0x27, 0x0a, 0x99, 0x41,
	0x1e, 0x1c, 0xba, 0xb0, 0xba, 0x35, 0x1f, 0x56, 0xd0, 0xfd, 0xad, 0xef, 0xee, 0x6c, 0x7d, 0xb3,
	0x0c, 0xb0, 0xb5, 0x0c, 0xfd, 0xcf, 0xe1, 0x70, 0x7d, 0x55, 0xd2, 0x83, 0xf5, 0x65, 0x8f, 0x1f,
	0xd8, 0x1f, 0xe3, 0xd1, 0xf9, 0xc5, 0xc5, 0xcb, 0xef, 0x8e, 0xbd, 0xfe, 0x6f, 0xd0, 0x7f, 0xdf,
	0x6b, 0x55, 0x81, 0xb0, 0x49, 0x4d, 0x50, 0x99, 0xda, 0x1e, 0xf7, 0x4d, 0x9e, 0x40, 0x37, 0x61,
	0xb4, 0x76, 0xba, 0x4a, 0x58, 0x27, 0x61, 0x57, 0xce, 0xeb, 0x13, 0xf0, 0x13, 0xb6, 0xb3, 0xc6,
	0xbd, 0x84, 0x6d, 0x56, 0xb8, 0xff, 0xa7, 0x07, 0xd1, 0x6e, 0x67, 0x96, 0xcc, 0x90, 0x0f, 0x15,
	0x72, 0xcc, 0x8c, 0x60, 0x0b, 0xfd, 0x3a, 0x57, 0x97, 0x52, 0xda, 0x37, 0x90, 0x4a, 0xac, 0x98,
	0x41, 0xdb, 0xd0, 0xf9, 0xec, 0xc7, 0x50, 0x43, 0x57, 0x58, 0x6e, 0x06, 0xdc, 0xdb, 0x1a, 0xf0,
	0x19, 0x1c, 0x29, 0x9b, 0x64, 0x5a, 0x05, 0x4d, 0x07, 0xad, 0xb0, 0x15, 0xf9, 0xb1, 0xef, 0xc0,
	0x2a, 0xda, 0xfa, 0xfb, 0x8b, 0x5f, 0x5f, 0xa4, 0xc2, 0xcc, 0x8a, 0xdb, 0x41, 0x92, 0x2f, 0xcf,
	0x5e, 0x7e, 0x3d, 0x3f, 0x4b, 0xf3, 0x53, 0xe7, 0xec, 0xa9, 0x42, 0x9d, 0x17, 0x2a, 0x41, 0x7d,
	0xe6, 0xfe, 0xa3, 0x6f, 0x8b, 0xe9, 0x99, 0x23, 0xfe, 0xf0, 0x1e, 0xfc, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x86, 0xbe, 0x25, 0x4f, 0xd3, 0x05, 0x00, 0x00,
}
