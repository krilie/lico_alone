package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/lico603/lico-my-site-user/common/context_util"
)

//用户注册，注册，normal用户注册
func UserRegister(ctx *context_util.Context, loginName string, password string) {
	//数据验证
	govalidator.StringMatches("", "")

}
