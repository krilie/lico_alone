package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestService_ChangeUserPassword(t *testing.T) {
	srv := NewUserService(config.Cfg.DB)
	err := srv.ChangeUserPassword(context.Background(), "1162588057191321600", "123", "11234")
	t.Log(err)
}

func TestService_UserLogin(t *testing.T) {
	srv := NewUserService(config.Cfg.DB)
	jwt, err := srv.UserLogin(context.Background(), "123", "123", "123")
	t.Log(jwt, err)
}

func TestService_RegisterNewUser(t *testing.T) {
	srv := NewUserService(config.Cfg.DB)
	err := srv.RegisterNewUser(context.Background(), "123", "123")
	t.Log(err)
}
