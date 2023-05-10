package validators

import (
	"go-server/constants"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateEmail 定义验证时间格式
func ValidateEmail(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(constants.RegEmail, date)
	if !ok {
		return false
	}
	return true
}
