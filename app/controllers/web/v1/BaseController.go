package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/service"
	"goapi/pkg/helpers"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"net/http"
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

func GetDATA(c *gin.Context) gin.H {
	value, _ := c.Get("data")
	data := value.(gin.H)
	return data
}

// PaginationHTML 生成搜索结果页面的分页 HTML 代码。
// 它接受当前页码、总页数和搜索关键字作为输入。
func PaginationHTML(currentPage, totalPages int, pageType, keyword string) string {
	// 对关键字进行 URL 编码
	encodedKeyword := url.QueryEscape(keyword)
	// 初始化 HTML 字符串
	html := "\n"
	// 首页链接
	html += fmt.Sprintf("<a href=\"/%s-%s----------1---.html\" class=\"page-number page-previous\" title=\"首页\">首页</a>\n", pageType, encodedKeyword)

	// 上一页链接
	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	html += fmt.Sprintf("<a href=\"/%s-%s----------%d---.html\" class=\"page-number page-previous\" title=\"上一页\">上一页</a>\n", pageType, encodedKeyword, prevPage)

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
			html += fmt.Sprintf("<a href=\"/%s-%s----------%d---.html\" class=\"page-number display\" title=\"第%d页\">%d</a>\n", pageType, encodedKeyword, i, i, i)
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

func NoPage(c *gin.Context) {
	DATA := service.Site(c)
	c.HTML(http.StatusNotFound, "404", DATA)
}

// 顶级栏目处理

func TopCategory(c *gin.Context, table string, DATA, PageData map[string]interface{}, macType models.MacType, listMacType []models.MacType, CurrentlyTrending []models.MacVod) {
	service.ListMacType(table, int(macType.TypeID), &listMacType)
	// 正在热播
	service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
		"type_id_1":  macType.TypeID,
		"vod_status": 1,
	}, "vod_hits desc", 16, &CurrentlyTrending)
	PageData["CurrentlyTrending"] = CurrentlyTrending
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
		PageData[Name] = BindList
	}
	DATA["subCategory"] = 1
	PageData["listMacType"] = listMacType
	DATA["PageData"] = PageData
	c.HTML(http.StatusOK, "v/category.html", DATA)
}

// 子栏目处理

func SubCategory(c *gin.Context, table string, DATA, PageData map[string]interface{}, macType models.MacType, listMacType []models.MacType) {
	service.ListMacType(table, int(macType.TypePid), &listMacType)
	var pageList page.PageList // 返回数据
	var params requests.Search // 搜索数据
	params.PageSize = 72       // 每页默认72条数据
	whereSub := cmap.New().Items()
	whereSub["type_id"] = macType.TypeID
	whereSub["vod_status"] = 1
	// 获取路由中的参数值
	models.MacVodMgr(mysql.DB).Where(whereSub).Count(&pageList.Total)
	// 设置分页参数
	pageList.CurrentPage = params.PageNum
	pageList.PageSize = params.PageSize
	page.InitPageList(&pageList)
	var listResult []models.MacVod
	models.MacVodMgr(mysql.DB).
		Where(whereSub).
		Offset(int(pageList.Offset)).
		Limit(int(pageList.PageSize)).
		Find(&listResult)
	pageList.List = listResult
	PageData["pageList"] = pageList
	PageData["PaginationHTML"] = PaginationHTML(int(pageList.CurrentPage), int(pageList.PageTotal), macType.TypeEn, params.Name)
	DATA["subCategory"] = 0
	PageData["listMacType"] = listMacType
	DATA["PageData"] = PageData
	c.HTML(http.StatusOK, "v/category.html", DATA)
}
