package main

import (
	"micro-service-learning/projectLayout/internal/biz"
	"micro-service-learning/projectLayout/internal/data"
	"micro-service-learning/projectLayout/internal/service"
)

func InitUserServiceWithoutWire() *service.UserService {
	db := data.NewDB()
	component := data.NewCache()
	userRepo := data.NewUserRepo(db, component)
	userBiz := biz.NewUserBiz(userRepo)
	userService := service.NewUserService(userBiz)
	return userService
}
