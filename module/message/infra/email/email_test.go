package email

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := config.LoadConfigByFile(`e://config.yaml`)
	t.Log(err)
	cfg := config.Cfg.Email
	email := NewEmail(cfg.Address, cfg.Host, cfg.Port, cfg.UserName, cfg.Password)
	err = email.SendEmail(context.Background(), "776334655@qq.com", "bbbbb", "测试消息")
	t.Log(err)
}
