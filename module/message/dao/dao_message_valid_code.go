package dao

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/message/model"
)

type IMessageValidCode interface {
	CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error
	DeleteMessageValidCode(ctx context.Context, id string) error
	GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error)
	GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string) (*model.MessageValidCode, error)
}

func (d *Dao) CreateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_valid_code.go:18", "CreateMessageValidCode")
	err := d.Dao.Db.Create(item).Error
	if err != nil {
		log.Error(err)
		return errs.ErrDbCreate.WithError(err)
	}
	return nil
}

func (d *Dao) UpdateMessageValidCode(ctx context.Context, item *model.MessageValidCode) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_valid_code.go:28", "UpdateMessageValidCode")
	err := d.Dao.Db.Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		log.Error(err)
		return errs.ErrDbCreate.WithError(err)
	}
	return nil
}

func (d *Dao) DeleteMessageValidCode(ctx context.Context, id string) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_valid_code.go:38", "DeleteMessageValidCode")
	err := d.Dao.Db.Where("id=?", id).Delete(&model.MessageValidCode{}).Error
	if err != nil {
		log.Error(err)
		return errs.ErrDbCreate.WithError(err)
	}
	return nil
}

func (d *Dao) GetMessageValidCodeById(ctx context.Context, id string) (*model.MessageValidCode, error) {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_valid_code.go:50", "GetMessageValidCodeById")
	item := &model.MessageValidCode{}
	err := d.Dao.Db.Where("id=?", id).Find(item).Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.ErrDbCreate.WithError(err)
	}
	return item, nil
}

func (d *Dao) GetLastMessageValidCodeByPhoneNum(ctx context.Context, phoneNum string) (*model.MessageValidCode, error) {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_valid_code.go:50", "GetMessageValidCodeById")
	item := &model.MessageValidCode{}
	err := d.Dao.Db.Where("phone_num=?", phoneNum).Order("create_time desc").First(item).Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.ErrDbCreate.WithError(err)
	}
	return item, nil
}
