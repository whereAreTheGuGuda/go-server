package model

import "database/sql"

type AccountEntity struct {
	BaseEntity
	Nickname string        `gorm:"nickname" json:"nickname"`                 // 用户名
	Password string        `gorm:"password,not null;unique" json:"password"` // 密码
	Phone    string        `gorm:"phone,not null;unique" json:"phone"`       //手机号
	Address  string        `gorm:"address" josn:"address"`                   //地址
	Sex      sql.NullInt16 `gorm:"sex,default=0" json:"sex"`                 //性别
	Birthday string        `gorm:"birthday" json:"birthday"`                 //生日
	Coin     sql.NullInt16 `gorm:"coin" json:"coin"`                         //金币
	Damin    sql.NullInt16 `gorm:"damin" json:"damin"`                       //钻石
	Star     sql.NullInt16 `gorm:"star" json:"star"`                         //关注
	Status   sql.NullInt64 `gorm:"status,default=1" json:"status"`           // 状态1是正常,0是禁用
}

func (t *AccountEntity) TableName() string {
	return "account"
}
