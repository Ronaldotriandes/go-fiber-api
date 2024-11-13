package handler

import (
	"github.com/Ronaldotriandes/go-fiber-api/models"
	"github.com/Ronaldotriandes/go-fiber-api/repository"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productRepository repository.ProductRepository
}

func NewProductHandler(productRepository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepository}
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	// Implementasi logika untuk membuat pengguna
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := h.productRepository.Create(product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Produt created successfully",
		"product": product,
	})
}

func (h *ProductHandler) GetAll(ctx *fiber.Ctx) error {
	// Implementasi logika untuk mendapatkan semua pengguna
	products, err := h.productRepository.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "product get successfully",
		"products": products,
	})
}
