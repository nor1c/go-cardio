package database

import (
	"fmt"

	"github.com/nor1c/go-cardio/crud-fiber/infra/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// models "github.com/nor1c/go-cardio/crud-fiber/infra/models"
)

var DB *gorm.DB

func CreateViewTables(db *gorm.DB) {
	const createCurrentStockViewTable = `	
		CREATE OR REPLACE VIEW 
			available_stock
		AS SELECT
			product_id,
			(select SUM(qty) as qty FROM product_quantities WHERE type="IN")-
			(select SUM(qty) as qty FROM product_quantities WHERE type="OUT") AS current_stock
		FROM product_quantities
		GROUP BY product_id
	`
	if err := db.Exec(createCurrentStockViewTable).Error; err != nil {
		panic(err)
	}
}

func Connect() {
	fmt.Println("🟡 Trying to connect to DB..")
	db, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/go-cardio-crud-fiber?parseTime=true"))
	if err != nil {
		panic(err)
	}

	CreateViewTables(db)

	db.AutoMigrate(
		&models.Seller{},
		&models.Product{},
		&models.ProductQuantity{},
	)

	DB = db

	fmt.Println("✅ Connected to DB!")
}
