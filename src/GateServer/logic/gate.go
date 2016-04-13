package logic

import (
	"github.com/funny/link"
	"os"
	"protos"
	"protos/game"
	"protos/gate"
	. "protos/msgid"
	. "tools"
	"tools/cfg"
	"tools/dispatch"
	"tools/protocol"
	"tools/util"
)

var (
	waitGroup         util.WaitGroupWrapper
	serverMsgDispatch dispatch.DispatchInterface
	adminAddr         string
)

const (
	max_userConnectNum = 3000 //最大连接数3000
)

var (
	//用户连接数量, 用于对用户进行流量控制
	g_userConnectNum uint64
	//用户登陆成功后, 记录gate 和 userID 之前的关系
	g_gateSession_UserID map[*link.Session]uint64
	g_userID_GateSession map[uint64]*link.Session

	g_gateSession_GameClientSession map[*link.Session]*link.Session
	g_gameClientSession_GateSession map[*link.Session]*link.Session

	developMode bool
)

func init() {
	developMode = cfg.IsDevelopMode()
	g_userConnectNum = 0
	g_gateSession_UserID = make(map[*link.Session]uint64)
	g_userID_GateSession = make(map[uint64]*link.Session)

	g_gateSession_GameClientSession = make(map[*link.Session]*link.Session)
	g_gameClientSession_GateSession = make(map[*link.Session]*link.Session)

	adminAddr = cfg.GetServerConfig().AdminServer.IP + ":" + cfg.GetServerConfig().AdminServer.Port

	handle := dispatch.NewHandleConditions()

	//client->gate消息处理
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidGateID,
		H: dispatch.Handle{
			uint16(MsgID_Gate_PingC2S): gatePing,
		},
	})

	//client->admin消息处理
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidAdminID,
		H: dispatch.Handle{
			uint16(MsgID_Admin_UserLoginC2S):    userLogin,
			uint16(MsgID_Admin_UserRegisterC2S): userRegister,
			uint16(MsgID_Admin_UserExitC2S):     userExit,
		},
	})

	//client->game消息处理
	handle.Add(dispatch.HandleFuncCondition{
		Condition: protos.IsValidGameID,
		H: func(gateSession *link.Session, msg []byte) {
			sendMsgToGameServer(gateSession, msg)
		},
	})

	serverMsgDispatch = dispatch.NewDispatch(handle)

}

func StartGateServer() {
	if developMode {
		WARN("gate server in develop mode")
	}

	game.InitGameMsgID()

	addr := "0.0.0.0:" + cfg.GetServerConfig().GateServer.Port

	server, err := link.Serve("tcp", addr, protocol.PackCodecType_Safe)

	if err != nil {
		ERR("FATAL: listen (%s) failed - %s", addr, err)
		os.Exit(-1)
	}

	waitGroup.Wrap(func() {
		protocol.TCPServer(server, func(gateSession *link.Session) {
			//上线
			INFO("user online : ", gateSession.Id())
			g_userConnectNum++
			gateSession.AddCloseCallback(gateSession, func() {
				g_userConnectNum--
				//下线
				gateUnRegisterUserID(gateSession)
				INFO("user offline: ", gateSession.Id())
			})
		}, serverMsgDispatch, "GateServer")
	})

	waitGroup.Wait()
}

func gatePing(gateSession *link.Session, msg protos.ProtoMsg) {
	rev_msg := msg.Body.(*gate.Gate_PingC2S)

	content := rev_msg.Content
	INFO("GatePing Content: ", *content)

	send_msg := protos.MarshalProtoMsg(&gate.Gate_PingS2C{
		Content: protos.String("Hi GateClient"),
	})
	gateSession.Send(send_msg)
}

func gateRegisterUserID(gateSession *link.Session, userID uint64) {
	g_gateSession_UserID[gateSession] = userID
	g_userID_GateSession[userID] = gateSession
}

//登出时
func gateUnRegisterUserID(gateSession *link.Session) {
	userID := g_gateSession_UserID[gateSession]

	if userID > 0 {
		delete(g_gateSession_UserID, gateSession)
		delete(g_userID_GateSession, userID)
	}
}
