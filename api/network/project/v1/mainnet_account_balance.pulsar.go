// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package projectv1

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: network/project/v1/mainnet_account_balance.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_network_project_v1_mainnet_account_balance_proto protoreflect.FileDescriptor

var file_network_project_v1_mainnet_account_balance_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0xda, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x1a, 0x4d, 0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x67, 0x6e, 0x69,
	0x74, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x50,
	0x58, 0xaa, 0x02, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_project_v1_mainnet_account_balance_proto_goTypes = []interface{}{}
	file_network_project_v1_mainnet_account_balance_proto_depIdxs = []int32{
		0, // [0:0] is the sub-list for method output_type
		0, // [0:0] is the sub-list for method input_type
		0, // [0:0] is the sub-list for extension type_name
		0, // [0:0] is the sub-list for extension extendee
		0, // [0:0] is the sub-list for field type_name
	}
)

func init() { file_network_project_v1_mainnet_account_balance_proto_init() }
func file_network_project_v1_mainnet_account_balance_proto_init() {
	if File_network_project_v1_mainnet_account_balance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_project_v1_mainnet_account_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_network_project_v1_mainnet_account_balance_proto_goTypes,
		DependencyIndexes: file_network_project_v1_mainnet_account_balance_proto_depIdxs,
	}.Build()
	File_network_project_v1_mainnet_account_balance_proto = out.File
	file_network_project_v1_mainnet_account_balance_proto_rawDesc = nil
	file_network_project_v1_mainnet_account_balance_proto_goTypes = nil
	file_network_project_v1_mainnet_account_balance_proto_depIdxs = nil
}
