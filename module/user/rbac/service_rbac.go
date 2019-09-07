package rbac

import (
	"github.com/krilie/lico_alone/common/errs"
	"github.com/mikespook/gorbac"
)

type RBAC struct {
	*gorbac.RBAC
}

func NewRBAC() *RBAC {
	return &RBAC{RBAC: gorbac.New()}
}

func (a *RBAC) HasRole(id string) (bool, error) {
	_, _, err := a.Get(id)
	if err != nil && err == gorbac.ErrRoleNotExist {
		return false, nil
	}
	if err != nil {
		return false, errs.NewInternal().WithError(err)
	}
	return true, nil
}

func (a *RBAC) GetRole(role string) (gorbac.Role, error) {
	r, _, err := a.RBAC.Get(role)
	if err != nil && err == gorbac.ErrRoleNotExist {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *RBAC) AddRole(id string) error {
	err := a.RBAC.Add(gorbac.NewStdRole(id))
	if err != nil && err == gorbac.ErrRoleExist {
		return nil
	}
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	return nil
}
