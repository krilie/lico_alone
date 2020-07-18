package ctl_user

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewUserCtrl)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewUserCtrl)
}
