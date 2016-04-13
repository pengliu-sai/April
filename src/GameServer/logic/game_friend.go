package logic

import (
	"GameServer/dao"
	"github.com/funny/link"
	"protos"
	"protos/game"
	. "tools"
)

func gameFriendList(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive FriendListC2S message")

	roleID := g_gameSession_RoleID[gameSession]
	if roleID <= 0 {
		//未登陆角色
		return
	}

	null_send := protos.MarshalProtoMsg(&game.Game_FriendListS2C{})

	//dao
	roleList, err := dao.FriendListByRoleID(roleID)
	if roleList == nil || err != nil {
		gameSession.Send(null_send)
		return
	}

	roleListLen := len(roleList)

	roleInfoList := make([]*game.RoleInfo, roleListLen)

	for i := 0; i < roleListLen; i++ {
		role := roleList[i]
		roleInfo := game.RoleInfo{}
		roleInfo.RoleID = &role.ID
		roleInfo.RoleName = &role.Name

		sexValue := game.SexEnum_value[role.Sex]
		_sex := game.SexEnum(sexValue)
		roleInfo.Sex = &_sex

		raceValue := game.RacesEnum_value[role.Race]
		_race := game.RacesEnum(raceValue)
		roleInfo.Race = &_race

		roleInfoList[i] = &roleInfo
	}

	send_msg := protos.MarshalProtoMsg(&game.Game_FriendListS2C{
		Friends: roleInfoList,
	})

	INFO("Game send FriendListS2C message")
	gameSession.Send(send_msg)
}

func gameAddFriend(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive AddFriendC2S message")

	roleID := g_gameSession_RoleID[gameSession]
	if roleID <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_AddFriendC2S)
	friendID := *rev_msg.FriendID

	falseResult := false
	false_msg := protos.MarshalProtoMsg(&game.Game_AddFriendS2C{
		Result: &falseResult,
	})

	if friendID <= 0 {
		gameSession.Send(false_msg)
		return
	}

	//先检查该角色是否已经添加该friend
	isExit, err := dao.CheckFriend(roleID, friendID)

	if isExit || err != nil {
		ERR("friend already exist!")
		gameSession.Send(false_msg)
		return
	}

	roleFriend, err := dao.AddFriend(roleID, friendID)
	if roleFriend == nil || err != nil {
		gameSession.Send(false_msg)
		return
	}

	trueResult := true
	send_msg := protos.MarshalProtoMsg(&game.Game_AddFriendS2C{
		Result: &trueResult,
	})

	INFO("Game send AddFriendS2C message")
	gameSession.Send(send_msg)
}
