package union_service

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	ConfigService "github.com/krilie/lico_alone/module/module-config/service"
	FileService "github.com/krilie/lico_alone/module/module-file/service"
	MessageService "github.com/krilie/lico_alone/module/module-message/service"
	UserService "github.com/krilie/lico_alone/module/module-user/service"
)

func init() {
	dig.Container.MustProvide(func(
		ModuleUser *UserService.UserService,
		ModuleConfig *ConfigService.ConfigService,
		ModuleFile *FileService.FileService,
		ModuleMessage *MessageService.MessageService,
		log *nlog.NLog,
	) *UnionService {
		return &UnionService{
			ModuleUser:    ModuleUser,
			ModuleConfig:  ModuleConfig,
			ModuleFile:    ModuleFile,
			ModuleMessage: ModuleMessage,
			log:           log,
		}
	})
}
