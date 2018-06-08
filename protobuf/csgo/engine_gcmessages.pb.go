// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine_gcmessages.proto

package csgo // import "github.com/13k/go-steam-resources/protobuf/csgo"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CEngineGotvSyncPacket struct {
	MatchId              *uint64  `protobuf:"varint,1,opt,name=match_id,json=matchId" json:"match_id,omitempty"`
	InstanceId           *uint32  `protobuf:"varint,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Signupfragment       *uint32  `protobuf:"varint,3,opt,name=signupfragment" json:"signupfragment,omitempty"`
	Currentfragment      *uint32  `protobuf:"varint,4,opt,name=currentfragment" json:"currentfragment,omitempty"`
	Tickrate             *float32 `protobuf:"fixed32,5,opt,name=tickrate" json:"tickrate,omitempty"`
	Tick                 *uint32  `protobuf:"varint,6,opt,name=tick" json:"tick,omitempty"`
	Rtdelay              *float32 `protobuf:"fixed32,8,opt,name=rtdelay" json:"rtdelay,omitempty"`
	Rcvage               *float32 `protobuf:"fixed32,9,opt,name=rcvage" json:"rcvage,omitempty"`
	KeyframeInterval     *float32 `protobuf:"fixed32,10,opt,name=keyframe_interval,json=keyframeInterval" json:"keyframe_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CEngineGotvSyncPacket) Reset()         { *m = CEngineGotvSyncPacket{} }
func (m *CEngineGotvSyncPacket) String() string { return proto.CompactTextString(m) }
func (*CEngineGotvSyncPacket) ProtoMessage()    {}
func (*CEngineGotvSyncPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_gcmessages_f7d518ac56cf1aa6, []int{0}
}
func (m *CEngineGotvSyncPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CEngineGotvSyncPacket.Unmarshal(m, b)
}
func (m *CEngineGotvSyncPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CEngineGotvSyncPacket.Marshal(b, m, deterministic)
}
func (dst *CEngineGotvSyncPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CEngineGotvSyncPacket.Merge(dst, src)
}
func (m *CEngineGotvSyncPacket) XXX_Size() int {
	return xxx_messageInfo_CEngineGotvSyncPacket.Size(m)
}
func (m *CEngineGotvSyncPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_CEngineGotvSyncPacket.DiscardUnknown(m)
}

var xxx_messageInfo_CEngineGotvSyncPacket proto.InternalMessageInfo

func (m *CEngineGotvSyncPacket) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetInstanceId() uint32 {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetSignupfragment() uint32 {
	if m != nil && m.Signupfragment != nil {
		return *m.Signupfragment
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetCurrentfragment() uint32 {
	if m != nil && m.Currentfragment != nil {
		return *m.Currentfragment
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetTickrate() float32 {
	if m != nil && m.Tickrate != nil {
		return *m.Tickrate
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetTick() uint32 {
	if m != nil && m.Tick != nil {
		return *m.Tick
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetRtdelay() float32 {
	if m != nil && m.Rtdelay != nil {
		return *m.Rtdelay
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetRcvage() float32 {
	if m != nil && m.Rcvage != nil {
		return *m.Rcvage
	}
	return 0
}

func (m *CEngineGotvSyncPacket) GetKeyframeInterval() float32 {
	if m != nil && m.KeyframeInterval != nil {
		return *m.KeyframeInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*CEngineGotvSyncPacket)(nil), "CEngineGotvSyncPacket")
}

func init() {
	proto.RegisterFile("engine_gcmessages.proto", fileDescriptor_engine_gcmessages_f7d518ac56cf1aa6)
}

var fileDescriptor_engine_gcmessages_f7d518ac56cf1aa6 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4e, 0x2a, 0x41,
	0x10, 0x85, 0xef, 0x70, 0x11, 0xb0, 0x8c, 0x7f, 0x9d, 0xa8, 0x2d, 0x1b, 0x27, 0x2e, 0xcc, 0x24,
	0x06, 0x26, 0x46, 0x9f, 0x40, 0x63, 0x0c, 0x3b, 0x83, 0x3b, 0x37, 0xa4, 0xe9, 0x29, 0x9a, 0xce,
	0x30, 0xdd, 0xa4, 0xba, 0x86, 0x84, 0x9d, 0x0f, 0xe5, 0x03, 0x1a, 0x5a, 0x06, 0x13, 0x76, 0xf5,
	0x9d, 0xf3, 0x9d, 0x4d, 0xc1, 0x15, 0x3a, 0x63, 0x1d, 0x4e, 0x8c, 0xae, 0x30, 0x04, 0x65, 0x30,
	0x0c, 0x97, 0xe4, 0xd9, 0xf7, 0x53, 0xe3, 0xbd, 0x59, 0x60, 0x1e, 0x69, 0x5a, 0xcf, 0xf2, 0x02,
	0x83, 0x26, 0xbb, 0x64, 0x4f, 0xbf, 0xc6, 0xed, 0x77, 0x0b, 0x2e, 0x5e, 0x5e, 0xe3, 0xfc, 0xcd,
	0xf3, 0xea, 0x63, 0xed, 0xf4, 0xbb, 0xd2, 0x25, 0xb2, 0xb8, 0x86, 0x5e, 0xa5, 0x58, 0xcf, 0x27,
	0xb6, 0x90, 0x49, 0x9a, 0x64, 0xed, 0x71, 0x37, 0xf2, 0xa8, 0x10, 0x37, 0x70, 0x64, 0x5d, 0x60,
	0xe5, 0x34, 0x6e, 0xda, 0x56, 0x9a, 0x64, 0xc7, 0x63, 0x68, 0xa2, 0x51, 0x21, 0xee, 0xe0, 0x24,
	0x58, 0xe3, 0xea, 0xe5, 0x8c, 0x94, 0xa9, 0xd0, 0xb1, 0xfc, 0x1f, 0x9d, 0xbd, 0x54, 0x64, 0x70,
	0xaa, 0x6b, 0x22, 0x74, 0xbc, 0x13, 0xdb, 0x51, 0xdc, 0x8f, 0x45, 0x1f, 0x7a, 0x6c, 0x75, 0x49,
	0x8a, 0x51, 0x1e, 0xa4, 0x49, 0xd6, 0x1a, 0xef, 0x58, 0x08, 0x68, 0x6f, 0x6e, 0xd9, 0x89, 0xd3,
	0x78, 0x0b, 0x09, 0x5d, 0xe2, 0x02, 0x17, 0x6a, 0x2d, 0x7b, 0x51, 0x6f, 0x50, 0x5c, 0x42, 0x87,
	0xf4, 0x4a, 0x19, 0x94, 0x87, 0xb1, 0xd8, 0x92, 0xb8, 0x87, 0xf3, 0x12, 0xd7, 0x33, 0x52, 0x15,
	0x4e, 0xac, 0x63, 0xa4, 0x95, 0x5a, 0x48, 0x88, 0xca, 0x59, 0x53, 0x8c, 0xb6, 0xf9, 0xf3, 0xd3,
	0x67, 0x6e, 0x2c, 0xcf, 0xeb, 0xe9, 0x50, 0xfb, 0x2a, 0x7f, 0x78, 0x2c, 0x73, 0xe3, 0x07, 0x81,
	0x51, 0x55, 0x03, 0xc2, 0xe0, 0x6b, 0xd2, 0x18, 0xfe, 0xbe, 0xae, 0x83, 0xf1, 0x5f, 0xc9, 0xbf,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xcf, 0xc5, 0x8c, 0xa8, 0x01, 0x00, 0x00,
}
