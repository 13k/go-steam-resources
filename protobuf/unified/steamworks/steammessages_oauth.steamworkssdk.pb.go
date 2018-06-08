// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_oauth.steamworkssdk.proto

package steamworks

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

type COAuthToken_ImplicitGrantNoPrompt_Request struct {
	Clientid             *string  `protobuf:"bytes,1,opt,name=clientid" json:"clientid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Request) Reset() {
	*m = COAuthToken_ImplicitGrantNoPrompt_Request{}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) String() string { return proto.CompactTextString(m) }
func (*COAuthToken_ImplicitGrantNoPrompt_Request) ProtoMessage()    {}
func (*COAuthToken_ImplicitGrantNoPrompt_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_oauth_steamworkssdk_4da66221a162949b, []int{0}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request.Unmarshal(m, b)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request.Marshal(b, m, deterministic)
}
func (dst *COAuthToken_ImplicitGrantNoPrompt_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request.Merge(dst, src)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) XXX_Size() int {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request.Size(m)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request.DiscardUnknown(m)
}

var xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Request proto.InternalMessageInfo

func (m *COAuthToken_ImplicitGrantNoPrompt_Request) GetClientid() string {
	if m != nil && m.Clientid != nil {
		return *m.Clientid
	}
	return ""
}

type COAuthToken_ImplicitGrantNoPrompt_Response struct {
	AccessToken          *string  `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RedirectUri          *string  `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) Reset() {
	*m = COAuthToken_ImplicitGrantNoPrompt_Response{}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) String() string {
	return proto.CompactTextString(m)
}
func (*COAuthToken_ImplicitGrantNoPrompt_Response) ProtoMessage() {}
func (*COAuthToken_ImplicitGrantNoPrompt_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_oauth_steamworkssdk_4da66221a162949b, []int{1}
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response.Unmarshal(m, b)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response.Marshal(b, m, deterministic)
}
func (dst *COAuthToken_ImplicitGrantNoPrompt_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response.Merge(dst, src)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) XXX_Size() int {
	return xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response.Size(m)
}
func (m *COAuthToken_ImplicitGrantNoPrompt_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response.DiscardUnknown(m)
}

var xxx_messageInfo_COAuthToken_ImplicitGrantNoPrompt_Response proto.InternalMessageInfo

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *COAuthToken_ImplicitGrantNoPrompt_Response) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func init() {
	proto.RegisterType((*COAuthToken_ImplicitGrantNoPrompt_Request)(nil), "COAuthToken_ImplicitGrantNoPrompt_Request")
	proto.RegisterType((*COAuthToken_ImplicitGrantNoPrompt_Response)(nil), "COAuthToken_ImplicitGrantNoPrompt_Response")
}

func init() {
	proto.RegisterFile("steammessages_oauth.steamworkssdk.proto", fileDescriptor_steammessages_oauth_steamworkssdk_4da66221a162949b)
}

var fileDescriptor_steammessages_oauth_steamworkssdk_4da66221a162949b = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6a, 0x14, 0x41,
	0x10, 0x65, 0x72, 0xd2, 0x49, 0x4e, 0x03, 0xc2, 0xb0, 0x17, 0x8b, 0xf5, 0xb0, 0x49, 0x94, 0x01,
	0x45, 0x21, 0x07, 0x2f, 0x9a, 0xa0, 0xec, 0x45, 0x65, 0x34, 0x20, 0x78, 0x18, 0x7a, 0xbb, 0x6b,
	0x67, 0x8a, 0xdd, 0xe9, 0x1e, 0xbb, 0xaa, 0xb3, 0xe4, 0x3a, 0xe0, 0x07, 0x79, 0x98, 0x0f, 0xf1,
	0x3b, 0xfc, 0x09, 0x99, 0xde, 0x8d, 0x1a, 0xd8, 0xc3, 0x5e, 0x5f, 0xbf, 0x7a, 0xfd, 0xea, 0xbd,
	0x4a, 0x67, 0x2c, 0xa8, 0xda, 0x16, 0x99, 0x55, 0x8d, 0x5c, 0x39, 0x15, 0xa4, 0x29, 0x22, 0xb6,
	0x71, 0x7e, 0xc5, 0x6c, 0x56, 0x45, 0xe7, 0x9d, 0xb8, 0x49, 0x71, 0x9f, 0x18, 0x2c, 0x2d, 0x09,
	0x4d, 0xb5, 0x50, 0x8c, 0xfb, 0xf8, 0xd3, 0x1f, 0x49, 0x7a, 0x76, 0xf9, 0xf1, 0x4d, 0x90, 0xe6,
	0x8b, 0x5b, 0xa1, 0xad, 0xe6, 0x6d, 0xb7, 0x26, 0x4d, 0xf2, 0xde, 0x2b, 0x2b, 0x1f, 0xdc, 0x27,
	0xef, 0xda, 0x4e, 0xaa, 0x12, 0xbf, 0x07, 0x64, 0xc9, 0xbe, 0xa6, 0x0f, 0xf4, 0x9a, 0xd0, 0x0a,
	0x99, 0x3c, 0x81, 0xe4, 0xf4, 0xe1, 0xdb, 0xd7, 0xfd, 0x90, 0x5f, 0x5c, 0x46, 0x0c, 0xe6, 0x57,
	0xb0, 0x74, 0x1e, 0x36, 0x0d, 0xe9, 0x06, 0xc4, 0x81, 0x76, 0xc1, 0x0a, 0x48, 0x83, 0x60, 0x43,
	0xbb, 0x40, 0x0f, 0x6e, 0x09, 0xc4, 0x1c, 0xd0, 0x80, 0x8c, 0x9f, 0x71, 0xf9, 0x57, 0x6d, 0xfa,
	0x2b, 0x49, 0xcf, 0x0f, 0xf1, 0xc1, 0x9d, 0xb3, 0x8c, 0xd9, 0xbb, 0xf4, 0x44, 0x69, 0x8d, 0xcc,
	0x55, 0x54, 0xda, 0x99, 0x79, 0xd2, 0x0f, 0xf9, 0xe3, 0x28, 0x02, 0x51, 0xe5, 0x19, 0xd4, 0xe3,
	0x38, 0x1a, 0x70, 0x16, 0x38, 0xc4, 0x81, 0xf2, 0x78, 0x3b, 0x18, 0x09, 0xd9, 0xb7, 0xf4, 0xc4,
	0xa3, 0x21, 0x8f, 0x5a, 0xaa, 0xe0, 0x29, 0x3f, 0x8a, 0x3a, 0x17, 0xfd, 0x90, 0xbf, 0x2c, 0x77,
	0x38, 0x39, 0x0b, 0xd7, 0xe5, 0x1c, 0x3a, 0xef, 0x6e, 0xc8, 0xa0, 0x01, 0x13, 0x3c, 0xd9, 0x1a,
	0xb6, 0xd6, 0xc1, 0x63, 0x4d, 0x2c, 0x5e, 0x8d, 0xbc, 0xa2, 0x3c, 0xbe, 0x53, 0xbb, 0xf6, 0xf4,
	0xe2, 0xe7, 0x51, 0x9a, 0xfe, 0x5b, 0x29, 0xfb, 0x9d, 0xa4, 0x8f, 0xf6, 0xae, 0x95, 0x9d, 0x17,
	0x07, 0x57, 0x30, 0x79, 0x5a, 0x1c, 0x1e, 0xd3, 0xf4, 0xb6, 0x1f, 0xf2, 0x10, 0xdf, 0x18, 0x94,
	0x05, 0xda, 0xb1, 0x61, 0x9b, 0x51, 0x8c, 0x0e, 0x4e, 0x63, 0x46, 0x20, 0xb7, 0x1d, 0xc2, 0x2c,
	0x42, 0xb3, 0xb3, 0xd8, 0xe3, 0x58, 0x1c, 0x77, 0xa8, 0xe3, 0x05, 0xdd, 0xad, 0x3b, 0xbf, 0x1a,
	0xb3, 0x5c, 0x60, 0xa3, 0xd6, 0xcb, 0xb1, 0x4f, 0x05, 0x81, 0xd1, 0xc3, 0x86, 0xa4, 0x71, 0x41,
	0xc6, 0x90, 0xda, 0x4e, 0xc8, 0xd6, 0x93, 0x57, 0xfd, 0x90, 0x3f, 0xff, 0x8c, 0xfe, 0x86, 0x34,
	0x82, 0x76, 0x56, 0x14, 0xd9, 0x31, 0xb6, 0x16, 0xa5, 0x71, 0x86, 0xc7, 0xfb, 0x68, 0x95, 0x55,
	0x35, 0xfe, 0x6f, 0x86, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x45, 0x90, 0x8d, 0x70, 0xe9, 0x02,
	0x00, 0x00,
}
