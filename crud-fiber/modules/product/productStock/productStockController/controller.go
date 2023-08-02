package productquantitycontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
	"github.com/nor1c/go-cardio/crud-fiber/infra/utils"

	service "github.com/nor1c/go-cardio/crud-fiber/modules/product/productStock/productStockService"
)

func AddStock(c *fiber.Ctx) error {
	stock := new(models.ProductQuantity)

	if err := c.BodyParser(stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Invalid request body.",
		})
	}

	if err := service.Add(stock); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to add stock to selected product!",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Product quantity increased!",
	})
}

func ReduceStock(c *fiber.Ctx) error {
	stock := new(models.ProductQuantity)

	if err := c.BodyParser(&stock); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Invalid request body.",
		})
	}

	if err := service.Reduce(stock); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to reduce stock of the selected product!",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Product quantity reduced!",
	})
}
