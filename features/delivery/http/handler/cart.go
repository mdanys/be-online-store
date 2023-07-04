package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type CartHandler struct {
	CartUsecase domain.CartUsecase
}

func (ch *CartHandler) CreateCart(c *fiber.Ctx) (err error) {
	userId, role := middleware.ExtractToken(c)
	if role != "customer" {
		return c.Status(fasthttp.StatusUnauthorized).SendString("Only customer")
	}

	var input domain.CartRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	input.UserID = &userId

	err = validate.Struct(input)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	err = ch.CartUsecase.CreateCart(c.Context(), input)
	if err != nil {
		if err.Error() == "not enough stock" {
			return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).SendString("Success")
}
