// Code generated by protoc-gen-go.
// source: src/protos/msgid/msgid.proto
// DO NOT EDIT!

/*
Package msgid is a generated protocol buffer package.

It is generated from these files:
	src/protos/msgid/msgid.proto

It has these top-level messages:
*/
package msgid

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MsgID int32

const (
	// ------[[system 1-900]]-----
	MsgID_System_LogC2S MsgID = 1
	MsgID_System_LogS2C MsgID = 2
	// ------[[gate 901-1000]]--------
	MsgID_Gate_PingC2S MsgID = 901
	MsgID_Gate_PingS2C MsgID = 902
	// -----[[admin 1001-2000]]-------
	MsgID_Admin_UserLoginC2S    MsgID = 1001
	MsgID_Admin_UserLoginS2C    MsgID = 1002
	MsgID_Admin_UserRegisterC2S MsgID = 1003
	MsgID_Admin_UserRegisterS2C MsgID = 1004
	MsgID_Admin_UserExitC2S     MsgID = 1005
	MsgID_Admin_UserExitS2C     MsgID = 1006
	// -----[[game 2001-5000]]-------
	MsgID_Game_PingC2S             MsgID = 2001
	MsgID_Game_PingS2C             MsgID = 2002
	MsgID_Game_RegisterUserIDC2S   MsgID = 2003
	MsgID_Game_RegisterUserIDS2C   MsgID = 2004
	MsgID_Game_RegisterRoleIDC2S   MsgID = 2005
	MsgID_Game_RegisterRoleIDS2C   MsgID = 2006
	MsgID_Game_RandomRoleNameC2S   MsgID = 2007
	MsgID_Game_RandomRoleNameS2C   MsgID = 2008
	MsgID_Game_RoleCreateC2S       MsgID = 2009
	MsgID_Game_RoleCreateS2C       MsgID = 2010
	MsgID_Game_RoleInfoListC2S     MsgID = 2011
	MsgID_Game_RoleInfoListS2C     MsgID = 2012
	MsgID_Game_EnterScenesC2S      MsgID = 2013
	MsgID_Game_EnterScenesS2C      MsgID = 2014
	MsgID_Game_ExitScenesC2S       MsgID = 2015
	MsgID_Game_ExitScenesS2C       MsgID = 2016
	MsgID_Game_RoleInfoByRoleIDC2S MsgID = 2017
	MsgID_Game_RoleInfoByRoleIDS2C MsgID = 2018
	MsgID_Game_FriendListC2S       MsgID = 2019
	MsgID_Game_FriendListS2C       MsgID = 2020
	MsgID_Game_AddFriendC2S        MsgID = 2021
	MsgID_Game_AddFriendS2C        MsgID = 2022
	// game-world通用的消息ID-->start
	MsgID_Game_SendChatC2S     MsgID = 2023
	MsgID_Game_SendChatS2C     MsgID = 2024
	MsgID_Game_Receive_ChatS2C MsgID = 2026
	// -----[[world 5001-6000]]------
	MsgID_World_PingC2S           MsgID = 5001
	MsgID_World_PingS2C           MsgID = 5002
	MsgID_World_RegisterRoleIDC2S MsgID = 5003
	MsgID_World_RegisterRoleIDS2C MsgID = 5004
	// world-game通用的消息ID-->start
	MsgID_World_SendChatC2S     MsgID = 5005
	MsgID_World_SendChatS2C     MsgID = 5006
	MsgID_World_Receive_ChatS2C MsgID = 5008
	// DB不要求回调, 不区分奇偶
	MsgID_DB_UpdateLastLoginTime MsgID = 6001
)

var MsgID_name = map[int32]string{
	1:    "System_LogC2S",
	2:    "System_LogS2C",
	901:  "Gate_PingC2S",
	902:  "Gate_PingS2C",
	1001: "Admin_UserLoginC2S",
	1002: "Admin_UserLoginS2C",
	1003: "Admin_UserRegisterC2S",
	1004: "Admin_UserRegisterS2C",
	1005: "Admin_UserExitC2S",
	1006: "Admin_UserExitS2C",
	2001: "Game_PingC2S",
	2002: "Game_PingS2C",
	2003: "Game_RegisterUserIDC2S",
	2004: "Game_RegisterUserIDS2C",
	2005: "Game_RegisterRoleIDC2S",
	2006: "Game_RegisterRoleIDS2C",
	2007: "Game_RandomRoleNameC2S",
	2008: "Game_RandomRoleNameS2C",
	2009: "Game_RoleCreateC2S",
	2010: "Game_RoleCreateS2C",
	2011: "Game_RoleInfoListC2S",
	2012: "Game_RoleInfoListS2C",
	2013: "Game_EnterScenesC2S",
	2014: "Game_EnterScenesS2C",
	2015: "Game_ExitScenesC2S",
	2016: "Game_ExitScenesS2C",
	2017: "Game_RoleInfoByRoleIDC2S",
	2018: "Game_RoleInfoByRoleIDS2C",
	2019: "Game_FriendListC2S",
	2020: "Game_FriendListS2C",
	2021: "Game_AddFriendC2S",
	2022: "Game_AddFriendS2C",
	2023: "Game_SendChatC2S",
	2024: "Game_SendChatS2C",
	2026: "Game_Receive_ChatS2C",
	5001: "World_PingC2S",
	5002: "World_PingS2C",
	5003: "World_RegisterRoleIDC2S",
	5004: "World_RegisterRoleIDS2C",
	5005: "World_SendChatC2S",
	5006: "World_SendChatS2C",
	5008: "World_Receive_ChatS2C",
	6001: "DB_UpdateLastLoginTime",
}
var MsgID_value = map[string]int32{
	"System_LogC2S":            1,
	"System_LogS2C":            2,
	"Gate_PingC2S":             901,
	"Gate_PingS2C":             902,
	"Admin_UserLoginC2S":       1001,
	"Admin_UserLoginS2C":       1002,
	"Admin_UserRegisterC2S":    1003,
	"Admin_UserRegisterS2C":    1004,
	"Admin_UserExitC2S":        1005,
	"Admin_UserExitS2C":        1006,
	"Game_PingC2S":             2001,
	"Game_PingS2C":             2002,
	"Game_RegisterUserIDC2S":   2003,
	"Game_RegisterUserIDS2C":   2004,
	"Game_RegisterRoleIDC2S":   2005,
	"Game_RegisterRoleIDS2C":   2006,
	"Game_RandomRoleNameC2S":   2007,
	"Game_RandomRoleNameS2C":   2008,
	"Game_RoleCreateC2S":       2009,
	"Game_RoleCreateS2C":       2010,
	"Game_RoleInfoListC2S":     2011,
	"Game_RoleInfoListS2C":     2012,
	"Game_EnterScenesC2S":      2013,
	"Game_EnterScenesS2C":      2014,
	"Game_ExitScenesC2S":       2015,
	"Game_ExitScenesS2C":       2016,
	"Game_RoleInfoByRoleIDC2S": 2017,
	"Game_RoleInfoByRoleIDS2C": 2018,
	"Game_FriendListC2S":       2019,
	"Game_FriendListS2C":       2020,
	"Game_AddFriendC2S":        2021,
	"Game_AddFriendS2C":        2022,
	"Game_SendChatC2S":         2023,
	"Game_SendChatS2C":         2024,
	"Game_Receive_ChatS2C":     2026,
	"World_PingC2S":            5001,
	"World_PingS2C":            5002,
	"World_RegisterRoleIDC2S":  5003,
	"World_RegisterRoleIDS2C":  5004,
	"World_SendChatC2S":        5005,
	"World_SendChatS2C":        5006,
	"World_Receive_ChatS2C":    5008,
	"DB_UpdateLastLoginTime":   6001,
}

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}
func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (x *MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgID_value, data, "MsgID")
	if err != nil {
		return err
	}
	*x = MsgID(value)
	return nil
}
func (MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("MsgID", MsgID_name, MsgID_value)
}

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xdb, 0x8a, 0xd4, 0x30,
	0x1c, 0xc6, 0x51, 0x90, 0x4a, 0x70, 0x49, 0x37, 0xba, 0x07, 0x75, 0xbd, 0x2e, 0x78, 0xb1, 0x0b,
	0xfb, 0x06, 0xbb, 0x33, 0xab, 0x0c, 0x8c, 0x22, 0x5b, 0x17, 0x2f, 0x4b, 0x99, 0xc4, 0x1a, 0x98,
	0xa6, 0x43, 0x53, 0xc4, 0x79, 0x00, 0x05, 0x8f, 0xf8, 0x52, 0x3e, 0x80, 0xe7, 0xf3, 0xf9, 0x3c,
	0x1e, 0xc0, 0x37, 0x30, 0x5f, 0xba, 0x9d, 0x4e, 0x99, 0xf4, 0x66, 0x60, 0xbe, 0xdf, 0xf7, 0x4b,
	0xff, 0x4d, 0x13, 0xb2, 0xa6, 0xf3, 0xc1, 0xc6, 0x28, 0xcf, 0x8a, 0x4c, 0x6f, 0xa4, 0x3a, 0x91,
	0xbc, 0xfc, 0x5d, 0xb7, 0xd1, 0xe9, 0x07, 0x87, 0xc9, 0xa1, 0x73, 0x3a, 0xe9, 0x75, 0xd9, 0x22,
	0x59, 0x08, 0xc7, 0xba, 0x10, 0x69, 0xd4, 0xcf, 0x92, 0xce, 0x66, 0xe8, 0x1f, 0x68, 0x46, 0xe1,
	0x66, 0xc7, 0x3f, 0x68, 0xa2, 0x23, 0x67, 0xe3, 0x42, 0x44, 0x17, 0xa4, 0xb2, 0xa5, 0xeb, 0x5e,
	0x23, 0x42, 0xe9, 0x86, 0xc7, 0x56, 0x08, 0xdb, 0xe2, 0xa9, 0x54, 0xd1, 0x9e, 0x16, 0xb9, 0x71,
	0xa5, 0x42, 0xf7, 0xa7, 0x0b, 0xc0, 0x98, 0x78, 0xec, 0x04, 0x59, 0xaa, 0xc1, 0xae, 0x48, 0xa4,
	0x79, 0x6c, 0x0e, 0xe9, 0x57, 0x0b, 0x83, 0xf7, 0xdb, 0x63, 0xcb, 0x64, 0xb1, 0x66, 0x3b, 0xd7,
	0x64, 0x01, 0xe7, 0x8f, 0x23, 0x47, 0xff, 0xef, 0xfe, 0xb0, 0x69, 0x3d, 0xff, 0x43, 0xda, 0x88,
	0xd0, 0x7a, 0x44, 0xd9, 0x49, 0xb2, 0x6c, 0xa3, 0xea, 0x61, 0x58, 0xa4, 0xd7, 0x45, 0xff, 0x71,
	0x1b, 0x84, 0xf9, 0x64, 0x1e, 0xee, 0x66, 0x43, 0x51, 0x9a, 0x4f, 0xdb, 0x20, 0xcc, 0x67, 0x33,
	0x30, 0x56, 0x3c, 0x4b, 0x81, 0xce, 0x9b, 0xbf, 0x30, 0x9f, 0xb7, 0x41, 0x98, 0x2f, 0x28, 0x36,
	0xb5, 0x84, 0x26, 0xee, 0xe4, 0xc2, 0x7c, 0x0b, 0x58, 0x2f, 0x5d, 0x00, 0xc6, 0x2b, 0xca, 0x8e,
	0x93, 0x63, 0x53, 0xd0, 0x53, 0x97, 0xb3, 0xbe, 0x19, 0x07, 0xce, 0x6b, 0x37, 0x82, 0xf5, 0x86,
	0xb2, 0x55, 0x72, 0xd4, 0xa2, 0x1d, 0x85, 0xfd, 0x1f, 0x08, 0x25, 0x34, 0xa4, 0xb7, 0x4e, 0x02,
	0xe7, 0x5d, 0x3d, 0x82, 0xfd, 0x04, 0x53, 0xe5, 0xbd, 0x0b, 0xc0, 0xf8, 0x40, 0xd9, 0x29, 0xb2,
	0xda, 0x18, 0x60, 0x7b, 0x5c, 0xef, 0xe1, 0xc7, 0x76, 0x0c, 0xfb, 0x53, 0xbd, 0xec, 0x99, 0x5c,
	0x0a, 0xc5, 0xab, 0xf7, 0xfa, 0xec, 0x02, 0x30, 0xbe, 0x50, 0x9c, 0x14, 0x0b, 0xb6, 0x38, 0x2f,
	0x19, 0x84, 0xaf, 0x8e, 0x1c, 0xfd, 0x6f, 0x94, 0x2d, 0x11, 0xdf, 0xe6, 0x21, 0xaa, 0x57, 0x62,
	0xbb, 0xfe, 0xf7, 0xf9, 0x18, 0xed, 0x1f, 0x33, 0xdb, 0x29, 0x06, 0x42, 0x5e, 0x15, 0x51, 0x85,
	0x26, 0x94, 0x31, 0xb2, 0x70, 0x29, 0xcb, 0x87, 0x7c, 0x7a, 0x16, 0x6f, 0x06, 0xcd, 0x0c, 0xbd,
	0x5b, 0x01, 0x5b, 0x23, 0x2b, 0x65, 0x36, 0x7f, 0xa6, 0x6e, 0xb7, 0x52, 0xb8, 0x77, 0x02, 0xbc,
	0x44, 0x49, 0x67, 0xa7, 0xbd, 0xeb, 0xc8, 0xd1, 0xbf, 0x17, 0xe0, 0xaa, 0x55, 0xab, 0x35, 0xe7,
	0xbd, 0x1f, 0xe0, 0x0c, 0x76, 0xb7, 0xa3, 0xbd, 0x11, 0x37, 0xe7, 0xa8, 0x1f, 0xeb, 0xc2, 0x5e,
	0xdf, 0x8b, 0x32, 0x15, 0xfe, 0xbf, 0xf5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x1d, 0x76,
	0xc0, 0x66, 0x04, 0x00, 0x00,
}