package module

import (
	article_module "github.com/krilie/lico_alone/module/module-blog-article/service"
	carousel_module "github.com/krilie/lico_alone/module/module-carousel/service"
	config_module "github.com/krilie/lico_alone/module/module-config/service"
	customer_module "github.com/krilie/lico_alone/module/module-customer/service"
	dynamic_module "github.com/krilie/lico_alone/module/module-dynamic-share/service"
	file_module "github.com/krilie/lico_alone/module/module-file/service"
	message_module "github.com/krilie/lico_alone/module/module-message/service"
	statistic_module "github.com/krilie/lico_alone/module/module-statistic/service"
	user_module "github.com/krilie/lico_alone/module/module-user/service"
)

func DigProviderModule() {
	article_module.DigModuleBlogArticleProviderAll()
	user_module.DigProviderAll()
	carousel_module.DigModuleCarouselProviderAll()
	config_module.DigModuleConfigProviderAll()
	file_module.DigModuleFileProviderAll()
	message_module.DigModuleMessageProviderAll()
	statistic_module.DigProviderAll()
	customer_module.DigModuleCustomerProviderAll()
	dynamic_module.DigModuleDynamicShareProviderAll()
}
