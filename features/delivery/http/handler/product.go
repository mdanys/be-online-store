package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/aws"
	"be-online-store/utils/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) (err error) {
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

	res, err := ph.ProductUsecase.CreateProduct(c.Context(), input)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (ph *ProductHandler) GetListProduct(c *fiber.Ctx) (err error) {
	page, err := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	limit, err := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	categoryId := c.QueryInt("category_id")

	res, err := ph.ProductUsecase.GetListProduct(c.Context(), page, limit, int64(categoryId))
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}
