package validator_rules

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//	var ZHPhone validator.Func = func
//	@description: 国内手机号校验规则
//	@author: mrx 2021-07-12 20:38:32
//	@param fl validator.FieldLevel
//	@return bool {
var ZHPhone validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)

	rules := `^(1[3|4|5|7|8|9][0-9]\d{4,8})$`

	if ok {
		reg := regexp.MustCompile(rules)
		if reg.MatchString(value) {
			return true
		}
	}

	return false
}

//	var IsString validator.Func = func
//	@description: 判断参数类型是否为字符串
//	@author: mrx 2021-07-12 20:38:37
//	@param fl validator.FieldLevel
//	@return bool {
var IsString validator.Func = func(fl validator.FieldLevel) bool {
	switch fl.Field().Interface().(type) {
	case string:
		return true
	default:
		return false
	}
}
