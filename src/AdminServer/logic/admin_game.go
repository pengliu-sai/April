package logic

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"tools/http_api"
)

func clearGameServerInfo(gameAddr string) {
	_lock.Lock()
	defer _lock.Unlock()

	gameServersRing.RemoveNode(gameAddr)

	var userIDMap = make(map[uint64]bool)
	for k, v := range g_userID_GameServerAddr {
		if v == gameAddr {
			//先收集, 后删除
			userIDMap[k] = true
		}
	}
	for m := range userIDMap {
		delete(g_userID_GameServerAddr, m)
	}
}

//用户下线
func (s *httpServer) doGame_UserOffline(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	userIDStr, err := reqParams.Get("userid")
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_USERID"}
	}

	gameAddr := g_userID_GameServerAddr[userID]
	if gameAddr == "" {
		return nil, http_api.Err{400, "USER_NOT_ONLINE"}
	}

	//清除登陆信息
	clearUserLoginInfo(userID)

	return "SUCCESS", nil
}

//游戏服务器下线
func (s *httpServer) doGame_GameServerOffline(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	gameAddr, err := reqParams.Get("gameserveraddr")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_GAMEADDR"}
	}

	//清理gameServer信息
	clearGameServerInfo(gameAddr)

	return "SUCCESS", nil
}
