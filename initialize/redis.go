package initialize

import (
	"core/global"
	"github.com/redis/go-redis/v9"
)

// Redis 获取Redis链接
func Redis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
