package page

import "math"

// PageList 分页返回数
type PageList struct {
	IndexPage    int64       `json:"index_page"`
	PagePrevious int64       `json:"page_previous"`
	CurrentPage  int64       `json:"current_page"`
	PageNext     int64       `json:"page_next"`
	Total        int64       `json:"total"`
	List         interface{} `json:"list"`
	PageNum      int64       `json:"pageNum"`   // 第几页
	PageSize     int64       `json:"pageSize"`  // 每页多少条
	PageTotal    int64       `json:"pageTotal"` // 总页数
	Offset       int64       `json:"-"`
}

// InitPageList 设置分页参数
func InitPageList(lists *PageList) {
	lists.IndexPage = 1
	lists.PagePrevious = 1
	lists.PageNext = 1
	// 当前页数
	if lists.CurrentPage <= 0 {
		lists.CurrentPage = 1
	}
	// 第一页
	if lists.PageNum == 0 {
		lists.PageNum = 1
	}
	// 每页取出多少条数据，默认10条
	if lists.PageSize == 0 {
		lists.PageSize = 10
	}
	// 总页数（最后一页）
	lists.PageTotal = int64(math.Ceil(float64(lists.Total) / float64(lists.PageSize)))
	// 上一页
	if lists.PageTotal > 1 && lists.CurrentPage > 1 {
		lists.PagePrevious = lists.CurrentPage - 1
	}
	// 下一页
	if lists.CurrentPage <= lists.PageTotal {
		lists.PageNext = lists.CurrentPage + 1
	}
	lists.Offset = (lists.CurrentPage - 1) * lists.PageSize
	if lists.Offset < 0 {
		lists.Offset = 0
	}
}
