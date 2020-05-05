package service

import (
	"context"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/module-user/model"
	"time"
)

// InitData 初始化数据
func (s *UserService) InitUserData(ctx context.Context) (err error) {
	err = s.Dao.Transaction(ctx, func(ctx context.Context) error {
		err = s.Dao.DeleteAllUserData(ctx)
		if err != nil {
			return err
		}
		admin := getInitAdminUserData()
		// 用户
		err = s.Dao.CreateUserMaster(ctx, &admin.user)
		if err != nil {
			return err
		}
		// 角色
		err = s.Dao.CreateRole(ctx, &admin.role)
		if err != nil {
			return err
		}
		// 权限
		err = s.Dao.CreatePerms(ctx, admin.permissions)
		if err != nil {
			return err
		}
		// 用户角色
		err = s.Dao.CreateUserRole(ctx, admin.user.Id, admin.role.Name)
		if err != nil {
			return err
		}
		// 角色权限
		for _, permission := range admin.permissions {
			err = s.Dao.CreateRolePerm(ctx, admin.role.Name, permission.Name)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// initAdminUserStruct 初始化结构
type initAdminUserStruct struct {
	user        model.UserMaster
	role        model.Role
	permissions []model.Permission
}

func NewPerm(name, method, path, desc string, sort int) model.Permission {
	return model.Permission{
		Name:        name,
		CreateTime:  time.Now(),
		Description: desc,
		RefMethod:   method,
		RefPath:     path,
		Sort:        sort,
	}
}

func getInitAdminUserData() initAdminUserStruct {
	return initAdminUserStruct{
		user: model.UserMaster{
			Model:      com_model.Model{Id: "00001", CreateTime: time.Now()},
			UpdateTime: time.Now(), LoginName: "admin", PhoneNum: "", Email: "",
			Password: pswd_util.GetMd5Password("123456", "2345r"),
			Picture:  "", Salt: "2345r",
		},
		role: model.Role{
			Name: "admin", CreateTime: time.Now(), Description: "超级管理员 初始勿动", MainPermissionName: "",
		},
		permissions: []model.Permission{
			NewPerm("test", "get", "test", "测试用", 0),
		},
	}
}
