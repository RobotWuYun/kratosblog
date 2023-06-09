// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/kratosblog/server/v1/category.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "kratosblog/api/kratosblog/data/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCategoryRequest) Reset() {
	*x = ListCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosblog_server_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryRequest) ProtoMessage() {}

func (x *ListCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosblog_server_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_kratosblog_server_v1_category_proto_rawDescGZIP(), []int{0}
}

type ListCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*v1.Category `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListCategoryReply) Reset() {
	*x = ListCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosblog_server_v1_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryReply) ProtoMessage() {}

func (x *ListCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosblog_server_v1_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryReply.ProtoReflect.Descriptor instead.
func (*ListCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_kratosblog_server_v1_category_proto_rawDescGZIP(), []int{1}
}

func (x *ListCategoryReply) GetList() []*v1.Category {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_kratosblog_server_v1_category_proto protoreflect.FileDescriptor

var file_api_kratosblog_server_v1_category_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x49, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x92, 0x01, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x85, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x44, 0x0a, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x26, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_kratosblog_server_v1_category_proto_rawDescOnce sync.Once
	file_api_kratosblog_server_v1_category_proto_rawDescData = file_api_kratosblog_server_v1_category_proto_rawDesc
)

func file_api_kratosblog_server_v1_category_proto_rawDescGZIP() []byte {
	file_api_kratosblog_server_v1_category_proto_rawDescOnce.Do(func() {
		file_api_kratosblog_server_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_kratosblog_server_v1_category_proto_rawDescData)
	})
	return file_api_kratosblog_server_v1_category_proto_rawDescData
}

var file_api_kratosblog_server_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_kratosblog_server_v1_category_proto_goTypes = []interface{}{
	(*ListCategoryRequest)(nil), // 0: api.kratosblog.server.v1.ListCategoryRequest
	(*ListCategoryReply)(nil),   // 1: api.kratosblog.server.v1.ListCategoryReply
	(*v1.Category)(nil),         // 2: api.kratosblog.data.v1.Category
}
var file_api_kratosblog_server_v1_category_proto_depIdxs = []int32{
	2, // 0: api.kratosblog.server.v1.ListCategoryReply.list:type_name -> api.kratosblog.data.v1.Category
	0, // 1: api.kratosblog.server.v1.Category.ListCategory:input_type -> api.kratosblog.server.v1.ListCategoryRequest
	1, // 2: api.kratosblog.server.v1.Category.ListCategory:output_type -> api.kratosblog.server.v1.ListCategoryReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_kratosblog_server_v1_category_proto_init() }
func file_api_kratosblog_server_v1_category_proto_init() {
	if File_api_kratosblog_server_v1_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_kratosblog_server_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryRequest); i {
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
		file_api_kratosblog_server_v1_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryReply); i {
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
			RawDescriptor: file_api_kratosblog_server_v1_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_kratosblog_server_v1_category_proto_goTypes,
		DependencyIndexes: file_api_kratosblog_server_v1_category_proto_depIdxs,
		MessageInfos:      file_api_kratosblog_server_v1_category_proto_msgTypes,
	}.Build()
	File_api_kratosblog_server_v1_category_proto = out.File
	file_api_kratosblog_server_v1_category_proto_rawDesc = nil
	file_api_kratosblog_server_v1_category_proto_goTypes = nil
	file_api_kratosblog_server_v1_category_proto_depIdxs = nil
}
