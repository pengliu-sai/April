package model

type RoleFriend struct {
	ID         uint64 `json:"id"`
	RoleID     uint64 `json:"role_id"`
	FriendID   uint64 `json:"friend_id"`
	CreateTime int64  `json:"create_time`
}

func NewRoleFriendModel() *RoleFriend {
	roleFriend := new(RoleFriend)
	return roleFriend
}
