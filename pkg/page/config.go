package page

import (
	"fmt"
	"math"
	"net/url"
)

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

// PaginationHTML 生成搜索结果页面的分页 HTML 代码。
// 它接受当前页码、总页数和搜索关键字作为输入。
func PaginationHTML(currentPage, totalPages int, keyword string) string {
	// 对关键字进行 URL 编码
	encodedKeyword := url.QueryEscape(keyword)
	// 初始化 HTML 字符串
	html := "\n"
	// 首页链接
	html += fmt.Sprintf("<a href=\"/search-%s----------1---.html\" class=\"page-number page-previous\" title=\"首页\">首页</a>\n", encodedKeyword)

	// 上一页链接
	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	html += fmt.Sprintf("<a href=\"/search-%s----------%d---.html\" class=\"page-number page-previous\" title=\"上一页\">上一页</a>\n", encodedKeyword, prevPage)

	// 计算分页的起始页和结束页
	startPage := currentPage - 2
	endPage := currentPage + 2

	// 调整起始页和结束页，确保最多显示 5 个页码链接
	if startPage < 1 {
		startPage = 1
		endPage = 5
	}

	if endPage > totalPages {
		endPage = totalPages
		if totalPages >= 5 {
			startPage = totalPages - 4
		}
	}

	// 生成页码链接
	for i := startPage; i <= endPage; i++ {
		if i == currentPage {
			// 高亮显示当前页
			html += fmt.Sprintf("<span class=\"page-number page-current display\">%d</span>\n", currentPage)
		} else {
			// 生成其他页码的链接
			html += fmt.Sprintf("<a href=\"/search-%s----------%d---.html\" class=\"page-number display\" title=\"第%d页\">%d</a>\n", encodedKeyword, i, i, i)
		}
	}

	// 下一页链接
	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}
	html += fmt.Sprintf("<a href=\"/search-%s----------%d---.html\" class=\"page-number page-next\" title=\"下一页\">下一页</a>\n", encodedKeyword, nextPage)

	// 尾页链接
	html += fmt.Sprintf("<a href=\"/search-%s----------%d---.html\" class=\"page-number page-next\" title=\"尾页\">尾页</a>\n", encodedKeyword, totalPages)
	return html
}
