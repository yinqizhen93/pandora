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
	db.InitDB2()
	service.InitService()
	router.InitRouter()
}

//func main() {
//	config.InitConfig()
//	dbc := db.InitDB2()
//	ctx := context.Background()
//	//user, _ := CreateCars(ctx, dbc)
//	//fmt.Println(user)
//	//err := QueryCars(ctx, user)
//	//if err := CreateGraph(ctx, dbc); err != nil {
//	//	panic(err)
//	//}
//
//	if err := QueryGithub(ctx, dbc); err != nil {
//		panic(err)
//	}
//
//	if err := QueryArielCars(ctx, dbc); err != nil {
//		panic(err)
//	}
//
//	if err := QueryGroupWithUsers(ctx, dbc); err != nil {
//		panic(err)
//	}
//
//}
