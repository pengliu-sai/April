package logic

import (
	"flag"
	"github.com/funny/link"
	"os"
	"protos"
	dbProto "protos/db"
	"protos/game"
	. "protos/msgid"
	"protos/world"
	. "tools"
	"tools/cfg"
	"tools/db"
	"tools/dispatch"
	"tools/protocol"
	"tools/redis"
	"tools/util"
	"protos/config"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

var (
	waitGroup         util.WaitGroupWrapper
	serverMsgDispatch dispatch.DispatchInterface
	gameIndex         int
)

var (
	g_gameSession_UserID map[*link.Session]uint64
	g_gameSession_RoleID map[*link.Session]uint64
	g_roleID_gameSession map[uint64]*link.Session
	g_roleID_roleName    map[uint64]string
)

var (
	g_gameSession_RoleInfoList map[*link.Session][]*game.RoleInfo
)

var (
	g_gameSession_WorldClientSession map[*link.Session]*link.Session
	g_worldClientSession_gameSession map[*link.Session]*link.Session

	developMode bool
)

var (
	globalConfig config.GlobalConfigInfo
	gameConfigFileName = os.Getenv("APRIL_PATH") + "data/config/config.data"
)

func loadGameConfig() {
	globalConfig = config.GlobalConfigInfo{}
	fileByte, err := ioutil.ReadFile(gameConfigFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = proto.Unmarshal(fileByte, &globalConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	INFO("globalConfig: ", *globalConfig.FbInfoList.FbList[0].FBName)

	INFO("loadGameConfigToMemory success!")

}

func init() {
	developMode = cfg.IsDevelopMode()

	game.InitGameMsgID()
	world.InitWorldMsgID()

	dbProto.InitDBMsgID()

}

func InitDB() {
	err := db.Connect("game_db")
	checkError(err)
}

func InitRedis() {
	err := redis.Connect("game_redis")
	checkError(err)
}

func InitGameServer() {
	g_gameSession_UserID = make(map[*link.Session]uint64)
	g_gameSession_RoleID = make(map[*link.Session]uint64)
	g_roleID_gameSession = make(map[uint64]*link.Session)
	g_roleID_roleName = make(map[uint64]string)

	g_gameSession_RoleInfoList = make(map[*link.Session][]*game.RoleInfo)

	g_gameSession_WorldClientSession = make(map[*link.Session]*link.Session)
	g_worldClientSession_gameSession = make(map[*link.Session]*link.Session)

	flag.IntVar(&gameIndex, "game_index", -1, "game server index")
	flag.Parse()
	if gameIndex <= 0 {
		ERR("please set game index")
		os.Exit(-1)
	}

	handle := dispatch.NewHandleConditions()
	//gate -> game 消息处理
	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidGameID,
		H: dispatch.Handle{
			uint16(MsgID_Game_PingC2S):           gamePing,
			uint16(MsgID_Game_RegisterUserIDC2S): gameRegisterUserID,
			uint16(MsgID_Game_RandomRoleNameC2S): gameRandomRoleName,
			uint16(MsgID_Game_RoleCreateC2S):     gameRoleCreate,
			uint16(MsgID_Game_RoleInfoListC2S):   gameRoleInfoList,
			uint16(MsgID_Game_EnterScenesC2S):    gameEnterScenes,
			uint16(MsgID_Game_ExitScenesC2S):     gameExitScenes,

			uint16(MsgID_Game_SendChatC2S):         gameSendChat,
			uint16(MsgID_Game_RoleInfoByRoleIDC2S): gameRoleInfoByRoleID,
			uint16(MsgID_Game_FriendListC2S):       gameFriendList,
			uint16(MsgID_Game_AddFriendC2S):        gameAddFriend,
			uint16(MsgID_Game_FBChapterListC2S): gameFBChapterList,
			uint16(MsgID_Game_FBSectionListC2S): gameFBSectionListByChapterID,

		},
	})

	serverMsgDispatch = dispatch.NewDispatch(handle)
}

func StartGameServer() {
	if developMode {
		WARN("game server in develop mode")
	}

	InitDB()
	InitRedis()
	InitGameServer()

	loadGameConfig()

	addr := "127.0.0.1" + ":" + cfg.GetServerConfig().GameServerList[gameIndex-1].Port

	server, err := link.Serve("tcp", addr, protocol.PackCodecType_Safe)

	if err != nil {
		ERR("FATAL: listen (%s) failed - %s", addr, err)
		os.Exit(-1)
	}

	waitGroup.Wrap(func() {
		protocol.TCPServer(server, func(gameSession *link.Session) {
			//上线
			INFO("user online : ", gameSession.Id())
			gameSession.AddCloseCallback(gameSession, func() {
				//下线
				gameUnRegisterUserID(gameSession)
				INFO("user offline: ", gameSession.Id())
			})
		}, serverMsgDispatch, "GameServer")
	})

	waitGroup.Wait()
}

func checkError(err error) {
	if err != nil {
		ERR("Fatal error: %v", err)
		os.Exit(-1)
	}
}

func gamePing(gameSession *link.Session, msg protos.ProtoMsg) {
	rev_msg := msg.Body.(*game.Game_PingC2S)
	content := rev_msg.Content
	INFO("GamePing Content: ", *content)

	send_msg := protos.MarshalProtoMsg(&game.Game_PingS2C{
		Content: protos.String("Hi GameClient"),
	})
	gameSession.Send(send_msg)
}

func gameRegisterUserID(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive registerUserIDC2S message")
	rev_msg := msg.Body.(*game.Game_RegisterUserIDC2S)
	userID := *rev_msg.UserID
	result := false

	if userID > 0 {
		g_gameSession_UserID[gameSession] = userID
		result = true
	}

	send_msg := protos.MarshalProtoMsg(&game.Game_RegisterUserIDS2C{
		Result: &result,
		UserID: &userID,
	})

	INFO("Game send registerUserIDS2C message")
	gameSession.Send(send_msg)
}

//登出时
func gameUnRegisterUserID(gameSession *link.Session) {
	//下线角色
	gameUnRegisterRoleID(gameSession)
	//下线用户
	delete(g_gameSession_UserID, gameSession)

}
