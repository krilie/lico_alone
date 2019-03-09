package user_base

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/jwt"
)

//验证用户是否有效
func UserValidate(ctx *context_util.Context, jwtToken string) error {
	_, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil
	} else {
		return err
	}
}
