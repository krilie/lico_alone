package service_common

import (
	"github.com/krilie/lico_alone/common/dig"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewCommonService)
}
