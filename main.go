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

func main() {
	//err := router.Router.Run(":5001")
	app := initApp()
	err := app.run()
	if err != nil {
		panic(err)
	}
}
