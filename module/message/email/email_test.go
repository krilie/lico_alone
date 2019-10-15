package email

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestSendEmail(t *testing.T) {
	cfg := config.Cfg.Email
	email := NewEmail(cfg.Address, cfg.Host, cfg.Port, cfg.UserName, cfg.Password)
	err := email.SendEmail(context.Background(), "776334655@qq.com", "aa", "测试消息")
	t.Log(err)
}
