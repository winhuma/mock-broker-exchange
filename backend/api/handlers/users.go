package handlers

import (
	"broker-exchange/api/services"
	"broker-exchange/myconfig/mymodels"
	"broker-exchange/myconfig/myvariable"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(c *fiber.Ctx) error {

	var userData mymodels.BodyUserLogin
	mybody := c.Body()
	err := json.Unmarshal(mybody, &userData)
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("data invalid format."))
	}

	mytoken, failMSG, err := services.GenerateToken(userData)
	if err != nil {
		return err
	}
	if failMSG != "" {
		return c.Status(400).JSON(mymodels.MyResponse(failMSG))
	}
	return c.Status(200).JSON(mymodels.MyResponse("Login success.", mytoken))
}

func UserGetMyBalance(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Get(myvariable.HeaderXUserID))
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("Invalid user."))
	}
	mydata, err := services.UserGetBalanceByUserID(userID)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(mymodels.MyResponse("Get data success.", mydata))
}
