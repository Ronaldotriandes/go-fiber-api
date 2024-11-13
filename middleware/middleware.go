package middleware

import (
	errors "github.com/Ronaldotriandes/go-fiber-api/internal"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*errors.ErrorApp); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"error": e.Message,
		})
	}
	errorMess := "Internal Server Errorss"
	if err != nil {
		errorMess = err.Error()
	}
	// Handle unexpected errors
	return c.Status(500).JSON(fiber.Map{
		"error": errorMess,
	})
}
