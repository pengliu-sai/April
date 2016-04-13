package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/world"
	. "tools"
)

func worldSendChat(worldSession *link.Session, msg protos.ProtoMsg) {
	INFO("World receive sendChatC2S message")

	sendRoleID := g_worldSession_roleID[worldSession]

	if sendRoleID <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	rev_msg := msg.Body.(*world.World_SendChatC2S)
	roleID := *rev_msg.RoleID
	content := *rev_msg.Content

	falseResult := false
	false_msg := protos.MarshalProtoMsg(&world.World_SendChatS2C{
		Result: &falseResult,
	})

	//查询接收者的session
	otherGameSession := g_roleID_worldSession[roleID]

	if otherGameSession == nil {
		ERR("worldServer don't find role: ", roleID)
		worldSession.Send(false_msg)
		return
	}

	send_chat_msg := protos.MarshalProtoMsg(&world.World_Receive_ChatS2C{
		RoleID:  protos.Uint64(roleID),
		Content: protos.String(content),
	})

	err := otherGameSession.Send(send_chat_msg)

	if err != nil {
		ERR("World send sendChatS2C error")
		worldSession.Send(false_msg)
		return
	}

	trueResult := true
	send_msg := protos.MarshalProtoMsg(&world.World_SendChatS2C{
		Result: &trueResult,
	})

	INFO("World send sendChatS2C message")
	worldSession.Send(send_msg)
}
