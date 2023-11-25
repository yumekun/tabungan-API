package redis_store

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (store *RedisStore) AddToStream(ctx context.Context, stream string, values interface{}) error {
	cmd := store.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: values,
	})

	_, err := cmd.Result()
	if err != nil {
		return err
	}

	return nil
}

