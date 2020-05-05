package cache

import "github.com/krilie/lico_alone/common/dig"

func init() {
	dig.MustProvide(func() *Cache {
		return NewCache()
	})
}
