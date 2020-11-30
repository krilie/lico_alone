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
	MustProvides(component.DigComponentProviderAllForTest)

func TestSendSms(t *testing.T) {
	container.MustInvoke(func(cfg *ncfg.NConfig) {
		sms := NewAliSms(cfg.Cfg.AliSms.Key, cfg.Cfg.AliSms.Secret)
		err := sms.SendRegisterSms(context.Background(), "3332234", "223")
		assert.Nil(t, err, err)
	})
}
