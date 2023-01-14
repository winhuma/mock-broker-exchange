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

	mytoken, userID, failMSG, err := services.GenerateToken(userData)
	if err != nil {
		return err
	}
	if failMSG != "" {
		return c.Status(400).JSON(mymodels.MyResponse(failMSG))
	}
	dBalance, err := services.UserGetBalanceByUserID(userID)
	if err != nil {
		return err
	}
	myres := map[string]interface{}{
		"token":        mytoken,
		"user_balance": dBalance,
	}
	return c.Status(200).JSON(mymodels.MyResponse("Login success.", myres))
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

func UserAddStarterBalance(c *fiber.Ctx) error {
	var dUser mymodels.BodyUserAddStarterBalance
	userID, err := strconv.Atoi(c.Get(myvariable.HeaderXUserID))
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("Invalid user."))
	}
	err = json.Unmarshal(c.Body(), &dUser)
	if err != nil {
		return c.Status(400).JSON(mymodels.MyResponse("data invalid format."))
	}
	failMSG, err := services.UserAddStarterBalance(userID, int(dUser.CurrencyID.Int64), dUser.BalanceStart.Decimal)
	if err != nil {
		return err
	}
	if failMSG != "" {
		return c.Status(400).JSON(mymodels.MyResponse(failMSG))
	}
	return c.Status(200).JSON(mymodels.MyResponse("Add data success."))
}
