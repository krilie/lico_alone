package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/domain"
	"time"
)

func (s *UserService) ChangeUserPassword(ctx context.Context, userId, oldPswd, newPswd string) error {
	user, err := domain.NewUser(ctx, s.Dao, userId)
	if err != nil {
		s.log.Errorf("change user password err:%v", err)
		return err
	}
	err = user.UpdatePassword(ctx, oldPswd, newPswd)
	if err != nil {
		s.log.Errorf("change user password err:%v", err)
		return err
	}
	return nil
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
	user := com_model.UserMaster{
		Model: com_model.Model{
			Id:         id_util.NextSnowflake(),
			CreateTime: time.Now(),
		},
		UpdateTime: time.Now(),
		LoginName:  phoneNum,
		PhoneNum:   phoneNum,
		Email:      nil,
		Password:   pswd_util.GetMd5Password(password, salt),
		Picture:    "",
		Salt:       salt,
	}
	err = s.Dao.CreateUserMaster(ctx, &user)
	return err
}

func (s *UserService) UserLogin(ctx context.Context, phoneNum, password, clientId string) (jwt string, err error) {
	user, err := domain.NewUserByPhoneNum(ctx, s.Dao, phoneNum)
	if err != nil {
		s.log.Errorf("user login err:%v", err)
		return "", err
	}
	ok := user.IsPasswordOk(password)
	if !ok {
		return "", errs.NewBadRequest().WithMsg("密码错误")
	}
	return user.NewJwt(clientId)
}
