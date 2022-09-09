package handler

import (
	"context"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/user/internal/repository"
	"wendaxitong/user/internal/service"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) MustEmbedUnimplementedUserServiceServer() {}

func (u *UserService) UserLogin(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.UserLogin(request)
	return &service.UserResponse{
		UserInfo: &service.UserModel{
			UserName: user.UserName,
			Password: "******",
			Phone:    user.Phone,
			Email:    user.Email,
		},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage}, nil
}

func (u *UserService) UserRegister(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.RegisterUserInfo(request)
	return &service.UserResponse{
		UserInfo: &service.UserModel{
			UserName: user.UserName,
			Password: "******",
			Phone:    user.Phone,
			Email:    user.Email,
		},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage}, nil
}

func (u *UserService) DeleteUser(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.DeleteUser(request)
	return &service.UserResponse{
		UserInfo:      nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) ModifyUserInfo(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.ModifyUserInfo(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.UserResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.UserResponse{
		UserInfo: &service.UserModel{
			UserId:   user.UserId,
			UserName: user.UserName,
			Password: "******",
			Phone:    user.Phone,
			Email:    user.Email,
		},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) GetUserIdByUserName(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.GetUserInfoByUserName(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.UserResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}

	return &service.UserResponse{
		UserInfo:      &service.UserModel{UserId: user.UserId},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) GetUserInfoByUserName(c context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.GetUserInfoByUserName(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.UserResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}

	return &service.UserResponse{
		UserInfo: &service.UserModel{
			UserId:     user.UserId,
			UserName:   user.UserName,
			Password:   user.Password,
			Phone:      user.Phone,
			Email:      user.Email,
			NumFans:    user.NumFans,
			NumIdols:   user.NumIdols,
			FansNames:  user.FansNames,
			IdolsNames: user.IdolsNames,
		},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) FollowUser(ctx context.Context, request *service.UserRequest2) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.FollowUser(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.UserResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}

	return &service.UserResponse{
		UserInfo:      &service.UserModel{},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}
