package mymodels

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type BodyUserLogin struct {
	Username     string              `json:"username"`
	BalanceStart decimal.NullDecimal `json:"balance_start"`
	CurrencyID   null.Int            `json:"currency_id"`
}

type DBUser struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
}

type ResUserBalance struct {
	ID           int             `json:"id"`
	UserID       int             `json:"user_id"`
	CurrencyID   int             `json:"currency_id"`
	CurrencyName string          `json:"currency_name"`
	Balance      decimal.Decimal `json:"balance"`
}

type DBUserBalace struct {
	ID         int             `json:"id" gorm:"primaryKey"`
	UserID     int             `json:"user_id"`
	CurrencyID int             `json:"currency_id"`
	Balance    decimal.Decimal `json:"balance"`
}

func (DBUser) TableName() string {
	return "users"
}

func (DBUserBalace) TableName() string {
	return "user_balance"
}
