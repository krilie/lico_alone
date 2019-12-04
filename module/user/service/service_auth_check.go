package service

// 权限接口 动态加载

func (s *Service) HasUser(id string) (bool, error) {
	panic("implement me")
}

func (s *Service) HasPermission(id, permission string) (bool, error) {
	panic("implement me")
}

func (s *Service) HasRole(userId, roleId string) (bool, error) {
	panic("implement me")
}
