package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/rbac"
)

type Service struct {
	Dao      *dao.Dao
	AuthRBAC *rbac.RBAC
}

func (s *Service) SetTx(ctx context.Context, tx *gorm.DB) (cdb.Service, error) {
	var log = clog.NewLog(ctx, "module/user/service/service.go.Service", "WithTx")
	log.Debug("new tx")
	txDao, err := s.Dao.Begin(tx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &Service{
		Dao:      &dao.Dao{Dao: txDao},
		AuthRBAC: s.AuthRBAC,
	}, err
}

func (s *Service) GetDb(ctx context.Context) *gorm.DB {
	return s.Dao.Db
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Dao:      dao.NewDao(&cfg),
		AuthRBAC: rbac.NewRBAC(),
	}
}
