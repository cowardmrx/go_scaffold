package request

import (
	"github.com/go-playground/validator/v10"
	"go_scaffold/global"
)

type Request struct{}

type RequestFunc interface {
	First() string
	Array() []string
}

//	@method: First
//	@description: 获取单一错误信息
//	@param: err error
//	@return: string
func (request *Request) First(err error) string {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	errMap := errs.Translate(global.ValidatorTrans)

	var errMsg string
	for _, v := range errMap {
		errMsg = v
	}

	return errMsg
}

//	@method: Array
//	@description: 获取全部错误信息
//	@param: err error
//	@return: []string
func (request *Request) Array(err error) []string {
	var errors []string

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return append(errors, err.Error())
	}

	errMap := errs.Translate(global.ValidatorTrans)

	for _, v := range errMap {
		errors = append(errors, v)
	}

	return errors
}
