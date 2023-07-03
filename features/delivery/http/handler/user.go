package handler

import (
	"be-online-store/domain"
	"be-online-store/utils/aws"
	"be-online-store/utils/middleware"
	"strconv"

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

func (uh *UserHandler) CreateUser(c *fiber.Ctx) (err error) {
	var input domain.UserRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("user_picture")
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	if file != nil {
		input.UserPicture, err = aws.UploadFile(file)
		if err != nil {
			return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
		}
	}

	res, err := uh.UserUsecase.CreateUser(c.Context(), input)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) (err error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	res, err := uh.UserUsecase.GetUserByID(c.Context(), int64(id))
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}

func (uh *UserHandler) GetAllUser(c *fiber.Ctx) (err error) {
	_, role := middleware.ExtractToken(c)
	if role != "admin" {
		return c.Status(fasthttp.StatusUnauthorized).SendString("Only admin")
	}

	page, err := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	limit, err := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).SendString(err.Error())
	}

	res, err := uh.UserUsecase.GetAllUser(c.Context(), page, limit)
	if err != nil {
		return c.Status(fasthttp.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fasthttp.StatusOK).JSON(res)
}
