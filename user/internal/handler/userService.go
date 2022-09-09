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

func (u *UserService) UserLogin(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
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

func (u *UserService) UserRegister(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
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

func (u *UserService) DeleteUser(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	msg := user.DeleteUser(request)
	return &service.UserResponse{
		UserInfo:      nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) ModifyUserInfo(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user *repository.UserInfo
	userInfo, msg := user.ModifyUserInfo(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.UserResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.UserResponse{
		UserInfo: &service.UserModel{
			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,
			Password: "******",
			Phone:    userInfo.Phone,
			Email:    userInfo.Email,
		},
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) GetUserInfoByUserName(ctx context.Context, request *service.UserRequest) (*service.UserResponse, error) {
	var user repository.UserInfo
	user, msg := user.GetUserInfoByUserName(request)
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
			NumTopic:   user.NumTopic,
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

func (u *UserService) CreateTopic(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	msg := topic.CreateTopic(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.TopicResponse{
		TopicInfo:     nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) DeleteTopic(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	msg := topic.DeleteTopic(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.TopicResponse{
		TopicInfo:     nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) CommentTopic(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	msg := topic.CommentTopic(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.TopicResponse{
		TopicInfo:     nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) DeleteComment(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	msg := topic.DeleteComment(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.TopicResponse{
		TopicInfo:     nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) LikesTopicOrComment(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	msg := topic.LikesTopicOrComment(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}
	return &service.TopicResponse{
		TopicInfo:     nil,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) QueryTopicInfo(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	topics, msg := topic.QueryTopicInfo(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}

	var topicInfo = make([]*service.TopicModel, 0, len(topics))
	for i := 0; i < len(topics); i++ {
		topicInfo = append(topicInfo, &service.TopicModel{
			Id:       topics[i].Id,
			UserName: topics[i].UserName,
			Content:  topics[i].Content,
			NumLikes: topics[i].NumLikes,
		})
	}

	return &service.TopicResponse{
		TopicInfo:     topicInfo,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}

func (u *UserService) QueryCommentInfo(ctx context.Context, request *service.TopicRequest) (*service.TopicResponse, error) {
	var topic repository.TopicInfo
	topics, msg := topic.QueryCommentInfo(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return &service.TopicResponse{
			StatusCode:    msg.StatusCode,
			StatusMessage: msg.StatusMessage,
		}, nil
	}

	var topicInfo = make([]*service.TopicModel, 0, len(topics))
	for i := 0; i < len(topics); i++ {
		topicInfo = append(topicInfo, &service.TopicModel{
			Id:       topics[i].Id,
			UserName: topics[i].UserName,
			Content:  topics[i].Content,
			NumLikes: topics[i].NumLikes,
			RootId:   topics[i].RootId,
			ParentId: topics[i].ParentId,
		})
	}

	return &service.TopicResponse{
		TopicInfo:     topicInfo,
		StatusCode:    msg.StatusCode,
		StatusMessage: msg.StatusMessage,
	}, nil
}
