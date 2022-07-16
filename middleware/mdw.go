package middleware

import (
	"github.com/google/wire"
	"pandora/ent"
	"pandora/service/cache"
	"pandora/service/logger"
)

type Middleware struct {
	logger logger.Logger
	db     *ent.Client
	cache  *cache.MemCache // only init cache when use
}

func NewMiddleware(logger logger.Logger, db *ent.Client) *Middleware {
	return &Middleware{logger: logger, db: db}
}

var ProviderSet = wire.NewSet(NewMiddleware)
