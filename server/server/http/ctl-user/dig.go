package ctl_user

import (
	"github.com/krilie/lico_alone/common/dig"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewUserCtrl)
}
