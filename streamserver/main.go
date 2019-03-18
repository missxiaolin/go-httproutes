package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/streamserver/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handlers.StreamHandler)

	router.POST("/upload/:vid-id", handlers.UploadHandler)

	return router;
}

func main()  {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}