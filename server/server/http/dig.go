package http

import (
	"github.com/krilie/lico_alone/common/dig"
	ctl_common "github.com/krilie/lico_alone/server/http/ctl-common"
	ctl_health_check "github.com/krilie/lico_alone/server/http/ctl-health-check"
	ctl_user "github.com/krilie/lico_alone/server/http/ctl-user"
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

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewController)
}
