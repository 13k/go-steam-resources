// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: underlords/valveextensions.proto

package underlords

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var file_underlords_valveextensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61000,
		Name:          "underlords.map_field",
		Tag:           "varint,61000,opt,name=map_field,def=0",
		Filename:      "underlords/valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61001,
		Name:          "underlords.map_key",
		Tag:           "varint,61001,opt,name=map_key,def=0",
		Filename:      "underlords/valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         61002,
		Name:          "underlords.diff_encode_field",
		Tag:           "varint,61002,opt,name=diff_encode_field,def=0",
		Filename:      "underlords/valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61003,
		Name:          "underlords.delta_ignore",
		Tag:           "varint,61003,opt,name=delta_ignore,def=0",
		Filename:      "underlords/valveextensions.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool map_field = 61000;
	E_MapField = &file_underlords_valveextensions_proto_extTypes[0]
	// optional bool map_key = 61001;
	E_MapKey = &file_underlords_valveextensions_proto_extTypes[1]
	// optional int32 diff_encode_field = 61002;
	E_DiffEncodeField = &file_underlords_valveextensions_proto_extTypes[2]
	// optional bool delta_ignore = 61003;
	E_DeltaIgnore = &file_underlords_valveextensions_proto_extTypes[3]
)

var File_underlords_valveextensions_proto protoreflect.FileDescriptor

var file_underlords_valveextensions_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x76, 0x65, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x43, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc8, 0xdc, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x70,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x3f, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x6b, 0x65, 0x79,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc9, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06,
	0x6d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x3a, 0x4e, 0x0a, 0x11, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0xdc, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x49, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcb, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x42, 0x3e, 0x48, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x70, 0x62, 0x2f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x6f, 0x72, 0x64, 0x73, 0x80, 0x01,
	0x00,
}

var file_underlords_valveextensions_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_underlords_valveextensions_proto_depIdxs = []int32{
	0, // 0: underlords.map_field:extendee -> google.protobuf.FieldOptions
	0, // 1: underlords.map_key:extendee -> google.protobuf.FieldOptions
	0, // 2: underlords.diff_encode_field:extendee -> google.protobuf.FieldOptions
	0, // 3: underlords.delta_ignore:extendee -> google.protobuf.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_underlords_valveextensions_proto_init() }
func file_underlords_valveextensions_proto_init() {
	if File_underlords_valveextensions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_underlords_valveextensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_underlords_valveextensions_proto_goTypes,
		DependencyIndexes: file_underlords_valveextensions_proto_depIdxs,
		ExtensionInfos:    file_underlords_valveextensions_proto_extTypes,
	}.Build()
	File_underlords_valveextensions_proto = out.File
	file_underlords_valveextensions_proto_rawDesc = nil
	file_underlords_valveextensions_proto_goTypes = nil
	file_underlords_valveextensions_proto_depIdxs = nil
}
