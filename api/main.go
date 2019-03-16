package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/handlers"
	"video_server/api/middleware"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	middleware.ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	//创建用户
	router.POST("/user", handlers.CreateUser)

	// 用户登录
	router.GET("/user/:userName", handlers.Login)

	// 用户资源
	//router.GET("user/:username/videos")

	//router.GET("user/:userName/videos/:vidId")

	return router
}

func main()  {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8080", mh)
}