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

type messageValidCode struct {
	*ndb.NDb
	log *nlog.NLog
}

type IMessageValidCode interface {
	CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	DeleteMessageValidCode(ctx context.Context, id string) error
	GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error)
	GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string, validType model.ValidCodeType) (*model.MessageValidCode, error)
}

func (d *messageValidCode) CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	err := d.GetDb(ctx).Create(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageValidCode) UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	err := d.GetDb(ctx).Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageValidCode) DeleteMessageValidCode(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(&model.MessageValidCode{}).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *messageValidCode) GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error) {
	item := &model.MessageValidCode{}
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

func (d *messageValidCode) GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string, validType model.ValidCodeType) (*model.MessageValidCode, error) {
	item := &model.MessageValidCode{}
	err := d.GetDb(ctx).Where("phone_num=? and type=?", phoneNum, validType.ToInt()).Order("created_at desc").First(item).Error
	if err != nil {
		d.log.Get(ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return item, nil
}
