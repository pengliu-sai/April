package logic

import (
	"github.com/funny/link"
	"protos"
	"protos/admin"
	"protos/game"
	. "protos/msgid"
	. "tools"
	"tools/dispatch"
	"tools/protocol"
)

func isValidGameID_Game_RegisterUserIDS2C(msgID uint16) bool {
	return msgID == uint16(MsgID_Game_RegisterUserIDS2C)
}

func isValidGameID_ButNot_Game_RegisterUserIDS2C(msgID uint16) bool {
	return !isValidGameID_Game_RegisterUserIDS2C(msgID) && protos.IsValidGameID(msgID)
}

func connectGameServer(userID uint64, gameServerAddr string, gateSession *link.Session) (*link.Session, error) {
	handle := dispatch.NewHandleConditions()

	handle.Add(dispatch.HandleCondition{
		Condition: isValidGameID_Game_RegisterUserIDS2C,
		H: dispatch.Handle{
			uint16(MsgID_Game_RegisterUserIDS2C): gameRegisterUserIDCallBack,
		},
	})

	handle.Add(dispatch.HandleFuncCondition{
		Condition: isValidGameID_ButNot_Game_RegisterUserIDS2C,
		H: func(gameClientSession *link.Session, msg []byte) {
			sendMsgToClient(gameClientSession, msg)
		},
	})

	clientMsgDispatch := dispatch.NewDispatch(handle)
	gameClientSession, err := link.Connect("tcp", gameServerAddr, protocol.PackCodecType_Safe)
	if err != nil {
		return nil, err
	}
	waitGroup.Wrap(func() {
		protocol.SessionReceive(gameClientSession, clientMsgDispatch)

	})

	g_gameClientSession_GateSession[gameClientSession] = gateSession
	g_gateSession_GameClientSession[gateSession] = gameClientSession

	return gameClientSession, nil
}

func gameRegisterUserIDCallBack(gameClientSession *link.Session, msg protos.ProtoMsg) {
	INFO("Gate receive registerUserIDS2C message")
	rev_msg := msg.Body.(*game.Game_RegisterUserIDS2C)
	result := *rev_msg.Result
	userID := *rev_msg.UserID
	null_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginS2C{})

	//send Admin_UserLoginS2C
	if !result {
		sendMsgToClient(gameClientSession, null_msg)
		return
	}

	send_msg := protos.MarshalProtoMsg(&admin.Admin_UserLoginS2C{
		UserID: protos.Uint64(userID),
	})

	INFO("Gate send userLoginS2C message")
	sendMsgToClient(gameClientSession, send_msg)

	gateSession := g_gameClientSession_GateSession[gameClientSession]
	if gateSession != nil {
		gateRegisterUserID(gateSession, userID)
	}
}

func sendMsgToClient(gameClientSession *link.Session, msg []byte) {
	gateSession := g_gameClientSession_GateSession[gameClientSession]
	if gateSession != nil {
		gateSession.Send(msg)
	}
}

func sendMsgToGameServer(gateSession *link.Session, msg []byte) {
	gameClientSession := g_gateSession_GameClientSession[gateSession]
	if gameClientSession != nil {
		gameClientSession.Send(msg)
	}
}

func sendMsgToGameServer_RegisterUserID(gateSession *link.Session, userID uint64) {
	send_msg := protos.MarshalProtoMsg(&game.Game_RegisterUserIDC2S{
		UserID: protos.Uint64(userID),
	})
	sendMsgToGameServer(gateSession, send_msg)
}
