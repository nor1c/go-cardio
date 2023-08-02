package productcontroller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
	"github.com/nor1c/go-cardio/crud-fiber/infra/utils"

	service "github.com/nor1c/go-cardio/crud-fiber/modules/product/productService"
)

func GetAll(c *fiber.Ctx) error {
	products := new([]models.Product)

	if err := service.GetAll(products); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Something went wrong when fetching the product list.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Data:    &products,
		Message: "",
	})
}

func GetDetail(c *fiber.Ctx) error {
	product := new(models.Product)
	fmt.Println(product)

	id := c.Params("id")

	if err := service.FindById(product, &id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to fetch selected product!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Data:    &product,
		Message: "",
	})
}

func Create(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Request body invalid",
		})
	}

	if err := service.Create(product); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to publish your product!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Data:    &product,
		Message: "Product published!",
	})
}

func Update(c *fiber.Ctx) error {
	product := new(models.Product)

	id := c.Params("id")

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Invalid request body!",
		})
	}

	if err := service.Update(product, &id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to update the selected product!",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Product detail updated!",
		Data:    &product,
	})
}

func Remove(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.Remove(&id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to remove the product!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Product successfully removed!",
	})
}
