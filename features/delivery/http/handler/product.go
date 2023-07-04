package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/aws"
	"be-online-store/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func (uh *ProductHandler) CreateProduct(c *fiber.Ctx) (err error) {
	_, role := middleware.ExtractToken(c)
	if role != "admin" {
		return c.Status(fasthttp.StatusUnauthorized).SendString("Only admin")
	}

	var input domain.ProductRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	err = validate.Struct(input)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	file, _ := c.FormFile("product_picture")

	if file != nil {
		s, err := aws.UploadFile(file)
		if err != nil {
			return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
		}

		input.ProductPicture = &s
	}

	res, err := uh.ProductUsecase.CreateProduct(c.Context(), input)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}
