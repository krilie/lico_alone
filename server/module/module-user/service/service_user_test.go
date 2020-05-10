package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"testing"
)

func TestUserService_RegisterNewUser(t *testing.T) {
	dig.Container.MustInvoke(func(svc *UserService) {
		err := svc.RegisterNewUser(context.NewContext(), "123", "123456")
		t.Log(err)
	})
}
