package main

import (
	"be-online-store/config"
	"be-online-store/features/delivery/http"
	"be-online-store/features/repository/mysql"
	"be-online-store/features/usecase"
	"be-online-store/utils/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	cfg := config.NewConfig()
	dbConn, _ := database.InitDatabase(cfg)

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	userRepo := mysql.NewMySQLUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo)

	http.RouteAPI(app, userUsecase)

	if err := app.Listen(":" + strconv.Itoa(config.NewConfig().ServerPort)); err != nil {
		log.Fatal(err)
	}
}
