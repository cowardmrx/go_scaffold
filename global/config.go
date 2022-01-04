package global

import (
	"gitee.ltd/lxh/logger"
	"github.com/jinzhu/configor"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
	"go_scaffold/config"
	"go_scaffold/pkg/utils"
	"gopkg.in/yaml.v3"
)

//	@method InitConfig
//	@description: read config file and init application config
func InitConfig(ctx *cli.Context) {
	switch ctx.String("cfs") {
	case ConfigLocal:
		local()
	case ConfigNacos:
		nacos()
	}
}

//	@method local
//	@description: load config fron local file
func local() {
	conf := new(config.ApplicationConf)

	if err := configor.New(&configor.Config{
		AutoReload: true,
		AutoReloadCallback: func(config interface{}) {
			conf.InitApplicationConf()

			InitDatabase()

			InitRedis()

			InitCache()

			InitValidatorTranslator()

			logger.InitLogger(logger.LogConfig{
				Mode:       logger.Dev,
				FileEnable: true,
			})
		},
	}).Load(conf, RootPath+"app.yaml"); err != nil {
		panic("read config failed:" + err.Error())
	}

	conf.InitApplicationConf()

	InitDatabase()

	InitRedis()

	InitCache()

	InitValidatorTranslator()

	logger.InitLogger(logger.LogConfig{
		Mode:       logger.Dev,
		FileEnable: true,
	})
}

//	@method nacos
//	@description: load config and registry server by nacos
func nacos() {

	// read config fron nacos
	getConfigFromNacos()

	// registry server
	registerToNacos()

}

//	@method getNacosConfig
//	@description: get nacos client config
//	@return vo.NacosClientParam
func getNacosConfig() vo.NacosClientParam {

	config.GetAppInfo()

	config.GetHttp()

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Nacos.Host, config.Nacos.Port),
	}

	cc := constant.ClientConfig{
		AppName:             config.AppInfo.Name,
		NamespaceId:         config.Nacos.NamespaceID,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		LogDir:              "logs",
		CacheDir:            "logs",
	}

	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

//	@method registerToNacos
//	@description: registry application to nacos
func registerToNacos() {
	config.GetNacosConfig()

	configParam := getNacosConfig()

	client, err := clients.NewNamingClient(configParam)

	if err != nil {
		panic(err.Error())
	}

	ip := config.AppInfo.Name

	if ips := utils.GetIps(); ips != nil {
		ip = ips[0]
	}

	state, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        cast.ToUint64(config.Http.Port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		ServiceName: config.AppInfo.Name,
		Ephemeral:   true,
		Metadata: map[string]string{
			"name": config.AppInfo.Name,
			"ip":   ip,
		},
	})

	if err != nil || !state {
		logger.Say.Panicf("服务注册失败，退出程序: %v", err.Error())
	}

	logger.Say.Info("服务注册成功")
	return
}

//	@method getConfigFromNacos
//	@description: load config from nacos
func getConfigFromNacos() {
	config.GetNacosConfig()

	configParam := getNacosConfig()

	client, err := clients.NewConfigClient(configParam)

	if err != nil {
		logger.Say.Panicf("创建配置链接失败: %v", err.Error())
	}

	// listen config change
	_ = client.ListenConfig(vo.ConfigParam{
		DataId: config.Nacos.ConfigName,
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			configChanged(data)
		},
	})

	configStr, err := client.GetConfig(vo.ConfigParam{
		DataId: config.Nacos.ConfigName,
		Group:  "DEFAULT_GROUP",
	})

	if err != nil {
		logger.Say.Panicf("读取配置文件失败: %v", err.Error())
	}

	configChanged(configStr)
	return

}

//	@method configChanged
//	@description: config changed
//	@param data string
func configChanged(data string) {

	conf := new(config.ApplicationConf)

	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		logger.Say.Panicf("配置文件读取失败：%v", err.Error())
	}

	conf.InitApplicationConf()

	InitDatabase()

	InitRedis()

	InitCache()

	InitValidatorTranslator()

	logger.InitLogger(logger.LogConfig{
		Mode:       logger.Dev,
		FileEnable: true,
	})

	logger.Say.Info("配置文件读取成功")
}
