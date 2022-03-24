// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: teamcloud/v1/teamcloud.proto

package teamcloudv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetTeamCloudInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTeamCloudInfoRequest) Reset() {
	*x = GetTeamCloudInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamCloudInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamCloudInfoRequest) ProtoMessage() {}

func (x *GetTeamCloudInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamCloudInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTeamCloudInfoRequest) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_teamcloud_proto_rawDescGZIP(), []int{0}
}

type GetTeamCloudInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status string             `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Data   *TeamCloudInfoData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTeamCloudInfoResponse) Reset() {
	*x = GetTeamCloudInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamCloudInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamCloudInfoResponse) ProtoMessage() {}

func (x *GetTeamCloudInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamCloudInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTeamCloudInfoResponse) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_teamcloud_proto_rawDescGZIP(), []int{1}
}

func (x *GetTeamCloudInfoResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTeamCloudInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetTeamCloudInfoResponse) GetData() *TeamCloudInfoData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TeamCloudInfoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageVersion    string `protobuf:"bytes,1,opt,name=imageVersion,proto3" json:"imageVersion,omitempty"`
	TemplateVersion string `protobuf:"bytes,2,opt,name=templateVersion,proto3" json:"templateVersion,omitempty"`
}

func (x *TeamCloudInfoData) Reset() {
	*x = TeamCloudInfoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamCloudInfoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamCloudInfoData) ProtoMessage() {}

func (x *TeamCloudInfoData) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_teamcloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamCloudInfoData.ProtoReflect.Descriptor instead.
func (*TeamCloudInfoData) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_teamcloud_proto_rawDescGZIP(), []int{2}
}

func (x *TeamCloudInfoData) GetImageVersion() string {
	if x != nil {
		return x.ImageVersion
	}
	return ""
}

func (x *TeamCloudInfoData) GetTemplateVersion() string {
	if x != nil {
		return x.TemplateVersion
	}
	return ""
}

var File_teamcloud_v1_teamcloud_proto protoreflect.FileDescriptor

var file_teamcloud_v1_teamcloud_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x61, 0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0xbb, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x41, 0x50, 0x49, 0x12, 0xaa, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6c,
	0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d,
	0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x3a,
	0x01, 0x2a, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2f,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamcloud_v1_teamcloud_proto_rawDescOnce sync.Once
	file_teamcloud_v1_teamcloud_proto_rawDescData = file_teamcloud_v1_teamcloud_proto_rawDesc
)

func file_teamcloud_v1_teamcloud_proto_rawDescGZIP() []byte {
	file_teamcloud_v1_teamcloud_proto_rawDescOnce.Do(func() {
		file_teamcloud_v1_teamcloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamcloud_v1_teamcloud_proto_rawDescData)
	})
	return file_teamcloud_v1_teamcloud_proto_rawDescData
}

var file_teamcloud_v1_teamcloud_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_teamcloud_v1_teamcloud_proto_goTypes = []interface{}{
	(*GetTeamCloudInfoRequest)(nil),  // 0: colbylwilliams.teamcloud.v1.GetTeamCloudInfoRequest
	(*GetTeamCloudInfoResponse)(nil), // 1: colbylwilliams.teamcloud.v1.GetTeamCloudInfoResponse
	(*TeamCloudInfoData)(nil),        // 2: colbylwilliams.teamcloud.v1.TeamCloudInfoData
}
var file_teamcloud_v1_teamcloud_proto_depIdxs = []int32{
	2, // 0: colbylwilliams.teamcloud.v1.GetTeamCloudInfoResponse.data:type_name -> colbylwilliams.teamcloud.v1.TeamCloudInfoData
	0, // 1: colbylwilliams.teamcloud.v1.TeamCloudAPI.GetTeamCloudInfo:input_type -> colbylwilliams.teamcloud.v1.GetTeamCloudInfoRequest
	1, // 2: colbylwilliams.teamcloud.v1.TeamCloudAPI.GetTeamCloudInfo:output_type -> colbylwilliams.teamcloud.v1.GetTeamCloudInfoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_teamcloud_proto_init() }
func file_teamcloud_v1_teamcloud_proto_init() {
	if File_teamcloud_v1_teamcloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teamcloud_v1_teamcloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamCloudInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teamcloud_v1_teamcloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamCloudInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teamcloud_v1_teamcloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamCloudInfoData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teamcloud_v1_teamcloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teamcloud_v1_teamcloud_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_teamcloud_proto_depIdxs,
		MessageInfos:      file_teamcloud_v1_teamcloud_proto_msgTypes,
	}.Build()
	File_teamcloud_v1_teamcloud_proto = out.File
	file_teamcloud_v1_teamcloud_proto_rawDesc = nil
	file_teamcloud_v1_teamcloud_proto_goTypes = nil
	file_teamcloud_v1_teamcloud_proto_depIdxs = nil
}