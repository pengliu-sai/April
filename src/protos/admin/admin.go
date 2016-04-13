package admin

import (
	"protos"
	. "protos/msgid"
)

func init() {
	protos.RegisterMsgID(uint16(MsgID_Admin_UserLoginC2S), Admin_UserLoginC2S{})
	protos.RegisterMsgID(uint16(MsgID_Admin_UserLoginS2C), Admin_UserLoginS2C{})

	protos.RegisterMsgID(uint16(MsgID_Admin_UserRegisterC2S), Admin_UserRegisterC2S{})
	protos.RegisterMsgID(uint16(MsgID_Admin_UserRegisterS2C), Admin_UserRegisterS2C{})

	protos.RegisterMsgID(uint16(MsgID_Admin_UserExitC2S), Admin_UserExitC2S{})
	protos.RegisterMsgID(uint16(MsgID_Admin_UserExitS2C), Admin_UserExitS2C{})

}
