package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

type ZUnionModule struct {
	Dao *dao.UserDao
	log *nlog.NLog
}

func NewZUnionModule(dao *dao.UserDao, log *nlog.NLog, cfg *ncfg.NConfig) *ZUnionModule {
	log = log.WithField(context_enum.Module.Str(), "module user service")
	return &ZUnionModule{
		Dao: dao,
		log: log,
	}
}
