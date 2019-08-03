package cache

import "testing"

func TestGCache(t *testing.T) {
	set := GCache.Set("ok", "ok2")
	if set != nil {
		t.Log(set)
	}
	get, set := GCache.Get("ok")
	t.Log(get, set)
}
