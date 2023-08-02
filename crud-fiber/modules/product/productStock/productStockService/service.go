package productquantityservice

import (
	"errors"

	"github.com/nor1c/go-cardio/crud-fiber/infra/database"
	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
)

func Add(stock *models.ProductQuantity) error {
	if err := database.DB.Create(&stock).Error; err != nil {
		return err
	}

	return nil
}

func Reduce(stock *models.ProductQuantity) error {
	if stock.Type != "OUT" {
		return errors.New("Stock type must be 'OUT'")
	}

	if err := database.DB.Create(&stock).Error; err != nil {
		return err
	}

	return nil
}
