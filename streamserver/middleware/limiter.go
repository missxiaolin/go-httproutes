package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
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


func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		response.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}


	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}
