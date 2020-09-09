package dao

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-customer/model"
)

type CustomerDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewCustomerDao(db *ndb.NDb, log *nlog.NLog) *CustomerDao {
	log = log.WithField(context_enum.Module.Str(), "module CustomerDao dao")
	err := db.GetDb(context.NewContext()).AutoMigrate(new(model.CustomerAccount))
	if err != nil {
		panic(err)
	}
	return &CustomerDao{
		NDb: db,
		log: log,
	}
}
