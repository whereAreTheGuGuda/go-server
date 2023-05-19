package vo

import (
	"go-server/model"
)

type AccountVo struct {
	Id        int64           `json:"id"`
	CreatedAt model.LocalTime `json:"createdAt"`
	UpdatedAt model.LocalTime `json:"updatedAt"`
	Nickname  string          `json:"nickname"` // 用户名
	Status    int64           `json:"status"`   // 状态1是正常,0是禁用
	Phone     string          `josn:"phone"`
	Sex       int64           `json:"sex"`
	Birthday  string          `json:"birthday"`
	Coin      int64           `json:"coin"`
	Damin     int64           `json:"damin"`
	Star      int64           `json:"star"`
	Nation    string          `json:"nation"`
	Province  string          `json:"province"`
	City      string          `json:"city"`
	Country   string          `json:"country"`
	Vip_Level int64           `json:"vip_level"`
}
