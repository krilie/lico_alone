package service

import (
	"context"
	"github.com/krilie/lico_alone/common/config"
	"testing"
)

func TestService_ChangeUserPassword(t *testing.T) {
	srv := NewService(config.Cfg)
	err := srv.ChangeUserPassword(context.Background(), "1162588057191321600", "123", "11234")
	t.Log(err)
}
