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
	MsgID_Game_SendChatC2S         MsgID = 2023
	MsgID_Game_SendChatS2C         MsgID = 2024
	MsgID_Game_Receive_ChatS2C     MsgID = 2026
	MsgID_Game_FBChapterListC2S    MsgID = 2027
	MsgID_Game_FBChapterListS2C    MsgID = 2028
	MsgID_Game_FBSectionListC2S    MsgID = 2029
	MsgID_Game_FBSectionListS2C    MsgID = 2030
	MsgID_Game_FBBattleC2S         MsgID = 2031
	MsgID_Game_FBBattleS2C         MsgID = 2032
	// -----[[world 5001-6000]]------
	MsgID_World_PingC2S           MsgID = 5001
	MsgID_World_PingS2C           MsgID = 5002
	MsgID_World_RegisterRoleIDC2S MsgID = 5003
	MsgID_World_RegisterRoleIDS2C MsgID = 5004
	MsgID_World_SendChatC2S       MsgID = 5005
	MsgID_World_SendChatS2C       MsgID = 5006
	MsgID_World_Receive_ChatS2C   MsgID = 5008
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
	2027: "Game_FBChapterListC2S",
	2028: "Game_FBChapterListS2C",
	2029: "Game_FBSectionListC2S",
	2030: "Game_FBSectionListS2C",
	2031: "Game_FBBattleC2S",
	2032: "Game_FBBattleS2C",
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
	"Game_FBChapterListC2S":    2027,
	"Game_FBChapterListS2C":    2028,
	"Game_FBSectionListC2S":    2029,
	"Game_FBSectionListS2C":    2030,
	"Game_FBBattleC2S":         2031,
	"Game_FBBattleS2C":         2032,
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
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0xdb, 0x6a, 0xd4, 0x40,
	0x18, 0xc7, 0x51, 0x90, 0xc5, 0xc1, 0x32, 0xdb, 0xd1, 0x1e, 0xd4, 0x7a, 0x1d, 0xf0, 0xa2, 0x85,
	0xbe, 0x41, 0xb3, 0xdb, 0xca, 0xc2, 0x2a, 0xd2, 0x58, 0xbc, 0x0c, 0x61, 0x33, 0xae, 0x03, 0x9b,
	0xc9, 0x92, 0x0c, 0x62, 0x1f, 0x40, 0xc1, 0x23, 0xbe, 0x9a, 0xe7, 0xf3, 0xb1, 0x1e, 0x5b, 0xd7,
	0xc3, 0x1b, 0x38, 0xff, 0x89, 0xd9, 0x6c, 0xc8, 0xcc, 0xcd, 0xc2, 0xfe, 0x7f, 0xdf, 0x6f, 0xf2,
	0xe5, 0xcb, 0x97, 0x90, 0x95, 0x3c, 0x1b, 0xac, 0x8d, 0xb3, 0x54, 0xa5, 0xf9, 0x5a, 0x92, 0x0f,
	0x45, 0x5c, 0xfc, 0xae, 0x9a, 0xe8, 0xec, 0xde, 0x51, 0x72, 0xe4, 0x7c, 0x3e, 0xec, 0x75, 0xd9,
	0x3c, 0x99, 0x0b, 0x76, 0x73, 0xc5, 0x93, 0xb0, 0x9f, 0x0e, 0x3b, 0xeb, 0x41, 0xfb, 0x50, 0x3d,
	0x0a, 0xd6, 0x3b, 0xed, 0xc3, 0x3a, 0x3a, 0x76, 0x2e, 0x52, 0x3c, 0xbc, 0x28, 0xa4, 0x29, 0xba,
	0xd1, 0xaa, 0x45, 0x28, 0xba, 0xd9, 0x62, 0x4b, 0x84, 0x6d, 0xc4, 0x89, 0x90, 0xe1, 0x4e, 0xce,
	0x33, 0xed, 0x0a, 0x89, 0xda, 0x1f, 0x36, 0x00, 0x63, 0xbf, 0xc5, 0x4e, 0x91, 0x85, 0x0a, 0x6c,
	0xf3, 0xa1, 0xd0, 0x97, 0xcd, 0x20, 0x1d, 0x38, 0x18, 0xbc, 0x9f, 0x2d, 0xb6, 0x48, 0xe6, 0x2b,
	0xb6, 0x79, 0x5d, 0x28, 0x38, 0x13, 0x4b, 0x8e, 0xfa, 0x5f, 0xff, 0x9b, 0x4d, 0xaa, 0xfe, 0x1f,
	0xd2, 0x5a, 0x84, 0xaa, 0x47, 0x94, 0x9d, 0x26, 0x8b, 0x26, 0x2a, 0x2f, 0x86, 0x43, 0x7a, 0x5d,
	0xd4, 0x3f, 0x76, 0x41, 0x98, 0x4f, 0x9a, 0x70, 0x3b, 0x1d, 0xf1, 0xc2, 0x7c, 0xea, 0x82, 0x30,
	0x9f, 0xcd, 0xc0, 0x48, 0xc6, 0x69, 0x02, 0x74, 0x41, 0xff, 0x85, 0xf9, 0xdc, 0x05, 0x61, 0xbe,
	0xa0, 0x18, 0x6a, 0x01, 0x75, 0xdc, 0xc9, 0xb8, 0x7e, 0x16, 0xb0, 0x5e, 0xda, 0x00, 0x8c, 0x57,
	0x94, 0x9d, 0x24, 0x27, 0xa6, 0xa0, 0x27, 0xaf, 0xa4, 0x7d, 0xdd, 0x0e, 0x9c, 0xd7, 0x76, 0x04,
	0xeb, 0x0d, 0x65, 0xcb, 0xe4, 0xb8, 0x41, 0x9b, 0x12, 0xf3, 0x1f, 0x70, 0xc9, 0x73, 0x48, 0x6f,
	0xad, 0x04, 0xce, 0xbb, 0xaa, 0x05, 0xf3, 0x08, 0xa6, 0xca, 0x7b, 0x1b, 0x80, 0xf1, 0x81, 0xb2,
	0x33, 0x64, 0xb9, 0xd6, 0x80, 0xbf, 0x5b, 0xcd, 0xf0, 0xa3, 0x1b, 0xc3, 0xfe, 0x54, 0x1d, 0xbb,
	0x95, 0x09, 0x2e, 0xe3, 0xf2, 0xbe, 0xf6, 0x6c, 0x00, 0xc6, 0x67, 0x8a, 0x4d, 0x31, 0x60, 0x23,
	0x8e, 0x0b, 0x06, 0xe1, 0x8b, 0x25, 0x47, 0xfd, 0x57, 0xca, 0x16, 0x48, 0xdb, 0xe4, 0x01, 0x4a,
	0xaf, 0x46, 0xe6, 0xfc, 0x6f, 0xcd, 0x18, 0xd5, 0xdf, 0x67, 0xc6, 0xc9, 0x07, 0x5c, 0x5c, 0xe3,
	0x61, 0x89, 0xf6, 0x29, 0xd6, 0xba, 0xe8, 0xc8, 0xd7, 0xe1, 0x58, 0x0f, 0xae, 0xec, 0xf6, 0xc0,
	0xc1, 0xcc, 0xca, 0xcf, 0xb2, 0x80, 0x0f, 0x94, 0x48, 0x65, 0xe9, 0x4d, 0x1c, 0xcc, 0xac, 0x7e,
	0xd5, 0xe1, 0x96, 0xef, 0x47, 0x4a, 0x8d, 0xcc, 0x92, 0xfc, 0x6e, 0xc6, 0xa8, 0xfe, 0x43, 0x19,
	0x23, 0x73, 0x97, 0xd3, 0x6c, 0x14, 0x4f, 0xdf, 0x94, 0x5b, 0x5e, 0x3d, 0x43, 0xdd, 0x6d, 0x8f,
	0xad, 0x90, 0xa5, 0x22, 0x6b, 0x6e, 0xfc, 0x1d, 0x27, 0x85, 0x7b, 0xd7, 0xc3, 0x88, 0x0b, 0x3a,
	0x3b, 0xcb, 0x7b, 0x96, 0x1c, 0xf5, 0xf7, 0x3d, 0xdc, 0x5d, 0x79, 0x5a, 0x7d, 0x9a, 0x0f, 0x3c,
	0xbc, 0x21, 0x5d, 0x3f, 0xdc, 0x19, 0xc7, 0x7a, 0xcb, 0xfb, 0x51, 0xae, 0xcc, 0xc7, 0xe5, 0x92,
	0x48, 0x78, 0xfb, 0xef, 0xea, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x00, 0x8b, 0xd0, 0x04,
	0x05, 0x00, 0x00,
}
