package email

import (
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewIEmail(cfg *ncfg.NConfig) IEmail {
	eCfg := cfg.Cfg.Email
	return NewEmail(eCfg.Address, eCfg.Host, eCfg.Port, eCfg.UserName, eCfg.Password)
}

func init() {
	dig.Container.MustProvide(NewIEmail)
}

// DigProvider provider
func DigProvider() {
	dig.Container.MustProvide(NewIEmail)
}
