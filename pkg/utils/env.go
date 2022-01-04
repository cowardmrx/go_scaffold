package utils

import (
	"github.com/spf13/cast"
	"os"
)

//	@method GetStringEnv
//	@description: 获取指定字符串类型的环境变量
//	@param name string
//	@param defaultEnv string
//	@return string
func GetStringEnv(name, defaultEnv string) string {

	val, ok := os.LookupEnv(name)

	if !ok {
		return defaultEnv
	}

	return val
}

//	@method GetIntEnv
//	@description: 获取指定数值类型的环境变量
//	@param name string
//	@param defaultEnv int
//	@return int
func GetIntEnv(name string, defaultEnv int) int {
	val, ok := os.LookupEnv(name)

	if !ok {
		return defaultEnv
	}

	return cast.ToInt(val)
}

//	@method GetBoolEnv
//	@description: 获取布尔类型的环境变量
//	@param name string
//	@param defaultEnv bool
//	@return bool
func GetBoolEnv(name string, defaultEnv bool) bool {
	val, ok := os.LookupEnv(name)

	if !ok {
		return defaultEnv
	}

	return cast.ToBool(val)
}
