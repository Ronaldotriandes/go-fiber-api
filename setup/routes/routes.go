package routes

import (
	"github.com/Ronaldotriandes/go-fiber-api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	api := app.Group("/api")

	// Setup individual route groups
	userRoutes(api, userHandler)
	productRoutes(api, productHandler)
}
func userRoutes(api fiber.Router, userHandle *handler.UserHandler) {
	users := api.Group("/users")
	users.Get("/", userHandle.GetAll)
	users.Post("/", userHandle.CreateUser)
}

func productRoutes(api fiber.Router, productHandler *handler.ProductHandler) {
	products := api.Group("/products")
	products.Get("/", productHandler.GetAll)
	products.Post("/", productHandler.CreateProduct)
}
