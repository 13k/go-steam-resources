// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content_manifest.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContentManifestPayload struct {
	Mappings         []*ContentManifestPayload_FileMapping `protobuf:"bytes,1,rep,name=mappings" json:"mappings,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *ContentManifestPayload) Reset()                    { *m = ContentManifestPayload{} }
func (m *ContentManifestPayload) String() string            { return proto.CompactTextString(m) }
func (*ContentManifestPayload) ProtoMessage()               {}
func (*ContentManifestPayload) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ContentManifestPayload) GetMappings() []*ContentManifestPayload_FileMapping {
	if m != nil {
		return m.Mappings
	}
	return nil
}

type ContentManifestPayload_FileMapping struct {
	Filename         *string                                         `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Size             *uint64                                         `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Flags            *uint32                                         `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	ShaFilename      []byte                                          `protobuf:"bytes,4,opt,name=sha_filename,json=shaFilename" json:"sha_filename,omitempty"`
	ShaContent       []byte                                          `protobuf:"bytes,5,opt,name=sha_content,json=shaContent" json:"sha_content,omitempty"`
	Chunks           []*ContentManifestPayload_FileMapping_ChunkData `protobuf:"bytes,6,rep,name=chunks" json:"chunks,omitempty"`
	Linktarget       *string                                         `protobuf:"bytes,7,opt,name=linktarget" json:"linktarget,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *ContentManifestPayload_FileMapping) Reset()         { *m = ContentManifestPayload_FileMapping{} }
func (m *ContentManifestPayload_FileMapping) String() string { return proto.CompactTextString(m) }
func (*ContentManifestPayload_FileMapping) ProtoMessage()    {}
func (*ContentManifestPayload_FileMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

func (m *ContentManifestPayload_FileMapping) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *ContentManifestPayload_FileMapping) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping) GetShaFilename() []byte {
	if m != nil {
		return m.ShaFilename
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetShaContent() []byte {
	if m != nil {
		return m.ShaContent
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetChunks() []*ContentManifestPayload_FileMapping_ChunkData {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping) GetLinktarget() string {
	if m != nil && m.Linktarget != nil {
		return *m.Linktarget
	}
	return ""
}

type ContentManifestPayload_FileMapping_ChunkData struct {
	Sha              []byte  `protobuf:"bytes,1,opt,name=sha" json:"sha,omitempty"`
	Crc              *uint32 `protobuf:"fixed32,2,opt,name=crc" json:"crc,omitempty"`
	Offset           *uint64 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	CbOriginal       *uint32 `protobuf:"varint,4,opt,name=cb_original,json=cbOriginal" json:"cb_original,omitempty"`
	CbCompressed     *uint32 `protobuf:"varint,5,opt,name=cb_compressed,json=cbCompressed" json:"cb_compressed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ContentManifestPayload_FileMapping_ChunkData) Reset() {
	*m = ContentManifestPayload_FileMapping_ChunkData{}
}
func (m *ContentManifestPayload_FileMapping_ChunkData) String() string {
	return proto.CompactTextString(m)
}
func (*ContentManifestPayload_FileMapping_ChunkData) ProtoMessage() {}
func (*ContentManifestPayload_FileMapping_ChunkData) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0, 0}
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetSha() []byte {
	if m != nil {
		return m.Sha
	}
	return nil
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCrc() uint32 {
	if m != nil && m.Crc != nil {
		return *m.Crc
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCbOriginal() uint32 {
	if m != nil && m.CbOriginal != nil {
		return *m.CbOriginal
	}
	return 0
}

func (m *ContentManifestPayload_FileMapping_ChunkData) GetCbCompressed() uint32 {
	if m != nil && m.CbCompressed != nil {
		return *m.CbCompressed
	}
	return 0
}

type ContentManifestMetadata struct {
	DepotId            *uint32 `protobuf:"varint,1,opt,name=depot_id,json=depotId" json:"depot_id,omitempty"`
	GidManifest        *uint64 `protobuf:"varint,2,opt,name=gid_manifest,json=gidManifest" json:"gid_manifest,omitempty"`
	CreationTime       *uint32 `protobuf:"varint,3,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	FilenamesEncrypted *bool   `protobuf:"varint,4,opt,name=filenames_encrypted,json=filenamesEncrypted" json:"filenames_encrypted,omitempty"`
	CbDiskOriginal     *uint64 `protobuf:"varint,5,opt,name=cb_disk_original,json=cbDiskOriginal" json:"cb_disk_original,omitempty"`
	CbDiskCompressed   *uint64 `protobuf:"varint,6,opt,name=cb_disk_compressed,json=cbDiskCompressed" json:"cb_disk_compressed,omitempty"`
	UniqueChunks       *uint32 `protobuf:"varint,7,opt,name=unique_chunks,json=uniqueChunks" json:"unique_chunks,omitempty"`
	CrcEncrypted       *uint32 `protobuf:"varint,8,opt,name=crc_encrypted,json=crcEncrypted" json:"crc_encrypted,omitempty"`
	CrcClear           *uint32 `protobuf:"varint,9,opt,name=crc_clear,json=crcClear" json:"crc_clear,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *ContentManifestMetadata) Reset()                    { *m = ContentManifestMetadata{} }
func (m *ContentManifestMetadata) String() string            { return proto.CompactTextString(m) }
func (*ContentManifestMetadata) ProtoMessage()               {}
func (*ContentManifestMetadata) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ContentManifestMetadata) GetDepotId() uint32 {
	if m != nil && m.DepotId != nil {
		return *m.DepotId
	}
	return 0
}

func (m *ContentManifestMetadata) GetGidManifest() uint64 {
	if m != nil && m.GidManifest != nil {
		return *m.GidManifest
	}
	return 0
}

func (m *ContentManifestMetadata) GetCreationTime() uint32 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *ContentManifestMetadata) GetFilenamesEncrypted() bool {
	if m != nil && m.FilenamesEncrypted != nil {
		return *m.FilenamesEncrypted
	}
	return false
}

func (m *ContentManifestMetadata) GetCbDiskOriginal() uint64 {
	if m != nil && m.CbDiskOriginal != nil {
		return *m.CbDiskOriginal
	}
	return 0
}

func (m *ContentManifestMetadata) GetCbDiskCompressed() uint64 {
	if m != nil && m.CbDiskCompressed != nil {
		return *m.CbDiskCompressed
	}
	return 0
}

func (m *ContentManifestMetadata) GetUniqueChunks() uint32 {
	if m != nil && m.UniqueChunks != nil {
		return *m.UniqueChunks
	}
	return 0
}

func (m *ContentManifestMetadata) GetCrcEncrypted() uint32 {
	if m != nil && m.CrcEncrypted != nil {
		return *m.CrcEncrypted
	}
	return 0
}

func (m *ContentManifestMetadata) GetCrcClear() uint32 {
	if m != nil && m.CrcClear != nil {
		return *m.CrcClear
	}
	return 0
}

type ContentManifestSignature struct {
	Signature        []byte `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ContentManifestSignature) Reset()                    { *m = ContentManifestSignature{} }
func (m *ContentManifestSignature) String() string            { return proto.CompactTextString(m) }
func (*ContentManifestSignature) ProtoMessage()               {}
func (*ContentManifestSignature) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ContentManifestSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ContentManifestPayload)(nil), "ContentManifestPayload")
	proto.RegisterType((*ContentManifestPayload_FileMapping)(nil), "ContentManifestPayload.FileMapping")
	proto.RegisterType((*ContentManifestPayload_FileMapping_ChunkData)(nil), "ContentManifestPayload.FileMapping.ChunkData")
	proto.RegisterType((*ContentManifestMetadata)(nil), "ContentManifestMetadata")
	proto.RegisterType((*ContentManifestSignature)(nil), "ContentManifestSignature")
}

func init() { proto.RegisterFile("content_manifest.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xc9, 0xf6, 0x5f, 0x3a, 0x6d, 0x51, 0x65, 0x50, 0x09, 0x05, 0x41, 0xe9, 0x5e, 0x72,
	0x80, 0x22, 0x71, 0xe2, 0x86, 0x44, 0x77, 0x57, 0x70, 0xa8, 0x40, 0x86, 0x7b, 0xe4, 0x38, 0xd3,
	0xc4, 0x6a, 0xe2, 0x84, 0xd8, 0x3d, 0x2c, 0x27, 0x5e, 0x00, 0x1e, 0x14, 0xf1, 0x10, 0xc8, 0x8e,
	0x93, 0x56, 0x2b, 0x0e, 0xdc, 0x3c, 0x33, 0xdf, 0xd8, 0xdf, 0x6f, 0xc6, 0xb0, 0xe0, 0xa5, 0xd4,
	0x28, 0x75, 0x54, 0x30, 0x29, 0xf6, 0xa8, 0xf4, 0xa6, 0xaa, 0x4b, 0x5d, 0xae, 0x7f, 0xf7, 0x60,
	0xb1, 0x6d, 0x4a, 0x3b, 0x57, 0xf9, 0xcc, 0x6e, 0xf3, 0x92, 0x25, 0xe4, 0x1d, 0xf8, 0x05, 0xab,
	0x2a, 0x21, 0x53, 0x15, 0x78, 0xab, 0x5e, 0x38, 0x79, 0x73, 0xb9, 0xf9, 0xb7, 0x74, 0x73, 0x23,
	0x72, 0xdc, 0x35, 0x5a, 0xda, 0x35, 0x2d, 0x7f, 0xf5, 0x60, 0x72, 0x56, 0x21, 0x4b, 0xf0, 0xf7,
	0x22, 0x47, 0xc9, 0x0a, 0x0c, 0xbc, 0x95, 0x17, 0x8e, 0x69, 0x17, 0x13, 0x02, 0x7d, 0x25, 0xbe,
	0x63, 0x70, 0xb1, 0xf2, 0xc2, 0x3e, 0xb5, 0x67, 0xf2, 0x10, 0x06, 0xfb, 0x9c, 0xa5, 0x2a, 0xe8,
	0xad, 0xbc, 0x70, 0x46, 0x9b, 0x80, 0xbc, 0x80, 0xa9, 0xca, 0x58, 0xd4, 0xdd, 0xd4, 0x5f, 0x79,
	0xe1, 0x94, 0x4e, 0x54, 0xc6, 0x6e, 0xda, 0xcb, 0x9e, 0x83, 0x09, 0x23, 0x87, 0x1c, 0x0c, 0xac,
	0x02, 0x54, 0xc6, 0x9c, 0x7d, 0x72, 0x0d, 0x43, 0x9e, 0x1d, 0xe5, 0x41, 0x05, 0x43, 0x0b, 0xf6,
	0xea, 0x3f, 0xc0, 0x36, 0x5b, 0xd3, 0x71, 0xc5, 0x34, 0xa3, 0xae, 0x99, 0x3c, 0x03, 0xc8, 0x85,
	0x3c, 0x68, 0x56, 0xa7, 0xa8, 0x83, 0x91, 0x45, 0x3a, 0xcb, 0x2c, 0x7f, 0x7a, 0x30, 0xee, 0xba,
	0xc8, 0x1c, 0x7a, 0x2a, 0x63, 0x96, 0x7c, 0x4a, 0xcd, 0xd1, 0x64, 0x78, 0xcd, 0x2d, 0xf3, 0x88,
	0x9a, 0x23, 0x59, 0xc0, 0xb0, 0xdc, 0xef, 0x15, 0x6a, 0xcb, 0xdc, 0xa7, 0x2e, 0x32, 0x44, 0x3c,
	0x8e, 0xca, 0x5a, 0xa4, 0x42, 0xb2, 0xdc, 0x32, 0xcf, 0x28, 0xf0, 0xf8, 0x93, 0xcb, 0x90, 0x4b,
	0x98, 0xf1, 0x38, 0xe2, 0x65, 0x51, 0xd5, 0xa8, 0x14, 0x26, 0x16, 0x7a, 0x46, 0xa7, 0x3c, 0xde,
	0x76, 0xb9, 0xf5, 0x9f, 0x0b, 0x78, 0x74, 0x07, 0x74, 0x87, 0x9a, 0x25, 0xc6, 0xdd, 0x63, 0xf0,
	0x13, 0xac, 0x4a, 0x1d, 0x89, 0xc4, 0x5a, 0x9c, 0xd1, 0x91, 0x8d, 0x3f, 0x26, 0x66, 0xe2, 0xa9,
	0x48, 0xba, 0x9f, 0xe3, 0x76, 0x34, 0x49, 0x45, 0xd2, 0xde, 0x62, 0x9f, 0xaf, 0x91, 0x69, 0x51,
	0xca, 0x48, 0x8b, 0x02, 0xdd, 0xca, 0xa6, 0x6d, 0xf2, 0xab, 0x28, 0x90, 0xbc, 0x86, 0x07, 0xed,
	0xd6, 0x54, 0x84, 0x92, 0xd7, 0xb7, 0x95, 0xc6, 0xc4, 0xc2, 0xf8, 0x94, 0x74, 0xa5, 0xeb, 0xb6,
	0x42, 0x42, 0x98, 0xf3, 0x38, 0x4a, 0x84, 0x3a, 0x9c, 0xd0, 0x07, 0xf6, 0xf1, 0xfb, 0x3c, 0xbe,
	0x12, 0xea, 0xd0, 0xe1, 0xbf, 0x04, 0xd2, 0x2a, 0xcf, 0x66, 0x30, 0xb4, 0xda, 0x79, 0xa3, 0x3d,
	0xcd, 0xc1, 0xb8, 0x3d, 0x4a, 0xf1, 0xed, 0x88, 0x91, 0xfb, 0x05, 0xa3, 0xc6, 0x6d, 0x93, 0xdc,
	0x36, 0xcb, 0xb5, 0x48, 0xfc, 0xcc, 0xa7, 0xdf, 0x22, 0xf1, 0x93, 0xc3, 0x27, 0x30, 0x36, 0x22,
	0x9e, 0x23, 0xab, 0x83, 0xb1, 0x15, 0xf8, 0xbc, 0xe6, 0x5b, 0x13, 0xaf, 0xdf, 0x42, 0x70, 0x67,
	0xda, 0x5f, 0x44, 0x2a, 0x99, 0x3e, 0xd6, 0x48, 0x9e, 0xc2, 0x58, 0xb5, 0x81, 0xfb, 0x12, 0xa7,
	0xc4, 0xfb, 0xc1, 0x07, 0xef, 0x87, 0x77, 0xef, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x06,
	0x35, 0x96, 0xb5, 0x03, 0x00, 0x00,
}
