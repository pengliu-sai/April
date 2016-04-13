package redis

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
	. "tools"
	"tools/cfg"
)

var RedisClient redis.Conn

func Connect(redisName string) error {
	redisListConf := cfg.GetServerConfig().Redis
	var find_db_index = -1
	for i := 0; i < len(redisListConf); i++ {
		if redisListConf[i].Name == redisName {
			find_db_index = i
			break
		}
	}

	if find_db_index < 0 {
		return errors.New("Can't find redis config")
	}
	redisConf := redisListConf[find_db_index]

	redisAddr := redisConf.IP + ":" + redisConf.Port
	conn, err := redis.DialTimeout("tcp", redisAddr, 0, 1*time.Second, 1*time.Second)
	conn.Send("AUTH", "liupeng")
	if err != nil {
		return err
	}

	RedisClient = conn
	return nil
}

const (
	DB_Write_Msgs = "DB_Write_Msgs"
)

func PushDBWriteMsg(msg []byte) {
	_, err := RedisClient.Do("RPUSH", DB_Write_Msgs, msg)
	if err != nil {
		ERR("PushDBWriteMsg: ", err)
	}
}

func PullDBWriteMsg() [][]byte {
	//datas, err := RedisClient.Do("LRANGE", DB_Write_Msgs, 0, -1)
	//if err != nil {
	//	ERR("PullDBWriteMsg: ", err)
	//}
	//RedisClient.Do("LTRIM", DB_Write_Msgs, len(datas), -1)
	return nil
}
