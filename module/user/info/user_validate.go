package info

import (
	"context"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/krilie/lico_alone/common/jwt"
)

//验证用户是否有效
func (User) Validate(ctx context.Context, jwtToken string) (jwt2.Claims, error) {
	claims, err := jwt.CheckJwtToken(jwtToken)
	if err != nil {
		return nil, err
	} else {
		return claims, nil
	}
}
