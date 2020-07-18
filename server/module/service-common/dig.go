package service_common

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewCommonService)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCommonService)
}
