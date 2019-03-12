package user_base

import (
	"fmt"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/jwt"
	"github.com/lico603/lico-my-site-user/common/random_token"
	"github.com/lico603/lico-my-site-user/common/uuid_util"
	"testing"
)

func TestUserLogin(t *testing.T) {
	//上下文对象
	var ctx context_util.Context
	ctx.TraceId = uuid_util.GetUuid()
	ctx.UserClaims = new(jwt.UserClaims)
	ctx.UserClaims.AppId = "123"
	//测试数据
	userName := random_token.GetAToken()
	userPswd := random_token.GetAToken()
	fmt.Println("user_name:", userName)
	fmt.Println("user_password:", userPswd)
	//注册
	err := UserBaseRegister(&ctx, userName, userPswd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//登录
	jwtString, err := UserLogin(&ctx, userName, userPswd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//检查token
	t.Log("jwt string" + jwtString)
	userClaims, err := jwt.CheckJwtToken(jwtString)
	if err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("jwt ok", userClaims)
	}
	//登出
	err = UserLogout(&ctx, jwtString)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
