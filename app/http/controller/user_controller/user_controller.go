package user_controller

import (
	"github.com/gin-gonic/gin"
	"go_scaffold/app/http/core/response"
	"go_scaffold/app/http/request/user_request"
	"go_scaffold/app/service/user_service"
)

type userController struct {
	response *response.Response
}

type UserController interface {
	Add(c *gin.Context)
}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) Add(c *gin.Context) {
	request := new(user_request.UserAddRequest)

	if err := c.ShouldBind(request); err != nil {
		u.response.ValidForm(c, request.First(err))
		return
	}

	service := user_service.NewUserService()

	if err := service.AddUser(request); err != nil {
		u.response.Failed(c, err.Error())
		return
	}

	u.response.Message(c, "success")
	return
}
