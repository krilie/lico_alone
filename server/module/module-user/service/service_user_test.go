package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAllForTest).
	MustProvides(DigModuleUserProviderAll)

func TestUserService_RegisterNewUser(t *testing.T) {
	container.MustInvoke(func(svc *UserModule) {
		err := svc.RegisterNewUser(context.EmptyAppCtx(), "123", "123456")
		t.Log(err)
	})
}
