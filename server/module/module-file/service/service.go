package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/module-file/dao"
	oss_s3 "github.com/krilie/lico_alone/module/module-file/third-api"
)

type Service struct {
	Dao       *dao.Dao
	FileSaver oss_s3.FileOperator
}

func (a *Service) NewWithTx(ctx context.Context, tx *gorm.DB) (service cdb.Service, err error) {
	return &Service{
		Dao:       &dao.Dao{Dao: &cdb.Dao{Db: tx}},
		FileSaver: a.FileSaver,
	}, nil
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Dao:       dao.NewDao(cfg.DB),
		FileSaver: oss_s3.NewOssLocal(cfg),
	}
}
