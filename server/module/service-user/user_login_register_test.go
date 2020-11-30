package service_user

import (
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-message/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserService_UserRegister(t *testing.T) {
	container.MustInvoke(func(userService *UserService) {
		err := userService.UserRegister(context.NewContext(), "mobile", "123", "123", "123")
		t.Log(err)
		assert.NotNil(t, err, "注册失败")
	})
}

func TestUserService_UserRegister2(t *testing.T) {
	container.MustInvoke(func(userService *UserService) {
		id := id_util.GetUuid()
		ctx := context.NewContext()
		phone := id_util.NextSnowflake()
		code := "123"
		err := userService.moduleMsg.Dao.CreateMessageValidCode(ctx, &model.MessageValidCode{
			Model:    com_model.Model{Id: id, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: gorm.DeletedAt{}},
			SendTime: time.Now(), PhoneNum: phone, Code: code, Type: model.MessageValidCodeTypeRegister.ToInt(),
		})
		CheckErr(t, err)
		err = userService.UserRegister(ctx, phone, "123", code, "")
		CheckErr(t, err)
		err = userService.moduleMsg.Dao.DeleteMessageValidCode(ctx, id)
		CheckErr(t, err)
		err = userService.moduleUser.Dao.DeleteUserByPhone(ctx, phone)
		CheckErr(t, err)
	})
}

func CheckErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
