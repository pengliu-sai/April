package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type ServerConfig struct {
	Area struct {
		ID   uint16 `json:"id"`
		Name string `json:"name"`
		Desc string `json:"desc"`
	} `json:"area"`
	GateServer struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Desc    string `json:"desc"`
		LogFile string `json:"log_file"`
	} `json:"gate_server"`
	AdminServer struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Desc    string `json:"desc"`
		LogFile string `json:"log_file"`
	} `json:"admin_server"`
	ChatServer struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Desc    string `json:"desc"`
		LogFile string `json:"log_file"`
	} `json:"chat_server"`
	WorldServer struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Desc    string `json:"desc"`
		LogFile string `json:"log_file"`
	} `json:"world_server"`
	LogServer struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		LogFile string `json:"log_file"`
	} `json:"log_server"`
	GameServerList []struct {
		Name    string `json:"name"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Desc    string `json:"desc"`
		LogFile string `json:"log_file"`
	} `json:"game_server_list"`
	Redis []struct {
		Name     string `json:"name"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		Password string `json:"password"`
		Desc     string `json:"desc"`
	} `json:"redis"`
	Mysql []struct {
		Name     string `json:"name"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		Charset  string `json:"charset"`
		Desc     string `json:"desc"`
	} `json:"mysql"`
	Base struct {
		Debug       bool   `json:"debug"`
		LogOutput   string `json:"log_output"`
		Desc        string `json:"desc"`
		DevelopMode bool   `json:"develop_mode"`
	} `json:"base"`
}

type UserConfig struct {
	WhiteUserList []struct {
		Name string `json:"name"`
	} `json:"white_user_list"`
	BlackUserList []struct {
		Name string `json:"name"`
	} `json:"black_user_list"`
}

var (
	_serverConfig ServerConfig
	_userConfig   UserConfig
	_lock         sync.Mutex
)

func init() {
	Reload()
}

func Reload() {
	serverConfigPath := os.Getenv("APRIL_PATH") + "data/server/config_server.json"
	userConfigPath := os.Getenv("APRIL_PATH") + "data/server/config_user.json"
	_lock.Lock()
	_load_server_config(serverConfigPath)
	_load_user_config(userConfigPath)
	_lock.Unlock()
}

func _load_server_config(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(path, err)
		return
	}

	jsonStr, err := ioutil.ReadAll(f)
	if err != nil {
	}

	err = json.Unmarshal(jsonStr, &_serverConfig)
	if err != nil {
		log.Println("config_server.json Unmarshal err", err)
	}
}

func _load_user_config(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(path, err)
		return
	}

	jsonStr, err := ioutil.ReadAll(f)
	if err != nil {
	}

	err = json.Unmarshal(jsonStr, &_userConfig)
	if err != nil {
		log.Println("config_user.json Unmarshal err", err)
	}
}

func GetServerConfig() *ServerConfig {
	_lock.Lock()
	defer _lock.Unlock()
	return &_serverConfig
}

func GetUserConfig() *UserConfig {
	_lock.Lock()
	defer _lock.Unlock()
	return &_userConfig
}

func IsDevelopMode() bool {
	return GetServerConfig().Base.DevelopMode
}
