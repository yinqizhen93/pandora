package config

import "github.com/google/wire"

type Config interface {
	Load()
	Get(key string) interface{}
	GetInt(key string) int
	GetString(key string) string
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	Watch()
}

func NewConfig() Config {
	return NewFileConfig()
}

var ProviderSet = wire.NewSet(NewConfig)
