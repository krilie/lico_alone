package user_api

import "context"

func (a *AppUser) UserLogin(ctx context.Context, phone, pswd, clientId string) (jwt string, err error) {
	return a.UserService.UserLogin(ctx, phone, pswd, clientId)
}

func (a *AppUser) UserRegister(ctx context.Context, phone, password, clientId string) error {
	return a.UserService.RegisterNewUser(ctx, phone, password)
}
