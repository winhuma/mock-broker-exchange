package myserver

import (
	"broker-exchange/libs/dbs"
	"broker-exchange/myconfig/mymodels"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/shopspring/decimal"
)

func SetUpMockData() {
	SetUpUsers()
	SetUpUserBalance()
	SetUpCurrency()
}

func SetUpUsers() {
	mydb := dbs.GetInstanceDB()
	mydb.AutoMigrate(&mymodels.DBUser{})
	fmt.Println("[PROCESS] SetUpUsers success")
}

func SetUpUserBalance() {
	mydb := dbs.GetInstanceDB()
	mydb.AutoMigrate(&mymodels.DBUserBalace{})
	fmt.Println("[PROCESS] SetUpUserBalance success")
}

func SetUpCurrency() {
	mydb := dbs.GetInstanceDB()
	mydb.AutoMigrate(&mymodels.DBCurrency{})

	var countCurrency int64
	err := mydb.Table(mymodels.DBCurrency.TableName(mymodels.DBCurrency{})).Count(&countCurrency).Error
	if err != nil {
		log.Fatal(err)
	}
	if countCurrency == 0 {
		var mydata []mymodels.DBCurrency
		data := ReadCSV("utils/mockdata/currency.csv")
		for index, row := range data {
			if index == 0 {
				continue
			}
			valueUSD, _ := decimal.NewFromString(row[1])
			var rowData = mymodels.DBCurrency{
				Name:     row[0],
				ValueUSD: valueUSD,
			}
			mydata = append(mydata, rowData)

		}
		err = mydb.Table(mymodels.DBCurrency.TableName(mymodels.DBCurrency{})).Create(&mydata).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("[PROCESS] SetUpCurrency success")
}

func ReadCSV(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}
