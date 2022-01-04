package serve

import (
	"gitee.ltd/lxh/logger"
	"github.com/urfave/cli/v2"
	"go_scaffold/app/http"
	"go_scaffold/cli/cmd"
	"go_scaffold/config"
	"go_scaffold/global"
)

//	@method HttpServe
//	@description: http server
//	@param cli *cli.Context
//	@return error
func HttpServe(ctx *cli.Context) error {

	if err := loadConfig(ctx); err != nil {
		return err
	}

	return http.Kernel()
}

//	@method loadConfig
//	@description: load config for this server
func loadConfig(ctx *cli.Context) error {
	global.InitConfig(ctx)

	// auto migrate is open
	if config.Database.AutoMigrate {
		err := cmd.NewGormCmd().Migrate(ctx)
		if err != nil {
			logger.Say.Panicf("migrate table failed exit, %v", err.Error())
			return err
		}
	}

	return nil
}
