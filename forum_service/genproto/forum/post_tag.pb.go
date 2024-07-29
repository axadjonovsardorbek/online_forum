// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: post_tag.proto

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

type PostTagsCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	TagId  string `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
}

func (x *PostTagsCreateReq) Reset() {
	*x = PostTagsCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTagsCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTagsCreateReq) ProtoMessage() {}

func (x *PostTagsCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTagsCreateReq.ProtoReflect.Descriptor instead.
func (*PostTagsCreateReq) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{0}
}

func (x *PostTagsCreateReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostTagsCreateReq) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type PostTagsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Post *PostRes `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	Tag  *TagRes  `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PostTagsRes) Reset() {
	*x = PostTagsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTagsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTagsRes) ProtoMessage() {}

func (x *PostTagsRes) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTagsRes.ProtoReflect.Descriptor instead.
func (*PostTagsRes) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{1}
}

func (x *PostTagsRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostTagsRes) GetPost() *PostRes {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostTagsRes) GetTag() *TagRes {
	if x != nil {
		return x.Tag
	}
	return nil
}

type PostTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	TagId  string `protobuf:"bytes,3,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
}

func (x *PostTags) Reset() {
	*x = PostTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTags) ProtoMessage() {}

func (x *PostTags) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTags.ProtoReflect.Descriptor instead.
func (*PostTags) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{2}
}

func (x *PostTags) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostTags) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostTags) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type PostTagsGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posttag []*PostTagsRes `protobuf:"bytes,1,rep,name=posttag,proto3" json:"posttag,omitempty"`
}

func (x *PostTagsGetAllRes) Reset() {
	*x = PostTagsGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTagsGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTagsGetAllRes) ProtoMessage() {}

func (x *PostTagsGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTagsGetAllRes.ProtoReflect.Descriptor instead.
func (*PostTagsGetAllRes) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{3}
}

func (x *PostTagsGetAllRes) GetPosttag() []*PostTagsRes {
	if x != nil {
		return x.Posttag
	}
	return nil
}

type PostTagsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Posttag *PostTagsCreateReq `protobuf:"bytes,2,opt,name=posttag,proto3" json:"posttag,omitempty"`
}

func (x *PostTagsUpdate) Reset() {
	*x = PostTagsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTagsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTagsUpdate) ProtoMessage() {}

func (x *PostTagsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTagsUpdate.ProtoReflect.Descriptor instead.
func (*PostTagsUpdate) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{4}
}

func (x *PostTagsUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostTagsUpdate) GetPosttag() *PostTagsCreateReq {
	if x != nil {
		return x.Posttag
	}
	return nil
}

type PopularTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PopularTag) Reset() {
	*x = PopularTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopularTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopularTag) ProtoMessage() {}

func (x *PopularTag) ProtoReflect() protoreflect.Message {
	mi := &file_post_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopularTag.ProtoReflect.Descriptor instead.
func (*PopularTag) Descriptor() ([]byte, []int) {
	return file_post_tag_proto_rawDescGZIP(), []int{5}
}

func (x *PopularTag) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PopularTag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PopularTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_post_tag_proto protoreflect.FileDescriptor

var file_post_tag_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x09, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x4a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x52, 0x07, 0x70, 0x6f, 0x73,
	0x74, 0x74, 0x61, 0x67, 0x22, 0x54, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x74, 0x61, 0x67, 0x22, 0x46, 0x0a, 0x0a, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x72, 0x54, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0x82, 0x03, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x66, 0x6f,
	0x72, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x66, 0x6f,
	0x72, 0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x79, 0x54, 0x61, 0x67, 0x12, 0x10, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x72, 0x54, 0x61, 0x67, 0x12, 0x0b, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x1a, 0x11, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x72, 0x54, 0x61, 0x67, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_post_tag_proto_rawDescOnce sync.Once
	file_post_tag_proto_rawDescData = file_post_tag_proto_rawDesc
)

func file_post_tag_proto_rawDescGZIP() []byte {
	file_post_tag_proto_rawDescOnce.Do(func() {
		file_post_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_tag_proto_rawDescData)
	})
	return file_post_tag_proto_rawDescData
}

var file_post_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_post_tag_proto_goTypes = []interface{}{
	(*PostTagsCreateReq)(nil), // 0: forum.PostTagsCreateReq
	(*PostTagsRes)(nil),       // 1: forum.PostTagsRes
	(*PostTags)(nil),          // 2: forum.PostTags
	(*PostTagsGetAllRes)(nil), // 3: forum.PostTagsGetAllRes
	(*PostTagsUpdate)(nil),    // 4: forum.PostTagsUpdate
	(*PopularTag)(nil),        // 5: forum.PopularTag
	(*PostRes)(nil),           // 6: forum.PostRes
	(*TagRes)(nil),            // 7: forum.TagRes
	(*GetByIdReq)(nil),        // 8: forum.GetByIdReq
	(*Filter)(nil),            // 9: forum.Filter
	(*GetFilter)(nil),         // 10: forum.GetFilter
	(*Void)(nil),              // 11: forum.Void
}
var file_post_tag_proto_depIdxs = []int32{
	6,  // 0: forum.PostTagsRes.post:type_name -> forum.PostRes
	7,  // 1: forum.PostTagsRes.tag:type_name -> forum.TagRes
	1,  // 2: forum.PostTagsGetAllRes.posttag:type_name -> forum.PostTagsRes
	0,  // 3: forum.PostTagsUpdate.posttag:type_name -> forum.PostTagsCreateReq
	0,  // 4: forum.PostTagsService.Create:input_type -> forum.PostTagsCreateReq
	8,  // 5: forum.PostTagsService.GetById:input_type -> forum.GetByIdReq
	9,  // 6: forum.PostTagsService.GetAll:input_type -> forum.Filter
	4,  // 7: forum.PostTagsService.Update:input_type -> forum.PostTagsUpdate
	8,  // 8: forum.PostTagsService.Delete:input_type -> forum.GetByIdReq
	10, // 9: forum.PostTagsService.GetPostByTag:input_type -> forum.GetFilter
	11, // 10: forum.PostTagsService.GetPopularTag:input_type -> forum.Void
	2,  // 11: forum.PostTagsService.Create:output_type -> forum.PostTags
	1,  // 12: forum.PostTagsService.GetById:output_type -> forum.PostTagsRes
	3,  // 13: forum.PostTagsService.GetAll:output_type -> forum.PostTagsGetAllRes
	2,  // 14: forum.PostTagsService.Update:output_type -> forum.PostTags
	11, // 15: forum.PostTagsService.Delete:output_type -> forum.Void
	3,  // 16: forum.PostTagsService.GetPostByTag:output_type -> forum.PostTagsGetAllRes
	5,  // 17: forum.PostTagsService.GetPopularTag:output_type -> forum.PopularTag
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_post_tag_proto_init() }
func file_post_tag_proto_init() {
	if File_post_tag_proto != nil {
		return
	}
	file_common_proto_init()
	file_post_proto_init()
	file_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_post_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTagsCreateReq); i {
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
		file_post_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTagsRes); i {
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
		file_post_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTags); i {
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
		file_post_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTagsGetAllRes); i {
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
		file_post_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTagsUpdate); i {
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
		file_post_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopularTag); i {
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
			RawDescriptor: file_post_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_tag_proto_goTypes,
		DependencyIndexes: file_post_tag_proto_depIdxs,
		MessageInfos:      file_post_tag_proto_msgTypes,
	}.Build()
	File_post_tag_proto = out.File
	file_post_tag_proto_rawDesc = nil
	file_post_tag_proto_goTypes = nil
	file_post_tag_proto_depIdxs = nil
}