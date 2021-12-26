package biz

import (
	"errors"
	"micro-service-learning/projectLayout/internal/data"
)

type UserBiz struct {
	repo *data.UserRepo
}

func NewUserBiz(repo *data.UserRepo) *UserBiz {
	return &UserBiz{
		repo: repo,
	}
}

func (ub *UserBiz) GetUserById(uid uint64) (*UserDO, error) {
	if uid == 0 {
		return nil, errors.New("invalid user id")
	}
	u, err := ub.repo.GetUser(uid)
	if err != nil {
		return nil, err
	}
	return &UserDO{
		Nickname: u.Nickname,
	}, nil
}

type UserDO struct {
	Nickname string
}

type MyService interface {
}

type Option func(db *service) MyService

func NewDB(opts ...Option) MyService {
	return &service{}
}

type service struct {
}
