package logic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"tools/cfg"
	. "tools/unitest"
)

var (
	adminAddr             string
	global_liupeng_userID = "847315546201395200"
)

func init() {
	adminConf := cfg.GetServerConfig().AdminServer
	adminAddr = adminConf.IP + ":" + adminConf.Port
}

/*-----------[[ gate ]]---------------------*/
//测试连接
func TestGateConnect(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/?ping=hi", adminAddr)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	Equal(t, string(body), "pong")

}

//测试注册
func TesGateUserRegister(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userRegister?name=%s&password=%s", adminAddr, "liupeng", "123456")
	req, _ := http.NewRequest("POST", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

//测试登陆
func TestGateUserLoginByUserName(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userLoginByUserName?name=%s&password=%s", adminAddr, "liupeng", "123456")
	req, _ := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

//登陆by name + password + userid
func TestGateUserLoginByUserID(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userLoginByUserID?name=%s&password=%s&userid=%s", adminAddr, "liupeng", "123456", "915441286658396160")
	req, _ := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

//用户退出
func TestGateUserExit(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/gate/userExit?userid=%s", adminAddr, global_liupeng_userID)
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	Equal(t, string(body), "SUCCESS")
	t.Log("body:", string(body))
}

/*-----------[[ admin ]]----------------*/
//标记删除
func TestAdminUserMarkDeleted(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/admin/userMarkDeleted?userid=%s", adminAddr, global_liupeng_userID)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

//用户恢复
func TestAdminUserRebirth(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/admin/userRebirth?userid=%s", adminAddr, global_liupeng_userID)
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

//删除用户
func TestAdminUserDelete(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/admin/userDelete?name=%s&password=%s", adminAddr, "liupeng", "123456")
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	t.Log("body:", string(body))
}

/*-----------[[ game ]]---------------------*/

//用户登出
func TestGameUserOffline(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/game/userOffline?userid=%s", adminAddr, global_liupeng_userID)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	Equal(t, string(body), "SUCCESS")
}

//gameServer下线
func TestGameServerOffline(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/game/gameServerOffline?gameserveraddr=%s", adminAddr, "127.0.0.1:8888")
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	Equal(t, string(body), "SUCCESS")
}

//测试跨区登陆通知用户下线
func TestOtherAdminUserOffline(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/otheradmin/userOffline?userid=%s", adminAddr, global_liupeng_userID)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	Equal(t, err, nil)
	Equal(t, resp.StatusCode, 200)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	Equal(t, string(body), "SUCCESS")
}
