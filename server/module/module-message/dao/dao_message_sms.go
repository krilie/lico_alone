package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-message/model"
)

type IMessageSms interface {
	CreateMessageSms(ctx context.Context, item *model.MessageSms) error
	UpdateMessageSms(ctx context.Context, item *model.MessageSms) error
	DeleteMessageSms(ctx context.Context, id string) error
	GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error)
}

func (d *MessageDao) CreateMessageSms(ctx context.Context, item *model.MessageSms) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) UpdateMessageSms(ctx context.Context, item *model.MessageSms) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) DeleteMessageSms(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageSms{}).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error) {
	panic("implement me")
}
