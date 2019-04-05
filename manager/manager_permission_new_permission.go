package manager

import (
	"github.com/asaskevich/govalidator"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/common/context_util"
	"github.com/lico603/lico_user/common/log"
	"github.com/lico603/lico_user/common/uuid_util"
	"github.com/lico603/lico_user/model"
	"time"
)

//给系统添加新的permission permission项,默认权限检查已经通过
func ManagerPermissionNewPermission(ctx *context_util.Context, pName string, pDescription string) (p *model.Permission, err error) {
	//检查参数
	if !govalidator.IsAlpha(pName) || len(pDescription) == 0 {
		log.Infoln("", "param error:", pName, pDescription)
		return nil, errs.ErrParam
	}
	//添加一个
	p = new(model.Permission)
	p.ID = uuid_util.GetUuid()
	p.CreateTime = time.Now()
	p.Description = pDescription
	p.Name = pName
	p.Version = 0
	err = model.Db.Create(p).Error
	if err != nil {
		return nil, err
	} else {
		return p, nil
	}
}
