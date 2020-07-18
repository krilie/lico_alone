package ncfg

import "github.com/krilie/lico_alone/common/dig"

func init() {
	dig.Container.MustProvide(NewNConfig)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewNConfig)
}
