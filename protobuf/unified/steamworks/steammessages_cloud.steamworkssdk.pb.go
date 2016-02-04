// Code generated by protoc-gen-go.
// source: steammessages_cloud.steamworkssdk.proto
// DO NOT EDIT!

/*
Package steamworks is a generated protocol buffer package.

It is generated from these files:
	steammessages_cloud.steamworkssdk.proto
	steammessages_oauth.steamworkssdk.proto
	steammessages_publishedfile.steamworkssdk.proto

It has these top-level messages:
	CCloud_GetUploadServerInfo_Request
	CCloud_GetUploadServerInfo_Response
	CCloud_GetFileDetails_Request
	CCloud_UserFile
	CCloud_GetFileDetails_Response
	CCloud_EnumerateUserFiles_Request
	CCloud_EnumerateUserFiles_Response
	CCloud_Delete_Request
	CCloud_Delete_Response
	COAuthToken_ImplicitGrantNoPrompt_Request
	COAuthToken_ImplicitGrantNoPrompt_Response
	CPublishedFile_Subscribe_Request
	CPublishedFile_Subscribe_Response
	CPublishedFile_Unsubscribe_Request
	CPublishedFile_Unsubscribe_Response
	CPublishedFile_Publish_Request
	CPublishedFile_Publish_Response
	CPublishedFile_GetDetails_Request
	PublishedFileDetails
	CPublishedFile_GetDetails_Response
	CPublishedFile_GetUserFiles_Request
	CPublishedFile_GetUserFiles_Response
	CPublishedFile_Update_Request
	CPublishedFile_Update_Response
	CPublishedFile_RefreshVotingQueue_Request
	CPublishedFile_RefreshVotingQueue_Response
*/
package steamworks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CCloud_GetUploadServerInfo_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Request) Reset()         { *m = CCloud_GetUploadServerInfo_Request{} }
func (m *CCloud_GetUploadServerInfo_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Request) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *CCloud_GetUploadServerInfo_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_GetUploadServerInfo_Response struct {
	ServerUrl        *string `protobuf:"bytes,1,opt,name=server_url" json:"server_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetUploadServerInfo_Response) Reset()         { *m = CCloud_GetUploadServerInfo_Response{} }
func (m *CCloud_GetUploadServerInfo_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_GetUploadServerInfo_Response) ProtoMessage()    {}
func (*CCloud_GetUploadServerInfo_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func (m *CCloud_GetUploadServerInfo_Response) GetServerUrl() string {
	if m != nil && m.ServerUrl != nil {
		return *m.ServerUrl
	}
	return ""
}

type CCloud_GetFileDetails_Request struct {
	Ugcid            *uint64 `protobuf:"varint,1,opt,name=ugcid" json:"ugcid,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_GetFileDetails_Request) Reset()                    { *m = CCloud_GetFileDetails_Request{} }
func (m *CCloud_GetFileDetails_Request) String() string            { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Request) ProtoMessage()               {}
func (*CCloud_GetFileDetails_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CCloud_GetFileDetails_Request) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_GetFileDetails_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_UserFile struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Ugcid            *uint64 `protobuf:"varint,2,opt,name=ugcid" json:"ugcid,omitempty"`
	Filename         *string `protobuf:"bytes,3,opt,name=filename" json:"filename,omitempty"`
	Timestamp        *uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	FileSize         *uint32 `protobuf:"varint,5,opt,name=file_size" json:"file_size,omitempty"`
	Url              *string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	SteamidCreator   *uint64 `protobuf:"fixed64,7,opt,name=steamid_creator" json:"steamid_creator,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_UserFile) Reset()                    { *m = CCloud_UserFile{} }
func (m *CCloud_UserFile) String() string            { return proto.CompactTextString(m) }
func (*CCloud_UserFile) ProtoMessage()               {}
func (*CCloud_UserFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CCloud_UserFile) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_UserFile) GetUgcid() uint64 {
	if m != nil && m.Ugcid != nil {
		return *m.Ugcid
	}
	return 0
}

func (m *CCloud_UserFile) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_UserFile) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CCloud_UserFile) GetFileSize() uint32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CCloud_UserFile) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *CCloud_UserFile) GetSteamidCreator() uint64 {
	if m != nil && m.SteamidCreator != nil {
		return *m.SteamidCreator
	}
	return 0
}

type CCloud_GetFileDetails_Response struct {
	Details          *CCloud_UserFile `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CCloud_GetFileDetails_Response) Reset()                    { *m = CCloud_GetFileDetails_Response{} }
func (m *CCloud_GetFileDetails_Response) String() string            { return proto.CompactTextString(m) }
func (*CCloud_GetFileDetails_Response) ProtoMessage()               {}
func (*CCloud_GetFileDetails_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CCloud_GetFileDetails_Response) GetDetails() *CCloud_UserFile {
	if m != nil {
		return m.Details
	}
	return nil
}

type CCloud_EnumerateUserFiles_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ExtendedDetails  *bool   `protobuf:"varint,2,opt,name=extended_details" json:"extended_details,omitempty"`
	Count            *uint32 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	StartIndex       *uint32 `protobuf:"varint,4,opt,name=start_index" json:"start_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Request) Reset()         { *m = CCloud_EnumerateUserFiles_Request{} }
func (m *CCloud_EnumerateUserFiles_Request) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Request) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *CCloud_EnumerateUserFiles_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetExtendedDetails() bool {
	if m != nil && m.ExtendedDetails != nil {
		return *m.ExtendedDetails
	}
	return false
}

func (m *CCloud_EnumerateUserFiles_Request) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CCloud_EnumerateUserFiles_Request) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

type CCloud_EnumerateUserFiles_Response struct {
	Files            []*CCloud_UserFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	TotalFiles       *uint32            `protobuf:"varint,2,opt,name=total_files" json:"total_files,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CCloud_EnumerateUserFiles_Response) Reset()         { *m = CCloud_EnumerateUserFiles_Response{} }
func (m *CCloud_EnumerateUserFiles_Response) String() string { return proto.CompactTextString(m) }
func (*CCloud_EnumerateUserFiles_Response) ProtoMessage()    {}
func (*CCloud_EnumerateUserFiles_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

func (m *CCloud_EnumerateUserFiles_Response) GetFiles() []*CCloud_UserFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *CCloud_EnumerateUserFiles_Response) GetTotalFiles() uint32 {
	if m != nil && m.TotalFiles != nil {
		return *m.TotalFiles
	}
	return 0
}

type CCloud_Delete_Request struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CCloud_Delete_Request) Reset()                    { *m = CCloud_Delete_Request{} }
func (m *CCloud_Delete_Request) String() string            { return proto.CompactTextString(m) }
func (*CCloud_Delete_Request) ProtoMessage()               {}
func (*CCloud_Delete_Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CCloud_Delete_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CCloud_Delete_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CCloud_Delete_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CCloud_Delete_Response) Reset()                    { *m = CCloud_Delete_Response{} }
func (m *CCloud_Delete_Response) String() string            { return proto.CompactTextString(m) }
func (*CCloud_Delete_Response) ProtoMessage()               {}
func (*CCloud_Delete_Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*CCloud_GetUploadServerInfo_Request)(nil), "CCloud_GetUploadServerInfo_Request")
	proto.RegisterType((*CCloud_GetUploadServerInfo_Response)(nil), "CCloud_GetUploadServerInfo_Response")
	proto.RegisterType((*CCloud_GetFileDetails_Request)(nil), "CCloud_GetFileDetails_Request")
	proto.RegisterType((*CCloud_UserFile)(nil), "CCloud_UserFile")
	proto.RegisterType((*CCloud_GetFileDetails_Response)(nil), "CCloud_GetFileDetails_Response")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Request)(nil), "CCloud_EnumerateUserFiles_Request")
	proto.RegisterType((*CCloud_EnumerateUserFiles_Response)(nil), "CCloud_EnumerateUserFiles_Response")
	proto.RegisterType((*CCloud_Delete_Request)(nil), "CCloud_Delete_Request")
	proto.RegisterType((*CCloud_Delete_Response)(nil), "CCloud_Delete_Response")
}

var fileDescriptor0 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x6f, 0x1c, 0x35,
	0x10, 0xc7, 0x75, 0x24, 0x97, 0xb6, 0x8e, 0x42, 0x83, 0x2b, 0xda, 0xd3, 0x21, 0x1a, 0x67, 0xd3,
	0xaa, 0xa9, 0x40, 0xa6, 0xaa, 0x00, 0x09, 0x90, 0x90, 0x7a, 0x39, 0xa8, 0x2a, 0x81, 0x90, 0x12,
	0x45, 0xa8, 0x20, 0xb4, 0xf2, 0xdd, 0xce, 0x5d, 0xac, 0x78, 0xed, 0xc5, 0xf6, 0x36, 0x01, 0xf5,
	0x01, 0xf1, 0x0c, 0x7f, 0x01, 0xfc, 0x1b, 0xbc, 0xf0, 0x1f, 0xf1, 0xc0, 0xff, 0xc0, 0xd8, 0xfb,
	0xe3, 0xf6, 0xc2, 0x91, 0xd0, 0xb7, 0x5b, 0x7b, 0x66, 0xfc, 0x99, 0x99, 0xef, 0xcc, 0x91, 0x07,
	0xce, 0x83, 0xc8, 0x73, 0x70, 0x4e, 0xcc, 0xc1, 0xa5, 0x53, 0x65, 0xca, 0x8c, 0xc7, 0xb3, 0x33,
	0x63, 0x4f, 0x9d, 0xcb, 0x4e, 0x79, 0x61, 0x8d, 0x37, 0x43, 0xbe, 0x6c, 0x58, 0x6a, 0x39, 0x93,
	0x90, 0xa5, 0x13, 0xe1, 0x60, 0x95, 0x7d, 0x92, 0x91, 0xe4, 0xe0, 0x20, 0x44, 0x4b, 0x9f, 0x82,
	0x3f, 0x2e, 0x94, 0x11, 0xd9, 0x11, 0xd8, 0x17, 0x60, 0x9f, 0xe9, 0x99, 0x49, 0x0f, 0xe1, 0xfb,
	0x12, 0x9c, 0xa7, 0x9f, 0x92, 0xbe, 0x28, 0x0a, 0x99, 0x0d, 0x7a, 0xac, 0xb7, 0xbf, 0x35, 0x7a,
	0xef, 0xe7, 0x3f, 0x06, 0xef, 0x3c, 0x29, 0x0a, 0xf6, 0x6c, 0xcc, 0xbc, 0x61, 0x67, 0x27, 0x72,
	0x7a, 0xc2, 0x04, 0x9b, 0x49, 0x05, 0xec, 0x4c, 0x2a, 0xc5, 0x26, 0xc0, 0xca, 0x18, 0x0b, 0x32,
	0x34, 0xe0, 0xc9, 0x47, 0x64, 0xef, 0xd2, 0x57, 0x5c, 0x61, 0xb4, 0x03, 0x4a, 0x09, 0x71, 0xf1,
	0x38, 0x2d, 0xad, 0x8a, 0x6f, 0xdd, 0x48, 0x7e, 0xe9, 0x91, 0xb7, 0x17, 0xbe, 0x9f, 0xe3, 0x13,
	0x63, 0xf0, 0x42, 0x2a, 0xd7, 0xc2, 0x7d, 0x42, 0xfa, 0xe5, 0x7c, 0x5a, 0xc3, 0xad, 0x8f, 0xde,
	0x45, 0xb8, 0x7d, 0x04, 0x33, 0x33, 0xe6, 0x4f, 0x80, 0x45, 0xd7, 0x0a, 0x0d, 0x51, 0xe7, 0xe0,
	0x59, 0x56, 0xf9, 0xb3, 0x99, 0xb1, 0x9c, 0xf2, 0x26, 0xb3, 0xd7, 0x62, 0x66, 0x3b, 0xe8, 0xfc,
	0x56, 0x93, 0x19, 0x7a, 0x47, 0xbf, 0x09, 0x28, 0xa3, 0xe7, 0x2e, 0x66, 0xf2, 0x6b, 0x8f, 0xdc,
	0xac, 0x71, 0x8e, 0x91, 0x35, 0xf0, 0xd0, 0xad, 0xa5, 0xea, 0x84, 0xcf, 0x8a, 0x27, 0x84, 0x5c,
	0xa7, 0xdb, 0xe4, 0x7a, 0x88, 0xa2, 0x45, 0x0e, 0x83, 0xb5, 0x90, 0x12, 0x7d, 0x83, 0xdc, 0xf0,
	0x12, 0x5b, 0xe4, 0x45, 0x5e, 0x0c, 0xd6, 0xa3, 0x11, 0x1e, 0x05, 0xa3, 0xd4, 0xc9, 0x1f, 0x61,
	0xd0, 0x8f, 0x61, 0x36, 0xc9, 0x5a, 0xa8, 0xc2, 0x46, 0x74, 0xb9, 0x43, 0x6e, 0xc6, 0xee, 0xc9,
	0x2c, 0x9d, 0x5a, 0x10, 0xde, 0xd8, 0xc1, 0x35, 0xbc, 0xd8, 0x48, 0x0e, 0xc8, 0xdd, 0xff, 0xaa,
	0x4e, 0x5d, 0xd4, 0x5d, 0x72, 0xad, 0xce, 0x38, 0xf2, 0x6d, 0x3e, 0xde, 0xe6, 0x17, 0x12, 0x48,
	0xfe, 0x5a, 0x23, 0xbb, 0xf5, 0xd9, 0x67, 0xba, 0xcc, 0xc1, 0x0a, 0x0f, 0xcd, 0xe5, 0xa2, 0xce,
	0xef, 0x2f, 0x8b, 0xe0, 0x3e, 0x96, 0x6a, 0x77, 0x21, 0x02, 0x68, 0x1c, 0xdb, 0xba, 0x39, 0x6c,
	0x01, 0xa7, 0xbf, 0xf5, 0xc8, 0x36, 0x9c, 0x7b, 0xd0, 0xa8, 0x85, 0xb4, 0x01, 0x09, 0x95, 0xb9,
	0x3e, 0xfa, 0xa9, 0x87, 0x21, 0x5e, 0xee, 0x7f, 0x55, 0x78, 0x69, 0xb4, 0x50, 0x0f, 0x19, 0xa6,
	0xc0, 0x1a, 0xdb, 0xb6, 0x4d, 0x13, 0x31, 0x3d, 0x65, 0x46, 0x77, 0xc2, 0xce, 0x4c, 0xa9, 0x33,
	0xce, 0xc6, 0x30, 0x13, 0xa5, 0xf2, 0xa1, 0x2f, 0x78, 0xaf, 0x7e, 0x60, 0x16, 0x7c, 0x69, 0x75,
	0x90, 0x1c, 0x9a, 0x8a, 0x40, 0x96, 0x31, 0xa1, 0x33, 0x76, 0xfc, 0xf4, 0x20, 0xfc, 0xac, 0x05,
	0xd1, 0x8d, 0x41, 0x5f, 0x92, 0xfe, 0x14, 0x7f, 0xf9, 0xd8, 0x99, 0xad, 0x51, 0x8e, 0x40, 0xb2,
	0x03, 0xf4, 0xa5, 0x38, 0x97, 0x79, 0x99, 0x33, 0xcc, 0x6d, 0x02, 0x36, 0x44, 0xb0, 0xe0, 0x9a,
	0x37, 0xab, 0xe7, 0x2a, 0x34, 0xe9, 0xd8, 0x54, 0x28, 0xb5, 0x0c, 0x25, 0x58, 0x5e, 0xfb, 0xa3,
	0xe3, 0x07, 0x8f, 0x1e, 0xd5, 0x4f, 0x37, 0x98, 0x9c, 0x5a, 0xb2, 0x89, 0x22, 0xb0, 0x3e, 0x95,
	0x98, 0xf1, 0x79, 0x94, 0xc2, 0xd6, 0xe8, 0x3b, 0x64, 0x78, 0xde, 0x61, 0x38, 0x0a, 0x16, 0x52,
	0xcf, 0x59, 0x34, 0x0a, 0x71, 0x27, 0x30, 0x97, 0xba, 0x2d, 0x38, 0x9a, 0x31, 0xe1, 0x97, 0x5f,
	0x0e, 0x79, 0x46, 0x2b, 0x1d, 0x1c, 0xeb, 0xc4, 0x95, 0x74, 0x9e, 0x27, 0xdf, 0xb4, 0x03, 0xbf,
	0xb2, 0xd5, 0xb5, 0x68, 0x76, 0x48, 0x3f, 0xc2, 0x62, 0xaf, 0xd7, 0x56, 0x49, 0x86, 0xde, 0x22,
	0x9b, 0xde, 0x78, 0xa1, 0xd2, 0xca, 0x2c, 0x4e, 0x4f, 0xf2, 0x9c, 0xbc, 0x59, 0xdb, 0x8d, 0x41,
	0x81, 0x87, 0x56, 0x3a, 0xdd, 0x19, 0x88, 0x63, 0xfd, 0xca, 0x73, 0x37, 0x20, 0xb7, 0x2f, 0x86,
	0xae, 0x50, 0x1f, 0xff, 0xbd, 0x4e, 0xfa, 0xf1, 0x86, 0xfe, 0xde, 0x23, 0xb7, 0x56, 0xec, 0x17,
	0xba, 0xc7, 0xaf, 0x5e, 0x71, 0xc3, 0x7b, 0xfc, 0x7f, 0x6c, 0xa8, 0xe4, 0x63, 0xe4, 0xfc, 0xf0,
	0x30, 0x36, 0xd0, 0x45, 0xd0, 0xe3, 0xc3, 0x2f, 0x9a, 0xfa, 0xe2, 0x42, 0x2d, 0x50, 0x24, 0x71,
	0x37, 0xb3, 0x6a, 0x8f, 0x85, 0x1d, 0x83, 0x32, 0x28, 0xf1, 0x8b, 0xd3, 0x33, 0xf2, 0xfa, 0xf2,
	0x8c, 0xd2, 0xbb, 0xfc, 0xd2, 0xcd, 0x36, 0xdc, 0xe1, 0x97, 0xcf, 0x76, 0x72, 0x0f, 0x71, 0x58,
	0x83, 0xd3, 0x0c, 0x4c, 0xd0, 0x43, 0x67, 0xed, 0x71, 0xfa, 0x67, 0x8f, 0xd0, 0x7f, 0x37, 0x9b,
	0x26, 0xfc, 0xca, 0x99, 0x1f, 0xee, 0xf1, 0xab, 0xc5, 0x92, 0x7c, 0x8b, 0x14, 0x5f, 0xb7, 0x06,
	0xae, 0xf3, 0xb6, 0xeb, 0x14, 0x20, 0x94, 0x49, 0xb0, 0xb9, 0x7c, 0x01, 0xba, 0x1a, 0xce, 0x31,
	0x6f, 0xd1, 0xcb, 0x22, 0xa8, 0x76, 0x31, 0x22, 0xc2, 0xa3, 0x69, 0xd8, 0x91, 0x9c, 0x66, 0x64,
	0xa3, 0xea, 0x38, 0xbd, 0xcd, 0x57, 0x8a, 0x6b, 0x78, 0x87, 0xaf, 0x56, 0x46, 0xf2, 0x10, 0xb9,
	0xee, 0x57, 0x87, 0xae, 0xf9, 0x7b, 0x9a, 0x59, 0x93, 0xc7, 0x6e, 0x05, 0xa4, 0x07, 0xae, 0xea,
	0x16, 0x1f, 0x46, 0xd3, 0x27, 0xb1, 0x6b, 0x72, 0x0a, 0x91, 0xfa, 0x28, 0xec, 0xdb, 0x3a, 0x97,
	0xd0, 0xd8, 0x38, 0x69, 0x8e, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x80, 0x16, 0x28, 0xba, 0x8b,
	0x07, 0x00, 0x00,
}
