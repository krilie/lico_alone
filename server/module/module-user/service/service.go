package service

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

type UserService struct {
	Dao *dao.UserDao
	log *nlog.NLog
}

func NewUserService(dao *dao.UserDao, log *nlog.NLog) *UserService {
	return &UserService{
		Dao: dao,
		log: log,
	}
}
