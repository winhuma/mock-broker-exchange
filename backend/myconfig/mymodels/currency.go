package mymodels

import (
	"github.com/shopspring/decimal"
)

type DBCurrency struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	ValueUSD decimal.Decimal `json:"value_usd"`
}

func (DBCurrency) TableName() string {
	return "currency"
}
