// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_partnerapps.steamclient.proto

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CPartnerApps_RequestUploadToken_Request struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_RequestUploadToken_Request) Reset() {
	*m = CPartnerApps_RequestUploadToken_Request{}
}
func (m *CPartnerApps_RequestUploadToken_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_RequestUploadToken_Request) ProtoMessage()    {}
func (*CPartnerApps_RequestUploadToken_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{0}
}

func (m *CPartnerApps_RequestUploadToken_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CPartnerApps_RequestUploadToken_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPartnerApps_RequestUploadToken_Response struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token,json=uploadToken" json:"upload_token,omitempty"`
	Location         *string `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,3,opt,name=routing_id,json=routingId" json:"routing_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_RequestUploadToken_Response) Reset() {
	*m = CPartnerApps_RequestUploadToken_Response{}
}
func (m *CPartnerApps_RequestUploadToken_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_RequestUploadToken_Response) ProtoMessage()    {}
func (*CPartnerApps_RequestUploadToken_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1}
}

func (m *CPartnerApps_RequestUploadToken_Response) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_RequestUploadToken_Response) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *CPartnerApps_RequestUploadToken_Response) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

type CPartnerApps_FinishUpload_Request struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token,json=uploadToken" json:"upload_token,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,2,opt,name=routing_id,json=routingId" json:"routing_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUpload_Request) Reset()         { *m = CPartnerApps_FinishUpload_Request{} }
func (m *CPartnerApps_FinishUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUpload_Request) ProtoMessage()    {}
func (*CPartnerApps_FinishUpload_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{2}
}

func (m *CPartnerApps_FinishUpload_Request) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_FinishUpload_Request) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

func (m *CPartnerApps_FinishUpload_Request) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_FinishUploadKVSign_Response struct {
	SignedInstallscript *string `protobuf:"bytes,1,opt,name=signed_installscript,json=signedInstallscript" json:"signed_installscript,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadKVSign_Response) Reset() {
	*m = CPartnerApps_FinishUploadKVSign_Response{}
}
func (m *CPartnerApps_FinishUploadKVSign_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUploadKVSign_Response) ProtoMessage()    {}
func (*CPartnerApps_FinishUploadKVSign_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{3}
}

func (m *CPartnerApps_FinishUploadKVSign_Response) GetSignedInstallscript() string {
	if m != nil && m.SignedInstallscript != nil {
		return *m.SignedInstallscript
	}
	return ""
}

type CPartnerApps_FinishUploadLegacyDRM_Request struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token,json=uploadToken" json:"upload_token,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,2,opt,name=routing_id,json=routingId" json:"routing_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,3,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Flags            *uint32 `protobuf:"varint,4,opt,name=flags" json:"flags,omitempty"`
	ToolName         *string `protobuf:"bytes,5,opt,name=tool_name,json=toolName" json:"tool_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) Reset() {
	*m = CPartnerApps_FinishUploadLegacyDRM_Request{}
}
func (m *CPartnerApps_FinishUploadLegacyDRM_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CPartnerApps_FinishUploadLegacyDRM_Request) ProtoMessage() {}
func (*CPartnerApps_FinishUploadLegacyDRM_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{4}
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetToolName() string {
	if m != nil && m.ToolName != nil {
		return *m.ToolName
	}
	return ""
}

type CPartnerApps_FinishUploadLegacyDRM_Response struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Response) Reset() {
	*m = CPartnerApps_FinishUploadLegacyDRM_Response{}
}
func (m *CPartnerApps_FinishUploadLegacyDRM_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CPartnerApps_FinishUploadLegacyDRM_Response) ProtoMessage() {}
func (*CPartnerApps_FinishUploadLegacyDRM_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{5}
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Response) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

type CPartnerApps_FinishUpload_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPartnerApps_FinishUpload_Response) Reset()         { *m = CPartnerApps_FinishUpload_Response{} }
func (m *CPartnerApps_FinishUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUpload_Response) ProtoMessage()    {}
func (*CPartnerApps_FinishUpload_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{6}
}

type CPartnerApps_FindDRMUploads_Request struct {
	AppId            *int32 `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPartnerApps_FindDRMUploads_Request) Reset()         { *m = CPartnerApps_FindDRMUploads_Request{} }
func (m *CPartnerApps_FindDRMUploads_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FindDRMUploads_Request) ProtoMessage()    {}
func (*CPartnerApps_FindDRMUploads_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{7}
}

func (m *CPartnerApps_FindDRMUploads_Request) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_ExistingDRMUpload struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	ActorId          *int32  `protobuf:"varint,3,opt,name=actor_id,json=actorId" json:"actor_id,omitempty"`
	SuppliedName     *string `protobuf:"bytes,5,opt,name=supplied_name,json=suppliedName" json:"supplied_name,omitempty"`
	Flags            *uint32 `protobuf:"varint,6,opt,name=flags" json:"flags,omitempty"`
	ModType          *string `protobuf:"bytes,7,opt,name=mod_type,json=modType" json:"mod_type,omitempty"`
	Timestamp        *uint32 `protobuf:"fixed32,8,opt,name=timestamp" json:"timestamp,omitempty"`
	OrigFileId       *string `protobuf:"bytes,9,opt,name=orig_file_id,json=origFileId" json:"orig_file_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_ExistingDRMUpload) Reset()                    { *m = CPartnerApps_ExistingDRMUpload{} }
func (m *CPartnerApps_ExistingDRMUpload) String() string            { return proto.CompactTextString(m) }
func (*CPartnerApps_ExistingDRMUpload) ProtoMessage()               {}
func (*CPartnerApps_ExistingDRMUpload) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

func (m *CPartnerApps_ExistingDRMUpload) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetActorId() int32 {
	if m != nil && m.ActorId != nil {
		return *m.ActorId
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetSuppliedName() string {
	if m != nil && m.SuppliedName != nil {
		return *m.SuppliedName
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetModType() string {
	if m != nil && m.ModType != nil {
		return *m.ModType
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetOrigFileId() string {
	if m != nil && m.OrigFileId != nil {
		return *m.OrigFileId
	}
	return ""
}

type CPartnerApps_FindDRMUploads_Response struct {
	Uploads          []*CPartnerApps_ExistingDRMUpload `protobuf:"bytes,1,rep,name=uploads" json:"uploads,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CPartnerApps_FindDRMUploads_Response) Reset()         { *m = CPartnerApps_FindDRMUploads_Response{} }
func (m *CPartnerApps_FindDRMUploads_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FindDRMUploads_Response) ProtoMessage()    {}
func (*CPartnerApps_FindDRMUploads_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{9}
}

func (m *CPartnerApps_FindDRMUploads_Response) GetUploads() []*CPartnerApps_ExistingDRMUpload {
	if m != nil {
		return m.Uploads
	}
	return nil
}

type CPartnerApps_Download_Request struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	AppId            *int32  `protobuf:"varint,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_Download_Request) Reset()                    { *m = CPartnerApps_Download_Request{} }
func (m *CPartnerApps_Download_Request) String() string            { return proto.CompactTextString(m) }
func (*CPartnerApps_Download_Request) ProtoMessage()               {}
func (*CPartnerApps_Download_Request) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func (m *CPartnerApps_Download_Request) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

func (m *CPartnerApps_Download_Request) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_Download_Response struct {
	DownloadUrl      *string `protobuf:"bytes,1,opt,name=download_url,json=downloadUrl" json:"download_url,omitempty"`
	AppId            *int32  `protobuf:"varint,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_Download_Response) Reset()         { *m = CPartnerApps_Download_Response{} }
func (m *CPartnerApps_Download_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_Download_Response) ProtoMessage()    {}
func (*CPartnerApps_Download_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{11}
}

func (m *CPartnerApps_Download_Response) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *CPartnerApps_Download_Response) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func init() {
	proto.RegisterType((*CPartnerApps_RequestUploadToken_Request)(nil), "CPartnerApps_RequestUploadToken_Request")
	proto.RegisterType((*CPartnerApps_RequestUploadToken_Response)(nil), "CPartnerApps_RequestUploadToken_Response")
	proto.RegisterType((*CPartnerApps_FinishUpload_Request)(nil), "CPartnerApps_FinishUpload_Request")
	proto.RegisterType((*CPartnerApps_FinishUploadKVSign_Response)(nil), "CPartnerApps_FinishUploadKVSign_Response")
	proto.RegisterType((*CPartnerApps_FinishUploadLegacyDRM_Request)(nil), "CPartnerApps_FinishUploadLegacyDRM_Request")
	proto.RegisterType((*CPartnerApps_FinishUploadLegacyDRM_Response)(nil), "CPartnerApps_FinishUploadLegacyDRM_Response")
	proto.RegisterType((*CPartnerApps_FinishUpload_Response)(nil), "CPartnerApps_FinishUpload_Response")
	proto.RegisterType((*CPartnerApps_FindDRMUploads_Request)(nil), "CPartnerApps_FindDRMUploads_Request")
	proto.RegisterType((*CPartnerApps_ExistingDRMUpload)(nil), "CPartnerApps_ExistingDRMUpload")
	proto.RegisterType((*CPartnerApps_FindDRMUploads_Response)(nil), "CPartnerApps_FindDRMUploads_Response")
	proto.RegisterType((*CPartnerApps_Download_Request)(nil), "CPartnerApps_Download_Request")
	proto.RegisterType((*CPartnerApps_Download_Response)(nil), "CPartnerApps_Download_Response")
}

func init() { proto.RegisterFile("steammessages_partnerapps.steamclient.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 907 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6f, 0xdc, 0x44,
	0x14, 0x96, 0x93, 0x26, 0xbb, 0xfb, 0x92, 0x72, 0x18, 0x1a, 0xe1, 0x2e, 0xb4, 0x9d, 0x3a, 0x41,
	0x2c, 0x4d, 0x64, 0x20, 0x1c, 0x10, 0x52, 0x0f, 0xd0, 0xa4, 0xa9, 0x56, 0xa5, 0x01, 0xb9, 0x2d,
	0x07, 0x10, 0xb2, 0x26, 0xf6, 0xac, 0x3b, 0xc2, 0xf6, 0x0c, 0x33, 0x63, 0xe8, 0x4a, 0x1c, 0x60,
	0x4f, 0x5c, 0xf8, 0x03, 0xb8, 0x73, 0x03, 0x21, 0x71, 0x58, 0x89, 0x3f, 0x0f, 0x8d, 0x7f, 0xec,
	0xda, 0x0e, 0xbb, 0x35, 0x87, 0xaa, 0xc7, 0xf7, 0xe6, 0xcd, 0x7b, 0xdf, 0x7c, 0xef, 0xcd, 0x37,
	0x03, 0x87, 0x4a, 0x53, 0x92, 0x24, 0x54, 0x29, 0x12, 0x51, 0xe5, 0x0b, 0x22, 0x75, 0x4a, 0x25,
	0x11, 0x42, 0xb9, 0xf9, 0x4a, 0x10, 0x33, 0x9a, 0x6a, 0x57, 0x48, 0xae, 0xf9, 0xf0, 0xa8, 0x19,
	0x9c, 0xa5, 0x6c, 0xc2, 0x68, 0xe8, 0x5f, 0x10, 0x45, 0x2f, 0x47, 0x3b, 0x5f, 0xc3, 0x3b, 0x27,
	0x5f, 0x14, 0xf9, 0x3e, 0x15, 0x42, 0xf9, 0x1e, 0xfd, 0x2e, 0xa3, 0x4a, 0x3f, 0x15, 0x31, 0x27,
	0xe1, 0x13, 0xfe, 0x2d, 0x4d, 0x2b, 0x17, 0x1a, 0x42, 0x7f, 0xc2, 0x62, 0x9a, 0x92, 0x84, 0xda,
	0x16, 0xb6, 0x46, 0x03, 0x6f, 0x61, 0xa3, 0x6b, 0xb0, 0x45, 0x84, 0x60, 0xa1, 0xbd, 0x81, 0xad,
	0xd1, 0x55, 0xaf, 0x30, 0x9c, 0x5f, 0x2c, 0x18, 0xbd, 0x38, 0xbb, 0x12, 0x3c, 0x55, 0x14, 0xdd,
	0x86, 0xdd, 0x2c, 0xf7, 0xfb, 0xda, 0x2c, 0xe4, 0x25, 0xae, 0x78, 0x3b, 0xd9, 0x32, 0xd6, 0x20,
	0x88, 0x79, 0x40, 0x34, 0xe3, 0x69, 0x5e, 0x68, 0xe0, 0x2d, 0x6c, 0x74, 0x03, 0x40, 0xf2, 0x4c,
	0xb3, 0x34, 0xf2, 0x59, 0x68, 0x6f, 0xe6, 0x9b, 0x07, 0xa5, 0x67, 0x1c, 0x3a, 0x3f, 0xc2, 0xed,
	0x06, 0x92, 0x33, 0x96, 0x32, 0xf5, 0xac, 0x00, 0xb2, 0x38, 0x61, 0x07, 0x08, 0xcd, 0x32, 0x1b,
	0xad, 0x32, 0x68, 0x0f, 0xb6, 0x89, 0x10, 0x15, 0x82, 0x82, 0x88, 0x71, 0xe8, 0x7c, 0xd3, 0xe2,
	0xa1, 0x5e, 0xfd, 0xe1, 0x97, 0x8f, 0x59, 0x54, 0xe3, 0xe1, 0x03, 0xb8, 0xa6, 0x58, 0x94, 0xd2,
	0xd0, 0x67, 0xa9, 0xd2, 0x24, 0x8e, 0x55, 0x20, 0x99, 0xd0, 0x25, 0xe5, 0xaf, 0x17, 0x6b, 0xe3,
	0xfa, 0x92, 0xf3, 0x8f, 0x05, 0x77, 0x56, 0xe6, 0xff, 0x8c, 0x46, 0x24, 0x98, 0x9e, 0x7a, 0x8f,
	0x5e, 0xfa, 0x31, 0xcd, 0x14, 0x4c, 0x62, 0x12, 0x29, 0xfb, 0x4a, 0xe1, 0xcd, 0x0d, 0xf4, 0x26,
	0x0c, 0x34, 0xe7, 0xb1, 0x9f, 0x0f, 0xce, 0x56, 0xd1, 0x36, 0xe3, 0x38, 0x27, 0x09, 0x75, 0xce,
	0xe0, 0xb0, 0x13, 0xf2, 0x92, 0x9c, 0x37, 0xa0, 0x67, 0x66, 0xce, 0x54, 0x2e, 0xf8, 0xd8, 0x36,
	0xe6, 0x38, 0x74, 0x0e, 0xc0, 0x59, 0xd7, 0xdf, 0x62, 0xbb, 0x73, 0x17, 0xf6, 0xdb, 0x51, 0xe1,
	0xa9, 0xf7, 0xa8, 0x08, 0x5b, 0x8c, 0x67, 0xed, 0x78, 0xa6, 0xc8, 0x56, 0xd5, 0xc5, 0x9f, 0x37,
	0xe0, 0x66, 0x63, 0xfb, 0xfd, 0xe7, 0x4c, 0x19, 0x46, 0x16, 0x29, 0x56, 0xe2, 0xab, 0xa5, 0xdc,
	0xa8, 0x33, 0x76, 0x1d, 0xfa, 0x24, 0xd0, 0x5c, 0x56, 0x54, 0x6e, 0x79, 0xbd, 0xdc, 0x1e, 0x87,
	0x68, 0x1f, 0xae, 0xaa, 0x4c, 0x88, 0xd8, 0x5c, 0xde, 0x1a, 0x75, 0xbb, 0x95, 0xf3, 0xbc, 0xbc,
	0x77, 0x05, 0xe3, 0xdb, 0x75, 0xc6, 0xaf, 0x43, 0x3f, 0xe1, 0xa1, 0xaf, 0xa7, 0x82, 0xda, 0xbd,
	0x7c, 0x57, 0x2f, 0xe1, 0xe1, 0x93, 0xa9, 0xa0, 0xe8, 0x2d, 0x18, 0x68, 0x96, 0x50, 0xa5, 0x49,
	0x22, 0xec, 0x3e, 0xb6, 0x46, 0x3d, 0x6f, 0xe9, 0x40, 0x18, 0x76, 0xb9, 0x64, 0x91, 0x5f, 0x9d,
	0x61, 0x90, 0x6f, 0x06, 0xe3, 0x3b, 0x2b, 0x78, 0x26, 0x70, 0xb0, 0x9e, 0xc1, 0xb2, 0x51, 0x1f,
	0x43, 0xaf, 0x98, 0x27, 0x65, 0x5b, 0x78, 0x73, 0xb4, 0x73, 0x7c, 0xcb, 0x5d, 0x4f, 0x9d, 0x57,
	0xc5, 0x3b, 0x9f, 0xc3, 0x8d, 0x46, 0xe8, 0x29, 0xff, 0x21, 0x6d, 0x5c, 0xd3, 0x8e, 0x24, 0x2f,
	0xfa, 0xf6, 0x55, 0xab, 0x6d, 0xb5, 0x84, 0x4b, 0xed, 0x09, 0x2b, 0x67, 0x26, 0xe3, 0x32, 0xed,
	0x4e, 0xe5, 0x7b, 0x2a, 0xe3, 0x15, 0xb9, 0x8f, 0xff, 0x04, 0xd8, 0xa9, 0xe5, 0x46, 0x7f, 0x5b,
	0x60, 0x97, 0x38, 0x8b, 0x8b, 0x5d, 0xd3, 0x3a, 0x34, 0x72, 0x3b, 0x6a, 0xed, 0xf0, 0x5d, 0xb7,
	0xab, 0x6e, 0x3a, 0x9f, 0xcc, 0xe6, 0xf6, 0xdd, 0x32, 0x00, 0x17, 0x1c, 0xe2, 0xfc, 0x5e, 0xe3,
	0x09, 0x97, 0xb8, 0xa1, 0x21, 0xd8, 0xf0, 0x54, 0xc5, 0xbc, 0x87, 0x8d, 0x94, 0xb0, 0x34, 0x42,
	0x7f, 0x59, 0xb0, 0x57, 0x26, 0x58, 0xb4, 0xe3, 0x95, 0x01, 0xa6, 0xcf, 0x69, 0x90, 0x69, 0x72,
	0x11, 0xd3, 0x16, 0x5a, 0x21, 0x79, 0x40, 0x95, 0x32, 0x80, 0xe7, 0x4b, 0xc0, 0x27, 0xf7, 0x1f,
	0xbc, 0x74, 0xc0, 0x0f, 0x66, 0x73, 0xfb, 0x64, 0x25, 0xe0, 0x20, 0x53, 0x9a, 0x27, 0x9d, 0x70,
	0xff, 0x61, 0x01, 0xba, 0x2c, 0xfd, 0xc8, 0x71, 0x5f, 0xf8, 0x34, 0xb5, 0xe1, 0xae, 0x79, 0x40,
	0x9c, 0xb3, 0xd9, 0xdc, 0xbe, 0x77, 0xc2, 0x93, 0x84, 0x69, 0x9c, 0x50, 0xfd, 0x8c, 0x1b, 0xb4,
	0x79, 0xb3, 0x31, 0x49, 0x5b, 0x03, 0x41, 0x26, 0x9a, 0x4a, 0xe3, 0x2e, 0xe1, 0x32, 0x85, 0x03,
	0x9e, 0x88, 0x98, 0x6a, 0x8a, 0x7e, 0xb7, 0x60, 0xaf, 0x5e, 0x67, 0xa9, 0x72, 0x87, 0x6e, 0xf7,
	0xd7, 0x66, 0x78, 0xe4, 0xfe, 0x0f, 0x81, 0x77, 0xde, 0x9f, 0xcd, 0xed, 0xa3, 0x26, 0xf8, 0xff,
	0x9a, 0x8a, 0x1a, 0xa9, 0xbf, 0xb5, 0x60, 0x2e, 0x26, 0xa2, 0x13, 0xaf, 0xfb, 0x6e, 0x87, 0x67,
	0xe3, 0xa3, 0xd9, 0xdc, 0xfe, 0xf0, 0x32, 0xa8, 0x15, 0x9d, 0xaf, 0x61, 0xfb, 0xd5, 0x82, 0xd7,
	0x9a, 0x0a, 0x89, 0x0e, 0xdc, 0x0e, 0x2f, 0xd0, 0xf0, 0x6d, 0xb7, 0x8b, 0xca, 0x3a, 0xee, 0x6c,
	0x6e, 0xdf, 0x31, 0x8b, 0x0a, 0x9f, 0xe3, 0x84, 0x2b, 0x8d, 0x25, 0x0d, 0x68, 0xaa, 0xb1, 0x90,
	0x8c, 0xcb, 0xb2, 0xa9, 0x0a, 0x5f, 0x4c, 0x71, 0xfe, 0x21, 0x43, 0x0f, 0xa1, 0x5f, 0x89, 0x1f,
	0xba, 0xe9, 0xae, 0x55, 0xd9, 0xe1, 0x2d, 0x77, 0xbd, 0x68, 0x0e, 0x8f, 0x67, 0x73, 0xdb, 0x7d,
	0x4c, 0xe5, 0xf7, 0x2c, 0xa0, 0x25, 0x2d, 0x2a, 0xe7, 0x85, 0x08, 0x81, 0x13, 0x92, 0x92, 0x88,
	0x26, 0x06, 0xcb, 0xc5, 0x14, 0x97, 0x3f, 0x56, 0x75, 0x6f, 0xf3, 0x27, 0xcb, 0xfa, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x12, 0xc6, 0xae, 0xcf, 0x0a, 0x00, 0x00,
}
