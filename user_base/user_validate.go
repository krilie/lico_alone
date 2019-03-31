package user_base

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/lico603/lico_user/common/context_util"
	"github.com/lico603/lico_user/common/jwt"
)

//验证用户是否有效
func UserValidate(ctx *context_util.Context, jwtToken string) (jwt2.Claims, error) {
	claims, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}
