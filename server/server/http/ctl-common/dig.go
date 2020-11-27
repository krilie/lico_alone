package ctl_common

import (
	"github.com/krilie/lico_alone/common/appdig"
)

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewCommonCtrl)
}
