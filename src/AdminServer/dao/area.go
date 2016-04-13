package dao

import (
	. "AdminServer/model"
	"errors"
	"strconv"
	"tools/db"
)

func AreaInfoByAreaID(areaID uint16) (*Area, error) {
	areaIDStr := strconv.FormatUint(uint64(areaID), 10)
	data := db.DBOrm.SetTable("area").Where("area_id = '" + areaIDStr + "'").FindOne()
	if data == nil {
		return nil, errors.New("sql area select fail")
	}

	if len(data) != 0 {
		return fullAreaModel(data[0]), nil
	}
	return nil, nil
}

func fullAreaModel(data map[string]string) *Area {
	model := NewAreaModel()

	id, _ := strconv.ParseUint(data["id"], 10, 64)

	model.ID = id

	model.Name = data["name"]

	model.IP = data["ip"]

	model.Port = data["port"]

	return model
}
