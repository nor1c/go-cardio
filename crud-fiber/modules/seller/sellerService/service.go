package sellerservice

import (
	"errors"

	"github.com/nor1c/go-cardio/crud-fiber/infra/database"
	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
)

func GetSellers() ([]models.Seller, error) {
	var sellers []models.Seller

	result := database.DB.Find(&sellers)
	if err := result.Error; err != nil {
		return nil, err
	}

	return sellers, nil
}

func GetSellerDetail(sellerId *string) (models.Seller, error) {
	var seller models.Seller

	result := database.DB.Preload("Products").Where("id=?", sellerId).First(&seller)
	if err := result.Error; err != nil {
		return seller, err
	}

	return seller, nil
}

func RegisterNewSeller(seller *models.Seller) (bool, error) {
	result := database.DB.Create(&seller)

	if result.Error != nil {
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}

func UpdateSeller(sellerId *string, seller *models.Seller) (bool, error) {
	result := database.DB.Where("id=?", &sellerId).Updates(&seller)

	if err := result.Error; err != nil {
		return false, err
	}

	return true, nil
}

func RemoveSeller(id *string, seller *models.Seller) bool {
	result := database.DB.Delete(&seller, &id)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}

func SetSellerToInactive() {

}
