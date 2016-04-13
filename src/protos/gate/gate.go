package gate

import (
	"protos"
	. "protos/msgid"
)

func init() {
	protos.RegisterMsgID(uint16(MsgID_Gate_PingC2S), Gate_PingC2S{})
	protos.RegisterMsgID(uint16(MsgID_Gate_PingS2C), Gate_PingS2C{})
}
