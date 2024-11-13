package main

import (
	"log"
	"time"

	"github.com/Ronaldotriandes/go-fiber-api/config"
	"github.com/Ronaldotriandes/go-fiber-api/database"
	"github.com/Ronaldotriandes/go-fiber-api/handler"
	"github.com/Ronaldotriandes/go-fiber-api/middleware"
	"github.com/Ronaldotriandes/go-fiber-api/repository"
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
	db := database.GetDB()
	userRepo := repository.NewUserRepository(db)
	userHandle := handler.NewUserHandler(userRepo)
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Use(logger.New())

	routesSetup(app, userHandle)

	log.Fatal(app.Listen(":" + cfg.AppPort))

}

func routesSetup(app *fiber.App, userHandle *handler.UserHandler) {
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("hello word")
	})
	users := api.Group("/users")
	users.Get("/", userHandle.GetAll)
	users.Post("/", userHandle.CreateUser)

}
