package sms

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewAliSms2(cfg *ncfg.NConfig) IAliSms {
	smsCfg := cfg.Cfg.AliSms
	return NewAliSms(smsCfg.Key, smsCfg.Secret)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewAliSms2)
}
