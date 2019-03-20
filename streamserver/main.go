package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/streamserver/handlers"
	"video_server/streamserver/response"
	"video_server/streamserver/util"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *util.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = util.NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handlers.StreamHandler)

	router.POST("/upload/:vid-id", handlers.UploadHandler)

	return router;
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		response.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main()  {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}