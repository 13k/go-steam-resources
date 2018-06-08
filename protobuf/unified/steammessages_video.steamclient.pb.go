// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_video.steamclient.proto

package unified

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

type CVideo_ClientGetVideoURL_Request struct {
	VideoId              *uint64  `protobuf:"varint,1,opt,name=video_id,json=videoId" json:"video_id,omitempty"`
	ClientCellid         *uint32  `protobuf:"varint,2,opt,name=client_cellid,json=clientCellid" json:"client_cellid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CVideo_ClientGetVideoURL_Request) Reset()         { *m = CVideo_ClientGetVideoURL_Request{} }
func (m *CVideo_ClientGetVideoURL_Request) String() string { return proto.CompactTextString(m) }
func (*CVideo_ClientGetVideoURL_Request) ProtoMessage()    {}
func (*CVideo_ClientGetVideoURL_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e, []int{0}
}
func (m *CVideo_ClientGetVideoURL_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Request.Unmarshal(m, b)
}
func (m *CVideo_ClientGetVideoURL_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Request.Marshal(b, m, deterministic)
}
func (dst *CVideo_ClientGetVideoURL_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVideo_ClientGetVideoURL_Request.Merge(dst, src)
}
func (m *CVideo_ClientGetVideoURL_Request) XXX_Size() int {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Request.Size(m)
}
func (m *CVideo_ClientGetVideoURL_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CVideo_ClientGetVideoURL_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CVideo_ClientGetVideoURL_Request proto.InternalMessageInfo

func (m *CVideo_ClientGetVideoURL_Request) GetVideoId() uint64 {
	if m != nil && m.VideoId != nil {
		return *m.VideoId
	}
	return 0
}

func (m *CVideo_ClientGetVideoURL_Request) GetClientCellid() uint32 {
	if m != nil && m.ClientCellid != nil {
		return *m.ClientCellid
	}
	return 0
}

type CVideo_ClientGetVideoURL_Response struct {
	VideoId              *uint64  `protobuf:"varint,1,opt,name=video_id,json=videoId" json:"video_id,omitempty"`
	VideoUrl             *string  `protobuf:"bytes,2,opt,name=video_url,json=videoUrl" json:"video_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CVideo_ClientGetVideoURL_Response) Reset()         { *m = CVideo_ClientGetVideoURL_Response{} }
func (m *CVideo_ClientGetVideoURL_Response) String() string { return proto.CompactTextString(m) }
func (*CVideo_ClientGetVideoURL_Response) ProtoMessage()    {}
func (*CVideo_ClientGetVideoURL_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e, []int{1}
}
func (m *CVideo_ClientGetVideoURL_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Response.Unmarshal(m, b)
}
func (m *CVideo_ClientGetVideoURL_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Response.Marshal(b, m, deterministic)
}
func (dst *CVideo_ClientGetVideoURL_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVideo_ClientGetVideoURL_Response.Merge(dst, src)
}
func (m *CVideo_ClientGetVideoURL_Response) XXX_Size() int {
	return xxx_messageInfo_CVideo_ClientGetVideoURL_Response.Size(m)
}
func (m *CVideo_ClientGetVideoURL_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CVideo_ClientGetVideoURL_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CVideo_ClientGetVideoURL_Response proto.InternalMessageInfo

func (m *CVideo_ClientGetVideoURL_Response) GetVideoId() uint64 {
	if m != nil && m.VideoId != nil {
		return *m.VideoId
	}
	return 0
}

func (m *CVideo_ClientGetVideoURL_Response) GetVideoUrl() string {
	if m != nil && m.VideoUrl != nil {
		return *m.VideoUrl
	}
	return ""
}

type CVideo_UnlockedH264_Notification struct {
	EncryptionKey        []byte   `protobuf:"bytes,1,opt,name=encryption_key,json=encryptionKey" json:"encryption_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CVideo_UnlockedH264_Notification) Reset()         { *m = CVideo_UnlockedH264_Notification{} }
func (m *CVideo_UnlockedH264_Notification) String() string { return proto.CompactTextString(m) }
func (*CVideo_UnlockedH264_Notification) ProtoMessage()    {}
func (*CVideo_UnlockedH264_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e, []int{2}
}
func (m *CVideo_UnlockedH264_Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVideo_UnlockedH264_Notification.Unmarshal(m, b)
}
func (m *CVideo_UnlockedH264_Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVideo_UnlockedH264_Notification.Marshal(b, m, deterministic)
}
func (dst *CVideo_UnlockedH264_Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVideo_UnlockedH264_Notification.Merge(dst, src)
}
func (m *CVideo_UnlockedH264_Notification) XXX_Size() int {
	return xxx_messageInfo_CVideo_UnlockedH264_Notification.Size(m)
}
func (m *CVideo_UnlockedH264_Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_CVideo_UnlockedH264_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_CVideo_UnlockedH264_Notification proto.InternalMessageInfo

func (m *CVideo_UnlockedH264_Notification) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

type CFovasVideo_ClientGetOPFSettings_Request struct {
	AppId                *uint32  `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	ClientCellid         *uint32  `protobuf:"varint,2,opt,name=client_cellid,json=clientCellid" json:"client_cellid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CFovasVideo_ClientGetOPFSettings_Request) Reset() {
	*m = CFovasVideo_ClientGetOPFSettings_Request{}
}
func (m *CFovasVideo_ClientGetOPFSettings_Request) String() string { return proto.CompactTextString(m) }
func (*CFovasVideo_ClientGetOPFSettings_Request) ProtoMessage()    {}
func (*CFovasVideo_ClientGetOPFSettings_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e, []int{3}
}
func (m *CFovasVideo_ClientGetOPFSettings_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request.Unmarshal(m, b)
}
func (m *CFovasVideo_ClientGetOPFSettings_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request.Marshal(b, m, deterministic)
}
func (dst *CFovasVideo_ClientGetOPFSettings_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request.Merge(dst, src)
}
func (m *CFovasVideo_ClientGetOPFSettings_Request) XXX_Size() int {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request.Size(m)
}
func (m *CFovasVideo_ClientGetOPFSettings_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Request proto.InternalMessageInfo

func (m *CFovasVideo_ClientGetOPFSettings_Request) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CFovasVideo_ClientGetOPFSettings_Request) GetClientCellid() uint32 {
	if m != nil && m.ClientCellid != nil {
		return *m.ClientCellid
	}
	return 0
}

type CFovasVideo_ClientGetOPFSettings_Response struct {
	AppId                *uint32  `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	OpfSettings          *string  `protobuf:"bytes,2,opt,name=opf_settings,json=opfSettings" json:"opf_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CFovasVideo_ClientGetOPFSettings_Response) Reset() {
	*m = CFovasVideo_ClientGetOPFSettings_Response{}
}
func (m *CFovasVideo_ClientGetOPFSettings_Response) String() string { return proto.CompactTextString(m) }
func (*CFovasVideo_ClientGetOPFSettings_Response) ProtoMessage()    {}
func (*CFovasVideo_ClientGetOPFSettings_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e, []int{4}
}
func (m *CFovasVideo_ClientGetOPFSettings_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response.Unmarshal(m, b)
}
func (m *CFovasVideo_ClientGetOPFSettings_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response.Marshal(b, m, deterministic)
}
func (dst *CFovasVideo_ClientGetOPFSettings_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response.Merge(dst, src)
}
func (m *CFovasVideo_ClientGetOPFSettings_Response) XXX_Size() int {
	return xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response.Size(m)
}
func (m *CFovasVideo_ClientGetOPFSettings_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CFovasVideo_ClientGetOPFSettings_Response proto.InternalMessageInfo

func (m *CFovasVideo_ClientGetOPFSettings_Response) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CFovasVideo_ClientGetOPFSettings_Response) GetOpfSettings() string {
	if m != nil && m.OpfSettings != nil {
		return *m.OpfSettings
	}
	return ""
}

func init() {
	proto.RegisterType((*CVideo_ClientGetVideoURL_Request)(nil), "CVideo_ClientGetVideoURL_Request")
	proto.RegisterType((*CVideo_ClientGetVideoURL_Response)(nil), "CVideo_ClientGetVideoURL_Response")
	proto.RegisterType((*CVideo_UnlockedH264_Notification)(nil), "CVideo_UnlockedH264_Notification")
	proto.RegisterType((*CFovasVideo_ClientGetOPFSettings_Request)(nil), "CFovasVideo_ClientGetOPFSettings_Request")
	proto.RegisterType((*CFovasVideo_ClientGetOPFSettings_Response)(nil), "CFovasVideo_ClientGetOPFSettings_Response")
}

func init() {
	proto.RegisterFile("steammessages_video.steamclient.proto", fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e)
}

var fileDescriptor_steammessages_video_steamclient_025c0dcb9779ee9e = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x4a,
	0x10, 0xc7, 0xb5, 0x7d, 0x6d, 0x5f, 0xbb, 0x4d, 0x9e, 0x54, 0xeb, 0x09, 0x19, 0x23, 0xc4, 0xd4,
	0x52, 0xa1, 0x45, 0x95, 0x85, 0x42, 0x55, 0x6e, 0x88, 0x26, 0x55, 0x4b, 0x68, 0x69, 0x23, 0x97,
	0x70, 0xb5, 0x36, 0xc9, 0x38, 0x59, 0xd5, 0xd9, 0x35, 0xde, 0x4d, 0x50, 0x38, 0x21, 0x4b, 0x48,
	0x9c, 0x10, 0x47, 0xc4, 0x81, 0x2f, 0xc0, 0xd9, 0x77, 0x24, 0x3e, 0x12, 0x5f, 0x00, 0x79, 0x9d,
	0xa4, 0x85, 0x02, 0xcd, 0x81, 0xe3, 0xce, 0x8c, 0x67, 0xfe, 0x33, 0xff, 0x9f, 0xe9, 0xba, 0xd2,
	0xc8, 0xfa, 0x7d, 0x54, 0x8a, 0x75, 0x51, 0x05, 0x43, 0xde, 0x41, 0xe9, 0x99, 0x58, 0x3b, 0xe2,
	0x28, 0xb4, 0x17, 0x27, 0x52, 0x4b, 0x67, 0xeb, 0xc7, 0xb2, 0x81, 0xe0, 0x21, 0xc7, 0x4e, 0xd0,
	0x62, 0x0a, 0x2f, 0x57, 0xbb, 0x1f, 0x08, 0x85, 0xda, 0xf3, 0xbc, 0x55, 0x50, 0x33, 0xf1, 0x03,
	0xd4, 0xe6, 0xd9, 0xf4, 0x8f, 0x02, 0x1f, 0x5f, 0x0c, 0x50, 0x69, 0xeb, 0x0e, 0x5d, 0x32, 0xd3,
	0x02, 0xde, 0xb1, 0x09, 0x90, 0x8d, 0xf9, 0x6a, 0x29, 0xcd, 0xec, 0x25, 0x53, 0x07, 0xf5, 0x3d,
	0xff, 0x5f, 0x93, 0xad, 0x77, 0xac, 0x43, 0x5a, 0x2e, 0xba, 0x07, 0x6d, 0x8c, 0x22, 0xde, 0xb1,
	0xe7, 0x80, 0x6c, 0x94, 0xab, 0xb7, 0xd3, 0xcc, 0x76, 0x6b, 0x18, 0x45, 0x50, 0xdf, 0x03, 0x19,
	0x42, 0x51, 0xb3, 0x05, 0xaf, 0x30, 0x91, 0xc0, 0x43, 0x18, 0x88, 0x33, 0x21, 0x5f, 0x0a, 0xbf,
	0x54, 0x24, 0x6a, 0xe6, 0x5b, 0xf7, 0x0d, 0xa1, 0x6b, 0x7f, 0x90, 0xa6, 0x62, 0x29, 0x14, 0xce,
	0xae, 0xed, 0x01, 0x5d, 0x2e, 0x0a, 0x07, 0x49, 0x64, 0x74, 0x2d, 0x57, 0x9d, 0x34, 0xb3, 0xaf,
	0x35, 0xfd, 0x23, 0x08, 0x65, 0x02, 0x26, 0x09, 0x7d, 0x26, 0x78, 0x88, 0x4a, 0xfb, 0x45, 0xd7,
	0x66, 0x12, 0xb9, 0xf5, 0xe9, 0x85, 0x9a, 0x22, 0x92, 0xed, 0x33, 0xec, 0x3c, 0xae, 0xec, 0x6c,
	0x07, 0xc7, 0x52, 0xf3, 0x90, 0xb7, 0x99, 0xe6, 0x52, 0x58, 0xeb, 0xf4, 0x3f, 0x14, 0xed, 0x64,
	0x14, 0xe7, 0xaf, 0xe0, 0x0c, 0x47, 0x46, 0x4b, 0xc9, 0x2f, 0x9f, 0x47, 0x0f, 0x71, 0xe4, 0x7e,
	0x24, 0x74, 0xa3, 0xb6, 0x2f, 0x87, 0x4c, 0xfd, 0xb4, 0xd7, 0x49, 0x63, 0xff, 0x14, 0xb5, 0xe6,
	0xa2, 0xab, 0xa6, 0x57, 0x5f, 0xa3, 0x8b, 0x2c, 0x8e, 0x27, 0x7b, 0x95, 0xab, 0x34, 0xcd, 0xec,
	0xc5, 0xdd, 0x38, 0xce, 0xb7, 0x5a, 0x60, 0x71, 0xfc, 0xb7, 0xef, 0xfd, 0x9e, 0xd0, 0xcd, 0x19,
	0xc4, 0x8d, 0xef, 0x3e, 0x83, 0xba, 0x47, 0xb4, 0x24, 0xe3, 0x30, 0x50, 0xe3, 0x6f, 0xc7, 0x47,
	0xbf, 0x99, 0x66, 0xf6, 0xf5, 0x27, 0xa7, 0x27, 0xc7, 0xd0, 0x8a, 0x64, 0x2b, 0x97, 0x77, 0xd2,
	0xd8, 0x87, 0xc9, 0x00, 0x7f, 0x45, 0xc6, 0xe1, 0xe4, 0x51, 0xf9, 0x4c, 0xe8, 0x82, 0x11, 0x63,
	0xbd, 0x23, 0x74, 0xf5, 0x12, 0x05, 0xd6, 0x9a, 0x77, 0x15, 0xbb, 0x8e, 0xeb, 0x5d, 0xc9, 0x90,
	0x5b, 0x49, 0x33, 0xdb, 0x3b, 0x40, 0x0d, 0xba, 0x87, 0xc0, 0x05, 0xd7, 0x9c, 0x45, 0x90, 0x73,
	0xa1, 0x25, 0xb4, 0xb0, 0xcb, 0x05, 0x28, 0x9d, 0x20, 0xeb, 0x73, 0xd1, 0x05, 0x56, 0x80, 0xe2,
	0xac, 0xa6, 0x99, 0x5d, 0x2e, 0x28, 0xeb, 0xa3, 0xee, 0xc9, 0x8e, 0xaa, 0x7c, 0x22, 0x74, 0xc5,
	0x44, 0x8a, 0x51, 0xd6, 0x5b, 0x42, 0x2d, 0x43, 0xc9, 0xe8, 0x22, 0x38, 0xe7, 0xa2, 0x7f, 0x8b,
	0x93, 0xb3, 0xe2, 0x1d, 0xcb, 0xa9, 0xba, 0x6a, 0x9a, 0xd9, 0x0f, 0x2f, 0xa6, 0x21, 0x4c, 0x64,
	0x1f, 0x14, 0x26, 0x43, 0x4c, 0x72, 0x89, 0x85, 0x8b, 0xa0, 0x7b, 0x4c, 0x43, 0xaf, 0xb2, 0xb3,
	0x0d, 0x3d, 0xa6, 0xa0, 0x85, 0x28, 0x60, 0x30, 0x9e, 0xe0, 0xcc, 0x7f, 0xc9, 0xec, 0xb9, 0xca,
	0x37, 0x42, 0xe9, 0xb9, 0xc1, 0xd6, 0x57, 0x42, 0xff, 0xff, 0x95, 0xc9, 0xd6, 0xa6, 0x37, 0x2b,
	0xa4, 0xce, 0x5d, 0x6f, 0x66, 0x64, 0xdc, 0x67, 0x69, 0x66, 0x37, 0x7c, 0xd4, 0x09, 0xc7, 0x21,
	0x9a, 0x5b, 0xe7, 0xce, 0x4f, 0xf0, 0x80, 0x29, 0x13, 0x1e, 0xec, 0x0e, 0x19, 0x8f, 0x58, 0x2b,
	0x42, 0x18, 0x72, 0x66, 0x2a, 0x8b, 0xce, 0xe6, 0x47, 0xbd, 0xbf, 0x73, 0x0f, 0x1a, 0x11, 0x1b,
	0x61, 0x02, 0xbb, 0x71, 0xec, 0xdc, 0x4a, 0x33, 0xfb, 0x86, 0xd1, 0x00, 0x85, 0x1d, 0xa7, 0x98,
	0x0c, 0x79, 0x1b, 0xe1, 0x69, 0x61, 0x4b, 0xf5, 0x9f, 0xd7, 0x84, 0x7c, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x5e, 0xa4, 0xbd, 0xd4, 0x3d, 0x05, 0x00, 0x00,
}
