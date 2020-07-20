// +build !auto_test

package sms

import (
	"context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	ncfg.DigProviderByCfgStrFromEnv()
	m.Run()
}

func TestSendSms(t *testing.T) {
	dig.Container.MustInvoke(func(cfg *ncfg.NConfig) {
		sms := NewAliSms(cfg.Cfg.AliSms.Key, cfg.Cfg.AliSms.Secret)
		err := sms.SendRegisterSms(context.Background(), "11232123", "11112222223333333333")
		assert.Nil(t, err, err)
	})
}
