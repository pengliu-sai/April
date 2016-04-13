package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/game"
	. "tools"
	"tools/cfg"
)

//进入场景
func gameEnterScenes(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive enterScenesC2S message")
	userID := g_gameSession_UserID[gameSession]
	if userID <= 0 {
		//未登陆用户
		ERR("user canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_EnterScenesC2S)
	roleID := *rev_msg.RoleID

	falseResult := false
	false_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesS2C{
		Result: &falseResult,
	})

	// 验证roleID
	// 在获取角色列表时, 会缓存角色信息.
	roleInfoList := g_gameSession_RoleInfoList[gameSession]
	if roleInfoList == nil || len(roleInfoList) <= 0 {
		ERR("Game roleInfoList is nil")
		gameSession.Send(false_msg)
		return
	}

	find_role := false
	roleInfoListLen := len(roleInfoList)
	for i := 0; i < roleInfoListLen; i++ {
		roleInfo := roleInfoList[i]
		if *roleInfo.RoleID == roleID {
			find_role = true
			break
		}
	}

	//如果传错参数, 则返回false
	if !find_role {
		ERR("Game roleId no in roleInfoList")
		gameSession.Send(false_msg)
		return
	}

	INFO("try connect WorldServer:", roleID)
	worldServerAddr := cfg.GetServerConfig().WorldServer.IP + ":" + cfg.GetServerConfig().WorldServer.Port
	//开启WorldServer连接
	worldClientSession, err := connectWorldServer(roleID, gameSession, worldServerAddr)

	//连接失败
	if worldClientSession == nil || err != nil {

		ERR("connect world server failure: ", worldServerAddr)
		gameSession.Send(false_msg)
		return
	}

	INFO("connect WorldServer success")

	INFO("Game send registerRoleID to WorldServer: ", roleID)
	//发送消息到WorldServer
	sendMsgToWorldServer_RegisterRoleID(gameSession, roleID)
}

//退出场景
func gameExitScenes(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive exitScenesC2S message")

	if g_gameSession_RoleID[gameSession] <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	gameUnRegisterRoleID(gameSession)

	result := true
	send_msg := protos.MarshalProtoMsg(&game.Game_ExitScenesS2C{
		Result: &result,
	})

	INFO("Game send exitScenesS2C message")
	gameSession.Send(send_msg)
}

func gameRegisterRoleID(gameSession *link.Session, roleID uint64) {
	g_gameSession_RoleID[gameSession] = roleID
	g_roleID_gameSession[roleID] = gameSession

	//查询roleName
	roleInfoList := g_gameSession_RoleInfoList[gameSession]
	if roleInfoList != nil {
		roleInfoListLen := len(roleInfoList)
		for i := 0; i < roleInfoListLen; i++ {
			if *roleInfoList[i].RoleID == roleID {
				g_roleID_roleName[roleID] = *roleInfoList[i].RoleName
				break
			}
		}
	}
}

//切换角色或退出到角色选择界面
func gameUnRegisterRoleID(gameSession *link.Session) {
	roleID := g_gameSession_RoleID[gameSession]
	if roleID > 0 {
		delete(g_gameSession_RoleID, gameSession)
		delete(g_roleID_gameSession, roleID)
	}
}
