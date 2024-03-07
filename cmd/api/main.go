package main

import (
	"log"
	"online-shop/apps/auth"
	"online-shop/apps/products"
	"online-shop/external/database"
	"online-shop/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectionPostgres(config.Cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connection")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})
	auth.Init(router, db)
	products.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
