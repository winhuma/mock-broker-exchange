package myserver

import (
	"broker-exchange/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", handlers.HelloProject)
	app.Post("/login", handlers.UserLogin)

	app.Get("/balance/me", handlers.UserGetMyBalance)

	app.Get("/currencies", handlers.CurrencyGetAll)

	app.Post("/orders", handlers.OrderCreate)
}
