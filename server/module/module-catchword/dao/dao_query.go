package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/sql_util"
	"github.com/krilie/lico_alone/module/module-catchword/model"
)

// QueryList 查询列表
func (t *CatchwordDao) QueryList(ctx context.Context, keyWord string, pageParam com_model.PageParams) (data []*model.Catchword, err error) {
	sql, args, err := sq.Select("*").
		From("tb_catchword").
		Where(sq.Eq{"delete_at": nil}).
		Where(sq.Or{sq.Like{"title": sql_util.Like(keyWord)}, sq.Like{"content": sql_util.Like(keyWord)}}).
		ToSql()
	if err != nil {
		panic(err)
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("args", args).Info("sql str build")
	data = make([]*model.Catchword, 0)
	err = t.GetDb(ctx).Raw(sql, args).Scan(&data).Error
	return data, err
}
