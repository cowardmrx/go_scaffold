package cli

import (
	"github.com/urfave/cli/v2"
	"go_scaffold/cli/serve"
	"go_scaffold/global"
)

//	@method App
//	@description: application serve & serve usage
func App() {
	// set application name
	global.App.Name = "go scaffold"

	// set application desc
	global.App.Usage = "a fast develop scaffold by Golang"

	// set author
	global.App.Authors = []*cli.Author{
		{
			Name:  "coward",
			Email: "2991148012@qq.com",
		},
	}

	// application commands
	global.App.Commands = []*cli.Command{
		{
			Name:    "http:serve",
			Aliases: []string{"app:serve"},
			Usage:   "start http server",
			Action:  serve.HttpServe,
		},
		{
			Name:    "cron:serve",
			Aliases: []string{"cron"},
			Usage:   "start cron server",
			Action:  serve.Cron,
		},
	}

	// global vars
	global.App.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config-sources",
			Aliases:     []string{"cfs"},
			Usage:       "config sources(from local or nacos)",
			Value:       "LOCAL",
			DefaultText: "LOCAL",
			Destination: &global.ConfigSources,
		},
	}
}
