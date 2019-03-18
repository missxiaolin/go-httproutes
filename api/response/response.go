package response

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	w.WriteHeader(sc)
	io.WriteString(w, resp)
}