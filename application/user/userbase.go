package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

func (AppUser) Login(ctx *context_util.Context, loginName, password string) (jwtString string, err error) {
	panic("implement me")
}

func (AppUser) Validate(ctx *context_util.Context, jwtToken string) (jwt.Claims, error) {
	panic("implement me")
}

func (AppUser) Register(ctx *context_util.Context, loginName string, password string) error {
	panic("implement me")
}

func (AppUser) Logout(ctx *context_util.Context, jwtToken string) error {
	panic("implement me")
}

func (AppUser) ValidateClientAccToken(ctx *context_util.Context, clientAccKey string) (key *model.ClientUserAccessToken, err error) {
	panic("implement me")
}

func (AppUser) GetInfo(ctx *context_util.Context, userId string) (map[string]string, error) {
	panic("implement me")
}
