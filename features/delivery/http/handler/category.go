package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

func (ch *CategoryHandler) CreateCategory(c *fiber.Ctx) (err error) {
	_, role := middleware.ExtractToken(c)
	if role != "admin" {
		return c.Status(fasthttp.StatusUnauthorized).SendString("Only admin")
	}

	var input domain.CategoryRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(input)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	err = ch.CategoryUsecase.CreateCategory(c.Context(), input.Name)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).SendString("Success")
}

func (ch *CategoryHandler) GetAllCategory(c *fiber.Ctx) (err error) {
	res, err := ch.CategoryUsecase.GetAllCategory(c.Context())
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}
