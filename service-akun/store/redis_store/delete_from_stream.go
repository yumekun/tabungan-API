package redis_store

import (
	"context"
)



func (store *RedisStore) DeleteFromStream(ctx context.Context, stream string, messageID string) error {
	cmd := store.redisClient.XDel(ctx, stream, messageID)

	_, err := cmd.Result()
	if err != nil {
		return err
	}

	return nil
}


