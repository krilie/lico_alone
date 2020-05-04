package init_data_service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

func init() {
	dig.Container.MustProvide(func(log *nlog.NLog, service *union_service.UnionService) *InitDataService {
		return NewInitDataService(log, service)
	})
}
