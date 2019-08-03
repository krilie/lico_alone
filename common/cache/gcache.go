package cache

import (
	"github.com/bluele/gcache"
)

var GCache gcache.Cache

func init() {
	// 可以缓存的条数 2万
	GCache = gcache.New(20000).LRU().Build()
}
