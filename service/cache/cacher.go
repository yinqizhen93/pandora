package cache

import "github.com/google/wire"

type Cacher interface {
	Get(key string) (interface{}, bool)
	Set(schema []string, key string, val interface{}, cost int64)
	DelBySchema(schema string)
}

func NewCacher() Cacher {
	return NewMemCache()
}

var ProviderSet = wire.NewSet(NewCacher)
