package service

import (
	"fmt"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/module/module-statistic/model"
	"testing"
)

func TestStatisticService_HandleBrokerWebStationVisited(t *testing.T) {
	dig.Container.MustInvoke(func(svc *StatisticService) {
		ctx := context2.NewContext()
		list := []*model.StatVisitorLogs{}
		err := svc.Dao.GetDb(ctx).Model(new(model.AddStatVisitorLogsModel)).Find(&list, "city=''").Error
		if err != nil {
			t.Log(err)
			return
		}
		for _, logs := range list {
			if logs.City == "" {
				fmt.Println("ip " + logs.Ip)
				info, err := svc.ipInfoApi.GetIpInfo(ctx, logs.Ip)
				if err != nil {
					t.Log(err)
					return
				}
				logs.RegionName = info.RegionName
				logs.City = info.City
				logs.Memo = info.RawResponse
				err = svc.Dao.GetDb(ctx).Model(new(model.StatVisitorLogs)).Update(logs).Error
				if err != nil {
					t.Log(err)
					return
				}
				fmt.Println("ip done " + logs.Ip)
			}
		}
	})
}
