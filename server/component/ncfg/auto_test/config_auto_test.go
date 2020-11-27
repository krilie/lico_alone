package auto_test

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	var defaultCfg = `
[ali_sms]
  key = "EEVVEEsss&&8"
  secret = ""
`
	ncfg.DigProviderByCfgStr(defaultCfg)
	m.Run()
}

func TestAutoNewNConfig2(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *ncfg.NConfig) {
		assert.Equal(t, "EEVVEEsss&&8", cfg.Cfg.AliSms.Key, "not equal")
	})
}
