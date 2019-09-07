package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestService_RegisterNewUser(t *testing.T) {
	srv := NewService(config.Cfg)
	err := srv.RegisterNewUser(context.Background(), "123", "123")
	t.Log(err)
}
