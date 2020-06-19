package com_model

// PaginationParam 分页查询条件
// @Param page_num query int true "page_num页索引"
// @Param page_size query int true "page_size页大小"
type PageParams struct {
	PageNum  int `form:"page_num" json:"page_num" xml:"page_num"  binding:"required"`    // 页索引
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"  binding:"required"` // 页大小
}

func (p *PageParams) CheckOkOrSetDefault() {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
}

// PaginationResult 分页查询结果
type PageInfo struct {
	TotalCount int `json:"total_count" swaggo:"true,总条数"` // 总数据条数
	TotalPage  int `json:"total_page" swaggo:"true,所有页数"`
	PageNum    int `json:"page_num" swaggo:"true,当前页码"`
	PageSize   int `json:"page_size" swggo:"true,页大小"`
}

type PageData struct {
	PageInfo PageInfo    `json:"page_info"` // 分页信息
	Data     interface{} `json:"data"`      // 列表数据
}
