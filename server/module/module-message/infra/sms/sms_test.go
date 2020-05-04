package sms

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendSms(t *testing.T) {
	sms := NewAliSms(config.Cfg.AliSms.Key, config.Cfg.AliSms.Secret)
	err := sms.SendRegisterSms(context.Background(), "11232123", "11112222223333333333")
	assert.Nil(t, err, err)
}
