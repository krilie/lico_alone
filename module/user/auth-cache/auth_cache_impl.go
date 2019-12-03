package auth_cache

import (
	"errors"
	"github.com/mikespook/gorbac"
)

func (a *AuthCache) AddRole(role *Role) error {
	err := a.RBAC.Add(role)
	return err
}

func (a *AuthCache) HasRole(roleId string) bool {
	_, _, err := a.RBAC.Get(roleId)
	if errors.Is(err, gorbac.ErrRoleNotExist) {
		return false
	} else {
		return true
	}
}

func (a *AuthCache) GetRole(roleId string) *Role {
	a.IsGranted()
	stdRole, _, err := a.RBAC.Get(roleId)
	if errors.Is(err, gorbac.ErrRoleNotExist) {
		return nil
	} else {
		return stdRole.(*Role)
	}
}
