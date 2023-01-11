package handlers

import (
	"broker-exchange/myconfig/mymodels"

	"github.com/gofiber/fiber/v2"
)

func HelloProject(c *fiber.Ctx) error {
	return c.Status(200).JSON(mymodels.MyResponse("hello"))
}
