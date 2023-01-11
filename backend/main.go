package main

import (
	"broker-exchange/libs/dbs"
	"broker-exchange/myconfig/myserver"
	"broker-exchange/myconfig/myvariable"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found")
	}
	myvariable.SetEnv()

	dbs.PostgresConnect(myvariable.ENV_DB_PROJECT)

	app := myserver.New()
	myserver.SetUpMockData()
	myserver.SetRoutes(app)
	myserver.RunServe(app, myvariable.ENV_PORT)
}
