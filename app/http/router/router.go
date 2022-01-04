package router

import (
	"github.com/cowardmrx/gin_cors"
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

var options []Option

//	@method IncludeRouters
//	@description: include other module router
//	@param opts ...Option
func IncludeRouters(opts ...Option) {
	options = append(opts)
}

//	@method RootRouter
//	@description: root router
//	@return *gin.Engine
func RootRouter() *gin.Engine {
	r := gin.New()

	// forward client IP addr
	r.ForwardedByClientIP = true

	// 跨域设置
	cors := &gin_cors.Cors{
		AccessControlAllowOrigins: []string{
			"http://localhost",
		},
		AccessControlAllowMethods: "GET,POST,UPDATE,PUT,DELETE",
	}

	// 跨域中间件
	r.Use(cors.CorsMiddleware())

	for _, opt := range options {
		opt(r)
	}

	return r
}
