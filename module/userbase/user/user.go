package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/userbase/model"
)

// 用户模块的逻辑部分

// jwt过期时间，从配置文件中取
var jwtExpDuration int64

func init() {
	jwtExpDuration = int64(config.GetInt("jwt.normal_exp_duration"))
	govalidator.SetFieldsRequiredByDefault(true)
}

type IUser interface {
	Login(ctx *context.Context, loginName, password string) (jwtString string, err error)
	Validate(ctx *context.Context, jwtToken string) (jwt.Claims, error)
	Register(ctx *context.Context, loginName string, password string) error
	Logout(ctx *context.Context, jwtToken string) error
	ValidateClientAccToken(ctx *context.Context, clientAccKey string) (key *model.ClientUserAccessToken, err error)
	GetInfo(ctx *context.Context, userId string) (map[string]string, error)
}
type User struct{}
