package main

import (
	"context"
	"database/sql"
	"fmt"

	"service-akun/handler"
	"service-akun/service"
	"service-akun/store/postgres_store/store/postgres_store"
	"service-akun/store/redis_store"
	util "service-akun/util/config"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	redisStore  := redis_store.NewRedisStore(redisClient)

	service := service.NewService(config,postgresStore,redisStore)

	handler := handler.NewHandler(service)

	app := fiber.New()

	corsConfig := cors.Config{
		AllowHeaders:     "Content-Type,Authorization",
		AllowCredentials: true,
	}

	app.Use(cors.New(corsConfig))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})

	app.Post("/daftar", handler.Daftar)
	app.Post("/tarik", handler.Tarik)
	app.Post("/tabung", handler.Tabung)
	app.Get("/saldo/:no_rekening", handler.Saldo)
	app.Get("/mutasi/:no_rekening", handler.Mutasi)

	host := config.Host
	port := config.Port

	err = app.Listen(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
 fmt.Printf("failed to listen at port '%s': %s", port, err)

		
		return
	}
}