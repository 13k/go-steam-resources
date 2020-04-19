// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: dota2/dota_hud_types.proto

package dota2

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type EHeroSelectionText int32

const (
	EHeroSelectionText_k_EHeroSelectionText_Invalid                              EHeroSelectionText = -1
	EHeroSelectionText_k_EHeroSelectionText_None                                 EHeroSelectionText = 0
	EHeroSelectionText_k_EHeroSelectionText_ChooseHero                           EHeroSelectionText = 1
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_YouFirst           EHeroSelectionText = 2
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_Planning_TheyFirst          EHeroSelectionText = 3
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_BanSelected_YouFirst        EHeroSelectionText = 4
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst       EHeroSelectionText = 5
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_YouPicking                  EHeroSelectionText = 6
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TheyPicking                 EHeroSelectionText = 7
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TeammateRandomed            EHeroSelectionText = 8
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_YouPicking_LosingGold       EHeroSelectionText = 9
	EHeroSelectionText_k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold      EHeroSelectionText = 10
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_ChooseCaptain           EHeroSelectionText = 11
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain EHeroSelectionText = 12
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouSelect               EHeroSelectionText = 13
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheySelect              EHeroSelectionText = 14
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_YouBan                  EHeroSelectionText = 15
	EHeroSelectionText_k_EHeroSelectionText_CaptainsMode_TheyBan                 EHeroSelectionText = 16
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_HeroReview               EHeroSelectionText = 17
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_RoundDisplay             EHeroSelectionText = 18
	EHeroSelectionText_k_EHeroSelectionText_RandomDraft_Waiting                  EHeroSelectionText = 19
)

// Enum value maps for EHeroSelectionText.
var (
	EHeroSelectionText_name = map[int32]string{
		-1: "k_EHeroSelectionText_Invalid",
		0:  "k_EHeroSelectionText_None",
		1:  "k_EHeroSelectionText_ChooseHero",
		2:  "k_EHeroSelectionText_AllDraft_Planning_YouFirst",
		3:  "k_EHeroSelectionText_AllDraft_Planning_TheyFirst",
		4:  "k_EHeroSelectionText_AllDraft_BanSelected_YouFirst",
		5:  "k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst",
		6:  "k_EHeroSelectionText_AllDraft_YouPicking",
		7:  "k_EHeroSelectionText_AllDraft_TheyPicking",
		8:  "k_EHeroSelectionText_AllDraft_TeammateRandomed",
		9:  "k_EHeroSelectionText_AllDraft_YouPicking_LosingGold",
		10: "k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold",
		11: "k_EHeroSelectionText_CaptainsMode_ChooseCaptain",
		12: "k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain",
		13: "k_EHeroSelectionText_CaptainsMode_YouSelect",
		14: "k_EHeroSelectionText_CaptainsMode_TheySelect",
		15: "k_EHeroSelectionText_CaptainsMode_YouBan",
		16: "k_EHeroSelectionText_CaptainsMode_TheyBan",
		17: "k_EHeroSelectionText_RandomDraft_HeroReview",
		18: "k_EHeroSelectionText_RandomDraft_RoundDisplay",
		19: "k_EHeroSelectionText_RandomDraft_Waiting",
	}
	EHeroSelectionText_value = map[string]int32{
		"k_EHeroSelectionText_Invalid":                              -1,
		"k_EHeroSelectionText_None":                                 0,
		"k_EHeroSelectionText_ChooseHero":                           1,
		"k_EHeroSelectionText_AllDraft_Planning_YouFirst":           2,
		"k_EHeroSelectionText_AllDraft_Planning_TheyFirst":          3,
		"k_EHeroSelectionText_AllDraft_BanSelected_YouFirst":        4,
		"k_EHeroSelectionText_AllDraft_BanSelected_TheyFirst":       5,
		"k_EHeroSelectionText_AllDraft_YouPicking":                  6,
		"k_EHeroSelectionText_AllDraft_TheyPicking":                 7,
		"k_EHeroSelectionText_AllDraft_TeammateRandomed":            8,
		"k_EHeroSelectionText_AllDraft_YouPicking_LosingGold":       9,
		"k_EHeroSelectionText_AllDraft_TheyPicking_LosingGold":      10,
		"k_EHeroSelectionText_CaptainsMode_ChooseCaptain":           11,
		"k_EHeroSelectionText_CaptainsMode_WaitingForChooseCaptain": 12,
		"k_EHeroSelectionText_CaptainsMode_YouSelect":               13,
		"k_EHeroSelectionText_CaptainsMode_TheySelect":              14,
		"k_EHeroSelectionText_CaptainsMode_YouBan":                  15,
		"k_EHeroSelectionText_CaptainsMode_TheyBan":                 16,
		"k_EHeroSelectionText_RandomDraft_HeroReview":               17,
		"k_EHeroSelectionText_RandomDraft_RoundDisplay":             18,
		"k_EHeroSelectionText_RandomDraft_Waiting":                  19,
	}
)

func (x EHeroSelectionText) Enum() *EHeroSelectionText {
	p := new(EHeroSelectionText)
	*p = x
	return p
}

func (x EHeroSelectionText) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EHeroSelectionText) Descriptor() protoreflect.EnumDescriptor {
	return file_dota2_dota_hud_types_proto_enumTypes[0].Descriptor()
}

func (EHeroSelectionText) Type() protoreflect.EnumType {
	return &file_dota2_dota_hud_types_proto_enumTypes[0]
}

func (x EHeroSelectionText) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EHeroSelectionText) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EHeroSelectionText(num)
	return nil
}

// Deprecated: Use EHeroSelectionText.Descriptor instead.
func (EHeroSelectionText) EnumDescriptor() ([]byte, []int) {
	return file_dota2_dota_hud_types_proto_rawDescGZIP(), []int{0}
}

var file_dota2_dota_hud_types_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50501,
		Name:          "dota2.hud_localize_token",
		Tag:           "bytes,50501,opt,name=hud_localize_token",
		Filename:      "dota2/dota_hud_types.proto",
	},
}

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional string hud_localize_token = 50501;
	E_HudLocalizeToken = &file_dota2_dota_hud_types_proto_extTypes[0]
)

var File_dota2_dota_hud_types_proto protoreflect.FileDescriptor

var file_dota2_dota_hud_types_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x68, 0x75, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x6f,
	0x74, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf7, 0x0f, 0x0a, 0x12, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x1c,
	0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x78, 0x74, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x6b, 0x5f, 0x45, 0x48, 0x65,
	0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x48, 0x0a, 0x1f, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72,
	0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x43,
	0x68, 0x6f, 0x6f, 0x73, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x10, 0x01, 0x1a, 0x23, 0xaa, 0xd4, 0x18,
	0x1f, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x48, 0x65, 0x72, 0x6f,
	0x12, 0x68, 0x0a, 0x2f, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66,
	0x74, 0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x59, 0x6f, 0x75, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x10, 0x02, 0x1a, 0x33, 0xaa, 0xd4, 0x18, 0x2f, 0x23, 0x44, 0x4f, 0x54, 0x41,
	0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x59, 0x6f, 0x75, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x6a, 0x0a, 0x30, 0x6b, 0x5f,
	0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65,
	0x78, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x50, 0x6c, 0x61, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x10, 0x03,
	0x1a, 0x34, 0xaa, 0xd4, 0x18, 0x30, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f,
	0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x54, 0x68, 0x65,
	0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x6e, 0x0a, 0x32, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72,
	0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41,
	0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x42, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x59, 0x6f, 0x75, 0x46, 0x69, 0x72, 0x73, 0x74, 0x10, 0x04, 0x1a, 0x36,
	0xaa, 0xd4, 0x18, 0x32, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66,
	0x74, 0x5f, 0x42, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x59, 0x6f,
	0x75, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x70, 0x0a, 0x33, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72,
	0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41,
	0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x42, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x10, 0x05, 0x1a,
	0x37, 0xaa, 0xd4, 0x18, 0x33, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61,
	0x66, 0x74, 0x5f, 0x42, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x54,
	0x68, 0x65, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x48,
	0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74,
	0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x59, 0x6f, 0x75, 0x50, 0x69, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x1a, 0x2c, 0xaa, 0xd4, 0x18, 0x28, 0x23, 0x44, 0x4f, 0x54,
	0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x59, 0x6f, 0x75, 0x50, 0x69, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x29, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41, 0x6c, 0x6c,
	0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x10, 0x07, 0x1a, 0x2d, 0xaa, 0xd4, 0x18, 0x29, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48,
	0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c,
	0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x6f, 0x0a, 0x2e, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x54, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x65, 0x64, 0x10, 0x08, 0x1a, 0x3b, 0xaa, 0xd4, 0x18, 0x37, 0x23, 0x44, 0x4f, 0x54,
	0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x54, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x65, 0x64, 0x5f, 0x50, 0x61, 0x6e, 0x6f, 0x72,
	0x61, 0x6d, 0x61, 0x12, 0x70, 0x0a, 0x33, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x44,
	0x72, 0x61, 0x66, 0x74, 0x5f, 0x59, 0x6f, 0x75, 0x50, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x4c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6c, 0x64, 0x10, 0x09, 0x1a, 0x37, 0xaa, 0xd4,
	0x18, 0x33, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f,
	0x59, 0x6f, 0x75, 0x50, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x4c, 0x6f, 0x73, 0x69, 0x6e,
	0x67, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x72, 0x0a, 0x34, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x41, 0x6c,
	0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x4c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6c, 0x64, 0x10, 0x0a, 0x1a,
	0x38, 0xaa, 0xd4, 0x18, 0x34, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61,
	0x66, 0x74, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x4c,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x68, 0x0a, 0x2f, 0x6b, 0x5f, 0x45,
	0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78,
	0x74, 0x5f, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x43,
	0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x10, 0x0b, 0x1a, 0x33,
	0xaa, 0xd4, 0x18, 0x2f, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e,
	0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x12, 0x7c, 0x0a, 0x39, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x43, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x72, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e,
	0x10, 0x0c, 0x1a, 0x3d, 0xaa, 0xd4, 0x18, 0x39, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x43, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x65, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69,
	0x6e, 0x12, 0x60, 0x0a, 0x2b, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x43, 0x61, 0x70, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x59, 0x6f, 0x75, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x10, 0x0d, 0x1a, 0x2f, 0xaa, 0xd4, 0x18, 0x2b, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x43, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x59, 0x6f, 0x75, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x12, 0x62, 0x0a, 0x2c, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x43, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x10, 0x0e, 0x1a, 0x30, 0xaa, 0xd4, 0x18, 0x2c, 0x23, 0x44, 0x4f, 0x54, 0x41,
	0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x68, 0x65,
	0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x5a, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x48, 0x65,
	0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f,
	0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x59, 0x6f, 0x75,
	0x42, 0x61, 0x6e, 0x10, 0x0f, 0x1a, 0x2c, 0xaa, 0xd4, 0x18, 0x28, 0x23, 0x44, 0x4f, 0x54, 0x41,
	0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x43, 0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x59, 0x6f, 0x75,
	0x42, 0x61, 0x6e, 0x12, 0x5c, 0x0a, 0x29, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x43, 0x61, 0x70, 0x74,
	0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x42, 0x61, 0x6e,
	0x10, 0x10, 0x1a, 0x2d, 0xaa, 0xd4, 0x18, 0x29, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x43, 0x61, 0x70,
	0x74, 0x61, 0x69, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x54, 0x68, 0x65, 0x79, 0x42, 0x61,
	0x6e, 0x12, 0x60, 0x0a, 0x2b, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x10, 0x11, 0x1a, 0x2f, 0xaa, 0xd4, 0x18, 0x2b, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x48, 0x65,
	0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x61, 0x0a, 0x2d, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x10, 0x12, 0x1a, 0x2e, 0xaa, 0xd4, 0x18, 0x2a, 0x23, 0x44, 0x4f, 0x54,
	0x41, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x41, 0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x57, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x48, 0x65, 0x72,
	0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x52,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x57, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x10, 0x13, 0x1a, 0x29, 0xaa, 0xd4, 0x18, 0x25, 0x23, 0x44, 0x4f, 0x54, 0x41, 0x5f,
	0x48, 0x65, 0x72, 0x6f, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x41,
	0x6c, 0x6c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x3a,
	0x51, 0x0a, 0x12, 0x68, 0x75, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc5, 0x8a, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x68, 0x75, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x35, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x6f, 0x74, 0x61, 0x32, 0x80, 0x01, 0x00,
}

var (
	file_dota2_dota_hud_types_proto_rawDescOnce sync.Once
	file_dota2_dota_hud_types_proto_rawDescData = file_dota2_dota_hud_types_proto_rawDesc
)

func file_dota2_dota_hud_types_proto_rawDescGZIP() []byte {
	file_dota2_dota_hud_types_proto_rawDescOnce.Do(func() {
		file_dota2_dota_hud_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota2_dota_hud_types_proto_rawDescData)
	})
	return file_dota2_dota_hud_types_proto_rawDescData
}

var file_dota2_dota_hud_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota2_dota_hud_types_proto_goTypes = []interface{}{
	(EHeroSelectionText)(0),             // 0: dota2.EHeroSelectionText
	(*descriptor.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_dota2_dota_hud_types_proto_depIdxs = []int32{
	1, // 0: dota2.hud_localize_token:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dota2_dota_hud_types_proto_init() }
func file_dota2_dota_hud_types_proto_init() {
	if File_dota2_dota_hud_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dota2_dota_hud_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_dota2_dota_hud_types_proto_goTypes,
		DependencyIndexes: file_dota2_dota_hud_types_proto_depIdxs,
		EnumInfos:         file_dota2_dota_hud_types_proto_enumTypes,
		ExtensionInfos:    file_dota2_dota_hud_types_proto_extTypes,
	}.Build()
	File_dota2_dota_hud_types_proto = out.File
	file_dota2_dota_hud_types_proto_rawDesc = nil
	file_dota2_dota_hud_types_proto_goTypes = nil
	file_dota2_dota_hud_types_proto_depIdxs = nil
}
