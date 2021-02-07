package dao

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"gorm.io/gorm"
	"time"
)

func (d *CustomerDao) CreateCustomerAccount(ctx context.Context, account *model.CustomerAccount) error {
	err := d.GetDb(ctx).Model(new(model.CustomerAccount)).Create(account).Error
	return err
}

func (d *CustomerDao) GetCustomerByCustomerTraceId(ctx context.Context, customerTraceId string) (*model.CustomerAccount, error) {
	customer := new(model.CustomerAccount)
	err := d.GetDb(ctx).First(&customer, "customer_trace_id=?", customerTraceId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (d *CustomerDao) GetCustomerByCustomerId(ctx context.Context, id string) (*model.CustomerAccount, error) {
	customer := new(model.CustomerAccount)
	err := d.GetDb(ctx).Where("id=?", id).Find(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (d *CustomerDao) DeleteCustomerByCustomerTraceId(ctx context.Context, customerTraceId string) error {
	err := d.GetDb(ctx).Where("customer_trace_id=?", customerTraceId).Delete(new(model.CustomerAccount)).Error
	return err
}

func (d *CustomerDao) DeleteCustomerById(ctx context.Context, id string) error {
	err := d.GetDb(ctx).Where("id=?", id).Delete(new(model.CustomerAccount)).Error
	return err
}

func (d *CustomerDao) IncreaseAccessTimes(ctx context.Context, id, ip, addr string) error {
	var sql = "update tb_customer_account set updated_at=?,last_access_ip=?,last_access_addr=?,access_times=access_times+1 where deleted_at is null and id=?"
	result := d.
		GetDb(ctx).
		Exec(sql, time.Now(), ip, addr, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
