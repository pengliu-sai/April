package logic

import (
	"github.com/funny/link"
	"os"
	"protos"
	. "protos/msgid"
	"protos/world"
	"sync"
	. "tools"
	"tools/cfg"
	"tools/db"
	"tools/dispatch"
	"tools/protocol"
	"tools/util"
)

var (
	waitGroup         util.WaitGroupWrapper
	serverMsgDispatch dispatch.DispatchInterface
)

var (
	g_roleID_worldSession map[uint64]*link.Session
	g_worldSession_roleID map[*link.Session]uint64

	developMode bool

	mx sync.Mutex
)

func init() {

	developMode = cfg.IsDevelopMode()

	world.InitWorldMsgID()
}

func InitDB() {
	err := db.Connect("game_db")
	checkError(err)
}

func InitWorldServer() {
	g_roleID_worldSession = make(map[uint64]*link.Session)
	g_worldSession_roleID = make(map[*link.Session]uint64)

	handle := dispatch.NewHandleConditions()
	//game -> world 消息处理
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidWorldID,
		H: dispatch.Handle{
			uint16(MsgID_World_PingC2S):           worldPing,
			uint16(MsgID_World_RegisterRoleIDC2S): worldRegisterRoleID,
			uint16(MsgID_World_SendChatC2S):       worldSendChat,
		},
	})

	serverMsgDispatch = dispatch.NewDispatch(handle)
}

func StartWorldServer() {
	if developMode {
		WARN("world server in develop mode")
	}

	InitDB()

	InitWorldServer()

	addr := "127.0.0.1" + ":" + cfg.GetServerConfig().WorldServer.Port

	server, err := link.Serve("tcp", addr, protocol.PackCodecType_Safe)

	if err != nil {
		ERR("FATAL: listen (%s) failed - %s", addr, err)
		os.Exit(-1)
	}

	waitGroup.Wrap(func() {
		protocol.TCPServer(server, func(worldSession *link.Session) {
			//上线
			INFO("game online : ", worldSession.Id())
			worldSession.AddCloseCallback(worldSession, func() {
				//下线
				worldUnRegisterRoleID(worldSession)
				INFO("game offline: ", worldSession.Id())
			})
		}, serverMsgDispatch, "WorldServer")
	})

	waitGroup.Wait()
}

func checkError(err error) {
	if err != nil {
		ERR("Fatal error: %v", err)
		os.Exit(-1)
	}
}

func worldPing(worldSession *link.Session, msg protos.ProtoMsg) {
	rev_msg := msg.Body.(*world.World_PingC2S)
	content := rev_msg.Content
	INFO("WorldPing Content: ", *content)

	send_msg := protos.MarshalProtoMsg(&world.World_PingS2C{
		Content: protos.String("Hi WorldClient"),
	})
	worldSession.Send(send_msg)
}

func worldRegisterRoleID(worldSession *link.Session, msg protos.ProtoMsg) {
	INFO("World receive registerRoleIDC2S message")
	rev_msg := msg.Body.(*world.World_RegisterRoleIDC2S)
	roleID := *rev_msg.RoleID
	result := false
	if roleID > 0 {
		g_roleID_worldSession[roleID] = worldSession
		g_worldSession_roleID[worldSession] = roleID
		result = true
	}

	send_msg := protos.MarshalProtoMsg(&world.World_RegisterRoleIDS2C{
		Result: &result,
		RoleID: protos.Uint64(roleID),
	})
	INFO("World send registerRoleIDS2C message")
	worldSession.Send(send_msg)
}

func worldUnRegisterRoleID(worldSession *link.Session) {
	defer mx.Unlock()
	mx.Lock()
	roleID := g_worldSession_roleID[worldSession]
	if roleID > 0 {
		delete(g_roleID_worldSession, roleID)
		delete(g_worldSession_roleID, worldSession)
	}
}
