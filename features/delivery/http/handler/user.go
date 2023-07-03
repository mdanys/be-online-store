package handler

import (
	"be-online-store/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func (uh *UserHandler) Login(c *fiber.Ctx) (err error) {
	var input domain.LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	res, err := uh.UserUsecase.GetUserLogin(c.Context(), input)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}
