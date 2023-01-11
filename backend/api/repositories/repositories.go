package repositories

import (
	"broker-exchange/libs/dbs"

	"gorm.io/gorm"
)

func BeginTransaction() *gorm.DB {
	return dbs.GetInstanceDB().Begin()
}
