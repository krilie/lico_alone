package dao

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/component/clog"
	"github.com/krilie/lico_alone/module/message/model"
)

type IMessageEmail interface {
	CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error
	DeleteMessageEmail(ctx context.Context, id string) error
	GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error)
}

func (d *Dao) CreateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_email.go:20", "CreateMessageEmail")
	err := d.Dao.Db.Create(item).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) UpdateMessageEmail(ctx context.Context, item *model.MessageEmail) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_email.go:20", "UpdateMessageEmail")
	err := d.Dao.Db.Omit("create_time").Where("id=?", item.Id).Update(item).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) DeleteMessageEmail(ctx context.Context, id string) error {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_email.go:20", "DeleteMessageEmail")
	err := d.Dao.Db.Where("id=?", id).Delete(&model.MessageEmail{}).Error
	if err != nil {
		log.Error(err)
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) GetMessageEmailById(ctx context.Context, id string) (*model.MessageEmail, error) {
	log := clog.NewLog(ctx, "module/message/dao/dao_message_email.go:47", "GetMessageEmailById")
	item := &model.MessageEmail{}
	err := d.Dao.Db.Where("id=?", id).Find(item).Error
	if err != nil {
		log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.NewErrDbCreate().WithError(err)
	}
	return item, nil
}
