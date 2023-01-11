package services

import (
	"broker-exchange/api/repositories"
	"broker-exchange/myconfig/mymodels"
	"broker-exchange/myconfig/myvariable"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func OrderCreate(userID int, userData mymodels.BodyOrderCreate) error {
	userBalance, err := repositories.UserBalanceGetByUserID(userID)
	if err != nil {
		return err
	}
	var mapUserBalance = map[int]mymodels.DBUserBalace{}
	for _, balance := range userBalance {
		mapUserBalance[balance.CurrencyID] = balance
	}

	dMyCurrency, err := repositories.CurrencyGetByID(userData.TargetCurrencyID)
	if err != nil {
		return err
	}
	dTargetCurrency, err := repositories.CurrencyGetByID(userData.TargetCurrencyID)
	if err != nil {
		return err
	}

	myTx := repositories.BeginTransaction()
	err = UpdateMybalance(myTx, userID, userData.Action, mapUserBalance[userData.MyCurrencyID], userData)
	if err != nil {
		return err
	}

	err = UpdateTargetBalance(myTx, userID, userData, mapUserBalance[userData.TargetCurrencyID], dMyCurrency.ValueUSD, dTargetCurrency.ValueUSD)
	if err != nil {
		return err
	}
	return myTx.Commit().Error
}

// ##############################################
func UpdateMybalance(mytx *gorm.DB, userID int, action string, mycurrencyUpdate mymodels.DBUserBalace, userData mymodels.BodyOrderCreate) error {
	if action == myvariable.VarActionType.BUY {
		mycurrencyUpdate.Balance = mycurrencyUpdate.Balance.Sub(userData.MyCurrencyValue.Decimal)
	} else if action == myvariable.VarActionType.SALE {
		mycurrencyUpdate.Balance = mycurrencyUpdate.Balance.Add(userData.MyCurrencyValue.Decimal)
	}
	err := repositories.UserBalanceUpdate(mytx, userID, userData.MyCurrencyID, mycurrencyUpdate.Balance)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTargetBalance(mytx *gorm.DB, userID int, userData mymodels.BodyOrderCreate, userTargetBalance mymodels.DBUserBalace, myrate decimal.Decimal, targetRate decimal.Decimal) error {
	var err error
	calRateTarget := userData.MyCurrencyValue.Decimal.Mul(myrate)
	ValueUpdate := targetRate.Div(calRateTarget)

	if userTargetBalance.ID == 0 {
		err = repositories.UserBalanceNewCurrency(mytx, userID, userData.TargetCurrencyID, ValueUpdate)
	} else {
		if userData.Action == myvariable.VarActionType.BUY {
			ValueUpdate = userTargetBalance.Balance.Add(ValueUpdate)
		} else if userData.Action == myvariable.VarActionType.SALE {
			ValueUpdate = userTargetBalance.Balance.Sub(ValueUpdate)
		}
		err = repositories.UserBalanceUpdate(mytx, userID, userData.TargetCurrencyID, ValueUpdate)
	}
	return err
}
