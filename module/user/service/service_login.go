package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/module/user/domain"
)

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
