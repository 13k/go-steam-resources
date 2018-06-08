// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steammessages_base.proto

package steam // import "github.com/13k/go-steam-resources/protobuf/steam"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CMsgProtoBufHeader struct {
	Steamid              *uint64  `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	ClientSessionid      *int32   `protobuf:"varint,2,opt,name=client_sessionid,json=clientSessionid" json:"client_sessionid,omitempty"`
	RoutingAppid         *uint32  `protobuf:"varint,3,opt,name=routing_appid,json=routingAppid" json:"routing_appid,omitempty"`
	JobidSource          *uint64  `protobuf:"fixed64,10,opt,name=jobid_source,json=jobidSource,def=18446744073709551615" json:"jobid_source,omitempty"`
	JobidTarget          *uint64  `protobuf:"fixed64,11,opt,name=jobid_target,json=jobidTarget,def=18446744073709551615" json:"jobid_target,omitempty"`
	TargetJobName        *string  `protobuf:"bytes,12,opt,name=target_job_name,json=targetJobName" json:"target_job_name,omitempty"`
	SeqNum               *int32   `protobuf:"varint,24,opt,name=seq_num,json=seqNum" json:"seq_num,omitempty"`
	Eresult              *int32   `protobuf:"varint,13,opt,name=eresult,def=2" json:"eresult,omitempty"`
	ErrorMessage         *string  `protobuf:"bytes,14,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	Ip                   *uint32  `protobuf:"varint,15,opt,name=ip" json:"ip,omitempty"`
	AuthAccountFlags     *uint32  `protobuf:"varint,16,opt,name=auth_account_flags,json=authAccountFlags" json:"auth_account_flags,omitempty"`
	TokenSource          *uint32  `protobuf:"varint,22,opt,name=token_source,json=tokenSource" json:"token_source,omitempty"`
	AdminSpoofingUser    *bool    `protobuf:"varint,23,opt,name=admin_spoofing_user,json=adminSpoofingUser" json:"admin_spoofing_user,omitempty"`
	TransportError       *int32   `protobuf:"varint,17,opt,name=transport_error,json=transportError,def=1" json:"transport_error,omitempty"`
	Messageid            *uint64  `protobuf:"varint,18,opt,name=messageid,def=18446744073709551615" json:"messageid,omitempty"`
	PublisherGroupId     *uint32  `protobuf:"varint,19,opt,name=publisher_group_id,json=publisherGroupId" json:"publisher_group_id,omitempty"`
	Sysid                *uint32  `protobuf:"varint,20,opt,name=sysid" json:"sysid,omitempty"`
	TraceTag             *uint64  `protobuf:"varint,21,opt,name=trace_tag,json=traceTag" json:"trace_tag,omitempty"`
	WebapiKeyId          *uint32  `protobuf:"varint,25,opt,name=webapi_key_id,json=webapiKeyId" json:"webapi_key_id,omitempty"`
	IsFromExternalSource *bool    `protobuf:"varint,26,opt,name=is_from_external_source,json=isFromExternalSource" json:"is_from_external_source,omitempty"`
	ForwardToSysid       []uint32 `protobuf:"varint,27,rep,name=forward_to_sysid,json=forwardToSysid" json:"forward_to_sysid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgProtoBufHeader) Reset()         { *m = CMsgProtoBufHeader{} }
func (m *CMsgProtoBufHeader) String() string { return proto.CompactTextString(m) }
func (*CMsgProtoBufHeader) ProtoMessage()    {}
func (*CMsgProtoBufHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{0}
}
func (m *CMsgProtoBufHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgProtoBufHeader.Unmarshal(m, b)
}
func (m *CMsgProtoBufHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgProtoBufHeader.Marshal(b, m, deterministic)
}
func (dst *CMsgProtoBufHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgProtoBufHeader.Merge(dst, src)
}
func (m *CMsgProtoBufHeader) XXX_Size() int {
	return xxx_messageInfo_CMsgProtoBufHeader.Size(m)
}
func (m *CMsgProtoBufHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgProtoBufHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgProtoBufHeader proto.InternalMessageInfo

const Default_CMsgProtoBufHeader_JobidSource uint64 = 18446744073709551615
const Default_CMsgProtoBufHeader_JobidTarget uint64 = 18446744073709551615
const Default_CMsgProtoBufHeader_Eresult int32 = 2
const Default_CMsgProtoBufHeader_TransportError int32 = 1
const Default_CMsgProtoBufHeader_Messageid uint64 = 18446744073709551615

func (m *CMsgProtoBufHeader) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetClientSessionid() int32 {
	if m != nil && m.ClientSessionid != nil {
		return *m.ClientSessionid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetRoutingAppid() uint32 {
	if m != nil && m.RoutingAppid != nil {
		return *m.RoutingAppid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetJobidSource() uint64 {
	if m != nil && m.JobidSource != nil {
		return *m.JobidSource
	}
	return Default_CMsgProtoBufHeader_JobidSource
}

func (m *CMsgProtoBufHeader) GetJobidTarget() uint64 {
	if m != nil && m.JobidTarget != nil {
		return *m.JobidTarget
	}
	return Default_CMsgProtoBufHeader_JobidTarget
}

func (m *CMsgProtoBufHeader) GetTargetJobName() string {
	if m != nil && m.TargetJobName != nil {
		return *m.TargetJobName
	}
	return ""
}

func (m *CMsgProtoBufHeader) GetSeqNum() int32 {
	if m != nil && m.SeqNum != nil {
		return *m.SeqNum
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgProtoBufHeader_Eresult
}

func (m *CMsgProtoBufHeader) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *CMsgProtoBufHeader) GetIp() uint32 {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetAuthAccountFlags() uint32 {
	if m != nil && m.AuthAccountFlags != nil {
		return *m.AuthAccountFlags
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetTokenSource() uint32 {
	if m != nil && m.TokenSource != nil {
		return *m.TokenSource
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetAdminSpoofingUser() bool {
	if m != nil && m.AdminSpoofingUser != nil {
		return *m.AdminSpoofingUser
	}
	return false
}

func (m *CMsgProtoBufHeader) GetTransportError() int32 {
	if m != nil && m.TransportError != nil {
		return *m.TransportError
	}
	return Default_CMsgProtoBufHeader_TransportError
}

func (m *CMsgProtoBufHeader) GetMessageid() uint64 {
	if m != nil && m.Messageid != nil {
		return *m.Messageid
	}
	return Default_CMsgProtoBufHeader_Messageid
}

func (m *CMsgProtoBufHeader) GetPublisherGroupId() uint32 {
	if m != nil && m.PublisherGroupId != nil {
		return *m.PublisherGroupId
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetSysid() uint32 {
	if m != nil && m.Sysid != nil {
		return *m.Sysid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetTraceTag() uint64 {
	if m != nil && m.TraceTag != nil {
		return *m.TraceTag
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetWebapiKeyId() uint32 {
	if m != nil && m.WebapiKeyId != nil {
		return *m.WebapiKeyId
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetIsFromExternalSource() bool {
	if m != nil && m.IsFromExternalSource != nil {
		return *m.IsFromExternalSource
	}
	return false
}

func (m *CMsgProtoBufHeader) GetForwardToSysid() []uint32 {
	if m != nil {
		return m.ForwardToSysid
	}
	return nil
}

type CMsgMulti struct {
	SizeUnzipped         *uint32  `protobuf:"varint,1,opt,name=size_unzipped,json=sizeUnzipped" json:"size_unzipped,omitempty"`
	MessageBody          []byte   `protobuf:"bytes,2,opt,name=message_body,json=messageBody" json:"message_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgMulti) Reset()         { *m = CMsgMulti{} }
func (m *CMsgMulti) String() string { return proto.CompactTextString(m) }
func (*CMsgMulti) ProtoMessage()    {}
func (*CMsgMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{1}
}
func (m *CMsgMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgMulti.Unmarshal(m, b)
}
func (m *CMsgMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgMulti.Marshal(b, m, deterministic)
}
func (dst *CMsgMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgMulti.Merge(dst, src)
}
func (m *CMsgMulti) XXX_Size() int {
	return xxx_messageInfo_CMsgMulti.Size(m)
}
func (m *CMsgMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgMulti.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgMulti proto.InternalMessageInfo

func (m *CMsgMulti) GetSizeUnzipped() uint32 {
	if m != nil && m.SizeUnzipped != nil {
		return *m.SizeUnzipped
	}
	return 0
}

func (m *CMsgMulti) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

type CMsgProtobufWrapped struct {
	MessageBody          []byte   `protobuf:"bytes,1,opt,name=message_body,json=messageBody" json:"message_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgProtobufWrapped) Reset()         { *m = CMsgProtobufWrapped{} }
func (m *CMsgProtobufWrapped) String() string { return proto.CompactTextString(m) }
func (*CMsgProtobufWrapped) ProtoMessage()    {}
func (*CMsgProtobufWrapped) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{2}
}
func (m *CMsgProtobufWrapped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgProtobufWrapped.Unmarshal(m, b)
}
func (m *CMsgProtobufWrapped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgProtobufWrapped.Marshal(b, m, deterministic)
}
func (dst *CMsgProtobufWrapped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgProtobufWrapped.Merge(dst, src)
}
func (m *CMsgProtobufWrapped) XXX_Size() int {
	return xxx_messageInfo_CMsgProtobufWrapped.Size(m)
}
func (m *CMsgProtobufWrapped) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgProtobufWrapped.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgProtobufWrapped proto.InternalMessageInfo

func (m *CMsgProtobufWrapped) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

type CMsgAuthTicket struct {
	Estate               *uint32  `protobuf:"varint,1,opt,name=estate" json:"estate,omitempty"`
	Eresult              *uint32  `protobuf:"varint,2,opt,name=eresult,def=2" json:"eresult,omitempty"`
	Steamid              *uint64  `protobuf:"fixed64,3,opt,name=steamid" json:"steamid,omitempty"`
	Gameid               *uint64  `protobuf:"fixed64,4,opt,name=gameid" json:"gameid,omitempty"`
	HSteamPipe           *uint32  `protobuf:"varint,5,opt,name=h_steam_pipe,json=hSteamPipe" json:"h_steam_pipe,omitempty"`
	TicketCrc            *uint32  `protobuf:"varint,6,opt,name=ticket_crc,json=ticketCrc" json:"ticket_crc,omitempty"`
	Ticket               []byte   `protobuf:"bytes,7,opt,name=ticket" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMsgAuthTicket) Reset()         { *m = CMsgAuthTicket{} }
func (m *CMsgAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgAuthTicket) ProtoMessage()    {}
func (*CMsgAuthTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{3}
}
func (m *CMsgAuthTicket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgAuthTicket.Unmarshal(m, b)
}
func (m *CMsgAuthTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgAuthTicket.Marshal(b, m, deterministic)
}
func (dst *CMsgAuthTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgAuthTicket.Merge(dst, src)
}
func (m *CMsgAuthTicket) XXX_Size() int {
	return xxx_messageInfo_CMsgAuthTicket.Size(m)
}
func (m *CMsgAuthTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgAuthTicket.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgAuthTicket proto.InternalMessageInfo

const Default_CMsgAuthTicket_Eresult uint32 = 2

func (m *CMsgAuthTicket) GetEstate() uint32 {
	if m != nil && m.Estate != nil {
		return *m.Estate
	}
	return 0
}

func (m *CMsgAuthTicket) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgAuthTicket_Eresult
}

func (m *CMsgAuthTicket) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CMsgAuthTicket) GetGameid() uint64 {
	if m != nil && m.Gameid != nil {
		return *m.Gameid
	}
	return 0
}

func (m *CMsgAuthTicket) GetHSteamPipe() uint32 {
	if m != nil && m.HSteamPipe != nil {
		return *m.HSteamPipe
	}
	return 0
}

func (m *CMsgAuthTicket) GetTicketCrc() uint32 {
	if m != nil && m.TicketCrc != nil {
		return *m.TicketCrc
	}
	return 0
}

func (m *CMsgAuthTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type CCDDBAppDetailCommon struct {
	Appid                 *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Name                  *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Icon                  *string  `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	Logo                  *string  `protobuf:"bytes,4,opt,name=logo" json:"logo,omitempty"`
	LogoSmall             *string  `protobuf:"bytes,5,opt,name=logo_small,json=logoSmall" json:"logo_small,omitempty"`
	Tool                  *bool    `protobuf:"varint,6,opt,name=tool" json:"tool,omitempty"`
	Demo                  *bool    `protobuf:"varint,7,opt,name=demo" json:"demo,omitempty"`
	Media                 *bool    `protobuf:"varint,8,opt,name=media" json:"media,omitempty"`
	CommunityVisibleStats *bool    `protobuf:"varint,9,opt,name=community_visible_stats,json=communityVisibleStats" json:"community_visible_stats,omitempty"`
	FriendlyName          *string  `protobuf:"bytes,10,opt,name=friendly_name,json=friendlyName" json:"friendly_name,omitempty"`
	Propagation           *string  `protobuf:"bytes,11,opt,name=propagation" json:"propagation,omitempty"`
	HasAdultContent       *bool    `protobuf:"varint,12,opt,name=has_adult_content,json=hasAdultContent" json:"has_adult_content,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *CCDDBAppDetailCommon) Reset()         { *m = CCDDBAppDetailCommon{} }
func (m *CCDDBAppDetailCommon) String() string { return proto.CompactTextString(m) }
func (*CCDDBAppDetailCommon) ProtoMessage()    {}
func (*CCDDBAppDetailCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{4}
}
func (m *CCDDBAppDetailCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCDDBAppDetailCommon.Unmarshal(m, b)
}
func (m *CCDDBAppDetailCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCDDBAppDetailCommon.Marshal(b, m, deterministic)
}
func (dst *CCDDBAppDetailCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCDDBAppDetailCommon.Merge(dst, src)
}
func (m *CCDDBAppDetailCommon) XXX_Size() int {
	return xxx_messageInfo_CCDDBAppDetailCommon.Size(m)
}
func (m *CCDDBAppDetailCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_CCDDBAppDetailCommon.DiscardUnknown(m)
}

var xxx_messageInfo_CCDDBAppDetailCommon proto.InternalMessageInfo

func (m *CCDDBAppDetailCommon) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCDDBAppDetailCommon) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetLogo() string {
	if m != nil && m.Logo != nil {
		return *m.Logo
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetLogoSmall() string {
	if m != nil && m.LogoSmall != nil {
		return *m.LogoSmall
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetTool() bool {
	if m != nil && m.Tool != nil {
		return *m.Tool
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetDemo() bool {
	if m != nil && m.Demo != nil {
		return *m.Demo
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetMedia() bool {
	if m != nil && m.Media != nil {
		return *m.Media
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetCommunityVisibleStats() bool {
	if m != nil && m.CommunityVisibleStats != nil {
		return *m.CommunityVisibleStats
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetFriendlyName() string {
	if m != nil && m.FriendlyName != nil {
		return *m.FriendlyName
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetPropagation() string {
	if m != nil && m.Propagation != nil {
		return *m.Propagation
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetHasAdultContent() bool {
	if m != nil && m.HasAdultContent != nil {
		return *m.HasAdultContent
	}
	return false
}

type CMsgAppRights struct {
	EditInfo                 *bool    `protobuf:"varint,1,opt,name=edit_info,json=editInfo" json:"edit_info,omitempty"`
	Publish                  *bool    `protobuf:"varint,2,opt,name=publish" json:"publish,omitempty"`
	ViewErrorData            *bool    `protobuf:"varint,3,opt,name=view_error_data,json=viewErrorData" json:"view_error_data,omitempty"`
	Download                 *bool    `protobuf:"varint,4,opt,name=download" json:"download,omitempty"`
	UploadCdkeys             *bool    `protobuf:"varint,5,opt,name=upload_cdkeys,json=uploadCdkeys" json:"upload_cdkeys,omitempty"`
	GenerateCdkeys           *bool    `protobuf:"varint,6,opt,name=generate_cdkeys,json=generateCdkeys" json:"generate_cdkeys,omitempty"`
	ViewFinancials           *bool    `protobuf:"varint,7,opt,name=view_financials,json=viewFinancials" json:"view_financials,omitempty"`
	ManageCeg                *bool    `protobuf:"varint,8,opt,name=manage_ceg,json=manageCeg" json:"manage_ceg,omitempty"`
	ManageSigning            *bool    `protobuf:"varint,9,opt,name=manage_signing,json=manageSigning" json:"manage_signing,omitempty"`
	ManageCdkeys             *bool    `protobuf:"varint,10,opt,name=manage_cdkeys,json=manageCdkeys" json:"manage_cdkeys,omitempty"`
	EditMarketing            *bool    `protobuf:"varint,11,opt,name=edit_marketing,json=editMarketing" json:"edit_marketing,omitempty"`
	EconomySupport           *bool    `protobuf:"varint,12,opt,name=economy_support,json=economySupport" json:"economy_support,omitempty"`
	EconomySupportSupervisor *bool    `protobuf:"varint,13,opt,name=economy_support_supervisor,json=economySupportSupervisor" json:"economy_support_supervisor,omitempty"`
	ManagePricing            *bool    `protobuf:"varint,14,opt,name=manage_pricing,json=managePricing" json:"manage_pricing,omitempty"`
	BroadcastLive            *bool    `protobuf:"varint,15,opt,name=broadcast_live,json=broadcastLive" json:"broadcast_live,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *CMsgAppRights) Reset()         { *m = CMsgAppRights{} }
func (m *CMsgAppRights) String() string { return proto.CompactTextString(m) }
func (*CMsgAppRights) ProtoMessage()    {}
func (*CMsgAppRights) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{5}
}
func (m *CMsgAppRights) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgAppRights.Unmarshal(m, b)
}
func (m *CMsgAppRights) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgAppRights.Marshal(b, m, deterministic)
}
func (dst *CMsgAppRights) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgAppRights.Merge(dst, src)
}
func (m *CMsgAppRights) XXX_Size() int {
	return xxx_messageInfo_CMsgAppRights.Size(m)
}
func (m *CMsgAppRights) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgAppRights.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgAppRights proto.InternalMessageInfo

func (m *CMsgAppRights) GetEditInfo() bool {
	if m != nil && m.EditInfo != nil {
		return *m.EditInfo
	}
	return false
}

func (m *CMsgAppRights) GetPublish() bool {
	if m != nil && m.Publish != nil {
		return *m.Publish
	}
	return false
}

func (m *CMsgAppRights) GetViewErrorData() bool {
	if m != nil && m.ViewErrorData != nil {
		return *m.ViewErrorData
	}
	return false
}

func (m *CMsgAppRights) GetDownload() bool {
	if m != nil && m.Download != nil {
		return *m.Download
	}
	return false
}

func (m *CMsgAppRights) GetUploadCdkeys() bool {
	if m != nil && m.UploadCdkeys != nil {
		return *m.UploadCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetGenerateCdkeys() bool {
	if m != nil && m.GenerateCdkeys != nil {
		return *m.GenerateCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetViewFinancials() bool {
	if m != nil && m.ViewFinancials != nil {
		return *m.ViewFinancials
	}
	return false
}

func (m *CMsgAppRights) GetManageCeg() bool {
	if m != nil && m.ManageCeg != nil {
		return *m.ManageCeg
	}
	return false
}

func (m *CMsgAppRights) GetManageSigning() bool {
	if m != nil && m.ManageSigning != nil {
		return *m.ManageSigning
	}
	return false
}

func (m *CMsgAppRights) GetManageCdkeys() bool {
	if m != nil && m.ManageCdkeys != nil {
		return *m.ManageCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetEditMarketing() bool {
	if m != nil && m.EditMarketing != nil {
		return *m.EditMarketing
	}
	return false
}

func (m *CMsgAppRights) GetEconomySupport() bool {
	if m != nil && m.EconomySupport != nil {
		return *m.EconomySupport
	}
	return false
}

func (m *CMsgAppRights) GetEconomySupportSupervisor() bool {
	if m != nil && m.EconomySupportSupervisor != nil {
		return *m.EconomySupportSupervisor
	}
	return false
}

func (m *CMsgAppRights) GetManagePricing() bool {
	if m != nil && m.ManagePricing != nil {
		return *m.ManagePricing
	}
	return false
}

func (m *CMsgAppRights) GetBroadcastLive() bool {
	if m != nil && m.BroadcastLive != nil {
		return *m.BroadcastLive
	}
	return false
}

type CCuratorPreferences struct {
	SupportedLanguages   *uint32  `protobuf:"varint,1,opt,name=supported_languages,json=supportedLanguages" json:"supported_languages,omitempty"`
	PlatformWindows      *bool    `protobuf:"varint,2,opt,name=platform_windows,json=platformWindows" json:"platform_windows,omitempty"`
	PlatformMac          *bool    `protobuf:"varint,3,opt,name=platform_mac,json=platformMac" json:"platform_mac,omitempty"`
	PlatformLinux        *bool    `protobuf:"varint,4,opt,name=platform_linux,json=platformLinux" json:"platform_linux,omitempty"`
	VrContent            *bool    `protobuf:"varint,5,opt,name=vr_content,json=vrContent" json:"vr_content,omitempty"`
	AdultContentViolence *bool    `protobuf:"varint,6,opt,name=adult_content_violence,json=adultContentViolence" json:"adult_content_violence,omitempty"`
	AdultContentSex      *bool    `protobuf:"varint,7,opt,name=adult_content_sex,json=adultContentSex" json:"adult_content_sex,omitempty"`
	TimestampUpdated     *uint32  `protobuf:"varint,8,opt,name=timestamp_updated,json=timestampUpdated" json:"timestamp_updated,omitempty"`
	TagidsCurated        []uint32 `protobuf:"varint,9,rep,name=tagids_curated,json=tagidsCurated" json:"tagids_curated,omitempty"`
	TagidsFiltered       []uint32 `protobuf:"varint,10,rep,name=tagids_filtered,json=tagidsFiltered" json:"tagids_filtered,omitempty"`
	WebsiteTitle         *string  `protobuf:"bytes,11,opt,name=website_title,json=websiteTitle" json:"website_title,omitempty"`
	WebsiteUrl           *string  `protobuf:"bytes,12,opt,name=website_url,json=websiteUrl" json:"website_url,omitempty"`
	DiscussionUrl        *string  `protobuf:"bytes,13,opt,name=discussion_url,json=discussionUrl" json:"discussion_url,omitempty"`
	ShowBroadcast        *bool    `protobuf:"varint,14,opt,name=show_broadcast,json=showBroadcast" json:"show_broadcast,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CCuratorPreferences) Reset()         { *m = CCuratorPreferences{} }
func (m *CCuratorPreferences) String() string { return proto.CompactTextString(m) }
func (*CCuratorPreferences) ProtoMessage()    {}
func (*CCuratorPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_steammessages_base_ecf731d44c7128ae, []int{6}
}
func (m *CCuratorPreferences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCuratorPreferences.Unmarshal(m, b)
}
func (m *CCuratorPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCuratorPreferences.Marshal(b, m, deterministic)
}
func (dst *CCuratorPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCuratorPreferences.Merge(dst, src)
}
func (m *CCuratorPreferences) XXX_Size() int {
	return xxx_messageInfo_CCuratorPreferences.Size(m)
}
func (m *CCuratorPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_CCuratorPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_CCuratorPreferences proto.InternalMessageInfo

func (m *CCuratorPreferences) GetSupportedLanguages() uint32 {
	if m != nil && m.SupportedLanguages != nil {
		return *m.SupportedLanguages
	}
	return 0
}

func (m *CCuratorPreferences) GetPlatformWindows() bool {
	if m != nil && m.PlatformWindows != nil {
		return *m.PlatformWindows
	}
	return false
}

func (m *CCuratorPreferences) GetPlatformMac() bool {
	if m != nil && m.PlatformMac != nil {
		return *m.PlatformMac
	}
	return false
}

func (m *CCuratorPreferences) GetPlatformLinux() bool {
	if m != nil && m.PlatformLinux != nil {
		return *m.PlatformLinux
	}
	return false
}

func (m *CCuratorPreferences) GetVrContent() bool {
	if m != nil && m.VrContent != nil {
		return *m.VrContent
	}
	return false
}

func (m *CCuratorPreferences) GetAdultContentViolence() bool {
	if m != nil && m.AdultContentViolence != nil {
		return *m.AdultContentViolence
	}
	return false
}

func (m *CCuratorPreferences) GetAdultContentSex() bool {
	if m != nil && m.AdultContentSex != nil {
		return *m.AdultContentSex
	}
	return false
}

func (m *CCuratorPreferences) GetTimestampUpdated() uint32 {
	if m != nil && m.TimestampUpdated != nil {
		return *m.TimestampUpdated
	}
	return 0
}

func (m *CCuratorPreferences) GetTagidsCurated() []uint32 {
	if m != nil {
		return m.TagidsCurated
	}
	return nil
}

func (m *CCuratorPreferences) GetTagidsFiltered() []uint32 {
	if m != nil {
		return m.TagidsFiltered
	}
	return nil
}

func (m *CCuratorPreferences) GetWebsiteTitle() string {
	if m != nil && m.WebsiteTitle != nil {
		return *m.WebsiteTitle
	}
	return ""
}

func (m *CCuratorPreferences) GetWebsiteUrl() string {
	if m != nil && m.WebsiteUrl != nil {
		return *m.WebsiteUrl
	}
	return ""
}

func (m *CCuratorPreferences) GetDiscussionUrl() string {
	if m != nil && m.DiscussionUrl != nil {
		return *m.DiscussionUrl
	}
	return ""
}

func (m *CCuratorPreferences) GetShowBroadcast() bool {
	if m != nil && m.ShowBroadcast != nil {
		return *m.ShowBroadcast
	}
	return false
}

var E_MsgpoolSoftLimit = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50000,
	Name:          "msgpool_soft_limit",
	Tag:           "varint,50000,opt,name=msgpool_soft_limit,json=msgpoolSoftLimit,def=32",
	Filename:      "steammessages_base.proto",
}

var E_MsgpoolHardLimit = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50001,
	Name:          "msgpool_hard_limit",
	Tag:           "varint,50001,opt,name=msgpool_hard_limit,json=msgpoolHardLimit,def=384",
	Filename:      "steammessages_base.proto",
}

var E_ForcePhpGeneration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50000,
	Name:          "force_php_generation",
	Tag:           "varint,50000,opt,name=force_php_generation,json=forcePhpGeneration,def=0",
	Filename:      "steammessages_base.proto",
}

var E_PhpOutputAlwaysNumber = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50020,
	Name:          "php_output_always_number",
	Tag:           "varint,50020,opt,name=php_output_always_number,json=phpOutputAlwaysNumber,def=0",
	Filename:      "steammessages_base.proto",
}

func init() {
	proto.RegisterType((*CMsgProtoBufHeader)(nil), "CMsgProtoBufHeader")
	proto.RegisterType((*CMsgMulti)(nil), "CMsgMulti")
	proto.RegisterType((*CMsgProtobufWrapped)(nil), "CMsgProtobufWrapped")
	proto.RegisterType((*CMsgAuthTicket)(nil), "CMsgAuthTicket")
	proto.RegisterType((*CCDDBAppDetailCommon)(nil), "CCDDBAppDetailCommon")
	proto.RegisterType((*CMsgAppRights)(nil), "CMsgAppRights")
	proto.RegisterType((*CCuratorPreferences)(nil), "CCuratorPreferences")
	proto.RegisterExtension(E_MsgpoolSoftLimit)
	proto.RegisterExtension(E_MsgpoolHardLimit)
	proto.RegisterExtension(E_ForcePhpGeneration)
	proto.RegisterExtension(E_PhpOutputAlwaysNumber)
}

func init() {
	proto.RegisterFile("steammessages_base.proto", fileDescriptor_steammessages_base_ecf731d44c7128ae)
}

var fileDescriptor_steammessages_base_ecf731d44c7128ae = []byte{
	// 1646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0x51, 0x73, 0x1b, 0xb7,
	0x11, 0x80, 0x4b, 0x59, 0x96, 0x49, 0x88, 0x47, 0xc9, 0x90, 0x6c, 0x5f, 0xed, 0x66, 0xc2, 0xaa,
	0x93, 0x56, 0x4d, 0x1b, 0xc9, 0xb2, 0x65, 0xcb, 0xd5, 0xf4, 0x45, 0xa2, 0xa3, 0xd8, 0xad, 0xe5,
	0x68, 0x8e, 0x72, 0x3c, 0xd3, 0x17, 0x0c, 0x78, 0xb7, 0x77, 0x44, 0x74, 0x77, 0x40, 0x00, 0x9c,
	0x68, 0xe6, 0xa9, 0x4f, 0xfd, 0x6d, 0x9d, 0xce, 0x74, 0x26, 0xfd, 0x0f, 0xfd, 0x01, 0xfd, 0x09,
	0x9d, 0x05, 0x70, 0x94, 0x94, 0xa4, 0xd3, 0x3c, 0xe9, 0xf0, 0xed, 0x62, 0xb9, 0xbb, 0xd8, 0x5d,
	0x40, 0x24, 0x36, 0x16, 0x78, 0x55, 0x81, 0x31, 0xbc, 0x00, 0xc3, 0x26, 0xdc, 0xc0, 0x8e, 0xd2,
	0xd2, 0xca, 0x87, 0xc3, 0x42, 0xca, 0xa2, 0x84, 0x5d, 0xb7, 0x9a, 0x34, 0xf9, 0x6e, 0x06, 0x26,
	0xd5, 0x42, 0x59, 0xa9, 0xbd, 0xc6, 0xd6, 0x3f, 0x56, 0x08, 0x1d, 0x9d, 0x9a, 0xe2, 0x0c, 0x57,
	0xc7, 0x4d, 0xfe, 0x0a, 0x78, 0x06, 0x9a, 0xc6, 0xe4, 0x8e, 0x33, 0x2a, 0xb2, 0xb8, 0x33, 0xec,
	0x6c, 0xaf, 0x24, 0xed, 0x92, 0xfe, 0x96, 0xac, 0xa7, 0xa5, 0x80, 0xda, 0x32, 0x03, 0xc6, 0x08,
	0x59, 0x8b, 0x2c, 0x5e, 0x1a, 0x76, 0xb6, 0x6f, 0x27, 0x6b, 0x9e, 0x8f, 0x5b, 0x4c, 0x7f, 0x45,
	0x22, 0x2d, 0x1b, 0x2b, 0xea, 0x82, 0x71, 0xa5, 0x44, 0x16, 0xdf, 0x1a, 0x76, 0xb6, 0xa3, 0xa4,
	0x1f, 0xe0, 0x11, 0x32, 0x7a, 0x40, 0xfa, 0x5f, 0xcb, 0x89, 0xc8, 0x98, 0x91, 0x8d, 0x4e, 0x21,
	0x26, 0xf8, 0x73, 0x87, 0x9b, 0x7b, 0x2f, 0xf6, 0xf7, 0x9f, 0x1f, 0xec, 0xef, 0x3f, 0x3e, 0x78,
	0x7a, 0xf0, 0xf8, 0x0f, 0xcf, 0x9e, 0xed, 0x3d, 0xdf, 0x7b, 0x96, 0xac, 0x3a, 0xcd, 0xb1, 0x53,
	0xbc, 0xda, 0x68, 0xb9, 0x2e, 0xc0, 0xc6, 0xab, 0xff, 0x77, 0xe3, 0xb9, 0x53, 0xa4, 0xbf, 0x26,
	0x6b, 0x7e, 0x0b, 0xfb, 0x5a, 0x4e, 0x58, 0xcd, 0x2b, 0x88, 0xfb, 0xc3, 0xce, 0x76, 0x2f, 0x89,
	0x3c, 0xfe, 0x93, 0x9c, 0xbc, 0xe5, 0x15, 0xd0, 0x07, 0xe4, 0x8e, 0x81, 0x6f, 0x58, 0xdd, 0x54,
	0x71, 0xec, 0x02, 0x5c, 0x31, 0xf0, 0xcd, 0xdb, 0xa6, 0xa2, 0x8f, 0xc8, 0x1d, 0xd0, 0x60, 0x9a,
	0xd2, 0xc6, 0x11, 0x0a, 0x0e, 0x3b, 0x4f, 0x92, 0x96, 0x60, 0xd0, 0xa0, 0xb5, 0xd4, 0x2c, 0x9c,
	0x47, 0x3c, 0x70, 0xb6, 0xfb, 0x0e, 0x9e, 0x7a, 0x46, 0x07, 0x64, 0x49, 0xa8, 0x78, 0xcd, 0xa5,
	0x63, 0x49, 0x28, 0xfa, 0x7b, 0x42, 0x79, 0x63, 0xa7, 0x8c, 0xa7, 0xa9, 0x6c, 0x6a, 0xcb, 0xf2,
	0x92, 0x17, 0x26, 0x5e, 0x77, 0xf2, 0x75, 0x94, 0x1c, 0x79, 0xc1, 0x09, 0x72, 0xfa, 0x4b, 0xd2,
	0xb7, 0xf2, 0x02, 0xea, 0x36, 0x65, 0xf7, 0x9d, 0xde, 0xaa, 0x63, 0x21, 0x39, 0x3b, 0x64, 0x83,
	0x67, 0x95, 0xa8, 0x99, 0x51, 0x52, 0xe6, 0x78, 0x02, 0x8d, 0x01, 0x1d, 0x3f, 0x18, 0x76, 0xb6,
	0xbb, 0xc9, 0x5d, 0x27, 0x1a, 0x07, 0xc9, 0x3b, 0x03, 0x9a, 0x7e, 0x4a, 0xd6, 0xac, 0xe6, 0xb5,
	0x51, 0x52, 0x5b, 0xe6, 0x5c, 0x8d, 0xef, 0xfa, 0xd0, 0xf6, 0x92, 0xc1, 0x42, 0xf2, 0x39, 0x0a,
	0xe8, 0x13, 0xd2, 0x0b, 0xb1, 0x89, 0x2c, 0xa6, 0xc3, 0xce, 0xf6, 0xf2, 0xff, 0xc8, 0xfa, 0x95,
	0x1a, 0x06, 0xa8, 0x9a, 0x49, 0x29, 0xcc, 0x14, 0x34, 0x2b, 0xb4, 0x6c, 0x14, 0x13, 0x59, 0xbc,
	0xe1, 0x03, 0x5c, 0x48, 0xbe, 0x40, 0xc1, 0xeb, 0x8c, 0x6e, 0x92, 0xdb, 0x66, 0x6e, 0x44, 0x16,
	0x6f, 0x3a, 0x05, 0xbf, 0xa0, 0x8f, 0x48, 0xcf, 0x6a, 0x9e, 0x02, 0xb3, 0xbc, 0x88, 0xef, 0xe1,
	0xef, 0x26, 0x5d, 0x07, 0xce, 0x79, 0x41, 0xb7, 0x48, 0x34, 0x83, 0x09, 0x57, 0x82, 0x5d, 0xc0,
	0x1c, 0x6d, 0xff, 0xdc, 0x27, 0xc5, 0xc3, 0x3f, 0xc3, 0xfc, 0x75, 0x46, 0x9f, 0x91, 0x07, 0xc2,
	0xb0, 0x5c, 0xcb, 0x8a, 0xc1, 0x07, 0x0b, 0xba, 0xe6, 0x65, 0x9b, 0xc2, 0x87, 0x2e, 0x31, 0x9b,
	0xc2, 0x9c, 0x68, 0x59, 0x7d, 0x1e, 0x84, 0x21, 0x97, 0xdb, 0x64, 0x3d, 0x97, 0x7a, 0xc6, 0x75,
	0xc6, 0xac, 0x64, 0xde, 0xb1, 0x47, 0xc3, 0x5b, 0xdb, 0x51, 0x32, 0x08, 0xfc, 0x5c, 0x8e, 0x91,
	0x6e, 0x8d, 0x49, 0x0f, 0x7b, 0xe9, 0xb4, 0x29, 0xad, 0xc0, 0x42, 0x30, 0xe2, 0x5b, 0x60, 0x4d,
	0xfd, 0xad, 0x50, 0x0a, 0x7c, 0x23, 0x45, 0x49, 0x1f, 0xe1, 0xbb, 0xc0, 0xf0, 0x28, 0x43, 0x92,
	0xd8, 0x44, 0x66, 0x73, 0xd7, 0x49, 0xfd, 0x64, 0x35, 0xb0, 0x63, 0x99, 0xcd, 0xb7, 0x5e, 0x90,
	0x8d, 0x45, 0x83, 0x4e, 0x9a, 0xfc, 0xbd, 0xe6, 0x3f, 0xba, 0xb3, 0xf3, 0xc3, 0x9d, 0xdf, 0x75,
	0xc8, 0x00, 0xb7, 0x1e, 0x35, 0x76, 0x7a, 0x2e, 0xd2, 0x0b, 0xb0, 0xf4, 0x3e, 0x59, 0x01, 0x63,
	0xb9, 0x85, 0xe0, 0x4d, 0x58, 0x5d, 0x2f, 0x69, 0x74, 0x21, 0xba, 0x51, 0xd2, 0xd7, 0x86, 0xc1,
	0xad, 0x9b, 0xc3, 0xe0, 0x3e, 0x59, 0x29, 0x78, 0x85, 0x75, 0xb0, 0xec, 0x04, 0x61, 0x45, 0x87,
	0xa4, 0x3f, 0x65, 0x4e, 0x89, 0x29, 0xa1, 0x20, 0xbe, 0xed, 0x7e, 0x8c, 0x4c, 0xc7, 0x88, 0xce,
	0x84, 0x02, 0xfa, 0x11, 0x21, 0xd6, 0xb9, 0xc4, 0x52, 0x9d, 0xc6, 0x2b, 0x4e, 0xde, 0xf3, 0x64,
	0xa4, 0x53, 0x34, 0xec, 0x17, 0xf1, 0x1d, 0x17, 0x57, 0x58, 0x6d, 0xfd, 0x67, 0x89, 0x6c, 0x8e,
	0x46, 0x2f, 0x5f, 0x1e, 0x1f, 0x29, 0xf5, 0x12, 0x2c, 0x17, 0xe5, 0x48, 0x56, 0x95, 0xac, 0xb1,
	0x64, 0xfc, 0x8c, 0xf1, 0x71, 0xf9, 0x05, 0xa5, 0x64, 0xd9, 0xf5, 0xf7, 0x92, 0xeb, 0x41, 0xf7,
	0x8d, 0x4c, 0xa4, 0xb2, 0x76, 0xa1, 0xf4, 0x12, 0xf7, 0x8d, 0xac, 0x94, 0x85, 0x74, 0x51, 0xf4,
	0x12, 0xf7, 0x8d, 0x1e, 0xe2, 0x5f, 0x66, 0x2a, 0x5e, 0x96, 0x2e, 0x82, 0x5e, 0xd2, 0x43, 0x32,
	0x46, 0x80, 0x5b, 0xac, 0x94, 0xa5, 0x73, 0xbd, 0x9b, 0xb8, 0x6f, 0x64, 0x19, 0x54, 0xd2, 0xf9,
	0xdc, 0x4d, 0xdc, 0x37, 0x3a, 0x56, 0x41, 0x26, 0x78, 0xdc, 0x75, 0xd0, 0x2f, 0xe8, 0x73, 0xf2,
	0x20, 0x95, 0x55, 0xd5, 0xd4, 0xc2, 0xce, 0xd9, 0xa5, 0x30, 0x62, 0x52, 0x02, 0xc3, 0xa3, 0x30,
	0x71, 0xcf, 0xe9, 0xdd, 0x5b, 0x88, 0xbf, 0xf2, 0xd2, 0x31, 0x0a, 0xb1, 0xa8, 0x72, 0x2d, 0xa0,
	0xce, 0xca, 0xb9, 0x9f, 0x5c, 0xc4, 0x4f, 0x97, 0x16, 0xba, 0xc1, 0x35, 0x24, 0xab, 0x4a, 0x4b,
	0xc5, 0x0b, 0x6e, 0x85, 0xac, 0xdd, 0x60, 0xec, 0x25, 0xd7, 0x11, 0xfd, 0x94, 0xdc, 0x9d, 0x72,
	0xc3, 0x78, 0xd6, 0x94, 0x96, 0xa5, 0xb2, 0xb6, 0x50, 0x5b, 0x37, 0x04, 0xbb, 0xc9, 0xda, 0x94,
	0x9b, 0x23, 0xe4, 0x23, 0x8f, 0xb7, 0xfe, 0xbe, 0x4c, 0x22, 0x57, 0x45, 0x4a, 0x25, 0xa2, 0x98,
	0x5a, 0x83, 0x8d, 0x08, 0x99, 0xb0, 0x4c, 0xd4, 0xb9, 0x74, 0xf9, 0xee, 0x26, 0x5d, 0x04, 0xaf,
	0xeb, 0x5c, 0x62, 0xb1, 0x84, 0x7e, 0x76, 0x59, 0xef, 0x26, 0xed, 0x12, 0xe7, 0xee, 0xa5, 0x80,
	0x99, 0x1f, 0x2f, 0x2c, 0xe3, 0x96, 0xbb, 0x33, 0xe8, 0x26, 0x11, 0x62, 0x37, 0x5b, 0x5e, 0x72,
	0xcb, 0xe9, 0x43, 0xd2, 0xcd, 0xe4, 0xac, 0x2e, 0x25, 0xf7, 0x65, 0xd5, 0x4d, 0x16, 0x6b, 0x8c,
	0xbf, 0x51, 0xf8, 0xc5, 0xd2, 0xec, 0x02, 0xe6, 0xc6, 0x9d, 0x4b, 0x37, 0xe9, 0x7b, 0x38, 0x72,
	0x8c, 0xfe, 0x86, 0xac, 0x15, 0x50, 0x83, 0xe6, 0x16, 0x5a, 0x35, 0x7f, 0x4a, 0x83, 0x16, 0x5f,
	0x29, 0x3a, 0x8f, 0x72, 0x51, 0xf3, 0x3a, 0x15, 0xbc, 0x34, 0xe1, 0xe8, 0x06, 0x88, 0x4f, 0x16,
	0x14, 0x6b, 0xa1, 0xe2, 0x35, 0xf6, 0x5a, 0x0a, 0x45, 0x38, 0xc9, 0x9e, 0x27, 0x23, 0x28, 0xe8,
	0x27, 0x64, 0x10, 0xc4, 0x46, 0x14, 0xb5, 0xa8, 0x8b, 0x70, 0x88, 0x91, 0xa7, 0x63, 0x0f, 0xd1,
	0xf9, 0xd6, 0x8a, 0xf7, 0x8a, 0x78, 0xe7, 0x83, 0x21, 0xef, 0xd3, 0x27, 0x64, 0xe0, 0x92, 0x5b,
	0x71, 0x7d, 0x01, 0x78, 0x4d, 0xba, 0xf3, 0xeb, 0x26, 0x11, 0xd2, 0xd3, 0x16, 0xa2, 0xeb, 0x90,
	0xca, 0x5a, 0x56, 0x73, 0x66, 0x1a, 0x85, 0xc3, 0x39, 0x9c, 0xdf, 0x20, 0xe0, 0xb1, 0xa7, 0xf4,
	0x8f, 0xe4, 0xe1, 0xf7, 0x14, 0xf1, 0x2f, 0xe8, 0x4b, 0x61, 0xa4, 0x76, 0xf7, 0x57, 0x37, 0x89,
	0x6f, 0xee, 0x19, 0x2f, 0xe4, 0xd7, 0x22, 0x53, 0x5a, 0xa4, 0xe8, 0xcd, 0xe0, 0x7a, 0x64, 0x67,
	0x1e, 0xa2, 0xda, 0x44, 0x4b, 0x9e, 0xa5, 0xdc, 0x58, 0x56, 0x8a, 0x4b, 0x70, 0x77, 0x5b, 0x37,
	0x89, 0x16, 0xf4, 0x8d, 0xb8, 0x84, 0xad, 0x7f, 0x2e, 0x93, 0x8d, 0xd1, 0xa8, 0xd1, 0xdc, 0x4a,
	0x7d, 0xa6, 0x21, 0x07, 0x0d, 0x75, 0x0a, 0x86, 0xee, 0x92, 0x8d, 0xe0, 0x1b, 0x64, 0xac, 0xe4,
	0x75, 0xd1, 0xe0, 0x43, 0x26, 0xb4, 0x32, 0x5d, 0x88, 0xde, 0xb4, 0x12, 0x7c, 0x84, 0xa8, 0x92,
	0xdb, 0x5c, 0xea, 0x8a, 0xcd, 0x44, 0x9d, 0xc9, 0x99, 0x09, 0xd5, 0xb6, 0xd6, 0xf2, 0xf7, 0x1e,
	0xe3, 0x9c, 0x5c, 0xa8, 0x56, 0x3c, 0x0d, 0x25, 0xb7, 0xda, 0xb2, 0x53, 0x9e, 0xa2, 0xf7, 0x0b,
	0x95, 0x52, 0xd4, 0xcd, 0x87, 0x50, 0x76, 0x51, 0x4b, 0xdf, 0x20, 0xc4, 0x22, 0xb8, 0xd4, 0x8b,
	0x6e, 0xf1, 0x85, 0xd7, 0xbb, 0xd4, 0xa1, 0x4f, 0xe8, 0x3e, 0xb9, 0x7f, 0xa3, 0x9f, 0xd8, 0xa5,
	0x90, 0x25, 0xc6, 0x17, 0x8a, 0x6f, 0x93, 0x5f, 0xeb, 0xaa, 0xaf, 0x82, 0x0c, 0x3b, 0xf1, 0xe6,
	0x2e, 0x03, 0x1f, 0x42, 0x11, 0xae, 0x5d, 0xdf, 0x30, 0x86, 0x0f, 0xf4, 0x77, 0xe4, 0xae, 0x15,
	0x15, 0x4e, 0xec, 0x4a, 0xb1, 0x46, 0x65, 0xdc, 0x42, 0xe6, 0x8a, 0x31, 0x4a, 0xd6, 0x17, 0x82,
	0x77, 0x9e, 0x63, 0x50, 0x96, 0x17, 0x22, 0x33, 0x2c, 0xc5, 0x84, 0x43, 0x16, 0xf7, 0xdc, 0x9d,
	0x15, 0x79, 0x3a, 0xf2, 0x10, 0xeb, 0x28, 0xa8, 0xe5, 0xa2, 0xb4, 0xa0, 0x21, 0x8b, 0x89, 0xbf,
	0xdb, 0x3c, 0x3e, 0x09, 0x14, 0x8b, 0x77, 0x06, 0x13, 0x23, 0x2c, 0x30, 0x2b, 0x6c, 0x09, 0x61,
	0xac, 0xf4, 0x03, 0x3c, 0x47, 0x46, 0x3f, 0x26, 0xab, 0xad, 0x52, 0xa3, 0xcb, 0xf0, 0xac, 0x22,
	0x01, 0xbd, 0xd3, 0x25, 0x7a, 0x95, 0x09, 0x93, 0x36, 0xee, 0x89, 0xe8, 0x74, 0x22, 0xff, 0xf4,
	0xba, 0xa2, 0x41, 0xcd, 0x4c, 0xe5, 0x8c, 0x2d, 0xca, 0xa7, 0x2d, 0x3b, 0xa4, 0xc7, 0x2d, 0x3c,
	0x1c, 0x13, 0x5a, 0x99, 0x42, 0x49, 0x89, 0xf7, 0x78, 0x8e, 0x95, 0x57, 0x09, 0x4b, 0x3f, 0xde,
	0xf1, 0xaf, 0xde, 0x9d, 0xf6, 0xd5, 0xbb, 0x13, 0x9e, 0x5d, 0x5f, 0x2a, 0x1c, 0x7e, 0x26, 0xfe,
	0xee, 0x6f, 0xb7, 0xdc, 0xdb, 0x66, 0xe9, 0xe9, 0x93, 0x64, 0x3d, 0x18, 0x18, 0xcb, 0xdc, 0xbe,
	0xc1, 0xed, 0x87, 0xe7, 0x57, 0x46, 0xa7, 0x78, 0xe7, 0xff, 0x44, 0xa3, 0xff, 0x0a, 0x46, 0x6f,
	0x3d, 0x7d, 0xb1, 0xbf, 0xb0, 0xfa, 0x8a, 0xeb, 0xcc, 0x5b, 0x7d, 0x4f, 0x36, 0x73, 0xa9, 0x53,
	0x60, 0x6a, 0xaa, 0x58, 0x18, 0x43, 0x38, 0x89, 0x7f, 0xf1, 0x03, 0xbb, 0x27, 0xa2, 0xfc, 0x9e,
	0xa7, 0xdd, 0xc3, 0xdb, 0x39, 0x2f, 0x0d, 0x24, 0xd4, 0x99, 0x38, 0x9b, 0xaa, 0x2f, 0x16, 0x06,
	0x0e, 0x19, 0x89, 0xd1, 0xa4, 0x6c, 0xac, 0x6a, 0x2c, 0xe3, 0xe5, 0x8c, 0xcf, 0x0d, 0xbe, 0x59,
	0x27, 0xa0, 0xe9, 0x47, 0x3f, 0x62, 0x1c, 0xca, 0xac, 0xb5, 0xfe, 0xef, 0x9b, 0xd6, 0xef, 0xa9,
	0xa9, 0xfa, 0xd2, 0x99, 0x39, 0x72, 0x56, 0xde, 0x3a, 0x23, 0xc7, 0x07, 0xaf, 0x3a, 0x7f, 0x79,
	0x5c, 0x08, 0x3b, 0x6d, 0x26, 0x3b, 0xa9, 0xac, 0x76, 0xf7, 0x9e, 0x5e, 0xec, 0x16, 0xf2, 0x33,
	0x77, 0xbb, 0x7f, 0xa6, 0xc1, 0x3f, 0xa0, 0xcc, 0xd5, 0xbf, 0x18, 0x4e, 0xf0, 0xd7, 0xce, 0xcf,
	0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x66, 0xfc, 0x12, 0x20, 0x97, 0x0c, 0x00, 0x00,
}
