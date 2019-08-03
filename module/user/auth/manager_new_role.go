package auth

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

//给系统添加新的角色
func (UserManage) NewRole(ctx context.Context, roleName string, roleDescription string) (role *model.Role, err error) {
	//检查参数
	if !govalidator.IsAlpha(roleName) || len(roleDescription) == 0 {
		log.Infoln("CreateNewRole", "param err:", roleName, roleDescription)
		return nil, errs.ErrParam
	}
	//生成新角色并添加
	role = new(model.Role)
	role.ID = id_util.GetUuid()
	role.Name = roleName
	role.Description = roleDescription
	role.CreateTime = time.Now()
	err = dao.Db.Create(role).Error
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}
