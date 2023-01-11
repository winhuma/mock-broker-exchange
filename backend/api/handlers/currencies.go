package handlers

import (
	"broker-exchange/api/services"
	"broker-exchange/myconfig/mymodels"

	"github.com/gofiber/fiber/v2"
)

func CurrencyGetAll(c *fiber.Ctx) error {
	mydata, err := services.CurrencyGetAll()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.MyResponse("success", mydata))
}
