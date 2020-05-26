package common_service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestCommonService_GetIcpInfo(t *testing.T) {
	dig.Container.MustInvoke(func(svc *CommonService) {
		info := svc.GetIcpInfo(context.NewContext())
		t.Log(str_util.ToJsonPretty(info))
	})
}
