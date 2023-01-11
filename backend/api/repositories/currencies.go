package repositories

import (
	"broker-exchange/libs/dbs"
	"broker-exchange/myconfig/mymodels"
)

func CurrencyGetAll() ([]mymodels.DBCurrency, error) {
	var mydata []mymodels.DBCurrency
	mydb := dbs.GetInstanceDB()
	err := mydb.Table(mymodels.DBCurrency.TableName(mymodels.DBCurrency{})).Scan(&mydata).Error
	if err != nil {
		return nil, err
	}
	return mydata, nil
}

func CurrencyGetByID(currencyID int) (mymodels.DBCurrency, error) {
	var mydata mymodels.DBCurrency
	mydb := dbs.GetInstanceDB()
	err := mydb.Table(
		mymodels.DBCurrency.TableName(mymodels.DBCurrency{})).
		Where("id=?", currencyID).Scan(&mydata).Error
	if err != nil {
		return mydata, err
	}
	return mydata, nil
}
