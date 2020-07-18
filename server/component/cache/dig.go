package cache

import "github.com/krilie/lico_alone/common/dig"

func init() {
	dig.Container.MustProvide(NewCache)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCache)
}
