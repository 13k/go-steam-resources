// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: steam/client/steammessages_contentsystem.steamclient.proto

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

type CContentServerDirectory_GetServersForSteamPipe_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellId       *uint32 `protobuf:"varint,1,opt,name=cell_id,json=cellId" json:"cell_id,omitempty"`
	MaxServers   *uint32 `protobuf:"varint,2,opt,name=max_servers,json=maxServers,def=20" json:"max_servers,omitempty"`
	IpOverride   *string `protobuf:"bytes,3,opt,name=ip_override,json=ipOverride" json:"ip_override,omitempty"`
	LauncherType *int32  `protobuf:"varint,4,opt,name=launcher_type,json=launcherType,def=0" json:"launcher_type,omitempty"`
}

// Default values for CContentServerDirectory_GetServersForSteamPipe_Request fields.
const (
	Default_CContentServerDirectory_GetServersForSteamPipe_Request_MaxServers   = uint32(20)
	Default_CContentServerDirectory_GetServersForSteamPipe_Request_LauncherType = int32(0)
)

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) Reset() {
	*x = CContentServerDirectory_GetServersForSteamPipe_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CContentServerDirectory_GetServersForSteamPipe_Request) ProtoMessage() {}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CContentServerDirectory_GetServersForSteamPipe_Request.ProtoReflect.Descriptor instead.
func (*CContentServerDirectory_GetServersForSteamPipe_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) GetCellId() uint32 {
	if x != nil && x.CellId != nil {
		return *x.CellId
	}
	return 0
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) GetMaxServers() uint32 {
	if x != nil && x.MaxServers != nil {
		return *x.MaxServers
	}
	return Default_CContentServerDirectory_GetServersForSteamPipe_Request_MaxServers
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) GetIpOverride() string {
	if x != nil && x.IpOverride != nil {
		return *x.IpOverride
	}
	return ""
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Request) GetLauncherType() int32 {
	if x != nil && x.LauncherType != nil {
		return *x.LauncherType
	}
	return Default_CContentServerDirectory_GetServersForSteamPipe_Request_LauncherType
}

type CContentServerDirectory_ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                     *string  `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	SourceId                 *int32   `protobuf:"varint,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	CellId                   *int32   `protobuf:"varint,3,opt,name=cell_id,json=cellId" json:"cell_id,omitempty"`
	Load                     *int32   `protobuf:"varint,4,opt,name=load" json:"load,omitempty"`
	WeightedLoad             *float32 `protobuf:"fixed32,5,opt,name=weighted_load,json=weightedLoad" json:"weighted_load,omitempty"`
	NumEntriesInClientList   *int32   `protobuf:"varint,6,opt,name=num_entries_in_client_list,json=numEntriesInClientList" json:"num_entries_in_client_list,omitempty"`
	SteamChinaOnly           *bool    `protobuf:"varint,7,opt,name=steam_china_only,json=steamChinaOnly" json:"steam_china_only,omitempty"`
	Host                     *string  `protobuf:"bytes,8,opt,name=host" json:"host,omitempty"`
	Vhost                    *string  `protobuf:"bytes,9,opt,name=vhost" json:"vhost,omitempty"`
	UseAsProxy               *bool    `protobuf:"varint,10,opt,name=use_as_proxy,json=useAsProxy" json:"use_as_proxy,omitempty"`
	ProxyRequestPathTemplate *string  `protobuf:"bytes,11,opt,name=proxy_request_path_template,json=proxyRequestPathTemplate" json:"proxy_request_path_template,omitempty"`
	HttpsSupport             *string  `protobuf:"bytes,12,opt,name=https_support,json=httpsSupport" json:"https_support,omitempty"`
	AllowedAppIds            []uint32 `protobuf:"varint,13,rep,name=allowed_app_ids,json=allowedAppIds" json:"allowed_app_ids,omitempty"`
	PreferredServer          *bool    `protobuf:"varint,14,opt,name=preferred_server,json=preferredServer" json:"preferred_server,omitempty"`
}

func (x *CContentServerDirectory_ServerInfo) Reset() {
	*x = CContentServerDirectory_ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CContentServerDirectory_ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CContentServerDirectory_ServerInfo) ProtoMessage() {}

func (x *CContentServerDirectory_ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CContentServerDirectory_ServerInfo.ProtoReflect.Descriptor instead.
func (*CContentServerDirectory_ServerInfo) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CContentServerDirectory_ServerInfo) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *CContentServerDirectory_ServerInfo) GetSourceId() int32 {
	if x != nil && x.SourceId != nil {
		return *x.SourceId
	}
	return 0
}

func (x *CContentServerDirectory_ServerInfo) GetCellId() int32 {
	if x != nil && x.CellId != nil {
		return *x.CellId
	}
	return 0
}

func (x *CContentServerDirectory_ServerInfo) GetLoad() int32 {
	if x != nil && x.Load != nil {
		return *x.Load
	}
	return 0
}

func (x *CContentServerDirectory_ServerInfo) GetWeightedLoad() float32 {
	if x != nil && x.WeightedLoad != nil {
		return *x.WeightedLoad
	}
	return 0
}

func (x *CContentServerDirectory_ServerInfo) GetNumEntriesInClientList() int32 {
	if x != nil && x.NumEntriesInClientList != nil {
		return *x.NumEntriesInClientList
	}
	return 0
}

func (x *CContentServerDirectory_ServerInfo) GetSteamChinaOnly() bool {
	if x != nil && x.SteamChinaOnly != nil {
		return *x.SteamChinaOnly
	}
	return false
}

func (x *CContentServerDirectory_ServerInfo) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

func (x *CContentServerDirectory_ServerInfo) GetVhost() string {
	if x != nil && x.Vhost != nil {
		return *x.Vhost
	}
	return ""
}

func (x *CContentServerDirectory_ServerInfo) GetUseAsProxy() bool {
	if x != nil && x.UseAsProxy != nil {
		return *x.UseAsProxy
	}
	return false
}

func (x *CContentServerDirectory_ServerInfo) GetProxyRequestPathTemplate() string {
	if x != nil && x.ProxyRequestPathTemplate != nil {
		return *x.ProxyRequestPathTemplate
	}
	return ""
}

func (x *CContentServerDirectory_ServerInfo) GetHttpsSupport() string {
	if x != nil && x.HttpsSupport != nil {
		return *x.HttpsSupport
	}
	return ""
}

func (x *CContentServerDirectory_ServerInfo) GetAllowedAppIds() []uint32 {
	if x != nil {
		return x.AllowedAppIds
	}
	return nil
}

func (x *CContentServerDirectory_ServerInfo) GetPreferredServer() bool {
	if x != nil && x.PreferredServer != nil {
		return *x.PreferredServer
	}
	return false
}

type CContentServerDirectory_GetServersForSteamPipe_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*CContentServerDirectory_ServerInfo `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Response) Reset() {
	*x = CContentServerDirectory_GetServersForSteamPipe_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CContentServerDirectory_GetServersForSteamPipe_Response) ProtoMessage() {}

func (x *CContentServerDirectory_GetServersForSteamPipe_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CContentServerDirectory_GetServersForSteamPipe_Response.ProtoReflect.Descriptor instead.
func (*CContentServerDirectory_GetServersForSteamPipe_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CContentServerDirectory_GetServersForSteamPipe_Response) GetServers() []*CContentServerDirectory_ServerInfo {
	if x != nil {
		return x.Servers
	}
	return nil
}

type CContentServerDirectory_GetDepotPatchInfo_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Depotid          *uint32 `protobuf:"varint,2,opt,name=depotid" json:"depotid,omitempty"`
	SourceManifestid *uint64 `protobuf:"varint,3,opt,name=source_manifestid,json=sourceManifestid" json:"source_manifestid,omitempty"`
	TargetManifestid *uint64 `protobuf:"varint,4,opt,name=target_manifestid,json=targetManifestid" json:"target_manifestid,omitempty"`
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) Reset() {
	*x = CContentServerDirectory_GetDepotPatchInfo_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CContentServerDirectory_GetDepotPatchInfo_Request) ProtoMessage() {}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CContentServerDirectory_GetDepotPatchInfo_Request.ProtoReflect.Descriptor instead.
func (*CContentServerDirectory_GetDepotPatchInfo_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP(), []int{3}
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) GetDepotid() uint32 {
	if x != nil && x.Depotid != nil {
		return *x.Depotid
	}
	return 0
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) GetSourceManifestid() uint64 {
	if x != nil && x.SourceManifestid != nil {
		return *x.SourceManifestid
	}
	return 0
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Request) GetTargetManifestid() uint64 {
	if x != nil && x.TargetManifestid != nil {
		return *x.TargetManifestid
	}
	return 0
}

type CContentServerDirectory_GetDepotPatchInfo_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvailable *bool `protobuf:"varint,1,opt,name=is_available,json=isAvailable" json:"is_available,omitempty"`
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Response) Reset() {
	*x = CContentServerDirectory_GetDepotPatchInfo_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CContentServerDirectory_GetDepotPatchInfo_Response) ProtoMessage() {}

func (x *CContentServerDirectory_GetDepotPatchInfo_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CContentServerDirectory_GetDepotPatchInfo_Response.ProtoReflect.Descriptor instead.
func (*CContentServerDirectory_GetDepotPatchInfo_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP(), []int{4}
}

func (x *CContentServerDirectory_GetDepotPatchInfo_Response) GetIsAvailable() bool {
	if x != nil && x.IsAvailable != nil {
		return *x.IsAvailable
	}
	return false
}

var File_steam_client_steammessages_contentsystem_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_contentsystem_steamclient_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x39, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x36, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x50, 0x69, 0x70, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x43, 0x65,
	0x6c, 0x6c, 0x20, 0x49, 0x44, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x45, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x3a, 0x02, 0x32, 0x30, 0x42, 0x20, 0x82, 0xb5, 0x18, 0x1c, 0x6d, 0x61, 0x78, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x69, 0x70, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x82, 0xb5, 0x18, 0x11, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0a, 0x69, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x42, 0x11, 0x82, 0xb5, 0x18, 0x0d, 0x6c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x20, 0x74, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x90, 0x04, 0x0a, 0x22, 0x43, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x61,
	0x64, 0x12, 0x3a, 0x0a, 0x1a, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x49, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x68, 0x69, 0x6e, 0x61, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x43, 0x68,
	0x69, 0x6e, 0x61, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x41, 0x73, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x70, 0x70, 0x49, 0x64, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x85, 0x01, 0x0a, 0x37, 0x43,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x50, 0x69, 0x70, 0x65, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x31, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x69, 0x64, 0x22, 0x57, 0x0a, 0x32, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x32, 0xff, 0x02, 0x0a, 0x16,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0xa5, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x50, 0x69, 0x70,
	0x65, 0x12, 0x44, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x50, 0x69, 0x70, 0x65, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x50, 0x69, 0x70, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x96,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x24, 0x82, 0xb5, 0x18, 0x20, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x43, 0x44, 0x4e, 0x20, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x3b, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x80, 0x01, 0x01,
}

var (
	file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescData = file_steam_client_steammessages_contentsystem_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_contentsystem_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_steam_client_steammessages_contentsystem_steamclient_proto_goTypes = []interface{}{
	(*CContentServerDirectory_GetServersForSteamPipe_Request)(nil),  // 0: steam.client.CContentServerDirectory_GetServersForSteamPipe_Request
	(*CContentServerDirectory_ServerInfo)(nil),                      // 1: steam.client.CContentServerDirectory_ServerInfo
	(*CContentServerDirectory_GetServersForSteamPipe_Response)(nil), // 2: steam.client.CContentServerDirectory_GetServersForSteamPipe_Response
	(*CContentServerDirectory_GetDepotPatchInfo_Request)(nil),       // 3: steam.client.CContentServerDirectory_GetDepotPatchInfo_Request
	(*CContentServerDirectory_GetDepotPatchInfo_Response)(nil),      // 4: steam.client.CContentServerDirectory_GetDepotPatchInfo_Response
}
var file_steam_client_steammessages_contentsystem_steamclient_proto_depIdxs = []int32{
	1, // 0: steam.client.CContentServerDirectory_GetServersForSteamPipe_Response.servers:type_name -> steam.client.CContentServerDirectory_ServerInfo
	0, // 1: steam.client.ContentServerDirectory.GetServersForSteamPipe:input_type -> steam.client.CContentServerDirectory_GetServersForSteamPipe_Request
	3, // 2: steam.client.ContentServerDirectory.GetDepotPatchInfo:input_type -> steam.client.CContentServerDirectory_GetDepotPatchInfo_Request
	2, // 3: steam.client.ContentServerDirectory.GetServersForSteamPipe:output_type -> steam.client.CContentServerDirectory_GetServersForSteamPipe_Response
	4, // 4: steam.client.ContentServerDirectory.GetDepotPatchInfo:output_type -> steam.client.CContentServerDirectory_GetDepotPatchInfo_Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_contentsystem_steamclient_proto_init() }
func file_steam_client_steammessages_contentsystem_steamclient_proto_init() {
	if File_steam_client_steammessages_contentsystem_steamclient_proto != nil {
		return
	}
	file_steam_client_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CContentServerDirectory_GetServersForSteamPipe_Request); i {
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
		file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CContentServerDirectory_ServerInfo); i {
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
		file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CContentServerDirectory_GetServersForSteamPipe_Response); i {
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
		file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CContentServerDirectory_GetDepotPatchInfo_Request); i {
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
		file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CContentServerDirectory_GetDepotPatchInfo_Response); i {
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
			RawDescriptor: file_steam_client_steammessages_contentsystem_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steam_client_steammessages_contentsystem_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_contentsystem_steamclient_proto_depIdxs,
		MessageInfos:      file_steam_client_steammessages_contentsystem_steamclient_proto_msgTypes,
	}.Build()
	File_steam_client_steammessages_contentsystem_steamclient_proto = out.File
	file_steam_client_steammessages_contentsystem_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_contentsystem_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_contentsystem_steamclient_proto_depIdxs = nil
}