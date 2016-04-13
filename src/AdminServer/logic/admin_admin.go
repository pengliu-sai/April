package logic

import (
	"AdminServer/dao"
	"github.com/julienschmidt/httprouter"
	"net/http"
	. "tools"
	"tools/http_api"
)

// admin管理员权限......

//删除玩家
func (s *httpServer) doAdmin_UserDelete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	name, err := reqParams.Get("name")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_NAME"}
	}

	password, err := reqParams.Get("password")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_PASSWORD"}
	}

	num, err := dao.UserDelete(name, password)
	if err != nil {
		return nil, http_api.Err{400, "CANOT_DELETE_USER"}
	}
	return num, nil
}

//标记删除
func (s *httpServer) doAdmin_UserMarkDeleted(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	userID, err := reqParams.Get("userid")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_USERID"}

	}

	num, err := dao.UserMarkDeleted(userID)
	if err != nil {
		return nil, http_api.Err{400, "CANOT_MARK_DELETED_USER"}
	}
	return num, nil

}

//恢复用户
func (s *httpServer) doAdmin_UserRebirth(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	userID, err := reqParams.Get("userid")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_USERID"}

	}

	num, err := dao.UserRebirth(userID)

	if err != nil {
		return nil, http_api.Err{400, "CANOT_REBIRTH_USER"}
	}
	return num, nil
}

//获取用户信息
func (s *httpServer) doAdmin_UserInfo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	name, err := reqParams.Get("name")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_NAME"}
	}

	password, err := reqParams.Get("password")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_PASSWORD"}
	}

	user, err := dao.UserInfo(name, password)
	if user == nil || err != nil {
		return nil, http_api.Err{400, "CANOT_SEARCH_USER"}
	}

	INFO("user search success:", user.Name)

	return user, nil
}

//游戏服务器下线
func (s *httpServer) doAdmin_GameServerOffline(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	gameAddr, err := reqParams.Get("gameaddr")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_GAMEADDR"}
	}

	//清理gameServer信息
	clearGameServerInfo(gameAddr)

	return "SUCCESS", nil
}
