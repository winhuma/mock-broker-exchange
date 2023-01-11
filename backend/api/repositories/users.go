package repositories

import (
	"broker-exchange/libs/dbs"
	"broker-exchange/myconfig/mymodels"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func UserGetByUsername(username string) (mydata mymodels.DBUser, err error) {
	mydb := dbs.GetInstanceDB()
	err = mydb.Table(
		mymodels.DBUser.TableName(mymodels.DBUser{})).
		Where("username=?", username).Scan(&mydata).Error
	if err != nil {
		return mydata, err
	}
	return mydata, nil
}

func UserCreate(userData mymodels.DBUser) (int, error) {
	mydb := dbs.GetInstanceDB()
	err := mydb.Table(mymodels.DBUser.TableName(mymodels.DBUser{})).
		Create(&userData).Error
	if err != nil {
		return userData.ID, err
	}
	return userData.ID, nil
}

func UserBalanceGetByUserID(userID int) ([]mymodels.DBUserBalace, error) {
	var mydata []mymodels.DBUserBalace
	mydb := dbs.GetInstanceDB()
	err := mydb.Raw(`select ub.*, c.name as currency_name
		from user_balance ub
		join currency c on c.id=ub.currency_id
		where ub.user_id=?`, userID).Scan(&mydata).Error
	if err != nil {
		return mydata, err
	}
	return mydata, nil
}

func UserBalanceNewCurrency(tx *gorm.DB, userID int, currencyID int, balance decimal.Decimal) error {
	mydb := dbs.GetInstanceDB()
	if tx != nil {
		mydb = tx
	}
	err := mydb.Table(
		mymodels.DBUserBalace.TableName(mymodels.DBUserBalace{})).
		Create(&mymodels.DBUserBalace{
			UserID:     userID,
			CurrencyID: currencyID,
			Balance:    balance,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func UserBalanceUpdate(tx *gorm.DB, userID int, currencyID int, balance decimal.Decimal) error {
	mydb := dbs.GetInstanceDB()
	if tx != nil {
		mydb = tx
	}
	err := mydb.Table(
		mymodels.DBUserBalace.TableName(mymodels.DBUserBalace{})).
		Where("user_id=? and currency_id=?", userID, currencyID).
		Updates(mymodels.DBUserBalace{Balance: balance}).Error
	if err != nil {
		return err
	}
	return nil
}
