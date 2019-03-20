package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"video_server/streamserver/defs"
	"video_server/streamserver/response"
)



/**
视频播放
 */
func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := defs.VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

/**
视频上传
 */
func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 上传大小设置
	r.Body = http.MaxBytesReader(w, r.Body, defs.MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	// 读出file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	// 写入
	err = ioutil.WriteFile(defs.VIDEO_DIR + fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
