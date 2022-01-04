package global

import (
	"github.com/cowardmrx/rds_cache_go"
	"github.com/go-redis/redis/v8"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var (
	App *cli.App
)

var (
	Database *gorm.DB
	Redis    *redis.Client
	Cache    rds_cache_go.Cache
)

var (
	RootPath = ""
)

var (
	// ConfigSources 配置来源
	ConfigSources string
)

const (
	// ConfigLocal 本地配置文件
	ConfigLocal = "LOCAL"
	// ConfigNacos nacos配置
	ConfigNacos = "NACOS"
)
