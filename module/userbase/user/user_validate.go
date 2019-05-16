package user

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/jwt"
)

//验证用户是否有效
func (User) UserValidate(ctx *context_util.Context, jwtToken string) (jwt2.Claims, error) {
	claims, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}
