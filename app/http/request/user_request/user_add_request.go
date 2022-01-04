package user_request

import "go_scaffold/app/http/request"

type UserAddRequest struct {
	request.Request
	Name string `json:"name" form:"name" binding:"required" label:"姓名"`
	Age  int    `json:"age" form:"age" binding:"required,min=10" label:"年龄"`
}
