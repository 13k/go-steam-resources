// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gametoolevents.proto

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

type ChangeMapToolEvent struct {
	Mapname              *string  `protobuf:"bytes,1,opt,name=mapname" json:"mapname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeMapToolEvent) Reset()         { *m = ChangeMapToolEvent{} }
func (m *ChangeMapToolEvent) String() string { return proto.CompactTextString(m) }
func (*ChangeMapToolEvent) ProtoMessage()    {}
func (*ChangeMapToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{0}
}
func (m *ChangeMapToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeMapToolEvent.Unmarshal(m, b)
}
func (m *ChangeMapToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeMapToolEvent.Marshal(b, m, deterministic)
}
func (dst *ChangeMapToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeMapToolEvent.Merge(dst, src)
}
func (m *ChangeMapToolEvent) XXX_Size() int {
	return xxx_messageInfo_ChangeMapToolEvent.Size(m)
}
func (m *ChangeMapToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeMapToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeMapToolEvent proto.InternalMessageInfo

func (m *ChangeMapToolEvent) GetMapname() string {
	if m != nil && m.Mapname != nil {
		return *m.Mapname
	}
	return ""
}

type TraceRayServerToolEvent struct {
	Start                *CMsgVector `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End                  *CMsgVector `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TraceRayServerToolEvent) Reset()         { *m = TraceRayServerToolEvent{} }
func (m *TraceRayServerToolEvent) String() string { return proto.CompactTextString(m) }
func (*TraceRayServerToolEvent) ProtoMessage()    {}
func (*TraceRayServerToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{1}
}
func (m *TraceRayServerToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceRayServerToolEvent.Unmarshal(m, b)
}
func (m *TraceRayServerToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceRayServerToolEvent.Marshal(b, m, deterministic)
}
func (dst *TraceRayServerToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceRayServerToolEvent.Merge(dst, src)
}
func (m *TraceRayServerToolEvent) XXX_Size() int {
	return xxx_messageInfo_TraceRayServerToolEvent.Size(m)
}
func (m *TraceRayServerToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceRayServerToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TraceRayServerToolEvent proto.InternalMessageInfo

func (m *TraceRayServerToolEvent) GetStart() *CMsgVector {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TraceRayServerToolEvent) GetEnd() *CMsgVector {
	if m != nil {
		return m.End
	}
	return nil
}

type ToolTraceRayResult struct {
	Hit                  *bool       `protobuf:"varint,1,opt,name=hit" json:"hit,omitempty"`
	Impact               *CMsgVector `protobuf:"bytes,2,opt,name=impact" json:"impact,omitempty"`
	Normal               *CMsgVector `protobuf:"bytes,3,opt,name=normal" json:"normal,omitempty"`
	Distance             *float32    `protobuf:"fixed32,4,opt,name=distance" json:"distance,omitempty"`
	Fraction             *float32    `protobuf:"fixed32,5,opt,name=fraction" json:"fraction,omitempty"`
	Ehandle              *int32      `protobuf:"varint,6,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ToolTraceRayResult) Reset()         { *m = ToolTraceRayResult{} }
func (m *ToolTraceRayResult) String() string { return proto.CompactTextString(m) }
func (*ToolTraceRayResult) ProtoMessage()    {}
func (*ToolTraceRayResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{2}
}
func (m *ToolTraceRayResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolTraceRayResult.Unmarshal(m, b)
}
func (m *ToolTraceRayResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolTraceRayResult.Marshal(b, m, deterministic)
}
func (dst *ToolTraceRayResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolTraceRayResult.Merge(dst, src)
}
func (m *ToolTraceRayResult) XXX_Size() int {
	return xxx_messageInfo_ToolTraceRayResult.Size(m)
}
func (m *ToolTraceRayResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolTraceRayResult.DiscardUnknown(m)
}

var xxx_messageInfo_ToolTraceRayResult proto.InternalMessageInfo

func (m *ToolTraceRayResult) GetHit() bool {
	if m != nil && m.Hit != nil {
		return *m.Hit
	}
	return false
}

func (m *ToolTraceRayResult) GetImpact() *CMsgVector {
	if m != nil {
		return m.Impact
	}
	return nil
}

func (m *ToolTraceRayResult) GetNormal() *CMsgVector {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *ToolTraceRayResult) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

func (m *ToolTraceRayResult) GetFraction() float32 {
	if m != nil && m.Fraction != nil {
		return *m.Fraction
	}
	return 0
}

func (m *ToolTraceRayResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type SpawnEntityToolEvent struct {
	EntityKeyvalues      []byte   `protobuf:"bytes,1,opt,name=entity_keyvalues,json=entityKeyvalues" json:"entity_keyvalues,omitempty"`
	Clientsideentity     *bool    `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnEntityToolEvent) Reset()         { *m = SpawnEntityToolEvent{} }
func (m *SpawnEntityToolEvent) String() string { return proto.CompactTextString(m) }
func (*SpawnEntityToolEvent) ProtoMessage()    {}
func (*SpawnEntityToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{3}
}
func (m *SpawnEntityToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnEntityToolEvent.Unmarshal(m, b)
}
func (m *SpawnEntityToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnEntityToolEvent.Marshal(b, m, deterministic)
}
func (dst *SpawnEntityToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnEntityToolEvent.Merge(dst, src)
}
func (m *SpawnEntityToolEvent) XXX_Size() int {
	return xxx_messageInfo_SpawnEntityToolEvent.Size(m)
}
func (m *SpawnEntityToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnEntityToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnEntityToolEvent proto.InternalMessageInfo

func (m *SpawnEntityToolEvent) GetEntityKeyvalues() []byte {
	if m != nil {
		return m.EntityKeyvalues
	}
	return nil
}

func (m *SpawnEntityToolEvent) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type SpawnEntityToolEventResult struct {
	Ehandle              *int32   `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnEntityToolEventResult) Reset()         { *m = SpawnEntityToolEventResult{} }
func (m *SpawnEntityToolEventResult) String() string { return proto.CompactTextString(m) }
func (*SpawnEntityToolEventResult) ProtoMessage()    {}
func (*SpawnEntityToolEventResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{4}
}
func (m *SpawnEntityToolEventResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnEntityToolEventResult.Unmarshal(m, b)
}
func (m *SpawnEntityToolEventResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnEntityToolEventResult.Marshal(b, m, deterministic)
}
func (dst *SpawnEntityToolEventResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnEntityToolEventResult.Merge(dst, src)
}
func (m *SpawnEntityToolEventResult) XXX_Size() int {
	return xxx_messageInfo_SpawnEntityToolEventResult.Size(m)
}
func (m *SpawnEntityToolEventResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnEntityToolEventResult.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnEntityToolEventResult proto.InternalMessageInfo

func (m *SpawnEntityToolEventResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyEntityToolEvent struct {
	Ehandle              *int32   `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyEntityToolEvent) Reset()         { *m = DestroyEntityToolEvent{} }
func (m *DestroyEntityToolEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyEntityToolEvent) ProtoMessage()    {}
func (*DestroyEntityToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{5}
}
func (m *DestroyEntityToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyEntityToolEvent.Unmarshal(m, b)
}
func (m *DestroyEntityToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyEntityToolEvent.Marshal(b, m, deterministic)
}
func (dst *DestroyEntityToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyEntityToolEvent.Merge(dst, src)
}
func (m *DestroyEntityToolEvent) XXX_Size() int {
	return xxx_messageInfo_DestroyEntityToolEvent.Size(m)
}
func (m *DestroyEntityToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyEntityToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyEntityToolEvent proto.InternalMessageInfo

func (m *DestroyEntityToolEvent) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyAllEntitiesToolEvent struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyAllEntitiesToolEvent) Reset()         { *m = DestroyAllEntitiesToolEvent{} }
func (m *DestroyAllEntitiesToolEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyAllEntitiesToolEvent) ProtoMessage()    {}
func (*DestroyAllEntitiesToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{6}
}
func (m *DestroyAllEntitiesToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyAllEntitiesToolEvent.Unmarshal(m, b)
}
func (m *DestroyAllEntitiesToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyAllEntitiesToolEvent.Marshal(b, m, deterministic)
}
func (dst *DestroyAllEntitiesToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyAllEntitiesToolEvent.Merge(dst, src)
}
func (m *DestroyAllEntitiesToolEvent) XXX_Size() int {
	return xxx_messageInfo_DestroyAllEntitiesToolEvent.Size(m)
}
func (m *DestroyAllEntitiesToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyAllEntitiesToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyAllEntitiesToolEvent proto.InternalMessageInfo

type RestartMapToolEvent struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartMapToolEvent) Reset()         { *m = RestartMapToolEvent{} }
func (m *RestartMapToolEvent) String() string { return proto.CompactTextString(m) }
func (*RestartMapToolEvent) ProtoMessage()    {}
func (*RestartMapToolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{7}
}
func (m *RestartMapToolEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartMapToolEvent.Unmarshal(m, b)
}
func (m *RestartMapToolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartMapToolEvent.Marshal(b, m, deterministic)
}
func (dst *RestartMapToolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartMapToolEvent.Merge(dst, src)
}
func (m *RestartMapToolEvent) XXX_Size() int {
	return xxx_messageInfo_RestartMapToolEvent.Size(m)
}
func (m *RestartMapToolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartMapToolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RestartMapToolEvent proto.InternalMessageInfo

type ToolEvent_GetEntityInfo struct {
	Ehandle              *int32   `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity     *bool    `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_GetEntityInfo) Reset()         { *m = ToolEvent_GetEntityInfo{} }
func (m *ToolEvent_GetEntityInfo) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfo) ProtoMessage()    {}
func (*ToolEvent_GetEntityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{8}
}
func (m *ToolEvent_GetEntityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_GetEntityInfo.Unmarshal(m, b)
}
func (m *ToolEvent_GetEntityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_GetEntityInfo.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_GetEntityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_GetEntityInfo.Merge(dst, src)
}
func (m *ToolEvent_GetEntityInfo) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_GetEntityInfo.Size(m)
}
func (m *ToolEvent_GetEntityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_GetEntityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_GetEntityInfo proto.InternalMessageInfo

func (m *ToolEvent_GetEntityInfo) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInfo) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInfoResult struct {
	Cppclass             *string     `protobuf:"bytes,1,opt,name=cppclass,def=shithead" json:"cppclass,omitempty"`
	Classname            *string     `protobuf:"bytes,2,opt,name=classname" json:"classname,omitempty"`
	Name                 *string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Origin               *CMsgVector `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Mins                 *CMsgVector `protobuf:"bytes,5,opt,name=mins" json:"mins,omitempty"`
	Maxs                 *CMsgVector `protobuf:"bytes,6,opt,name=maxs" json:"maxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ToolEvent_GetEntityInfoResult) Reset()         { *m = ToolEvent_GetEntityInfoResult{} }
func (m *ToolEvent_GetEntityInfoResult) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfoResult) ProtoMessage()    {}
func (*ToolEvent_GetEntityInfoResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{9}
}
func (m *ToolEvent_GetEntityInfoResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_GetEntityInfoResult.Unmarshal(m, b)
}
func (m *ToolEvent_GetEntityInfoResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_GetEntityInfoResult.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_GetEntityInfoResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_GetEntityInfoResult.Merge(dst, src)
}
func (m *ToolEvent_GetEntityInfoResult) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_GetEntityInfoResult.Size(m)
}
func (m *ToolEvent_GetEntityInfoResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_GetEntityInfoResult.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_GetEntityInfoResult proto.InternalMessageInfo

const Default_ToolEvent_GetEntityInfoResult_Cppclass string = "shithead"

func (m *ToolEvent_GetEntityInfoResult) GetCppclass() string {
	if m != nil && m.Cppclass != nil {
		return *m.Cppclass
	}
	return Default_ToolEvent_GetEntityInfoResult_Cppclass
}

func (m *ToolEvent_GetEntityInfoResult) GetClassname() string {
	if m != nil && m.Classname != nil {
		return *m.Classname
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetOrigin() *CMsgVector {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMins() *CMsgVector {
	if m != nil {
		return m.Mins
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMaxs() *CMsgVector {
	if m != nil {
		return m.Maxs
	}
	return nil
}

type ToolEvent_GetEntityInputs struct {
	Ehandle              *int32   `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity     *bool    `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_GetEntityInputs) Reset()         { *m = ToolEvent_GetEntityInputs{} }
func (m *ToolEvent_GetEntityInputs) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputs) ProtoMessage()    {}
func (*ToolEvent_GetEntityInputs) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{10}
}
func (m *ToolEvent_GetEntityInputs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_GetEntityInputs.Unmarshal(m, b)
}
func (m *ToolEvent_GetEntityInputs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_GetEntityInputs.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_GetEntityInputs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_GetEntityInputs.Merge(dst, src)
}
func (m *ToolEvent_GetEntityInputs) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_GetEntityInputs.Size(m)
}
func (m *ToolEvent_GetEntityInputs) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_GetEntityInputs.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_GetEntityInputs proto.InternalMessageInfo

func (m *ToolEvent_GetEntityInputs) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInputs) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInputsResult struct {
	InputList            []string `protobuf:"bytes,1,rep,name=input_list,json=inputList" json:"input_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_GetEntityInputsResult) Reset()         { *m = ToolEvent_GetEntityInputsResult{} }
func (m *ToolEvent_GetEntityInputsResult) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputsResult) ProtoMessage()    {}
func (*ToolEvent_GetEntityInputsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{11}
}
func (m *ToolEvent_GetEntityInputsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_GetEntityInputsResult.Unmarshal(m, b)
}
func (m *ToolEvent_GetEntityInputsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_GetEntityInputsResult.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_GetEntityInputsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_GetEntityInputsResult.Merge(dst, src)
}
func (m *ToolEvent_GetEntityInputsResult) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_GetEntityInputsResult.Size(m)
}
func (m *ToolEvent_GetEntityInputsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_GetEntityInputsResult.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_GetEntityInputsResult proto.InternalMessageInfo

func (m *ToolEvent_GetEntityInputsResult) GetInputList() []string {
	if m != nil {
		return m.InputList
	}
	return nil
}

type ToolEvent_FireEntityInput struct {
	Ehandle              *int32   `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity     *bool    `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	InputName            *string  `protobuf:"bytes,3,opt,name=input_name,json=inputName" json:"input_name,omitempty"`
	InputParam           *string  `protobuf:"bytes,4,opt,name=input_param,json=inputParam" json:"input_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_FireEntityInput) Reset()         { *m = ToolEvent_FireEntityInput{} }
func (m *ToolEvent_FireEntityInput) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_FireEntityInput) ProtoMessage()    {}
func (*ToolEvent_FireEntityInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{12}
}
func (m *ToolEvent_FireEntityInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_FireEntityInput.Unmarshal(m, b)
}
func (m *ToolEvent_FireEntityInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_FireEntityInput.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_FireEntityInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_FireEntityInput.Merge(dst, src)
}
func (m *ToolEvent_FireEntityInput) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_FireEntityInput.Size(m)
}
func (m *ToolEvent_FireEntityInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_FireEntityInput.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_FireEntityInput proto.InternalMessageInfo

func (m *ToolEvent_FireEntityInput) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_FireEntityInput) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

func (m *ToolEvent_FireEntityInput) GetInputName() string {
	if m != nil && m.InputName != nil {
		return *m.InputName
	}
	return ""
}

func (m *ToolEvent_FireEntityInput) GetInputParam() string {
	if m != nil && m.InputParam != nil {
		return *m.InputParam
	}
	return ""
}

type ToolEvent_SFMRecordingStateChanged struct {
	Isrecording          *bool    `protobuf:"varint,1,opt,name=isrecording" json:"isrecording,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_SFMRecordingStateChanged) Reset()         { *m = ToolEvent_SFMRecordingStateChanged{} }
func (m *ToolEvent_SFMRecordingStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMRecordingStateChanged) ProtoMessage()    {}
func (*ToolEvent_SFMRecordingStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{13}
}
func (m *ToolEvent_SFMRecordingStateChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_SFMRecordingStateChanged.Unmarshal(m, b)
}
func (m *ToolEvent_SFMRecordingStateChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_SFMRecordingStateChanged.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_SFMRecordingStateChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_SFMRecordingStateChanged.Merge(dst, src)
}
func (m *ToolEvent_SFMRecordingStateChanged) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_SFMRecordingStateChanged.Size(m)
}
func (m *ToolEvent_SFMRecordingStateChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_SFMRecordingStateChanged.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_SFMRecordingStateChanged proto.InternalMessageInfo

func (m *ToolEvent_SFMRecordingStateChanged) GetIsrecording() bool {
	if m != nil && m.Isrecording != nil {
		return *m.Isrecording
	}
	return false
}

type ToolEvent_SFMToolActiveStateChanged struct {
	Isactive             *bool    `protobuf:"varint,1,opt,name=isactive" json:"isactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToolEvent_SFMToolActiveStateChanged) Reset()         { *m = ToolEvent_SFMToolActiveStateChanged{} }
func (m *ToolEvent_SFMToolActiveStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMToolActiveStateChanged) ProtoMessage()    {}
func (*ToolEvent_SFMToolActiveStateChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_gametoolevents_115765871d01f2d3, []int{14}
}
func (m *ToolEvent_SFMToolActiveStateChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged.Unmarshal(m, b)
}
func (m *ToolEvent_SFMToolActiveStateChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged.Marshal(b, m, deterministic)
}
func (dst *ToolEvent_SFMToolActiveStateChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged.Merge(dst, src)
}
func (m *ToolEvent_SFMToolActiveStateChanged) XXX_Size() int {
	return xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged.Size(m)
}
func (m *ToolEvent_SFMToolActiveStateChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged.DiscardUnknown(m)
}

var xxx_messageInfo_ToolEvent_SFMToolActiveStateChanged proto.InternalMessageInfo

func (m *ToolEvent_SFMToolActiveStateChanged) GetIsactive() bool {
	if m != nil && m.Isactive != nil {
		return *m.Isactive
	}
	return false
}

func init() {
	proto.RegisterType((*ChangeMapToolEvent)(nil), "ChangeMapToolEvent")
	proto.RegisterType((*TraceRayServerToolEvent)(nil), "TraceRayServerToolEvent")
	proto.RegisterType((*ToolTraceRayResult)(nil), "ToolTraceRayResult")
	proto.RegisterType((*SpawnEntityToolEvent)(nil), "SpawnEntityToolEvent")
	proto.RegisterType((*SpawnEntityToolEventResult)(nil), "SpawnEntityToolEventResult")
	proto.RegisterType((*DestroyEntityToolEvent)(nil), "DestroyEntityToolEvent")
	proto.RegisterType((*DestroyAllEntitiesToolEvent)(nil), "DestroyAllEntitiesToolEvent")
	proto.RegisterType((*RestartMapToolEvent)(nil), "RestartMapToolEvent")
	proto.RegisterType((*ToolEvent_GetEntityInfo)(nil), "ToolEvent_GetEntityInfo")
	proto.RegisterType((*ToolEvent_GetEntityInfoResult)(nil), "ToolEvent_GetEntityInfoResult")
	proto.RegisterType((*ToolEvent_GetEntityInputs)(nil), "ToolEvent_GetEntityInputs")
	proto.RegisterType((*ToolEvent_GetEntityInputsResult)(nil), "ToolEvent_GetEntityInputsResult")
	proto.RegisterType((*ToolEvent_FireEntityInput)(nil), "ToolEvent_FireEntityInput")
	proto.RegisterType((*ToolEvent_SFMRecordingStateChanged)(nil), "ToolEvent_SFMRecordingStateChanged")
	proto.RegisterType((*ToolEvent_SFMToolActiveStateChanged)(nil), "ToolEvent_SFMToolActiveStateChanged")
}

func init() {
	proto.RegisterFile("gametoolevents.proto", fileDescriptor_gametoolevents_115765871d01f2d3)
}

var fileDescriptor_gametoolevents_115765871d01f2d3 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0xfe, 0x2d, 0x05, 0x7e, 0xed, 0xa9, 0x89, 0x64, 0x45, 0x58, 0xab, 0x0d, 0x75, 0xf1, 0x02,
	0x4d, 0x68, 0x15, 0xa3, 0x17, 0x5e, 0x89, 0x08, 0xc6, 0x28, 0xc6, 0x6c, 0x89, 0x17, 0x7a, 0xd1,
	0x1c, 0x76, 0x0f, 0xdb, 0x09, 0xbb, 0x33, 0x9b, 0x99, 0x69, 0xa1, 0x77, 0xbe, 0x8a, 0xaf, 0xe2,
	0x83, 0xf8, 0x2c, 0x66, 0x66, 0xb6, 0x7f, 0x80, 0xf6, 0xc2, 0x84, 0xbb, 0x39, 0xdf, 0x77, 0xfe,
	0xcd, 0x77, 0xce, 0x0c, 0xac, 0xa7, 0x98, 0x93, 0x16, 0x22, 0xa3, 0x21, 0x71, 0xad, 0xda, 0x85,
	0x14, 0x5a, 0x34, 0x36, 0x38, 0xe9, 0x0b, 0x21, 0xcf, 0x4f, 0x51, 0x91, 0x1e, 0x15, 0x54, 0xe2,
	0x61, 0x1b, 0xfc, 0x83, 0x3e, 0xf2, 0x94, 0x8e, 0xb1, 0x38, 0x11, 0x22, 0x3b, 0x34, 0x41, 0x7e,
	0x00, 0xff, 0xe7, 0x58, 0x70, 0xcc, 0x29, 0xf0, 0x5a, 0xde, 0x4e, 0x2d, 0x1a, 0x9b, 0xe1, 0x0f,
	0xd8, 0x3c, 0x91, 0x18, 0x53, 0x84, 0xa3, 0x2e, 0xc9, 0x21, 0xc9, 0x69, 0xd0, 0x63, 0x58, 0x51,
	0x1a, 0xa5, 0xb6, 0x21, 0xf5, 0xbd, 0x7a, 0xfb, 0xe0, 0x58, 0xa5, 0xdf, 0x28, 0xd6, 0x42, 0x46,
	0x8e, 0xf1, 0x9b, 0x50, 0x21, 0x9e, 0x04, 0x4b, 0x37, 0x1d, 0x0c, 0x1e, 0xfe, 0xf6, 0xc0, 0x37,
	0xf9, 0xc6, 0x15, 0x22, 0x52, 0x83, 0x4c, 0xfb, 0x6b, 0x50, 0xe9, 0x33, 0x97, 0xb6, 0x1a, 0x99,
	0xa3, 0xbf, 0x0d, 0xab, 0x2c, 0x2f, 0x30, 0xd6, 0xf3, 0x52, 0x95, 0x94, 0x71, 0xe2, 0x42, 0xe6,
	0x98, 0x05, 0x95, 0x39, 0x4e, 0x8e, 0xf2, 0x1b, 0x50, 0x4d, 0x98, 0xd2, 0xc8, 0x63, 0x0a, 0x96,
	0x5b, 0xde, 0xce, 0x52, 0x34, 0xb1, 0x0d, 0x77, 0x26, 0x31, 0xd6, 0x4c, 0xf0, 0x60, 0xc5, 0x71,
	0x63, 0xdb, 0x28, 0x44, 0x7d, 0xe4, 0x49, 0x46, 0xc1, 0x6a, 0xcb, 0xdb, 0x59, 0x89, 0xc6, 0x66,
	0x98, 0xc3, 0x7a, 0xb7, 0xc0, 0x0b, 0x7e, 0xc8, 0x35, 0xd3, 0xa3, 0xa9, 0x3c, 0x4f, 0x61, 0x8d,
	0x2c, 0xd4, 0x3b, 0xa7, 0xd1, 0x10, 0xb3, 0x01, 0x29, 0x7b, 0xa5, 0x3b, 0xd1, 0x5d, 0x87, 0x7f,
	0x1a, 0xc3, 0xfe, 0x33, 0x58, 0x8b, 0x33, 0x66, 0xa6, 0xc7, 0x12, 0x72, 0xa4, 0xbd, 0x68, 0x35,
	0xba, 0x81, 0x87, 0xaf, 0xa1, 0x31, 0xaf, 0x5c, 0x29, 0xdd, 0x4c, 0x9b, 0xde, 0xd5, 0x36, 0xf7,
	0x60, 0xe3, 0x3d, 0x29, 0x2d, 0xc5, 0xe8, 0x7a, 0xa3, 0x8b, 0x63, 0x9a, 0xf0, 0xb0, 0x8c, 0xd9,
	0xcf, 0x32, 0x1b, 0xc6, 0x48, 0x4d, 0x02, 0xc3, 0xfb, 0x70, 0x2f, 0x22, 0x3b, 0xe8, 0xd9, 0x65,
	0x0a, 0x7b, 0xb0, 0x39, 0x31, 0x7a, 0x1f, 0x48, 0xbb, 0x7a, 0x1f, 0xf9, 0x99, 0x58, 0x5c, 0xea,
	0x9f, 0x24, 0xf8, 0xe3, 0x41, 0x73, 0x41, 0x85, 0x52, 0x86, 0x27, 0x50, 0x8d, 0x8b, 0x22, 0xce,
	0x50, 0x39, 0xcd, 0x6b, 0x6f, 0xaa, 0xaa, 0xcf, 0x74, 0x9f, 0x30, 0x89, 0x26, 0x8c, 0xff, 0x08,
	0x6a, 0xf6, 0x60, 0xf7, 0x7e, 0xc9, 0xee, 0xfd, 0x14, 0xf0, 0x7d, 0x58, 0xb6, 0x44, 0xc5, 0x12,
	0xf6, 0x6c, 0x56, 0x4c, 0x48, 0x96, 0x32, 0x6e, 0x77, 0xe7, 0xfa, 0x8a, 0x39, 0xca, 0xdf, 0x82,
	0xe5, 0x9c, 0x71, 0x65, 0x57, 0xe8, 0x9a, 0x8b, 0x25, 0xac, 0x03, 0x5e, 0x2a, 0xbb, 0x48, 0x37,
	0x1c, 0xf0, 0x52, 0x85, 0x08, 0x0f, 0xe6, 0xde, 0xaf, 0x18, 0x68, 0x75, 0x4b, 0x1a, 0xbe, 0x85,
	0xad, 0x85, 0x25, 0x4a, 0x11, 0x9b, 0x00, 0xcc, 0xd8, 0xbd, 0x8c, 0x29, 0xf3, 0x1a, 0x2b, 0x46,
	0x1f, 0x8b, 0x7c, 0x66, 0x4a, 0x87, 0xbf, 0xbc, 0xd9, 0x2e, 0x8f, 0x98, 0xa4, 0x99, 0x1c, 0xb7,
	0xd3, 0xe5, 0xb4, 0x85, 0x99, 0x49, 0xb8, 0x16, 0xbe, 0x98, 0x71, 0x6c, 0x41, 0xdd, 0xd1, 0x05,
	0x4a, 0xcc, 0xed, 0x4c, 0x6a, 0x91, 0x8b, 0xf8, 0x6a, 0x90, 0xf0, 0x08, 0xc2, 0x69, 0x8b, 0xdd,
	0xa3, 0xe3, 0x88, 0x62, 0x21, 0x13, 0xc6, 0xd3, 0xae, 0x46, 0x4d, 0xee, 0x23, 0x4c, 0xfc, 0x16,
	0xd4, 0x99, 0x92, 0x63, 0xaa, 0xfc, 0x77, 0x66, 0xa1, 0x70, 0x1f, 0xb6, 0xaf, 0xe4, 0x31, 0xc6,
	0x7e, 0xac, 0xd9, 0x90, 0xae, 0x24, 0x6a, 0x40, 0x95, 0x29, 0xb4, 0x78, 0x99, 0x65, 0x62, 0xbf,
	0x7b, 0xf5, 0xfd, 0x79, 0xca, 0x74, 0x7f, 0x70, 0xda, 0x8e, 0x45, 0xde, 0x79, 0xf1, 0xf2, 0xbc,
	0x93, 0x8a, 0x5d, 0xa5, 0x09, 0xf3, 0x5d, 0x49, 0x4a, 0x0c, 0x64, 0x4c, 0xaa, 0x63, 0xff, 0xe8,
	0xd3, 0xc1, 0x59, 0x27, 0x11, 0x1a, 0xf7, 0x7e, 0x7a, 0xff, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x8b, 0x96, 0xb7, 0xde, 0x05, 0x00, 0x00,
}