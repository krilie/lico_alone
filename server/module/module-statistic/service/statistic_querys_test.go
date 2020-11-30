package service

import (
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"testing"
)

func TestStatisticService_QueryAllVisitorLonLat(t *testing.T) {
	container.MustInvoke(func(svc *StatisticService) {
		ctx := context2.NewContext()
		lat, err := svc.QueryAllVisitorLonLat(ctx)
		t.Log(err)
		t.Log(str_util.ToJsonPretty(lat), len(lat))
	})
}
