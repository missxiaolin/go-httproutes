package handles

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/dbops"
	"video_server/scheduler/util"
)

func VidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		util.SendResponse(w, 400, "video id should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		util.SendResponse(w, 500, "Internal server error")
		return
	}

	util.SendResponse(w, 200, "success")
	return
}
