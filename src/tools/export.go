package tools

import (
	"fmt"
	"log"
	"strings"
	"tools/cfg"
	"tools/logger"
)

var (
	debug_open bool
)

func init() {
	if cfg.GetServerConfig().Base.Debug {
		debug_open = true
	}
}

//------------------------------------------------ 严重程度由高到低
func ERR(v ...interface{}) {
	log.Printf("\033[1;4;31m[ERROR] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func WARN(v ...interface{}) {
	log.Printf("\033[1;33m[WARN] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func INFO(v ...interface{}) {
	log.Printf("\033[32m[INFO] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func INFO_F(format string, v ...interface{}) {
	log.Printf("\033[32m[INFO] %v \033[0m\n", strings.TrimRight(fmt.Sprintf(format, v...), "\n"))
}

func NOTICE(v ...interface{}) {
	log.Printf("[NOTICE] %v\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func DEBUG(v ...interface{}) {
	if debug_open {
		log.Printf("\033[1;35m[DEBUG] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
	}
}

func SetGateLogFile() {
	logger.StartLogger(cfg.GetServerConfig().GateServer.LogFile)
	log.SetPrefix("[GateServer]")
}

func SetAdminLogFile() {
	logger.StartLogger(cfg.GetServerConfig().AdminServer.LogFile)
	log.SetPrefix("[AdminServer]")
}

func SetGameLogFile(index int) {
	logger.StartLogger(cfg.GetServerConfig().GameServerList[index].LogFile)
	log.SetPrefix("[GameServer-" + string(index) + "]")
}

func SetWorldLogFile() {
	logger.StartLogger(cfg.GetServerConfig().WorldServer.LogFile)
	log.SetPrefix("[WorldServer]")
}

func SetChatLogFile() {
	logger.StartLogger(cfg.GetServerConfig().ChatServer.LogFile)
	log.SetPrefix("[ChatServer]")
}
