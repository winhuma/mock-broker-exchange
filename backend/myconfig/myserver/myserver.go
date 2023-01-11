package myserver

import (
	"broker-exchange/myconfig/myvariable"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt"
)

func New() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(AuthJWT())
	app.Use(logger.New())

	return app
}

func RunServe(app *fiber.App, port string) {
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}

// ############################################
func AuthJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := strings.Split(c.Get("Authorization"), " ")
		if len(authHeader) < 2 {
			return c.Next()
		}
		tokenString := authHeader[1]
		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			return c.Next()
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			userID, ok := claims["user_id"].(string)
			if ok {
				c.Request().Header.Add(myvariable.HeaderXUserID, fmt.Sprint(userID))
			}
		}
		return c.Next()
	}
}
