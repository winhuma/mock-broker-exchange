package mymodels

import "github.com/shopspring/decimal"

type BodyOrderCreate struct {
	MyCurrencyID     int                 `json:"my_currency_id"`
	TargetCurrencyID int                 `json:"target_currency_id"`
	MyCurrencyValue  decimal.NullDecimal `json:"my_currency_value"`
	Action           string              `json:"action"`
}

type ActionType struct {
	SALE string
	BUY  string
}

type DBOrder struct {
	ID int `json:"id"`
}
