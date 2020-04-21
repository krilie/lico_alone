package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	auth_cache "github.com/krilie/lico_alone/module/user/auth-cache"
	"github.com/krilie/lico_alone/module/user/dao"
)

type Service struct {
	Dao      *dao.Dao
	AuthRBAC *auth_cache.AuthCache
}

func (s *Service) NewWithTx(ctx context.Context, tx *gorm.DB) (cdb.Service, error) {
	return &Service{
		Dao:      &dao.Dao{Dao: &cdb.Dao{Db: tx}},
		AuthRBAC: s.AuthRBAC,
	}, nil
}

func (s *Service) GetDb(ctx context.Context) *gorm.DB {
	return s.Dao.Db
}

func NewService(cfg config.DB) *Service {
	return &Service{
		Dao:      dao.NewDao(cfg),
		AuthRBAC: auth_cache.NewAuthCache(),
	}
}

func (s *Service) WithTrans(ctx context.Context, oriService cdb.Service, txFunc func(ctx context.Context, service cdb.Service) error) (err error) {
	return cdb.WithTrans(ctx, oriService, func(ctx context.Context, service cdb.Service) error {
		return txFunc(ctx, service.(*Service))
	})
}
