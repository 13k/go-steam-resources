// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_client_enums.proto

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

type ETournamentTemplate int32

const (
	ETournamentTemplate_k_ETournamentTemplate_None          ETournamentTemplate = 0
	ETournamentTemplate_k_ETournamentTemplate_AutomatedWin3 ETournamentTemplate = 1
)

var ETournamentTemplate_name = map[int32]string{
	0: "k_ETournamentTemplate_None",
	1: "k_ETournamentTemplate_AutomatedWin3",
}
var ETournamentTemplate_value = map[string]int32{
	"k_ETournamentTemplate_None":          0,
	"k_ETournamentTemplate_AutomatedWin3": 1,
}

func (x ETournamentTemplate) Enum() *ETournamentTemplate {
	p := new(ETournamentTemplate)
	*p = x
	return p
}
func (x ETournamentTemplate) String() string {
	return proto.EnumName(ETournamentTemplate_name, int32(x))
}
func (x *ETournamentTemplate) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentTemplate_value, data, "ETournamentTemplate")
	if err != nil {
		return err
	}
	*x = ETournamentTemplate(value)
	return nil
}
func (ETournamentTemplate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{0}
}

type ETournamentGameState int32

const (
	ETournamentGameState_k_ETournamentGameState_Unknown              ETournamentGameState = 0
	ETournamentGameState_k_ETournamentGameState_Canceled             ETournamentGameState = 1
	ETournamentGameState_k_ETournamentGameState_Scheduled            ETournamentGameState = 2
	ETournamentGameState_k_ETournamentGameState_Active               ETournamentGameState = 3
	ETournamentGameState_k_ETournamentGameState_RadVictory           ETournamentGameState = 20
	ETournamentGameState_k_ETournamentGameState_DireVictory          ETournamentGameState = 21
	ETournamentGameState_k_ETournamentGameState_RadVictoryByForfeit  ETournamentGameState = 22
	ETournamentGameState_k_ETournamentGameState_DireVictoryByForfeit ETournamentGameState = 23
	ETournamentGameState_k_ETournamentGameState_ServerFailure        ETournamentGameState = 40
	ETournamentGameState_k_ETournamentGameState_NotNeeded            ETournamentGameState = 41
)

var ETournamentGameState_name = map[int32]string{
	0:  "k_ETournamentGameState_Unknown",
	1:  "k_ETournamentGameState_Canceled",
	2:  "k_ETournamentGameState_Scheduled",
	3:  "k_ETournamentGameState_Active",
	20: "k_ETournamentGameState_RadVictory",
	21: "k_ETournamentGameState_DireVictory",
	22: "k_ETournamentGameState_RadVictoryByForfeit",
	23: "k_ETournamentGameState_DireVictoryByForfeit",
	40: "k_ETournamentGameState_ServerFailure",
	41: "k_ETournamentGameState_NotNeeded",
}
var ETournamentGameState_value = map[string]int32{
	"k_ETournamentGameState_Unknown":              0,
	"k_ETournamentGameState_Canceled":             1,
	"k_ETournamentGameState_Scheduled":            2,
	"k_ETournamentGameState_Active":               3,
	"k_ETournamentGameState_RadVictory":           20,
	"k_ETournamentGameState_DireVictory":          21,
	"k_ETournamentGameState_RadVictoryByForfeit":  22,
	"k_ETournamentGameState_DireVictoryByForfeit": 23,
	"k_ETournamentGameState_ServerFailure":        40,
	"k_ETournamentGameState_NotNeeded":            41,
}

func (x ETournamentGameState) Enum() *ETournamentGameState {
	p := new(ETournamentGameState)
	*p = x
	return p
}
func (x ETournamentGameState) String() string {
	return proto.EnumName(ETournamentGameState_name, int32(x))
}
func (x *ETournamentGameState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentGameState_value, data, "ETournamentGameState")
	if err != nil {
		return err
	}
	*x = ETournamentGameState(value)
	return nil
}
func (ETournamentGameState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{1}
}

type ETournamentTeamState int32

const (
	ETournamentTeamState_k_ETournamentTeamState_Unknown      ETournamentTeamState = 0
	ETournamentTeamState_k_ETournamentTeamState_Node1        ETournamentTeamState = 1
	ETournamentTeamState_k_ETournamentTeamState_NodeMax      ETournamentTeamState = 1024
	ETournamentTeamState_k_ETournamentTeamState_Eliminated   ETournamentTeamState = 14003
	ETournamentTeamState_k_ETournamentTeamState_Forfeited    ETournamentTeamState = 14004
	ETournamentTeamState_k_ETournamentTeamState_Finished1st  ETournamentTeamState = 15001
	ETournamentTeamState_k_ETournamentTeamState_Finished2nd  ETournamentTeamState = 15002
	ETournamentTeamState_k_ETournamentTeamState_Finished3rd  ETournamentTeamState = 15003
	ETournamentTeamState_k_ETournamentTeamState_Finished4th  ETournamentTeamState = 15004
	ETournamentTeamState_k_ETournamentTeamState_Finished5th  ETournamentTeamState = 15005
	ETournamentTeamState_k_ETournamentTeamState_Finished6th  ETournamentTeamState = 15006
	ETournamentTeamState_k_ETournamentTeamState_Finished7th  ETournamentTeamState = 15007
	ETournamentTeamState_k_ETournamentTeamState_Finished8th  ETournamentTeamState = 15008
	ETournamentTeamState_k_ETournamentTeamState_Finished9th  ETournamentTeamState = 15009
	ETournamentTeamState_k_ETournamentTeamState_Finished10th ETournamentTeamState = 15010
	ETournamentTeamState_k_ETournamentTeamState_Finished11th ETournamentTeamState = 15011
	ETournamentTeamState_k_ETournamentTeamState_Finished12th ETournamentTeamState = 15012
	ETournamentTeamState_k_ETournamentTeamState_Finished13th ETournamentTeamState = 15013
	ETournamentTeamState_k_ETournamentTeamState_Finished14th ETournamentTeamState = 15014
	ETournamentTeamState_k_ETournamentTeamState_Finished15th ETournamentTeamState = 15015
	ETournamentTeamState_k_ETournamentTeamState_Finished16th ETournamentTeamState = 15016
)

var ETournamentTeamState_name = map[int32]string{
	0:     "k_ETournamentTeamState_Unknown",
	1:     "k_ETournamentTeamState_Node1",
	1024:  "k_ETournamentTeamState_NodeMax",
	14003: "k_ETournamentTeamState_Eliminated",
	14004: "k_ETournamentTeamState_Forfeited",
	15001: "k_ETournamentTeamState_Finished1st",
	15002: "k_ETournamentTeamState_Finished2nd",
	15003: "k_ETournamentTeamState_Finished3rd",
	15004: "k_ETournamentTeamState_Finished4th",
	15005: "k_ETournamentTeamState_Finished5th",
	15006: "k_ETournamentTeamState_Finished6th",
	15007: "k_ETournamentTeamState_Finished7th",
	15008: "k_ETournamentTeamState_Finished8th",
	15009: "k_ETournamentTeamState_Finished9th",
	15010: "k_ETournamentTeamState_Finished10th",
	15011: "k_ETournamentTeamState_Finished11th",
	15012: "k_ETournamentTeamState_Finished12th",
	15013: "k_ETournamentTeamState_Finished13th",
	15014: "k_ETournamentTeamState_Finished14th",
	15015: "k_ETournamentTeamState_Finished15th",
	15016: "k_ETournamentTeamState_Finished16th",
}
var ETournamentTeamState_value = map[string]int32{
	"k_ETournamentTeamState_Unknown":      0,
	"k_ETournamentTeamState_Node1":        1,
	"k_ETournamentTeamState_NodeMax":      1024,
	"k_ETournamentTeamState_Eliminated":   14003,
	"k_ETournamentTeamState_Forfeited":    14004,
	"k_ETournamentTeamState_Finished1st":  15001,
	"k_ETournamentTeamState_Finished2nd":  15002,
	"k_ETournamentTeamState_Finished3rd":  15003,
	"k_ETournamentTeamState_Finished4th":  15004,
	"k_ETournamentTeamState_Finished5th":  15005,
	"k_ETournamentTeamState_Finished6th":  15006,
	"k_ETournamentTeamState_Finished7th":  15007,
	"k_ETournamentTeamState_Finished8th":  15008,
	"k_ETournamentTeamState_Finished9th":  15009,
	"k_ETournamentTeamState_Finished10th": 15010,
	"k_ETournamentTeamState_Finished11th": 15011,
	"k_ETournamentTeamState_Finished12th": 15012,
	"k_ETournamentTeamState_Finished13th": 15013,
	"k_ETournamentTeamState_Finished14th": 15014,
	"k_ETournamentTeamState_Finished15th": 15015,
	"k_ETournamentTeamState_Finished16th": 15016,
}

func (x ETournamentTeamState) Enum() *ETournamentTeamState {
	p := new(ETournamentTeamState)
	*p = x
	return p
}
func (x ETournamentTeamState) String() string {
	return proto.EnumName(ETournamentTeamState_name, int32(x))
}
func (x *ETournamentTeamState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentTeamState_value, data, "ETournamentTeamState")
	if err != nil {
		return err
	}
	*x = ETournamentTeamState(value)
	return nil
}
func (ETournamentTeamState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{2}
}

type ETournamentState int32

const (
	ETournamentState_k_ETournamentState_Unknown                     ETournamentState = 0
	ETournamentState_k_ETournamentState_CanceledByAdmin             ETournamentState = 1
	ETournamentState_k_ETournamentState_Completed                   ETournamentState = 2
	ETournamentState_k_ETournamentState_Merged                      ETournamentState = 3
	ETournamentState_k_ETournamentState_ServerFailure               ETournamentState = 4
	ETournamentState_k_ETournamentState_TeamAbandoned               ETournamentState = 5
	ETournamentState_k_ETournamentState_TeamTimeoutForfeit          ETournamentState = 6
	ETournamentState_k_ETournamentState_TeamTimeoutRefund           ETournamentState = 7
	ETournamentState_k_ETournamentState_ServerFailureGrantedVictory ETournamentState = 8
	ETournamentState_k_ETournamentState_TeamTimeoutGrantedVictory   ETournamentState = 9
	ETournamentState_k_ETournamentState_InProgress                  ETournamentState = 100
	ETournamentState_k_ETournamentState_WaitingToMerge              ETournamentState = 101
)

var ETournamentState_name = map[int32]string{
	0:   "k_ETournamentState_Unknown",
	1:   "k_ETournamentState_CanceledByAdmin",
	2:   "k_ETournamentState_Completed",
	3:   "k_ETournamentState_Merged",
	4:   "k_ETournamentState_ServerFailure",
	5:   "k_ETournamentState_TeamAbandoned",
	6:   "k_ETournamentState_TeamTimeoutForfeit",
	7:   "k_ETournamentState_TeamTimeoutRefund",
	8:   "k_ETournamentState_ServerFailureGrantedVictory",
	9:   "k_ETournamentState_TeamTimeoutGrantedVictory",
	100: "k_ETournamentState_InProgress",
	101: "k_ETournamentState_WaitingToMerge",
}
var ETournamentState_value = map[string]int32{
	"k_ETournamentState_Unknown":                     0,
	"k_ETournamentState_CanceledByAdmin":             1,
	"k_ETournamentState_Completed":                   2,
	"k_ETournamentState_Merged":                      3,
	"k_ETournamentState_ServerFailure":               4,
	"k_ETournamentState_TeamAbandoned":               5,
	"k_ETournamentState_TeamTimeoutForfeit":          6,
	"k_ETournamentState_TeamTimeoutRefund":           7,
	"k_ETournamentState_ServerFailureGrantedVictory": 8,
	"k_ETournamentState_TeamTimeoutGrantedVictory":   9,
	"k_ETournamentState_InProgress":                  100,
	"k_ETournamentState_WaitingToMerge":              101,
}

func (x ETournamentState) Enum() *ETournamentState {
	p := new(ETournamentState)
	*p = x
	return p
}
func (x ETournamentState) String() string {
	return proto.EnumName(ETournamentState_name, int32(x))
}
func (x *ETournamentState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentState_value, data, "ETournamentState")
	if err != nil {
		return err
	}
	*x = ETournamentState(value)
	return nil
}
func (ETournamentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{3}
}

type ETournamentNodeState int32

const (
	ETournamentNodeState_k_ETournamentNodeState_Unknown             ETournamentNodeState = 0
	ETournamentNodeState_k_ETournamentNodeState_Canceled            ETournamentNodeState = 1
	ETournamentNodeState_k_ETournamentNodeState_TeamsNotYetAssigned ETournamentNodeState = 2
	ETournamentNodeState_k_ETournamentNodeState_InBetweenGames      ETournamentNodeState = 3
	ETournamentNodeState_k_ETournamentNodeState_GameInProgress      ETournamentNodeState = 4
	ETournamentNodeState_k_ETournamentNodeState_A_Won               ETournamentNodeState = 5
	ETournamentNodeState_k_ETournamentNodeState_B_Won               ETournamentNodeState = 6
	ETournamentNodeState_k_ETournamentNodeState_A_WonByForfeit      ETournamentNodeState = 7
	ETournamentNodeState_k_ETournamentNodeState_B_WonByForfeit      ETournamentNodeState = 8
	ETournamentNodeState_k_ETournamentNodeState_A_Bye               ETournamentNodeState = 9
	ETournamentNodeState_k_ETournamentNodeState_A_Abandoned         ETournamentNodeState = 10
	ETournamentNodeState_k_ETournamentNodeState_ServerFailure       ETournamentNodeState = 11
	ETournamentNodeState_k_ETournamentNodeState_A_TimeoutForfeit    ETournamentNodeState = 12
	ETournamentNodeState_k_ETournamentNodeState_A_TimeoutRefund     ETournamentNodeState = 13
)

var ETournamentNodeState_name = map[int32]string{
	0:  "k_ETournamentNodeState_Unknown",
	1:  "k_ETournamentNodeState_Canceled",
	2:  "k_ETournamentNodeState_TeamsNotYetAssigned",
	3:  "k_ETournamentNodeState_InBetweenGames",
	4:  "k_ETournamentNodeState_GameInProgress",
	5:  "k_ETournamentNodeState_A_Won",
	6:  "k_ETournamentNodeState_B_Won",
	7:  "k_ETournamentNodeState_A_WonByForfeit",
	8:  "k_ETournamentNodeState_B_WonByForfeit",
	9:  "k_ETournamentNodeState_A_Bye",
	10: "k_ETournamentNodeState_A_Abandoned",
	11: "k_ETournamentNodeState_ServerFailure",
	12: "k_ETournamentNodeState_A_TimeoutForfeit",
	13: "k_ETournamentNodeState_A_TimeoutRefund",
}
var ETournamentNodeState_value = map[string]int32{
	"k_ETournamentNodeState_Unknown":             0,
	"k_ETournamentNodeState_Canceled":            1,
	"k_ETournamentNodeState_TeamsNotYetAssigned": 2,
	"k_ETournamentNodeState_InBetweenGames":      3,
	"k_ETournamentNodeState_GameInProgress":      4,
	"k_ETournamentNodeState_A_Won":               5,
	"k_ETournamentNodeState_B_Won":               6,
	"k_ETournamentNodeState_A_WonByForfeit":      7,
	"k_ETournamentNodeState_B_WonByForfeit":      8,
	"k_ETournamentNodeState_A_Bye":               9,
	"k_ETournamentNodeState_A_Abandoned":         10,
	"k_ETournamentNodeState_ServerFailure":       11,
	"k_ETournamentNodeState_A_TimeoutForfeit":    12,
	"k_ETournamentNodeState_A_TimeoutRefund":     13,
}

func (x ETournamentNodeState) Enum() *ETournamentNodeState {
	p := new(ETournamentNodeState)
	*p = x
	return p
}
func (x ETournamentNodeState) String() string {
	return proto.EnumName(ETournamentNodeState_name, int32(x))
}
func (x *ETournamentNodeState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentNodeState_value, data, "ETournamentNodeState")
	if err != nil {
		return err
	}
	*x = ETournamentNodeState(value)
	return nil
}
func (ETournamentNodeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{4}
}

type EDOTAGroupMergeResult int32

const (
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_OK                   EDOTAGroupMergeResult = 0
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_FAILED_GENERIC       EDOTAGroupMergeResult = 1
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NOT_LEADER           EDOTAGroupMergeResult = 2
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS     EDOTAGroupMergeResult = 3
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_TOO_MANY_COACHES     EDOTAGroupMergeResult = 4
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_ENGINE_MISMATCH      EDOTAGroupMergeResult = 5
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NO_SUCH_GROUP        EDOTAGroupMergeResult = 6
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN EDOTAGroupMergeResult = 7
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_ALREADY_INVITED      EDOTAGroupMergeResult = 8
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NOT_INVITED          EDOTAGroupMergeResult = 9
)

var EDOTAGroupMergeResult_name = map[int32]string{
	0: "k_EDOTAGroupMergeResult_OK",
	1: "k_EDOTAGroupMergeResult_FAILED_GENERIC",
	2: "k_EDOTAGroupMergeResult_NOT_LEADER",
	3: "k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS",
	4: "k_EDOTAGroupMergeResult_TOO_MANY_COACHES",
	5: "k_EDOTAGroupMergeResult_ENGINE_MISMATCH",
	6: "k_EDOTAGroupMergeResult_NO_SUCH_GROUP",
	7: "k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN",
	8: "k_EDOTAGroupMergeResult_ALREADY_INVITED",
	9: "k_EDOTAGroupMergeResult_NOT_INVITED",
}
var EDOTAGroupMergeResult_value = map[string]int32{
	"k_EDOTAGroupMergeResult_OK":                   0,
	"k_EDOTAGroupMergeResult_FAILED_GENERIC":       1,
	"k_EDOTAGroupMergeResult_NOT_LEADER":           2,
	"k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS":     3,
	"k_EDOTAGroupMergeResult_TOO_MANY_COACHES":     4,
	"k_EDOTAGroupMergeResult_ENGINE_MISMATCH":      5,
	"k_EDOTAGroupMergeResult_NO_SUCH_GROUP":        6,
	"k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN": 7,
	"k_EDOTAGroupMergeResult_ALREADY_INVITED":      8,
	"k_EDOTAGroupMergeResult_NOT_INVITED":          9,
}

func (x EDOTAGroupMergeResult) Enum() *EDOTAGroupMergeResult {
	p := new(EDOTAGroupMergeResult)
	*p = x
	return p
}
func (x EDOTAGroupMergeResult) String() string {
	return proto.EnumName(EDOTAGroupMergeResult_name, int32(x))
}
func (x *EDOTAGroupMergeResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDOTAGroupMergeResult_value, data, "EDOTAGroupMergeResult")
	if err != nil {
		return err
	}
	*x = EDOTAGroupMergeResult(value)
	return nil
}
func (EDOTAGroupMergeResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dota_client_enums_41177145f6264518, []int{5}
}

func init() {
	proto.RegisterEnum("ETournamentTemplate", ETournamentTemplate_name, ETournamentTemplate_value)
	proto.RegisterEnum("ETournamentGameState", ETournamentGameState_name, ETournamentGameState_value)
	proto.RegisterEnum("ETournamentTeamState", ETournamentTeamState_name, ETournamentTeamState_value)
	proto.RegisterEnum("ETournamentState", ETournamentState_name, ETournamentState_value)
	proto.RegisterEnum("ETournamentNodeState", ETournamentNodeState_name, ETournamentNodeState_value)
	proto.RegisterEnum("EDOTAGroupMergeResult", EDOTAGroupMergeResult_name, EDOTAGroupMergeResult_value)
}

func init() {
	proto.RegisterFile("dota_client_enums.proto", fileDescriptor_dota_client_enums_41177145f6264518)
}

var fileDescriptor_dota_client_enums_41177145f6264518 = []byte{
	// 945 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd6, 0xdb, 0x72, 0xdb, 0x44,
	0x18, 0x07, 0xf0, 0xb8, 0x4e, 0xd3, 0x64, 0x81, 0x99, 0x6f, 0x96, 0x96, 0x0e, 0x0c, 0x2d, 0x6d,
	0xd3, 0x1c, 0xea, 0xb6, 0x71, 0x0e, 0x40, 0xe1, 0x52, 0xb6, 0x15, 0xdb, 0x43, 0x2c, 0x67, 0x6c,
	0xa5, 0x99, 0x70, 0x81, 0x46, 0xf1, 0x7e, 0xb1, 0x34, 0xb1, 0x76, 0x33, 0xd2, 0xaa, 0x25, 0x77,
	0x7d, 0x0d, 0xce, 0x67, 0xe8, 0x3d, 0xbc, 0x03, 0x37, 0xbc, 0x07, 0x37, 0x3c, 0x04, 0x23, 0x47,
	0xb1, 0x6c, 0x9d, 0xa2, 0xeb, 0xfd, 0x69, 0xa5, 0xfd, 0xbe, 0xdd, 0xff, 0x8a, 0xdc, 0x66, 0x42,
	0x9a, 0xc6, 0x60, 0x64, 0x23, 0x97, 0x06, 0x72, 0xdf, 0xf1, 0x36, 0xce, 0x5c, 0x21, 0x45, 0xe5,
	0x0b, 0xf2, 0xb6, 0xaa, 0x0b, 0xdf, 0xe5, 0xa6, 0x83, 0x5c, 0xea, 0xe8, 0x9c, 0x8d, 0x4c, 0x89,
	0xf4, 0x2e, 0x79, 0xef, 0xd4, 0x48, 0x19, 0x30, 0x34, 0xc1, 0x11, 0xe6, 0xe8, 0x1a, 0x59, 0x4e,
	0x1f, 0x57, 0x7c, 0x29, 0x1c, 0x53, 0x22, 0x3b, 0xb4, 0xf9, 0x0e, 0x94, 0x2a, 0xaf, 0xcb, 0xe4,
	0xe6, 0x94, 0x6b, 0x9a, 0x0e, 0xf6, 0x65, 0xf0, 0x86, 0x07, 0xe4, 0xee, 0xcc, 0x0c, 0x93, 0x11,
	0xe3, 0x80, 0x9f, 0x72, 0xf1, 0x92, 0xc3, 0x1c, 0x5d, 0x26, 0x1f, 0x64, 0x98, 0xba, 0xc9, 0x07,
	0x38, 0x42, 0x06, 0x25, 0xfa, 0x90, 0xdc, 0xcb, 0x40, 0xfd, 0x81, 0x85, 0xcc, 0x0f, 0xd4, 0x35,
	0x7a, 0x9f, 0xdc, 0xc9, 0x50, 0xca, 0x40, 0xda, 0x2f, 0x10, 0xca, 0x74, 0x85, 0xdc, 0xcf, 0x20,
	0x3d, 0x93, 0x3d, 0xb7, 0x07, 0x52, 0xb8, 0xe7, 0x70, 0x93, 0xae, 0x92, 0x07, 0x19, 0xac, 0x61,
	0xbb, 0x78, 0xe9, 0x6e, 0xd1, 0x0d, 0x52, 0xb9, 0x72, 0xba, 0xda, 0xf9, 0xae, 0x70, 0x4f, 0xd0,
	0x96, 0xf0, 0x0e, 0xad, 0x92, 0xc7, 0x57, 0xcf, 0x1b, 0x3d, 0x70, 0x9b, 0xae, 0x93, 0x87, 0x59,
	0x0b, 0x47, 0xf7, 0x05, 0xba, 0xbb, 0xa6, 0x3d, 0xf2, 0x5d, 0x84, 0xf5, 0x9c, 0x12, 0x69, 0x42,
	0x6a, 0x88, 0x0c, 0x19, 0x3c, 0xaa, 0xfc, 0xbb, 0x30, 0xd3, 0x2a, 0x1d, 0x4d, 0x27, 0xbd, 0x55,
	0x93, 0x91, 0xa9, 0x56, 0xdd, 0x23, 0xef, 0x67, 0x18, 0x4d, 0x30, 0xdc, 0x82, 0x12, 0x5d, 0xce,
	0x9c, 0x25, 0x10, 0x1d, 0xf3, 0x4b, 0x78, 0xb5, 0x48, 0x57, 0x63, 0x3d, 0x88, 0x90, 0x3a, 0xb2,
	0x1d, 0x9b, 0x07, 0x3b, 0x0b, 0xfe, 0x74, 0xe8, 0x4a, 0x6c, 0x45, 0x91, 0x0b, 0xeb, 0x83, 0x0c,
	0xfe, 0x72, 0xe8, 0x5a, 0xac, 0x57, 0x53, 0xcc, 0xe6, 0xb6, 0x67, 0x21, 0xdb, 0xf2, 0x24, 0x7c,
	0xe5, 0x17, 0x80, 0xdb, 0x9c, 0xc1, 0xd7, 0x45, 0xe0, 0x8e, 0xcb, 0xe0, 0x9b, 0x22, 0xf0, 0x43,
	0x69, 0xc1, 0xb7, 0x45, 0xe0, 0x47, 0xd2, 0x82, 0xef, 0x8a, 0xc0, 0x8f, 0xa5, 0x05, 0xdf, 0x17,
	0x81, 0xcf, 0xa4, 0x05, 0x3f, 0x14, 0x81, 0x9f, 0x48, 0x0b, 0x7e, 0x2c, 0x02, 0x3f, 0x95, 0x16,
	0xfc, 0xe4, 0xd3, 0xf5, 0x44, 0x2e, 0x24, 0x0a, 0xbe, 0x29, 0x2d, 0xf8, 0xb9, 0x90, 0xdc, 0x92,
	0x16, 0xfc, 0x52, 0x48, 0x6e, 0x4b, 0x0b, 0x7e, 0x2d, 0x24, 0x77, 0xa4, 0x05, 0xbf, 0x15, 0x92,
	0x41, 0x7b, 0x7e, 0x2f, 0x24, 0x83, 0xfe, 0xfc, 0x51, 0x48, 0x06, 0x0d, 0x7a, 0xed, 0x57, 0xfe,
	0x2b, 0x13, 0x98, 0x82, 0x17, 0xa7, 0x2c, 0x1e, 0xb9, 0xf1, 0x13, 0x16, 0xcf, 0x9d, 0xd9, 0x20,
	0xac, 0x9d, 0x2b, 0xcc, 0xb1, 0x39, 0x94, 0x12, 0x27, 0x31, 0x74, 0xc2, 0x39, 0x1b, 0xa1, 0x1c,
	0x67, 0xe1, 0x1d, 0xf2, 0x6e, 0x8a, 0xe8, 0xa0, 0x3b, 0x44, 0x06, 0xe5, 0x44, 0x5a, 0xa4, 0x65,
	0xca, 0x7c, 0x86, 0x0a, 0xd6, 0xac, 0x1c, 0x9b, 0x9c, 0x09, 0x8e, 0x0c, 0xae, 0xd3, 0x47, 0x64,
	0x25, 0x43, 0xe9, 0xb6, 0x83, 0xc2, 0x97, 0x97, 0x71, 0xb6, 0x90, 0x88, 0xb3, 0x04, 0xed, 0xe1,
	0x89, 0xcf, 0x19, 0xdc, 0xa0, 0xdb, 0x64, 0xe3, 0xaa, 0x0f, 0x6c, 0xba, 0x26, 0x97, 0x38, 0x49,
	0xed, 0x45, 0xba, 0x49, 0x9e, 0xe4, 0xcf, 0x1e, 0x7b, 0x62, 0x29, 0x71, 0x63, 0x5c, 0x3c, 0xd1,
	0xe6, 0xfb, 0xae, 0x18, 0xba, 0xe8, 0x79, 0xc0, 0x12, 0x37, 0xc6, 0x05, 0x39, 0x34, 0x6d, 0x69,
	0xf3, 0xa1, 0x2e, 0xc6, 0x15, 0x05, 0xac, 0xfc, 0x33, 0x3f, 0x13, 0xac, 0x41, 0xdc, 0xa5, 0x07,
	0xeb, 0x64, 0x24, 0xe7, 0x0e, 0x8c, 0xcc, 0xd4, 0x1d, 0x18, 0xbf, 0x6b, 0x22, 0x14, 0xac, 0xd0,
	0xd3, 0x84, 0x3c, 0x42, 0xa9, 0x78, 0x9e, 0x3d, 0xe4, 0xe3, 0x1d, 0x10, 0x6f, 0x4b, 0xe4, 0xdb,
	0xbc, 0x86, 0xf2, 0x25, 0x22, 0x0f, 0xee, 0x08, 0x0f, 0xca, 0x39, 0x34, 0x10, 0x53, 0xe5, 0x98,
	0x4f, 0xec, 0xbc, 0x88, 0x2a, 0xc6, 0xa1, 0xe0, 0x70, 0x3d, 0x47, 0xd4, 0xc6, 0x62, 0x21, 0xe7,
	0x75, 0xe3, 0x39, 0xa2, 0xfb, 0xef, 0x46, 0x0e, 0xad, 0xcd, 0xd2, 0xc5, 0xdc, 0x2f, 0xab, 0x9d,
	0x23, 0x2c, 0x25, 0x4e, 0xd7, 0xb4, 0x88, 0x36, 0x34, 0x49, 0xec, 0xd2, 0xc8, 0xcd, 0x1e, 0x90,
	0x37, 0xe8, 0x63, 0xb2, 0x96, 0x39, 0x63, 0x6c, 0xf3, 0xbf, 0x49, 0x2b, 0x64, 0xf5, 0x2a, 0x1c,
	0x6e, 0xff, 0xb7, 0x2a, 0x7f, 0x97, 0xc9, 0x2d, 0xb5, 0xd1, 0xd5, 0x95, 0xa6, 0x2b, 0xfc, 0xb3,
	0xf1, 0x26, 0xeb, 0xa1, 0xe7, 0x8f, 0x64, 0x18, 0x21, 0x69, 0x43, 0x46, 0xf7, 0x33, 0x98, 0x0b,
	0xdf, 0x92, 0x3a, 0xbe, 0xab, 0xb4, 0xf7, 0xd4, 0x86, 0xd1, 0x54, 0x35, 0xb5, 0xd7, 0xae, 0x43,
	0x29, 0x2c, 0x48, 0xaa, 0xd5, 0xba, 0xba, 0xb1, 0xa7, 0x2a, 0x0d, 0xb5, 0x07, 0xd7, 0xe8, 0x13,
	0xb2, 0x9e, 0xe5, 0xf4, 0x6e, 0xd7, 0xe8, 0x28, 0xda, 0x91, 0xb1, 0xbf, 0xa7, 0x1c, 0xa9, 0xbd,
	0x3e, 0x94, 0x0b, 0xe9, 0x7a, 0x57, 0xa9, 0xb7, 0xd4, 0x3e, 0xcc, 0x87, 0x25, 0x4c, 0xd5, 0xaa,
	0xd6, 0x6c, 0x6b, 0xaa, 0xd1, 0x69, 0xf7, 0x3b, 0x8a, 0x5e, 0x6f, 0x4d, 0xa2, 0x26, 0xe3, 0x83,
	0x8d, 0xfe, 0x41, 0xbd, 0x65, 0x34, 0x7b, 0xdd, 0x83, 0x7d, 0x58, 0x08, 0xc3, 0x20, 0xbd, 0x4e,
	0x7a, 0x4b, 0xed, 0x5d, 0xc0, 0xf1, 0x3a, 0xbb, 0xfb, 0xaa, 0x06, 0x37, 0xf2, 0xbe, 0x44, 0xd9,
	0xeb, 0xa9, 0x4a, 0xe3, 0xc8, 0x68, 0x6b, 0xcf, 0xdb, 0xba, 0xda, 0x80, 0xc5, 0xf0, 0xe7, 0x38,
	0xb3, 0x74, 0x97, 0x70, 0xa9, 0xf6, 0xac, 0x55, 0xfa, 0x7c, 0x73, 0x68, 0x4b, 0xcb, 0x3f, 0xde,
	0x18, 0x08, 0xa7, 0xba, 0xb5, 0x73, 0x5a, 0x1d, 0x8a, 0xa7, 0x9e, 0x44, 0xd3, 0x79, 0xea, 0xa2,
	0x27, 0x7c, 0x77, 0x80, 0x5e, 0x75, 0xfc, 0xa7, 0x7e, 0xec, 0x9f, 0x54, 0x83, 0x9f, 0xf8, 0xed,
	0x57, 0xa5, 0xb9, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x3e, 0xcb, 0xfe, 0xcf, 0x0b, 0x00,
	0x00,
}
