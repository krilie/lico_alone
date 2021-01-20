package auto_test

import (
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoNewNConfig2(t *testing.T) {
	var defaultCfg = `
[ali_sms]
  key = "EEVVEEsss&&8"
  secret = ""
`
	cfg := ncfg.NewNConfigByCfgStr(defaultCfg).GetAliSmsCfg()
	assert.Equal(t, "EEVVEEsss&&8", cfg.Key, "not equal")
}
