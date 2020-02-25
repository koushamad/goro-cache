package cache

import (
	"github.com/gadelkareem/cachita"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/app/helper"
	"sync"
	"time"
)

var (
	cache cachita.Cache
	once sync.Once
	)

func Boot() {
	once.Do(func() {
		var err error
		var driver = conf.Get("cache.driver")
		switch driver {
		case "file" :
			cache, err = cachita.NewFileCache(helper.CachePath(), 0, 0)
			throw.Fatal("Cache file driver error", 101, err)
			break
		case "redis":
			cache, err = cachita.Redis(conf.Get("redis.host") + ":" + conf.Get("redis.port"))
			throw.Fatal("Cache redis driver error", 102, err)
			break
		default:
			cache = cachita.Memory()
		}
	})
}

func Exists(key string) bool {
	return cache.Exists(key)
}

func Get(key string, i interface{}) {
	err := cache.Get(key, i)
	throw.Fatal("Cannot get cash", 103, err)
}

func Put(key string, i interface{}, ttl time.Duration){
	err := cache.Put(key, i, ttl*time.Minute);
	throw.Fatal("Cannot put cash", 104, err)
}

func Incr(key string, ttl time.Duration) int64 {
	data, err := cache.Incr(key, ttl)
	throw.Fatal("Cannot incr cash", 105, err)
	return data
}

func Invalidate(key string) {
	err := cache.Invalidate(key)
	throw.Fatal("Cannot invalidate cash",106, err)
}

