package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/nor1c/go-cardio/crud-fiber/infra/database"
	appRoute "github.com/nor1c/go-cardio/crud-fiber/infra/route"
	"github.com/nor1c/go-cardio/crud-fiber/infra/utils"
)

var validate = validator.New()

func main() {
	// connect to database on app start
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: "CRUD with Fiber",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	// routes
	appRoute.Main(app)

	log.Fatal(app.Listen(":8080"))
}
