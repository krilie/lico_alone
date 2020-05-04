package dao

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/message/model"
)

type IMessageEmail interface {
	CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	DeleteMessageEmail(ctx context.Context, id string) error
	GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error)
}

func (d *MessageDao) CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) DeleteMessageEmail(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageEmail{}).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error) {
	item := &model.MessageEmail{}
	err := d.GetDb(ctx).Where("id=?", id).Find(item).Error
	if err != nil {
		d.log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return item, nil
}
