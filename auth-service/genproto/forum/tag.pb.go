// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: tag.proto

package forum

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

type TagCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TagCreateReq) Reset() {
	*x = TagCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagCreateReq) ProtoMessage() {}

func (x *TagCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagCreateReq.ProtoReflect.Descriptor instead.
func (*TagCreateReq) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TagRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TagRes) Reset() {
	*x = TagRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRes) ProtoMessage() {}

func (x *TagRes) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRes.ProtoReflect.Descriptor instead.
func (*TagRes) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TagGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag []*TagRes `protobuf:"bytes,1,rep,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TagGetAllRes) Reset() {
	*x = TagGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagGetAllRes) ProtoMessage() {}

func (x *TagGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagGetAllRes.ProtoReflect.Descriptor instead.
func (*TagGetAllRes) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{2}
}

func (x *TagGetAllRes) GetTag() []*TagRes {
	if x != nil {
		return x.Tag
	}
	return nil
}

type TagUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tag *TagCreateReq `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TagUpdate) Reset() {
	*x = TagUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagUpdate) ProtoMessage() {}

func (x *TagUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagUpdate.ProtoReflect.Descriptor instead.
func (*TagUpdate) Descriptor() ([]byte, []int) {
	return file_tag_proto_rawDescGZIP(), []int{3}
}

func (x *TagUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagUpdate) GetTag() *TagCreateReq {
	if x != nil {
		return x.Tag
	}
	return nil
}

var File_tag_proto protoreflect.FileDescriptor

var file_tag_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x22, 0x42, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x52, 0x03, 0x74, 0x61, 0x67, 0x32, 0xf4, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x11, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x13,
	0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x10, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x1a, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x66,
	0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_proto_rawDescOnce sync.Once
	file_tag_proto_rawDescData = file_tag_proto_rawDesc
)

func file_tag_proto_rawDescGZIP() []byte {
	file_tag_proto_rawDescOnce.Do(func() {
		file_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_proto_rawDescData)
	})
	return file_tag_proto_rawDescData
}

var file_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tag_proto_goTypes = []interface{}{
	(*TagCreateReq)(nil), // 0: forum.TagCreateReq
	(*TagRes)(nil),       // 1: forum.TagRes
	(*TagGetAllRes)(nil), // 2: forum.TagGetAllRes
	(*TagUpdate)(nil),    // 3: forum.TagUpdate
	(*GetByIdReq)(nil),   // 4: forum.GetByIdReq
	(*Filter)(nil),       // 5: forum.Filter
	(*Void)(nil),         // 6: forum.Void
}
var file_tag_proto_depIdxs = []int32{
	1, // 0: forum.TagGetAllRes.tag:type_name -> forum.TagRes
	0, // 1: forum.TagUpdate.tag:type_name -> forum.TagCreateReq
	0, // 2: forum.TagService.Create:input_type -> forum.TagCreateReq
	4, // 3: forum.TagService.GetById:input_type -> forum.GetByIdReq
	5, // 4: forum.TagService.GetAll:input_type -> forum.Filter
	3, // 5: forum.TagService.Update:input_type -> forum.TagUpdate
	4, // 6: forum.TagService.Delete:input_type -> forum.GetByIdReq
	1, // 7: forum.TagService.Create:output_type -> forum.TagRes
	1, // 8: forum.TagService.GetById:output_type -> forum.TagRes
	2, // 9: forum.TagService.GetAll:output_type -> forum.TagGetAllRes
	1, // 10: forum.TagService.Update:output_type -> forum.TagRes
	6, // 11: forum.TagService.Delete:output_type -> forum.Void
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tag_proto_init() }
func file_tag_proto_init() {
	if File_tag_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagCreateReq); i {
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
		file_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRes); i {
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
		file_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagGetAllRes); i {
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
		file_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagUpdate); i {
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
			RawDescriptor: file_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_proto_goTypes,
		DependencyIndexes: file_tag_proto_depIdxs,
		MessageInfos:      file_tag_proto_msgTypes,
	}.Build()
	File_tag_proto = out.File
	file_tag_proto_rawDesc = nil
	file_tag_proto_goTypes = nil
	file_tag_proto_depIdxs = nil
}