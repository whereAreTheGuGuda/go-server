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
