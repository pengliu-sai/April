package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/admin"
	"protos/game"
	"protos/gate"
	. "protos/msgid"
	"testing"
	. "tools"
	"tools/cfg"
	"tools/dispatch"
	"tools/protocol"
	. "tools/unitest"
)

var (
	clientMsgDispatch dispatch.DispatchInterface
	gateAddr          string
	clientSession     *link.Session

	global_t *testing.T

	clientSession_liupeng *link.Session
	clientSession_god     *link.Session
)

/*

	测试数据.
	用户: 刘鹏, God
	角色: 杰伦是刘鹏的, 叮当是上帝的, 喵
*/

var (
	global_userID_liupeng        = uint64(915441286658396160)
	global_userID_god            = uint64(915664487518834688)
	global_liupeng_roleID_jielun = uint64(915771140146728960)
	global_god_roleID_dingdang   = uint64(915787632821145600)
)

func init() {
	game.InitGameMsgID()

	gateAddr = cfg.GetServerConfig().GateServer.IP + ":" + cfg.GetServerConfig().GateServer.Port
	handle := dispatch.NewHandleConditions()

	/* gate -> client  */
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidGateID,
		H: dispatch.Handle{
			uint16(MsgID_Gate_PingS2C): _gatePingCallBack,
		},
	})

	/* admin -> client */
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidAdminID,
		H: dispatch.Handle{
			uint16(MsgID_Admin_UserLoginS2C):    _userLoginCallBack,
			uint16(MsgID_Admin_UserRegisterS2C): _userRegisterCallBack,
			uint16(MsgID_Admin_UserExitS2C):     _userExitCallBack,
		},
	})

	/* game -> client */
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidGameID,
		H: dispatch.Handle{
			uint16(MsgID_Game_RandomRoleNameS2C):   _randomRoleNameCallBack,
			uint16(MsgID_Game_RoleCreateS2C):       _roleCreateCallBack,
			uint16(MsgID_Game_RoleInfoListS2C):     _roleInfoListCallBack,
			uint16(MsgID_Game_EnterScenesS2C):      _enterScenesCallBack,
			uint16(MsgID_Game_SendChatS2C):         _sendChatCallBack,
			uint16(MsgID_Game_Receive_ChatS2C):     _receiveChatCallBack,
			uint16(MsgID_Game_RoleInfoByRoleIDS2C): _roleInfoByRoleIDCallBack,
			uint16(MsgID_Game_FriendListS2C):       _friendListCallBack,
			uint16(MsgID_Game_AddFriendS2C):        _addFriendCallBack,
		},
	})

	clientMsgDispatch = dispatch.NewDispatch(handle)
}

func connectGateServer() {
	var err error
	clientSession, err = link.Connect("tcp", gateAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectGateServer_liupeng() {
	var err error
	clientSession_liupeng, err = link.Connect("tcp", gateAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_liupeng, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectGateServer_god() {
	var err error
	clientSession_god, err = link.Connect("tcp", gateAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_god, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func TestGateConnect(t *testing.T) {
	connectGateServer()

	global_t = t

	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&gate.Gate_PingC2S{
		Content: protos.String("Hi GateServer"),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _gatePingCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	rev_msg := msg.Body.(*gate.Gate_PingS2C)
	Equal(global_t, *(rev_msg.Content), "Hi GateClient")
}

//func TestInitData_liupeng(t *testing.T) {
//	//连接gate
//	connectGateServer_liupeng()
//
//	//注册liupeng
//	registerUser_liupeng()
//
//	//登陆
//	loginUser_liupeng()
//
//	//创建角色
//	roleCreate_liupeng_jielun()
//}

//func TestInitData_god(t *testing.T) {
//	connectGateServer_god()
//
//	registerUser_god()
//
//	loginUser_god()
//
//	roleCreate_god_dingdang()
//}

func registerUser_liupeng() {
	waitGroup.Add(1)
	DEBUG("send userRegisterC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserRegisterC2S{
		Name:     protos.String("liupeng"),
		Password: protos.String("123456"),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()

}

func registerUser_god() {
	waitGroup.Add(1)
	DEBUG("send userRegisterC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserRegisterC2S{
		Name:     protos.String("god"),
		Password: protos.String("654321"),
	})
	clientSession_god.Send(send_msg)
	waitGroup.Wait()
}

func loginUser_liupeng() {
	waitGroup.Add(1)
	DEBUG("send userLoginC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginC2S{
		Name:     protos.String("liupeng"),
		Password: protos.String("123456"),
	})

	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func loginUser_god() {
	waitGroup.Add(1)
	DEBUG("send userLoginC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginC2S{
		Name:     protos.String("god"),
		Password: protos.String("654321"),
	})

	clientSession_god.Send(send_msg)
	waitGroup.Wait()
}

func roleCreate_liupeng_jielun() {
	waitGroup.Add(1)
	DEBUG("send roleCreateC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateC2S{
		Name: protos.String("杰伦"),
		Sex:  game.SexEnum_male.Enum(),
		Race: game.RacesEnum_archer.Enum(),
	})
	clientSession_liupeng.Send(send_msg)

	waitGroup.Wait()
}

func roleCreate_god_dingdang() {
	waitGroup.Add(1)
	DEBUG("send roleCreateC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateC2S{
		Name: protos.String("叮当"),
		Sex:  game.SexEnum_male.Enum(),
		Race: game.RacesEnum_archer.Enum(),
	})
	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

func roleInfoList_liupeng() {
	waitGroup.Add(1)
	DEBUG("send roleInfoListC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListC2S{})
	clientSession_liupeng.Send(send_msg)

	waitGroup.Wait()
}

func roleInfoList_god() {
	waitGroup.Add(1)
	DEBUG("send roleInfoListC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListC2S{})
	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

func enterScenes_liupeng_jielun() {
	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession_liupeng.Send(send_msg)

	waitGroup.Wait()
}

func enterScenes_god_dingdang() {
	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_god_roleID_dingdang),
	})
	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

/*-------------------------[[gate -> admin]] --------------------*/
//测试用户登陆
func TestAdminUserLogin(t *testing.T) {
	connectGateServer()

	global_t = t
	waitGroup.Add(1)
	DEBUG("send userLoginC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginC2S{
		Name:     protos.String("liupeng"),
		Password: protos.String("123456"),
	})

	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _userLoginCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive userLoginS2C message")
	rev_msg := msg.Body.(*admin.Admin_UserLoginS2C)
	err_msg := "userLogin failure"
	if rev_msg != nil && rev_msg.UserID != nil {
		Assert(global_t, *rev_msg.UserID > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func TestAdminUserExit(t *testing.T) {
	connectGateServer()

	global_t = t

	waitGroup.Add(1)
	DEBUG("send userExitC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserExitC2S{
		Optional_UserID: protos.Uint64(global_userID_liupeng),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _userExitCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive userExitC2S message")
	rev_msg := msg.Body.(*admin.Admin_UserExitS2C)
	err_msg := "userExit error"
	if rev_msg != nil && rev_msg.Result != nil {
		Equal(global_t, *rev_msg.Result, true)
		Assert(global_t, *rev_msg.Result, err_msg)
	}
	DEBUG(err_msg)
}

//测试用户注册
func TestAdminUserRegister(t *testing.T) {
	connectGateServer()

	global_t = t

	waitGroup.Add(1)
	DEBUG("send userRegisterC2S message")
	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserRegisterC2S{
		Name:     protos.String("miaoming"),
		Password: protos.String("123456"),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _userRegisterCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive userRegisterS2C message")
	rev_msg := msg.Body.(*admin.Admin_UserRegisterS2C)

	err_msg := "register user failure"

	if rev_msg != nil && rev_msg.UserID != nil && rev_msg.Name != nil {
		Assert(global_t, *rev_msg.UserID > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

/*------------------[[gate -> game]]-------------------*/
//随机角色名称
func TestGameRandomRoleName(t *testing.T) {
	//用户先登陆游戏
	TestAdminUserLogin(t)
	waitGroup.Add(1)
	DEBUG("send randomRoleC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RandomRoleNameC2S{
		Sex: game.SexEnum_male.Enum(),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _randomRoleNameCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive randomRoleS2C message")
	rev_msg := msg.Body.(*game.Game_RandomRoleNameS2C)
	err_msg := "random role error"
	if rev_msg != nil && rev_msg.Name != nil {
		Assert(global_t, *rev_msg.Name != "", err_msg)
	}
	DEBUG(err_msg)
}

//创建角色
func TestGameRoleCreate(t *testing.T) {
	//用户先登陆游戏
	TestAdminUserLogin(t)
	waitGroup.Add(1)
	DEBUG("send roleCreateC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateC2S{
		Name: protos.String("加菲"),
		Sex:  game.SexEnum_male.Enum(),
		Race: game.RacesEnum_archer.Enum(),
	})
	clientSession.Send(send_msg)

	waitGroup.Wait()
}

func _roleCreateCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive roleCreateS2C message")
	rev_msg := msg.Body.(*game.Game_RoleCreateS2C)
	err_msg := "roleCreate failure"
	if rev_msg != nil && rev_msg.RoleID != nil {
		Assert(global_t, *rev_msg.RoleID > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

//获取角色列表
func TestGameRoleInfoList(t *testing.T) {
	//用户先登陆游戏
	TestAdminUserLogin(t)
	waitGroup.Add(1)
	DEBUG("send roleInfoListC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListC2S{})
	clientSession.Send(send_msg)

	waitGroup.Wait()
}

func _roleInfoListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive roleInfoListS2C message")
	rev_msg := msg.Body.(*game.Game_RoleInfoListS2C)
	err_msg := "roleInfoList failure"
	if rev_msg != nil && rev_msg.Roles != nil {
		Assert(global_t, len(rev_msg.Roles) > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

//进入游戏世界
func TestGameEnterScenes(t *testing.T) {
	//获取角色列表
	TestGameRoleInfoList(t)

	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession.Send(send_msg)

	waitGroup.Wait()
}

func _enterScenesCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive enterScenesS2C message")
	rev_msg := msg.Body.(*game.Game_EnterScenesS2C)
	err_msg := "enter scenes error"
	if rev_msg != nil && rev_msg.Result != nil {
		Equal(global_t, *rev_msg.Result, true)
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}

//杰伦上线
func TestOnline_jielun(t *testing.T) {
	//连接gate
	connectGateServer_liupeng()
	loginUser_liupeng()
	roleInfoList_liupeng()
	enterScenes_liupeng_jielun()
}

//叮当上线
func TestOnline_dingdang(t *testing.T) {
	connectGateServer_god()
	loginUser_god()
	roleInfoList_god()
	enterScenes_god_dingdang()
}

//聊天
func TestGameSendChat_Interactive(t *testing.T) {
	//杰伦上线
	TestOnline_jielun(t)
	//叮当上线
	TestOnline_dingdang(t)

	waitGroup.Add(4)
	DEBUG("send sendChatC2S message")
	//叮当向杰伦say Hi!
	send_msg := protos.MarshalProtoMsg(&game.Game_SendChatC2S{
		RoleID:  protos.Uint64(global_liupeng_roleID_jielun),
		Content: protos.String("你好, 杰伦!"),
	})

	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

func _sendChatCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive sendChatS2C message")
	rev_msg := msg.Body.(*game.Game_SendChatS2C)
	err_msg := "send chat error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}

}

func _receiveChatCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive receiveChatS2C message")
	rev_msg := msg.Body.(*game.Game_Receive_ChatS2C)
	if rev_msg != nil && rev_msg.Content != nil {
		if *rev_msg.Content == "你好, 杰伦!" {
			// 来自叮当的消息
			send_msg := protos.MarshalProtoMsg(&game.Game_SendChatC2S{
				RoleID:  protos.Uint64(global_god_roleID_dingdang),
				Content: protos.String("你好, 叮当!"),
			})

			clientSession_liupeng.Send(send_msg)
		} else if *rev_msg.Content == "你好, 叮当!" {
			DEBUG("receiveChatCallBack success")
		} else {
			ERR("sendChatC2S error")
		}
	}
}

func TestGameRoleInfoByRoleID(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	DEBUG("send RoleInfoByRoleIDC2S message")
	//杰伦找叮当
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoByRoleIDC2S{
		RoleID: protos.Uint64(global_god_roleID_dingdang),
	})

	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _roleInfoByRoleIDCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive RoleInfoByRoleIDS2C message")
	rev_msg := msg.Body.(*game.Game_RoleInfoByRoleIDS2C)
	err_msg := "roleInfoByRoleID error"
	if rev_msg != nil && rev_msg.Role != nil {
		Assert(global_t, *rev_msg.Role.RoleID > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func TestGameFriendListCallBack(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	DEBUG("send FriendListC2S message")

	send_msg := protos.MarshalProtoMsg(&game.Game_FriendListC2S{})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _friendListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive FriendListS2C message")
	rev_msg := msg.Body.(*game.Game_FriendListS2C)
	err_msg := "friendList error"
	if rev_msg != nil && rev_msg.Friends != nil {
		Assert(global_t, len(rev_msg.Friends) > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func TestGameAddFriend(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	DEBUG("send AddFriendC2S message")

	send_msg := protos.MarshalProtoMsg(&game.Game_AddFriendC2S{
		FriendID: protos.Uint64(global_god_roleID_dingdang),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _addFriendCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive AddFriendS2C message")
	rev_msg := msg.Body.(*game.Game_AddFriendS2C)
	err_msg := "addFriend error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}
