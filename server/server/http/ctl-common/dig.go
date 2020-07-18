package ctl_common

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewCommonCtrl)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCommonCtrl)
}
