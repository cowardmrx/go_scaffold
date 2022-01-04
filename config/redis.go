package config

import "fmt"

type redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

//	@method GetRedisDSN
//	@description: get redis connect string
//	@receiver r
//	@return string
func (r *redis) GetRedisDSN() string {
	return fmt.Sprintf("%s:%v", r.Host, r.Port)
}
