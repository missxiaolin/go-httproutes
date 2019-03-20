package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
	"video_server/streamserver/defs"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := defs.VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		//response.SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	
}
