package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/model"
)

type messageSms struct {
	*ndb.NDb
	log *nlog.NLog
}

type IMessageSms interface {
	CreateMessageSms(ctx context.Context, item *model.MessageSms) error
	UpdateMessageSms(ctx context.Context, item *model.MessageSms) error
	DeleteMessageSms(ctx context.Context, id string) error
	GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error)
}

func (d *messageSms) CreateMessageSms(ctx context.Context, item *model.MessageSms) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageSms) UpdateMessageSms(ctx context.Context, item *model.MessageSms) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Updates(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageSms) DeleteMessageSms(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageSms{}).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageSms) GetMessageSmsById(ctx context.Context, id string) (*model.MessageSms, error) {
	panic("implement me")
}
