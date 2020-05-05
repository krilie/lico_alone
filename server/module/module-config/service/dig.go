package service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-config/dao"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, dao *dao.ConfigDao) *ConfigService {
		return NewService(log, dao)
	})
}
