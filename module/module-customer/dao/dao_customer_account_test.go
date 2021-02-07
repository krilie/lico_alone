package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	com_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAllForTest).
	MustProvide(NewCustomerDao)

func TestAutoCustomerDao_CreateCustomerAccount(t *testing.T) {
	container.MustInvoke(func(dao *CustomerDao) {
		test := AddCustomerDataForTest(t, context2.EmptyAppCtx(), dao)
		assert.NotNil(t, test, "should not nil")
	})
}

func TestAutoCustomerDao_DeleteCustomerByCustomerTraceId(t *testing.T) {
	container.MustInvoke(func(dao *CustomerDao) {
		test := AddCustomerDataForTest(t, context2.EmptyAppCtx(), dao)
		err := dao.DeleteCustomerByCustomerTraceId(context2.EmptyAppCtx(), test.CustomerTraceId)
		assert.Equal(t, nil, err, "should no err")
	})
}

func TestAutoCustomerDao_DeleteCustomerById(t *testing.T) {
	container.MustInvoke(func(dao *CustomerDao) {
		test := AddCustomerDataForTest(t, context2.EmptyAppCtx(), dao)
		err := dao.DeleteCustomerById(context2.EmptyAppCtx(), test.Id)
		assert.Equal(t, nil, err, "should no err")
	})
}

func TestAutoCustomerDao_GetCustomerByCustomerTraceId(t *testing.T) {
	container.MustInvoke(func(dao *CustomerDao) {
		test := AddCustomerDataForTest(t, context2.EmptyAppCtx(), dao)
		customerInfo, err := dao.GetCustomerByCustomerTraceId(context2.EmptyAppCtx(), test.CustomerTraceId)
		assert.Equal(t, nil, err, "should no err")
		assert.NotNil(t, customerInfo, "should not nil")
		assert.Equal(t, test.CustomerTraceId, customerInfo.CustomerTraceId, "should same")
	})
}

func TestAutoCustomerDao_IncreaseAccessTimes(t *testing.T) {
	container.MustInvoke(func(dao *CustomerDao) {
		test := AddCustomerDataForTest(t, context2.EmptyAppCtx(), dao)
		err := dao.IncreaseAccessTimes(context2.EmptyAppCtx(), test.Id, "123", "1234")
		assert.Nil(t, err, "should not nil")
	})
}

func AddCustomerDataForTest(t *testing.T, ctx context.Context, dao *CustomerDao) *model.CustomerAccount {
	customerModel := &model.CustomerAccount{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Time{},
				Valid: false,
			},
		},
		CustomerTraceId: id_util.GetUuid(),
		LoginName:       id_util.NextSnowflake(),
		Password:        "123",
		LastAccessIp:    "123",
		Mobile:          "123",
		Email:           "123",
		Other:           "123",
	}
	err := dao.CreateCustomerAccount(context2.EmptyAppCtx(), customerModel)
	assert.Equal(t, nil, err, "should no err")
	return customerModel
}
