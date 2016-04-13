package model

type Role struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Race       string `json:"race"`
	UserID     uint64 `json:"user_id"`
	CreateTime int64  `json:"create_time`
}

func NewRoleModel() *Role {
	role := new(Role)
	return role
}
