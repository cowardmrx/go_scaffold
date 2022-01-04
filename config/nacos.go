package config

import (
	"github.com/spf13/cast"
	"go_scaffold/pkg/utils"
)

type nacos struct {
	Host        string `yaml:"host"`
	Port        uint64 `yaml:"port"`
	NamespaceID string `yaml:"namespace_id"`
	ConfigName  string `yaml:"config_name"`
}

//	@method GetNacosConfig
//	@description: 获取nacos配置
func GetNacosConfig() {
	host := utils.GetStringEnv("NACOS_HOST", "192.168.0.151")
	port := utils.GetIntEnv("NACOS_PORT", 8848)
	namespaceID := utils.GetStringEnv("NAMESPACE_ID", "")
	configName := utils.GetStringEnv("CONFIG_NAME", "app.yaml")

	Nacos = &nacos{
		Host:        host,
		Port:        cast.ToUint64(port),
		NamespaceID: namespaceID,
		ConfigName:  configName,
	}

	return
}
