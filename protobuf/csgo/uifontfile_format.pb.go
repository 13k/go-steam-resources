// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uifontfile_format.proto

package csgo // import "github.com/13k/go-steam-resources/protobuf/csgo"

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
	return fileDescriptor_uifontfile_format_e7f059f4d2d744b5, []int{0}
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
	return fileDescriptor_uifontfile_format_e7f059f4d2d744b5, []int{1}
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
	return fileDescriptor_uifontfile_format_e7f059f4d2d744b5, []int{1, 0}
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
	proto.RegisterFile("uifontfile_format.proto", fileDescriptor_uifontfile_format_e7f059f4d2d744b5)
}

var fileDescriptor_uifontfile_format_e7f059f4d2d744b5 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xcd, 0x3c, 0x19, 0xb7, 0xa9, 0x61, 0x68, 0xf1, 0x54, 0x86, 0x60, 0x0f, 0x5b, 0x8b,
	0x0a, 0x7e, 0x80, 0x4d, 0xa7, 0x5e, 0x44, 0x0a, 0x7a, 0xf0, 0x60, 0xc8, 0xb2, 0xd7, 0x5a, 0xb6,
	0xe6, 0x95, 0xe4, 0x55, 0xd8, 0xcd, 0x8f, 0xeb, 0xc7, 0x90, 0x76, 0xdd, 0x26, 0xb2, 0x63, 0x7e,
	0xff, 0x5f, 0xfe, 0x2f, 0x79, 0xfc, 0xac, 0xcc, 0x12, 0x34, 0x94, 0x64, 0x0b, 0x90, 0x09, 0xda,
	0x5c, 0x51, 0x58, 0x58, 0x24, 0xec, 0x6b, 0xde, 0x19, 0xbf, 0x3e, 0x4d, 0xd0, 0xd0, 0x24, 0x5b,
	0xc0, 0xcb, 0x48, 0x5c, 0xf0, 0x6e, 0x65, 0xca, 0x5a, 0x35, 0x2a, 0x07, 0x8f, 0xf9, 0x2c, 0x38,
	0x88, 0xdb, 0x49, 0xe3, 0x3c, 0xab, 0x1c, 0xc4, 0x80, 0x0b, 0x2c, 0xc0, 0xd0, 0xb2, 0xa8, 0xfa,
	0x0c, 0xc9, 0x99, 0x22, 0xe5, 0xb5, 0x7c, 0x16, 0xb4, 0xe3, 0xe3, 0x75, 0x52, 0xb5, 0xde, 0x29,
	0x52, 0xfd, 0x1f, 0xc6, 0x7b, 0x7f, 0xa7, 0x28, 0x3d, 0x57, 0x69, 0x35, 0xec, 0x92, 0x1f, 0x15,
	0xab, 0x83, 0xfc, 0x02, 0xeb, 0x32, 0x34, 0x1e, 0xf3, 0x5b, 0x41, 0x27, 0xee, 0x36, 0xf8, 0x6d,
	0x45, 0xc5, 0x07, 0xef, 0x81, 0xd1, 0x76, 0x59, 0x10, 0xcc, 0xe4, 0xe6, 0x7d, 0xce, 0x6b, 0xf9,
	0xfb, 0xc1, 0xe1, 0xf5, 0x20, 0xdc, 0xd5, 0x5e, 0xc1, 0xfb, 0xf5, 0xa5, 0xed, 0x0f, 0x63, 0x01,
	0xff, 0xa1, 0x3b, 0x7f, 0xe0, 0xa7, 0xbb, 0x6d, 0x31, 0xe4, 0x5b, 0x5f, 0x6a, 0x34, 0x04, 0x86,
	0x5c, 0xbd, 0x93, 0x76, 0x7c, 0xb2, 0x49, 0xc6, 0x4d, 0x30, 0xba, 0x7d, 0x64, 0xef, 0x51, 0x9a,
	0xd1, 0x67, 0x39, 0x0d, 0x35, 0xe6, 0xd1, 0xd5, 0xcd, 0x3c, 0x4a, 0x71, 0xe8, 0x08, 0x54, 0x3e,
	0xb4, 0xe0, 0xb0, 0xb4, 0x1a, 0x5c, 0x54, 0x2f, 0x7f, 0x5a, 0x26, 0x91, 0x76, 0x29, 0x7e, 0xb3,
	0xbd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x81, 0x97, 0xff, 0xa1, 0x01, 0x00, 0x00,
}
