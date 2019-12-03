/// auth_cache 权限的保存
package auth_cache

import (
	"errors"
	"github.com/mikespook/gorbac"
	"time"
)

type Role struct {
	*gorbac.StdRole
	Name       string
	ParentName string
	CreateTime time.Time
}

type Permission struct {
	Name       string
	CreateTime time.Time
	RefMethod  string
	RefPath    string
}

func (p *Permission) ID() string {
	return p.Name
}

func (p *Permission) Match(perm gorbac.Permission) bool {
	return p.Name == perm.ID()
}

type AuthCache struct {
	*gorbac.RBAC
}

func NewAuthCache() *AuthCache {
	return &AuthCache{RBAC: gorbac.New()}
}

func (a *AuthCache) AddRole(role string) error {
	err := a.RBAC.Add(gorbac.NewStdRole(role))
	return err
}

func (a *AuthCache) HasRole(role string) bool {
	_, _, err := a.RBAC.Get(role)
	if errors.Is(err, gorbac.ErrRoleNotExist) {
		return false
	} else {
		return true
	}
}
func (a *AuthCache) GetRole(role string) gorbac.Role {
	stdRole, _, err := a.RBAC.Get(role)
	if errors.Is(err, gorbac.ErrRoleNotExist) {
		return nil
	} else {
		return stdRole
	}
}
func (a *AuthCache) AttachPermToRole(role, perm string) {
	stdRole := a.GetRole(role)
	if stdRole == nil {
		stdRole = gorbac.NewStdRole(role)
		if err := a.RBAC.Add(stdRole.ID()); err != nil {
			return
		}
	}
}
