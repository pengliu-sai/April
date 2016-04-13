package model

type Area struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Port   string `json:"port"`
	AreaID uint16 `json:"area_id"`
}

func NewAreaModel() *Area {
	admin := new(Area)
	return admin
}
