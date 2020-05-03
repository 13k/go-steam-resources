// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: dota2/clientmessages.proto

package dota2

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

type EBaseClientMessages int32

const (
	EBaseClientMessages_CM_CustomGameEvent                  EBaseClientMessages = 280
	EBaseClientMessages_CM_CustomGameEventBounce            EBaseClientMessages = 281
	EBaseClientMessages_CM_ClientUIEvent                    EBaseClientMessages = 282
	EBaseClientMessages_CM_DevPaletteVisibilityChanged      EBaseClientMessages = 283
	EBaseClientMessages_CM_WorldUIControllerHasPanelChanged EBaseClientMessages = 284
	EBaseClientMessages_CM_RotateAnchor                     EBaseClientMessages = 285
	EBaseClientMessages_CM_MAX_BASE                         EBaseClientMessages = 300
)

// Enum value maps for EBaseClientMessages.
var (
	EBaseClientMessages_name = map[int32]string{
		280: "CM_CustomGameEvent",
		281: "CM_CustomGameEventBounce",
		282: "CM_ClientUIEvent",
		283: "CM_DevPaletteVisibilityChanged",
		284: "CM_WorldUIControllerHasPanelChanged",
		285: "CM_RotateAnchor",
		300: "CM_MAX_BASE",
	}
	EBaseClientMessages_value = map[string]int32{
		"CM_CustomGameEvent":                  280,
		"CM_CustomGameEventBounce":            281,
		"CM_ClientUIEvent":                    282,
		"CM_DevPaletteVisibilityChanged":      283,
		"CM_WorldUIControllerHasPanelChanged": 284,
		"CM_RotateAnchor":                     285,
		"CM_MAX_BASE":                         300,
	}
)

func (x EBaseClientMessages) Enum() *EBaseClientMessages {
	p := new(EBaseClientMessages)
	*p = x
	return p
}

func (x EBaseClientMessages) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EBaseClientMessages) Descriptor() protoreflect.EnumDescriptor {
	return file_dota2_clientmessages_proto_enumTypes[0].Descriptor()
}

func (EBaseClientMessages) Type() protoreflect.EnumType {
	return &file_dota2_clientmessages_proto_enumTypes[0]
}

func (x EBaseClientMessages) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EBaseClientMessages) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EBaseClientMessages(num)
	return nil
}

// Deprecated: Use EBaseClientMessages.Descriptor instead.
func (EBaseClientMessages) EnumDescriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{0}
}

type EClientUIEvent int32

const (
	EClientUIEvent_EClientUIEvent_Invalid        EClientUIEvent = 0
	EClientUIEvent_EClientUIEvent_DialogFinished EClientUIEvent = 1
	EClientUIEvent_EClientUIEvent_FireOutput     EClientUIEvent = 2
)

// Enum value maps for EClientUIEvent.
var (
	EClientUIEvent_name = map[int32]string{
		0: "EClientUIEvent_Invalid",
		1: "EClientUIEvent_DialogFinished",
		2: "EClientUIEvent_FireOutput",
	}
	EClientUIEvent_value = map[string]int32{
		"EClientUIEvent_Invalid":        0,
		"EClientUIEvent_DialogFinished": 1,
		"EClientUIEvent_FireOutput":     2,
	}
)

func (x EClientUIEvent) Enum() *EClientUIEvent {
	p := new(EClientUIEvent)
	*p = x
	return p
}

func (x EClientUIEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EClientUIEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_dota2_clientmessages_proto_enumTypes[1].Descriptor()
}

func (EClientUIEvent) Type() protoreflect.EnumType {
	return &file_dota2_clientmessages_proto_enumTypes[1]
}

func (x EClientUIEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EClientUIEvent) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EClientUIEvent(num)
	return nil
}

// Deprecated: Use EClientUIEvent.Descriptor instead.
func (EClientUIEvent) EnumDescriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{1}
}

type CClientMsg_CustomGameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventName *string `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data      []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (x *CClientMsg_CustomGameEvent) Reset() {
	*x = CClientMsg_CustomGameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_CustomGameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_CustomGameEvent) ProtoMessage() {}

func (x *CClientMsg_CustomGameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_CustomGameEvent.ProtoReflect.Descriptor instead.
func (*CClientMsg_CustomGameEvent) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{0}
}

func (x *CClientMsg_CustomGameEvent) GetEventName() string {
	if x != nil && x.EventName != nil {
		return *x.EventName
	}
	return ""
}

func (x *CClientMsg_CustomGameEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CClientMsg_CustomGameEventBounce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventName   *string `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	Data        []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	PlayerIndex *int32  `protobuf:"varint,3,opt,name=player_index,json=playerIndex" json:"player_index,omitempty"`
}

func (x *CClientMsg_CustomGameEventBounce) Reset() {
	*x = CClientMsg_CustomGameEventBounce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_CustomGameEventBounce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_CustomGameEventBounce) ProtoMessage() {}

func (x *CClientMsg_CustomGameEventBounce) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_CustomGameEventBounce.ProtoReflect.Descriptor instead.
func (*CClientMsg_CustomGameEventBounce) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{1}
}

func (x *CClientMsg_CustomGameEventBounce) GetEventName() string {
	if x != nil && x.EventName != nil {
		return *x.EventName
	}
	return ""
}

func (x *CClientMsg_CustomGameEventBounce) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CClientMsg_CustomGameEventBounce) GetPlayerIndex() int32 {
	if x != nil && x.PlayerIndex != nil {
		return *x.PlayerIndex
	}
	return 0
}

type CClientMsg_ClientUIEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event         *EClientUIEvent `protobuf:"varint,1,opt,name=event,enum=dota2.EClientUIEvent,def=0" json:"event,omitempty"`
	EntEhandle    *uint32         `protobuf:"varint,2,opt,name=ent_ehandle,json=entEhandle" json:"ent_ehandle,omitempty"`
	ClientEhandle *uint32         `protobuf:"varint,3,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	Data1         *string         `protobuf:"bytes,4,opt,name=data1" json:"data1,omitempty"`
	Data2         *string         `protobuf:"bytes,5,opt,name=data2" json:"data2,omitempty"`
}

// Default values for CClientMsg_ClientUIEvent fields.
const (
	Default_CClientMsg_ClientUIEvent_Event = EClientUIEvent_EClientUIEvent_Invalid
)

func (x *CClientMsg_ClientUIEvent) Reset() {
	*x = CClientMsg_ClientUIEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_ClientUIEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_ClientUIEvent) ProtoMessage() {}

func (x *CClientMsg_ClientUIEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_ClientUIEvent.ProtoReflect.Descriptor instead.
func (*CClientMsg_ClientUIEvent) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{2}
}

func (x *CClientMsg_ClientUIEvent) GetEvent() EClientUIEvent {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return Default_CClientMsg_ClientUIEvent_Event
}

func (x *CClientMsg_ClientUIEvent) GetEntEhandle() uint32 {
	if x != nil && x.EntEhandle != nil {
		return *x.EntEhandle
	}
	return 0
}

func (x *CClientMsg_ClientUIEvent) GetClientEhandle() uint32 {
	if x != nil && x.ClientEhandle != nil {
		return *x.ClientEhandle
	}
	return 0
}

func (x *CClientMsg_ClientUIEvent) GetData1() string {
	if x != nil && x.Data1 != nil {
		return *x.Data1
	}
	return ""
}

func (x *CClientMsg_ClientUIEvent) GetData2() string {
	if x != nil && x.Data2 != nil {
		return *x.Data2
	}
	return ""
}

type CClientMsg_DevPaletteVisibilityChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visible *bool `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
}

func (x *CClientMsg_DevPaletteVisibilityChangedEvent) Reset() {
	*x = CClientMsg_DevPaletteVisibilityChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_DevPaletteVisibilityChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_DevPaletteVisibilityChangedEvent) ProtoMessage() {}

func (x *CClientMsg_DevPaletteVisibilityChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_DevPaletteVisibilityChangedEvent.ProtoReflect.Descriptor instead.
func (*CClientMsg_DevPaletteVisibilityChangedEvent) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{3}
}

func (x *CClientMsg_DevPaletteVisibilityChangedEvent) GetVisible() bool {
	if x != nil && x.Visible != nil {
		return *x.Visible
	}
	return false
}

type CClientMsg_WorldUIControllerHasPanelChangedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasPanel        *bool   `protobuf:"varint,1,opt,name=has_panel,json=hasPanel" json:"has_panel,omitempty"`
	ClientEhandle   *uint32 `protobuf:"varint,2,opt,name=client_ehandle,json=clientEhandle" json:"client_ehandle,omitempty"`
	LiteralHandType *uint32 `protobuf:"varint,3,opt,name=literal_hand_type,json=literalHandType" json:"literal_hand_type,omitempty"`
}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) Reset() {
	*x = CClientMsg_WorldUIControllerHasPanelChangedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) ProtoMessage() {}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_WorldUIControllerHasPanelChangedEvent.ProtoReflect.Descriptor instead.
func (*CClientMsg_WorldUIControllerHasPanelChangedEvent) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{4}
}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetHasPanel() bool {
	if x != nil && x.HasPanel != nil {
		return *x.HasPanel
	}
	return false
}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetClientEhandle() uint32 {
	if x != nil && x.ClientEhandle != nil {
		return *x.ClientEhandle
	}
	return 0
}

func (x *CClientMsg_WorldUIControllerHasPanelChangedEvent) GetLiteralHandType() uint32 {
	if x != nil && x.LiteralHandType != nil {
		return *x.LiteralHandType
	}
	return 0
}

type CClientMsg_RotateAnchor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Angle *float32 `protobuf:"fixed32,1,opt,name=angle" json:"angle,omitempty"`
}

func (x *CClientMsg_RotateAnchor) Reset() {
	*x = CClientMsg_RotateAnchor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota2_clientmessages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientMsg_RotateAnchor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientMsg_RotateAnchor) ProtoMessage() {}

func (x *CClientMsg_RotateAnchor) ProtoReflect() protoreflect.Message {
	mi := &file_dota2_clientmessages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientMsg_RotateAnchor.ProtoReflect.Descriptor instead.
func (*CClientMsg_RotateAnchor) Descriptor() ([]byte, []int) {
	return file_dota2_clientmessages_proto_rawDescGZIP(), []int{5}
}

func (x *CClientMsg_RotateAnchor) GetAngle() float32 {
	if x != nil && x.Angle != nil {
		return *x.Angle
	}
	return 0
}

var File_dota2_clientmessages_proto protoreflect.FileDescriptor

var file_dota2_clientmessages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x6f,
	0x74, 0x61, 0x32, 0x22, 0x4f, 0x0a, 0x1a, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73,
	0x67, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x78, 0x0a, 0x20, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xd3,
	0x01, 0x0a, 0x18, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x74,
	0x61, 0x32, 0x2e, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x3a, 0x16, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x45, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x45, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61,
	0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x61, 0x74, 0x61, 0x32, 0x22, 0x47, 0x0a, 0x2b, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x5f, 0x44, 0x65, 0x76, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x56, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0xa2, 0x01,
	0x0a, 0x30, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x5f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x55, 0x49, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61,
	0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67,
	0x5f, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x6e,
	0x67, 0x6c, 0x65, 0x2a, 0xdb, 0x01, 0x0a, 0x13, 0x45, 0x42, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x12, 0x43,
	0x4d, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x10, 0x98, 0x02, 0x12, 0x1d, 0x0a, 0x18, 0x43, 0x4d, 0x5f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x10, 0x99, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x4d, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x55, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x9a, 0x02, 0x12, 0x23, 0x0a, 0x1e, 0x43, 0x4d,
	0x5f, 0x44, 0x65, 0x76, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x10, 0x9b, 0x02, 0x12,
	0x28, 0x0a, 0x23, 0x43, 0x4d, 0x5f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x55, 0x49, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x48, 0x61, 0x73, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x10, 0x9c, 0x02, 0x12, 0x14, 0x0a, 0x0f, 0x43, 0x4d, 0x5f,
	0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x10, 0x9d, 0x02, 0x12,
	0x10, 0x0a, 0x0b, 0x43, 0x4d, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x10, 0xac,
	0x02, 0x2a, 0x6e, 0x0a, 0x0e, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12,
	0x21, 0x0a, 0x1d, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x46, 0x69, 0x72, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x10,
	0x02, 0x42, 0x39, 0x48, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x80, 0x01, 0x00,
}

var (
	file_dota2_clientmessages_proto_rawDescOnce sync.Once
	file_dota2_clientmessages_proto_rawDescData = file_dota2_clientmessages_proto_rawDesc
)

func file_dota2_clientmessages_proto_rawDescGZIP() []byte {
	file_dota2_clientmessages_proto_rawDescOnce.Do(func() {
		file_dota2_clientmessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota2_clientmessages_proto_rawDescData)
	})
	return file_dota2_clientmessages_proto_rawDescData
}

var file_dota2_clientmessages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_dota2_clientmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dota2_clientmessages_proto_goTypes = []interface{}{
	(EBaseClientMessages)(0),                                 // 0: dota2.EBaseClientMessages
	(EClientUIEvent)(0),                                      // 1: dota2.EClientUIEvent
	(*CClientMsg_CustomGameEvent)(nil),                       // 2: dota2.CClientMsg_CustomGameEvent
	(*CClientMsg_CustomGameEventBounce)(nil),                 // 3: dota2.CClientMsg_CustomGameEventBounce
	(*CClientMsg_ClientUIEvent)(nil),                         // 4: dota2.CClientMsg_ClientUIEvent
	(*CClientMsg_DevPaletteVisibilityChangedEvent)(nil),      // 5: dota2.CClientMsg_DevPaletteVisibilityChangedEvent
	(*CClientMsg_WorldUIControllerHasPanelChangedEvent)(nil), // 6: dota2.CClientMsg_WorldUIControllerHasPanelChangedEvent
	(*CClientMsg_RotateAnchor)(nil),                          // 7: dota2.CClientMsg_RotateAnchor
}
var file_dota2_clientmessages_proto_depIdxs = []int32{
	1, // 0: dota2.CClientMsg_ClientUIEvent.event:type_name -> dota2.EClientUIEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dota2_clientmessages_proto_init() }
func file_dota2_clientmessages_proto_init() {
	if File_dota2_clientmessages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dota2_clientmessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_CustomGameEvent); i {
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
		file_dota2_clientmessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_CustomGameEventBounce); i {
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
		file_dota2_clientmessages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_ClientUIEvent); i {
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
		file_dota2_clientmessages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_DevPaletteVisibilityChangedEvent); i {
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
		file_dota2_clientmessages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_WorldUIControllerHasPanelChangedEvent); i {
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
		file_dota2_clientmessages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CClientMsg_RotateAnchor); i {
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
			RawDescriptor: file_dota2_clientmessages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota2_clientmessages_proto_goTypes,
		DependencyIndexes: file_dota2_clientmessages_proto_depIdxs,
		EnumInfos:         file_dota2_clientmessages_proto_enumTypes,
		MessageInfos:      file_dota2_clientmessages_proto_msgTypes,
	}.Build()
	File_dota2_clientmessages_proto = out.File
	file_dota2_clientmessages_proto_rawDesc = nil
	file_dota2_clientmessages_proto_goTypes = nil
	file_dota2_clientmessages_proto_depIdxs = nil
}
