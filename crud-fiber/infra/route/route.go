package route

import (
	"github.com/gofiber/fiber/v2"
	productRoute "github.com/nor1c/go-cardio/crud-fiber/modules/product/productRoute"
	sellerRoute "github.com/nor1c/go-cardio/crud-fiber/modules/seller/sellerRoute"
)

func Main(app *fiber.App) {
	sellerRoute.Main(app)
	productRoute.Main(app)
}
