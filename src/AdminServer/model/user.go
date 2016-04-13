package model

type User struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	Money           int32  `json:"money"`
	CreateTime      int64  `json:"create_time`
	LastLoginTime   int64  `json:"last_login_time"`
	LastLoginAreaID uint16 `json:"last_login_area_id"`

	GameServerAddr string `json:"game_server_addr"`
}

func NewUserModel() *User {
	user := new(User)
	return user
}
