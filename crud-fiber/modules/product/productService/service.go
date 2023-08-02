package productservice

import (
	"github.com/nor1c/go-cardio/crud-fiber/infra/database"
	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
)

func GetAll(sellers *[]models.Product) error {
	if err := database.DB.Find(&sellers).Error; err != nil {
		return err
	}

	return nil
}

func FindById(product *models.Product, id *string) error {
	if err := database.DB.Where("id=?", &id).Preload("AvailableStock").First(&product).Error; err != nil {
		return err
	}

	return nil
}

func Create(product *models.Product) error {
	if err := database.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func Update(product *models.Product, id *string) error {
	if err := database.DB.Where("id", &id).Updates(&product).Error; err != nil {
		return err
	}

	return nil
}

func Remove(id *string) error {
	product := new(models.Product)

	if err := database.DB.Delete(&product, &id).Error; err != nil {
		return err
	}

	return nil
}
