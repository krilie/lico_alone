package dao

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-message/model"
	"gorm.io/gorm"
)

type messageEmail struct {
	*ndb.NDb
	log *nlog.NLog
}

type IMessageEmail interface {
	CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	DeleteMessageEmail(ctx context.Context, id string) error
	GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error)
}

func (d *messageEmail) CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageEmail) UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Updates(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageEmail) DeleteMessageEmail(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageEmail{}).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageEmail) GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error) {
	item := &model.MessageEmail{}
	err := d.GetDb(ctx).Where("id=?", id).Find(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return item, nil
}
