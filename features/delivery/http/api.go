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

func RouteAPI(app *fiber.App, user domain.UserUsecase, category domain.CategoryUsecase, product domain.ProductUsecase, cart domain.CartUsecase, order domain.OrderUsecase) {
	handlerUser := &handler.UserHandler{UserUsecase: user}
	handlerCategory := &handler.CategoryHandler{CategoryUsecase: category}
	handlerProduct := &handler.ProductHandler{ProductUsecase: product}
	handlerCart := &handler.CartHandler{CartUsecase: cart}
	handlerOrder := &handler.OrderHandler{OrderUsecase: order}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Login
	app.Post("/login", handlerUser.Login)

	// User
	app.Post("/user", handlerUser.CreateUser)
	app.Get("/user", authRequired(), handlerUser.GetAllUser)
	app.Patch("/user", authRequired(), handlerUser.UpdateUser)
	app.Get("/user/:id", handlerUser.GetUserByID)

	// Category
	app.Post("/category", authRequired(), handlerCategory.CreateCategory)
	app.Get("/category", handlerCategory.GetAllCategory)

	// Product
	app.Post("/product", authRequired(), handlerProduct.CreateProduct)
	app.Get("/product", handlerProduct.GetListProduct)

	// Cart
	app.Post("/cart", authRequired(), handlerCart.CreateCart)
	app.Get("/cart", authRequired(), handlerCart.GetCartByUserID)
	app.Delete("/cart/:id", authRequired(), handlerCart.DeleteCart)

	// Order
	app.Post("/order", authRequired(), handlerOrder.CreateOrder)
	app.Get("/order", authRequired(), handlerOrder.GetOrderByUserID)
	app.Put("/order/:order_id", authRequired(), handlerOrder.UpdateOrderStatus)
	app.Get("/order/:order_id", authRequired(), handlerOrder.GetOrderByOrderID)
}
