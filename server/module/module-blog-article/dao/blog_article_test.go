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
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Title:       "test",
			Pv:          0,
			Content:     "test",
			Picture:     "test",
			Description: "",
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

func TestBlogArticleDao_UpdateArticle(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		err := dao.UpdateArticle(context.NewContext(), &model.Article{
			Model: com_model.Model{
				Id:        "b93b348a-0d93-45d9-9cc4-5b1e4fe5407b",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Title:       "123",
			Pv:          10,
			Content:     "231",
			Picture:     "123",
			Description: "",
		})
		dao.log.Get(ctx).Info("info")
		t.Log(err)
	})
}

func TestBlogArticleDao_QueryArticleById(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		id, err := dao.QueryArticleById(context.NewContext(), "12")
		t.Log(id, err)
	})
}

func TestBlogArticleDao_UpdateArticleSample(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		err := dao.UpdateArticleSample(context.NewContext(), &model.UpdateArticleModel{
			Id:      "11",
			Title:   "22",
			Content: "33",
			Picture: "44",
		})
		t.Log(err)
	})
}
