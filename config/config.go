package config

var (
	AppInfo  *app
	Http     *http
	Database *database
	Redis    *redis
	Nacos    *nacos
)

type ApplicationConf struct {
	App      *app
	Http     *http
	Database *database
	Redis    *redis
	Nacos    *nacos
}

//	@method InitApplicationConf
//	@description: init application config
//	@receiver appConf
func (appConf *ApplicationConf) InitApplicationConf() {
	AppInfo = appConf.App
	Http = appConf.Http
	Database = appConf.Database
	Redis = appConf.Redis
	Nacos = appConf.Nacos
}
