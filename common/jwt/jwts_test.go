package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestNewJwtToken(t *testing.T) {
	var userClaims UserClaims
	userClaims.NickName = "ii"
	userClaims.UserRoles = "normal"
	userClaims.ClientId = "123"
	userClaims.Iss = "sys"
	userClaims.UserId = "34"
	userClaims.Jti = "45"
	userClaims.Iat = time.Now().Unix()
	//userClaims.Exp = time.Now().Add(time.Hour).Unix()
	userClaims.Exp = time.Now().Unix() + -1
	jwtToken, e := GetNewJwtToken(&userClaims)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(jwtToken)
	}
	claims, e := CheckJwtToken(jwtToken)
	if e != nil {
		t.Error(e)
		eV := e.(*jwt.ValidationError)
		if eV.Inner == ErrIatTime {
			t.Log("ok err time iat exp")
		} else if eV.Inner == ErrTimeExp {
			t.Log("ok err time exp")
		}
	} else {
		t.Log(claims)
	}
}
