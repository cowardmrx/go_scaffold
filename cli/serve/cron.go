package serve

import (
	"github.com/urfave/cli/v2"
	"go_scaffold/app/cron"
)

//	@method Cron
//	@description: cron serve
//	@param ctx *cli.Context
//	@return error
func Cron(ctx *cli.Context) error {

	if err := loadConfig(ctx); err != nil {
		return err
	}

	return cron.Kernel()
}
