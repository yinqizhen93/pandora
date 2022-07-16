// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"pandora/api/handler"
	"pandora/middleware"
	"pandora/router"
	"pandora/service/cache"
	"pandora/service/db"
	"pandora/service/logger"
)

// Injectors from wire.go:

func initApp(addr ...string) *App {
	engine := router.NewEngine()
	loggerLogger := logger.NewLogger()
	cacher := cache.NewCacher()
	client := db.NewEntClient(cacher)
	handlerHandler := handler.NewHandler(loggerLogger, client)
	middlewareMiddleware := middleware.NewMiddleware(loggerLogger, client)
	appRouter := router.NewAppRouter(engine, handlerHandler, middlewareMiddleware)
	app := NewApp(appRouter, addr...)
	return app
}
