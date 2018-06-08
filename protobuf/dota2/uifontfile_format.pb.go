// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uifontfile_format.proto

package dota2 // import "github.com/13k/go-steam-resources/protobuf/dota2"

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

type CUIFontFilePB struct {
	FontFileName         *string  `protobuf:"bytes,1,opt,name=font_file_name,json=fontFileName" json:"font_file_name,omitempty"`
	OpentypeFontData     []byte   `protobuf:"bytes,2,opt,name=opentype_font_data,json=opentypeFontData" json:"opentype_font_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CUIFontFilePB) Reset()         { *m = CUIFontFilePB{} }
func (m *CUIFontFilePB) String() string { return proto.CompactTextString(m) }
func (*CUIFontFilePB) ProtoMessage()    {}
func (*CUIFontFilePB) Descriptor() ([]byte, []int) {
	return fileDescriptor_uifontfile_format_03661485664839fa, []int{0}
}
func (m *CUIFontFilePB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CUIFontFilePB.Unmarshal(m, b)
}
func (m *CUIFontFilePB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CUIFontFilePB.Marshal(b, m, deterministic)
}
func (dst *CUIFontFilePB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CUIFontFilePB.Merge(dst, src)
}
func (m *CUIFontFilePB) XXX_Size() int {
	return xxx_messageInfo_CUIFontFilePB.Size(m)
}
func (m *CUIFontFilePB) XXX_DiscardUnknown() {
	xxx_messageInfo_CUIFontFilePB.DiscardUnknown(m)
}

var xxx_messageInfo_CUIFontFilePB proto.InternalMessageInfo

func (m *CUIFontFilePB) GetFontFileName() string {
	if m != nil && m.FontFileName != nil {
		return *m.FontFileName
	}
	return ""
}

func (m *CUIFontFilePB) GetOpentypeFontData() []byte {
	if m != nil {
		return m.OpentypeFontData
	}
	return nil
}

type CUIFontFilePackagePB struct {
	PackageVersion       *uint32                                        `protobuf:"varint,1,req,name=package_version,json=packageVersion" json:"package_version,omitempty"`
	EncryptedFontFiles   []*CUIFontFilePackagePB_CUIEncryptedFontFilePB `protobuf:"bytes,2,rep,name=encrypted_font_files,json=encryptedFontFiles" json:"encrypted_font_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *CUIFontFilePackagePB) Reset()         { *m = CUIFontFilePackagePB{} }
func (m *CUIFontFilePackagePB) String() string { return proto.CompactTextString(m) }
func (*CUIFontFilePackagePB) ProtoMessage()    {}
func (*CUIFontFilePackagePB) Descriptor() ([]byte, []int) {
	return fileDescriptor_uifontfile_format_03661485664839fa, []int{1}
}
func (m *CUIFontFilePackagePB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CUIFontFilePackagePB.Unmarshal(m, b)
}
func (m *CUIFontFilePackagePB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CUIFontFilePackagePB.Marshal(b, m, deterministic)
}
func (dst *CUIFontFilePackagePB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CUIFontFilePackagePB.Merge(dst, src)
}
func (m *CUIFontFilePackagePB) XXX_Size() int {
	return xxx_messageInfo_CUIFontFilePackagePB.Size(m)
}
func (m *CUIFontFilePackagePB) XXX_DiscardUnknown() {
	xxx_messageInfo_CUIFontFilePackagePB.DiscardUnknown(m)
}

var xxx_messageInfo_CUIFontFilePackagePB proto.InternalMessageInfo

func (m *CUIFontFilePackagePB) GetPackageVersion() uint32 {
	if m != nil && m.PackageVersion != nil {
		return *m.PackageVersion
	}
	return 0
}

func (m *CUIFontFilePackagePB) GetEncryptedFontFiles() []*CUIFontFilePackagePB_CUIEncryptedFontFilePB {
	if m != nil {
		return m.EncryptedFontFiles
	}
	return nil
}

type CUIFontFilePackagePB_CUIEncryptedFontFilePB struct {
	EncryptedContents    []byte   `protobuf:"bytes,1,opt,name=encrypted_contents,json=encryptedContents" json:"encrypted_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) Reset() {
	*m = CUIFontFilePackagePB_CUIEncryptedFontFilePB{}
}
func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) String() string {
	return proto.CompactTextString(m)
}
func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) ProtoMessage() {}
func (*CUIFontFilePackagePB_CUIEncryptedFontFilePB) Descriptor() ([]byte, []int) {
	return fileDescriptor_uifontfile_format_03661485664839fa, []int{1, 0}
}
func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB.Unmarshal(m, b)
}
func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB.Marshal(b, m, deterministic)
}
func (dst *CUIFontFilePackagePB_CUIEncryptedFontFilePB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB.Merge(dst, src)
}
func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) XXX_Size() int {
	return xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB.Size(m)
}
func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) XXX_DiscardUnknown() {
	xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB.DiscardUnknown(m)
}

var xxx_messageInfo_CUIFontFilePackagePB_CUIEncryptedFontFilePB proto.InternalMessageInfo

func (m *CUIFontFilePackagePB_CUIEncryptedFontFilePB) GetEncryptedContents() []byte {
	if m != nil {
		return m.EncryptedContents
	}
	return nil
}

func init() {
	proto.RegisterType((*CUIFontFilePB)(nil), "CUIFontFilePB")
	proto.RegisterType((*CUIFontFilePackagePB)(nil), "CUIFontFilePackagePB")
	proto.RegisterType((*CUIFontFilePackagePB_CUIEncryptedFontFilePB)(nil), "CUIFontFilePackagePB.CUIEncryptedFontFilePB")
}

func init() {
	proto.RegisterFile("uifontfile_format.proto", fileDescriptor_uifontfile_format_03661485664839fa)
}

var fileDescriptor_uifontfile_format_03661485664839fa = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x2d, 0x9e, 0xac, 0x80, 0xda, 0x10, 0x5d, 0x3c, 0x2d, 0xc4, 0xc4, 0x1d, 0x80, 0x29,
	0x1e, 0xbc, 0x83, 0xa2, 0x5e, 0x8c, 0x59, 0xa2, 0x07, 0x0f, 0x36, 0xa5, 0xbc, 0xcd, 0x05, 0xd6,
	0xb7, 0xb4, 0x6f, 0x26, 0xdc, 0xfc, 0xb8, 0x7e, 0x0c, 0xb3, 0x31, 0xc0, 0x18, 0x8e, 0xfd, 0xfd,
	0x7f, 0xfd, 0xbf, 0xf6, 0xf1, 0xb3, 0x22, 0x8d, 0xd1, 0x50, 0x9c, 0x2e, 0x40, 0xc6, 0x68, 0x33,
	0x45, 0x83, 0xdc, 0x22, 0x61, 0x57, 0xf3, 0xd6, 0xf8, 0xf5, 0x69, 0x82, 0x86, 0x26, 0xe9, 0x02,
	0x5e, 0x46, 0xe2, 0x82, 0xb7, 0x4b, 0x53, 0x56, 0xaa, 0x51, 0x19, 0x78, 0xcc, 0x67, 0xc1, 0x41,
	0xd4, 0x8c, 0x6b, 0xe7, 0x59, 0x65, 0x20, 0x7a, 0x5c, 0x60, 0x0e, 0x86, 0x96, 0x79, 0xd9, 0x67,
	0x48, 0xce, 0x14, 0x29, 0xaf, 0xe1, 0xb3, 0xa0, 0x19, 0x1d, 0xaf, 0x93, 0xb2, 0xf5, 0x4e, 0x91,
	0xea, 0xfe, 0x30, 0xde, 0xf9, 0x3b, 0x45, 0xe9, 0xb9, 0x4a, 0xca, 0x61, 0x97, 0xfc, 0x28, 0x5f,
	0x1d, 0xe4, 0x17, 0x58, 0x97, 0xa2, 0xf1, 0x98, 0xdf, 0x08, 0x5a, 0x51, 0xbb, 0xc6, 0x6f, 0x2b,
	0x2a, 0x3e, 0x78, 0x07, 0x8c, 0xb6, 0xcb, 0x9c, 0x60, 0x26, 0x37, 0xef, 0x73, 0x5e, 0xc3, 0xdf,
	0x0f, 0x0e, 0x87, 0xbd, 0xc1, 0xae, 0xf6, 0x12, 0xde, 0xaf, 0x2f, 0x6d, 0x7f, 0x18, 0x09, 0xf8,
	0x0f, 0xdd, 0xf9, 0x03, 0x3f, 0xdd, 0x6d, 0x8b, 0x3e, 0xdf, 0xfa, 0x52, 0xa3, 0x21, 0x30, 0xe4,
	0xaa, 0x9d, 0x34, 0xa3, 0x93, 0x4d, 0x32, 0xae, 0x83, 0xd1, 0xed, 0x23, 0x7b, 0xbf, 0x4a, 0x52,
	0xfa, 0x2c, 0xa6, 0x03, 0x8d, 0x59, 0x78, 0x7d, 0x33, 0x0f, 0x13, 0xec, 0x3b, 0x02, 0x95, 0xf5,
	0x2d, 0x38, 0x2c, 0xac, 0x06, 0x17, 0x56, 0xcb, 0x9f, 0x16, 0x71, 0x38, 0x43, 0x52, 0xc3, 0x6f,
	0xb6, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x21, 0x4b, 0xbe, 0xd5, 0xa2, 0x01, 0x00, 0x00,
}
