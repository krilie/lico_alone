package cache

import (
	"github.com/bluele/gcache"
	"github.com/krilie/lico_alone/common/errs"
)

var GCache gcache.Cache

func init() {
	// 可以缓存的条数 2万
	GCache = gcache.New(20000).LRU().Build()
}

// cache
func WithCache(cache gcache.Cache, prefix string, key string, fn func(key string) (value interface{}, err error)) (interface{}, error) {
	get, err := cache.Get(prefix + key)
	if err == nil {
		return get, nil
	}
	if err != gcache.KeyNotFoundError {
		return nil, errs.NewInternal().WithError(err)
	}
	value, err := fn(key)
	if err != nil {
		return nil, err
	}
	err = cache.Set(prefix+key, value)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return value, nil
}
