package middleware

import (
	"github.com/google/wire"
	"pandora/ent"
	"pandora/service/access"
	"pandora/service/cache"
	"pandora/service/logger"
	"pandora/service/sse"
	ws "pandora/service/websocket"
)

type Middleware struct {
	logger     logger.Logger
	db         *ent.Client
	cache      cache.Cacher
	accessCtrl access.RBAC  // only init rbac when use
	sse        *sse.SSEvent // only init when use
	ws         *ws.Hub      // only init when use
}

func NewMiddleware(logger logger.Logger, db *ent.Client, cache cache.Cacher, s *sse.SSEvent, ws *ws.Hub) *Middleware {
	return &Middleware{logger: logger, db: db, cache: cache, sse: s, ws: ws}
}

var ProviderSet = wire.NewSet(NewMiddleware)
