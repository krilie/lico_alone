package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

type UserService struct {
	Dao *dao.UserDao
	log *nlog.NLog
}

func NewUserService(dao *dao.UserDao, log *nlog.NLog) *UserService {
	log = log.WithField(context_enum.Module.Str(), "module user service")
	return &UserService{
		Dao: dao,
		log: log,
	}
}
