package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/game"
	. "protos/msgid"
	"protos/world"
	. "tools"
	"tools/dispatch"
	"tools/protocol"
)

func connectWorldServer(roleID uint64, gameSession *link.Session, worldServerAddr string) (*link.Session, error) {
	handle := dispatch.NewHandleConditions()
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidWorldID,
		H: dispatch.Handle{
			uint16(MsgID_World_RegisterRoleIDS2C): worldRegisterRoleIDCallBack,
			uint16(MsgID_World_SendChatS2C):       worldSendChatCallBack,
			uint16(MsgID_World_Receive_ChatS2C):   worldReceiveChatCallBack,
		},
	})

	clientMsgDispatch := dispatch.NewDispatch(handle)

	worldClientSession, err := link.Connect("tcp", worldServerAddr, protocol.PackCodecType_Safe)
	if err != nil {
		return nil, err
	}
	waitGroup.Wrap(func() {
		protocol.SessionReceive(worldClientSession, clientMsgDispatch)

	})

	g_gameSession_WorldClientSession[gameSession] = worldClientSession
	g_worldClientSession_gameSession[worldClientSession] = gameSession

	return worldClientSession, nil
}

func worldRegisterRoleIDCallBack(worldClientSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive registerRoleIDS2C message")
	rev_msg := msg.Body.(*world.World_RegisterRoleIDS2C)
	result := *rev_msg.Result
	roleID := *rev_msg.RoleID

	send_msg := protos.MarshalProtoMsg(&game.Game_EnterScenesS2C{
		Result: &result,
	})

	INFO("Game send gameEneterScenesS2C message")
	sendMsgToClient(worldClientSession, send_msg)

	gameSession := g_worldClientSession_gameSession[worldClientSession]
	if gameSession != nil {
		gameRegisterRoleID(gameSession, roleID)

	}
}

func worldSendChatCallBack(worldClientSession *link.Session, msg protos.ProtoMsg) {

	rev_msg := msg.Body.(*world.World_SendChatS2C)

	send_msg := protos.MarshalProtoMsg(&game.Game_SendChatS2C{
		Result: rev_msg.Result,
	})
	sendMsgToClient(worldClientSession, send_msg)

}

func worldReceiveChatCallBack(worldClientSession *link.Session, msg protos.ProtoMsg) {

	rev_msg := msg.Body.(*world.World_Receive_ChatS2C)
	roleID := *rev_msg.RoleID
	content := *rev_msg.Content
	//通过roleID, 查找roleName
	roleName := g_roleID_roleName[*rev_msg.RoleID]

	send_msg := protos.MarshalProtoMsg(&game.Game_Receive_ChatS2C{
		RoleID:   protos.Uint64(roleID),
		RoleName: protos.String(roleName),
		Content:  protos.String(content),
	})

	sendMsgToClient(worldClientSession, send_msg)
}

func sendMsgToClient(worldClientSession *link.Session, msg []byte) {
	gameSession := g_worldClientSession_gameSession[worldClientSession]
	if gameSession != nil {
		gameSession.Send(msg)
	}
}

func sendMsgToWorldServer(gameSession *link.Session, msg []byte) {
	worldClientSession := g_gameSession_WorldClientSession[gameSession]
	if worldClientSession != nil {
		worldClientSession.Send(msg)
	}
}

func sendMsgToWorldServer_RegisterRoleID(gameSession *link.Session, roleID uint64) {
	send_msg := protos.MarshalProtoMsg(&world.World_RegisterRoleIDC2S{
		RoleID: &roleID,
	})
	sendMsgToWorldServer(gameSession, send_msg)
}
