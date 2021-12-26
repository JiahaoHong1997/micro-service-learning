package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micro-service-learning/projectLayout/api"
	"micro-service-learning/projectLayout/internal/biz"
)

type UserService struct {
	api.UnimplementedUserServiceServer
	biz *biz.UserBiz
}

func NewUserService(biz *biz.UserBiz) *UserService {
	return &UserService{
		biz: biz,
	}
}

func (u *UserService) UserInfo(ctx context.Context, request *api.UserInfoRequest) (*api.UserInfoReply, error) {
	usr, err := u.biz.GetUserById(request.Uid)
	if err != nil {
		return nil, status.Error(codes.Code(1), "system error")
	}

	return &api.UserInfoReply{
		User: &api.User{
			Nickname: usr.Nickname,
		},
	}, nil
}
