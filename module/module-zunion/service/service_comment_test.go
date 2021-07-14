package service

import (
	"context"
	"github.com/krilie/lico_alone/common/appdig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-zunion/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvides(DigModuleZUnionProviderAll)

func TestUserService_RegisterNewUser(t *testing.T) {
	container.MustInvoke(func(svc *ZUnionModule) {
		err := svc.AddComment(context.Background(), &model.TbComment{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    gorm.DeletedAt{Valid: false},
			Id:           id_util.GetUuid(),
			UserId:       id_util.GetUuid(),
			CommentId:    id_util.GetUuid(),
			TargetId:     id_util.GetUuid(),
			Content:      "test 123",
			LikeCount:    23,
			DislikeCount: 1,
			IsCheck:      false,
		})
		println(err)
	})
}
