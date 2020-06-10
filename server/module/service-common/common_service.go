package service_common

import (
	"context"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	service3 "github.com/krilie/lico_alone/module/module-carousel/service"
	"github.com/krilie/lico_alone/module/module-config/model"
	"github.com/krilie/lico_alone/module/module-config/service"
)

type CommonService struct {
	configService  *service.ConfigModule
	moduleArticle  *service2.BlogArticleModule
	moduleCarousel *service3.CarouselModule
}

func NewCommonService(moduleArticle *service2.BlogArticleModule, configService *service.ConfigModule, moduleCarousel *service3.CarouselModule) *CommonService {
	return &CommonService{configService: configService, moduleArticle: moduleArticle, moduleCarousel: moduleCarousel}
}

func (a *CommonService) GetIcpInfo(ctx context.Context) (info model.IcpInfo) {
	_, err := a.configService.GetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), &info)
	if err != nil {
		return model.IcpInfo{Name: "", Link: "", Label: ""}
	}
	return info
}
