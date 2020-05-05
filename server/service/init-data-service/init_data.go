package init_data_service

import (
	"context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

type InitDataService struct {
	log          *nlog.NLog
	unionService *union_service.UnionService
}

func (initData *InitDataService) GetNDb(ctx context.Context) *ndb.NDb {
	return initData.unionService.ModuleConfig.Dao.NDb
}

func NewInitDataService(log *nlog.NLog, union_service *union_service.UnionService) *InitDataService {
	return &InitDataService{
		log:          log,
		unionService: union_service,
	}
}
