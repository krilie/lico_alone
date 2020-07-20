package ctl_common

import (
	"github.com/krilie/lico_alone/common/dig"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCommonCtrl)
}
