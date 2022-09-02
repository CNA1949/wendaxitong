package handler

import (
	"context"
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
