package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/game"
	"protos/world"
	. "tools"
)

func gameSendChat(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive sendChatC2S message")
	sendRoleID := g_gameSession_RoleID[gameSession]

	if sendRoleID <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_SendChatC2S)
	roleID := *rev_msg.RoleID
	content := *rev_msg.Content

	falseResult := false
	false_msg := protos.MarshalProtoMsg(&game.Game_SendChatS2C{
		Result: &falseResult,
	})

	//查询接收者的session
	otherGameSession := g_roleID_gameSession[roleID]

	//查询发送者的姓名
	sendRoleName := g_roleID_roleName[sendRoleID]

	//如果不是同一服务器, 则转发给世界服务器处理...
	if otherGameSession == nil {
		send_msg := protos.MarshalProtoMsg(&world.World_SendChatC2S{
			RoleID:  protos.Uint64(roleID),
			Content: protos.String(content),
		})
		INFO("Game send sendChatC2S message to World")
		sendMsgToWorldServer(gameSession, send_msg)
		return
	}

	// 如果接收这存在本服, 则直接发送
	send_chat_msg := protos.MarshalProtoMsg(&game.Game_Receive_ChatS2C{
		RoleID:   protos.Uint64(sendRoleID),
		RoleName: protos.String(sendRoleName),
		Content:  protos.String(content),
	})

	err := otherGameSession.Send(send_chat_msg)
	INFO("Game send receiveChatS2C message")

	if err != nil {
		gameSession.Send(false_msg)
		return
	}

	trueResult := true
	send_msg := protos.MarshalProtoMsg(&game.Game_SendChatS2C{
		Result: &trueResult,
	})

	INFO("Game send sendChatS2C message")
	gameSession.Send(send_msg)
}
