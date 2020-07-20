package module

import (
	article_module "github.com/krilie/lico_alone/module/module-blog-article/service"
	carousel_module "github.com/krilie/lico_alone/module/module-carousel/service"
	config_module "github.com/krilie/lico_alone/module/module-config/service"
	file_module "github.com/krilie/lico_alone/module/module-file/service"
	message_module "github.com/krilie/lico_alone/module/module-message/service"
	statistic_module "github.com/krilie/lico_alone/module/module-statistic/service"
	user_module "github.com/krilie/lico_alone/module/module-user/service"
)

func DigProviderModule() {
	article_module.DigProviderWithDao()
	user_module.DigProviderAll()
	carousel_module.DigProviderWithDao()
	config_module.DigProviderWithDao()
	file_module.DigProviderWithDao()
	message_module.DigProviderAll()
	statistic_module.DigProviderAll()
}
