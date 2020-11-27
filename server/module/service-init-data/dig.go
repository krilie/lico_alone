package service_init_data

import (
	"github.com/krilie/lico_alone/common/appdig"
)

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewInitDataService)
}
