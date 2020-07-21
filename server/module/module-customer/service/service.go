package service

import (
	context2 "github.com/krilie/lico_alone/common/context"
	infra_ip "github.com/krilie/lico_alone/common/thirdtools/infra-ip"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

type CustomerModule struct {
	Dao   *dao.CustomerDao
	log   *nlog.NLog
	ipApi infra_ip.IIpInfo
}

func NewCustomerModule(broker *broker.Broker, dao *dao.CustomerDao, log *nlog.NLog) *CustomerModule {
	var svc = &CustomerModule{
		Dao:   dao,
		log:   log,
		ipApi: infra_ip.NewIpInfoApiOne(),
	}
	broker.MustRegister(context2.NewContext(), svc.HandleBrokerWebStationVisited)
	return svc
}
