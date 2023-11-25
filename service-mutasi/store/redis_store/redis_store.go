package redis_store

import (
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	redisClient *redis.Client
}

func NewRedisStore(redisClient *redis.Client) *RedisStore {
	return &RedisStore{
		redisClient: redisClient,
	}
}