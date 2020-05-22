package common_service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
	union_service "github.com/krilie/lico_alone/service/union-service"
)

type CommonService struct {
	UnionService *union_service.UnionService
}

func NewCommonService(unionService *union_service.UnionService) *CommonService {
	return &CommonService{UnionService: unionService}
}

func (c *CommonService) GetIcpInfo(ctx context.Context) (info model.IcpInfo) {
	_, err := c.UnionService.ModuleConfig.GetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), &info)
	if err != nil {
		return model.IcpInfo{Name: "", Link: "", Label: ""}
	}
	return info
}
