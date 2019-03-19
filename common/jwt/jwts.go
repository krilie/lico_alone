package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lico603/lico-my-site-user/common/config"
)

func init() {
	if config.GetString("jwt.hs256_key") == "" {
		SetHs256Key("fasdfasdrewq^&(*()&*(^%*&FLSJDF")
	} else {
		SetHs256Key(config.GetString("jwt.hs256_key"))
	}
}

var hs256Key []byte

//先设置hs256key 与config 解耦
func SetHs256Key(key string) {
	hs256Key = []byte(key)
}

//創建一個新的jwt,
func GetNewJwtToken(userClaims *UserClaims) (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = userClaims
	return jwtToken.SignedString(hs256Key) //TODO: 這個換成rsa加密
}

//檢查jwt是否有效
func CheckJwtToken(jwtString string) (userClaims UserClaims, err error) {
	_, err = jwt.ParseWithClaims(jwtString, &userClaims, func(_ *jwt.Token) (i interface{}, e error) {
		return hs256Key, nil
	})
	return
}
