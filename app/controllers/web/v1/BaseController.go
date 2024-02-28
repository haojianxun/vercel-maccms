package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/helpers"
	"net/url"
	"strings"
	"time"
)

type BaseController struct {
}

type GroupV1 struct {
	BaseController
	IndexController
	VideosController
	ArttypeController
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

func extractParameters(url string) (movie, actor, area, director string, pageNum, year int64) {
	// 使用字符串分割来提取参数
	parts := strings.Split(url, "---")
	if len(parts) < 5 {
		return "", "", "", "", 1, 0
	}

	// 根据前缀判断参数类型
	switch {
	case strings.HasPrefix(parts[0], "-"):
		movie = strings.ReplaceAll(parts[0], "-", "")
	case strings.HasPrefix(parts[0], "--"):
		actor = strings.ReplaceAll(parts[0], "-", "")
	}
	area = strings.ReplaceAll(parts[1], "-", "")
	director = strings.ReplaceAll(parts[2], "-", "")
	pageNum = helpers.StringToInt64(strings.ReplaceAll(parts[3], "-", ""))
	year = helpers.StringToInt64(strings.ReplaceAll(parts[4], "-", ""))
	return movie, actor, area, director, pageNum, year
}

func PageMs(c *gin.Context) {
	startTime, _ := c.Get("startTime")
	fmt.Println("当前耗时", time.Now().UnixMilli()-startTime.(int64), "ms")
}
