package global

import (
	"gitee.ltd/lxh/logger"
	"go_scaffold/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"time"
)

//	@method InitDatabase
//	@description: connect database
func InitDatabase() {

	var conn *gorm.DB
	var err error

	// 自定义日志配置，使用整合后的zap
	newLogger := gormLogger.New(
		log.New(logger.NewGormLogger(), "", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			IgnoreRecordNotFoundError: false,           // 忽略没找到结果的错误
			LogLevel:                  gormLogger.Info, // Log level
			Colorful:                  false,           // Disable color
		},
	)

	switch config.Database.Driver {
	case config.MysqlDriver:
		conn, err = gorm.Open(mysql.Open(config.Database.GetDatabaseDSN()), &gorm.Config{
			Logger: newLogger,
		})
	}

	if err != nil {
		panic("connect database failed:" + err.Error())
	}

	Database = conn
}
