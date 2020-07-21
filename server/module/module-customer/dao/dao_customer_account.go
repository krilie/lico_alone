package dao

import (
	"context"
	"errors"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"gorm.io/gorm"
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
