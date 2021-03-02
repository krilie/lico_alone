package auto_test

import (
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAutoNewNConfig2(t *testing.T) {
	var defaultCfg = `
[ali_sms]
  key = "EEVVEEsss&&8"
  secret = ""
`
	cfg := ncfg.NewNConfig()
	err := cfg.LoadFromConfigTomlStr(defaultCfg)
	require.Nil(t, err)
	cfgAli := cfg.GetAliSmsCfg()
	assert.Equal(t, "EEVVEEsss&&8", cfgAli.Key, "not equal")
}
