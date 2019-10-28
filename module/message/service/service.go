package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/message/dao"
	"github.com/krilie/lico_alone/module/message/infra/email"
	"github.com/krilie/lico_alone/module/message/infra/sms"
)

type Service struct {
	Dao *dao.Dao
	email *email.Email
	sms *sms.AliSms
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
		Dao: &dao.Dao{Dao: txDao},
	}, err
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Dao: dao.NewDao(cfg.DB),
		email:email.NewEmail(cfg.Email.Address,cfg.Email.Host,cfg.Email.Port,cfg.Email.UserName,cfg.Email.Password),
		sms:sms.NewAliSms()
	}
}
