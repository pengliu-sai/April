package logic

import (
	"GameServer/dao"
	"github.com/funny/link"
	"protos"
	"protos/game"
	. "tools"
	"tools/random"
)

func gameRandomRoleName(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive randomRoleNameC2S message")
	userID := g_gameSession_UserID[gameSession]
	if userID <= 0 {
		//未登陆用户
		ERR("user canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_RandomRoleNameC2S)
	//根据性别 生成随机名
	INFO("rev_msg sex: ", *rev_msg.Sex)
	//测试数据
	var names = [...]string{"杰伦", "叮当", "加菲"}
	index := random.RandomInt31n(3)

	roleName := names[index]

	send_msg := protos.MarshalProtoMsg(&game.Game_RandomRoleNameS2C{
		Name: protos.String(roleName),
	})

	INFO("Game send randomRoleNameS2C message")
	gameSession.Send(send_msg)
}

func gameRoleCreate(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive roleCreateC2S message")
	userID := g_gameSession_UserID[gameSession]
	if userID <= 0 {
		//未登陆用户
		ERR("user canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_RoleCreateC2S)

	name := *rev_msg.Name
	sex := *rev_msg.Sex
	race := *rev_msg.Race

	null_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateS2C{})

	//存入DB, 写入role -> userID
	role, err := dao.RoleCreate(name, sex.String(), race.String(), userID)
	if role == nil || err != nil {
		gameSession.Send(null_msg)
		return
	}

	sexValue := game.SexEnum_value[role.Sex]
	_sex := game.SexEnum(sexValue)

	raceValue := game.RacesEnum_value[role.Race]
	_race := game.RacesEnum(raceValue)

	send_msg := protos.MarshalProtoMsg(&game.Game_RoleCreateS2C{
		RoleID:   protos.Uint64(role.ID),
		RoleName: protos.String(role.Name),
		Sex:      &_sex,
		Race:     &_race,
	})

	INFO("Game send roleCreateS2C message")
	gameSession.Send(send_msg)
}

func gameRoleInfoList(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive roleInfoListC2S message")

	userID := g_gameSession_UserID[gameSession]
	if userID <= 0 {
		//未登陆用户
		ERR("user canot login")
		return
	}

	null_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListS2C{})

	//从db读取roles
	roleList, err := dao.RoleListByUserID(userID)
	if roleList == nil || err != nil {
		gameSession.Send(null_msg)
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

	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoListS2C{
		Roles: roleInfoList,
	})

	g_gameSession_RoleInfoList[gameSession] = roleInfoList

	INFO("Game send roleInfoListS2C message")
	gameSession.Send(send_msg)
}

func gameRoleInfoByRoleID(gameSession *link.Session, msg protos.ProtoMsg) {
	INFO("Game receive roleInfoByRoleIDC2S message")

	if g_gameSession_RoleID[gameSession] <= 0 {
		//未登陆角色
		ERR("role canot login")
		return
	}

	rev_msg := msg.Body.(*game.Game_RoleInfoByRoleIDC2S)
	roleID := *rev_msg.RoleID

	null_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoByRoleIDS2C{})
	if roleID <= 0 {
		gameSession.Send(null_msg)
		return
	}

	//dao查询..
	role, err := dao.RoleInfoByRoleID(roleID)
	if role == nil || err != nil {
		gameSession.Send(null_msg)
		return
	}

	roleInfo := game.RoleInfo{}
	roleInfo.RoleID = &role.ID
	roleInfo.RoleName = &role.Name

	sexValue := game.SexEnum_value[role.Sex]
	_sex := game.SexEnum(sexValue)
	roleInfo.Sex = &_sex

	raceValue := game.SexEnum_value[role.Race]
	_race := game.RacesEnum(raceValue)
	roleInfo.Race = &_race

	send_msg := protos.MarshalProtoMsg(&game.Game_RoleInfoByRoleIDS2C{
		Role: &roleInfo,
	})

	gameSession.Send(send_msg)
	INFO("Game send roleInfoByRoleIDS2C message")
}
