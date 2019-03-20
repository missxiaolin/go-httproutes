package routes

import (
	"github.com/julienschmidt/httprouter"
	"video_server/streamserver/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handlers.StreamHandler)

	router.POST("/upload/:vid-id", handlers.UploadHandler)

	return router;
}
