package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

func init() {
	dig.Container.MustProvide(func(dao *dao.UserDao, log *nlog.NLog) *UserService {
		return NewUserService(dao, log)
	})
}
