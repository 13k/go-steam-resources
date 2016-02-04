// Code generated by protoc-gen-go.
// source: steammessages_player.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CPlayer_GetGameBadgeLevels_Request struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Request) Reset()         { *m = CPlayer_GetGameBadgeLevels_Request{} }
func (m *CPlayer_GetGameBadgeLevels_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Request) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{0}
}

func (m *CPlayer_GetGameBadgeLevels_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPlayer_GetGameBadgeLevels_Response struct {
	PlayerLevel      *uint32                                      `protobuf:"varint,1,opt,name=player_level" json:"player_level,omitempty"`
	Badges           []*CPlayer_GetGameBadgeLevels_Response_Badge `protobuf:"bytes,2,rep,name=badges" json:"badges,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response) Reset()         { *m = CPlayer_GetGameBadgeLevels_Response{} }
func (m *CPlayer_GetGameBadgeLevels_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1}
}

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
	Level            *int32  `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Series           *int32  `protobuf:"varint,2,opt,name=series" json:"series,omitempty"`
	BorderColor      *uint32 `protobuf:"varint,3,opt,name=border_color" json:"border_color,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetGameBadgeLevels_Response_Badge) Reset() {
	*m = CPlayer_GetGameBadgeLevels_Response_Badge{}
}
func (m *CPlayer_GetGameBadgeLevels_Response_Badge) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetGameBadgeLevels_Response_Badge) ProtoMessage()    {}
func (*CPlayer_GetGameBadgeLevels_Response_Badge) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{1, 0}
}

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
	MinLastPlayed    *uint32 `protobuf:"varint,1,opt,name=min_last_played" json:"min_last_played,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Request) Reset()         { *m = CPlayer_GetLastPlayedTimes_Request{} }
func (m *CPlayer_GetLastPlayedTimes_Request) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Request) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{2}
}

func (m *CPlayer_GetLastPlayedTimes_Request) GetMinLastPlayed() uint32 {
	if m != nil && m.MinLastPlayed != nil {
		return *m.MinLastPlayed
	}
	return 0
}

type CPlayer_GetLastPlayedTimes_Response struct {
	Games            []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response) Reset()         { *m = CPlayer_GetLastPlayedTimes_Response{} }
func (m *CPlayer_GetLastPlayedTimes_Response) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{3}
}

func (m *CPlayer_GetLastPlayedTimes_Response) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayer_GetLastPlayedTimes_Response_Game struct {
	Appid            *int32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	LastPlaytime     *uint32 `protobuf:"varint,2,opt,name=last_playtime" json:"last_playtime,omitempty"`
	Playtime_2Weeks  *int32  `protobuf:"varint,3,opt,name=playtime_2weeks" json:"playtime_2weeks,omitempty"`
	PlaytimeForever  *int32  `protobuf:"varint,4,opt,name=playtime_forever" json:"playtime_forever,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPlayer_GetLastPlayedTimes_Response_Game) Reset() {
	*m = CPlayer_GetLastPlayedTimes_Response_Game{}
}
func (m *CPlayer_GetLastPlayedTimes_Response_Game) String() string { return proto.CompactTextString(m) }
func (*CPlayer_GetLastPlayedTimes_Response_Game) ProtoMessage()    {}
func (*CPlayer_GetLastPlayedTimes_Response_Game) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{3, 0}
}

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
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayer_AcceptSSA_Request) Reset()                    { *m = CPlayer_AcceptSSA_Request{} }
func (m *CPlayer_AcceptSSA_Request) String() string            { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Request) ProtoMessage()               {}
func (*CPlayer_AcceptSSA_Request) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

type CPlayer_AcceptSSA_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayer_AcceptSSA_Response) Reset()                    { *m = CPlayer_AcceptSSA_Response{} }
func (m *CPlayer_AcceptSSA_Response) String() string            { return proto.CompactTextString(m) }
func (*CPlayer_AcceptSSA_Response) ProtoMessage()               {}
func (*CPlayer_AcceptSSA_Response) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

type CPlayer_LastPlayedTimes_Notification struct {
	Games            []*CPlayer_GetLastPlayedTimes_Response_Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *CPlayer_LastPlayedTimes_Notification) Reset()         { *m = CPlayer_LastPlayedTimes_Notification{} }
func (m *CPlayer_LastPlayedTimes_Notification) String() string { return proto.CompactTextString(m) }
func (*CPlayer_LastPlayedTimes_Notification) ProtoMessage()    {}
func (*CPlayer_LastPlayedTimes_Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{6}
}

func (m *CPlayer_LastPlayedTimes_Notification) GetGames() []*CPlayer_GetLastPlayedTimes_Response_Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type CPlayerClient_GetSystemInformation_Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayerClient_GetSystemInformation_Request) Reset() {
	*m = CPlayerClient_GetSystemInformation_Request{}
}
func (m *CPlayerClient_GetSystemInformation_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayerClient_GetSystemInformation_Request) ProtoMessage() {}
func (*CPlayerClient_GetSystemInformation_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{7}
}

type CClientSystemInfo struct {
	Cpu                *CClientSystemInfo_CPU       `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	VideoCard          *CClientSystemInfo_VideoCard `protobuf:"bytes,2,opt,name=video_card" json:"video_card,omitempty"`
	OperatingSystem    *string                      `protobuf:"bytes,3,opt,name=operating_system" json:"operating_system,omitempty"`
	Os_64Bit           *bool                        `protobuf:"varint,4,opt,name=os_64bit" json:"os_64bit,omitempty"`
	SystemRamMb        *int32                       `protobuf:"varint,5,opt,name=system_ram_mb" json:"system_ram_mb,omitempty"`
	AudioDevice        *string                      `protobuf:"bytes,6,opt,name=audio_device" json:"audio_device,omitempty"`
	AudioDriverVersion *string                      `protobuf:"bytes,7,opt,name=audio_driver_version" json:"audio_driver_version,omitempty"`
	XXX_unrecognized   []byte                       `json:"-"`
}

func (m *CClientSystemInfo) Reset()                    { *m = CClientSystemInfo{} }
func (m *CClientSystemInfo) String() string            { return proto.CompactTextString(m) }
func (*CClientSystemInfo) ProtoMessage()               {}
func (*CClientSystemInfo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *CClientSystemInfo) GetCpu() *CClientSystemInfo_CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CClientSystemInfo) GetVideoCard() *CClientSystemInfo_VideoCard {
	if m != nil {
		return m.VideoCard
	}
	return nil
}

func (m *CClientSystemInfo) GetOperatingSystem() string {
	if m != nil && m.OperatingSystem != nil {
		return *m.OperatingSystem
	}
	return ""
}

func (m *CClientSystemInfo) GetOs_64Bit() bool {
	if m != nil && m.Os_64Bit != nil {
		return *m.Os_64Bit
	}
	return false
}

func (m *CClientSystemInfo) GetSystemRamMb() int32 {
	if m != nil && m.SystemRamMb != nil {
		return *m.SystemRamMb
	}
	return 0
}

func (m *CClientSystemInfo) GetAudioDevice() string {
	if m != nil && m.AudioDevice != nil {
		return *m.AudioDevice
	}
	return ""
}

func (m *CClientSystemInfo) GetAudioDriverVersion() string {
	if m != nil && m.AudioDriverVersion != nil {
		return *m.AudioDriverVersion
	}
	return ""
}

type CClientSystemInfo_CPU struct {
	SpeedMhz           *int32  `protobuf:"varint,1,opt,name=speed_mhz" json:"speed_mhz,omitempty"`
	Vendor             *string `protobuf:"bytes,2,opt,name=vendor" json:"vendor,omitempty"`
	LogicalProcessors  *int32  `protobuf:"varint,3,opt,name=logical_processors" json:"logical_processors,omitempty"`
	PhysicalProcessors *int32  `protobuf:"varint,4,opt,name=physical_processors" json:"physical_processors,omitempty"`
	Hyperthreading     *bool   `protobuf:"varint,5,opt,name=hyperthreading" json:"hyperthreading,omitempty"`
	Fcmov              *bool   `protobuf:"varint,6,opt,name=fcmov" json:"fcmov,omitempty"`
	Sse2               *bool   `protobuf:"varint,7,opt,name=sse2" json:"sse2,omitempty"`
	Sse3               *bool   `protobuf:"varint,8,opt,name=sse3" json:"sse3,omitempty"`
	Ssse3              *bool   `protobuf:"varint,9,opt,name=ssse3" json:"ssse3,omitempty"`
	Sse4A              *bool   `protobuf:"varint,10,opt,name=sse4a" json:"sse4a,omitempty"`
	Sse41              *bool   `protobuf:"varint,11,opt,name=sse41" json:"sse41,omitempty"`
	Sse42              *bool   `protobuf:"varint,12,opt,name=sse42" json:"sse42,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CClientSystemInfo_CPU) Reset()                    { *m = CClientSystemInfo_CPU{} }
func (m *CClientSystemInfo_CPU) String() string            { return proto.CompactTextString(m) }
func (*CClientSystemInfo_CPU) ProtoMessage()               {}
func (*CClientSystemInfo_CPU) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8, 0} }

func (m *CClientSystemInfo_CPU) GetSpeedMhz() int32 {
	if m != nil && m.SpeedMhz != nil {
		return *m.SpeedMhz
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *CClientSystemInfo_CPU) GetLogicalProcessors() int32 {
	if m != nil && m.LogicalProcessors != nil {
		return *m.LogicalProcessors
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetPhysicalProcessors() int32 {
	if m != nil && m.PhysicalProcessors != nil {
		return *m.PhysicalProcessors
	}
	return 0
}

func (m *CClientSystemInfo_CPU) GetHyperthreading() bool {
	if m != nil && m.Hyperthreading != nil {
		return *m.Hyperthreading
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetFcmov() bool {
	if m != nil && m.Fcmov != nil {
		return *m.Fcmov
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse2() bool {
	if m != nil && m.Sse2 != nil {
		return *m.Sse2
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse3() bool {
	if m != nil && m.Sse3 != nil {
		return *m.Sse3
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSsse3() bool {
	if m != nil && m.Ssse3 != nil {
		return *m.Ssse3
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse4A() bool {
	if m != nil && m.Sse4A != nil {
		return *m.Sse4A
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse41() bool {
	if m != nil && m.Sse41 != nil {
		return *m.Sse41
	}
	return false
}

func (m *CClientSystemInfo_CPU) GetSse42() bool {
	if m != nil && m.Sse42 != nil {
		return *m.Sse42
	}
	return false
}

type CClientSystemInfo_VideoCard struct {
	Driver           *string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DriverVersion    *string `protobuf:"bytes,2,opt,name=driver_version" json:"driver_version,omitempty"`
	DriverDate       *uint32 `protobuf:"varint,3,opt,name=driver_date" json:"driver_date,omitempty"`
	DirectxVersion   *string `protobuf:"bytes,4,opt,name=directx_version" json:"directx_version,omitempty"`
	OpenglVersion    *string `protobuf:"bytes,5,opt,name=opengl_version" json:"opengl_version,omitempty"`
	Vendorid         *int32  `protobuf:"varint,6,opt,name=vendorid" json:"vendorid,omitempty"`
	Deviceid         *int32  `protobuf:"varint,7,opt,name=deviceid" json:"deviceid,omitempty"`
	VramMb           *int32  `protobuf:"varint,8,opt,name=vram_mb" json:"vram_mb,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CClientSystemInfo_VideoCard) Reset()                    { *m = CClientSystemInfo_VideoCard{} }
func (m *CClientSystemInfo_VideoCard) String() string            { return proto.CompactTextString(m) }
func (*CClientSystemInfo_VideoCard) ProtoMessage()               {}
func (*CClientSystemInfo_VideoCard) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8, 1} }

func (m *CClientSystemInfo_VideoCard) GetDriver() string {
	if m != nil && m.Driver != nil {
		return *m.Driver
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetDriverVersion() string {
	if m != nil && m.DriverVersion != nil {
		return *m.DriverVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetDriverDate() uint32 {
	if m != nil && m.DriverDate != nil {
		return *m.DriverDate
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetDirectxVersion() string {
	if m != nil && m.DirectxVersion != nil {
		return *m.DirectxVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetOpenglVersion() string {
	if m != nil && m.OpenglVersion != nil {
		return *m.OpenglVersion
	}
	return ""
}

func (m *CClientSystemInfo_VideoCard) GetVendorid() int32 {
	if m != nil && m.Vendorid != nil {
		return *m.Vendorid
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetDeviceid() int32 {
	if m != nil && m.Deviceid != nil {
		return *m.Deviceid
	}
	return 0
}

func (m *CClientSystemInfo_VideoCard) GetVramMb() int32 {
	if m != nil && m.VramMb != nil {
		return *m.VramMb
	}
	return 0
}

type CPlayerClient_GetSystemInformation_Response struct {
	SystemInfo       *CClientSystemInfo `protobuf:"bytes,1,opt,name=system_info" json:"system_info,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CPlayerClient_GetSystemInformation_Response) Reset() {
	*m = CPlayerClient_GetSystemInformation_Response{}
}
func (m *CPlayerClient_GetSystemInformation_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CPlayerClient_GetSystemInformation_Response) ProtoMessage() {}
func (*CPlayerClient_GetSystemInformation_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{9}
}

func (m *CPlayerClient_GetSystemInformation_Response) GetSystemInfo() *CClientSystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
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
	proto.RegisterType((*CPlayer_LastPlayedTimes_Notification)(nil), "CPlayer_LastPlayedTimes_Notification")
	proto.RegisterType((*CPlayerClient_GetSystemInformation_Request)(nil), "CPlayerClient_GetSystemInformation_Request")
	proto.RegisterType((*CClientSystemInfo)(nil), "CClientSystemInfo")
	proto.RegisterType((*CClientSystemInfo_CPU)(nil), "CClientSystemInfo.CPU")
	proto.RegisterType((*CClientSystemInfo_VideoCard)(nil), "CClientSystemInfo.VideoCard")
	proto.RegisterType((*CPlayerClient_GetSystemInformation_Response)(nil), "CPlayerClient_GetSystemInformation_Response")
}

var fileDescriptor13 = []byte{
	// 1067 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x05, 0x2d, 0xcb, 0x96, 0x46, 0x76, 0x9c, 0x6c, 0x9c, 0x94, 0xa1, 0x5d, 0x60, 0x41, 0xa7,
	0xf9, 0x70, 0x1c, 0xa2, 0x55, 0x82, 0xa2, 0x68, 0x0b, 0x04, 0xb6, 0x0f, 0x41, 0x01, 0x23, 0x48,
	0xe3, 0x38, 0xa7, 0x02, 0xec, 0x8a, 0x5c, 0xc9, 0x44, 0x48, 0xae, 0xca, 0x5d, 0x29, 0x55, 0x4f,
	0x45, 0xce, 0xbd, 0x15, 0xfd, 0x1b, 0xbd, 0xb9, 0xe7, 0x00, 0xfd, 0x23, 0xfd, 0x03, 0x3d, 0xf6,
	0xde, 0xd9, 0x5d, 0x52, 0x1f, 0x96, 0x9d, 0xaa, 0x39, 0x18, 0xd6, 0xee, 0xbc, 0x1d, 0xbe, 0x79,
	0x3b, 0xf3, 0x16, 0xee, 0x48, 0xc5, 0x59, 0x96, 0x71, 0x29, 0x59, 0x8f, 0xcb, 0xb0, 0x9f, 0xb2,
	0x11, 0x2f, 0x02, 0xb3, 0x19, 0xa5, 0x09, 0xcf, 0x55, 0xd0, 0x2f, 0x84, 0x12, 0xde, 0xde, 0x2c,
	0x6e, 0x90, 0x27, 0xdd, 0x84, 0xc7, 0x61, 0x87, 0x49, 0x3e, 0x8f, 0xf6, 0x1f, 0x81, 0x7f, 0xf8,
	0xdc, 0xa4, 0x0a, 0x9f, 0x72, 0xf5, 0x94, 0x65, 0xfc, 0x80, 0xc5, 0x3d, 0x7e, 0xc4, 0x87, 0x3c,
	0x95, 0xe1, 0x0b, 0xfe, 0xc3, 0x80, 0x4b, 0x45, 0xd6, 0xa1, 0xce, 0xfa, 0xfd, 0x24, 0x76, 0x1d,
	0xea, 0xdc, 0x5b, 0xf7, 0xcf, 0x1c, 0xd8, 0x79, 0xef, 0x29, 0xd9, 0x17, 0xb9, 0xe4, 0x64, 0x13,
	0xd6, 0x2c, 0xcd, 0x30, 0xd5, 0x11, 0x7b, 0x9a, 0x7c, 0x09, 0x2b, 0x1d, 0x8d, 0x96, 0xee, 0x12,
	0xad, 0xdd, 0x6b, 0xb5, 0x77, 0x83, 0x05, 0x72, 0x05, 0x66, 0xd3, 0xfb, 0x1a, 0xea, 0xe6, 0x87,
	0x66, 0x34, 0xc9, 0x59, 0x27, 0x57, 0x60, 0x45, 0xf2, 0x22, 0x31, 0x39, 0xf5, 0x1a, 0xbf, 0xdc,
	0x11, 0x45, 0x8c, 0x39, 0x23, 0x91, 0x8a, 0xc2, 0xad, 0x19, 0xde, 0x6f, 0x9d, 0x99, 0x6a, 0x8f,
	0x98, 0x54, 0x66, 0x15, 0xbf, 0x4c, 0x50, 0xaf, 0x71, 0xb5, 0xdf, 0xc1, 0x46, 0x96, 0xe4, 0x61,
	0x8a, 0x61, 0x2b, 0x73, 0x59, 0xf7, 0xc1, 0xe1, 0xdb, 0x33, 0xf7, 0xc9, 0xcb, 0x53, 0x4e, 0x33,
	0x21, 0x15, 0x2d, 0x78, 0x84, 0x3a, 0x52, 0x0d, 0x7b, 0x68, 0x61, 0x54, 0x61, 0x1e, 0xaa, 0x10,
	0x60, 0x35, 0xa6, 0x2c, 0x2d, 0x38, 0x8b, 0x47, 0xf4, 0x75, 0x2e, 0xde, 0x48, 0xca, 0x3a, 0x62,
	0xa0, 0xfc, 0x77, 0xb3, 0xe2, 0xcd, 0x93, 0x28, 0xc5, 0xfb, 0x02, 0xea, 0x3d, 0x14, 0x43, 0xe2,
	0xb7, 0xb5, 0x4a, 0xf7, 0x83, 0x05, 0x0e, 0x05, 0x5a, 0x3e, 0x2f, 0x84, 0x65, 0xfd, 0x7f, 0xf6,
	0xd6, 0xea, 0xe4, 0x06, 0xac, 0x8f, 0x4b, 0xd2, 0x44, 0x8d, 0x54, 0xeb, 0xe4, 0x23, 0xd8, 0xa8,
	0x76, 0xc2, 0xf6, 0x1b, 0xce, 0x5f, 0x4b, 0xa3, 0x56, 0x9d, 0xb8, 0x70, 0x75, 0x1c, 0xe8, 0x8a,
	0x02, 0xd5, 0x2e, 0xdc, 0x65, 0x1d, 0xf1, 0xb7, 0xe0, 0x56, 0x45, 0x66, 0x3f, 0x8a, 0x78, 0x5f,
	0x1d, 0x1f, 0xef, 0x57, 0xea, 0xf9, 0xdb, 0xe0, 0x5d, 0x14, 0xb4, 0x04, 0xfd, 0xef, 0xe1, 0x76,
	0x15, 0x3d, 0x5f, 0xc4, 0x33, 0xa1, 0xb0, 0x55, 0x23, 0xa6, 0x12, 0x91, 0x7f, 0x78, 0xf5, 0xfe,
	0x1e, 0xec, 0x96, 0xd8, 0x43, 0x73, 0x09, 0xfa, 0xc4, 0xf1, 0x08, 0x3b, 0x3f, 0xfb, 0x26, 0xc7,
	0x32, 0x32, 0x93, 0x7f, 0xcc, 0xf6, 0x9f, 0x65, 0xb8, 0x76, 0x68, 0x81, 0x13, 0x10, 0xd9, 0x81,
	0x5a, 0xd4, 0x1f, 0x18, 0xdd, 0x5a, 0xed, 0x9b, 0xc1, 0x1c, 0x00, 0xd9, 0x9c, 0x90, 0x4f, 0x01,
	0x86, 0x49, 0xcc, 0x45, 0x18, 0xb1, 0x22, 0x36, 0x62, 0xb6, 0xda, 0xdb, 0x17, 0x60, 0x5f, 0x69,
	0xd0, 0x21, 0x62, 0xb4, 0xa2, 0xa2, 0xcf, 0x0b, 0x64, 0x90, 0xf7, 0x42, 0x69, 0x10, 0x46, 0xeb,
	0x26, 0xb9, 0x0a, 0x0d, 0x21, 0xc3, 0xcf, 0x1f, 0x77, 0x12, 0x65, 0x34, 0x6e, 0xe8, 0xdb, 0xb2,
	0x88, 0xb0, 0x60, 0x59, 0x98, 0x75, 0xdc, 0x7a, 0xd5, 0xd8, 0x6c, 0x10, 0x27, 0x22, 0x8c, 0xf9,
	0x30, 0x89, 0xb8, 0xbb, 0x62, 0x8e, 0x6f, 0xc3, 0x66, 0xb9, 0x5b, 0x24, 0x78, 0x4d, 0x21, 0xfe,
	0x49, 0xac, 0xd2, 0x5d, 0xd5, 0x51, 0xef, 0x2f, 0x07, 0x6a, 0x9a, 0xf0, 0x35, 0x68, 0xca, 0x3e,
	0x47, 0x2f, 0xc8, 0x4e, 0x7f, 0x9a, 0xcc, 0xcd, 0x90, 0xe7, 0x31, 0x4e, 0xc8, 0x92, 0x49, 0xe4,
	0x01, 0x49, 0x45, 0x0f, 0x2f, 0x21, 0x0d, 0xd1, 0x1f, 0x22, 0x34, 0x11, 0x51, 0x54, 0xfd, 0xb0,
	0x05, 0xd7, 0xfb, 0xa7, 0x23, 0x79, 0x3e, 0x68, 0x5a, 0x82, 0xdc, 0x84, 0x2b, 0xa7, 0x23, 0xac,
	0x4d, 0x9d, 0xea, 0x8e, 0xc7, 0xfa, 0x0c, 0xdf, 0x86, 0xee, 0xc1, 0x6e, 0x94, 0x89, 0xa1, 0x21,
	0xda, 0x20, 0x6b, 0xb0, 0x2c, 0x25, 0x6f, 0x1b, 0x62, 0xd5, 0xea, 0x91, 0xdb, 0xa8, 0xa0, 0xd2,
	0x2c, 0x9b, 0x93, 0x25, 0x7f, 0xcc, 0x5c, 0x98, 0x5e, 0x7e, 0xe6, 0xb6, 0xa6, 0x97, 0x6d, 0x77,
	0x4d, 0x2f, 0xbd, 0xdf, 0x1d, 0x68, 0x4e, 0x74, 0xc6, 0xaa, 0xac, 0x10, 0xa6, 0xca, 0xa6, 0x26,
	0x77, 0x4e, 0x18, 0x5b, 0xed, 0x75, 0x68, 0x95, 0xfb, 0x31, 0x53, 0xdc, 0x9a, 0x84, 0x9e, 0x87,
	0x38, 0xc1, 0xd9, 0x56, 0x3f, 0x8e, 0xd1, 0xcb, 0x55, 0x16, 0xbc, 0xbd, 0xbc, 0x97, 0x8e, 0xf7,
	0xeb, 0xd5, 0xdd, 0x59, 0x0d, 0x71, 0xd2, 0x56, 0x8c, 0x18, 0xb8, 0x63, 0xaf, 0x07, 0x77, 0x56,
	0xcd, 0xce, 0x06, 0xac, 0x0e, 0xcb, 0x7b, 0x6c, 0x98, 0x11, 0x7a, 0x05, 0x0f, 0x16, 0xea, 0xd2,
	0xd2, 0x0c, 0xee, 0x42, 0xab, 0xec, 0x86, 0x04, 0xc3, 0x65, 0x63, 0x92, 0xf9, 0x66, 0x6b, 0xff,
	0x5d, 0x83, 0x15, 0x9b, 0x97, 0xfc, 0xe1, 0x00, 0x99, 0x77, 0x54, 0xb2, 0x13, 0xfc, 0xb7, 0xe1,
	0x7b, 0xb7, 0x17, 0xf1, 0x64, 0xff, 0x04, 0xfd, 0xf0, 0xdb, 0x17, 0x5c, 0x0d, 0x8a, 0x5c, 0x1a,
	0xdb, 0x3b, 0xd6, 0xef, 0x0b, 0x35, 0x30, 0x2a, 0xba, 0x94, 0xd1, 0x01, 0xba, 0xf2, 0x9e, 0x09,
	0x99, 0x04, 0xd4, 0x78, 0x36, 0xc5, 0x02, 0xcd, 0x9e, 0x9e, 0xf0, 0x3d, 0xca, 0xf2, 0x98, 0x26,
	0x5d, 0x9a, 0xa8, 0xbb, 0x12, 0x23, 0x49, 0x4a, 0x7e, 0x73, 0xc0, 0xb5, 0x85, 0xcd, 0x0f, 0xfb,
	0x2c, 0xfd, 0x4b, 0x1c, 0x7c, 0x96, 0xfe, 0x65, 0x76, 0xe1, 0x07, 0x48, 0x7f, 0x17, 0x01, 0x96,
	0xfb, 0x79, 0x1f, 0x97, 0x63, 0x9a, 0x2c, 0x8a, 0xc4, 0x20, 0x57, 0x24, 0x82, 0xe6, 0xd8, 0xd1,
	0x88, 0x17, 0x5c, 0x6a, 0x81, 0xde, 0x56, 0xf0, 0x1e, 0x07, 0xfc, 0x18, 0xbf, 0x7a, 0xeb, 0x04,
	0x75, 0xa1, 0x89, 0xd4, 0xa9, 0x31, 0x8e, 0xe3, 0x62, 0xe5, 0x3b, 0xde, 0xf7, 0x1e, 0x62, 0xf8,
	0xfe, 0x3e, 0xc5, 0xb8, 0x6e, 0x20, 0xc3, 0x41, 0x83, 0xa4, 0xd4, 0x20, 0xab, 0xaf, 0x7d, 0x53,
	0x29, 0xf6, 0x2c, 0x6b, 0xff, 0x5a, 0x83, 0xb5, 0xe9, 0x3e, 0x22, 0xbf, 0x38, 0x70, 0xc3, 0x38,
	0xe9, 0xe8, 0xbc, 0x72, 0x9f, 0x04, 0x8b, 0x38, 0xaf, 0xd7, 0x0a, 0x9e, 0x89, 0x31, 0xd9, 0x27,
	0xc8, 0xe6, 0xab, 0xe9, 0x30, 0xed, 0x16, 0x22, 0x33, 0xec, 0x90, 0x81, 0x12, 0xd5, 0x43, 0x87,
	0x97, 0x9d, 0xe1, 0x13, 0x51, 0xbd, 0x89, 0x9a, 0xa1, 0x11, 0x91, 0xfc, 0xe9, 0xc0, 0xe6, 0x45,
	0xad, 0x4d, 0x1e, 0x04, 0x8b, 0xbb, 0xb4, 0xb7, 0x17, 0xfc, 0x8f, 0x61, 0xf1, 0x9f, 0x23, 0xe9,
	0xa3, 0xf2, 0xa8, 0xe5, 0xab, 0xc5, 0x9d, 0x70, 0x9e, 0x7a, 0xa0, 0xb5, 0xbc, 0xc9, 0x24, 0x87,
	0x7d, 0xa2, 0x0d, 0x40, 0x77, 0xae, 0xa4, 0x76, 0xe8, 0x3c, 0xdd, 0x29, 0x77, 0xe6, 0xc4, 0xaf,
	0x92, 0xe4, 0x53, 0xfa, 0xc8, 0x77, 0x67, 0xee, 0xd2, 0x41, 0xed, 0x67, 0xc7, 0xf9, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0xe2, 0xe2, 0x67, 0xb1, 0x09, 0x00, 0x00,
}
