package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/sqlutil"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/module/module-catchword/model"
	"time"
)

// QueryList 查询列表
func (t *CatchwordDao) QueryList(ctx context.Context, keyWord string, pageParam com_model.PageParams) (pageInfo *com_model.PageInfo, data []*model.Catchword, err error) {
	sqlBuilder := sq.Select().
		From("tb_catchword").
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Or{sq.Like{"title": sqlutil.Like(keyWord)}, sq.Like{"content": sqlutil.Like(keyWord)}})
	// countSql
	countSql, countArgs := sqlBuilder.Columns("count(1)").MustSql()
	// dataSql
	dataSql, dataArgs := sqlBuilder.
		Columns("*").
		OrderBy("created_at desc,id asc").
		Offset(uint64(pageParam.OffSet())).
		Limit(uint64(pageParam.Limit())).
		MustSql()
	// query total and data
	t.log.Get(ctx).WithField("sql", dataSql).WithField("args", dataArgs).Info("sql str build")

	data = make([]*model.Catchword, 0)
	totalCount, err := ndb.GetPageDataFormSql(ctx, t.GetDb(ctx), countSql, dataSql, countArgs, dataArgs, &data)
	if err != nil {
		return nil, nil, err
	}
	return com_model.NewPageInfo(pageParam.PageNum, pageParam.PageSize, totalCount), data, err
}

// QueryListForWebShow 查询列表
func (t *CatchwordDao) QueryListForWebShow(ctx context.Context, keyWord string, from time.Time, limit int) (data []*model.Catchword, err error) {
	sql, args, err := sq.Select("*").
		From("tb_catchword").
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Or{sq.Like{"title": sqlutil.Like(keyWord)}, sq.Like{"content": sqlutil.Like(keyWord)}}).
		Where(sq.LtOrEq{"create_at": from}).
		OrderBy("create_at desc,id asc").
		Limit(uint64(limit)).
		ToSql()
	if err != nil {
		panic(err)
	}
	t.log.Get(ctx).WithField("sql", sql).WithField("args", args).Info("sql str build")
	data = make([]*model.Catchword, 0)
	err = t.GetDb(ctx).Raw(sql, args...).Scan(&data).Error
	return data, err
}
