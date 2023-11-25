package redis_store

import (
	"context"

	"github.com/go-redis/redis/v8"
)


func (store *RedisStore) GetStreams(ctx context.Context, stream string, count int) ([]redis.XStream, error) {
	cmd := store.redisClient.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "0"},
		Count:   int64(count),
		Block:   0,
	})

	streams, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	if len(streams) > 0 {
		return streams, nil
	}

	return nil, nil
}