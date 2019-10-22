package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/module/user/domain"
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
