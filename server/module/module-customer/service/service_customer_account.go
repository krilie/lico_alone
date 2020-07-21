package service

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"gorm.io/gorm"
	"time"
)

// CreateCustomerAccount 创建一个用户
func (svc *CustomerModule) CreateCustomerAccount(ctx context.Context, item *model.CreateCustomerAccountModel) (id string, err error) {
	id = id_util.GetUuid()
	err = svc.Dao.CreateCustomerAccount(ctx, &model.CustomerAccount{
		Model: com_model.Model{
			Id: id, CreatedAt: time.Now(), UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		CustomerTraceId: item.CustomerTraceId, LoginName: item.LoginName, Password: item.Password,
		LastAccessIp: item.LastAccessIp, LastAccessAddr: svc.ipApi.GetIpInfoRegionCityOrEmpty(ctx, item.LastAccessIp), AccessTimes: 0,
		Mobile: item.Mobile, Email: item.Email, Other: item.Other,
	})
	return id, err
}

func (svc *CustomerModule) GetOrCreateCustomerAccountByTraceId(ctx context.Context, traceId, ip string) (*model.CustomerAccount, error) {
	customerAccount, err := svc.Dao.GetCustomerByCustomerTraceId(ctx, traceId)
	if err != nil {
		return nil, err
	}
	if customerAccount != nil {
		return customerAccount, nil
	}
	customerId, err := svc.CreateCustomerAccount(ctx, &model.CreateCustomerAccountModel{
		CustomerTraceId: traceId,
		LoginName:       id_util.NextSnowflake(),
		Password:        "",
		LastAccessIp:    ip,
		Mobile:          "",
		Email:           "",
		Other:           "auto create",
	})
	if err != nil {
		return nil, err
	}
	return svc.Dao.GetCustomerByCustomerId(ctx, customerId)
}

func (svc *CustomerModule) IncreaseCustomerAccessTimesByTraceId(ctx context.Context, traceId, ip string) error {
	customer, err := svc.GetOrCreateCustomerAccountByTraceId(ctx, traceId, ip)
	if err != nil {
		return err
	}
	err = svc.Dao.IncreaseAccessTimes(ctx, customer.Id, ip, svc.ipApi.GetIpInfoRegionCityOrEmpty(ctx, ip))
	return err
}
