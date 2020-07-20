// +build !auto_test

package email

import (
	"context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
	"testing"
)

func TestMain(m *testing.M) {
	ncfg.DigProviderByCfgStrFromEnv()
	m.Run()
}

func TestSendEmail(t *testing.T) {
	dig.Container.MustInvoke(func(cfg2 *ncfg.NConfig) {
		cfg := cfg2.Cfg.Email
		email := NewEmail(cfg.Address, cfg.Host, cfg.Port, cfg.UserName, cfg.Password)
		err := email.SendEmail(context.Background(), "776334655@qq.com", "bbbbb", "测试消息")
		t.Log(err)
	})

}
