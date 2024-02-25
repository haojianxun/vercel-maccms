package page

// PageList 分页返回数
type PageList struct {
	CurrentPage int64       `json:"current_page"`
	Total       int64       `json:"total"`
	List        interface{} `json:"list"`
	PageNum     int64       `json:"pageNum"`  // 第几页
	PageSize    int64       `json:"pageSize"` // 每页多少条
	Offset      int64       `json:"-"`
}

// InitPageList 设置分页参数
func InitPageList(lists *PageList) {
	// 当前页数
	if lists.CurrentPage <= 0 {
		lists.CurrentPage = 1
	} // 第一页
	if lists.PageNum == 0 {
		lists.PageNum = 1
	}
	// 每页取出多少条数据，默认10条
	if lists.PageSize == 0 {
		lists.PageSize = 10
	}

	lists.Offset = (lists.CurrentPage - 1) * lists.PageSize
	if lists.Offset < 0 {
		lists.Offset = 0
	}
}
