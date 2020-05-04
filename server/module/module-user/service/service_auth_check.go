package service

// 权限接口 动态加载

func (s *UserService) HasUser(id string) (bool, error) {
	panic("implement me")
}

func (s *UserService) HasPermission(id, method, path string) (bool, error) {
	panic("implement me")
}

func (s *UserService) HasRole(userId, roleId string) (bool, error) {
	panic("implement me")
}
