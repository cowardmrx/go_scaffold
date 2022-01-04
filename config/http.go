package config

import (
	"github.com/spf13/cast"
	"go_scaffold/pkg/utils"
)

type http struct {
	Port string `yaml:"port"`
}

func GetHttp() {
	port := utils.GetIntEnv("PORT", 8989)

	Http = &http{
		Port: cast.ToString(port),
	}

	return
}
