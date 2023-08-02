package sellerroute

import (
	"github.com/gofiber/fiber/v2"

	sellerController "github.com/nor1c/go-cardio/crud-fiber/modules/seller/sellerController"
)

func Main(app *fiber.App) {
	seller := app.Group("/sellers")

	seller.Get("/", sellerController.GetAll)
	seller.Get("/:id", sellerController.FindById)
	seller.Post("/", sellerController.Create)
	seller.Put("/:id", sellerController.Update)
	seller.Delete("/:id", sellerController.Remove)
}
