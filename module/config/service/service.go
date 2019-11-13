package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/config/dao"
)

// 系统配置服务
type Service struct {
	Dao *dao.Dao
}

func (a *Service) SetTx(ctx context.Context, tx *gorm.DB) (cdb.Service, error) {
	var log = clog.NewLog(ctx, "module/user/service/service.go.Service", "WithTx")
	log.Debug("new tx")
	txDao, err := a.Dao.Begin(tx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &Service{
		Dao: &dao.Dao{Dao: txDao},
	}, err
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.DB) *Service {
	return &Service{
		Dao: dao.NewDao(cfg),
	}
}
