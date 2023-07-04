package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type OrderHandler struct {
	OrderUsecase domain.OrderUsecase
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) (err error) {
	_, role := middleware.ExtractToken(c)
	if role != "customer" {
		return c.Status(fasthttp.StatusUnauthorized).SendString("Only customer")
	}

	var input []int64
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	link, err := oh.OrderUsecase.CreateOrder(c.Context(), input...)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(fiber.Map{"status": "Success", "redict_link": link})
}
