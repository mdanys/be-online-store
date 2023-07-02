package main

import (
	"be-online-store/config"
	"be-online-store/features/delivery/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	http.RouteAPI(app)

	if err := app.Listen(":" + strconv.Itoa(config.NewConfig().ServerPort)); err != nil {
		log.Fatal(err)
	}
}
