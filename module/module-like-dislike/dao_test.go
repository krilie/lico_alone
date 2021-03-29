package module_like_dislike

import (
	"github.com/krilie/lico_alone/common/appdig"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/component"
	"github.com/stretchr/testify/assert"
	"testing"
)

var container = appdig.
	NewAppDig().
	MustProvides(component.DigComponentProviderAll).
	MustProvide(NewLikeDisLikeDao)

func TestNewLikeDisLikeDao(t *testing.T) {
	container.MustInvoke(func(dao *LikeDisLikeDao) {
		println(dao.Ping())
	})
}

func TestAutoNewLikeDisLikeDao(t *testing.T) {
	container.MustInvoke(func(dao *LikeDisLikeDao) {
		var ctx = context2.EmptyAppCtx()
		var params = LikeDisLikeModelParams{
			UserId:       id_util.NextSnowflake(),
			BusinessType: "123",
			BusinessId:   id_util.NextSnowflake(),
			GiveType:     "like",
		}
		// check if has
		has, err := dao.HasLikeDisLikeRecord(ctx, params)
		assert.Nil(t, err, "should no err")
		assert.False(t, has, "should false")
		// add record
		err = dao.AddLikeDisLikeRecord(ctx, params)
		assert.Nil(t, err, "should no err")
		// check if has
		has, err = dao.HasLikeDisLikeRecord(ctx, params)
		assert.Nil(t, err, "should no err")
		assert.True(t, has, "should true")
		// remove if has
		err = dao.RemoveLikeDisLikeRecord(ctx, params)
		assert.Nil(t, err, "should no err")
		// check if has
		has, err = dao.HasLikeDisLikeRecord(ctx, params)
		assert.Nil(t, err, "should no err")
		assert.False(t, has, "should false")
	})
}

func TestAutoLikeDisLikeDao_GetLikeDiskLikeResult(t *testing.T) {
	container.MustInvoke(func(dao *LikeDisLikeDao) {
		var ctx = context2.EmptyAppCtx()
		bId1 := id_util.GetUuid()
		bId2 := id_util.GetUuid()
		var params = []LikeDisLikeModelParams{
			{UserId: "1", BusinessType: "one", BusinessId: bId1, GiveType: "like"},
			{UserId: "2", BusinessType: "one", BusinessId: bId1, GiveType: "like"},
			{UserId: "3", BusinessType: "one", BusinessId: bId1, GiveType: "dislike"},
			{UserId: "4", BusinessType: "two", BusinessId: bId2, GiveType: "dislike"},
		}
		for _, param := range params {
			_ = dao.AddLikeDisLikeRecord(ctx, param)
		}
		result, err := dao.GetLikeDiskLikeResult(ctx, "one", []string{bId1})
		assert.Nil(t, err, "should no err")
		var expResult = []LikeDisLikeModelResult{
			{BusinessType: "one", BusinessId: bId1, GiveType: "like", Count: 2},
			{BusinessType: "one", BusinessId: bId1, GiveType: "dislike", Count: 1},
		}
		assert.Equal(t, expResult, result)
	})
}
