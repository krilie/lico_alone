package sms

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
)

func init() {
	dig.Container.MustProvide(func(cfg *ncfg.NConfig) IAliSms {
		smsCfg := cfg.Cfg.AliSms
		return NewAliSms(smsCfg.Key, smsCfg.Secret)
	})
}
