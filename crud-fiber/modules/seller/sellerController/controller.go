package sellercontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
	"github.com/nor1c/go-cardio/crud-fiber/infra/utils"
	service "github.com/nor1c/go-cardio/crud-fiber/modules/seller/sellerService"
)

func GetAll(c *fiber.Ctx) error {
	sellers, err := service.GetSellers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Success: true,
		Message: "",
		Data:    sellers,
	})
}

func FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	seller, err := service.GetSellerDetail(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Success: true,
		Message: "",
		Data:    seller,
	})
}

func Create(c *fiber.Ctx) error {
	seller := new(models.Seller)

	if err := c.BodyParser(&seller); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Request body invalid!",
			"err":     err.Error(),
		})
	}

	// validate user input
	validationErr := utils.ValidateStruct(seller)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErr)
	}

	// insert user input to the database
	_, err := service.RegisterNewSeller(seller)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "Seller successfully registered!",
		"seller":  seller,
	})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var seller *models.Seller

	if err := c.BodyParser(&seller); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve your input, please contact Administrator or try again later!",
		})
	}

	// make sure that the seller exists
	_, err := service.GetSellerDetail(&id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(&utils.ErrorResponse{
				Success: false,
				Message: "Seller not found!",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		}
	}

	_, err = service.UpdateSeller(&id, seller)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Seller detail successfully updated!",
		Data:    seller,
	})
}

func Remove(c *fiber.Ctx) error {
	var seller *models.Seller
	id := c.Params("id")

	isRemoved := service.RemoveSeller(&id, seller)
	if !isRemoved {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.ErrorResponse{
			Success: false,
			Message: "Failed to remove seller, please try again later.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.SuccessResponse{
		Success: true,
		Message: "Seller removed!",
	})
}
