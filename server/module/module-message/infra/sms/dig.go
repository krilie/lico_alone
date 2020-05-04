package sms

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(func(cfg *config.Config) IAliSms {
		smsCfg := cfg.AliSms
		return NewAliSms(smsCfg.Key, smsCfg.Secret)
	})
}
