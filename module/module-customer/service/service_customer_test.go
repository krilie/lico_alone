package service

import (
	"github.com/krilie/lico_alone/common/appdig"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-customer/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvides(DigModuleCustomerProviderAll)

func TestAutoCustomerModule_CreateCustomerAccount(t *testing.T) {
	container.MustInvoke(func(svc *CustomerModule) {
		account, err := svc.CreateCustomerAccount(context2.EmptyAppCtx(), &model.CreateCustomerAccountModel{
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

func TestAutoCustomerModule_GetOrCreateCustomerAccountByTraceId(t *testing.T) {
	container.MustInvoke(func(svc *CustomerModule) {
		var traceId = id_util.GetUuid()
		account, err := svc.GetOrCreateCustomerAccountByTraceId(context2.EmptyAppCtx(), traceId, "123")
		assert.Nil(t, err, "should nil")
		assert.NotNil(t, account, "not nil")
		assert.Equal(t, traceId, account.CustomerTraceId, "应该有值")
		assert.Equal(t, "123", account.LastAccessIp, "应该有值")
	})
}

func TestAutoCustomerModule_IncreaseCustomerAccessTimesByTraceId(t *testing.T) {
	container.MustInvoke(func(svc *CustomerModule) {
		err := svc.IncreaseCustomerAccessTimesByTraceId(context2.EmptyAppCtx(), id_util.NextSnowflake(), "4432")
		assert.Nil(t, err, "should nil")
	})
}
