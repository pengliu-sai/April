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

var roleGuid *guid.Guid = guid.NewGuid()

func RoleCreate(name string, sex string, race string, userID uint64) (*Role, error) {
	return insertRole(roleGuid.NewID(cfg.GetServerConfig().Area.ID), name, sex, race, userID)
}

func insertRole(id uint64, name string, sex string, race string, userID uint64) (*Role, error) {

	nowTime := time.Now().Unix()

	var value = make(map[string]interface{})
	value["id"] = id
	value["name"] = name
	value["sex"] = sex
	value["race"] = race
	value["user_id"] = userID
	value["create_time"] = nowTime

	_, err := db.DBOrm.SetTable("role").Insert(value)
	if err != nil {
		return nil, err
	}

	model := NewRoleModel()
	model.ID = id
	model.Name = name
	model.Sex = sex
	model.Race = race
	model.CreateTime = nowTime

	return model, nil
}

func RoleListByUserID(userID uint64) ([]*Role, error) {
	userIDStr := strconv.FormatUint(userID, 10)

	data := db.DBOrm.SetTable("role").Where("user_id = '" + userIDStr + "'").FindAll()
	if data == nil {
		return nil, errors.New("sql role select fail")
	}

	dataLen := len(data)

	roleList := make([]*Role, dataLen)
	for i := 0; i < dataLen; i++ {
		roleList[i] = fullRoleModel(data[i])
	}

	return roleList, nil
}

func fullRoleModel(data map[string]string) *Role {
	model := NewRoleModel()

	id, _ := strconv.ParseUint(data["id"], 10, 64)

	model.ID = id

	model.Name = data["name"]
	model.Sex = data["sex"]
	model.Race = data["race"]

	userID, _ := strconv.ParseUint(data["user_id"], 10, 64)
	model.UserID = userID

	return model
}

func RoleInfoByRoleID(roleID uint64) (*Role, error) {
	roleIDStr := strconv.FormatUint(roleID, 10)
	data := db.DBOrm.SetTable("role").Where("id = '" + roleIDStr + "'").FindOne()
	if data == nil {
		return nil, errors.New("sql user select fail")
	}

	if len(data) != 0 {
		return fullRoleModel(data[0]), nil
	}
	return nil, nil
}
