package config

import "go_scaffold/pkg/utils"

type app struct {
	Name string `yaml:"name"`
}

func GetAppInfo() {
	Name := utils.GetStringEnv("APP_NAME", "go_scaffold")

	AppInfo = &app{
		Name: Name,
	}

	return
}
