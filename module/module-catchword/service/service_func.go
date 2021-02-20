package service

import (
	"context"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	model2 "github.com/krilie/lico_alone/module/module-catchword/model"
	"gorm.io/gorm"
	"time"
)

func (a *CatchwordModule) AddCatchword(ctx context.Context, model *model2.AddCatchwordModel) (string, error) {
	var id = id_util.GetUuid()
	err := a.Dao.CreateCatchword(ctx, &model2.Catchword{
		Model: com_model.Model{
			Id:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{Valid: false},
		},
		Sort:    model.Sort,
		Title:   model.Title,
		Content: model.Content,
	})
	return id, err
}

func (a *CatchwordModule) DeleteCatchword(ctx context.Context, id string) error {
	err := a.Dao.DeleteCatchwordById(ctx, id)
	return err
}

func (a *CatchwordModule) UpdateCatchword(ctx context.Context, model *model2.UpdateCatchwordModel) error {
	return a.Dao.UpdateCatchwordById(ctx, &model2.Catchword{
		Model:   com_model.Model{Id: model.Id},
		Sort:    model.Sort,
		Title:   model.Title,
		Content: model.Content,
	})
}
