package utils

import (
	"chatserver/pkg/database"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func getVersionKey(prefix string) string {
	return fmt.Sprintf("%s:version", prefix)
}

func GenerateCacheKey(prefix, key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	version, err := database.RedisClient.Get(ctx, getVersionKey(prefix)).Int64()
	if err != nil {
		version = 1
	}

	return fmt.Sprintf("%s:v%d:%s", prefix, version, key)
}

func InvalidateCache(prefix string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return database.RedisClient.Incr(ctx, getVersionKey(prefix)).Err()
}

func RedisSet[T any](key string, value T, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return database.RedisClient.Set(ctx, key, data, expiration).Err()
}

func RedisGet[T any](key string) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := database.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func RedisDelete(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return database.RedisClient.Del(ctx, keys...).Err()
}

func RedisExists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := database.RedisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}
