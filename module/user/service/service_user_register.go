package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

func (s *Service) RegisterNewUser(ctx context.Context, phoneNum, password string) error {
	log := clog.NewLog(ctx, "module/user/service/service_user_register.go:14", "RegisterNewUser")
	if phoneNum == "" || password == "" {
		return errs.NewBadRequest().WithMsg("参数错误")
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
