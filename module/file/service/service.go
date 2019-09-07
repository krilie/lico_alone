package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/file/dao"
	oss_s3 "github.com/krilie/lico_alone/module/file/oss-s3"
)

type Service struct {
	Dao       *dao.Dao
	FileSaver oss_s3.FileOperator
}

func (a *Service) SetTx(ctx context.Context, tx *gorm.DB) (service cdb.Service, err error) {
	var log = clog.NewLog(ctx, "module/file/service/service.go.set_tx", "SetTx")
	log.Debug("new tx")
	txDao, err := a.Dao.Begin(tx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &Service{
		Dao:       &dao.Dao{Dao: txDao},
		FileSaver: a.FileSaver,
	}, err
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Dao:       dao.NewDao(&cfg),
		FileSaver: oss_s3.NewOssLocal(cfg),
	}
}
