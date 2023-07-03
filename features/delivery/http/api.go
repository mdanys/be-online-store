package http

import (
	"be-online-store/domain"
	"be-online-store/features/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteAPI(app *fiber.App, user domain.UserUsecase) {
	handlerUser := &handler.UserHandler{UserUsecase: user}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", handlerUser.Login)
}
