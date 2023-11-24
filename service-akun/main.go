package main

import (
	"database/sql"
	"fmt"

	"service-akun/db/store/postgres_store"
	"service-akun/handler"
	"service-akun/service"
	util "service-akun/util/config"

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


	store := postgres_store.NewPostgresStore(conn)

	service := service.NewService(store)

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