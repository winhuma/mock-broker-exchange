package services

import (
	"broker-exchange/api/repositories"
	"broker-exchange/myconfig/mymodels"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userData mymodels.BodyUserLogin) (string, string, error) {
	var mytoken string
	var failMSG string

	dUser, err := repositories.UserGetByUsername(userData.Username)
	if err != nil {
		return mytoken, failMSG, err
	}
	if dUser.ID == 0 {
		newUserID, failMSG, err := NewUser(userData)
		if err != nil || failMSG != "" {
			return mytoken, failMSG, err
		}
		dUser.ID = newUserID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(map[string]interface{}{
		"user_id":    fmt.Sprint(dUser.ID),
		"username":   userData.Username,
		"expires_at": time.Now().Add(5 * time.Minute).Unix(),
	}))

	mytoken, err = token.SignedString([]byte("MySignature"))
	if err != nil {
		return mytoken, failMSG, err
	}
	return mytoken, failMSG, nil
}

func NewUser(userData mymodels.BodyUserLogin) (int, string, error) {
	var failMSG string
	var err error
	var userID int
	if !userData.BalanceStart.Valid {
		failMSG = "Key amount not empty."
		return userID, failMSG, err
	}
	userID, err = repositories.UserCreate(mymodels.DBUser{
		Username: userData.Username,
	})
	if err != nil {
		return userID, failMSG, err
	}

	err = repositories.UserBalanceNewCurrency(nil, userID, int(userData.CurrencyID.Int64), userData.BalanceStart.Decimal)
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
