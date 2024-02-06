package v1

import (
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service"
	"goapi/pkg/mysql"
	"net/http"
	"time"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	table := "index.html"
	PageData := cmap.New().Items()
	value, _ := c.Get("data")
	DATA := value.(gin.H)
	// 获取当前时间
	now := time.Now()
	// 获取当天凌晨00:00的时间戳
	midnightUnix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	// 批量查询数据
	var (
		TodayCount                                                                    int64
		CurrentlyTrending, listDianYing, MonthDianYing, listDianShiJu, MonthDianShiJu []models.MacVod
		listDongMan, MonthDongMan, listZongYi, MonthZongYi, NewsAll, NewsVideos       []models.MacVod
		NewsDianShiJu, NewsDongMan, NewsZongYi                                        []models.MacVod
	)
	DB := models.MacVodMgr(mysql.DB)

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
	service.ListMacVod(table, "listDianYing", 1, 14, &listDianYing)
	// 电影月榜
	service.MonthListMacVod(table, "MonthDianYing", 1, 10, &MonthDianYing)

	// 首页电视剧
	service.ListMacVod(table, "listDianShiJu", 2, 14, &listDianShiJu)
	// 电视剧月榜
	service.MonthListMacVod(table, "MonthDianShiJu", 2, 10, &MonthDianShiJu)

	// 首页动漫
	service.ListMacVod(table, "listDongMan", 4, 14, &listDongMan)
	// 动漫月榜
	service.MonthListMacVod(table, "MonthDongMan", 4, 10, &MonthDongMan)

	// 首页综艺
	service.ListMacVod(table, "listZongYi", 3, 14, &listZongYi)
	// 综艺月榜
	service.MonthListMacVod(table, "MonthZongYi", 3, 10, &MonthZongYi)

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
	c.HTML(http.StatusOK, "index.html", DATA)
}

func (h *IndexController) Web(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "web"
	data["title"] = "Web"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "web.html", data)
}

func (h *IndexController) App(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "app"
	data["title"] = "App"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "label/app.html", data)
}

func (h *IndexController) Search(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "search"
	data["title"] = "Search"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "search.html", data)
}
