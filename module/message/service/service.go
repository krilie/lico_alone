package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/module/message/dao"
	"github.com/krilie/lico_alone/module/message/infra/email"
	"github.com/krilie/lico_alone/module/message/infra/sms"
)

type Service struct {
	Dao   *dao.Dao
	email email.IEmail
	sms   sms.IAliSms
}

func (a *Service) NewWithTx(ctx context.Context, tx *gorm.DB) (service cdb.Service, err error) {
	return &Service{
		Dao:   &dao.Dao{Dao: &cdb.Dao{Db: tx}},
		email: a.email,
		sms:   a.sms,
	}, nil
}

func (a *Service) GetDb(ctx context.Context) *gorm.DB {
	return a.Dao.Db
}

func NewService(cfg config.Config) *Service {
	return &Service{
		Dao:   dao.NewDao(cfg.DB),
		email: email.NewEmail(cfg.Email.Address, cfg.Email.Host, cfg.Email.Port, cfg.Email.UserName, cfg.Email.Password),
		sms:   sms.NewAliSms(cfg.AliSms.Key, cfg.AliSms.Secret),
	}
}
