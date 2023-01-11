package myvariable

import "broker-exchange/myconfig/mymodels"

var (
	HeaderXUserID = "X-USER-ID"
)

var VarActionType = mymodels.ActionType{
	SALE: "SALE",
	BUY:  "BUY",
}
