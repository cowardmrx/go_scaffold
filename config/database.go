package config

import "fmt"

const (
	MysqlDriver = "mysql"
)

type database struct {
	Driver      string `yaml:"driver"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Name        string `yaml:"name"`
	AutoMigrate bool   `yaml:"auto_migrate"`
}

//	@method GetDatabaseDSN
//	@description: get database connect string
//	@receiver db
//	@return string
func (db *database) GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Password, db.Host, db.Port, db.Name)
}
