package service

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewMessageModule)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewMessageModule)
}
