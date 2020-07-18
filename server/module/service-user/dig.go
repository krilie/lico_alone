package service_user

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewUserService)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewUserService)
}
