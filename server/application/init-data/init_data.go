package init_data

import (
	"context"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/model"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/module/config/model"
	model2 "github.com/krilie/lico_alone/module/user/model"
	"time"
)

// InitData
func (a *Init) InitData(ctx context.Context) {
	// 开始事务
	err := cdb.WithTrans(ctx, a, func(ctx context.Context, service cdb.Service) error {
		a := service.(*Init)
		// 是否初始化过
		isInit, err := a.ConfigService.GetValueBool(ctx, model.CommonIsInitData)
		if err != nil {
			return err
		}
		if isInit == nil {
			err := a.ConfigService.SetValueBool(ctx, model.CommonIsInitData, false)
			if err != nil {
				return err
			}
		} else {
			if *isInit {
				return nil
			}
		}
		err = a.ConfigService.SetValueBool(ctx, model.CommonIsInitData, true)
		if err != nil {
			return err
		}
		// 初始化第一个角色
		err = a.UserService.Dao.CreateRole(ctx, &model2.Role{
			Name:        "root",
			ParentName:  "",
			CreateTime:  time.Now(),
			Description: "root role",
		})
		if err != nil {
			return err
		}
		// 初始化权限
		perms := initPermissionData()
		for e := range perms {
			err := a.UserService.Dao.CreatePerm(ctx, perms[e])
			if err != nil {
				return err
			}
			err = a.UserService.Dao.CreateRolePerm(ctx, "root", perms[e].Name)
			if err != nil {
				return err
			}
		}
		// 初始化第一个用户
		adminId := "00001"
		salt := pswd_util.GetSalt(6)
		user := model2.UserMaster{
			Model:      model.Model{Id: adminId, CreateTime: time.Now()},
			UpdateTime: time.Now(),
			LoginName:  "admin",
			PhoneNum:   "",
			Email:      nil,
			Password:   pswd_util.GetMd5Password("admin", salt),
			Picture:    "",
			Salt:       salt,
		}
		err = a.UserService.Dao.CreateUserMaster(ctx, &user)
		if err != nil {
			return err
		}
		// 关联权限用户
		err = a.UserService.Dao.CreateUserRole(adminId, "root")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(errs.NewInternal().WithError(err))
	}
}

// initPermissionData 权限初始数据
func initPermissionData() []*model2.Permission {
	var newPerm = func(name, des, method, path string, sort int) *model2.Permission {
		return &model2.Permission{
			Name:        name,
			CreateTime:  time.Now(),
			Description: des,
			RefMethod:   method,
			RefPath:     path,
			Sort:        sort,
		}
	}
	return []*model2.Permission{
		newPerm("user.update_info", "更新信息", "POST", "/api/user/update_info", 1),
	}
}
