package main

import (
	"log"
	"time"

	"github.com/Ronaldotriandes/go-fiber-api/config"
	"github.com/Ronaldotriandes/go-fiber-api/database"
	"github.com/Ronaldotriandes/go-fiber-api/handler"
	"github.com/Ronaldotriandes/go-fiber-api/middleware"
	"github.com/Ronaldotriandes/go-fiber-api/repository"
	"github.com/Ronaldotriandes/go-fiber-api/setup/routes"
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
	productRepo := repository.NewProductRepository(db)

	userHandle := handler.NewUserHandler(userRepo)
	productHandle := handler.NewProductHandler(productRepo)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Use(logger.New())
	routes.SetupRoutes(app, userHandle, productHandle)

	log.Fatal(app.Listen(":" + cfg.AppPort))

}
