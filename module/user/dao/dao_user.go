package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

type IUser interface {
	GetUserMasterById(ctx context.Context, userId string) (*model.UserMaster, error)
	GetUserMasterByPhoneNum(ctx context.Context, phoneNum string) (*model.UserMaster, error)
	GetUserMasterByLoginName(ctx context.Context, loginName string) (*model.UserMaster, error)
	CreateUserMaster(ctx context.Context, master *model.UserMaster) error
	UpdateUserMaster(ctx context.Context, user *model.UserMaster) error
	PhoneNumExists(ctx context.Context, phoneNum string) (bool, error)
	GetAllValidUserId(ctx context.Context) ([]string, error)
}

func (d *Dao) GetUserMasterById(ctx context.Context, userId string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.Db.Model(new(model.UserMaster)).Where("id=?", userId).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.ErrDbQuery.New().WithError(err)
	}
	return &user, nil
}

func (d *Dao) GetUserMasterByPhoneNum(ctx context.Context, phoneNum string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.Db.Model(new(model.UserMaster)).Where(&model.UserMaster{PhoneNum: &phoneNum}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.ErrDbQuery.New().WithError(err)
	}
	return &user, nil
}

func (d *Dao) GetUserMasterByLoginName(ctx context.Context, loginName string) (*model.UserMaster, error) {
	var user model.UserMaster
	err := d.Db.Model(new(model.UserMaster)).Where(&model.UserMaster{LoginName: loginName}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, errs.ErrDbQuery.New().WithError(err)
	}
	return &user, nil
}

func (d *Dao) CreateUserMaster(ctx context.Context, master *model.UserMaster) error {
	err := d.Db.Model(&model.UserMaster{}).Create(master).Error
	if err != nil {
		return errs.NewErrDbCreate().WithError(err)
	}
	return nil
}

func (d *Dao) UpdateUserMaster(ctx context.Context, user *model.UserMaster) error {
	if user.Id == "" {
		return errs.NewBadRequest().WithMsg("no primary key on update user master.")
	}
	user.UpdateTime = time.Now()
	err := d.Db.Model(&model.UserMaster{}).Save(user).Error
	if err != nil {
		return errs.NewErrDbUpdate().WithError(err)
	}
	return nil
}

func (d *Dao) PhoneNumExists(ctx context.Context, phoneNum string) (bool, error) {
	count := 0
	err := d.Db.Model(&model.UserMaster{}).Where(&model.UserMaster{PhoneNum: &phoneNum}).Count(&count).Error
	if err != nil {
		return false, errs.NewErrDbQuery().WithError(err)
	}
	return count == 1, nil
}

// GetAllValidUserId 取到有效的用户id
func (d *Dao) GetAllValidUserId(ctx context.Context) ([]string, error) {
	var list []*model.UserMaster
	err := d.Db.Model(&model.UserMaster{}).Select("id").Find(&list).Error
	if err != nil {
		return nil, errs.NewErrDbQuery().WithError(err)
	}
	var retList []string
	for _, v := range list {
		retList = append(retList, v.Id)
	}
	return retList, nil
}
