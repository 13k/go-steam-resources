// Code generated by protoc-gen-go.
// source: steammessages_linkfilter.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CCommunity_GetLinkFilterHashPrefixes_Request struct {
	HitType          *uint32 `protobuf:"varint,1,opt,name=hit_type" json:"hit_type,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Start            *uint64 `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Request) Reset() {
	*m = CCommunity_GetLinkFilterHashPrefixes_Request{}
}
func (m *CCommunity_GetLinkFilterHashPrefixes_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CCommunity_GetLinkFilterHashPrefixes_Request) ProtoMessage() {}
func (*CCommunity_GetLinkFilterHashPrefixes_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{0}
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Request) GetHitType() uint32 {
	if m != nil && m.HitType != nil {
		return *m.HitType
	}
	return 0
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Request) GetStart() uint64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

type CCommunity_GetLinkFilterHashPrefixes_Response struct {
	HashPrefixes     []uint32 `protobuf:"varint,1,rep,name=hash_prefixes" json:"hash_prefixes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Response) Reset() {
	*m = CCommunity_GetLinkFilterHashPrefixes_Response{}
}
func (m *CCommunity_GetLinkFilterHashPrefixes_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CCommunity_GetLinkFilterHashPrefixes_Response) ProtoMessage() {}
func (*CCommunity_GetLinkFilterHashPrefixes_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{1}
}

func (m *CCommunity_GetLinkFilterHashPrefixes_Response) GetHashPrefixes() []uint32 {
	if m != nil {
		return m.HashPrefixes
	}
	return nil
}

type CCommunity_GetLinkFilterHashes_Request struct {
	HitType          *uint32 `protobuf:"varint,1,opt,name=hit_type" json:"hit_type,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Start            *uint64 `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCommunity_GetLinkFilterHashes_Request) Reset() {
	*m = CCommunity_GetLinkFilterHashes_Request{}
}
func (m *CCommunity_GetLinkFilterHashes_Request) String() string { return proto.CompactTextString(m) }
func (*CCommunity_GetLinkFilterHashes_Request) ProtoMessage()    {}
func (*CCommunity_GetLinkFilterHashes_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{2}
}

func (m *CCommunity_GetLinkFilterHashes_Request) GetHitType() uint32 {
	if m != nil && m.HitType != nil {
		return *m.HitType
	}
	return 0
}

func (m *CCommunity_GetLinkFilterHashes_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCommunity_GetLinkFilterHashes_Request) GetStart() uint64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

type CCommunity_GetLinkFilterHashes_Response struct {
	Hashes           [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CCommunity_GetLinkFilterHashes_Response) Reset() {
	*m = CCommunity_GetLinkFilterHashes_Response{}
}
func (m *CCommunity_GetLinkFilterHashes_Response) String() string { return proto.CompactTextString(m) }
func (*CCommunity_GetLinkFilterHashes_Response) ProtoMessage()    {}
func (*CCommunity_GetLinkFilterHashes_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{3}
}

func (m *CCommunity_GetLinkFilterHashes_Response) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type CCommunity_GetLinkFilterListVersion_Request struct {
	HitType          *uint32 `protobuf:"varint,1,opt,name=hit_type" json:"hit_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCommunity_GetLinkFilterListVersion_Request) Reset() {
	*m = CCommunity_GetLinkFilterListVersion_Request{}
}
func (m *CCommunity_GetLinkFilterListVersion_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CCommunity_GetLinkFilterListVersion_Request) ProtoMessage() {}
func (*CCommunity_GetLinkFilterListVersion_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{4}
}

func (m *CCommunity_GetLinkFilterListVersion_Request) GetHitType() uint32 {
	if m != nil && m.HitType != nil {
		return *m.HitType
	}
	return 0
}

type CCommunity_GetLinkFilterListVersion_Response struct {
	Version          *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Count            *uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCommunity_GetLinkFilterListVersion_Response) Reset() {
	*m = CCommunity_GetLinkFilterListVersion_Response{}
}
func (m *CCommunity_GetLinkFilterListVersion_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CCommunity_GetLinkFilterListVersion_Response) ProtoMessage() {}
func (*CCommunity_GetLinkFilterListVersion_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{5}
}

func (m *CCommunity_GetLinkFilterListVersion_Response) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *CCommunity_GetLinkFilterListVersion_Response) GetCount() uint64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*CCommunity_GetLinkFilterHashPrefixes_Request)(nil), "CCommunity_GetLinkFilterHashPrefixes_Request")
	proto.RegisterType((*CCommunity_GetLinkFilterHashPrefixes_Response)(nil), "CCommunity_GetLinkFilterHashPrefixes_Response")
	proto.RegisterType((*CCommunity_GetLinkFilterHashes_Request)(nil), "CCommunity_GetLinkFilterHashes_Request")
	proto.RegisterType((*CCommunity_GetLinkFilterHashes_Response)(nil), "CCommunity_GetLinkFilterHashes_Response")
	proto.RegisterType((*CCommunity_GetLinkFilterListVersion_Request)(nil), "CCommunity_GetLinkFilterListVersion_Request")
	proto.RegisterType((*CCommunity_GetLinkFilterListVersion_Response)(nil), "CCommunity_GetLinkFilterListVersion_Response")
}

var fileDescriptor9 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe4, 0x55, 0xcd, 0x4e, 0x14, 0x41,
	0x10, 0xce, 0xf0, 0xe3, 0x4f, 0x47, 0x2e, 0xc3, 0x65, 0xdd, 0x8b, 0x95, 0x31, 0x11, 0x22, 0x4b,
	0x2b, 0x10, 0xe3, 0xc1, 0x03, 0xd9, 0xc5, 0xf0, 0xa3, 0xa8, 0x08, 0x84, 0x93, 0xc9, 0x66, 0x66,
	0xb6, 0x86, 0x69, 0xd9, 0x9d, 0x5e, 0xbb, 0x7b, 0x40, 0x3c, 0x11, 0x2f, 0x3e, 0x82, 0x27, 0x5f,
	0xc0, 0xbb, 0x07, 0x5f, 0xc0, 0x9b, 0x17, 0x9f, 0xc8, 0xea, 0x9e, 0x99, 0x05, 0x02, 0xab, 0xac,
	0x57, 0x8f, 0xd3, 0x5d, 0xfd, 0xd5, 0x57, 0xf5, 0x7d, 0x55, 0xc3, 0xee, 0x6b, 0x83, 0x61, 0xaf,
	0x87, 0x5a, 0x87, 0xfb, 0xa8, 0xdb, 0x5d, 0x91, 0x1d, 0x24, 0xa2, 0x6b, 0x50, 0x71, 0x77, 0x11,
	0x77, 0x05, 0x66, 0x86, 0xf7, 0x95, 0x34, 0xb2, 0xde, 0x38, 0x1f, 0x9b, 0x67, 0x22, 0x11, 0xd8,
	0x69, 0x47, 0xa1, 0xc6, 0x8b, 0xd1, 0xc1, 0x97, 0x31, 0xd6, 0x58, 0x59, 0x91, 0xbd, 0x1e, 0xc5,
	0x99, 0xe3, 0xf6, 0x1a, 0x9a, 0x4d, 0x02, 0x5f, 0x75, 0xe0, 0xeb, 0xa1, 0x4e, 0xb7, 0x14, 0x26,
	0xe2, 0x3d, 0xe1, 0x6c, 0xe3, 0xbb, 0x1c, 0xb5, 0xf1, 0xd7, 0xd8, 0x8d, 0x54, 0x98, 0xb6, 0x39,
	0xee, 0x63, 0xcd, 0x03, 0x6f, 0x76, 0xaa, 0xf5, 0xe8, 0xe3, 0xb7, 0xda, 0xc2, 0x6e, 0x8a, 0xa0,
	0xd0, 0x28, 0x81, 0x87, 0xd8, 0x01, 0x8a, 0xd0, 0x70, 0x24, 0xba, 0x5d, 0x88, 0x10, 0x0a, 0xa6,
	0x74, 0x6a, 0x24, 0x98, 0x54, 0x68, 0xb0, 0x6f, 0xb9, 0xff, 0x86, 0x4d, 0xc6, 0x32, 0xcf, 0x4c,
	0x6d, 0xcc, 0xa1, 0xbc, 0x20, 0x94, 0x0d, 0x8b, 0x92, 0xe5, 0xbd, 0x08, 0x15, 0xc8, 0xa4, 0x40,
	0xa1, 0x47, 0x15, 0x2e, 0x88, 0x0c, 0x42, 0xd0, 0x22, 0xdb, 0xef, 0x22, 0x44, 0xa1, 0x89, 0x53,
	0x0e, 0x3b, 0x7d, 0x8c, 0x45, 0x72, 0x0c, 0x0f, 0x21, 0x91, 0x0a, 0x32, 0x09, 0x5d, 0xd1, 0x13,
	0x86, 0xfb, 0xab, 0x6c, 0x52, 0x9b, 0x50, 0x99, 0xda, 0x38, 0xa1, 0x4f, 0xb4, 0x1e, 0x13, 0xfa,
	0x92, 0x45, 0x77, 0x87, 0x84, 0x00, 0x2e, 0x37, 0x1c, 0xa5, 0x98, 0x55, 0xf8, 0xf6, 0xd4, 0xa5,
	0xa4, 0x34, 0x0e, 0x1d, 0x35, 0x0f, 0x3e, 0x79, 0x6c, 0xfe, 0x8a, 0xfd, 0xd1, 0x7d, 0x99, 0x69,
	0xf4, 0xf7, 0xd8, 0x54, 0x4a, 0x17, 0xed, 0x7e, 0x79, 0x43, 0x5d, 0x1a, 0xa7, 0xfa, 0x96, 0x89,
	0xc1, 0x13, 0xcb, 0x20, 0x11, 0x4a, 0x1b, 0x58, 0x5a, 0x84, 0xc8, 0xe6, 0xa2, 0x32, 0x0d, 0x1d,
	0xee, 0xac, 0x37, 0x17, 0xc0, 0x3e, 0x43, 0x77, 0x84, 0x61, 0x9c, 0x5a, 0x2e, 0x44, 0xcd, 0x75,
	0x1d, 0x3b, 0x3c, 0xf8, 0x3c, 0xc6, 0xee, 0xfd, 0x89, 0xc9, 0xff, 0xab, 0xd1, 0x5b, 0x36, 0xf3,
	0xd7, 0xc6, 0x94, 0xe2, 0x2c, 0xb3, 0x6b, 0x45, 0x97, 0x9d, 0x2a, 0xb7, 0x5a, 0x0f, 0x28, 0xe7,
	0x5c, 0x93, 0x08, 0x91, 0x24, 0xb6, 0x9a, 0x42, 0x01, 0x4a, 0x96, 0xab, 0x8c, 0xba, 0x91, 0x28,
	0xd9, 0x03, 0x99, 0xab, 0x4a, 0x07, 0x1e, 0x9c, 0x78, 0x6c, 0x6e, 0x58, 0xb2, 0x4d, 0x42, 0xd9,
	0x43, 0xa5, 0x85, 0xcc, 0x06, 0x52, 0xbc, 0xbe, 0x20, 0xc5, 0xc0, 0x08, 0x87, 0x45, 0x68, 0x95,
	0xd4, 0xf6, 0x66, 0xd0, 0xfa, 0x19, 0x0d, 0x71, 0x9a, 0x67, 0x07, 0xa7, 0x0a, 0x55, 0x9c, 0x78,
	0xf0, 0xd5, 0x1b, 0x3e, 0xb2, 0xe7, 0x29, 0x94, 0x45, 0xaf, 0xb2, 0xeb, 0x65, 0x2e, 0x47, 0xe1,
	0x66, 0xe1, 0x86, 0xa6, 0x4b, 0x0c, 0x51, 0x4e, 0x0f, 0x21, 0xb7, 0x8a, 0x39, 0x27, 0x6e, 0x3c,
	0x1d, 0x98, 0x32, 0x96, 0x99, 0x09, 0x45, 0x56, 0x7a, 0x85, 0xfb, 0x8b, 0x67, 0xdd, 0x30, 0xd1,
	0xba, 0x4b, 0x28, 0x77, 0x2e, 0x71, 0x83, 0x7b, 0x5d, 0xd6, 0xb1, 0xf8, 0x63, 0x92, 0x4d, 0x0f,
	0xb8, 0x9e, 0xf2, 0xf4, 0x7f, 0x79, 0xec, 0xf6, 0xd0, 0x61, 0xf2, 0xe7, 0xf9, 0x28, 0x3b, 0xa9,
	0xce, 0xf9, 0x48, 0x23, 0x1a, 0xec, 0x11, 0xf3, 0x6d, 0x8a, 0x23, 0xa7, 0x9e, 0x55, 0x1e, 0xaa,
	0x91, 0x2d, 0xb5, 0x20, 0x27, 0x3a, 0xe7, 0x8a, 0xa2, 0x03, 0xae, 0xa0, 0x86, 0x35, 0x7b, 0xae,
	0xd1, 0x85, 0x14, 0x9b, 0x14, 0x62, 0x9a, 0x56, 0x6a, 0x1f, 0xf7, 0xbf, 0x7b, 0x6c, 0xfa, 0x12,
	0xfb, 0xf9, 0x33, 0xfc, 0x6a, 0x83, 0x5b, 0x9f, 0xe5, 0x57, 0x34, 0x72, 0xf0, 0x8a, 0x4a, 0x78,
	0x7e, 0xb1, 0x84, 0x7f, 0xe7, 0xfe, 0xd3, 0x63, 0xb5, 0x61, 0x56, 0xf2, 0x1b, 0x7c, 0x04, 0xcf,
	0xd7, 0xe7, 0xf9, 0x28, 0xf6, 0x0c, 0x76, 0xa9, 0x94, 0xad, 0x4b, 0x4b, 0xe9, 0xa0, 0x8e, 0x95,
	0x88, 0x2a, 0x6b, 0x56, 0x03, 0x53, 0x2d, 0x48, 0x37, 0x21, 0x95, 0x59, 0x07, 0x9b, 0x92, 0x96,
	0xcf, 0x07, 0xe4, 0xf5, 0x97, 0x84, 0xfa, 0xac, 0x09, 0x1a, 0xd5, 0xa1, 0x88, 0x8b, 0x92, 0x15,
	0xc6, 0x52, 0x75, 0x2c, 0x5a, 0x27, 0x34, 0x21, 0x84, 0x91, 0xcc, 0x0d, 0xec, 0xd8, 0x5f, 0x22,
	0x0c, 0xf8, 0x42, 0x9f, 0xac, 0x6b, 0x3b, 0x02, 0xf6, 0x27, 0x5b, 0xee, 0x45, 0xfa, 0x6c, 0x8d,
	0x9f, 0x78, 0xde, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x3d, 0xc5, 0x86, 0x88, 0x07, 0x00,
	0x00,
}
