package user_api

import (
	"errors"
	all_service "github.com/krilie/lico_alone/application/all-service"
	service2 "github.com/krilie/lico_alone/module/message/service"
	"github.com/krilie/lico_alone/module/module-user/service"
	"github.com/mikespook/gorbac"
)

type AppUser struct {
	UserService *service.UserService
	Message     *service2.Service
}

func (a *AppUser) HasUser(id string) (bool, error) {
	_, _, err := a.UserService.AuthRBAC.Get(id)
	if err != nil {
		if errors.Is(err, gorbac.ErrRoleNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
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
	return &AppUser{UserService: allSrv.UserService, Message: allSrv.Message}
}
