package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/module-user/model"
	"gorm.io/gorm"
	"time"
)

type IUser interface {
	GetUserMasterById(ctx context.Context, userId string) (*model.UserMaster, error)
	GetUserMasterByPhoneNum(ctx context.Context, phoneNum string) (*model.UserMaster, error)
	GetUserMasterByLoginName(ctx context.Context, loginName string) (*model.UserMaster, error)
	CreateUserMaster(ctx context.Context, master *model.UserMaster) error
	UpdateUserMaster(ctx context.Context, user *model.UserMaster) error
	UpdateUserPassword(ctx context.Context, userId, md5edPswd, salt string) error
	IsPhoneNumExists(ctx context.Context, phoneNum string) (bool, error)
	GetAllValidUserId(ctx context.Context) ([]string, error)
	DeleteUserByPhone(ctx context.Context, phone string) error
}

func (d *UserDao) GetUserMasterById(ctx context.Context, userId string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.GetDb(ctx).Model(new(model.UserMaster)).Where("id=?", userId).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return &user, nil
}

func (d *UserDao) GetUserMasterByPhoneNum(ctx context.Context, phoneNum string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.GetDb(ctx).Where("phone_num=?", phoneNum).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return &user, nil
}

func (d *UserDao) GetUserMasterByLoginName(ctx context.Context, loginName string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.GetDb(ctx).Model(new(model.UserMaster)).Where(&model.UserMaster{LoginName: loginName}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.NewInternal().WithError(err)
	}
	return &user, nil
}

func (d *UserDao) CreateUserMaster(ctx context.Context, master *model.UserMaster) error {
	err := d.GetDb(ctx).Model(&model.UserMaster{}).Create(master).Error
	if err != nil {
		d.log.Get(ctx).Errorf("create user master db err:%v", err)
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) UpdateUserMaster(ctx context.Context, user *model.UserMaster) error {
	if user.Id == "" {
		return errs.NewNormal().WithMsg("no primary key on update user master.")
	}
	user.UpdatedAt = time.Now()
	err := d.GetDb(ctx).Model(&model.UserMaster{}).Save(user).Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) IsPhoneNumExists(ctx context.Context, phoneNum string) (bool, error) {
	count := 0
	err := d.GetDb(ctx).Model(&model.UserMaster{}).Where(&model.UserMaster{PhoneNum: phoneNum}).Count(&count).Error
	if err != nil {
		return false, errs.NewInternal().WithError(err)
	}
	return count == 1, nil
}

// GetAllValidUserId 取到有效的用户id
func (d *UserDao) GetAllValidUserId(ctx context.Context) ([]string, error) {
	var list []*model.UserMaster
	err := d.GetDb(ctx).Model(&model.UserMaster{}).Select("id").Find(&list).Error
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	var retList []string
	for _, v := range list {
		retList = append(retList, v.Id)
	}
	return retList, nil
}

func (d *UserDao) UpdateUserPassword(ctx context.Context, userId, md5edPswd, salt string) error {
	if userId == "" {
		return errs.NewNormal().WithMsg("no primary key on update user master.")
	}
	err := d.GetDb(ctx).
		Model(&model.UserMaster{}).
		Where(&model.UserMaster{Model: com_model.Model{Id: userId}}).
		UpdateColumns(map[string]interface{}{"update_time": time.Now(), "password": md5edPswd, "salt": salt}).
		Error
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}

func (d *UserDao) DeleteUserByPhone(ctx context.Context, phone string) error {
	err := d.GetDb(ctx).Model(new(model.UserMaster)).Delete(&model.UserMaster{
		PhoneNum: phone,
	}).Error
	return err
}
