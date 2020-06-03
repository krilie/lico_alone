package service_common

import (
	"context"
	"github.com/krilie/lico_alone/module/module-config/model"
	"github.com/krilie/lico_alone/module/module-config/service"
)

type CommonService struct {
	configService *service.ConfigModule
}

func NewCommonService(configService *service.ConfigModule) *CommonService {
	return &CommonService{configService: configService}
}

func (c *CommonService) GetIcpInfo(ctx context.Context) (info model.IcpInfo) {
	_, err := c.configService.GetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), &info)
	if err != nil {
		return model.IcpInfo{Name: "", Link: "", Label: ""}
	}
	return info
}
