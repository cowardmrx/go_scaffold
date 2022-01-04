package global

import "github.com/cowardmrx/rds_cache_go"

//	@method InitCache
//	@description: init cache default use redis [only support redis]
func InitCache() {
	cache := rds_cache_go.NewCache(rds_cache_go.WithRedisClient(Redis), rds_cache_go.WithDB(10))

	Cache = cache
}
