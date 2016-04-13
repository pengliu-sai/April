package db

import (
	"fmt"
	"testing"
	"time"
)

func startMysql(t *testing.T) {
	err := Connect("admin_db")
	if err != nil {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	DBOrm.SetTable("user").Delete("1=1")

	var last_login_time = time.Now().Add(-time.Minute * 20).Unix()
	for i := 1; i <= 100; i++ {
		var value = make(map[string]interface{})
		value["id"] = i
		value["name"] = fmt.Sprintf("user%d", i)
		value["money"] = int(2) // 新手送2元代金券
		value["create_time"] = time.Now().Unix()
		value["last_login_server_id"] = int(1) // 最后一次登陆的服务器
		value["last_login_time"] = last_login_time
		_, err := DBOrm.SetTable("user").Insert(value)
		if err != nil {
			t.Error(err)
		}
	}
}

//func TestDel(t *testing.T) {
//	startMysql(t)
//	defer DBOrm.DbClose()
//	DBOrm.SetTable("user").Delete("1=1")
//}

func TestDelByCondition(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	_, err := DBOrm.SetTable("user").Delete("id = 99")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	var value = make(map[string]interface{})
	value["name"] = fmt.Sprintf("update_user%d", 98)
	_, err := DBOrm.SetTable("user").Where("name = 'user98'").Update(value)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchAll(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	data := DBOrm.SetTable("user").FindAll()
	Print(data)
}

func TestSearchAllByCondition(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	data := DBOrm.SetTable("user").Where("name = 'user70'").FindAll()
	Print(data)
}

func TestSearchOne(t *testing.T) {
	startMysql(t)
	defer DBOrm.DbClose()
	data := DBOrm.SetTable("user").FindOne()
	Print(data)
}
