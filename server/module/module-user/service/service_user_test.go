package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"testing"
)

func TestUserService_RegisterNewUser(t *testing.T) {
	appdig.Container.MustInvoke(func(svc *UserModule) {
		err := svc.RegisterNewUser(context.NewContext(), "123", "123456")
		t.Log(err)
	})
}
