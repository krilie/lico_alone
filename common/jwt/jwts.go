package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

//創建一個新的jwt,
func NewJwtToken(userClaims *UserClaims) (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = userClaims
	return jwtToken.SignedString([]byte("12ef")) //TODO: 這個換成rsa加密
}

//檢查jwt是否有效
func CheckJwtToken(jwtString string) (userClaims UserClaims, err error) {
	_, err = jwt.ParseWithClaims(jwtString, &userClaims, func(_ *jwt.Token) (i interface{}, e error) {
		return []byte("12ef"), nil
	})
	return
}
