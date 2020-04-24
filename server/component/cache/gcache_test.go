package cache

import (
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestGCache(t *testing.T) {
	dig.MustInvoke(func(cache *Cache) {
		_ = cache.Set("ok", "ok")
		get, err := cache.Get("ok")
		t.Log(get, err)
	})
}
