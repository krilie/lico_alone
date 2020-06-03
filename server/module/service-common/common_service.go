package service_common

import (
	"context"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	"github.com/krilie/lico_alone/module/module-config/model"
	"github.com/krilie/lico_alone/module/module-config/service"
)

type CommonService struct {
	configService *service.ConfigModule
	moduleArticle *service2.BlogArticleModule
}

func NewCommonService(moduleArticle *service2.BlogArticleModule, configService *service.ConfigModule) *CommonService {
	return &CommonService{configService: configService, moduleArticle: moduleArticle}
}

func (c *CommonService) GetIcpInfo(ctx context.Context) (info model.IcpInfo) {
	_, err := c.configService.GetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), &info)
	if err != nil {
		return model.IcpInfo{Name: "", Link: "", Label: ""}
	}
	return info
}
