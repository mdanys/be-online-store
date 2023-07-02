package middleware

import (
	"be-online-store/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func GenerateToken(id int, role string) (s string, err error) {
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

func ExtractToken(c *fiber.Ctx) (id int, role string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id = claims["id"].(int)
	role = claims["role"].(string)

	return
}
