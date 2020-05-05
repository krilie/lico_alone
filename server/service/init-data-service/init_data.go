package init_data_service

import (
	"github.com/krilie/lico_alone/component/nlog"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

type InitDataService struct {
	log          *nlog.NLog
	unionService *union_service.UnionService
}

func NewInitDataService(log *nlog.NLog, union_service *union_service.UnionService) *InitDataService {
	return &InitDataService{
		log:          log,
		unionService: union_service,
	}
}
