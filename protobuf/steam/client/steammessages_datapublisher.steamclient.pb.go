// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: steam/client/steammessages_datapublisher.steamclient.proto

package client

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CDataPublisher_ClientContentCorruptionReport_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid           *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Depotid         *uint32 `protobuf:"varint,2,opt,name=depotid" json:"depotid,omitempty"`
	DownloadSource  *string `protobuf:"bytes,3,opt,name=download_source,json=downloadSource" json:"download_source,omitempty"`
	Objectid        *string `protobuf:"bytes,4,opt,name=objectid" json:"objectid,omitempty"`
	Cellid          *uint32 `protobuf:"varint,5,opt,name=cellid" json:"cellid,omitempty"`
	IsManifest      *bool   `protobuf:"varint,6,opt,name=is_manifest,json=isManifest" json:"is_manifest,omitempty"`
	ObjectSize      *uint64 `protobuf:"varint,7,opt,name=object_size,json=objectSize" json:"object_size,omitempty"`
	CorruptionType  *uint32 `protobuf:"varint,8,opt,name=corruption_type,json=corruptionType" json:"corruption_type,omitempty"`
	UsedHttps       *bool   `protobuf:"varint,9,opt,name=used_https,json=usedHttps" json:"used_https,omitempty"`
	OcProxyDetected *bool   `protobuf:"varint,10,opt,name=oc_proxy_detected,json=ocProxyDetected" json:"oc_proxy_detected,omitempty"`
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) Reset() {
	*x = CDataPublisher_ClientContentCorruptionReport_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDataPublisher_ClientContentCorruptionReport_Notification) ProtoMessage() {}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDataPublisher_ClientContentCorruptionReport_Notification.ProtoReflect.Descriptor instead.
func (*CDataPublisher_ClientContentCorruptionReport_Notification) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetDepotid() uint32 {
	if x != nil && x.Depotid != nil {
		return *x.Depotid
	}
	return 0
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetDownloadSource() string {
	if x != nil && x.DownloadSource != nil {
		return *x.DownloadSource
	}
	return ""
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetObjectid() string {
	if x != nil && x.Objectid != nil {
		return *x.Objectid
	}
	return ""
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetCellid() uint32 {
	if x != nil && x.Cellid != nil {
		return *x.Cellid
	}
	return 0
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetIsManifest() bool {
	if x != nil && x.IsManifest != nil {
		return *x.IsManifest
	}
	return false
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetObjectSize() uint64 {
	if x != nil && x.ObjectSize != nil {
		return *x.ObjectSize
	}
	return 0
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetCorruptionType() uint32 {
	if x != nil && x.CorruptionType != nil {
		return *x.CorruptionType
	}
	return 0
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetUsedHttps() bool {
	if x != nil && x.UsedHttps != nil {
		return *x.UsedHttps
	}
	return false
}

func (x *CDataPublisher_ClientContentCorruptionReport_Notification) GetOcProxyDetected() bool {
	if x != nil && x.OcProxyDetected != nil {
		return *x.OcProxyDetected
	}
	return false
}

type CValveHWSurvey_GetSurveySchedule_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Surveydatetoken        *string `protobuf:"bytes,1,opt,name=surveydatetoken" json:"surveydatetoken,omitempty"`
	Surveydatetokenversion *uint64 `protobuf:"fixed64,2,opt,name=surveydatetokenversion" json:"surveydatetokenversion,omitempty"`
}

func (x *CValveHWSurvey_GetSurveySchedule_Request) Reset() {
	*x = CValveHWSurvey_GetSurveySchedule_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CValveHWSurvey_GetSurveySchedule_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CValveHWSurvey_GetSurveySchedule_Request) ProtoMessage() {}

func (x *CValveHWSurvey_GetSurveySchedule_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CValveHWSurvey_GetSurveySchedule_Request.ProtoReflect.Descriptor instead.
func (*CValveHWSurvey_GetSurveySchedule_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CValveHWSurvey_GetSurveySchedule_Request) GetSurveydatetoken() string {
	if x != nil && x.Surveydatetoken != nil {
		return *x.Surveydatetoken
	}
	return ""
}

func (x *CValveHWSurvey_GetSurveySchedule_Request) GetSurveydatetokenversion() uint64 {
	if x != nil && x.Surveydatetokenversion != nil {
		return *x.Surveydatetokenversion
	}
	return 0
}

type CValveHWSurvey_GetSurveySchedule_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Surveydatetoken        *uint32 `protobuf:"varint,1,opt,name=surveydatetoken" json:"surveydatetoken,omitempty"`
	Surveydatetokenversion *uint64 `protobuf:"fixed64,2,opt,name=surveydatetokenversion" json:"surveydatetokenversion,omitempty"`
}

func (x *CValveHWSurvey_GetSurveySchedule_Response) Reset() {
	*x = CValveHWSurvey_GetSurveySchedule_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CValveHWSurvey_GetSurveySchedule_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CValveHWSurvey_GetSurveySchedule_Response) ProtoMessage() {}

func (x *CValveHWSurvey_GetSurveySchedule_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CValveHWSurvey_GetSurveySchedule_Response.ProtoReflect.Descriptor instead.
func (*CValveHWSurvey_GetSurveySchedule_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CValveHWSurvey_GetSurveySchedule_Response) GetSurveydatetoken() uint32 {
	if x != nil && x.Surveydatetoken != nil {
		return *x.Surveydatetoken
	}
	return 0
}

func (x *CValveHWSurvey_GetSurveySchedule_Response) GetSurveydatetokenversion() uint64 {
	if x != nil && x.Surveydatetokenversion != nil {
		return *x.Surveydatetokenversion
	}
	return 0
}

var File_steam_client_steammessages_datapublisher_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_datapublisher_steamclient_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x39, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x04, 0x0a, 0x39, 0x43, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f,
	0x74, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x82, 0xb5,
	0x18, 0x1c, 0x68, 0x6f, 0x73, 0x74, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0e,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1c, 0x82, 0xb5, 0x18, 0x18, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x20, 0x53, 0x48, 0x41, 0x20,
	0x6f, 0x72, 0x20, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x20, 0x49, 0x44, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x65, 0x6c, 0x6c,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x20, 0x43, 0x65, 0x6c, 0x6c, 0x20, 0x49, 0x44, 0x52, 0x06, 0x63, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x29, 0x82, 0xb5, 0x18, 0x25, 0x54,
	0x68, 0x65, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x20, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2c, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x20, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0a, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0x82, 0xb5, 0x18, 0x14, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x62, 0x79, 0x74, 0x65, 0x73, 0x52,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x63,
	0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x1e, 0x82, 0xb5, 0x18, 0x1a, 0x53, 0x65, 0x65, 0x20, 0x45, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x19, 0x82, 0xb5, 0x18, 0x15, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x77, 0x61, 0x73, 0x20, 0x48, 0x54,
	0x54, 0x50, 0x53, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x48, 0x74, 0x74, 0x70, 0x73, 0x12, 0x4b,
	0x0a, 0x11, 0x6f, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1f, 0x82, 0xb5, 0x18, 0x1b, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x0f, 0x6f, 0x63, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x28,
	0x43, 0x56, 0x61, 0x6c, 0x76, 0x65, 0x48, 0x57, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x29, 0x43,
	0x56, 0x61, 0x6c, 0x76, 0x65, 0x48, 0x57, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x16, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xbd, 0x01, 0x0a, 0x0d, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x82, 0x01, 0x0a,
	0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x72, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x47,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x27, 0x82, 0xb5, 0x18, 0x23, 0x44, 0x61, 0x74, 0x61, 0x20, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x20, 0x28, 0x44, 0x50, 0x29, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xc6, 0x01, 0x0a, 0x0d, 0x56,
	0x61, 0x6c, 0x76, 0x65, 0x48, 0x57, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0xa1, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x36, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x56, 0x61, 0x6c, 0x76, 0x65, 0x48, 0x57, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x5f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x56, 0x61, 0x6c, 0x76, 0x65,
	0x48, 0x57, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xb5, 0x18, 0x17, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20,
	0x49, 0x20, 0x72, 0x75, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x1a, 0x11, 0x82, 0xb5, 0x18, 0x0d, 0x56, 0x61, 0x6c, 0x76, 0x65, 0x48, 0x57, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x42, 0x3c, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x80, 0x01,
	0x01,
}

var (
	file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescData = file_steam_client_steammessages_datapublisher_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_datapublisher_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_steam_client_steammessages_datapublisher_steamclient_proto_goTypes = []interface{}{
	(*CDataPublisher_ClientContentCorruptionReport_Notification)(nil), // 0: steam.client.CDataPublisher_ClientContentCorruptionReport_Notification
	(*CValveHWSurvey_GetSurveySchedule_Request)(nil),                  // 1: steam.client.CValveHWSurvey_GetSurveySchedule_Request
	(*CValveHWSurvey_GetSurveySchedule_Response)(nil),                 // 2: steam.client.CValveHWSurvey_GetSurveySchedule_Response
	(*NoResponse)(nil), // 3: steam.client.NoResponse
}
var file_steam_client_steammessages_datapublisher_steamclient_proto_depIdxs = []int32{
	0, // 0: steam.client.DataPublisher.ClientContentCorruptionReport:input_type -> steam.client.CDataPublisher_ClientContentCorruptionReport_Notification
	1, // 1: steam.client.ValveHWSurvey.GetSurveySchedule:input_type -> steam.client.CValveHWSurvey_GetSurveySchedule_Request
	3, // 2: steam.client.DataPublisher.ClientContentCorruptionReport:output_type -> steam.client.NoResponse
	2, // 3: steam.client.ValveHWSurvey.GetSurveySchedule:output_type -> steam.client.CValveHWSurvey_GetSurveySchedule_Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_datapublisher_steamclient_proto_init() }
func file_steam_client_steammessages_datapublisher_steamclient_proto_init() {
	if File_steam_client_steammessages_datapublisher_steamclient_proto != nil {
		return
	}
	file_steam_client_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CDataPublisher_ClientContentCorruptionReport_Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CValveHWSurvey_GetSurveySchedule_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CValveHWSurvey_GetSurveySchedule_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steam_client_steammessages_datapublisher_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_steam_client_steammessages_datapublisher_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_datapublisher_steamclient_proto_depIdxs,
		MessageInfos:      file_steam_client_steammessages_datapublisher_steamclient_proto_msgTypes,
	}.Build()
	File_steam_client_steammessages_datapublisher_steamclient_proto = out.File
	file_steam_client_steammessages_datapublisher_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_datapublisher_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_datapublisher_steamclient_proto_depIdxs = nil
}
