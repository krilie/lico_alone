// +build auto_test

package service

import (
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-customer/dao"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	dao.DigProvider()
	DigProvider()
	m.Run()
}

func TestCustomerModule_CreateCustomerAccount(t *testing.T) {
	dig.Container.MustInvoke(func(svc *CustomerModule) {
		account, err := svc.CreateCustomerAccount(context2.NewContext(), &model.CreateCustomerAccountModel{
			CustomerTraceId: id_util.GetUuid(),
			LoginName:       id_util.NextSnowflake(),
			Password:        "323",
			LastAccessIp:    "234",
			Mobile:          "34",
			Email:           "45",
			Other:           "56",
		})
		assert.Nil(t, err, "should nil")
		assert.NotEqual(t, "", account, "应该有值")
	})
}

func TestCustomerModule_GetOrCreateCustomerAccountByTraceId(t *testing.T) {
	dig.Container.MustInvoke(func(svc *CustomerModule) {
		var traceId = id_util.GetUuid()
		account, err := svc.GetOrCreateCustomerAccountByTraceId(context2.NewContext(), traceId, "123")
		assert.Nil(t, err, "should nil")
		assert.NotNil(t, account, "not nil")
		assert.Equal(t, traceId, account.CustomerTraceId, "应该有值")
		assert.Equal(t, "123", account.LastAccessIp, "应该有值")
	})
}
