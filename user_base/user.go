package user_base

import (
	"github.com/asaskevich/govalidator"
	"github.com/lico603/lico-my-site-user/common/config"
)

// 用户模块的逻辑部分

// jwt过期时间，从配置文件中取
var jwtExpDuration int64

func init() {
	jwtExpDuration = int64(config.GetInt("jwt.normal_exp_duration"))
	govalidator.SetFieldsRequiredByDefault(true)
}
