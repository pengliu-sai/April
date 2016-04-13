package db

import (
	"protos"
	. "protos/msgid"
)

func InitDBMsgID() {
	protos.RegisterMsgID(uint16(MsgID_DB_UpdateLastLoginTime), DB_UpdateLastLoginTime{})
}
