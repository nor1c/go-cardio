package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	models "github.com/nor1c/go-cardio/crud-fiber/infra/models"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/go-cardio-crud-fiber"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Seller{},
		&models.Product{},
		&models.ProductQuantity{},
	)

	DB = db
}
