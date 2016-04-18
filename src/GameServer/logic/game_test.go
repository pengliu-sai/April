package logic

import (
	"protos"
	. "protos/msgid"
	"testing"
	"tools/cfg"
	"tools/dispatch"

	"github.com/funny/link"
	"protos/game"
	. "tools"
	"tools/protocol"
	. "tools/unitest"
)

var (
	clientMsgDispatch dispatch.DispatchInterface
	gameAddr          string
	clientSession     *link.Session

	global_t *testing.T
)

/*

	测试数据.
	用户: 刘鹏, God
	角色: 杰伦是刘鹏的, 叮当是上帝的, 喵
*/

var (
	global_gameIndex             int
	global_gameIndex_liupeng     int
	global_gameIndex_god         int
	global_userID_liupeng        = uint64(915441286658396160)
	global_userID_god            = uint64(915664487518834688)
	global_liupeng_roleID_jielun = uint64(915771140146728960)
	global_god_roleID_dingdang   = uint64(915787632821145600)
)

var (
	clientSession_liupeng *link.Session
	clientSession_god     *link.Session
)

func init() {
	global_gameIndex = 0
	gameAddr = cfg.GetServerConfig().GameServerList[global_gameIndex].IP + ":" + cfg.GetServerConfig().GameServerList[global_gameIndex].Port
	handle := dispatch.NewHandleConditions()

	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidGameID,
		H: dispatch.Handle{
			uint16(MsgID_Game_PingS2C):             _pingCallBack,
			uint16(MsgID_Game_RegisterUserIDS2C):   _registerUserIDCallBack,
			uint16(MsgID_Game_RandomRoleNameS2C):   _randomRoleNameCallBack,
			uint16(MsgID_Game_RoleCreateS2C):       _roleCreateCallBack,
			uint16(MsgID_Game_RoleInfoListS2C):     _roleInfoListCallBack,
			uint16(MsgID_Game_EnterScenesS2C):      _enterScenesCallBack,
			uint16(MsgID_Game_ExitScenesS2C):       _exitScenesCallBack,
			uint16(MsgID_Game_SendChatS2C):         _sendChatCallBack,
			uint16(MsgID_Game_Receive_ChatS2C):     _receiveChatCallBack,
			uint16(MsgID_Game_RoleInfoByRoleIDS2C): _roleInfoByRoleIDCallBack,
			uint16(MsgID_Game_FriendListS2C):       _friendListCallBack,
			uint16(MsgID_Game_AddFriendS2C):        _addFriendCallBack,
			uint16(MsgID_Game_FBChapterListS2C):    _fbChapterListCallBack,
			uint16(MsgID_Game_FBSectionListS2C):    _fbSectionListCallBack,
		},
	})
	clientMsgDispatch = dispatch.NewDispatch(handle)
}

func connectGameServer() {
	var err error
	clientSession, err = link.Connect("tcp", gameAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectGameServer_liupeng() {
	gameAddr = cfg.GetServerConfig().GameServerList[global_gameIndex_liupeng].IP + ":" + cfg.GetServerConfig().GameServerList[global_gameIndex_liupeng].Port

	var err error
	clientSession_liupeng, err = link.Connect("tcp", gameAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_liupeng, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectGameServer_god() {
	gameAddr = cfg.GetServerConfig().GameServerList[global_gameIndex_god].IP + ":" + cfg.GetServerConfig().GameServerList[global_gameIndex_god].Port

	var err error
	clientSession_god, err = link.Connect("tcp", gameAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_god, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func TestGameConnect(t *testing.T) {
	connectGameServer()

	global_t = t
	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_PingC2S{
		Content: protos.String("Hi GameServer"),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _pingCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	rev_msg := msg.Body.(*game.Game_PingS2C)
	Equal(global_t, *rev_msg.Content, "Hi GameClient")
}

func registerUserID_liupeng() {
	waitGroup.Add(1)
	DEBUG("send registerUserIDC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RegisterUserIDC2S{
		UserID: protos.Uint64(global_userID_liupeng),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func registerUserID_god() {
	waitGroup.Add(1)
	DEBUG("send registerUserIDC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RegisterUserIDC2S{
		UserID: protos.Uint64(global_userID_god),
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

func enterScenes_jielun() {
	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func enterScenes_dingdang() {
	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_god_roleID_dingdang),
	})
	clientSession_god.Send(send_msg)
	waitGroup.Wait()
}

func Test_RegisterUserID(t *testing.T) {
	connectGameServer_liupeng()

	global_t = t

	waitGroup.Add(1)
	DEBUG("send registerUserIDC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RegisterUserIDC2S{
		UserID: protos.Uint64(global_userID_liupeng),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _registerUserIDCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive registerUserIDS2C message")
	rev_msg := msg.Body.(*game.Game_RegisterUserIDS2C)
	err_msg := "registerUserID error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_RandomRoleName(t *testing.T) {
	connectGameServer_liupeng()

	enterScenes_jielun()

	global_t = t
	waitGroup.Add(1)
	DEBUG("send randomRoleNameC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RandomRoleNameC2S{
		Sex: game.SexEnum_male.Enum(),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _randomRoleNameCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive randomRoleNameS2C message")
	rev_msg := msg.Body.(*game.Game_RandomRoleNameS2C)
	err_msg := "randomRoleName error"
	if rev_msg != nil && rev_msg.Name != nil {
		Assert(global_t, len(*rev_msg.Name) > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_RoleCreate(t *testing.T) {
	connectGameServer()

	global_t = t
	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateC2S{
		Name: protos.String("叮当"),
		Sex:  game.SexEnum_male.Enum(),
		Race: game.RacesEnum_archer.Enum(),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _roleCreateCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	rev_msg := msg.Body.(*game.Game_RoleCreateS2C)
	errStr := "roleCreate failure"
	if rev_msg.RoleID != nil {
		Assert(global_t, *rev_msg.RoleID > 0, errStr)
	} else {
		global_t.Error(errStr)
	}
}

func Test_RoleInfoList(t *testing.T) {
	connectGameServer_liupeng()

	registerUserID_liupeng()

	global_t = t
	waitGroup.Add(1)
	DEBUG("send roleInfoListC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListC2S{})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _roleInfoListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive roleInfoListS2C message")
	rev_msg := msg.Body.(*game.Game_RoleInfoListS2C)
	err_msg := "roleInfoList failure"
	if rev_msg.Roles != nil {
		Assert(global_t, len(rev_msg.Roles) > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_EnterScenes(t *testing.T) {
	Test_RoleInfoList(t)

	waitGroup.Add(1)
	DEBUG("send enterScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _enterScenesCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive enterScenesS2C message")
	rev_msg := msg.Body.(*game.Game_EnterScenesS2C)
	err_msg := "enterScenes error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)

	} else {
		ERR(err_msg)
	}
}

func Test_ExitScenes(t *testing.T) {
	Test_EnterScenes(t)
	waitGroup.Add(1)
	DEBUG("send exitScenesC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_ExitScenesC2S{})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _exitScenesCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive exitScenesS2C message")
	rev_msg := msg.Body.(*game.Game_ExitScenesS2C)
	err_msg := "exitScenes error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}

//杰伦上线
func TestOnline_jielun(t *testing.T) {
	//连接game
	global_t = t
	connectGameServer_liupeng()
	registerUserID_liupeng()
	roleInfoList_liupeng()
	enterScenes_jielun()
}

//叮当上线
func TestOnline_dingdang(t *testing.T) {
	connectGameServer_god()
	registerUserID_god()
	roleInfoList_god()
	enterScenes_dingdang()
}

// 连接同一台服务器
func Test_Chat_Interactive(t *testing.T) {
	TestOnline_jielun(t)

	TestOnline_dingdang(t)

	global_t = t

	waitGroup.Add(4)

	//叮当向杰伦say Hi!
	send_msg := protos.MarshalProtoMsg(&game.Game_SendChatC2S{
		RoleID:  protos.Uint64(global_liupeng_roleID_jielun),
		Content: protos.String("你好, 杰伦!"),
	})

	DEBUG("send sendChatC2S message")
	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

func _sendChatCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive SendChatS2C message")
	rev_msg := msg.Body.(*game.Game_SendChatS2C)
	err_msg := "sendChat error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}

func _receiveChatCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive ReceiveChatS2C message")

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

//跨服聊天
func Test_Chat_World_Interactive(t *testing.T) {
	global_gameIndex_liupeng = 2
	TestOnline_jielun(t)
	global_gameIndex_god = 3
	TestOnline_dingdang(t)

	waitGroup.Add(4)

	//叮当向杰伦say Hi!
	send_msg := protos.MarshalProtoMsg(&game.Game_SendChatC2S{
		RoleID:  protos.Uint64(global_liupeng_roleID_jielun),
		Content: protos.String("你好, 杰伦!"),
	})

	DEBUG("send sendChatC2S message")

	clientSession_god.Send(send_msg)

	waitGroup.Wait()
}

func Test_RoleInfoByRoleID(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoByRoleIDC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	DEBUG("send roleInfoByRoleIDC2S message")
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _roleInfoByRoleIDCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive roleInfoByRoleIDS2C message")

	rev_msg := msg.Body.(*game.Game_RoleInfoByRoleIDS2C)
	err_msg := "roleInfoByRoleID error"
	if rev_msg != nil && rev_msg.Role != nil {
		Assert(global_t, *rev_msg.Role.RoleID > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_FriendList(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)

	DEBUG("send friendListC2S message")
	send_msg := protos.MarshalProtoMsg(&game.Game_FriendListC2S{})
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _friendListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive friendListS2C message")
	rev_msg := msg.Body.(*game.Game_FriendListS2C)
	err_msg := "friendList error"
	if rev_msg.Friends != nil {
		Assert(global_t, len(rev_msg.Friends) > 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_AddFriend(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_AddFriendC2S{
		FriendID: protos.Uint64(global_god_roleID_dingdang),
	})
	DEBUG("send addFriendC2S message")
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _addFriendCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive addFriendS2C message")
	rev_msg := msg.Body.(*game.Game_AddFriendS2C)
	err_msg := "addFriend error"
	if rev_msg != nil && rev_msg.Result != nil {
		Assert(global_t, *rev_msg.Result, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_FBChapterList(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_FBChapterListC2S{})
	DEBUG("send fbChapterList message")
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}


func _fbChapterListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive fbChapterListS2C message")
	rev_msg := msg.Body.(*game.Game_FBChapterListS2C)
	err_msg := "fbChapterList error"
	if rev_msg != nil && rev_msg.FbChapterList != nil {
		Assert(global_t, len(rev_msg.FbChapterList)> 0, err_msg)
	} else {
		ERR(err_msg)
	}
}

func Test_FBSectionList(t *testing.T) {
	TestOnline_jielun(t)

	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&game.Game_FBSectionListC2S{
		ChapterID: protos.Int64(1),
	})
	DEBUG("send fbSectionList message")
	clientSession_liupeng.Send(send_msg)
	waitGroup.Wait()
}

func _fbSectionListCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive fbSectionListS2C message")
	rev_msg := msg.Body.(*game.Game_FBSectionListS2C)
	err_msg := "fbSectionList error"
	if rev_msg != nil && rev_msg.FbInfoList != nil {
		Assert(global_t, len(rev_msg.FbInfoList) > 0, err_msg)
	} else {
		ERR(err_msg)
	}

}