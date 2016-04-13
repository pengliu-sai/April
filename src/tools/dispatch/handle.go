package dispatch

import (
	"github.com/funny/binary"
	"github.com/funny/link"

	"container/list"
	"protos"
	. "tools"
)

type HandleInterface interface {
	DealMsg(session *link.Session, msg []byte)
}

//通用Handle
type Handle map[uint16]func(*link.Session, protos.ProtoMsg)

func (this Handle) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	INFO("msgID: ", msgID)
	var protoMsg protos.ProtoMsg
	if protos.IsValidSystemID(msgID) || protos.IsValidGateID(msgID) || protos.IsValidAdminID(msgID) || protos.IsValidGameID(msgID) || protos.IsValidWorldID(msgID) {
		protoMsg = protos.UnmarshalProtoMsg(msg)
	}

	if protoMsg == protos.NullProtoMsg {
		ERR("收到空的proto消息: ", msgID)
		return
	}

	if f, exists := this[msgID]; exists {
		f(session, protoMsg)
	} else {
		ERR("该proto消息未注册handle: ", msgID)
	}
}

//条件Handle
type HandleCondition struct {
	Condition func(uint16) bool
	H         Handle
}

func (this HandleCondition) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	if this.Condition(msgID) {
		this.H.DealMsg(session, msg)
	}
}

//条件函数
type HandleFuncCondition struct {
	Condition func(uint16) bool
	H         func(*link.Session, []byte)
}

func (this HandleFuncCondition) DealMsg(session *link.Session, msg []byte) {
	msgID := binary.GetUint16LE(msg[:2])
	if this.Condition(msgID) {
		this.H(session, msg)
	}
}

//函数
type HandleFunc struct {
	H func(*link.Session, []byte)
}

func (this HandleFunc) DealMsg(session *link.Session, msg []byte) {
	this.H(session, msg)
}

//Handles
type HandleConditions struct {
	hList *list.List
}

func NewHandleConditions() HandleConditions {
	return HandleConditions{
		hList: list.New(),
	}
}

func (this HandleConditions) Add(handle HandleInterface) {
	this.hList.PushBack(handle)
}

func (this HandleConditions) DealMsg(session *link.Session, msg []byte) {
	for e := this.hList.Front(); e != nil; e = e.Next() {
		e.Value.(HandleInterface).DealMsg(session, msg)
	}
}
