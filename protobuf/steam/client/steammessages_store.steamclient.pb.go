// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_store.steamclient.proto

package client // import "github.com/13k/go-steam-resources/protobuf/steam/client"

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

type CStore_GetLocalizedNameForTags_Request struct {
	Language             *string  `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Tagids               []uint32 `protobuf:"varint,2,rep,name=tagids" json:"tagids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CStore_GetLocalizedNameForTags_Request) Reset() {
	*m = CStore_GetLocalizedNameForTags_Request{}
}
func (m *CStore_GetLocalizedNameForTags_Request) String() string { return proto.CompactTextString(m) }
func (*CStore_GetLocalizedNameForTags_Request) ProtoMessage()    {}
func (*CStore_GetLocalizedNameForTags_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_store_steamclient_b2c3a820870fa754, []int{0}
}
func (m *CStore_GetLocalizedNameForTags_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Request.Unmarshal(m, b)
}
func (m *CStore_GetLocalizedNameForTags_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Request.Marshal(b, m, deterministic)
}
func (dst *CStore_GetLocalizedNameForTags_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Request.Merge(dst, src)
}
func (m *CStore_GetLocalizedNameForTags_Request) XXX_Size() int {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Request.Size(m)
}
func (m *CStore_GetLocalizedNameForTags_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CStore_GetLocalizedNameForTags_Request proto.InternalMessageInfo

func (m *CStore_GetLocalizedNameForTags_Request) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *CStore_GetLocalizedNameForTags_Request) GetTagids() []uint32 {
	if m != nil {
		return m.Tagids
	}
	return nil
}

type CStore_GetLocalizedNameForTags_Response struct {
	Tags                 []*CStore_GetLocalizedNameForTags_Response_Tag `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *CStore_GetLocalizedNameForTags_Response) Reset() {
	*m = CStore_GetLocalizedNameForTags_Response{}
}
func (m *CStore_GetLocalizedNameForTags_Response) String() string { return proto.CompactTextString(m) }
func (*CStore_GetLocalizedNameForTags_Response) ProtoMessage()    {}
func (*CStore_GetLocalizedNameForTags_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_store_steamclient_b2c3a820870fa754, []int{1}
}
func (m *CStore_GetLocalizedNameForTags_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response.Unmarshal(m, b)
}
func (m *CStore_GetLocalizedNameForTags_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response.Marshal(b, m, deterministic)
}
func (dst *CStore_GetLocalizedNameForTags_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Response.Merge(dst, src)
}
func (m *CStore_GetLocalizedNameForTags_Response) XXX_Size() int {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response.Size(m)
}
func (m *CStore_GetLocalizedNameForTags_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CStore_GetLocalizedNameForTags_Response proto.InternalMessageInfo

func (m *CStore_GetLocalizedNameForTags_Response) GetTags() []*CStore_GetLocalizedNameForTags_Response_Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CStore_GetLocalizedNameForTags_Response_Tag struct {
	Tagid                *uint32  `protobuf:"varint,1,opt,name=tagid" json:"tagid,omitempty"`
	EnglishName          *string  `protobuf:"bytes,2,opt,name=english_name,json=englishName" json:"english_name,omitempty"`
	Name                 *string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CStore_GetLocalizedNameForTags_Response_Tag) Reset() {
	*m = CStore_GetLocalizedNameForTags_Response_Tag{}
}
func (m *CStore_GetLocalizedNameForTags_Response_Tag) String() string {
	return proto.CompactTextString(m)
}
func (*CStore_GetLocalizedNameForTags_Response_Tag) ProtoMessage() {}
func (*CStore_GetLocalizedNameForTags_Response_Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_store_steamclient_b2c3a820870fa754, []int{1, 0}
}
func (m *CStore_GetLocalizedNameForTags_Response_Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag.Unmarshal(m, b)
}
func (m *CStore_GetLocalizedNameForTags_Response_Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag.Marshal(b, m, deterministic)
}
func (dst *CStore_GetLocalizedNameForTags_Response_Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag.Merge(dst, src)
}
func (m *CStore_GetLocalizedNameForTags_Response_Tag) XXX_Size() int {
	return xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag.Size(m)
}
func (m *CStore_GetLocalizedNameForTags_Response_Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_CStore_GetLocalizedNameForTags_Response_Tag proto.InternalMessageInfo

func (m *CStore_GetLocalizedNameForTags_Response_Tag) GetTagid() uint32 {
	if m != nil && m.Tagid != nil {
		return *m.Tagid
	}
	return 0
}

func (m *CStore_GetLocalizedNameForTags_Response_Tag) GetEnglishName() string {
	if m != nil && m.EnglishName != nil {
		return *m.EnglishName
	}
	return ""
}

func (m *CStore_GetLocalizedNameForTags_Response_Tag) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CStore_GetLocalizedNameForTags_Request)(nil), "CStore_GetLocalizedNameForTags_Request")
	proto.RegisterType((*CStore_GetLocalizedNameForTags_Response)(nil), "CStore_GetLocalizedNameForTags_Response")
	proto.RegisterType((*CStore_GetLocalizedNameForTags_Response_Tag)(nil), "CStore_GetLocalizedNameForTags_Response.Tag")
}

func init() {
	proto.RegisterFile("steammessages_store.steamclient.proto", fileDescriptor_steammessages_store_steamclient_b2c3a820870fa754)
}

var fileDescriptor_steammessages_store_steamclient_b2c3a820870fa754 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xd5, 0xc6, 0x09, 0x82, 0x0d, 0x69, 0x56, 0x08, 0x2c, 0x37, 0x2c, 0x41, 0x24, 0x16, 0x4a,
	0x6c, 0x11, 0x0a, 0x1a, 0x0a, 0x3e, 0x24, 0xd2, 0x20, 0x0a, 0x93, 0x0a, 0x21, 0x59, 0x13, 0x7b,
	0xbc, 0x59, 0x61, 0xef, 0x06, 0xcf, 0x9a, 0x82, 0x0a, 0xe5, 0x57, 0xf0, 0x47, 0x52, 0x53, 0xf2,
	0xb7, 0x4e, 0xd9, 0xe4, 0x4e, 0x3a, 0xe9, 0x4e, 0x97, 0x6e, 0xdf, 0xbc, 0x99, 0xb7, 0x6f, 0x46,
	0x8f, 0xbf, 0x20, 0x87, 0xd0, 0x34, 0x48, 0x04, 0x0a, 0x29, 0x27, 0x67, 0x5b, 0x4c, 0x7c, 0xad,
	0xa8, 0x35, 0x1a, 0x97, 0x6c, 0x5b, 0xeb, 0x6c, 0x34, 0xbb, 0xde, 0xd6, 0x19, 0x5d, 0x69, 0x2c,
	0xf3, 0x35, 0xd0, 0x0d, 0xdd, 0xe3, 0xef, 0x7c, 0xf2, 0xf1, 0xeb, 0x41, 0x29, 0x5f, 0xa2, 0xfb,
	0x6c, 0x0b, 0xa8, 0xf5, 0x6f, 0x2c, 0xbf, 0x40, 0x83, 0x9f, 0x6c, 0xbb, 0x02, 0x45, 0x79, 0x86,
	0x3f, 0x3b, 0x24, 0x27, 0x22, 0x7e, 0xbf, 0x06, 0xa3, 0x3a, 0x50, 0x18, 0x32, 0xc9, 0xe2, 0x07,
	0xd9, 0x15, 0x16, 0x8f, 0xf9, 0x3d, 0x07, 0x4a, 0x97, 0x14, 0xf6, 0x64, 0x10, 0x8f, 0xb2, 0x13,
	0x1a, 0xff, 0x63, 0x7c, 0x7a, 0xa7, 0x3c, 0x6d, 0xad, 0x21, 0x14, 0xef, 0x78, 0xdf, 0x81, 0xa2,
	0x90, 0xc9, 0x20, 0x1e, 0x2e, 0x66, 0xc9, 0x99, 0x73, 0xc9, 0x0a, 0x54, 0xe6, 0x27, 0xa3, 0x8c,
	0x07, 0x2b, 0x50, 0xe2, 0x11, 0x1f, 0xf8, 0xef, 0xbd, 0xcb, 0x51, 0x76, 0x04, 0xe2, 0x19, 0x7f,
	0x88, 0x46, 0xd5, 0x9a, 0x36, 0xb9, 0x81, 0x06, 0xc3, 0x9e, 0x5f, 0x61, 0x78, 0xaa, 0x1d, 0x94,
	0x85, 0xe0, 0x7d, 0x4f, 0x05, 0x9e, 0xf2, 0xef, 0xc5, 0x7f, 0xc6, 0x07, 0xde, 0x88, 0xf8, 0xcb,
	0xf8, 0x93, 0x5b, 0xcc, 0x88, 0x69, 0x72, 0xde, 0x11, 0xa3, 0xf8, 0xdc, 0xb5, 0xc6, 0x2f, 0x77,
	0xfb, 0x70, 0xb2, 0x44, 0x47, 0xd2, 0x81, 0x92, 0x07, 0x2f, 0x24, 0xb5, 0x91, 0x20, 0x4b, 0x5d,
	0x55, 0xd8, 0xa2, 0x71, 0xf2, 0xf2, 0xfc, 0xd1, 0xf3, 0xdd, 0x3e, 0x7c, 0xfa, 0x5e, 0x12, 0xb6,
	0xbf, 0x74, 0x81, 0xd2, 0x59, 0x09, 0x45, 0x81, 0x44, 0xd2, 0x67, 0x44, 0x96, 0xe0, 0x20, 0xf9,
	0xf0, 0xf6, 0xdb, 0x1b, 0xa5, 0xdd, 0xa6, 0x5b, 0x27, 0x85, 0x6d, 0xd2, 0x57, 0xaf, 0x7f, 0xa4,
	0xca, 0xce, 0x7d, 0x20, 0xe6, 0x2d, 0x92, 0xed, 0xda, 0x02, 0x29, 0xf5, 0xa1, 0x58, 0x77, 0x55,
	0xea, 0x89, 0xf4, 0x18, 0x95, 0x3f, 0x8c, 0x5d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x94, 0xfe, 0x73,
	0xc0, 0x7d, 0x02, 0x00, 0x00,
}