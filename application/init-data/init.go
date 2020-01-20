// Package init 一些数据的初始化
package init_data

import (
	"context"
	"github.com/jinzhu/gorm"
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/common/cdb"
	configService "github.com/krilie/lico_alone/module/config/service"
	userService "github.com/krilie/lico_alone/module/user/service"
)

type Init struct {
	UserService   *userService.Service
	ConfigService *configService.Service
}

func (a *Init) SetTx(ctx context.Context, tx *gorm.DB) (srv cdb.Service, err error) {
	newUserService, err := a.UserService.SetTx(ctx, tx)
	if err != nil {
		return nil, err
	}
	configure, err := a.ConfigService.SetTx(ctx, tx)
	if err != nil {
		return nil, err
	}
	return &Init{
		UserService:   newUserService.(*userService.Service),
		ConfigService: configure.(*configService.Service),
	}, nil
}

func (a *Init) GetDb(ctx context.Context) *gorm.DB {
	return a.UserService.GetDb(ctx) // 返回其中一个db
}

func NewInit(allSrv *all_service.AllService) *Init {
	return &Init{
		UserService:   allSrv.UserService,
		ConfigService: allSrv.ConfigService,
	}
}
