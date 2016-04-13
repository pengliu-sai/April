package system

import (
	"protos"
	. "protos/msgid"
)

func init() {
	protos.RegisterMsgID(uint16(MsgID_System_LogC2S), System_LogC2S{})
	protos.RegisterMsgID(uint16(MsgID_System_LogS2C), System_LogS2C{})
}
