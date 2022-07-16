package handler

import (
	"github.com/google/wire"
	"pandora/ent"
	"pandora/service/logger"
)

type Handler struct {
	logger logger.Logger
	db     *ent.Client
}

func NewHandler(logger logger.Logger, db *ent.Client) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

var ProviderSet = wire.NewSet(NewHandler)
