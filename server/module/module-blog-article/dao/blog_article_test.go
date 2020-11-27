package dao

import (
	"fmt"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/module/module-blog-article/model"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	component.DigProviderTest()
	DigProvider()
	m.Run()
}

func TestBlogArticleDao_CreateArticle(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		err := dao.CreateArticle(context.NewContext(), &model.Article{
			Model: com_model.Model{
				Id:        id_util.GetUuid(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			Title:       "123" + id_util.NextSnowflake(),
			Pv:          0,
			Content:     "123" + id_util.NextSnowflake(),
			Picture:     "123" + id_util.NextSnowflake(),
			Description: "123" + id_util.NextSnowflake(),
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
		ctx := context.NewContext()
		err := dao.UpdateArticle(ctx, &model.Article{
			Model: com_model.Model{
				Id:        "b93b348a-0d93-45d9-9cc4-5b1e4fe5407b",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			Title:       "123" + id_util.NextSnowflake(),
			Pv:          10,
			Content:     "231" + id_util.NextSnowflake(),
			Picture:     "123" + id_util.NextSnowflake(),
			Description: "11234" + id_util.NextSnowflake(),
		})
		dao.log.Get(ctx).Info("info")
		t.Log(err)
	})
}

func TestBlogArticleDao_QueryArticleById(t *testing.T) {
	dig.Container.MustInvoke(func(dao *BlogArticleDao) {
		id, err := dao.GetArticleById(context.NewContext(), "12")
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
