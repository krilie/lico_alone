package sms

import (
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewAliSms2(cfg *ncfg.NConfig) IAliSms {
	smsCfg := cfg.Cfg.AliSms
	return NewAliSms(smsCfg.Key, smsCfg.Secret)
}
