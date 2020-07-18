package service

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewConfigModule)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewConfigModule)
}
