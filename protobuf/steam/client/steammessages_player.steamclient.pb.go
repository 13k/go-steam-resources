// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_player.steamclient.proto

package client // import "github.com/13k/go-steam-resources/protobuf/steam/client"

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

type CPlayer_GetGameBadgeLevels_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Request) Reset()         { *m = CPlayer_GetGameBadgeLevels_Request{} }
func (m *CPlayer_GetGameBadgeLevels_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Request) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{0}
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetGameBadgeLevels_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Merge(dst, src)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Request proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPlayer_GetGameBadgeLevels_Response struct {
	PlayerLevel          *uint32                                      `protobuf:"varint,1,opt,name=player_level,json=playerLevel" json:"player_level,omitempty"`
	Badges               []*CPlayer_GetGameBadgeLevels_Response_Badge `protobuf:"bytes,2,rep,name=badges" json:"badges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response) Reset()         { *m = CPlayer_GetGameBadgeLevels_Response{} }
func (m *CPlayer_GetGameBadgeLevels_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{1}
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetGameBadgeLevels_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Merge(dst, src)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Response) GetPlayerLevel() uint32 {
	if m != nil && m.PlayerLevel != nil {
		return *m.PlayerLevel
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response) GetBadges() []*CPlayer_GetGameBadgeLevels_Response_Badge {
	if m != nil {
		return m.Badges
	}
	return nil
}

type CPlayer_GetGameBadgeLevels_Response_Badge struct {
	Level                *int32   `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Series               *int32   `protobuf:"varint,2,opt,name=series" json:"series,omitempty"`
	BorderColor          *uint32  `protobuf:"varint,3,opt,name=border_color,json=borderColor" json:"border_color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) Reset() {
	*m = CPlayer_GetGameBadgeLevels_Response_Badge{}
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response_Badge) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response_Badge) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{1, 0}
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Unmarshal(m, b)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Merge(dst, src)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.Size(m)
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetGameBadgeLevels_Response_Badge proto.InternalMessageInfo

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetSeries() int32 {
	if m != nil && m.Series != nil {
		return *m.Series
	}
	return 0
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) GetBorderColor() uint32 {
	if m != nil && m.BorderColor != nil {
		return *m.BorderColor
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Request struct {
	MinLastPlayed        *uint32  `protobuf:"varint,1,opt,name=min_last_played,json=minLastPlayed" json:"min_last_played,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Request) Reset()         { *m = CPlayer_GetLastPlayedTimes_Request{} }
func (m *CPlayer_GetLastPlayedTimes_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Request) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{2}
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetLastPlayedTimes_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Merge(dst, src)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Request proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Request) GetMinLastPlayed() uint32 {
	if m != nil && m.MinLastPlayed != nil {
		return *m.MinLastPlayed
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Response struct {
	Games                []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response) Reset()         { *m = CPlayer_GetLastPlayedTimes_Response{} }
func (m *CPlayer_GetLastPlayedTimes_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{3}
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetLastPlayedTimes_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Merge(dst, src)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Response) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayer_GetLastPlayedTimes_Response_Game struct {
	Appid                *int32   `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	LastPlaytime         *uint32  `protobuf:"varint,2,opt,name=last_playtime,json=lastPlaytime" json:"last_playtime,omitempty"`
	Playtime_2Weeks      *int32   `protobuf:"varint,3,opt,name=playtime_2weeks,json=playtime2weeks" json:"playtime_2weeks,omitempty"`
	PlaytimeForever      *int32   `protobuf:"varint,4,opt,name=playtime_forever,json=playtimeForever" json:"playtime_forever,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) Reset() {
	*m = CPlayer_GetLastPlayedTimes_Response_Game{}
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response_Game) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response_Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{3, 0}
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Unmarshal(m, b)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Merge(dst, src)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.Size(m)
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetLastPlayedTimes_Response_Game proto.InternalMessageInfo

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetAppid() int32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetLastPlaytime() uint32 {
	if m != nil && m.LastPlaytime != nil {
		return *m.LastPlaytime
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytime_2Weeks() int32 {
	if m != nil && m.Playtime_2Weeks != nil {
		return *m.Playtime_2Weeks
	}
	return 0
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) GetPlaytimeForever() int32 {
	if m != nil && m.PlaytimeForever != nil {
		return *m.PlaytimeForever
	}
	return 0
}

type CPlayer_AcceptSSA_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_AcceptSSA_Request) Reset()         { *m = CPlayer_AcceptSSA_Request{} }
func (m *CPlayer_AcceptSSA_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Request) ProtoMessage()    {}
func (*CPlayer_AcceptSSA_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{4}
}
func (m *CPlayer_AcceptSSA_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Unmarshal(m, b)
}
func (m *CPlayer_AcceptSSA_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Marshal(b, m, deterministic)
}
func (dst *CPlayer_AcceptSSA_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_AcceptSSA_Request.Merge(dst, src)
}
func (m *CPlayer_AcceptSSA_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_AcceptSSA_Request.Size(m)
}
func (m *CPlayer_AcceptSSA_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_AcceptSSA_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_AcceptSSA_Request proto.InternalMessageInfo

type CPlayer_AcceptSSA_Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_AcceptSSA_Response) Reset()         { *m = CPlayer_AcceptSSA_Response{} }
func (m *CPlayer_AcceptSSA_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Response) ProtoMessage()    {}
func (*CPlayer_AcceptSSA_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{5}
}
func (m *CPlayer_AcceptSSA_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Unmarshal(m, b)
}
func (m *CPlayer_AcceptSSA_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Marshal(b, m, deterministic)
}
func (dst *CPlayer_AcceptSSA_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_AcceptSSA_Response.Merge(dst, src)
}
func (m *CPlayer_AcceptSSA_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_AcceptSSA_Response.Size(m)
}
func (m *CPlayer_AcceptSSA_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_AcceptSSA_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_AcceptSSA_Response proto.InternalMessageInfo

type CPlayer_GetNicknameList_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetNicknameList_Request) Reset()         { *m = CPlayer_GetNicknameList_Request{} }
func (m *CPlayer_GetNicknameList_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetNicknameList_Request) ProtoMessage()    {}
func (*CPlayer_GetNicknameList_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{6}
}
func (m *CPlayer_GetNicknameList_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetNicknameList_Request.Unmarshal(m, b)
}
func (m *CPlayer_GetNicknameList_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetNicknameList_Request.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetNicknameList_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetNicknameList_Request.Merge(dst, src)
}
func (m *CPlayer_GetNicknameList_Request) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetNicknameList_Request.Size(m)
}
func (m *CPlayer_GetNicknameList_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetNicknameList_Request.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetNicknameList_Request proto.InternalMessageInfo

type CPlayer_GetNicknameList_Response struct {
	Nicknames            []*CPlayer_GetNicknameList_Response_PlayerNickname `protobuf:"bytes,1,rep,name=nicknames" json:"nicknames,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *CPlayer_GetNicknameList_Response) Reset()         { *m = CPlayer_GetNicknameList_Response{} }
func (m *CPlayer_GetNicknameList_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetNicknameList_Response) ProtoMessage()    {}
func (*CPlayer_GetNicknameList_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{7}
}
func (m *CPlayer_GetNicknameList_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response.Unmarshal(m, b)
}
func (m *CPlayer_GetNicknameList_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetNicknameList_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetNicknameList_Response.Merge(dst, src)
}
func (m *CPlayer_GetNicknameList_Response) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response.Size(m)
}
func (m *CPlayer_GetNicknameList_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetNicknameList_Response.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetNicknameList_Response proto.InternalMessageInfo

func (m *CPlayer_GetNicknameList_Response) GetNicknames() []*CPlayer_GetNicknameList_Response_PlayerNickname {
	if m != nil {
		return m.Nicknames
	}
	return nil
}

type CPlayer_GetNicknameList_Response_PlayerNickname struct {
	Accountid            *uint32  `protobuf:"fixed32,1,opt,name=accountid" json:"accountid,omitempty"`
	Nickname             *string  `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_GetNicknameList_Response_PlayerNickname) Reset() {
	*m = CPlayer_GetNicknameList_Response_PlayerNickname{}
}
func (m *CPlayer_GetNicknameList_Response_PlayerNickname) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayer_GetNicknameList_Response_PlayerNickname) ProtoMessage() {}
func (*CPlayer_GetNicknameList_Response_PlayerNickname) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{7, 0}
}
func (m *CPlayer_GetNicknameList_Response_PlayerNickname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname.Unmarshal(m, b)
}
func (m *CPlayer_GetNicknameList_Response_PlayerNickname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname.Marshal(b, m, deterministic)
}
func (dst *CPlayer_GetNicknameList_Response_PlayerNickname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname.Merge(dst, src)
}
func (m *CPlayer_GetNicknameList_Response_PlayerNickname) XXX_Size() int {
	return xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname.Size(m)
}
func (m *CPlayer_GetNicknameList_Response_PlayerNickname) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_GetNicknameList_Response_PlayerNickname proto.InternalMessageInfo

func (m *CPlayer_GetNicknameList_Response_PlayerNickname) GetAccountid() uint32 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *CPlayer_GetNicknameList_Response_PlayerNickname) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

type CPlayer_LastPlayedTimes_Notification struct {
	Games                []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CPlayer_LastPlayedTimes_Notification) Reset()         { *m = CPlayer_LastPlayedTimes_Notification{} }
func (m *CPlayer_LastPlayedTimes_Notification) String() string { return proto.CompactTextString(m) }
func (*CPlayer_LastPlayedTimes_Notification) ProtoMessage()    {}
func (*CPlayer_LastPlayedTimes_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{8}
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Unmarshal(m, b)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Marshal(b, m, deterministic)
}
func (dst *CPlayer_LastPlayedTimes_Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Merge(dst, src)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_Size() int {
	return xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.Size(m)
}
func (m *CPlayer_LastPlayedTimes_Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_LastPlayedTimes_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_LastPlayedTimes_Notification proto.InternalMessageInfo

func (m *CPlayer_LastPlayedTimes_Notification) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayer_FriendNicknameChanged_Notification struct {
	Accountid            *uint32  `protobuf:"fixed32,1,opt,name=accountid" json:"accountid,omitempty"`
	Nickname             *string  `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPlayer_FriendNicknameChanged_Notification) Reset() {
	*m = CPlayer_FriendNicknameChanged_Notification{}
}
func (m *CPlayer_FriendNicknameChanged_Notification) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayer_FriendNicknameChanged_Notification) ProtoMessage() {}
func (*CPlayer_FriendNicknameChanged_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e, []int{9}
}
func (m *CPlayer_FriendNicknameChanged_Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification.Unmarshal(m, b)
}
func (m *CPlayer_FriendNicknameChanged_Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification.Marshal(b, m, deterministic)
}
func (dst *CPlayer_FriendNicknameChanged_Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification.Merge(dst, src)
}
func (m *CPlayer_FriendNicknameChanged_Notification) XXX_Size() int {
	return xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification.Size(m)
}
func (m *CPlayer_FriendNicknameChanged_Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_CPlayer_FriendNicknameChanged_Notification proto.InternalMessageInfo

func (m *CPlayer_FriendNicknameChanged_Notification) GetAccountid() uint32 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *CPlayer_FriendNicknameChanged_Notification) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func init() {
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Request)(nil), "CPlayer_GetGameBadgeLevels_Request")
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Response)(nil), "CPlayer_GetGameBadgeLevels_Response")
	proto.RegisterType((*CPlayer_GetGameBadgeLevels_Response_Badge)(nil), "CPlayer_GetGameBadgeLevels_Response.Badge")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Request)(nil), "CPlayer_GetLastPlayedTimes_Request")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Response)(nil), "CPlayer_GetLastPlayedTimes_Response")
	proto.RegisterType((*CPlayer_GetLastPlayedTimes_Response_Game)(nil), "CPlayer_GetLastPlayedTimes_Response.Game")
	proto.RegisterType((*CPlayer_AcceptSSA_Request)(nil), "CPlayer_AcceptSSA_Request")
	proto.RegisterType((*CPlayer_AcceptSSA_Response)(nil), "CPlayer_AcceptSSA_Response")
	proto.RegisterType((*CPlayer_GetNicknameList_Request)(nil), "CPlayer_GetNicknameList_Request")
	proto.RegisterType((*CPlayer_GetNicknameList_Response)(nil), "CPlayer_GetNicknameList_Response")
	proto.RegisterType((*CPlayer_GetNicknameList_Response_PlayerNickname)(nil), "CPlayer_GetNicknameList_Response.PlayerNickname")
	proto.RegisterType((*CPlayer_LastPlayedTimes_Notification)(nil), "CPlayer_LastPlayedTimes_Notification")
	proto.RegisterType((*CPlayer_FriendNicknameChanged_Notification)(nil), "CPlayer_FriendNicknameChanged_Notification")
}

func init() {
	proto.RegisterFile("steammessages_player.steamclient.proto", fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e)
}

var fileDescriptor_steammessages_player_steamclient_588f02fe3c090d5e = []byte{
	// 947 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0x97, 0xd3, 0x6e, 0x20, 0x93, 0xa4, 0x41, 0x23, 0x40, 0xae, 0x53, 0xc4, 0xc4, 0x29, 0x4d,
	0x13, 0x12, 0x07, 0xc2, 0x01, 0x04, 0x95, 0xa2, 0x64, 0xa5, 0x46, 0x82, 0x28, 0x02, 0xa7, 0x95,
	0x10, 0x97, 0xd5, 0xac, 0xfd, 0xbc, 0x3b, 0x5a, 0xdb, 0xb3, 0xcc, 0x8c, 0x53, 0xe5, 0x86, 0x56,
	0x1c, 0xb9, 0x20, 0xd1, 0x6f, 0xc0, 0x57, 0x30, 0x47, 0xc4, 0xa7, 0xe1, 0x13, 0x70, 0x46, 0x68,
	0x66, 0x6c, 0x6f, 0x9c, 0x6d, 0x9a, 0x15, 0x3d, 0xce, 0xef, 0xfd, 0x99, 0xdf, 0xfb, 0xbd, 0x79,
	0xcf, 0x46, 0x8f, 0xa4, 0x02, 0x9a, 0x65, 0x20, 0x25, 0x1d, 0x80, 0xec, 0x8d, 0x53, 0x7a, 0x09,
	0x22, 0x30, 0x60, 0x94, 0x32, 0xc8, 0x55, 0x30, 0x16, 0x5c, 0x71, 0x6f, 0xb7, 0xed, 0x57, 0xe4,
	0x2c, 0x61, 0x10, 0xf7, 0xfa, 0x54, 0xc2, 0xac, 0xb7, 0xff, 0x25, 0xf2, 0xbb, 0xdf, 0x9a, 0x54,
	0xbd, 0x13, 0x50, 0x27, 0x34, 0x83, 0x63, 0x1a, 0x0f, 0xe0, 0x14, 0x2e, 0x20, 0x95, 0xbd, 0x10,
	0x7e, 0x2c, 0x40, 0x2a, 0xfc, 0x2e, 0xea, 0xd0, 0xf1, 0x98, 0xc5, 0xae, 0x43, 0x9c, 0xc7, 0xab,
	0xa1, 0x3d, 0xf8, 0x7f, 0x3b, 0x68, 0xf3, 0xb5, 0xc1, 0x72, 0xcc, 0x73, 0x09, 0x78, 0x03, 0xad,
	0x58, 0xb6, 0xbd, 0x54, 0x5b, 0xaa, 0x24, 0xcb, 0x16, 0x33, 0xce, 0xf8, 0x18, 0x2d, 0xf6, 0x75,
	0xa8, 0x74, 0x17, 0xc8, 0x9d, 0xc7, 0xcb, 0x07, 0x3b, 0xc1, 0x1c, 0x89, 0x03, 0x03, 0x86, 0x55,
	0xa4, 0xf7, 0x3d, 0xea, 0x18, 0x40, 0xb3, 0x9d, 0x5e, 0xd4, 0x09, 0xed, 0x01, 0xbf, 0x8f, 0x16,
	0x25, 0x08, 0x66, 0xae, 0xd0, 0x70, 0x75, 0xd2, 0xec, 0xfa, 0x5c, 0xc4, 0x20, 0x7a, 0x11, 0x4f,
	0xb9, 0x70, 0xef, 0x58, 0x76, 0x16, 0xeb, 0x6a, 0xc8, 0xff, 0xd5, 0x69, 0xa9, 0x74, 0x4a, 0xa5,
	0x32, 0xa7, 0xf8, 0x19, 0xcb, 0x60, 0xaa, 0xd2, 0x08, 0xad, 0x65, 0x2c, 0xef, 0xa5, 0x54, 0x2a,
	0xdb, 0x9e, 0x4a, 0xaf, 0xe3, 0xee, 0xa4, 0x74, 0x0f, 0x9f, 0x0d, 0x81, 0x64, 0x5c, 0x2a, 0x22,
	0x20, 0x82, 0x5c, 0x11, 0xed, 0xb6, 0x67, 0xdd, 0x88, 0x62, 0x19, 0x10, 0x35, 0x04, 0x62, 0x7b,
	0x43, 0x68, 0x2a, 0x80, 0xc6, 0x97, 0x64, 0x94, 0xf3, 0x17, 0x92, 0xd0, 0x3e, 0x2f, 0x54, 0xb8,
	0x9a, 0xb1, 0x7c, 0x7a, 0xb3, 0xff, 0x6f, 0x5b, 0xfc, 0x59, 0x4e, 0x95, 0xf8, 0x87, 0xa8, 0x33,
	0xa0, 0x19, 0x48, 0xd7, 0x31, 0xc2, 0x6e, 0x07, 0x73, 0x04, 0x05, 0x5a, 0xf1, 0xd0, 0xc6, 0x79,
	0x2f, 0x1d, 0x74, 0x57, 0x9f, 0xdb, 0x8f, 0xa0, 0x53, 0x3d, 0x02, 0xbc, 0x89, 0x56, 0x9b, 0x82,
	0x75, 0x19, 0x46, 0xdd, 0xd5, 0x70, 0x25, 0xad, 0x72, 0x6b, 0x0c, 0x6f, 0xa1, 0xb5, 0xda, 0xde,
	0x3b, 0x78, 0x01, 0x30, 0x92, 0x46, 0xe6, 0x4e, 0x78, 0xaf, 0x86, 0x2d, 0x8a, 0xb7, 0xd1, 0x3b,
	0x8d, 0x63, 0xc2, 0x05, 0x5c, 0x80, 0x70, 0xef, 0x1a, 0xcf, 0x26, 0xc1, 0x53, 0x0b, 0xfb, 0xeb,
	0xe8, 0x7e, 0x5d, 0xca, 0x51, 0x14, 0xc1, 0x58, 0x9d, 0x9f, 0x1f, 0xd5, 0xad, 0xf0, 0x1f, 0x20,
	0xef, 0x55, 0x46, 0x5b, 0x9e, 0xbf, 0x81, 0x3e, 0xbc, 0xa2, 0xc2, 0x19, 0x8b, 0x46, 0x39, 0xcd,
	0xe0, 0x94, 0x49, 0xd5, 0x24, 0xf8, 0xd3, 0x41, 0xe4, 0x66, 0x9f, 0x4a, 0xdb, 0x33, 0xb4, 0x94,
	0x57, 0x86, 0x5a, 0xdf, 0x4f, 0x82, 0xdb, 0xa2, 0x02, 0x6b, 0xaf, 0x6d, 0xe1, 0x34, 0x85, 0xf7,
	0x35, 0xba, 0xd7, 0x36, 0xe2, 0x07, 0x68, 0x89, 0x46, 0x11, 0x2f, 0x72, 0x55, 0xe9, 0xfe, 0x56,
	0x38, 0x05, 0xb0, 0x87, 0xde, 0xae, 0x83, 0x8d, 0xec, 0x4b, 0x61, 0x73, 0xf6, 0x07, 0xe8, 0x61,
	0xcd, 0xe4, 0x7a, 0x9b, 0xcf, 0xb8, 0x62, 0x09, 0x8b, 0xa8, 0x62, 0x3c, 0x7f, 0xe3, 0xf7, 0xe1,
	0x27, 0x68, 0xa7, 0x0e, 0x79, 0x2a, 0x18, 0xe4, 0x71, 0x4d, 0xbe, 0x3b, 0xa4, 0xf9, 0x00, 0xe2,
	0xf6, 0x75, 0xff, 0xbb, 0xa0, 0x83, 0x9f, 0x3b, 0x68, 0xd1, 0xde, 0x83, 0xff, 0x70, 0x10, 0x9e,
	0xdd, 0x0b, 0x78, 0x33, 0xb8, 0x7d, 0x95, 0x79, 0x0f, 0xe7, 0xd9, 0x2c, 0xfe, 0xf3, 0x49, 0xe9,
	0x7e, 0x17, 0x82, 0x2a, 0x44, 0x2e, 0xcd, 0x60, 0x9e, 0xeb, 0xcd, 0x49, 0x8c, 0x1b, 0xe1, 0x09,
	0xa1, 0xa4, 0x90, 0x20, 0x76, 0x8d, 0xc9, 0x24, 0x20, 0x66, 0xc9, 0x90, 0x84, 0x0b, 0x83, 0x69,
	0x69, 0x76, 0x09, 0xcd, 0x63, 0xc2, 0x12, 0xc2, 0xd4, 0x96, 0x24, 0x09, 0x67, 0x29, 0x7e, 0xe9,
	0x20, 0xb7, 0x6b, 0x46, 0x7c, 0x56, 0xdd, 0x36, 0xfd, 0x1b, 0x76, 0x4c, 0x9b, 0xfe, 0x4d, 0xfd,
	0xf1, 0x83, 0x49, 0xe9, 0xee, 0x9c, 0x80, 0xb2, 0xdc, 0xaf, 0x6f, 0x1a, 0xd9, 0xd0, 0xac, 0x94,
	0xc7, 0x11, 0x5a, 0x6a, 0xc6, 0x04, 0x7b, 0xc1, 0x8d, 0x73, 0xe5, 0xad, 0x07, 0xaf, 0x19, 0xab,
	0x0f, 0x26, 0xa5, 0x7b, 0xff, 0xb9, 0x04, 0x41, 0x98, 0xd4, 0xa9, 0x61, 0xac, 0x58, 0x3e, 0xb0,
	0xf2, 0x9d, 0x1f, 0xe1, 0xdf, 0x1c, 0xb4, 0x76, 0x6d, 0x28, 0x30, 0x09, 0x6e, 0x19, 0x44, 0x6f,
	0xe3, 0xd6, 0x81, 0xf2, 0x9f, 0x4c, 0x4a, 0xf7, 0x8b, 0x69, 0xb5, 0x4c, 0x2a, 0xdd, 0xa2, 0x66,
	0xac, 0x88, 0x1a, 0x32, 0x69, 0xfa, 0x45, 0x86, 0xd4, 0xd6, 0xce, 0xd5, 0x10, 0x84, 0x81, 0xa4,
	0xb7, 0x37, 0x29, 0xdd, 0xed, 0x23, 0x22, 0x41, 0x5c, 0xb0, 0x08, 0x8c, 0x59, 0x73, 0x97, 0x52,
	0x73, 0xb7, 0x6d, 0xb7, 0x5f, 0x2a, 0x12, 0x53, 0x45, 0x0f, 0xfe, 0x59, 0x40, 0x2b, 0x96, 0x90,
	0x6d, 0x24, 0xfe, 0xc5, 0x41, 0xef, 0x99, 0x27, 0x7e, 0x79, 0xbd, 0xa1, 0x1f, 0x05, 0xf3, 0x4c,
	0xa0, 0xb7, 0x1c, 0x9c, 0xf1, 0xa6, 0x96, 0xc3, 0x49, 0xe9, 0x7e, 0x75, 0xd5, 0x4c, 0x12, 0xc1,
	0x33, 0xc3, 0x0e, 0x04, 0x51, 0xbc, 0xfe, 0x42, 0xf0, 0x84, 0x64, 0x5c, 0x40, 0xfd, 0x31, 0xd1,
	0x0c, 0x4d, 0x6f, 0xf1, 0xef, 0x0e, 0x5a, 0xb7, 0x74, 0x5e, 0x39, 0x8d, 0xf8, 0xe3, 0x60, 0xfe,
	0x69, 0x6d, 0x53, 0xfb, 0x66, 0x52, 0xba, 0x27, 0x73, 0x50, 0x53, 0x43, 0xaa, 0x08, 0x25, 0x89,
	0x49, 0xbf, 0x25, 0x9b, 0x46, 0x18, 0xf5, 0x23, 0x7b, 0x8d, 0xa7, 0x5f, 0xe8, 0xa3, 0x19, 0x75,
	0xeb, 0x14, 0xf9, 0x95, 0x5b, 0xe4, 0x5f, 0xa5, 0xbb, 0x70, 0xfc, 0xe4, 0x87, 0xcf, 0x07, 0x4c,
	0x0d, 0x8b, 0x7e, 0x10, 0xf1, 0x6c, 0xff, 0xd3, 0xcf, 0x46, 0xfb, 0x03, 0xbe, 0x67, 0x7e, 0x67,
	0xf6, 0x04, 0x48, 0x5e, 0x88, 0x08, 0xe4, 0xbe, 0xf9, 0xa5, 0xe9, 0x17, 0xc9, 0xbe, 0x31, 0xec,
	0xdb, 0x64, 0x3f, 0x39, 0xce, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xff, 0xc6, 0x53, 0x3c,
	0x09, 0x00, 0x00,
}
