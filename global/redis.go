package global

import (
	"github.com/go-redis/redis/v8"
	"go_scaffold/config"
)

//	@method InitRedis
//	@description: connect redis
func InitRedis() {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Redis.GetRedisDSN(),
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	Redis = conn
}
