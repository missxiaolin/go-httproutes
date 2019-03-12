package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/handlers"
)

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	//创建用户
	router.GET("/user", handlers.CreateUser)

	// 用户登录
	router.GET("/user/:userName", handlers.Login)

	// 用户资源
	//router.GET("user/:username/videos")

	//router.GET("user/:userName/videos/:vidId")

	return router
}

func main()  {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}