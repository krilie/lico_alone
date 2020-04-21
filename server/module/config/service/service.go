package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/config/dao"
)

// 系统配置服务
type Service struct {
	Dao *dao.Dao
}

func (a *Service) NewWithTx(ctx context.Context, tx *gorm.DB) (cdb.Service, error) {
	return &Service{
		Dao: &dao.Dao{Dao: &cdb.Dao{Db: tx}},
	}, nil
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.DB) *Service {
	return &Service{
		Dao: dao.NewDao(cfg),
	}
}
