package dto

type AccountDto struct {
	Phone    string `json:"phone" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type CreateAccountDto struct {
	AccountDto
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6,max=16"`
}

type ModifyAccountPassword struct {
	Password        string `json:"password" binding:"required,min=6,max=16"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6,max=16"`
}

type EditAccountInfo struct {
	Coin     int64  `json:"coin" binding:"omitempty"`
	Nation   string `json:"nation" binding:"omitempty"`   //国家
	Province string `json:"province" binding:"omitempty"` //省
	City     string `json:"city" binding:"omitempty"`     //市
	Country  string `json:"country" binding:"omitempty"`  //县
	Birthday string `json:"birthday" binding:"omitempty"` //生日
}
