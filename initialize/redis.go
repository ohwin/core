package initialize

import (
	"github.com/ohwin/core/global"
	"github.com/redis/go-redis/v9"
)

// Redis 获取Redis链接
func Redis() {
	config := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})
	global.Redis = &global.RDB{
		Client: client,
	}
}
