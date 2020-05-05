package dao

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-message/model"
)

type IMessageValidCode interface {
	CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	DeleteMessageValidCode(ctx context.Context, id string) error
	GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error)
	GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string, validType int) (*model.MessageValidCode, error)
}

func (d *MessageDao) CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) DeleteMessageValidCode(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageValidCode{}).Error
	if err != nil {
		d.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *MessageDao) GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error) {
	item := &model.MessageValidCode{}
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

func (d *MessageDao) GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string, validType model.ValidCodeType) (*model.MessageValidCode, error) {
	item := &model.MessageValidCode{}
	err := d.GetDb(ctx).Where("phone_num=? and type=?", phoneNum, validType.ToInt()).Order("create_time desc").First(item).Error
	if err != nil {
		d.log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return item, nil
}
