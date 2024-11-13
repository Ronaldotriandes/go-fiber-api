package handler

import (
	"github.com/Ronaldotriandes/go-fiber-api/models"
	"github.com/Ronaldotriandes/go-fiber-api/repository"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := h.userRepository.Create(user); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	users, err := h.userRepository.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User get successfully",
		"users":   users,
	})
}
