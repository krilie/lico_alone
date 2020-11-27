package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-customer/dao"
)

type DynamicShareModule struct {
	Dao *dao.CustomerDao
	log *nlog.NLog
}

func NewDynamicShareModule(dao *dao.CustomerDao, log *nlog.NLog) *DynamicShareModule {
	var svc = &DynamicShareModule{
		Dao: dao,
		log: log,
	}
	return svc
}
