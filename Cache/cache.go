package Cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Memory = cache.New(5*time.Minute, 10*time.Minute)

func Set(key string, value interface{}){
	Memory.Set(key, value, cache.DefaultExpiration)
}

func Get(key string) (interface{}, bool) {
	return Memory.Get(key)
}

