// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: steam/client/steammessages_shader.steamclient.proto

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

type CShader_RegisterShader_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid      *uint32                                  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	GpuDesc    *string                                  `protobuf:"bytes,2,opt,name=gpu_desc,json=gpuDesc" json:"gpu_desc,omitempty"`
	DriverDesc *string                                  `protobuf:"bytes,3,opt,name=driver_desc,json=driverDesc" json:"driver_desc,omitempty"`
	Shaders    []*CShader_RegisterShader_Request_Shader `protobuf:"bytes,4,rep,name=shaders" json:"shaders,omitempty"`
}

func (x *CShader_RegisterShader_Request) Reset() {
	*x = CShader_RegisterShader_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_RegisterShader_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_RegisterShader_Request) ProtoMessage() {}

func (x *CShader_RegisterShader_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_RegisterShader_Request.ProtoReflect.Descriptor instead.
func (*CShader_RegisterShader_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CShader_RegisterShader_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CShader_RegisterShader_Request) GetGpuDesc() string {
	if x != nil && x.GpuDesc != nil {
		return *x.GpuDesc
	}
	return ""
}

func (x *CShader_RegisterShader_Request) GetDriverDesc() string {
	if x != nil && x.DriverDesc != nil {
		return *x.DriverDesc
	}
	return ""
}

func (x *CShader_RegisterShader_Request) GetShaders() []*CShader_RegisterShader_Request_Shader {
	if x != nil {
		return x.Shaders
	}
	return nil
}

type CShader_RegisterShader_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestedCodeids []uint32 `protobuf:"varint,1,rep,name=requested_codeids,json=requestedCodeids" json:"requested_codeids,omitempty"`
}

func (x *CShader_RegisterShader_Response) Reset() {
	*x = CShader_RegisterShader_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_RegisterShader_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_RegisterShader_Response) ProtoMessage() {}

func (x *CShader_RegisterShader_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_RegisterShader_Response.ProtoReflect.Descriptor instead.
func (*CShader_RegisterShader_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CShader_RegisterShader_Response) GetRequestedCodeids() []uint32 {
	if x != nil {
		return x.RequestedCodeids
	}
	return nil
}

type CShader_SendShader_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid   *uint32                                  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Shaders []*CShader_SendShader_Request_ShaderCode `protobuf:"bytes,2,rep,name=shaders" json:"shaders,omitempty"`
}

func (x *CShader_SendShader_Request) Reset() {
	*x = CShader_SendShader_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_SendShader_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_SendShader_Request) ProtoMessage() {}

func (x *CShader_SendShader_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_SendShader_Request.ProtoReflect.Descriptor instead.
func (*CShader_SendShader_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CShader_SendShader_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CShader_SendShader_Request) GetShaders() []*CShader_SendShader_Request_ShaderCode {
	if x != nil {
		return x.Shaders
	}
	return nil
}

type CShader_SendShader_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CShader_SendShader_Response) Reset() {
	*x = CShader_SendShader_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_SendShader_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_SendShader_Response) ProtoMessage() {}

func (x *CShader_SendShader_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_SendShader_Response.ProtoReflect.Descriptor instead.
func (*CShader_SendShader_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{3}
}

type CShader_GetBucketManifest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid      *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	GpuDesc    *string `protobuf:"bytes,2,opt,name=gpu_desc,json=gpuDesc" json:"gpu_desc,omitempty"`
	DriverDesc *string `protobuf:"bytes,3,opt,name=driver_desc,json=driverDesc" json:"driver_desc,omitempty"`
}

func (x *CShader_GetBucketManifest_Request) Reset() {
	*x = CShader_GetBucketManifest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_GetBucketManifest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_GetBucketManifest_Request) ProtoMessage() {}

func (x *CShader_GetBucketManifest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_GetBucketManifest_Request.ProtoReflect.Descriptor instead.
func (*CShader_GetBucketManifest_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{4}
}

func (x *CShader_GetBucketManifest_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CShader_GetBucketManifest_Request) GetGpuDesc() string {
	if x != nil && x.GpuDesc != nil {
		return *x.GpuDesc
	}
	return ""
}

func (x *CShader_GetBucketManifest_Request) GetDriverDesc() string {
	if x != nil && x.DriverDesc != nil {
		return *x.DriverDesc
	}
	return ""
}

type CShader_GetBucketManifest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manifestid *uint64 `protobuf:"varint,1,opt,name=manifestid" json:"manifestid,omitempty"`
	Depotsize  *uint32 `protobuf:"varint,2,opt,name=depotsize" json:"depotsize,omitempty"`
	Bucketid   *uint64 `protobuf:"varint,3,opt,name=bucketid" json:"bucketid,omitempty"`
}

func (x *CShader_GetBucketManifest_Response) Reset() {
	*x = CShader_GetBucketManifest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_GetBucketManifest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_GetBucketManifest_Response) ProtoMessage() {}

func (x *CShader_GetBucketManifest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_GetBucketManifest_Response.ProtoReflect.Descriptor instead.
func (*CShader_GetBucketManifest_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{5}
}

func (x *CShader_GetBucketManifest_Response) GetManifestid() uint64 {
	if x != nil && x.Manifestid != nil {
		return *x.Manifestid
	}
	return 0
}

func (x *CShader_GetBucketManifest_Response) GetDepotsize() uint32 {
	if x != nil && x.Depotsize != nil {
		return *x.Depotsize
	}
	return 0
}

func (x *CShader_GetBucketManifest_Response) GetBucketid() uint64 {
	if x != nil && x.Bucketid != nil {
		return *x.Bucketid
	}
	return 0
}

type CShader_RegisterShader_Request_Shader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheKeySha   []byte `protobuf:"bytes,1,opt,name=cache_key_sha,json=cacheKeySha" json:"cache_key_sha,omitempty"`
	ShaderCodeSha []byte `protobuf:"bytes,2,opt,name=shader_code_sha,json=shaderCodeSha" json:"shader_code_sha,omitempty"`
}

func (x *CShader_RegisterShader_Request_Shader) Reset() {
	*x = CShader_RegisterShader_Request_Shader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_RegisterShader_Request_Shader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_RegisterShader_Request_Shader) ProtoMessage() {}

func (x *CShader_RegisterShader_Request_Shader) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_RegisterShader_Request_Shader.ProtoReflect.Descriptor instead.
func (*CShader_RegisterShader_Request_Shader) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CShader_RegisterShader_Request_Shader) GetCacheKeySha() []byte {
	if x != nil {
		return x.CacheKeySha
	}
	return nil
}

func (x *CShader_RegisterShader_Request_Shader) GetShaderCodeSha() []byte {
	if x != nil {
		return x.ShaderCodeSha
	}
	return nil
}

type CShader_SendShader_Request_ShaderCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShaderCodeSha []byte `protobuf:"bytes,1,opt,name=shader_code_sha,json=shaderCodeSha" json:"shader_code_sha,omitempty"`
	ShaderCode    []byte `protobuf:"bytes,2,opt,name=shader_code,json=shaderCode" json:"shader_code,omitempty"`
}

func (x *CShader_SendShader_Request_ShaderCode) Reset() {
	*x = CShader_SendShader_Request_ShaderCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CShader_SendShader_Request_ShaderCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CShader_SendShader_Request_ShaderCode) ProtoMessage() {}

func (x *CShader_SendShader_Request_ShaderCode) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_shader_steamclient_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CShader_SendShader_Request_ShaderCode.ProtoReflect.Descriptor instead.
func (*CShader_SendShader_Request_ShaderCode) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CShader_SendShader_Request_ShaderCode) GetShaderCodeSha() []byte {
	if x != nil {
		return x.ShaderCodeSha
	}
	return nil
}

func (x *CShader_SendShader_Request_ShaderCode) GetShaderCode() []byte {
	if x != nil {
		return x.ShaderCode
	}
	return nil
}

var File_steam_client_steammessages_shader_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_shader_steamclient_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x68, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x1a, 0x39, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97,
	0x02, 0x0a, 0x1e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x70, 0x75, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x70, 0x75, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x4d, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x73, 0x68, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x1a, 0x54, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x53, 0x68, 0x61,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x68, 0x61, 0x22, 0x4e, 0x0a, 0x1f, 0x43, 0x53, 0x68, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x1a, 0x43, 0x53, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x4d, 0x0a,
	0x07, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53,
	0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x07, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x55, 0x0a, 0x0a,
	0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x68, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x75, 0x0a, 0x21, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x70, 0x75, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x70, 0x75, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x22, 0x7e, 0x0a, 0x22, 0x43, 0x53, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x64, 0x32, 0xed, 0x04, 0x0a, 0x06, 0x53, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x12, 0xd8, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x82, 0xb5, 0x18, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x20, 0x6a, 0x75, 0x73, 0x74, 0x20, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x20, 0x70,
	0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x20, 0x67, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x73, 0x68, 0x61, 0x64,
	0x65, 0x72, 0x20, 0x63, 0x61, 0x63, 0x68, 0x65, 0x20, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x69, 0x6e,
	0x67, 0x20, 0x75, 0x73, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x68, 0x65, 0x6d, 0x12,
	0xa9, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53,
	0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x46, 0x82, 0xb5, 0x18, 0x42, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20,
	0x69, 0x73, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x75, 0x73, 0x20, 0x61, 0x63,
	0x74, 0x75, 0x61, 0x6c, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x73, 0x68,
	0x61, 0x64, 0x65, 0x72, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x77,
	0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0xc7, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x12, 0x2f, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x53, 0x68, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xb5, 0x18, 0x4b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x6b, 0x6e, 0x6f, 0x77, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x20, 0x49, 0x44, 0x20, 0x74,
	0x6f, 0x20, 0x66, 0x65, 0x74, 0x63, 0x68, 0x20, 0x28, 0x69, 0x66, 0x20, 0x61, 0x6e, 0x79, 0x29,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x27, 0x73, 0x20,
	0x64, 0x65, 0x70, 0x6f, 0x74, 0x1a, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x53, 0x68, 0x61, 0x64, 0x65,
	0x72, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x42, 0x3c, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x80, 0x01, 0x01,
}

var (
	file_steam_client_steammessages_shader_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_shader_steamclient_proto_rawDescData = file_steam_client_steammessages_shader_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_shader_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_shader_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_shader_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_shader_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_shader_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_shader_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_steam_client_steammessages_shader_steamclient_proto_goTypes = []interface{}{
	(*CShader_RegisterShader_Request)(nil),        // 0: steam.client.CShader_RegisterShader_Request
	(*CShader_RegisterShader_Response)(nil),       // 1: steam.client.CShader_RegisterShader_Response
	(*CShader_SendShader_Request)(nil),            // 2: steam.client.CShader_SendShader_Request
	(*CShader_SendShader_Response)(nil),           // 3: steam.client.CShader_SendShader_Response
	(*CShader_GetBucketManifest_Request)(nil),     // 4: steam.client.CShader_GetBucketManifest_Request
	(*CShader_GetBucketManifest_Response)(nil),    // 5: steam.client.CShader_GetBucketManifest_Response
	(*CShader_RegisterShader_Request_Shader)(nil), // 6: steam.client.CShader_RegisterShader_Request.Shader
	(*CShader_SendShader_Request_ShaderCode)(nil), // 7: steam.client.CShader_SendShader_Request.ShaderCode
}
var file_steam_client_steammessages_shader_steamclient_proto_depIdxs = []int32{
	6, // 0: steam.client.CShader_RegisterShader_Request.shaders:type_name -> steam.client.CShader_RegisterShader_Request.Shader
	7, // 1: steam.client.CShader_SendShader_Request.shaders:type_name -> steam.client.CShader_SendShader_Request.ShaderCode
	0, // 2: steam.client.Shader.RegisterShader:input_type -> steam.client.CShader_RegisterShader_Request
	2, // 3: steam.client.Shader.SendShader:input_type -> steam.client.CShader_SendShader_Request
	4, // 4: steam.client.Shader.GetBucketManifest:input_type -> steam.client.CShader_GetBucketManifest_Request
	1, // 5: steam.client.Shader.RegisterShader:output_type -> steam.client.CShader_RegisterShader_Response
	3, // 6: steam.client.Shader.SendShader:output_type -> steam.client.CShader_SendShader_Response
	5, // 7: steam.client.Shader.GetBucketManifest:output_type -> steam.client.CShader_GetBucketManifest_Response
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_shader_steamclient_proto_init() }
func file_steam_client_steammessages_shader_steamclient_proto_init() {
	if File_steam_client_steammessages_shader_steamclient_proto != nil {
		return
	}
	file_steam_client_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_RegisterShader_Request); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_RegisterShader_Response); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_SendShader_Request); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_SendShader_Response); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_GetBucketManifest_Request); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_GetBucketManifest_Response); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_RegisterShader_Request_Shader); i {
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
		file_steam_client_steammessages_shader_steamclient_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CShader_SendShader_Request_ShaderCode); i {
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
			RawDescriptor: file_steam_client_steammessages_shader_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steam_client_steammessages_shader_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_shader_steamclient_proto_depIdxs,
		MessageInfos:      file_steam_client_steammessages_shader_steamclient_proto_msgTypes,
	}.Build()
	File_steam_client_steammessages_shader_steamclient_proto = out.File
	file_steam_client_steammessages_shader_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_shader_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_shader_steamclient_proto_depIdxs = nil
}
