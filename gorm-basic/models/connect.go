package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config := mysql.Config{
		DSN: "root:@tcp(localhost:3306)/go-cardio-gorm-setup?charset=utf8&parseTime=true",
	}

	db, err := gorm.Open(mysql.New(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Author{}, &Book{})

	DB = db

	log.Println("DB connected!")
}
