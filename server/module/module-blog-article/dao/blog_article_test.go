package dao

import (
	"fmt"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"testing"
	"time"
)

func TestBlogArticleDao_CreateArticle(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		err := dao.CreateArticle(context.NewContext(), &model.Article{
			Model: com_model.Model{
				ID:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Title:   "test",
			Pv:      0,
			Content: "test",
			Picture: "test",
		})
		fmt.Println(err)
	})
}

func TestBlogArticleDao_DeleteArticleById(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		id, err := dao.DeleteArticleById(context.NewContext(), "11")
		fmt.Println(id)
		fmt.Println(err)
	})
}
