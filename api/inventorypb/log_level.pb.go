// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: inventorypb/log_level.proto

package inventorypb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Log level for exporters
type LogLevel int32

const (
	LogLevel_auto  LogLevel = 0
	LogLevel_fatal LogLevel = 1
	LogLevel_error LogLevel = 2
	LogLevel_warn  LogLevel = 3
	LogLevel_info  LogLevel = 4
	LogLevel_debug LogLevel = 5
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "auto",
		1: "fatal",
		2: "error",
		3: "warn",
		4: "info",
		5: "debug",
	}
	LogLevel_value = map[string]int32{
		"auto":  0,
		"fatal": 1,
		"error": 2,
		"warn":  3,
		"info":  4,
		"debug": 5,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_inventorypb_log_level_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_inventorypb_log_level_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_inventorypb_log_level_proto_rawDescGZIP(), []int{0}
}

var File_inventorypb_log_level_proto protoreflect.FileDescriptor

var file_inventorypb_log_level_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x6c, 0x6f,
	0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2a, 0x49, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6f, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x66, 0x61, 0x74, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x61, 0x72, 0x6e, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x10, 0x05, 0x42, 0x8a, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0d, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0xa2, 0x02,
	0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0xca, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0xe2, 0x02, 0x15, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventorypb_log_level_proto_rawDescOnce sync.Once
	file_inventorypb_log_level_proto_rawDescData = file_inventorypb_log_level_proto_rawDesc
)

func file_inventorypb_log_level_proto_rawDescGZIP() []byte {
	file_inventorypb_log_level_proto_rawDescOnce.Do(func() {
		file_inventorypb_log_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventorypb_log_level_proto_rawDescData)
	})
	return file_inventorypb_log_level_proto_rawDescData
}

var (
	file_inventorypb_log_level_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_inventorypb_log_level_proto_goTypes   = []interface{}{
		(LogLevel)(0), // 0: inventory.LogLevel
	}
)

var file_inventorypb_log_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventorypb_log_level_proto_init() }
func file_inventorypb_log_level_proto_init() {
	if File_inventorypb_log_level_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventorypb_log_level_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventorypb_log_level_proto_goTypes,
		DependencyIndexes: file_inventorypb_log_level_proto_depIdxs,
		EnumInfos:         file_inventorypb_log_level_proto_enumTypes,
	}.Build()
	File_inventorypb_log_level_proto = out.File
	file_inventorypb_log_level_proto_rawDesc = nil
	file_inventorypb_log_level_proto_goTypes = nil
	file_inventorypb_log_level_proto_depIdxs = nil
}
