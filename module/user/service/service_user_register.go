package service

import (
	"context"
	"github.com/krilie/lico_alone/common/cmodel"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

func (s *Service) RegisterNewUser(ctx context.Context, loginName, password string) error {
	master, err := s.Dao.GetUserMasterByLoginName(ctx, loginName)
	if err != nil {
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
		LoginName:  loginName,
		PhoneNum:   nil,
		Email:      nil,
		Password:   pswd_util.GetMd5Password(password, salt),
		Picture:    nil,
		Salt:       salt,
	}
	err = s.Dao.CreateUserMaster(ctx, &user)
	return err
}
