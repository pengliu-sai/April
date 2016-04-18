package game

import (
	"protos"
	. "protos/msgid"
)

func InitGameMsgID() {

	protos.RegisterMsgID(uint16(MsgID_Game_PingC2S), Game_PingC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_PingS2C), Game_PingS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RegisterUserIDC2S), Game_RegisterUserIDC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RegisterUserIDS2C), Game_RegisterUserIDS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RegisterRoleIDC2S), Game_RegisterRoleIDC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RegisterRoleIDS2C), Game_RegisterRoleIDS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RandomRoleNameC2S), Game_RandomRoleNameC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RandomRoleNameS2C), Game_RandomRoleNameS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RoleCreateC2S), Game_RoleCreateC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RoleCreateS2C), Game_RoleCreateS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RoleInfoListC2S), Game_RoleInfoListC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RoleInfoListS2C), Game_RoleInfoListS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_EnterScenesC2S), Game_EnterScenesC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_EnterScenesS2C), Game_EnterScenesS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_ExitScenesC2S), Game_ExitScenesC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_ExitScenesS2C), Game_ExitScenesS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_SendChatC2S), Game_SendChatC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_SendChatS2C), Game_SendChatS2C{})
	protos.RegisterMsgID(uint16(MsgID_Game_Receive_ChatS2C), Game_Receive_ChatS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_RoleInfoByRoleIDC2S), Game_RoleInfoByRoleIDC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_RoleInfoByRoleIDS2C), Game_RoleInfoByRoleIDS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_FriendListC2S), Game_FriendListC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_FriendListS2C), Game_FriendListS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_AddFriendC2S), Game_AddFriendC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_AddFriendS2C), Game_AddFriendS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_FBChapterListC2S), Game_FBChapterListC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_FBChapterListS2C), Game_FBChapterListS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_FBSectionListC2S), Game_FBSectionListC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_FBSectionListS2C), Game_FBSectionListS2C{})

	protos.RegisterMsgID(uint16(MsgID_Game_FBBattleC2S), Game_FBBattleC2S{})
	protos.RegisterMsgID(uint16(MsgID_Game_FBBattleS2C), Game_FBBattleS2C{})

}
