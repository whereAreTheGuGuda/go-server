package vo

type CommunityListVo struct {
	Id        int    `json:"id"`
	AccountId int    `json:"account_id"`
	Nickname  string `json:"nickname"` // 用户名
	Status    int64  `json:"status"`   // 状态1是正常,0是禁用
	Context   string `json:"context"`
}
