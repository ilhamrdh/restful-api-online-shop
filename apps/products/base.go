package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	productRouter := router.Group("products")
	{
		productRouter.Post("", controller.CreateProduct)
		productRouter.Get("", controller.GetListProducts)
		productRouter.Get("/sku/:sku", controller.GetProductDetail)
	}
}
