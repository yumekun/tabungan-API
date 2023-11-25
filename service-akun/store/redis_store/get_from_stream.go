package redis_store

import (
	"context"

	"github.com/go-redis/redis/v8"
)




func (store *RedisStore) GetFromStream(ctx context.Context, stream string, count int) ([]redis.XMessage, error) {
	cmd := store.redisClient.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "0"},
		Count:   int64(count),
		Block:   0,
	})

	streams, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	if len(streams) > 0 && len(streams[0].Messages) > 0 {
		messages := streams[0].Messages

		return messages, nil
	}

	return nil, nil
}

