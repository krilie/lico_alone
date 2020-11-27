package cache

import (
	"errors"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	DigProvider()
	m.Run()
}

func TestAutoGCache(t *testing.T) {
	dig.Container.MustInvoke(func(cache *Cache) {
		err := cache.Set("ok", "ok")
		assert.Equal(t, nil, err, "should no error", err)
		get, err := cache.Get("ok")
		assert.Equal(t, nil, err, "should no error", err)
		assert.Equal(t, "ok", get, "should a ok str", errors.New("this is err"))
		t.Log(get, err)
	})
}
