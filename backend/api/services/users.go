package services

import (
	"broker-exchange/api/repositories"
	"broker-exchange/myconfig/mymodels"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shopspring/decimal"
)

func GenerateToken(userData mymodels.BodyUserLogin) (string, int, string, error) {
	var mytoken string
	var failMSG string

	dUser, err := repositories.UserGetByUsername(userData.Username)
	if err != nil {
		return mytoken, dUser.ID, failMSG, err
	}
	if dUser.ID == 0 {
		newID, _, err := CreateNewUser(userData)
		if err != nil {
			return mytoken, dUser.ID, failMSG, err
		}
		dUser.ID = newID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(map[string]interface{}{
		"user_id":    fmt.Sprint(dUser.ID),
		"username":   userData.Username,
		"expires_at": time.Now().Add(5 * time.Minute).Unix(),
	}))

	mytoken, err = token.SignedString([]byte("MySecret"))
	if err != nil {
		return mytoken, dUser.ID, failMSG, err
	}
	return mytoken, dUser.ID, failMSG, nil
}

func UserAddStarterBalance(userID int, currencyID int, balance decimal.Decimal) (string, error) {
	var failMSG string
	var err error
	mybalance, err := repositories.UserBalanceGetByUserID(userID)
	if err != nil {
		return failMSG, err
	}
	for _, b := range mybalance {
		if b.CurrencyID == currencyID {
			failMSG = "Currency is duplicate"
			return failMSG, err
		}
	}
	err = repositories.UserBalanceNewCurrency(nil, userID, currencyID, balance)
	if err != nil {
		return failMSG, err
	}
	return failMSG, nil
}

func CreateNewUser(userData mymodels.BodyUserLogin) (int, string, error) {
	var failMSG string
	var err error
	var userID int
	userID, err = repositories.UserCreate(mymodels.DBUser{
		Username: userData.Username,
	})
	if err != nil {
		return userID, failMSG, err
	}
	return userID, failMSG, err
}

func UserGetBalanceByUserID(userID int) (interface{}, error) {
	mybalance, err := repositories.UserBalanceGetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return mybalance, nil
}
