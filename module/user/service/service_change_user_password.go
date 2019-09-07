package service

import (
	"context"
	"github.com/krilie/lico_alone/module/user/domain"
)

func (s *Service) ChangeUserPassword(ctx context.Context, userId, oldPswd, newPswd string) error {
	user, err := domain.NewUser(ctx, s.Dao, userId)
	if err != nil {
		return err
	}
	err = user.UpdatePassword(ctx, oldPswd, newPswd)
	if err != nil {
		return err
	}
	return nil
}
