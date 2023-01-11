package dbs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MyDB *gorm.DB

func PostgresConnect(connectionString string) *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Panic(err)
	}
	MyDB = db
	return db
}

func GetInstanceDB() *gorm.DB {
	return MyDB
}
