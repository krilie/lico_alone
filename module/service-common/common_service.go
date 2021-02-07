package service_common

import (
	"context"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/broker/messages"
	service2 "github.com/krilie/lico_alone/module/module-blog-article/service"
	service3 "github.com/krilie/lico_alone/module/module-carousel/service"
	service6 "github.com/krilie/lico_alone/module/module-catchword/service"
	"github.com/krilie/lico_alone/module/module-config/model"
	"github.com/krilie/lico_alone/module/module-config/service"
	service4 "github.com/krilie/lico_alone/module/module-customer/service"
	service5 "github.com/krilie/lico_alone/module/module-dynamic-share/service"
	"time"
)

type CommonService struct {
	configService      *service.ConfigModule
	ModuleArticle      *service2.BlogArticleModule
	ModuleCarousel     *service3.CarouselModule
	moduleCustomer     *service4.CustomerModule
	moduleDynamicShare *service5.DynamicShareModule
	ModuleCatchword    *service6.CatchwordModule
	broker             *broker.Broker
}

func NewCommonService(
	broker *broker.Broker,
	moduleCustomer *service4.CustomerModule,
	moduleArticle *service2.BlogArticleModule,
	configService *service.ConfigModule,
	moduleCarousel *service3.CarouselModule,
	moduleDynamicShare *service5.DynamicShareModule,
	moduleCatchword *service6.CatchwordModule) *CommonService {
	return &CommonService{
		configService:      configService,
		ModuleArticle:      moduleArticle,
		ModuleCarousel:     moduleCarousel,
		moduleCustomer:     moduleCustomer,
		moduleDynamicShare: moduleDynamicShare,
		ModuleCatchword:    moduleCatchword,
		broker:             broker,
	}
}

func (a *CommonService) GetIcpInfo(ctx context.Context) (info model.IcpInfo) {
	_, err := a.configService.GetJsonValue(ctx, model.ConfigItemsIcpInfo.Val(), &info)
	if err != nil {
		return model.IcpInfo{Name: "", Link: "", Label: ""}
	}
	return info
}

func (a *CommonService) WebVisited(ctx context.Context, ip, traceId string) {
	a.broker.MustSend(ctx, &messages.WebStationVisitedMessage{
		Ctx:        ctx,
		AccessTime: time.Now(),
		Ip:         ip,
		TraceId:    traceId,
	})
}

func (a *CommonService) GetAboutApp(ctx context.Context) (string, error) {
	str, err := a.configService.GetValueStr(ctx, model.ConfigItemsAboutApp.Val())
	if err != nil {
		return "", err
	}
	if str == nil {
		str = new(string)
	}
	return *str, nil
}
