package ctl_common

import (
	"github.com/krilie/lico_alone/common/dig"
	common_service "github.com/krilie/lico_alone/service/common-service"
)

func init() {
	dig.Container.MustProvide(func(commonService *common_service.CommonService) *CommonCtrl {
		return &CommonCtrl{
			CommonService: commonService,
		}
	})
}
