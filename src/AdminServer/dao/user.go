package dao

import (
	. "AdminServer/model"
	"errors"
	"strconv"
	"time"
	"tools/cfg"
	"tools/db"
	"tools/guid"
)

var userGuid *guid.Guid = guid.NewGuid()

func UserLoginByName(name string, password string) (*User, error) {
	data := db.DBOrm.SetTable("user").Where("name = '" + name + "' and password = '" + password + "' and is_deleted is NULL").FindOne()
	if data == nil {
		return nil, errors.New("sql user select fail")
	}

	if len(data) != 0 {
		return fullUserModel(data[0]), nil
	}
	return nil, nil
}

func UserRegister(name string, password string) (*User, error) {
	return insertUser(userGuid.NewID(cfg.GetServerConfig().Area.ID), name, password)
}

func UserDelete(name string, password string) (int, error) {
	num, err := db.DBOrm.SetTable("user").Delete("name = '" + name + "' and password = '" + password + "'")
	if err != nil {
		return 0, err
	}
	return num, nil
}

func UserMarkDeleted(userID string) (int, error) {
	var value = make(map[string]interface{})
	value["is_deleted"] = 1
	num, err := db.DBOrm.SetTable("user").Where("id = '" + userID + "'").Update(value)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func UserRebirth(userID string) (int, error) {
	num, err := db.DBOrm.SetTable("user").Where("id = '" + userID + "'").SetNull("is_deleted")
	if err != nil {
		return 0, err
	}
	return num, nil
}

func UserInfo(name string, password string) (*User, error) {
	data := db.DBOrm.SetTable("user").Where("name = '" + name + "' and password = '" + password + "'").FindOne()
	if data == nil {
		return nil, errors.New("sql user select fail")
	}

	if len(data) != 0 {
		return fullUserModel(data[1]), nil
	}
	return nil, nil
}

func fullUserModel(data map[string]string) *User {
	model := NewUserModel()

	id, _ := strconv.ParseUint(data["id"], 10, 64)

	model.ID = id

	model.Name = data["name"]

	money, _ := strconv.Atoi(data["money"])
	model.Money = int32(money)

	create_time, _ := strconv.ParseInt(data["create_time"], 10, 64)
	model.CreateTime = create_time

	last_login_time, _ := strconv.ParseInt(data["last_login_time"], 10, 64)
	model.LastLoginTime = last_login_time

	lastLoginAreaID, _ := strconv.Atoi(data["last_login_area_id"])

	model.LastLoginAreaID = uint16(lastLoginAreaID)

	return model
}

func insertUser(id uint64, name string, password string) (*User, error) {
	addMoney := 2
	nowTime := time.Now().Unix()
	var value = make(map[string]interface{})
	value["id"] = id
	value["name"] = name
	value["password"] = password
	value["money"] = addMoney
	value["create_time"] = nowTime
	value["last_login_time"] = nowTime
	value["last_login_area_id"] = 0

	_, err := db.DBOrm.SetTable("user").Insert(value)
	if err != nil {
		return nil, err
	}

	model := NewUserModel()
	model.ID = id
	model.Name = name
	model.Money = int32(addMoney)
	model.CreateTime = nowTime
	model.LastLoginTime = nowTime
	model.LastLoginAreaID = 0

	return model, nil
}
