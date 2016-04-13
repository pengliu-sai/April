package logic

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	. "tools"
	"tools/http_api"
)

func notifyUserOffline(userID uint64) {
	_lock.Lock()
	defer _lock.Unlock()
	if gameAddr, ok := g_userID_GameServerAddr[userID]; ok {
		url := fmt.Sprintf("http://%s/admin/userOffline?userid=%s", gameAddr, userID)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			ERR("User Offline failure: ", userID)
		}
	}
}

//跨区登陆, 通知其它区下线该用户
func (s *httpServer) doOtherAdmin_UserOffline(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	userIDStr, err := reqParams.Get("userid")
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_USERID"}
	}

	notifyUserOffline(userID)

	//清除登陆信息
	clearUserLoginInfo(userID)

	return "SUCCESS", nil
}
