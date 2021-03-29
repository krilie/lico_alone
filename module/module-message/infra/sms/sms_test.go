package sms

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll)

func TestSendSms(t *testing.T) {
	container.MustInvoke(func(cfg *ncfg.NConfig) {
		aliSmsCfg := cfg.GetAliSmsCfg()
		sms := NewAliSms(aliSmsCfg.Key, aliSmsCfg.Secret)
		err := sms.SendRegisterSms(context.Background(), "3332234", "223")
		assert.Nil(t, err, err)
	})
}
