package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

type UserModule struct {
	Dao *dao.UserDao
	log *nlog.NLog
}

func NewUserModule(dao *dao.UserDao, log *nlog.NLog) *UserModule {
	log = log.WithField(context_enum.Module.Str(), "module user service")
	return &UserModule{
		Dao: dao,
		log: log,
	}
}
