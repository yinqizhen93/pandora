package service

import (
	"pandora/service/cache"
)

func InitService() {
	InitCasbin()
	cache.InitCache()
}
