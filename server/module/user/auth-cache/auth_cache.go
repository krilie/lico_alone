/// auth_cache 权限的保存
package auth_cache

import (
	"github.com/mikespook/gorbac"
	"time"
)

type Role struct {
	*gorbac.StdRole
	Name       string
	ParentName string
	CreateTime time.Time
}

func NewRole(name, pName string, createTime time.Time) *Role {
	return &Role{StdRole: gorbac.NewStdRole(name), Name: name, ParentName: pName, CreateTime: createTime}
}

type Permission struct {
	Name       string
	CreateTime time.Time
	RefMethod  string
	RefPath    string
}

func NewPermission(name, refMethod, refPath string, createTime time.Time) *Permission {
	return &Permission{Name: name, CreateTime: createTime, RefMethod: refMethod, RefPath: refPath}
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
