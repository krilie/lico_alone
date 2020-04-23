package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/message/model"
)

type IMessageSms interface {
	CreateMessageSms(ctx context.Context, item *model.MessageSms) error
	UpdateMessageSms(ctx context.Context, item *model.MessageSms) error
	DeleteMessageSms(ctx context.Context, id string) error
	GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error)
}

func (d *Dao) CreateMessageSms(ctx context.Context, item *model.MessageSms) error {
	log := nlog.NewLog(ctx, "module/message/dao/dao_message_sms.go:17", "CreateMessageSms")
	err := d.Dao.Db.Create(item).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) UpdateMessageSms(ctx context.Context, item *model.MessageSms) error {
	log := nlog.NewLog(ctx, "module/message/dao/dao_message_sms.go:27", "UpdateMessageSms")
	err := d.Dao.Db.Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) DeleteMessageSms(ctx context.Context, id string) error {
	log := nlog.NewLog(ctx, "module/message/dao/dao_message_sms.go:37", "DeleteMessageSms")
	err := d.Dao.Db.Where("id=?", id).Delete(&model.MessageSms{}).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error) {
	panic("implement me")
}
