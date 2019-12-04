package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/user/domain"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

func (s *Service) ChangeUserPassword(ctx context.Context, userId, oldPswd, newPswd string) error {
	log := clog.NewLog(ctx, "module/user/service/service_change_user_password.go:9", "ChangeUserPassword")
	user, err := domain.NewUser(ctx, s.Dao, userId)
	if err != nil {
		log.Errorf("change user password err:%v", err)
		return err
	}
	err = user.UpdatePassword(ctx, oldPswd, newPswd)
	if err != nil {
		log.Errorf("change user password err:%v", err)
		return err
	}
	return nil
}

func (s *Service) RegisterNewUser(ctx context.Context, phoneNum, password string) error {
	log := clog.NewLog(ctx, "module/user/service/service_user_register.go:14", "RegisterNewUser")
	if phoneNum == "" {
		return errs.NewBadRequest().WithMsg("手机号不能为空")
	}
	if password == "" {
		password = id_util.GetUuid()
	}
	master, err := s.Dao.GetUserMasterByPhoneNum(ctx, phoneNum)
	if err != nil {
		log.Errorf("register new user err:%v", err)
		return err
	}
	if master != nil {
		return errs.NewBadRequest().WithMsg("此手机号已注册")
	}
	salt := pswd_util.GetSalt(6)
	user := model.UserMaster{
		Model: cmodel.Model{
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

func (s *Service) UserLogin(ctx context.Context, phoneNum, password, clientId string) (jwt string, err error) {
	log := clog.NewLog(ctx, "module/user/service/service_login.go:10", "UserLogin")
	user, err := domain.NewUserByPhoneNum(ctx, s.Dao, phoneNum)
	if err != nil {
		log.Errorf("user login err:%v", err)
		return "", err
	}
	ok := user.IsPasswordOk(password)
	if !ok {
		return "", errs.NewBadRequest().WithMsg("密码错误")
	}
	return user.NewJwt(clientId)
}
