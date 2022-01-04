package http

import (
	"fmt"
	"gitee.ltd/lxh/logger"
	"go_scaffold/app/http/router"
	"go_scaffold/app/http/router/route"
	"go_scaffold/config"
	"net/http"
)

//	@method Kernel
//	@description: http kernel
//	@return error
func Kernel() error {

	router.IncludeRouters(route.NewUserRouter)

	handler := router.RootRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Http.Port),
		Handler: handler,
	}

	logger.Say.Infof("http server running in : %v", server.Addr)

	return server.ListenAndServe()
}
