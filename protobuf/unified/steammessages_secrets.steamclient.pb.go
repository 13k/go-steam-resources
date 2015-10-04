// Code generated by protoc-gen-go.
// source: steammessages_secrets.steamclient.proto
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

type EKeyEscrowUsage int32

const (
	EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice EKeyEscrowUsage = 0
)

var EKeyEscrowUsage_name = map[int32]string{
	0: "k_EKeyEscrowUsageStreamingDevice",
}
var EKeyEscrowUsage_value = map[string]int32{
	"k_EKeyEscrowUsageStreamingDevice": 0,
}

func (x EKeyEscrowUsage) Enum() *EKeyEscrowUsage {
	p := new(EKeyEscrowUsage)
	*p = x
	return p
}
func (x EKeyEscrowUsage) String() string {
	return proto.EnumName(EKeyEscrowUsage_name, int32(x))
}
func (x *EKeyEscrowUsage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EKeyEscrowUsage_value, data, "EKeyEscrowUsage")
	if err != nil {
		return err
	}
	*x = EKeyEscrowUsage(value)
	return nil
}

type CKeyEscrow_Request struct {
	RsaOaepShaTicket []byte           `protobuf:"bytes,1,opt,name=rsa_oaep_sha_ticket" json:"rsa_oaep_sha_ticket,omitempty"`
	Password         []byte           `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Usage            *EKeyEscrowUsage `protobuf:"varint,3,opt,name=usage,enum=EKeyEscrowUsage,def=0" json:"usage,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CKeyEscrow_Request) Reset()         { *m = CKeyEscrow_Request{} }
func (m *CKeyEscrow_Request) String() string { return proto.CompactTextString(m) }
func (*CKeyEscrow_Request) ProtoMessage()    {}

const Default_CKeyEscrow_Request_Usage EKeyEscrowUsage = EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice

func (m *CKeyEscrow_Request) GetRsaOaepShaTicket() []byte {
	if m != nil {
		return m.RsaOaepShaTicket
	}
	return nil
}

func (m *CKeyEscrow_Request) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CKeyEscrow_Request) GetUsage() EKeyEscrowUsage {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return Default_CKeyEscrow_Request_Usage
}

type CKeyEscrow_Ticket struct {
	Password         []byte           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Identifier       *uint64          `protobuf:"varint,2,opt,name=identifier" json:"identifier,omitempty"`
	Payload          []byte           `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	Timestamp        *uint32          `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Usage            *EKeyEscrowUsage `protobuf:"varint,5,opt,name=usage,enum=EKeyEscrowUsage,def=0" json:"usage,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CKeyEscrow_Ticket) Reset()         { *m = CKeyEscrow_Ticket{} }
func (m *CKeyEscrow_Ticket) String() string { return proto.CompactTextString(m) }
func (*CKeyEscrow_Ticket) ProtoMessage()    {}

const Default_CKeyEscrow_Ticket_Usage EKeyEscrowUsage = EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice

func (m *CKeyEscrow_Ticket) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CKeyEscrow_Ticket) GetIdentifier() uint64 {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return 0
}

func (m *CKeyEscrow_Ticket) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CKeyEscrow_Ticket) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CKeyEscrow_Ticket) GetUsage() EKeyEscrowUsage {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return Default_CKeyEscrow_Ticket_Usage
}

type CKeyEscrow_Response struct {
	Ticket           *CKeyEscrow_Ticket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CKeyEscrow_Response) Reset()         { *m = CKeyEscrow_Response{} }
func (m *CKeyEscrow_Response) String() string { return proto.CompactTextString(m) }
func (*CKeyEscrow_Response) ProtoMessage()    {}

func (m *CKeyEscrow_Response) GetTicket() *CKeyEscrow_Ticket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func init() {
	proto.RegisterEnum("EKeyEscrowUsage", EKeyEscrowUsage_name, EKeyEscrowUsage_value)
}