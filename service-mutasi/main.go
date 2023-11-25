package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"service-mutasi/redis_consumer"
	"service-mutasi/store/postgres_store/store/postgres_store"
	"service-mutasi/store/redis_store"
	util "service-mutasi/util/config"
	"syscall"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("cannot load config: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		fmt.Printf("cannot connect to db: %s", err)
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisServiceAddress,
		Password: config.RedisPassword,
		DB:       0,
	})

	// ping redis
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("cannot connect to redis: %s", err)

		return
	}
	postgresStore := postgres_store.NewPostgresStore(conn)
	redisStore := redis_store.NewRedisStore(redisClient)

	redisConsumer := redis_consumer.NewRedisConsumer(config, postgresStore, redisStore)

	go redisConsumer.Run(context.Background())

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("exit", <-errs)
}