// Code generated by protoc-gen-go.
// source: steammessages_gameservers.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CGameServers_GetServerList_Request struct {
	Filter           *string `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Limit            *uint32 `protobuf:"varint,2,opt,name=limit,def=100" json:"limit,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGameServers_GetServerList_Request) Reset()         { *m = CGameServers_GetServerList_Request{} }
func (m *CGameServers_GetServerList_Request) String() string { return proto.CompactTextString(m) }
func (*CGameServers_GetServerList_Request) ProtoMessage()    {}

const Default_CGameServers_GetServerList_Request_Limit uint32 = 100

func (m *CGameServers_GetServerList_Request) GetFilter() string {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return ""
}

func (m *CGameServers_GetServerList_Request) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return Default_CGameServers_GetServerList_Request_Limit
}

type CGameServers_GetServerList_Response struct {
	Servers          []*CGameServers_GetServerList_Response_Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
	XXX_unrecognized []byte                                        `json:"-"`
}

func (m *CGameServers_GetServerList_Response) Reset()         { *m = CGameServers_GetServerList_Response{} }
func (m *CGameServers_GetServerList_Response) String() string { return proto.CompactTextString(m) }
func (*CGameServers_GetServerList_Response) ProtoMessage()    {}

func (m *CGameServers_GetServerList_Response) GetServers() []*CGameServers_GetServerList_Response_Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type CGameServers_GetServerList_Response_Server struct {
	Addr             *string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Gameport         *uint32 `protobuf:"varint,2,opt,name=gameport" json:"gameport,omitempty"`
	Specport         *uint32 `protobuf:"varint,3,opt,name=specport" json:"specport,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,4,opt,name=steamid" json:"steamid,omitempty"`
	Name             *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Appid            *uint32 `protobuf:"varint,6,opt,name=appid" json:"appid,omitempty"`
	Gamedir          *string `protobuf:"bytes,7,opt,name=gamedir" json:"gamedir,omitempty"`
	Version          *string `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	Product          *string `protobuf:"bytes,9,opt,name=product" json:"product,omitempty"`
	Region           *int32  `protobuf:"varint,10,opt,name=region" json:"region,omitempty"`
	Players          *int32  `protobuf:"varint,11,opt,name=players" json:"players,omitempty"`
	MaxPlayers       *int32  `protobuf:"varint,12,opt,name=max_players" json:"max_players,omitempty"`
	Bots             *int32  `protobuf:"varint,13,opt,name=bots" json:"bots,omitempty"`
	Map              *string `protobuf:"bytes,14,opt,name=map" json:"map,omitempty"`
	Secure           *bool   `protobuf:"varint,15,opt,name=secure" json:"secure,omitempty"`
	Dedicated        *bool   `protobuf:"varint,16,opt,name=dedicated" json:"dedicated,omitempty"`
	Os               *string `protobuf:"bytes,17,opt,name=os" json:"os,omitempty"`
	Gametype         *string `protobuf:"bytes,18,opt,name=gametype" json:"gametype,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGameServers_GetServerList_Response_Server) Reset() {
	*m = CGameServers_GetServerList_Response_Server{}
}
func (m *CGameServers_GetServerList_Response_Server) String() string {
	return proto.CompactTextString(m)
}
func (*CGameServers_GetServerList_Response_Server) ProtoMessage() {}

func (m *CGameServers_GetServerList_Response_Server) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetGameport() uint32 {
	if m != nil && m.Gameport != nil {
		return *m.Gameport
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetSpecport() uint32 {
	if m != nil && m.Specport != nil {
		return *m.Specport
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetGamedir() string {
	if m != nil && m.Gamedir != nil {
		return *m.Gamedir
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetRegion() int32 {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetPlayers() int32 {
	if m != nil && m.Players != nil {
		return *m.Players
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetMaxPlayers() int32 {
	if m != nil && m.MaxPlayers != nil {
		return *m.MaxPlayers
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetBots() int32 {
	if m != nil && m.Bots != nil {
		return *m.Bots
	}
	return 0
}

func (m *CGameServers_GetServerList_Response_Server) GetMap() string {
	if m != nil && m.Map != nil {
		return *m.Map
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetSecure() bool {
	if m != nil && m.Secure != nil {
		return *m.Secure
	}
	return false
}

func (m *CGameServers_GetServerList_Response_Server) GetDedicated() bool {
	if m != nil && m.Dedicated != nil {
		return *m.Dedicated
	}
	return false
}

func (m *CGameServers_GetServerList_Response_Server) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *CGameServers_GetServerList_Response_Server) GetGametype() string {
	if m != nil && m.Gametype != nil {
		return *m.Gametype
	}
	return ""
}

type CGameServers_GetServerSteamIDsByIP_Request struct {
	ServerIps        []string `protobuf:"bytes,1,rep,name=server_ips" json:"server_ips,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CGameServers_GetServerSteamIDsByIP_Request) Reset() {
	*m = CGameServers_GetServerSteamIDsByIP_Request{}
}
func (m *CGameServers_GetServerSteamIDsByIP_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CGameServers_GetServerSteamIDsByIP_Request) ProtoMessage() {}

func (m *CGameServers_GetServerSteamIDsByIP_Request) GetServerIps() []string {
	if m != nil {
		return m.ServerIps
	}
	return nil
}

type CGameServers_IPsWithSteamIDs_Response struct {
	Servers          []*CGameServers_IPsWithSteamIDs_Response_Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *CGameServers_IPsWithSteamIDs_Response) Reset()         { *m = CGameServers_IPsWithSteamIDs_Response{} }
func (m *CGameServers_IPsWithSteamIDs_Response) String() string { return proto.CompactTextString(m) }
func (*CGameServers_IPsWithSteamIDs_Response) ProtoMessage()    {}

func (m *CGameServers_IPsWithSteamIDs_Response) GetServers() []*CGameServers_IPsWithSteamIDs_Response_Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type CGameServers_IPsWithSteamIDs_Response_Server struct {
	Addr             *string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,2,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGameServers_IPsWithSteamIDs_Response_Server) Reset() {
	*m = CGameServers_IPsWithSteamIDs_Response_Server{}
}
func (m *CGameServers_IPsWithSteamIDs_Response_Server) String() string {
	return proto.CompactTextString(m)
}
func (*CGameServers_IPsWithSteamIDs_Response_Server) ProtoMessage() {}

func (m *CGameServers_IPsWithSteamIDs_Response_Server) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *CGameServers_IPsWithSteamIDs_Response_Server) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CGameServers_GetServerIPsBySteamID_Request struct {
	ServerSteamids   []uint64 `protobuf:"fixed64,1,rep,name=server_steamids" json:"server_steamids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CGameServers_GetServerIPsBySteamID_Request) Reset() {
	*m = CGameServers_GetServerIPsBySteamID_Request{}
}
func (m *CGameServers_GetServerIPsBySteamID_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CGameServers_GetServerIPsBySteamID_Request) ProtoMessage() {}

func (m *CGameServers_GetServerIPsBySteamID_Request) GetServerSteamids() []uint64 {
	if m != nil {
		return m.ServerSteamids
	}
	return nil
}
