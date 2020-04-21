package cmodel

// PaginationParam 分页查询条件
type PageParams struct {
	Index int `json:"index" swaggo:"false,页索引"` // 页索引
	Size  int `json:"size" swaggo:"false,页大小"`  // 页大小
}

// PaginationResult 分页查询结果
type PageInfo struct {
	TotalCount int `json:"total" swaggo:"true,总条数"` // 总数据条数
	TotalPage  int `json:"total_page" swaggo:"true,所有页数"`
	NowIndex   int `json:"current" swaggo:"true,当前页码"`
	PageSize   int `json:"page_size" swggo:"true,页大小"`
}

type PageData struct {
	PageInfo PageInfo    `json:"page_info"` // 分页信息
	Data     interface{} `json:"data"`      // 列表数据
}
