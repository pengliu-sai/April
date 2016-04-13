package redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func startRedis(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func Test_Set(t *testing.T) {
	startRedis(t)

	v, err := RedisClient.Do("SET", "hello", "world2")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v)
}

func Test_Get(t *testing.T) {
	startRedis(t)

	v, err := redis.String(RedisClient.Do("GET", "hello"))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(v)
}

func Test_SetNX(t *testing.T) {
	startRedis(t)

	key := "aaa"

	imap := map[string]string{"key1": "111", "key2": "222"}

	value, _ := json.Marshal(imap)

	v, err := RedisClient.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}

func Test_GetNX(t *testing.T) {
	startRedis(t)

	key := "aaa"

	v, err := redis.Bytes(RedisClient.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return
	}

	var imap map[string]string

	errShal := json.Unmarshal(v, &imap)
	if errShal != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imap)
}

func Test_LPush(t *testing.T) {
	startRedis(t)

	key := "redlist"

	v, err := RedisClient.Do("LPUSH", key, "qqq")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	v, err = RedisClient.Do("LPUSH", key, "www")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func Test_LRange(t *testing.T) {
	startRedis(t)

	key := "redlist"

	value, err := redis.Values(RedisClient.Do("LRANGE", key, "0", "100"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range value {
		fmt.Println(string(v.([]byte)))
	}
}

func Test_PushDBWriteMsg(t *testing.T) {
	startRedis(t)

	//

	//PushDBWriteMsg()
}
