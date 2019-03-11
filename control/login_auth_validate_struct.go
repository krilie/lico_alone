package control

import (
	"github.com/deckarep/golang-set"
)

//登录和权限检查
type auth struct {
	role       mapset.Set
	permission mapset.Set
}

//权限保存体
type validate map[string]auth

var authMap validate

//添加一个role
func (v validate) addRole(url, role string) bool {
	if _, ok := v[url]; !ok {
		//不存在
		v[url] = auth{}
	}
	return v[url].addRole(role)
}

//添加一个permission
func (v validate) addPermission(url, permission string) bool {
	if _, ok := v[url]; !ok {
		//不存在
		v[url] = auth{}
	}
	return v[url].addPermission(permission)
}

func (a *auth) addRole(role string) bool {
	if a.role == nil {
		a.role = mapset.NewThreadUnsafeSet()
	}
	return a.role.Add(role)
}
func (a *auth) addPermission(permission string) bool {
	if a.permission == nil {
		a.permission = mapset.NewThreadUnsafeSet()
	}
	return a.permission.Add(permission)
}

//添加一个permission
func (v validate) hasRole(url, role string) bool {
	if _, ok := v[url]; !ok {
		//不存在
		return false
	}
	if v[url].role == nil {
		return false
	}
	return v[url].role.Contains(role)
}

//添加一个permission
func (v validate) hasPermission(url, permission string) bool {
	if _, ok := v[url]; !ok {
		//不存在
		return false
	}
	if v[url].permission == nil {
		return false
	}
	return v[url].permission.Contains(permission)
}
