package cmd

import (
	"github.com/urfave/cli/v2"
	"go_scaffold/app/model"
	"go_scaffold/global"
)

type GormCmd struct{}

type GormCmdFunc interface {
	Migrate(ctx *cli.Context) error
}

//	@method NewGormCmd
//	@description: new gorm cmd
//	@return *GormCmd
func NewGormCmd() *GormCmd {
	return &GormCmd{}
}

//	@method Migrate
//	@description: gorm auto migrate
//	@receiver g
//	@return error
func (g *GormCmd) Migrate(ctx *cli.Context) error {

	return global.Database.AutoMigrate(
		&model.UserModel{}, // 用户模型
	)
}
