package auth

import (
	"net/http"
	infrafiber "online-shop/infra/fiber"
	"online-shop/infra/response"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (c Controller) Register(ctx *fiber.Ctx) error {
	var req = RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(http.StatusBadRequest),
			infrafiber.WithMessage("register fail"),
		).Send(ctx)
	}

	if err := c.service.Register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithMessage("register success"),
	).Send(ctx)

}
func (c Controller) Login(ctx *fiber.Ctx) error {
	var req = LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("login fail"),
		).Send(ctx)
	}
	token, err := c.service.Login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithHttpCode(http.StatusCreated),
		infrafiber.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
		infrafiber.WithMessage("login success"),
	).Send(ctx)
}
