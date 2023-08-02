package productroute

import (
	"github.com/gofiber/fiber/v2"

	productController "github.com/nor1c/go-cardio/crud-fiber/modules/product/productController"
	productStockController "github.com/nor1c/go-cardio/crud-fiber/modules/product/productStock/productStockController"
)

func Main(app *fiber.App) {
	product := app.Group("/products")

	product.Get("/", productController.GetAll)
	product.Get("/:id", productController.GetDetail)
	product.Post("/", productController.Create)
	product.Put("/:id", productController.Update)
	product.Delete("/:id", productController.Remove)

	quantity := product.Group("/stock")
	quantity.Post("/add", productStockController.AddStock)
	quantity.Post("/reduce", productStockController.ReduceStock)
}
