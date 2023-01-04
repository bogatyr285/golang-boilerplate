// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api/v1/doer/doer.proto

package doer

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type DoAwesomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *DoAwesomeRequest) Reset() {
	*x = DoAwesomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_doer_doer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoAwesomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoAwesomeRequest) ProtoMessage() {}

func (x *DoAwesomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_doer_doer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoAwesomeRequest.ProtoReflect.Descriptor instead.
func (*DoAwesomeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_doer_doer_proto_rawDescGZIP(), []int{0}
}

func (x *DoAwesomeRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type DoAwesomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DoAwesomeResponse) Reset() {
	*x = DoAwesomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_doer_doer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoAwesomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoAwesomeResponse) ProtoMessage() {}

func (x *DoAwesomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_doer_doer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoAwesomeResponse.ProtoReflect.Descriptor instead.
func (*DoAwesomeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_doer_doer_proto_rawDescGZIP(), []int{1}
}

func (x *DoAwesomeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_v1_doer_doer_proto protoreflect.FileDescriptor

var file_api_v1_doer_doer_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x65, 0x72, 0x2f, 0x64, 0x6f,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x64, 0x6f, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x10,
	0x44, 0x6f, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x05, 0x18, 0x0a, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x6f, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x6f, 0x0a, 0x07, 0x44, 0x6f, 0x65, 0x72,
	0x41, 0x50, 0x49, 0x12, 0x64, 0x0a, 0x09, 0x44, 0x6f, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x65, 0x72, 0x2e, 0x44,
	0x6f, 0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x65, 0x72, 0x2e, 0x44, 0x6f,
	0x41, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x61,
	0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x67, 0x61, 0x74, 0x79, 0x72, 0x32, 0x38, 0x35, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_doer_doer_proto_rawDescOnce sync.Once
	file_api_v1_doer_doer_proto_rawDescData = file_api_v1_doer_doer_proto_rawDesc
)

func file_api_v1_doer_doer_proto_rawDescGZIP() []byte {
	file_api_v1_doer_doer_proto_rawDescOnce.Do(func() {
		file_api_v1_doer_doer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_doer_doer_proto_rawDescData)
	})
	return file_api_v1_doer_doer_proto_rawDescData
}

var file_api_v1_doer_doer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_doer_doer_proto_goTypes = []interface{}{
	(*DoAwesomeRequest)(nil),  // 0: api.v1.doer.DoAwesomeRequest
	(*DoAwesomeResponse)(nil), // 1: api.v1.doer.DoAwesomeResponse
}
var file_api_v1_doer_doer_proto_depIdxs = []int32{
	0, // 0: api.v1.doer.DoerAPI.DoAwesome:input_type -> api.v1.doer.DoAwesomeRequest
	1, // 1: api.v1.doer.DoerAPI.DoAwesome:output_type -> api.v1.doer.DoAwesomeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_doer_doer_proto_init() }
func file_api_v1_doer_doer_proto_init() {
	if File_api_v1_doer_doer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_doer_doer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoAwesomeRequest); i {
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
		file_api_v1_doer_doer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoAwesomeResponse); i {
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
			RawDescriptor: file_api_v1_doer_doer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_doer_doer_proto_goTypes,
		DependencyIndexes: file_api_v1_doer_doer_proto_depIdxs,
		MessageInfos:      file_api_v1_doer_doer_proto_msgTypes,
	}.Build()
	File_api_v1_doer_doer_proto = out.File
	file_api_v1_doer_doer_proto_rawDesc = nil
	file_api_v1_doer_doer_proto_goTypes = nil
	file_api_v1_doer_doer_proto_depIdxs = nil
}
