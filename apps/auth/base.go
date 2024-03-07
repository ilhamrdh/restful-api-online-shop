package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	authRouter := router.Group("auth")
	{
		authRouter.Post("register", controller.Register)
		authRouter.Post("login", controller.Login)
	}
}
