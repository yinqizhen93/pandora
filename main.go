package main

import (
	"pandora/router"
	"pandora/service/config"
)

type App struct {
	addr   []string
	router *router.AppRouter
}

func NewApp(ae *router.AppRouter, conf config.Config) *App {
	addr := conf.GetString("server.address")
	return &App{
		addr:   []string{addr},
		router: ae,
	}
}

func (a App) run() error {
	return a.router.Run(a.addr...)
}

// @title           PANDORA API DOCUMENT
// @version         1.0
// @description     pandora api文档
// @host      localhost:5001
// @BasePath  /
// securityDefinitions.basic  BasicAuth -- 接口查看授权
func main() {
	//err := router.Router.Run(":5001")
	app := initApp()
	err := app.run()
	if err != nil {
		panic(err)
	}
}
