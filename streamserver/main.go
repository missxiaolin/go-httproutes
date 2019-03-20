package main

import (
	"net/http"
	"video_server/streamserver/middleware"
	"video_server/streamserver/routes"
)

func main()  {
	r := routes.RegisterHandlers()
	mh := middleware.NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}