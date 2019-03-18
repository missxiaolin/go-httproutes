package routes

import (
	"github.com/julienschmidt/httprouter"
	"video_server/api/handlers"
)

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()
	//创建用户
	router.POST("/user", handlers.CreateUser)

	// 用户登录
	router.GET("/user/:userName", handlers.Login)

	return router
}