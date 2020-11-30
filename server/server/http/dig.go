package http

import (
	ctl_common "github.com/krilie/lico_alone/server/http/ctl-common"
	ctl_health_check "github.com/krilie/lico_alone/server/http/ctl-health-check"
	ctl_user "github.com/krilie/lico_alone/server/http/ctl-user"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"github.com/krilie/lico_alone/server/http/middleware"
)

type Controllers struct {
	commonCtrl      *ctl_common.CommonCtrl
	userCtrl        *ctl_user.UserCtrl
	healthCheckCtrl *ctl_health_check.HealthCheckCtrl
}

func NewController(commonCtrl *ctl_common.CommonCtrl, userCtrl *ctl_user.UserCtrl, healthCheckCtrl *ctl_health_check.HealthCheckCtrl) *Controllers {
	return &Controllers{
		commonCtrl:      commonCtrl,
		userCtrl:        userCtrl,
		healthCheckCtrl: healthCheckCtrl,
	}
}

var DigControllerProviderAll = []interface{}{
	ginutil.NewGinUtils,
	middleware.NewGinMiddleware,
	ctl_common.NewCommonCtrl,
	ctl_user.NewUserCtrl,
	ctl_health_check.NewHealthCheckCtl,
	NewController,
}
