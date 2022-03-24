//
//TeamCloud
//
//API for working with a TeamCloud instance.
//
//The version of the OpenAPI document: v1
//
//Contact: colbyw@microsoft.com
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: teamcloud/v1/adapter_service.proto

package teamcloudv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_teamcloud_v1_adapter_service_proto protoreflect.FileDescriptor

var file_teamcloud_v1_adapter_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c,
	0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37,
	0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x72, 0x0a, 0x0a, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x64, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x3d, 0x2e,
	0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x59, 0x5a, 0x57,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x62, 0x79,
	0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x76, 0x31, 0x50, 0x01, 0x50, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_teamcloud_v1_adapter_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                    // 0: google.protobuf.Empty
	(*AdapterInformationListDataResult)(nil), // 1: colbylwilliams.teamcloud.v1.AdapterInformationListDataResult
}
var file_teamcloud_v1_adapter_service_proto_depIdxs = []int32{
	0, // 0: colbylwilliams.teamcloud.v1.AdapterAPI.GetAdapters:input_type -> google.protobuf.Empty
	1, // 1: colbylwilliams.teamcloud.v1.AdapterAPI.GetAdapters:output_type -> colbylwilliams.teamcloud.v1.AdapterInformationListDataResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_adapter_service_proto_init() }
func file_teamcloud_v1_adapter_service_proto_init() {
	if File_teamcloud_v1_adapter_service_proto != nil {
		return
	}
	file_teamcloud_v1_adapter_information_list_data_result_proto_init()
	file_teamcloud_v1_error_result_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teamcloud_v1_adapter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teamcloud_v1_adapter_service_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_adapter_service_proto_depIdxs,
	}.Build()
	File_teamcloud_v1_adapter_service_proto = out.File
	file_teamcloud_v1_adapter_service_proto_rawDesc = nil
	file_teamcloud_v1_adapter_service_proto_goTypes = nil
	file_teamcloud_v1_adapter_service_proto_depIdxs = nil
}
