package route

import (
	"github.com/gin-gonic/gin"
	"go_scaffold/app/http/controller/user_controller"
)

func NewUserRouter(r *gin.Engine) {
	userHandler := user_controller.NewUserController()

	route := r.Group("user")
	{
		route.POST("add", userHandler.Add)
	}
}
