package cfg

import (
	"fmt"
	"testing"
)

func TestServerConfig(t *testing.T) {
	config := GetServerConfig()
	fmt.Println("areaID: ", config.Area.ID)
	fmt.Println("config: ", config)
	fmt.Println("developMode:", config.Base.DevelopMode)
}

func TestUserConfig(t *testing.T) {
	config := GetUserConfig()
	fmt.Println("config", config)
}
