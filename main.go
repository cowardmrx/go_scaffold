package main

import (
	"github.com/urfave/cli/v2"
	"go_scaffold/global"
	"go_scaffold/initialize"
	"log"
	"os"
	"sort"
)

func init() {
	global.App = cli.NewApp()

	initialize.Init()
}

func main() {
	sort.Sort(cli.CommandsByName(global.App.Commands))
	err := global.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
