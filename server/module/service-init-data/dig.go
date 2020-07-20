package service_init_data

import (
	"github.com/krilie/lico_alone/common/dig"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewInitDataService)
}
