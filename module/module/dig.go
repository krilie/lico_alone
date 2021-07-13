package module

import (
	article_module "github.com/krilie/lico_alone/module/module-blog-article/service"
	carousel_module "github.com/krilie/lico_alone/module/module-carousel/service"
	"github.com/krilie/lico_alone/module/module-catchword/service"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	customer_module "github.com/krilie/lico_alone/module/module-customer/service"
	dynamic_module "github.com/krilie/lico_alone/module/module-dynamic-share/service"
	file_module "github.com/krilie/lico_alone/module/module-file/service"
	module_like_dislike "github.com/krilie/lico_alone/module/module-like-dislike"
	message_module "github.com/krilie/lico_alone/module/module-message/service"
	statistic_module "github.com/krilie/lico_alone/module/module-statistic/service"
	user_module "github.com/krilie/lico_alone/module/module-user/service"
	service2 "github.com/krilie/lico_alone/module/module-zunion/service"
)

var DigProviderModuleAll = BuildProviderList()

func BuildProviderList() []interface{} {
	var list = []interface{}{
		article_module.DigModuleBlogArticleProviderAll,
		user_module.DigModuleUserProviderAll,
		carousel_module.DigModuleCarouselProviderAll,
		file_module.DigModuleFileProviderAll,
		message_module.DigModuleMessageProviderAll,
		statistic_module.DigModuleStatisticProviderAll,
		customer_module.DigModuleCustomerProviderAll,
		dynamic_module.DigModuleDynamicShareProviderAll,
		ConfigService.DigModuleConfigProviderAll,
		module_like_dislike.DigModuleLikeDisLikeProviderAll,
		service.DigModuleCatchwordProviderAll,
		service2.DigModuleZUnionProviderAll,
	}
	var targetList []interface{}
	for i := range list {
		targetList = append(targetList, list[i].([]interface{})...)
	}
	return targetList
}
