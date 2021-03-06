package service

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/dao"
	"time"
)

type UserModule struct {
	Dao            *dao.UserDao
	log            *nlog.NLog
	jwtSecret      []byte
	jwtExpDuration time.Duration
}

func NewUserModule(dao *dao.UserDao, log *nlog.NLog, cfg *ncfg.NConfig) *UserModule {

	var jwtCfg = cfg.GetJwtCfg()

	log = log.WithField(context_enum.Module.Str(), "module user service")
	return &UserModule{
		Dao:            dao,
		log:            log,
		jwtSecret:      []byte(jwtCfg.HS256key),
		jwtExpDuration: time.Duration(jwtCfg.NormalExpDuration) * time.Second,
	}
}
