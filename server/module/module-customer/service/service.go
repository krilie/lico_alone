package service

import (
	infra_ip "github.com/krilie/lico_alone/common/thirdtools/infra-ip"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

type CustomerModule struct {
	Dao   *dao.CustomerDao
	log   *nlog.NLog
	ipApi infra_ip.IIpInfo
}

func NewCustomerModule(dao *dao.CustomerDao, log *nlog.NLog) *CustomerModule {
	return &CustomerModule{
		Dao:   dao,
		log:   log,
		ipApi: infra_ip.NewIpInfoApiOne(),
	}
}
