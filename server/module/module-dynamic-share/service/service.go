package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-dynamic-share/dao"
)

type DynamicShareModule struct {
	Dao *dao.DynamicShareDao
	log *nlog.NLog
}

func NewDynamicShareModule(dao *dao.DynamicShareDao, log *nlog.NLog) *DynamicShareModule {
	var svc = &DynamicShareModule{
		Dao: dao,
		log: log,
	}
	return svc
}
