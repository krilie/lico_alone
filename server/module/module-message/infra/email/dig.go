package email

import (
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
)

func init() {
	dig.Container.MustProvide(func(cfg *config.Config) IEmail {
		eCfg := cfg.Email
		return NewEmail(eCfg.Address, eCfg.Host, eCfg.Port, eCfg.UserName, eCfg.Password)
	})
}
