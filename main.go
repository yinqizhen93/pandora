package main

import (
	"pandora/config"
	"pandora/db"
	"pandora/logs"
	"pandora/router"
	"pandora/service"
)

func main() {
	err := router.Router.Run(":5001")
	if err != nil {
		panic(err)
	}
}

func init() {
	config.InitConfig()
	logs.InitLogger()
	db.InitDB()
	service.InitService()
	router.InitRouter()
}
