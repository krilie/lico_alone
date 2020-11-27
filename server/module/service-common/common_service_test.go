package service_common

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest() // 数据库 配置 消息队列 等
	module.DigProviderModule()  // 服务 模块 包括其dao层
	DigProvider()
}

func TestCommonService_GetIcpInfo(t *testing.T) {
	dig.Container.MustInvoke(func(svc *CommonService) {
		info := svc.GetIcpInfo(context.NewContext())
		t.Log(str_util.ToJsonPretty(info))
	})
}

func TestHello(t *testing.T) {
	t.Log("hello")
}
