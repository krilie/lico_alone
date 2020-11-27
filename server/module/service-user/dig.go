package service_user

import (
	"github.com/krilie/lico_alone/common/appdig"
)

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewUserService)
}
