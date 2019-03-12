package middle_funcs

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
func (v validate) addRole(url string, role ...string) {
	if _, ok := v[url]; !ok {
		//不存在
		v[url] = auth{}
	}
	v[url].addRole(role...)
}

//添加一个permission
func (v validate) addPermission(url string, permission ...string) {
	if _, ok := v[url]; !ok {
		//不存在
		v[url] = auth{}
	}
	v[url].addPermission(permission...)
}

func (a *auth) addRole(role ...string) {
	if a.role == nil {
		a.role = mapset.NewThreadUnsafeSet()
	}
	for _, val := range role {
		a.role.Add(val)
	}
}
func (a *auth) addPermission(permission ...string) {
	if a.permission == nil {
		a.permission = mapset.NewThreadUnsafeSet()
	}
	for _, val := range permission {
		a.permission.Add(val)
	}
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

//添加一个permission
func (v validate) hasUrl(url string) bool {
	if _, ok := v[url]; !ok {
		//不存在
		return false
	}
	return true
}
