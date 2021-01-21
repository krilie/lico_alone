package service

import (
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/strutil"
	"testing"
)

func TestStatisticService_QueryAllVisitorLonLat(t *testing.T) {
	container.MustInvoke(func(svc *StatisticService) {
		ctx := context2.EmptyAppCtx()
		lat, err := svc.QueryAllVisitorLonLat(ctx)
		t.Log(err)
		t.Log(strutil.ToJsonPretty(lat), len(lat))
	})
}
