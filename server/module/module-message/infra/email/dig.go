package email

import (
	"github.com/krilie/lico_alone/component/ncfg"
)

func NewIEmail(cfg *ncfg.NConfig) IEmail {
	var emailConfig = cfg.GetEmailCfg()
	return NewEmail(emailConfig.Address, emailConfig.Host, emailConfig.Port, emailConfig.UserName, emailConfig.Password)
}
