package user

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/jwt"
)

//用户登出，只判断是否jwtToken有效并返回
func UserLogout(ctx *context_util.Context, jwtToken string) error {
	_, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil
	} else {
		return err
	}
}
