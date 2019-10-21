package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestService_UserLogin(t *testing.T) {
	srv := NewService(config.Cfg.DB)
	jwt, err := srv.UserLogin(context.Background(), "123", "123", "123")
	t.Log(jwt, err)
}
