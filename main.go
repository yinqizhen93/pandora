package main

import (
	"pandora/router"
)

type App struct {
	addr   []string
	router *router.AppRouter
}

func NewApp(ae *router.AppRouter, addr ...string) *App {
	return &App{
		addr:   addr,
		router: ae,
	}
}

func (a App) run() error {
	return a.router.Run(a.addr...)
}

func main() {
	//err := router.Router.Run(":5001")
	addr := ":5001"
	app := initApp(addr)
	err := app.run()
	if err != nil {
		panic(err)
	}
}
