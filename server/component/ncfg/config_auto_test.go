// +build auto_test

package ncfg

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	var defaultCfg = `
[ali_sms]
  key = "EEVVEEsss&&8"
  secret = ""
`
	DigProviderByCfgStr(defaultCfg)
	m.Run()
}

func TestNewNConfig2(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *NConfig) {
		assert.Equal(t, "EEVVEEsss&&8", cfg.Cfg.AliSms.Key, "not equal")
	})
}
