package handlers

import (
	"broker-exchange/api/services"
	"broker-exchange/myconfig/mymodels"
	"broker-exchange/myconfig/myvariable"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func OrderCreate(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Get(myvariable.HeaderXUserID))
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("Invalid user."))
	}
	var mydata mymodels.BodyOrderCreate
	err = json.Unmarshal(c.Body(), &mydata)
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("data invalid format."))
	}
	err = services.OrderCreate(userID, mydata)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.MyResponse("Create order success."))
}
