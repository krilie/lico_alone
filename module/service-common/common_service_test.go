package service_common

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/jsonutil"
	"github.com/krilie/lico_alone/component"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvide(NewCommonService)

func TestCommonService_GetIcpInfo(t *testing.T) {
	container.MustInvoke(func(svc *CommonService) {
		info := svc.GetIcpInfo(context.EmptyAppCtx())
		t.Log(jsonutil.ToJsonPretty(info))
	})
}

func TestHello(t *testing.T) {
	t.Log("hello")
}
