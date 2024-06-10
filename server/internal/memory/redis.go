package memory

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	ACCOUNT_PREFIX = "_account"
)

func Set(redisClient *redis.Client, key, value string) error {
	ctx := context.Background()
	err := redisClient.Set(ctx, key, value, 0)
	return err.Err()
}

func Get(redisClient *redis.Client, key string) (string, error) {
	ctx := context.Background()
	value := redisClient.Get(ctx, key)
	if value.Err() != nil {
		return "", value.Err()
	}

	return value.String(), nil
}
