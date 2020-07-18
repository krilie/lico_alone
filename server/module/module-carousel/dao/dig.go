package dao

import "github.com/krilie/lico_alone/common/dig"

func init() {
	dig.Container.MustProvide(NewCarouseDao)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCarouseDao)
}
