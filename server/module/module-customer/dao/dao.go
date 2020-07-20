package dao

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
)

type CustomerDao struct {
	*ndb.NDb
	log *nlog.NLog
}

func NewCustomerDao(db *ndb.NDb, log *nlog.NLog) *CustomerDao {
	log = log.WithField(context_enum.Module.Str(), "module CustomerDao dao")
	return &CustomerDao{
		NDb: db,
		log: log,
	}
}
