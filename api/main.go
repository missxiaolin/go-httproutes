package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/middleware"
	"video_server/api/routes"
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

func main()  {
	r := routes.RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8080", mh)
}