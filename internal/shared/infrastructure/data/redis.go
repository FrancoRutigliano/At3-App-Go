package data

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnnection() (*redis.Client, error) {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
