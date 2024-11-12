package main

import (
	"log"
	"time"

	"github.com/Ronaldotriandes/go-fiber-api/src/config"
	"github.com/Ronaldotriandes/go-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type SomeStruct struct {
	Name string
	Age  uint8
}

func main() {
	cfg := config.NewConfig()

	database.DBConnect(cfg)
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Use(logger.New())

	routesSetup(app)

	log.Fatal(app.Listen(":" + cfg.AppPort))

}

func routesSetup(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("hello word")
	})

}
