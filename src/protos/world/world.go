package world

import (
	"protos"
	. "protos/msgid"
)

// 因为World 和 Game 公用一部分消息结构体, 默认的init方法会覆盖, 所以自定义方法
func InitWorldMsgID() {
	protos.RegisterMsgID(uint16(MsgID_World_PingC2S), World_PingC2S{})
	protos.RegisterMsgID(uint16(MsgID_World_PingS2C), World_PingS2C{})

	protos.RegisterMsgID(uint16(MsgID_World_RegisterRoleIDC2S), World_RegisterRoleIDC2S{})
	protos.RegisterMsgID(uint16(MsgID_World_RegisterRoleIDS2C), World_RegisterRoleIDS2C{})

	//通用部分
	protos.RegisterMsgID(uint16(MsgID_World_SendChatC2S), World_SendChatC2S{})
	protos.RegisterMsgID(uint16(MsgID_World_SendChatS2C), World_SendChatS2C{})
	protos.RegisterMsgID(uint16(MsgID_World_Receive_ChatS2C), World_Receive_ChatS2C{})
}
