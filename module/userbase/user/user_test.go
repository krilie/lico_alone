package user

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/id_util"
	"github.com/krilie/lico_alone/common/jwt"
	"github.com/krilie/lico_alone/common/random"
	"testing"
)

func TestUserLogin(t *testing.T) {

	userBase := User{}

	//上下文对象
	var ctx context.Context
	ctx.TraceId = id_util.GetUuid()
	ctx.UserClaims = new(jwt.UserClaims)
	ctx.UserClaims.ClientId = "123"
	//测试数据
	userName := random.GetAToken()
	userPswd := random.GetAToken()
	fmt.Println("user_name:", userName)
	fmt.Println("user_password:", userPswd)
	//注册
	err := userBase.Register(&ctx, userName, userPswd)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	//登录
	jwtString, err := userBase.Login(&ctx, userName, userPswd)
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
	err = userBase.Logout(&ctx, jwtString)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
