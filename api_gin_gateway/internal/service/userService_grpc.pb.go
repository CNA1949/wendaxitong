// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: userService.proto

package service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ModifyUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserInfoByUserName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FollowUser(ctx context.Context, in *UserRequest2, opts ...grpc.CallOption) (*UserResponse, error)
	CreateTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	DeleteTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	CommentTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	DeleteComment(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	LikesTopicOrComment(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	QueryTopicInfo(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
	QueryCommentInfo(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ModifyUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/ModifyUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByUserName(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/GetUserInfoByUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FollowUser(ctx context.Context, in *UserRequest2, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CommentTopic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/CommentTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteComment(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LikesTopicOrComment(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/LikesTopicOrComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryTopicInfo(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/QueryTopicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryCommentInfo(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (*TopicResponse, error) {
	out := new(TopicResponse)
	err := c.cc.Invoke(ctx, "/service.UserService/QueryCommentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserLogin(context.Context, *UserRequest) (*UserResponse, error)
	UserRegister(context.Context, *UserRequest) (*UserResponse, error)
	DeleteUser(context.Context, *UserRequest) (*UserResponse, error)
	ModifyUserInfo(context.Context, *UserRequest) (*UserResponse, error)
	GetUserInfoByUserName(context.Context, *UserRequest) (*UserResponse, error)
	FollowUser(context.Context, *UserRequest2) (*UserResponse, error)
	CreateTopic(context.Context, *TopicRequest) (*TopicResponse, error)
	DeleteTopic(context.Context, *TopicRequest) (*TopicResponse, error)
	CommentTopic(context.Context, *TopicRequest) (*TopicResponse, error)
	DeleteComment(context.Context, *TopicRequest) (*TopicResponse, error)
	LikesTopicOrComment(context.Context, *TopicRequest) (*TopicResponse, error)
	QueryTopicInfo(context.Context, *TopicRequest) (*TopicResponse, error)
	QueryCommentInfo(context.Context, *TopicRequest) (*TopicResponse, error)
	MustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) ModifyUserInfo(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoByUserName(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUserName not implemented")
}
func (UnimplementedUserServiceServer) FollowUser(context.Context, *UserRequest2) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedUserServiceServer) CreateTopic(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedUserServiceServer) DeleteTopic(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedUserServiceServer) CommentTopic(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentTopic not implemented")
}
func (UnimplementedUserServiceServer) DeleteComment(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedUserServiceServer) LikesTopicOrComment(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikesTopicOrComment not implemented")
}
func (UnimplementedUserServiceServer) QueryTopicInfo(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTopicInfo not implemented")
}
func (UnimplementedUserServiceServer) QueryCommentInfo(context.Context, *TopicRequest) (*TopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCommentInfo not implemented")
}
func (UnimplementedUserServiceServer) MustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	MustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/ModifyUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ModifyUserInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/GetUserInfoByUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByUserName(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FollowUser(ctx, req.(*UserRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CommentTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CommentTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/CommentTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CommentTopic(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteComment(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LikesTopicOrComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LikesTopicOrComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/LikesTopicOrComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LikesTopicOrComment(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryTopicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryTopicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/QueryTopicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryTopicInfo(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryCommentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryCommentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UserService/QueryCommentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryCommentInfo(ctx, req.(*TopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "ModifyUserInfo",
			Handler:    _UserService_ModifyUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfoByUserName",
			Handler:    _UserService_GetUserInfoByUserName_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _UserService_FollowUser_Handler,
		},
		{
			MethodName: "CreateTopic",
			Handler:    _UserService_CreateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _UserService_DeleteTopic_Handler,
		},
		{
			MethodName: "CommentTopic",
			Handler:    _UserService_CommentTopic_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _UserService_DeleteComment_Handler,
		},
		{
			MethodName: "LikesTopicOrComment",
			Handler:    _UserService_LikesTopicOrComment_Handler,
		},
		{
			MethodName: "QueryTopicInfo",
			Handler:    _UserService_QueryTopicInfo_Handler,
		},
		{
			MethodName: "QueryCommentInfo",
			Handler:    _UserService_QueryCommentInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userService.proto",
}
