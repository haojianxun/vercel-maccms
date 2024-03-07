package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/service"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"net/http"
	"strings"
	"time"
)

type IndexController struct {
	BaseController
}

// PresTrain loading 加载页面
func (h *IndexController) PresTrain(c *gin.Context) {
	DATA := GetDATA(c)
	c.HTML(http.StatusOK, "prestrain.html", DATA)
}

func (h *IndexController) DPlayer(c *gin.Context) {
	DATA := GetDATA(c)
	c.HTML(http.StatusOK, "dplayer.html", DATA)
}

func (h *IndexController) Web(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "web"
	DATA["title"] = "Web"
	PageMs(c)
	c.HTML(http.StatusOK, "web.html", DATA)
}

func (h *IndexController) About(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "about"
	DATA["title"] = "About"
	PageMs(c)
	c.HTML(http.StatusOK, "about.html", DATA)
}

func (h *IndexController) App(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "app"
	DATA["title"] = "App"
	PageMs(c)
	c.HTML(http.StatusOK, "label/app.html", DATA)
}

func (h *IndexController) Robots(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "robots"
	DATA["title"] = "Robots"
	PageMs(c)
	c.HTML(http.StatusOK, "robots.txt", DATA)
}

func (h *IndexController) Google(c *gin.Context) {
	c.String(http.StatusOK, "google-site-verification: google8fcf6226304ee6af.html")
}

func (h *IndexController) Index(c *gin.Context) {
	table := "index.html"
	PageData := cmap.New().Items()
	DATA := GetDATA(c)
	// 获取当前时间
	now := time.Now()
	// 获取当天凌晨00:00的时间戳
	midnightUnix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	// 批量查询数据
	var (
		TodayCount                                                                    int64
		listMacType                                                                   []models.MacType
		CurrentlyTrending, listDianYing, MonthDianYing, listDianShiJu, MonthDianShiJu []models.MacVod
		listDongMan, MonthDongMan, listZongYi, MonthZongYi, NewsAll, NewsVideos       []models.MacVod
		NewsDianShiJu, NewsDongMan, NewsZongYi                                        []models.MacVod
	)
	DB := models.MacVodMgr(mysql.DB)
	// 查询所有分类
	service.ListMacType("common", -1, &listMacType)
	// 正在热播
	CacheCurrentlyTrending, _ := service.GetTable(table, "CurrentlyTrending", []models.MacVod{})
	if CacheCurrentlyTrending == nil {
		DB.Debug().Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
		service.SaveTable(table, "CurrentlyTrending", CurrentlyTrending)
	} else {
		CurrentlyTrending = *CacheCurrentlyTrending.(*[]models.MacVod)
	}

	// 今日更新
	DB.Debug().Where("vod_time >= ?", midnightUnix).Count(&TodayCount)
	PageData["TodayCount"] = TodayCount

	// 首页电影
	service.ListWhereMacVod(table, "listDianYing", map[string]interface{}{
		"type_id_1":  1,
		"vod_status": 1,
	}, "vod_hits desc", 14, &listDianYing)
	// 电影月榜
	service.ListWhereMacVod(table, "MonthDianYing", map[string]interface{}{
		"type_id_1":  1,
		"vod_status": 1,
	}, "vod_hits_month desc", 10, &MonthDianYing)

	// 首页电视剧
	service.ListWhereMacVod(table, "listDianShiJu", map[string]interface{}{
		"type_id_1":  2,
		"vod_status": 1,
	}, "vod_hits desc", 14, &listDianShiJu)
	// 电视剧月榜
	service.ListWhereMacVod(table, "MonthDianShiJu", map[string]interface{}{
		"type_id_1":  2,
		"vod_status": 1,
	}, "vod_hits_month desc", 10, &MonthDianShiJu)

	// 首页动漫
	service.ListWhereMacVod(table, "listDongMan", map[string]interface{}{
		"type_id_1":  4,
		"vod_status": 1,
	}, "vod_hits desc", 14, &listDongMan)
	// 动漫月榜
	service.ListWhereMacVod(table, "MonthDongMan", map[string]interface{}{
		"type_id_1":  4,
		"vod_status": 1,
	}, "vod_hits_month desc", 10, &MonthDongMan)

	// 首页综艺
	service.ListWhereMacVod(table, "listZongYi", map[string]interface{}{
		"type_id_1":  3,
		"vod_status": 1,
	}, "vod_hits desc", 14, &listZongYi)
	// 综艺月榜
	service.ListWhereMacVod(table, "MonthZongYi", map[string]interface{}{
		"type_id_1":  3,
		"vod_status": 1,
	}, "vod_hits_month desc", 10, &MonthZongYi)

	// 最新影片
	listNewVideos := make(map[string][]models.MacVod)
	CachelistNewVideos, _ := service.GetTable(table, "listNewVideos", map[string][]models.MacVod{})
	if CachelistNewVideos == nil {
		DB.Debug().Order("vod_time desc").Limit(8).Find(&NewsAll)
		DB.Debug().Where("type_id_1", 1).Order("vod_time desc").Limit(8).Find(&NewsVideos)
		DB.Debug().Where("type_id_1", 2).Order("vod_time desc").Limit(8).Find(&NewsDianShiJu)
		DB.Debug().Where("type_id_1", 4).Order("vod_time desc").Limit(8).Find(&NewsDongMan)
		DB.Debug().Where("type_id_1", 3).Order("vod_time desc").Limit(8).Find(&NewsZongYi)
		listNewVideos = map[string][]models.MacVod{
			"NewsAll":       NewsAll,
			"NewsVideos":    NewsVideos,
			"NewsDianShiJu": NewsDianShiJu,
			"NewsDongMan":   NewsDongMan,
			"NewsZongYi":    NewsZongYi,
		}
		service.SaveTable(table, "listNewVideos", listNewVideos)
	} else {
		listNewVideos = *CachelistNewVideos.(*map[string][]models.MacVod)
	}

	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending
	PageData["listDianYing"] = listDianYing
	PageData["MonthDianYing"] = MonthDianYing
	PageData["listDianShiJu"] = listDianShiJu
	PageData["MonthDianShiJu"] = MonthDianShiJu
	PageData["listDongMan"] = listDongMan
	PageData["MonthDongMan"] = MonthDongMan
	PageData["listZongYi"] = listZongYi
	PageData["MonthZongYi"] = MonthZongYi
	PageData["listNewVideos"] = listNewVideos
	DATA["PageData"] = PageData
	DATA["page"] = "index"
	PageMs(c)
	c.HTML(http.StatusOK, "index.html", DATA)
}

func (h *IndexController) Search(c *gin.Context) {
	var (
		params   requests.Search // 搜索数据
		pageList page.PageList   // 返回数据
	)
	params.Name = c.Query("wd")
	search := strings.ReplaceAll(c.Param("search"), ".html", "")
	movie, actor, area, director, pageNum, year := extractParameters(search)
	fmt.Printf("Movie: %s, Actor: %s, Area: %s, Director: %s, Page: %s, Year: %s\n", movie, actor, area, director, pageNum, year)
	fmt.Println(search)
	fmt.Println(search)
	fmt.Println(search)
	if len(movie) > 0 {
		params.Name = movie
	}
	if pageNum > 0 {
		params.PageNum = pageNum
	}
	//table := "search.html"
	PageData := cmap.New().Items()
	DATA := GetDATA(c)
	where := cmap.New().Items()
	where["vod_status"] = 1
	if len(actor) > 0 {
		where["vod_actor LIKE"] = fmt.Sprintf("%%%v%%", actor)
	}
	if len(director) > 0 {
		where["vod_director LIKE"] = fmt.Sprintf("%%%v%%", director)
	}
	if len(area) > 0 {
		where["vod_area LIKE"] = fmt.Sprintf("%%%v%%", area)
	}
	if year > 0 {
		where["vod_year"] = year
	}
	// 获取路由中的参数值
	models.MacVodMgr(mysql.DB).Where("vod_name LIKE ?", fmt.Sprintf("%%%v%%", params.Name)).Where(where).Count(&pageList.Total)
	// 设置分页参数
	pageList.CurrentPage = params.PageNum
	pageList.PageSize = params.PageSize
	page.InitPageList(&pageList)
	var listSearch []models.MacVod
	err := models.MacVodMgr(mysql.DB).
		Where(
			"vod_name LIKE ?",
			fmt.Sprintf("%%%v%%", params.Name),
		).
		Where(where).
		Offset(int(pageList.Offset)).
		Limit(int(pageList.PageSize)).
		Find(&listSearch).Error
	if err != nil {
		NoPage(c)
		return
	}
	pageList.List = listSearch
	PageData["pageList"] = pageList
	PageData["wd"] = params.Name
	PageData["PaginationHTML"] = PaginationHTML(int(pageList.CurrentPage), int(pageList.PageTotal), "search", params.Name)
	var listMacType []models.MacType
	// 查询所有分类
	service.ListMacType("common", -1, &listMacType)
	PageData["listMacType"] = listMacType

	DATA["PageData"] = PageData
	DATA["page"] = "search"
	c.HTML(http.StatusOK, "search.html", DATA)
}
