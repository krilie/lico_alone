package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/module-user/model"
	"time"
)

func (s *UserService) ChangeUserPassword(ctx context.Context, userId, oldPswd, newPswd string) error {
	user, err := s.Dao.GetUserMasterById(ctx, userId)
	if err != nil {
		s.log.Errorf("change user password err:%v", err)
		return err
	}
	if user == nil {
		s.log.Warnf("change user password no user find id:%v", userId)
	}
	if !pswd_util.IsPasswordOk(oldPswd, user.Password, user.Salt) {
		return errs.NewBadRequest().WithMsg("password err")
	}
	user.Password = pswd_util.GetMd5Password(newPswd, user.Salt)
	err = s.Dao.UpdateUserMaster(ctx, user)
	return err
}

func (s *UserService) RegisterNewUser(ctx context.Context, phoneNum, password string) error {
	if phoneNum == "" {
		return errs.NewBadRequest().WithMsg("手机号不能为空")
	}
	if password == "" {
		password = id_util.GetUuid()
	}
	master, err := s.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		s.log.Errorf("register new user err:%v", err)
		return err
	}
	if master != nil {
		return errs.NewBadRequest().WithMsg("此手机号已注册")
	}
	salt := pswd_util.GetSalt(6)
	user := &model.UserMaster{
		Model: com_model.Model{
			Id:         id_util.NextSnowflake(),
			CreateTime: time.Now(),
		},
		UpdateTime: time.Now(),
		LoginName:  phoneNum,
		PhoneNum:   phoneNum,
		Email:      "",
		Password:   pswd_util.GetMd5Password(password, salt),
		Picture:    "",
		Salt:       salt,
	}
	err = s.Dao.CreateUserMaster(ctx, user)
	return err
}

func (s *UserService) UserLogin(ctx context.Context, phoneNum, password, clientId string) (jwt string, err error) {
	userMaster, err := s.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if userMaster == nil {
		return "", errs.NewUnauthorized().WithMsg("无此用户")
	}
	if !pswd_util.IsPasswordOk(password, userMaster.Password, userMaster.Salt) {
		return "", errs.NewBadRequest().WithMsg("密码错误")
	}
	return user.NewJwt(clientId)
}
