package ncfg

import "github.com/krilie/lico_alone/common/dig"

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewNConfig)
}

func DigProviderByCfgStr(cfgStr string) {
	dig.Container.MustProvide(NewNConfig)
	dig.Container.MustInvoke(func(nCfg *NConfig) {
		err := nCfg.LoadFromConfigTomlStr(cfgStr)
		if err != nil {
			panic(err)
		}
	})
}
