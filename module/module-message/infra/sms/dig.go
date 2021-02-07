package sms

import (
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewAliSms2(cfg *ncfg.NConfig) IAliSms {
	var smsCfg = cfg.GetAliSmsCfg()
	return NewAliSms(smsCfg.Key, smsCfg.Secret)
}
