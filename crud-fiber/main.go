package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nor1c/go-cardio/crud-fiber/infra/database"
	sellerRoute "github.com/nor1c/go-cardio/crud-fiber/modules/seller/sellerRoute"
)

func main() {
	database.Connect()

	//
	app := fiber.New()

	sellerRoute.Main(app)

	app.Listen(":8080")
}
