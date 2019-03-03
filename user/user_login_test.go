package user

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/jwt"
	"github.com/lico603/lico-my-site-user/common/string_util"
	"github.com/lico603/lico-my-site-user/common/uuid_util"
	"testing"
)

func TestUserLogin(t *testing.T) {
	var ctx context_util.Context

	ctx.StackId = uuid_util.GetUuid()
	ctx.Auth = new(context_util.AuthInfo)
	ctx.Auth.AppId = string_util.NewString("123")
	ctx.Auth.AppId = string_util.NewString("123")

	jwtString, err := UserLogin(&ctx, "ii", "12345678")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(jwtString)
	}
	userClaims, err := jwt.CheckJwtToken(jwtString)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(userClaims)
	}
}
