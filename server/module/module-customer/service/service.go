package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

type CustomerModule struct {
	Dao *dao.CustomerDao
	log *nlog.NLog
}

func NewCustomerModule(dao *dao.CustomerDao, log *nlog.NLog) *CustomerModule {
	return &CustomerModule{
		Dao: dao,
		log: log,
	}
}
