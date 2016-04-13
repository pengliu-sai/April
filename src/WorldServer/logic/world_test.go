package logic

import (
	"github.com/funny/link"
	"protos"
	. "protos/msgid"
	"testing"
	"tools/cfg"
	"tools/dispatch"

	"protos/world"
	. "tools"
	"tools/protocol"
	. "tools/unitest"
)

var (
	clientMsgDispatch dispatch.DispatchInterface
	worldAddr         string
	clientSession     *link.Session

	global_t *testing.T
)

/*

	测试数据.
	用户: 刘鹏, God
	角色: 杰伦是刘鹏的, 叮当是上帝的, 喵
*/

var (
	global_liupeng_roleID_jielun = uint64(915771140146728960)
	global_god_roleID_dingdang   = uint64(915787632821145600)
)

var (
	clientSession_jielun   *link.Session
	clientSession_dingdang *link.Session
)

func init() {
	worldAddr = cfg.GetServerConfig().WorldServer.IP + ":" + cfg.GetServerConfig().WorldServer.Port
	handle := dispatch.NewHandleConditions()

	handle.Add(dispatch.HandleCondition{
		Condition: protos.IsValidWorldID,
		H: dispatch.Handle{
			uint16(MsgID_World_PingS2C):           _pingCallBack,
			uint16(MsgID_World_RegisterRoleIDS2C): _registerRoleIDCallBack,
			uint16(MsgID_World_SendChatS2C):       _sendChatCallBack,
			uint16(MsgID_World_Receive_ChatS2C):   _receiveChatCallBack,
		},
	})
	clientMsgDispatch = dispatch.NewDispatch(handle)
}

func connectWorldServer() {
	var err error
	clientSession, err = link.Connect("tcp", worldAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectWorldServer_jielun() {
	var err error
	clientSession_jielun, err = link.Connect("tcp", worldAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_jielun, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func connectWorldServer_dingdang() {
	var err error
	clientSession_dingdang, err = link.Connect("tcp", worldAddr, protocol.PackCodecType_Safe)
	if err != nil {
		panic(err)
	}
	waitGroup.Wrap(func() {
		go protocol.SessionReceive(clientSession_dingdang, clientMsgDispatch)

	})
	waitGroup.Wait()
}

func registerRoleID_Jielun() {
	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&world.World_RegisterRoleIDC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession_jielun.Send(send_msg)
	waitGroup.Wait()
}

func registerRoleID_dingdang() {
	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&world.World_RegisterRoleIDC2S{
		RoleID: protos.Uint64(global_god_roleID_dingdang),
	})
	clientSession_dingdang.Send(send_msg)
}

func TestConnectWorldServer(t *testing.T) {
	connectWorldServer()

	global_t = t
	waitGroup.Add(1)
	send_msg := protos.MarshalProtoMsg(&world.World_PingC2S{
		Content: protos.String("Hi WorldServer"),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _pingCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	rev_msg := msg.Body.(*world.World_PingS2C)
	Equal(global_t, *rev_msg.Content, "Hi WorldClient")
}

func Test_RegisterRoleID(t *testing.T) {
	connectWorldServer()

	global_t = t
	waitGroup.Add(1)
	DEBUG("send registerRoleIDC2S message")
	send_msg := protos.MarshalProtoMsg(&world.World_RegisterRoleIDC2S{
		RoleID: protos.Uint64(global_liupeng_roleID_jielun),
	})
	clientSession.Send(send_msg)
	waitGroup.Wait()
}

func _registerRoleIDCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive registerRoleIDS2C message")
	rev_msg := msg.Body.(*world.World_RegisterRoleIDS2C)
	err_msg := "registerRoleID error"
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
	connectWorldServer_jielun()
	registerRoleID_Jielun()
}

//叮当上线
func TestOnline_dingdang(t *testing.T) {
	connectWorldServer_dingdang()
	registerRoleID_dingdang()
}

func Test_Chat_Interactive(t *testing.T) {
	TestOnline_jielun(t)
	TestOnline_dingdang(t)

	waitGroup.Add(4)

	//叮当向杰伦say Hi!
	send_msg := protos.MarshalProtoMsg(&world.World_SendChatC2S{
		RoleID:  protos.Uint64(global_liupeng_roleID_jielun),
		Content: protos.String("你好, 杰伦!"),
	})

	DEBUG("send sendChatC2S message")
	clientSession_dingdang.Send(send_msg)

	waitGroup.Wait()

}

func _sendChatCallBack(session *link.Session, msg protos.ProtoMsg) {
	defer waitGroup.Done()
	DEBUG("receive SendChatS2C message")
	rev_msg := msg.Body.(*world.World_SendChatS2C)
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

	rev_msg := msg.Body.(*world.World_Receive_ChatS2C)

	if rev_msg != nil && rev_msg.Content != nil {
		if *rev_msg.Content == "你好, 杰伦!" {
			// 来自叮当的消息
			send_msg := protos.MarshalProtoMsg(&world.World_SendChatC2S{
				RoleID:  protos.Uint64(global_god_roleID_dingdang),
				Content: protos.String("你好, 叮当!"),
			})

			clientSession_jielun.Send(send_msg)
		} else if *rev_msg.Content == "你好, 叮当!" {
			DEBUG("receiveChatCallBack success")
		} else {
			ERR("sendChatC2S error")
		}
	}
}
