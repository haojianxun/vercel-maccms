package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/service"
	"goapi/pkg/helpers"
	"net/http"
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

func GetDATA(c *gin.Context) gin.H {
	value, _ := c.Get("data")
	data := value.(gin.H)
	return data
}

// PaginationHTML 生成搜索结果页面的分页 HTML 代码。
// 它接受当前页码、总页数和搜索关键字作为输入。
func PaginationHTML(currentPage, totalPages int, pageType string, params map[string]string) string {
	// 初始化 HTML 字符串
	html := "\n"
	// search-:电影名称-:演员名称-:大陆-:导演-:年份-:第几页.html
	urlsParams := pageType +
		"-" + params["movie"] +
		"-" + params["actor"] +
		"-" + params["area"] +
		"-" + params["director"] +
		"-" + params["year"]
	//"-" + params["pageNum"]
	fmt.Println(urlsParams)
	// 首页链接
	html += fmt.Sprintf("<a href=\"/%s-1.html\" class=\"page-number page-previous\" title=\"首页\">首页</a>\n", urlsParams)

	// 上一页链接
	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	html += fmt.Sprintf("<a href=\"/%s-%d.html\" class=\"page-number page-previous\" title=\"上一页\">上一页</a>\n", urlsParams, prevPage)

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
			html += fmt.Sprintf("<a href=\"/%s-%d.html\" class=\"page-number display\" title=\"第%d页\">%d</a>\n", urlsParams, i, i, i)
		}
	}

	// 下一页链接
	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}
	html += fmt.Sprintf("<a href=\"/%s-%d.html\" class=\"page-number page-next\" title=\"下一页\">下一页</a>\n", urlsParams, nextPage)

	// 尾页链接
	html += fmt.Sprintf("<a href=\"/%s-%d.html\" class=\"page-number page-next\" title=\"尾页\">尾页</a>\n", urlsParams, totalPages)
	return html
}

func extractParameters(url string) (movie, actor, area, director string, pageNum, year int64) {
	// search-:电影名称-:演员名称-:大陆-:导演-:年份-:第几页.html
	strParams := strings.ReplaceAll(url, ".html", "")
	params := strings.Split(strParams, "-")
	if len(params) > 1 {
		movie = params[1]
	}
	if len(params) > 2 {
		actor = params[2]
	}
	if len(params) > 3 {
		area = params[3]
	}
	if len(params) > 4 {
		director = params[4]
	}
	if len(params) > 5 {
		year = helpers.StringToInt64(params[5])
	}
	if len(params) > 6 {
		pageNum = helpers.StringToInt64(params[6])
		if pageNum == 0 {
			pageNum = 1
		}
	}
	return movie, actor, area, director, pageNum, year
}

func PageMs(c *gin.Context) {
	startTime, _ := c.Get("startTime")
	fmt.Println("当前耗时", time.Now().UnixMilli()-startTime.(int64), "ms")
}

func NoPage(c *gin.Context) {
	DATA := service.Site(c)
	c.HTML(http.StatusNotFound, "404", DATA)
}
