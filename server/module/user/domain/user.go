package domain

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/jwt"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

type user struct {
	dao        *dao.Dao
	id         string
	userMaster *model.UserMaster
}

func NewUser(ctx context.Context, dao *dao.Dao, userId string) (*user, error) {
	master, err := dao.GetUserMasterById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if master == nil {
		return nil, errs.NewNotFound().WithMsg("未找到用户")
	}
	return &user{
		dao:        dao,
		id:         userId,
		userMaster: master,
	}, nil
}

func NewUserByPhoneNum(ctx context.Context, dao *dao.Dao, phoneNum string) (*user, error) {
	master, err := dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		return nil, err
	}
	if master == nil {
		return nil, errs.NewNotFound().WithMsg("未找到用户")
	}
	return &user{
		dao:        dao,
		id:         master.Id,
		userMaster: master,
	}, nil
}

func (a *user) HasPassword(ctx context.Context) bool {
	return a.userMaster.Password != ""
}

func (a *user) SetPassword(ctx context.Context, password string) error {
	salt := pswd_util.GetSalt(6)
	if password != "" {
		password = pswd_util.GetMd5Password(password, salt)
	}
	err := a.dao.Db.Model(&model.UserMaster{}).Where("id=?", a.id).UpdateColumns(&model.UserMaster{Password: password, Salt: salt}).Error
	if err != nil {
		return errs.NewErrDbUpdate().WithError(err)
	}
	return nil
}

func (a *user) UpdatePassword(ctx context.Context, oldPswd, newPswd string) error {
	newSalt := pswd_util.GetSalt(6)
	if newPswd == "" {
		return errs.NewBadRequest().WithMsg("请给出新密码")
	}
	if !pswd_util.IsPasswordOk(oldPswd, a.userMaster.Password, a.userMaster.Salt) {
		return errs.NewBadRequest().WithMsg("旧密码错误")
	}
	err := a.dao.Db.Model(&model.UserMaster{}).Where("id=?", a.id).UpdateColumns(&model.UserMaster{Password: pswd_util.GetMd5Password(newPswd, newSalt), Salt: newSalt}).Error
	if err != nil {
		return errs.NewErrDbUpdate().WithError(err)
	}
	return nil
}

func (a *user) GetEntity(ctx context.Context) *model.UserMaster {
	return a.userMaster
}

func (a *user) UpdatePicture(ctx context.Context, newPic string) (oldPic string, err error) {
	if a.userMaster.Picture != "" {
		oldPic = a.userMaster.Picture
	}
	a.userMaster.Picture = newPic
	err = a.dao.UpdateUserMaster(ctx, a.userMaster)
	return oldPic, err
}

func (a *user) IsPasswordOk(pswd string) bool {
	if a.userMaster.Password == "" {
		return false
	}
	return pswd_util.IsPasswordOk(pswd, a.userMaster.Password, a.userMaster.Salt)
}

func (a *user) NewJwt(clientId string) (string, error) {
	claims := jwt.UserClaims{
		ClientId: clientId,
		UserId:   a.userMaster.Id,
		Iat:      time.Now().Unix(),
		Exp:      time.Now().Add(time.Hour * 24 * 7).Unix(),
		Jti:      id_util.GetUuid(),
		Iss:      "sys",
	}
	jwtToken, err := jwt.GetNewJwtToken(&claims)
	if err != nil {
		return "", errs.NewInternal().WithError(err).WithMsg("凭证生成失败")
	}
	return jwtToken, nil
}

func (a *user) Roles(ctx context.Context) ([]*model.UserRole, error) {
	return a.dao.GetUserRolesByUserId(ctx, a.id)
}

func (a *user) AssignRole(roleName string) error {
	err := a.dao.CreateUserRole(a.id, roleName)
	return err
}

func (a *user) HasRole(roleName string) (bool, error) {
	b, e := a.dao.HasUserRoleByName(a.id, roleName)
	return b, e
}

func (a *user) RemoveRole(roleId string) error {
	b, err := a.HasRole(roleId)
	if err != nil {
		return err
	}
	if !b {
		return errs.NewBadRequest().WithMsg("你没有此角色")
	}
	err = a.dao.Db.Model(&model.UserRole{}).Delete(&model.UserRole{}).Error
	if err != nil {
		return errs.NewErrDbDelete().WithError(err)
	}
	return nil
}
