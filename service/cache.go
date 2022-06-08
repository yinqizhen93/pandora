   package service

import (
	"github.com/patrickmn/go-cache"
	"time"
)

func cache2() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)


	// ...
	// foo can then be passed around freely as a string

	// Want performance? Store pointers!
	//c.Set("foo", &MyStruct, cache.DefaultExpiration)
	//if x, found := c.Get("foo"); found {
	//	foo := x.(*MyStruct)
	//	// ...
	//}
}
