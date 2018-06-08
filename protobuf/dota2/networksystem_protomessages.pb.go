// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networksystem_protomessages.proto

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

type NetMessageSplitscreenUserChanged struct {
	Slot                 *uint32  `protobuf:"varint,1,opt,name=slot" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetMessageSplitscreenUserChanged) Reset()         { *m = NetMessageSplitscreenUserChanged{} }
func (m *NetMessageSplitscreenUserChanged) String() string { return proto.CompactTextString(m) }
func (*NetMessageSplitscreenUserChanged) ProtoMessage()    {}
func (*NetMessageSplitscreenUserChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63, []int{0}
}
func (m *NetMessageSplitscreenUserChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetMessageSplitscreenUserChanged.Unmarshal(m, b)
}
func (m *NetMessageSplitscreenUserChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetMessageSplitscreenUserChanged.Marshal(b, m, deterministic)
}
func (dst *NetMessageSplitscreenUserChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetMessageSplitscreenUserChanged.Merge(dst, src)
}
func (m *NetMessageSplitscreenUserChanged) XXX_Size() int {
	return xxx_messageInfo_NetMessageSplitscreenUserChanged.Size(m)
}
func (m *NetMessageSplitscreenUserChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_NetMessageSplitscreenUserChanged.DiscardUnknown(m)
}

var xxx_messageInfo_NetMessageSplitscreenUserChanged proto.InternalMessageInfo

func (m *NetMessageSplitscreenUserChanged) GetSlot() uint32 {
	if m != nil && m.Slot != nil {
		return *m.Slot
	}
	return 0
}

type NetMessageConnectionClosed struct {
	Reason               *uint32  `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetMessageConnectionClosed) Reset()         { *m = NetMessageConnectionClosed{} }
func (m *NetMessageConnectionClosed) String() string { return proto.CompactTextString(m) }
func (*NetMessageConnectionClosed) ProtoMessage()    {}
func (*NetMessageConnectionClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63, []int{1}
}
func (m *NetMessageConnectionClosed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetMessageConnectionClosed.Unmarshal(m, b)
}
func (m *NetMessageConnectionClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetMessageConnectionClosed.Marshal(b, m, deterministic)
}
func (dst *NetMessageConnectionClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetMessageConnectionClosed.Merge(dst, src)
}
func (m *NetMessageConnectionClosed) XXX_Size() int {
	return xxx_messageInfo_NetMessageConnectionClosed.Size(m)
}
func (m *NetMessageConnectionClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_NetMessageConnectionClosed.DiscardUnknown(m)
}

var xxx_messageInfo_NetMessageConnectionClosed proto.InternalMessageInfo

func (m *NetMessageConnectionClosed) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type NetMessageConnectionCrashed struct {
	Reason               *uint32  `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetMessageConnectionCrashed) Reset()         { *m = NetMessageConnectionCrashed{} }
func (m *NetMessageConnectionCrashed) String() string { return proto.CompactTextString(m) }
func (*NetMessageConnectionCrashed) ProtoMessage()    {}
func (*NetMessageConnectionCrashed) Descriptor() ([]byte, []int) {
	return fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63, []int{2}
}
func (m *NetMessageConnectionCrashed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetMessageConnectionCrashed.Unmarshal(m, b)
}
func (m *NetMessageConnectionCrashed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetMessageConnectionCrashed.Marshal(b, m, deterministic)
}
func (dst *NetMessageConnectionCrashed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetMessageConnectionCrashed.Merge(dst, src)
}
func (m *NetMessageConnectionCrashed) XXX_Size() int {
	return xxx_messageInfo_NetMessageConnectionCrashed.Size(m)
}
func (m *NetMessageConnectionCrashed) XXX_DiscardUnknown() {
	xxx_messageInfo_NetMessageConnectionCrashed.DiscardUnknown(m)
}

var xxx_messageInfo_NetMessageConnectionCrashed proto.InternalMessageInfo

func (m *NetMessageConnectionCrashed) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type NetMessagePacketStart struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetMessagePacketStart) Reset()         { *m = NetMessagePacketStart{} }
func (m *NetMessagePacketStart) String() string { return proto.CompactTextString(m) }
func (*NetMessagePacketStart) ProtoMessage()    {}
func (*NetMessagePacketStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63, []int{3}
}
func (m *NetMessagePacketStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetMessagePacketStart.Unmarshal(m, b)
}
func (m *NetMessagePacketStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetMessagePacketStart.Marshal(b, m, deterministic)
}
func (dst *NetMessagePacketStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetMessagePacketStart.Merge(dst, src)
}
func (m *NetMessagePacketStart) XXX_Size() int {
	return xxx_messageInfo_NetMessagePacketStart.Size(m)
}
func (m *NetMessagePacketStart) XXX_DiscardUnknown() {
	xxx_messageInfo_NetMessagePacketStart.DiscardUnknown(m)
}

var xxx_messageInfo_NetMessagePacketStart proto.InternalMessageInfo

type NetMessagePacketEnd struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetMessagePacketEnd) Reset()         { *m = NetMessagePacketEnd{} }
func (m *NetMessagePacketEnd) String() string { return proto.CompactTextString(m) }
func (*NetMessagePacketEnd) ProtoMessage()    {}
func (*NetMessagePacketEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63, []int{4}
}
func (m *NetMessagePacketEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetMessagePacketEnd.Unmarshal(m, b)
}
func (m *NetMessagePacketEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetMessagePacketEnd.Marshal(b, m, deterministic)
}
func (dst *NetMessagePacketEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetMessagePacketEnd.Merge(dst, src)
}
func (m *NetMessagePacketEnd) XXX_Size() int {
	return xxx_messageInfo_NetMessagePacketEnd.Size(m)
}
func (m *NetMessagePacketEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_NetMessagePacketEnd.DiscardUnknown(m)
}

var xxx_messageInfo_NetMessagePacketEnd proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NetMessageSplitscreenUserChanged)(nil), "NetMessageSplitscreenUserChanged")
	proto.RegisterType((*NetMessageConnectionClosed)(nil), "NetMessageConnectionClosed")
	proto.RegisterType((*NetMessageConnectionCrashed)(nil), "NetMessageConnectionCrashed")
	proto.RegisterType((*NetMessagePacketStart)(nil), "NetMessagePacketStart")
	proto.RegisterType((*NetMessagePacketEnd)(nil), "NetMessagePacketEnd")
}

func init() {
	proto.RegisterFile("networksystem_protomessages.proto", fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63)
}

var fileDescriptor_networksystem_protomessages_fcdb3c3721ad3a63 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xcf, 0x3f, 0x4b, 0xc3, 0x60,
	0x10, 0xc7, 0x71, 0x03, 0xe2, 0xf0, 0x80, 0x4b, 0xa4, 0x2a, 0xba, 0xd4, 0x4c, 0x2e, 0x6d, 0xfc,
	0x57, 0x5f, 0x80, 0xc1, 0x51, 0x11, 0x8b, 0x8b, 0x8b, 0x5c, 0x9f, 0xfc, 0x4c, 0x42, 0x92, 0xbb,
	0x72, 0x77, 0x41, 0xdc, 0x7c, 0xe9, 0x42, 0x2c, 0x14, 0x44, 0xdc, 0xee, 0x7b, 0xc7, 0x67, 0xb8,
	0x70, 0xc6, 0xf0, 0x0f, 0xd1, 0xd6, 0x3e, 0xcd, 0xd1, 0xbf, 0xad, 0x55, 0x5c, 0x7a, 0x98, 0x51,
	0x05, 0x9b, 0x8f, 0x95, 0xdd, 0x86, 0xe9, 0x23, 0xfc, 0xe1, 0x67, 0xb9, 0x5c, 0x77, 0x8d, 0x5b,
	0x54, 0x80, 0x5f, 0x0c, 0x5a, 0xd4, 0xc4, 0x15, 0xca, 0x34, 0x0d, 0xbb, 0xd6, 0x89, 0x1f, 0x27,
	0xd3, 0xe4, 0x7c, 0xff, 0x79, 0x9c, 0xb3, 0x9b, 0x70, 0xb2, 0x75, 0x85, 0x30, 0x23, 0x7a, 0x23,
	0x5c, 0x74, 0x62, 0x28, 0xd3, 0xc3, 0xb0, 0xa7, 0x20, 0x13, 0xde, 0x98, 0x4d, 0x65, 0x8b, 0x70,
	0xfa, 0xa7, 0x52, 0xb2, 0xfa, 0x1f, 0x76, 0x14, 0x26, 0x5b, 0xf6, 0x44, 0xb1, 0x85, 0x2f, 0x9d,
	0xd4, 0xb3, 0x49, 0x38, 0xf8, 0x7d, 0xb8, 0xe7, 0xf2, 0x6e, 0xf1, 0x7a, 0x51, 0x35, 0x5e, 0x0f,
	0xab, 0x79, 0x94, 0x3e, 0xbf, 0xbc, 0x6e, 0xf3, 0x4a, 0x66, 0xe6, 0xa0, 0x7e, 0xa6, 0x30, 0x19,
	0x34, 0xc2, 0xf2, 0xf1, 0xff, 0xd5, 0xf0, 0x9e, 0x97, 0xe2, 0x74, 0xf5, 0x95, 0xec, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0xdc, 0x73, 0x67, 0x2f, 0x01, 0x00, 0x00,
}
