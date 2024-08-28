package utils

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// retrieve data from redis
func GetCache(ctx context.Context, rdb *redis.Client, key string) ([]byte, error) {
	log.Printf("Getting cache for key : %s", key)
	result, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

// set data in Redis with an expiration time
func SetCache(ctx context.Context, rdb *redis.Client, key string, value []byte, expiration time.Duration) error {
	err := rdb.Set(ctx, key, value, expiration).Err()
	return err
}
