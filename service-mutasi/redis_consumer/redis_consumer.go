package redis_consumer

import (
	"service-mutasi/service"
	db "service-mutasi/store/postgres_store/store"
	"service-mutasi/store/redis_store"
	util "service-mutasi/util/config"
)

type store struct {
	redis *redis_store.RedisStore
}

func newStore(redisStore *redis_store.RedisStore) store {
	return store{
		redis: redisStore,
	}
}

type RedisConsumer struct {
	config  util.Config
	streamName string
	service service.IService
	store   store
}

func NewRedisConsumer(config util.Config,streamName string,postgresStore db.IStore, redisStore *redis_store.RedisStore) *RedisConsumer {
	service := service.NewService(postgresStore)
	store := newStore(redisStore)

	return &RedisConsumer{
		config:  config,
		streamName: streamName,
		service: service,
		store:   store,
	}
}