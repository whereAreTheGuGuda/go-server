package model

import "database/sql"

type AccountEntity struct {
	BaseEntity
	Nickname  string        `gorm:"nickname" json:"nickname"`                 // 用户名
	Password  string        `gorm:"password,not null;unique" json:"password"` // 密码
	Phone     string        `gorm:"phone,not null;unique" json:"phone"`       //手机号
	Sex       sql.NullInt64 `gorm:"sex,default=0" json:"sex"`                 //性别
	Birthday  string        `gorm:"birthday" json:"birthday"`                 //生日
	Coin      sql.NullInt64 `gorm:"coin" json:"coin"`                         //金币
	Damin     sql.NullInt64 `gorm:"damin" json:"damin"`                       //钻石
	Star      sql.NullInt64 `gorm:"star" json:"star"`                         //关注
	Status    sql.NullInt64 `gorm:"status,default=1" json:"status"`           // 状态1是正常,0是禁用
	Nation    string        `gorm:"naton,detault=\"中国\"" json:"nation"`       //国家
	Province  string        `gorm:"province,default=\"北京\"" json:"province"`  //省
	City      string        `gorm:"city,default=\"北京\"" json:"city"`          //市
	Country   string        `gorm:"country default=\"朝阳区\"" json:"country"`   //县
	Vip_Level sql.NullInt64 `grom:"vip_level,default=0" json:"vip_level"`     //会员等级
}

func (t *AccountEntity) TableName() string {
	return "account"
}
