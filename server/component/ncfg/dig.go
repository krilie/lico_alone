package ncfg

import (
	"github.com/krilie/lico_alone/common/dig"
	"os"
)

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewNConfig)
}

func DigProviderByCfgStr(cfgStr string) {
	dig.Container.MustProvide(func() *NConfig {
		cfg := NewNConfig()
		err := cfg.LoadFromConfigTomlStr(cfgStr)
		if err != nil {
			panic(err)
		}
		return cfg
	})
}

// json config
func DigProviderByCfgStrFromEnv() {
	dig.Container.MustProvide(func() *NConfig {
		cfg := NewNConfig()
		cfgStr := os.Getenv("MYAPP_TEST_CONFIG")
		err := cfg.LoadFromConfigJsonStr(cfgStr)
		if err != nil {
			panic(err)
		}
		return cfg
	})
}
