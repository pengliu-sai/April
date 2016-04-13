package dao

import (
	. "GameServer/model"
	"errors"
	"strconv"
	"time"
	"tools/cfg"
	"tools/db"
	"tools/guid"
)

var friendGuid *guid.Guid = guid.NewGuid()

func AddFriend(roleID uint64, friendID uint64) (*RoleFriend, error) {
	return insertFriend(friendGuid.NewID(cfg.GetServerConfig().Area.ID), roleID, friendID)

}

func insertFriend(id uint64, roleID uint64, friendID uint64) (*RoleFriend, error) {
	nowTime := time.Now().Unix()

	var value = make(map[string]interface{})
	value["id"] = id
	value["role_id"] = roleID
	value["friend_id"] = friendID
	value["create_time"] = nowTime

	_, err := db.DBOrm.SetTable("role_friend").Insert(value)
	if err != nil {
		return nil, err
	}

	model := NewRoleFriendModel()
	model.ID = id
	model.RoleID = roleID
	model.FriendID = friendID
	model.CreateTime = nowTime

	return model, nil
}

//inner查询出friends info信息
func FriendListByRoleID(roleID uint64) ([]*Role, error) {
	roleIDStr := strconv.FormatUint(roleID, 10)

	data := db.DBOrm.SetTable("role_friend").Fileds("role.id", "role.name", "role.sex", "role.race", "role.user_id", "role.create_time").Where("role_friend.role_id = '"+roleIDStr+"'").Join("role", "role_friend.friend_id = role.id").FindAll()
	if data == nil {
		return nil, errors.New("sql role_friend select fail")
	}

	dataLen := len(data)

	roleList := make([]*Role, dataLen)
	for i := 0; i < dataLen; i++ {
		roleList[i] = fullRoleModel(data[i])
	}

	return roleList, nil
}

func CheckFriend(roleID uint64, friendID uint64) (bool, error) {
	isExist := false

	roleIDStr := strconv.FormatUint(roleID, 10)
	friendIDStr := strconv.FormatUint(friendID, 10)

	data := db.DBOrm.SetTable("role_friend").Where("role_id = '" + roleIDStr + "' and friend_id = '" + friendIDStr + "'").FindOne()
	if data == nil {
		return isExist, errors.New("sql role_friend select fail")
	}

	dataLen := len(data)

	if dataLen > 0 {
		isExist = true
	}

	return isExist, nil
}
