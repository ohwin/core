package global

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RDB struct {
	Client redis.UniversalClient
}

func (r *RDB) Set(key string, value interface{}, expiration time.Duration) error {
	if err := r.Client.Set(context.TODO(), key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RDB) Get(key string) (string, error) {
	result, err := r.Client.Get(context.TODO(), key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
