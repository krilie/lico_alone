package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/module-user/dao"
)

type Service struct {
	Dao *dao.Dao
}

func (s *Service) NewWithTx(ctx context.Context, tx *gorm.DB) (cdb.Service, error) {
	return &Service{
		Dao: &dao.Dao{Dao: &cdb.Dao{Db: tx}},
	}, nil
}

func (s *Service) GetDb(ctx context.Context) *gorm.DB {
	return s.Dao.Db
}

func NewService(cfg config.DB) *Service {
	return &Service{
		Dao: dao.NewDao(cfg),
	}
}
