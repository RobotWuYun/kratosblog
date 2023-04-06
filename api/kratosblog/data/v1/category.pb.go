// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/kratosblog/data/v1/category.proto

package v1

import (
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

type CategoryStatus int32

const (
	CategoryStatus_CategoryStatusUnknown CategoryStatus = 0
	CategoryStatus_CategoryStatusNormal  CategoryStatus = 1
	CategoryStatus_CategoryStatusHide    CategoryStatus = 2
)

// Enum value maps for CategoryStatus.
var (
	CategoryStatus_name = map[int32]string{
		0: "CategoryStatusUnknown",
		1: "CategoryStatusNormal",
		2: "CategoryStatusHide",
	}
	CategoryStatus_value = map[string]int32{
		"CategoryStatusUnknown": 0,
		"CategoryStatusNormal":  1,
		"CategoryStatusHide":    2,
	}
)

func (x CategoryStatus) Enum() *CategoryStatus {
	p := new(CategoryStatus)
	*p = x
	return p
}

func (x CategoryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CategoryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_kratosblog_data_v1_category_proto_enumTypes[0].Descriptor()
}

func (CategoryStatus) Type() protoreflect.EnumType {
	return &file_api_kratosblog_data_v1_category_proto_enumTypes[0]
}

func (x CategoryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CategoryStatus.Descriptor instead.
func (CategoryStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_kratosblog_data_v1_category_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover        string `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Color        string `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	ArticleCount string `protobuf:"bytes,6,opt,name=article_count,json=articleCount,proto3" json:"article_count,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_kratosblog_data_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_api_kratosblog_data_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_api_kratosblog_data_v1_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Category) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Category) GetArticleCount() string {
	if x != nil {
		return x.ArticleCount
	}
	return ""
}

var File_api_kratosblog_data_v1_category_proto protoreflect.FileDescriptor

var file_api_kratosblog_data_v1_category_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22,
	0xa1, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x5d, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x64, 0x65,
	0x10, 0x02, 0x42, 0x40, 0x0a, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x24,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_kratosblog_data_v1_category_proto_rawDescOnce sync.Once
	file_api_kratosblog_data_v1_category_proto_rawDescData = file_api_kratosblog_data_v1_category_proto_rawDesc
)

func file_api_kratosblog_data_v1_category_proto_rawDescGZIP() []byte {
	file_api_kratosblog_data_v1_category_proto_rawDescOnce.Do(func() {
		file_api_kratosblog_data_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_kratosblog_data_v1_category_proto_rawDescData)
	})
	return file_api_kratosblog_data_v1_category_proto_rawDescData
}

var file_api_kratosblog_data_v1_category_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_kratosblog_data_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_kratosblog_data_v1_category_proto_goTypes = []interface{}{
	(CategoryStatus)(0), // 0: api.kratosblog.data.v1.CategoryStatus
	(*Category)(nil),    // 1: api.kratosblog.data.v1.Category
}
var file_api_kratosblog_data_v1_category_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_kratosblog_data_v1_category_proto_init() }
func file_api_kratosblog_data_v1_category_proto_init() {
	if File_api_kratosblog_data_v1_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_kratosblog_data_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_api_kratosblog_data_v1_category_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_kratosblog_data_v1_category_proto_goTypes,
		DependencyIndexes: file_api_kratosblog_data_v1_category_proto_depIdxs,
		EnumInfos:         file_api_kratosblog_data_v1_category_proto_enumTypes,
		MessageInfos:      file_api_kratosblog_data_v1_category_proto_msgTypes,
	}.Build()
	File_api_kratosblog_data_v1_category_proto = out.File
	file_api_kratosblog_data_v1_category_proto_rawDesc = nil
	file_api_kratosblog_data_v1_category_proto_goTypes = nil
	file_api_kratosblog_data_v1_category_proto_depIdxs = nil
}