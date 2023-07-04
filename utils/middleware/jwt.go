package middleware

import (
	"be-online-store/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

func GenerateToken(id int64, role string) (s string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["expired"] = time.Now().Add(time.Hour * 24 * 7)

	s, err = token.SignedString([]byte(config.NewConfig().JWTSecret))
	if err != nil {
		log.Error("error on token signed string: ", err.Error())
		return
	}

	return
}

func ExtractToken(c *fiber.Ctx) (id int64, role string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id = int64(claims["id"].(float64))
	role = claims["role"].(string)

	return
}
