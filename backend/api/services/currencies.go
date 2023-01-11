package services

import "broker-exchange/api/repositories"

func CurrencyGetAll() (interface{}, error) {
	mydata, err := repositories.CurrencyGetAll()
	if err != nil {
		return nil, err
	}
	return mydata, nil
}
