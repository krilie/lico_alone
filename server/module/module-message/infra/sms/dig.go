package sms

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewAliSms2(cfg *ncfg.NConfig) IAliSms {
	smsCfg := cfg.Cfg.AliSms
	return NewAliSms(smsCfg.Key, smsCfg.Secret)
}

func init() {
	dig.Container.MustProvide(NewAliSms2)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewAliSms2)
}
