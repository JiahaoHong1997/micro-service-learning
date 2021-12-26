//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"micro-service-learning/projectLayout/internal/biz"
	"micro-service-learning/projectLayout/internal/data"
	"micro-service-learning/projectLayout/internal/service"
)

func InitUserService() *service.UserService {
	wire.Build(service.NewUserService, biz.NewUserBiz, data.NewUserRepo, data.NewDB, data.NewCache)
	return &service.UserService{}
}
