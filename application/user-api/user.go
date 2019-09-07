package user_api

import (
	all_service "github.com/krilie/lico_alone/application/all-service"
	"github.com/krilie/lico_alone/module/user/service"
	"github.com/mikespook/gorbac"
)

type AppUser struct {
	UserService *service.Service
}

func (a *AppUser) HasUser(id string) (bool, error) {
	return a.UserService.AuthRBAC.HasRole(id)
}

func (a *AppUser) HasPermission(id, permission string) (bool, error) {
	rslt := a.UserService.AuthRBAC.IsGranted(id, gorbac.NewStdPermission(permission), nil)
	return rslt, nil
}

func (a *AppUser) HasRole(userId, roleId string) (bool, error) {
	strings, err := a.UserService.AuthRBAC.GetParents(userId)
	if err != nil {
		return false, err
	}
	for _, v := range strings {
		if roleId == v {
			return true, nil
		}
	}
	return false, nil
}

func NewAppUser(allSrv *all_service.AllService) *AppUser {
	return &AppUser{UserService: allSrv.UserService}
}
