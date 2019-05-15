package user_base

import (
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
)

// 用户模块的逻辑部分

// jwt过期时间，从配置文件中取
var jwtExpDuration int64

func init() {
	jwtExpDuration = int64(config.GetInt("jwt.normal_exp_duration"))
	govalidator.SetFieldsRequiredByDefault(true)
}

type IUser interface {
	UserLogin(ctx *context_util.Context, loginName, password string) (jwtString string, err error)
	UserValidate(ctx *context_util.Context, jwtToken string) (jwt.Claims, error)
	UserBaseRegister(ctx *context_util.Context, loginName string, password string) error
	UserLogout(ctx *context_util.Context, jwtToken string) error
	UserValidateClientAccToken(ctx *context_util.Context, clientAccKey string) (key *model.ClientUserAccessToken, err error)
	UserBaseGetInfo(ctx *context_util.Context, userId string) (map[string]string, error)
}
type User struct{}
