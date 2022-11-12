//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"pandora/api/handler"
	mdw "pandora/middleware"
	"pandora/router"
	"pandora/service/cache"
	"pandora/service/config"
	"pandora/service/db"
	"pandora/service/logger"
	"pandora/service/sse"
	"pandora/service/websocket"
)

func initApp(addr ...string) *App {
	panic(wire.Build(
		db.ProviderSet,
		logger.ProviderSet,
		handler.ProviderSet,
		router.ProviderSet,
		mdw.ProviderSet,
		cache.ProviderSet,
		config.ProviderSet,
		sse.ProviderSet,
		ws.ProviderSet,
		//access.ProviderSet,
		NewApp,
	))
}
