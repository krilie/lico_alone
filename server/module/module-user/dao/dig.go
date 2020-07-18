package dao

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewUserDao)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewUserDao)
}
