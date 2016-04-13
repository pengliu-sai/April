package logic

import (
	"AdminServer/model"
	"encoding/json"
	"fmt"
	"github.com/funny/link"
	"io/ioutil"
	"net/http"
	"protos"
	"protos/admin"
	"strconv"
	. "tools"
)

//登陆时不分派服务器, 在选择区时,在分配服务器.
//玩家只登陆一次, 登陆后获取sessionID 和 token 授权...
//在授权的有效期内,  玩家可以登陆各种区...

//登陆后选择服务器...
func userLogin(gateSession *link.Session, msg protos.ProtoMsg) {
	INFO("Gate receive userLoginC2S message")
	rev_msg := msg.Body.(*admin.Admin_UserLoginC2S)
	name := *rev_msg.Name
	password := *rev_msg.Password

	null_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginS2C{})

	if name == "" || password == "" {
		//如果帐号和密码为空, 发送空消息回去
		INFO("name or password is null")
		gateSession.Send(null_msg)
		return
	}

	//发动到adminServer进行登陆
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userLoginIn?name=%s&password=%s", adminAddr, name, password)
	req, _ := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		ERR("userLoginIn error: ", err.Error())
		gateSession.Send(null_msg)
		return
	}

	//解析登陆结果, 取出userID, gameAddr
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var user model.User

	if err := json.Unmarshal(body, &user); err != nil {
		ERR("userLogin json Unmarshal err")
		gateSession.Send(null_msg)
		return
	}

	//用户登陆完成, 并且连接上世界服务器后, 并向gameServer注册UserID之后, 才算正正的登陆成功.
	//连接游戏服务器
	INFO("try connect GameServer :", user.ID)
	gameClientSession, err := connectGameServer(user.ID, user.GameServerAddr, gateSession)

	if gameClientSession == nil || err != nil {
		ERR("connectGameServer error")
		gateSession.Send(null_msg)
		return
	}
	INFO("connect GameServer success")
	INFO("Gate send registerUserID to GameServer: ", user.ID)
	//向gameServer注册userID
	sendMsgToGameServer_RegisterUserID(gateSession, user.ID)
}

func userRegister(gateSession *link.Session, msg protos.ProtoMsg) {
	INFO("Gate receive userRegisterC2S message")
	rev_msg := msg.Body.(*admin.Admin_UserRegisterC2S)
	name := *rev_msg.Name
	password := *rev_msg.Password

	null_msg := protos.MarshalProtoMsg(&admin.Admin_UserRegisterS2C{})

	if name == "" || password == "" {
		//如果帐号和密码为空, 发送空消息回去
		INFO("name or password is null")
		gateSession.Send(null_msg)
		return
	}

	//发到adminServer进行注册
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userRegister?name=%s&password=%s", adminAddr, name, password)
	req, _ := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		ERR("userRegister error: ", name)
		gateSession.Send(null_msg)
		return
	}

	//解析登陆结果, 取出userID, Name
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var user model.User

	if err := json.Unmarshal(body, &user); err != nil {
		ERR("userRegister json Unmarshal err")
		gateSession.Send(null_msg)
		return
	}

	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserRegisterS2C{
		UserID: protos.Uint64(user.ID),
		Name:   protos.String(user.Name),
	})

	INFO("Gate send userRegisterS2C message")
	gateSession.Send(send_msg)
}

func userExit(gateSession *link.Session, msg protos.ProtoMsg) {
	INFO("Gate receive userExitC2S message")
	rev_msg := msg.Body.(*admin.Admin_UserExitC2S)
	optional_userID := *rev_msg.Optional_UserID

	falseResult := false

	false_msg := protos.MarshalProtoMsg(&admin.Admin_UserExitS2C{
		Result: &falseResult,
	})

	userID := g_gateSession_UserID[gateSession]

	//开发模式
	if developMode {
		if optional_userID <= 0 {
			gateSession.Send(false_msg)
			return
		}
		userID = optional_userID
	}

	//发到adminServer进行退出
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userExit?userid=%s", adminAddr, strconv.FormatUint(userID, 10))
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		gateSession.Send(false_msg)
		return
	}

	//解析登陆结果, 取出userID, Name
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	result := false

	if string(body) == "SUCCESS" {
		result = true
	}

	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserExitS2C{
		Result: &result,
	})

	INFO("Gate send userExitS2C message")
	gateSession.Send(send_msg)
}
