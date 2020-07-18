package service_init_data

import (
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(NewInitDataService)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewInitDataService)
}
