package cache

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

var Cache *MemCache

func InitCache() {
	Cache = NewMemCache()
}

type MemCache struct {
	Cache *ristretto.Cache
	urls  map[string][]string // url :[]schema
}

//var Cache *ristretto.Cache

func NewMemCache() *MemCache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	c := MemCache{
		Cache: cache,
		urls:  make(map[string][]string),
	}
	return &c
}

func (c *MemCache) Set(schema []string, key string, val interface{}, cost int64) {
	c.urls[key] = schema
	c.Cache.Set(key, val, cost) // todo cost 需要动态设置
	c.Cache.Wait()
}

func (c *MemCache) Get(key string) (interface{}, bool) {
	return c.Cache.Get(key)
}

func (c *MemCache) DelBySchema(schema string) {
	fmt.Printf("start to delete cache data for %s\n", schema)
	fmt.Println(c)
	for url, schemas := range c.urls {
		for _, s := range schemas {
			if s == schema {
				c.Cache.Del(url)
				break
			}
		}
	}
}
