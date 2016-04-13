package protos

import (
	"github.com/funny/binary"
	"github.com/golang/protobuf/proto"

	"reflect"
)
import (
	. "tools"
)

type ProtoMsg struct {
	ID   uint16
	Body interface{}
}

var (
	NullProtoMsg ProtoMsg                = ProtoMsg{0, nil}
	MsgObjectMap map[uint16]reflect.Type = make(map[uint16]reflect.Type)
	MsgIDMap     map[reflect.Type]uint16 = make(map[reflect.Type]uint16)
)

func RegisterMsgID(msgID uint16, data interface{}) {
	msgType := reflect.TypeOf(data)

	MsgObjectMap[msgID] = msgType
	MsgIDMap[reflect.TypeOf(reflect.New(msgType).Interface())] = msgID
}

//根据消息ID获取消息实体
func GetMsgObject(msgID uint16) proto.Message {
	if msgType, exists := MsgObjectMap[msgID]; exists {
		return reflect.New(msgType).Interface().(proto.Message)
	} else {
		ERR("No MsgID:", msgID)
	}
	return nil
}

//根据一条消息获取消息ID
func GetMsgID(msg interface{}) uint16 {
	msgType := reflect.TypeOf(msg)
	if msgID, exists := MsgIDMap[msgType]; exists {
		return msgID
	} else {
		ERR("No MsgType:", msgType)

	}
	return 0
}

//序列化
func MarshalProtoMsg(args proto.Message) []byte {
	msgID := GetMsgID(args)
	msgBody, _ := proto.Marshal(args)

	result := make([]byte, 2+len(msgBody))
	binary.PutUint16LE(result[:2], msgID)
	copy(result[2:], msgBody)

	return result
}

//反序列化
func UnmarshalProtoMsg(msg []byte) ProtoMsg {
	if len(msg) < 2 {
		return NullProtoMsg
	}

	msgID := binary.GetUint16LE(msg[:2])
	msgBody := GetMsgObject(msgID)
	if msgBody == nil {
		return NullProtoMsg
	}

	err := proto.Unmarshal(msg[2:], msgBody)
	if err != nil {
		return NullProtoMsg
	}

	return ProtoMsg{
		ID:   msgID,
		Body: msgBody,
	}
}

//封装消息String类型字段
func String(param string) *string {
	return proto.String(param)
}

//封装消息Uint64类型字段
func Uint64(param uint64) *uint64 {
	return proto.Uint64(param)
}

//封装消息Int64类型字段
func Int64(param int64) *int64 {
	return proto.Int64(param)
}

//封装消息Int32类型字段
func Int32(param int32) *int32 {
	return proto.Int32(param)
}

//封装消息Uint32类型字段
func Uint32(param uint32) *uint32 {
	return proto.Uint32(param)
}

//是否是系统消息
func IsValidSystemID(msgID uint16) bool {
	return msgID >= 1 && msgID <= 900
}

func IsValidGateID(msgID uint16) bool {
	return msgID >= 901 && msgID <= 1000
}

func IsValidAdminID(msgID uint16) bool {
	return msgID >= 1001 && msgID <= 2000
}

func IsValidGameID(msgID uint16) bool {
	return msgID >= 2001 && msgID <= 5000
}

func IsValidWorldID(msgID uint16) bool {
	return msgID >= 5001 && msgID <= 6000
}
