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

// logError 用h.logError替换 h.logger.Error 少些输入， 可以添加除写日志外其他东西
func (h Handler) logError(s string, kvs ...logger.Pair) {
	// add some other action
	h.logger.Error(s, kvs...)
}

func (h Handler) log(s string, kvs ...logger.Pair) {
	h.logger.Info(s, kvs...)
}
