package http

import (
	"be-online-store/config"
	"be-online-store/domain"
	"be-online-store/features/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func authRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(config.NewConfig().JWTSecret),
	})
}

func RouteAPI(app *fiber.App, user domain.UserUsecase) {
	handlerUser := &handler.UserHandler{UserUsecase: user}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", handlerUser.Login)
	app.Post("/user", handlerUser.CreateUser)
	app.Get("/user", authRequired(), handlerUser.GetAllUser)
	app.Patch("/user", authRequired(), handlerUser.UpdateUser)
	app.Get("/user/:id", handlerUser.GetUserByID)
}
