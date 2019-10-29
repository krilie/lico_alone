package sms

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestSendSms(t *testing.T) {
	t.Log(str_util.ToJsonPretty(config.Cfg))
	sms := NewAliSms(config.Cfg.AliSms.Key, config.Cfg.AliSms.Secret)
	err := sms.SendRegisterSms(context.Background(), "18761438228", "33456")
	t.Log(err)
}
