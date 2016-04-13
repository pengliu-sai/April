package logic

import (
	"AdminServer/dao"
	"AdminServer/model"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	. "tools"
	"tools/cfg"
	"tools/http_api"
	. "tools/redis"
)

const (
	DB_UserPwd_Key    = "DB_UserPwd_"
	DB_UserAreaID_Key = "DB_UserAreaID_"
)

//黑名单检查
func checkInBlackUserList(name string) bool {
	blackUserConf := cfg.GetUserConfig().BlackUserList
	exist := false
	blackUserLen := len(blackUserConf)
	for i := 0; i < blackUserLen; i++ {
		if blackUserConf[i].Name == name {
			exist = true
			break
		}
	}
	return exist
}

func clearUserLoginInfo(userID uint64) {
	_lock.Lock()
	defer _lock.Unlock()
	delete(g_userID_GameServerAddr, userID)

}

//跨区登陆检查
func checkCrossAreaLogin(userID uint64, areaID uint16) {
	if areaID > 0 && cfg.GetServerConfig().Area.ID != areaID {
		//从数据库中取出areaID 对应的addr.
		admin, err := dao.AreaInfoByAreaID(areaID)
		if err != nil {
			INFO("can't search area by areaID: ", areaID)
			return
		}

		if admin != nil {
			//向其它admin通知用户下线.
			client := http.Client{}
			url := fmt.Sprintf("http://%s:%s//admin/userExit?userid=%s", admin.IP, admin.Port, userID)
			req, _ := http.NewRequest("GET", url, nil)
			resp, err := client.Do(req)
			if resp.StatusCode != 200 {
				ERR("admin user exit  failure: ", areaID, ":", err)
			}
		}
	}
}

//分配游戏服务器
func allocGameServer(userID uint64) (string, bool) {
	_lock.Lock()
	defer _lock.Unlock()
	//判断是否已经分配过
	if v, ok := g_userID_GameServerAddr[userID]; ok {
		return v, true
	}

	//重新分配
	serverAddr, ok := gameServersRing.GetNode(string(userID))
	if ok {
		//加入缓存
		g_userID_GameServerAddr[userID] = serverAddr
		return serverAddr, true
	}

	return "", false
}

//用户登陆
func (s *httpServer) doGate_UserLoginByUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	INFO("Admin receive userLoginByUserNameC2S message")
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

	//黑名单
	exist := checkInBlackUserList(name)
	if exist {
		return nil, http_api.Err{400, "USER_IN_BLACK_LIST"}
	}

	user, err := dao.UserLoginByName(name, password)
	if user == nil {
		return nil, http_api.Err{400, "CANOT_FIND_USER"}
	}

	//跨区登陆检查
	checkCrossAreaLogin(user.ID, user.LastLoginAreaID)

	//分配游戏服务器
	gameServerAddr, ok := allocGameServer(user.ID)

	if !ok {
		return nil, http_api.Err{400, "ALLOC_GAMESERVER_FAILURE"}
	}

	user.GameServerAddr = gameServerAddr

	//密码存入redis
	userPwdKey := DB_UserPwd_Key + strconv.FormatUint(user.ID, 10)

	v, err := RedisClient.Do("SET", userPwdKey, password)
	if err != nil {
		return nil, http_api.Err{400, "REDIS_SET_USERPWD_FAILURE"}
	}

	if v != "OK" {
		return nil, http_api.Err{400, "REDIS_SET_USERPWD_FAILURE"}
	}

	//设置密码过期时间
	RedisClient.Do("EXPIRE", userPwdKey, 30)

	//areaID存入redis
	userAreaIDKey := DB_UserAreaID_Key + strconv.FormatUint(user.ID, 10)
	v, err = RedisClient.Do("SET", userAreaIDKey, cfg.GetServerConfig().Area.ID)
	if err != nil {
		return nil, http_api.Err{400, "REDIS_SET_AREAID_FAILURE"}
	}

	if v != "OK" {
		return nil, http_api.Err{400, "REDIS_SET_AREAID_FAILURE"}
	}

	INFO("Admin send userLoginByUserNameS2C message")
	return user, nil
}

func (s *httpServer) doGate_UserLoginByUserID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	INFO("Admin receive userLoginByUserIDC2S message")
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	name, err := reqParams.Get("name")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_NAME"}
	}

	userIDStr, err := reqParams.Get("userid")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_USERID"}
	}

	password, err := reqParams.Get("password")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_PASSWORD"}
	}

	//黑名单
	exist := checkInBlackUserList(name)
	if exist {
		return nil, http_api.Err{400, "USER_IN_BLACK_LIST"}
	}

	//从redis取出密码
	userPwdKey := DB_UserPwd_Key + userIDStr
	userPwdValue, err := redis.String(RedisClient.Do("GET", userPwdKey))
	if err != nil {
		return nil, http_api.Err{400, "REDIS_GET_USERPWD_FAILURE"}
	}

	if userPwdValue == "" {
		return nil, http_api.Err{400, "USER_PASSWORD_EXPIRED"}
	}

	if userPwdValue != password {
		return nil, http_api.Err{400, "USER_PASSWORD_INCORRECT"}
	}

	//验证通过..

	//从redis中取出之前登陆的areaID
	userAreaIDKey := DB_UserAreaID_Key + userIDStr
	userAreaIDValue, err := redis.String(RedisClient.Do("GET", userAreaIDKey))
	if err != nil {
		return nil, http_api.Err{400, "REDIS_GET_AREAID_FAILURE"}
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, http_api.Err{400, "PARSE_UINT_USERID_FAILURE"}
	}

	lastLoginAreaID := uint16(0)

	//如果之前有登陆其它区
	if userAreaIDValue != "" {
		_areaID, err := strconv.ParseUint(userAreaIDValue, 10, 16)
		if err != nil {
			return nil, http_api.Err{400, "PARSE_UINT_AREAID_FAILURE"}
		}
		lastLoginAreaID = uint16(_areaID)
		//跨区检查
		checkCrossAreaLogin(userID, lastLoginAreaID)
	}

	//分配游戏服务器
	gameServerAddr, ok := allocGameServer(userID)

	if !ok {
		return nil, http_api.Err{400, "ALLOC_GAMESERVER_FAILURE"}
	}

	user := model.NewUserModel()
	user.ID = userID
	user.GameServerAddr = gameServerAddr

	INFO("Admin send userLoginByUserIDS2C message")
	return user, nil
}

//用户注册
func (s *httpServer) doGate_UserRegister(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	INFO("Admin receive userRegisterC2S message")
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

	user, err := dao.UserRegister(name, password)
	if user == nil || err != nil {
		return nil, http_api.Err{400, "CANOT_REGISTER_USER"}
	}

	INFO("Admin send userRegisterS2C message")
	return user, nil
}

//用户主动退出
func (s *httpServer) doGate_UserExit(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	INFO("Admin receive userRegisterC2S message")
	INFO("Admin receive userExitC2S message")

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

	INFO("Admin send userExitS2C message")
	return "SUCCESS", nil
}
