package services

import (
	"broker-exchange/api/repositories"
	"broker-exchange/myconfig/mymodels"
	"broker-exchange/myconfig/myvariable"
	"fmt"

	"github.com/shopspring/decimal"
)

func OrderCreate(userID int, userData mymodels.BodyOrderCreate) (string, error) {

	var failMSG string
	var mapUserCurrBalance = map[int]mymodels.DBUserBalace{}
	var mapCurrency = map[int]mymodels.DBCurrency{}
	userBalance, err := repositories.UserBalanceGetByUserID(userID)
	if err != nil {
		return failMSG, err
	}
	dCurrency, err := repositories.CurrencyGetAll()
	if err != nil {
		return failMSG, err
	}
	for _, balance := range userBalance {
		mapUserCurrBalance[balance.CurrencyID] = balance
	}
	for _, currency := range dCurrency {
		mapCurrency[currency.ID] = currency
	}

	if userData.Action == myvariable.VarActionType.BUY {
		failMSG, err = OrderBuy(userID, userData, mapUserCurrBalance, mapCurrency)
	} else {
		failMSG, err = OrderSale(userID, userData, mapUserCurrBalance, mapCurrency)
	}

	return failMSG, err
}

// ##############################################
func OrderBuy(userID int, userData mymodels.BodyOrderCreate, mapUserCurrBalance map[int]mymodels.DBUserBalace, mapCurrency map[int]mymodels.DBCurrency) (string, error) {
	var failMSG string
	var err error
	var mybalance = userData.MyCurrencyValue.Decimal
	var rateUserSelect = mapCurrency[userData.MyCurrencyID].ValueUSD
	var ratetarget = mapCurrency[userData.TargetCurrencyID].ValueUSD
	var calRate = ratetarget.Div(rateUserSelect)
	var myvalueSelectUpdate decimal.Decimal

	myTx := repositories.BeginTransaction()

	myvalueSelectUpdate = mybalance.Div(calRate)
	dUpdate := mapUserCurrBalance[userData.MyCurrencyID].Balance.Sub(mybalance)
	if dUpdate.LessThan(decimal.NewFromInt(0)) {
		failMSG = fmt.Sprintf("%s not enaugh", mapCurrency[userData.MyCurrencyID].Name)
		return failMSG, err
	}
	err = repositories.UserBalanceUpdate(myTx, userID, userData.MyCurrencyID, dUpdate)
	if err != nil {
		return failMSG, err
	}

	if mapUserCurrBalance[userData.TargetCurrencyID].ID == 0 {
		err = repositories.UserBalanceNewCurrency(nil, userID, userData.TargetCurrencyID, myvalueSelectUpdate)
		if err != nil {
			return failMSG, err
		}
	} else {
		dTargetUpdate := mapUserCurrBalance[userData.TargetCurrencyID].Balance.Add(myvalueSelectUpdate)
		err = repositories.UserBalanceUpdate(myTx, userID, userData.TargetCurrencyID, dTargetUpdate)
		if err != nil {
			return failMSG, err
		}
	}

	return failMSG, myTx.Commit().Error
}

func OrderSale(userID int, userData mymodels.BodyOrderCreate, mapUserCurrBalance map[int]mymodels.DBUserBalace, mapCurrency map[int]mymodels.DBCurrency) (string, error) {
	var failMSG string
	var err error
	var mybalance = userData.MyCurrencyValue.Decimal
	var rateUserSelect = mapCurrency[userData.MyCurrencyID].ValueUSD
	var ratetarget = mapCurrency[userData.TargetCurrencyID].ValueUSD
	var calRate = ratetarget.Div(rateUserSelect)
	var myvalueSelectUpdate decimal.Decimal

	myTx := repositories.BeginTransaction()

	myvalueSelectUpdate = mybalance.Mul(calRate)
	dTargetUpdate := mapUserCurrBalance[userData.TargetCurrencyID].Balance.Sub(mybalance)
	if dTargetUpdate.LessThan(decimal.NewFromInt(0)) {
		failMSG = fmt.Sprintf("%s not enaugh", mapCurrency[userData.MyCurrencyID].Name)
		return failMSG, err
	}
	err = repositories.UserBalanceUpdate(myTx, userID, userData.TargetCurrencyID, dTargetUpdate)
	if err != nil {
		return failMSG, err
	}
	dMeUpdate := mapUserCurrBalance[userData.MyCurrencyID].Balance.Add(myvalueSelectUpdate)
	err = repositories.UserBalanceUpdate(myTx, userID, userData.MyCurrencyID, dMeUpdate)
	if err != nil {
		return failMSG, err
	}
	return failMSG, myTx.Commit().Error
}
