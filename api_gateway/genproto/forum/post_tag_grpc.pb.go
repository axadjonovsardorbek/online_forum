// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: post_tag.proto

package forum

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PostTagsService_Create_FullMethodName        = "/forum.PostTagsService/Create"
	PostTagsService_GetById_FullMethodName       = "/forum.PostTagsService/GetById"
	PostTagsService_GetAll_FullMethodName        = "/forum.PostTagsService/GetAll"
	PostTagsService_Update_FullMethodName        = "/forum.PostTagsService/Update"
	PostTagsService_Delete_FullMethodName        = "/forum.PostTagsService/Delete"
	PostTagsService_GetPostByTag_FullMethodName  = "/forum.PostTagsService/GetPostByTag"
	PostTagsService_GetPopularTag_FullMethodName = "/forum.PostTagsService/GetPopularTag"
)

// PostTagsServiceClient is the client API for PostTagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostTagsServiceClient interface {
	Create(ctx context.Context, in *PostTagsCreateReq, opts ...grpc.CallOption) (*PostTags, error)
	GetById(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*PostTagsRes, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*PostTagsGetAllRes, error)
	Update(ctx context.Context, in *PostTagsUpdate, opts ...grpc.CallOption) (*PostTags, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
	GetPostByTag(ctx context.Context, in *GetFilter, opts ...grpc.CallOption) (*PostTagsGetAllRes, error)
	GetPopularTag(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PopularTag, error)
}

type postTagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostTagsServiceClient(cc grpc.ClientConnInterface) PostTagsServiceClient {
	return &postTagsServiceClient{cc}
}

func (c *postTagsServiceClient) Create(ctx context.Context, in *PostTagsCreateReq, opts ...grpc.CallOption) (*PostTags, error) {
	out := new(PostTags)
	err := c.cc.Invoke(ctx, PostTagsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) GetById(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*PostTagsRes, error) {
	out := new(PostTagsRes)
	err := c.cc.Invoke(ctx, PostTagsService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*PostTagsGetAllRes, error) {
	out := new(PostTagsGetAllRes)
	err := c.cc.Invoke(ctx, PostTagsService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) Update(ctx context.Context, in *PostTagsUpdate, opts ...grpc.CallOption) (*PostTags, error) {
	out := new(PostTags)
	err := c.cc.Invoke(ctx, PostTagsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PostTagsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) GetPostByTag(ctx context.Context, in *GetFilter, opts ...grpc.CallOption) (*PostTagsGetAllRes, error) {
	out := new(PostTagsGetAllRes)
	err := c.cc.Invoke(ctx, PostTagsService_GetPostByTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postTagsServiceClient) GetPopularTag(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PopularTag, error) {
	out := new(PopularTag)
	err := c.cc.Invoke(ctx, PostTagsService_GetPopularTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostTagsServiceServer is the server API for PostTagsService service.
// All implementations must embed UnimplementedPostTagsServiceServer
// for forward compatibility
type PostTagsServiceServer interface {
	Create(context.Context, *PostTagsCreateReq) (*PostTags, error)
	GetById(context.Context, *GetByIdReq) (*PostTagsRes, error)
	GetAll(context.Context, *Filter) (*PostTagsGetAllRes, error)
	Update(context.Context, *PostTagsUpdate) (*PostTags, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	GetPostByTag(context.Context, *GetFilter) (*PostTagsGetAllRes, error)
	GetPopularTag(context.Context, *Void) (*PopularTag, error)
	mustEmbedUnimplementedPostTagsServiceServer()
}

// UnimplementedPostTagsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostTagsServiceServer struct {
}

func (UnimplementedPostTagsServiceServer) Create(context.Context, *PostTagsCreateReq) (*PostTags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPostTagsServiceServer) GetById(context.Context, *GetByIdReq) (*PostTagsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPostTagsServiceServer) GetAll(context.Context, *Filter) (*PostTagsGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPostTagsServiceServer) Update(context.Context, *PostTagsUpdate) (*PostTags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPostTagsServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPostTagsServiceServer) GetPostByTag(context.Context, *GetFilter) (*PostTagsGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostByTag not implemented")
}
func (UnimplementedPostTagsServiceServer) GetPopularTag(context.Context, *Void) (*PopularTag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularTag not implemented")
}
func (UnimplementedPostTagsServiceServer) mustEmbedUnimplementedPostTagsServiceServer() {}

// UnsafePostTagsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostTagsServiceServer will
// result in compilation errors.
type UnsafePostTagsServiceServer interface {
	mustEmbedUnimplementedPostTagsServiceServer()
}

func RegisterPostTagsServiceServer(s grpc.ServiceRegistrar, srv PostTagsServiceServer) {
	s.RegisterService(&PostTagsService_ServiceDesc, srv)
}

func _PostTagsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTagsCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).Create(ctx, req.(*PostTagsCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).GetById(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTagsUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).Update(ctx, req.(*PostTagsUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).Delete(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_GetPostByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).GetPostByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_GetPostByTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).GetPostByTag(ctx, req.(*GetFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostTagsService_GetPopularTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostTagsServiceServer).GetPopularTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostTagsService_GetPopularTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostTagsServiceServer).GetPopularTag(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// PostTagsService_ServiceDesc is the grpc.ServiceDesc for PostTagsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostTagsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forum.PostTagsService",
	HandlerType: (*PostTagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PostTagsService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _PostTagsService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PostTagsService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PostTagsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PostTagsService_Delete_Handler,
		},
		{
			MethodName: "GetPostByTag",
			Handler:    _PostTagsService_GetPostByTag_Handler,
		},
		{
			MethodName: "GetPopularTag",
			Handler:    _PostTagsService_GetPopularTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_tag.proto",
}
