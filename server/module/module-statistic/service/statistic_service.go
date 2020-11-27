package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	infra_ip "github.com/krilie/lico_alone/common/thirdtools/infra-ip"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-statistic/dao"
)

type StatisticService struct {
	Dao       *dao.StatisticDao
	broker    *broker.Broker
	log       *nlog.NLog
	ipInfoApi infra_ip.IIpInfo
}

func NewStatisticService(broker *broker.Broker, Dao *dao.StatisticDao, log *nlog.NLog) *StatisticService {
	log = log.WithField(context_enum.Module.Str(), "StatisticService")
	var svc = &StatisticService{
		Dao:       Dao,
		broker:    broker,
		log:       log,
		ipInfoApi: infra_ip.NewIpInfoApiOne(),
	}
	broker.MustRegister(context.NewContext(), svc.HandleBrokerWebStationVisited)
	broker.MustRegister(context.NewContext(), svc.HandleBrokerArticleVisitorMessage)
	return svc
}

// DigProvider provider
func DigProvider() {
	appdig.Container.MustProvide(NewStatisticService)
}

func DigProviderAll() {
	dao.DigProvider()
	DigProvider()
}
