package config

import (
	"github.com/google/wire"
	"os"
)

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

var RootPath = getRootPath()

func getRootPath() string {
	if cp := os.Getenv("ROOTPATH"); cp != "" {
		return cp
	}
	return "."
}
