package user_base

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/jwt"
)

//用户登出，只判断是否jwtToken有效并返回
func (User) UserLogout(ctx *context_util.Context, jwtToken string) error {
	_, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil
	} else {
		return err
	}
}
