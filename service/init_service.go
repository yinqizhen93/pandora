package service

import "pandora/service/cache"

func InitService() {
	//InitLogger()
	InitCasbin()
	cache.InitCache()
}
