package model

import "database/sql"

type CommunityEntity struct {
	BaseEntity
	AccountId int           `gorm:"account_id,not null" json:"user_id"` // 用户id
	Nickname  string        `gorm:"nickname,not null" json:"nickname"`  // 用户名
	Context   string        `gorm:"context,not null" json:"context"`    // 帖子内容
	Status    sql.NullInt64 `gorm:"status,default=1" json:"status"`     // 状态1是正常, 0是禁用
}

func (t *CommunityEntity) TableName() string {
	return "community"
}
