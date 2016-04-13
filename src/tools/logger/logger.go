package logger

import (
	"log"
	"os"
	"strings"
	"tools/cfg"
)

func StartLogger(path string) {
	if !strings.HasPrefix(path, "/") {
		path = os.Getenv("APRIL_PATH") + "logs/" + path

	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("cannot open logfile %v\n", err)
	}

	var r Repeater
	config := cfg.GetServerConfig()
	switch config.Base.LogOutput {
	case "terminal":
		r.out1 = os.Stdout
	case "file":
		r.out2 = file
	case "terminal+file":
		r.out1 = os.Stdout
		r.out2 = file
	}
	log.SetOutput(&r)
}
