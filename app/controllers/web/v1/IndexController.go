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
	IndexData := cmap.New().Items()
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
	IndexData["TodayCount"] = TodayCount

	// 首页电影
	CachelistDianYing, _ := service.GetTable(table, "listDianYing", []models.MacVod{})
	if CachelistDianYing == nil {
		DB.Debug().Where("type_id_1", 1).Order("vod_hits desc").Limit(14).Find(&listDianYing)
		service.SaveTable(table, "listDianYing", listDianYing)
	} else {
		listDianYing = *CachelistDianYing.(*[]models.MacVod)
	}

	// 电影月榜
	CacheMonthDianYing, _ := service.GetTable(table, "MonthDianYing", []models.MacVod{})
	if CacheMonthDianYing == nil {
		DB.Debug().Where("type_id_1", 1).Order("vod_hits_month desc").Limit(10).Find(&MonthDianYing)
		service.SaveTable(table, "MonthDianYing", MonthDianYing)
	} else {
		MonthDianYing = *CacheMonthDianYing.(*[]models.MacVod)
	}

	// 首页电视剧
	CachelistDianShiJu, _ := service.GetTable(table, "listDianShiJu", []models.MacVod{})
	if CachelistDianShiJu == nil {
		DB.Debug().Where("type_id_1", 2).Order("vod_hits desc").Limit(14).Find(&listDianShiJu)
		service.SaveTable(table, "listDianShiJu", listDianShiJu)
	} else {
		listDianShiJu = *CachelistDianShiJu.(*[]models.MacVod)
	}

	// 电视剧月榜
	CacheMonthDianShiJu, _ := service.GetTable(table, "MonthDianShiJu", []models.MacVod{})
	if CacheMonthDianShiJu == nil {
		DB.Debug().Where("type_id_1", 2).Order("vod_hits_month desc").Limit(10).Find(&MonthDianShiJu)
		service.SaveTable(table, "MonthDianShiJu", MonthDianShiJu)
	} else {
		MonthDianShiJu = *CacheMonthDianShiJu.(*[]models.MacVod)
	}

	// 首页动漫
	CachelistDongManu, _ := service.GetTable(table, "listDongMan", []models.MacVod{})
	if CachelistDongManu == nil {
		DB.Debug().Where("type_id", 4).Order("vod_hits desc").Limit(14).Find(&listDongMan)
		service.SaveTable(table, "listDongMan", listDongMan)
	} else {
		listDongMan = *CachelistDongManu.(*[]models.MacVod)
	}

	// 动漫月榜
	CacheMonthDongMan, _ := service.GetTable(table, "MonthDongMan", []models.MacVod{})
	if CacheMonthDongMan == nil {
		DB.Debug().Where("type_id", 4).Order("vod_hits_month desc").Limit(10).Find(&MonthDongMan)
		service.SaveTable(table, "MonthDongMan", MonthDongMan)
	} else {
		MonthDongMan = *CacheMonthDongMan.(*[]models.MacVod)
	}

	// 首页综艺
	CachelistZongYi, _ := service.GetTable(table, "listZongYi", []models.MacVod{})
	if CachelistZongYi == nil {
		DB.Debug().Where("type_id", 3).Order("vod_hits desc").Limit(14).Find(&listZongYi)
		service.SaveTable(table, "listZongYi", listZongYi)
	} else {
		listZongYi = *CachelistZongYi.(*[]models.MacVod)
	}

	// 综艺月榜
	CacheMonthZongYi, _ := service.GetTable(table, "MonthZongYi", []models.MacVod{})
	if CacheMonthZongYi == nil {
		DB.Debug().Where("type_id", 3).Order("vod_hits_month desc").Limit(10).Find(&MonthZongYi)
		service.SaveTable(table, "MonthZongYi", MonthZongYi)
	} else {
		MonthZongYi = *CacheMonthZongYi.(*[]models.MacVod)
	}

	// 最新影片
	listNewVideos := make(map[string][]models.MacVod)
	CachelistNewVideos, _ := service.GetTable(table, "listNewVideos", map[string][]models.MacVod{})
	if CachelistNewVideos == nil {
		DB.Debug().Order("vod_time desc").Limit(8).Find(&NewsAll)
		DB.Debug().Where("type_id_1", 1).Order("vod_time desc").Limit(8).Find(&NewsVideos)
		DB.Debug().Where("type_id_1", 2).Order("vod_time desc").Limit(8).Find(&NewsDianShiJu)
		DB.Debug().Where("type_id", 4).Order("vod_time desc").Limit(8).Find(&NewsDongMan)
		DB.Debug().Where("type_id", 3).Order("vod_time desc").Limit(8).Find(&NewsZongYi)
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

	IndexData["CurrentlyTrending"] = CurrentlyTrending
	IndexData["listDianYing"] = listDianYing
	IndexData["MonthDianYing"] = MonthDianYing
	IndexData["listDianShiJu"] = listDianShiJu
	IndexData["MonthDianShiJu"] = MonthDianShiJu
	IndexData["listDongMan"] = listDongMan
	IndexData["MonthDongMan"] = MonthDongMan
	IndexData["listZongYi"] = listZongYi
	IndexData["MonthZongYi"] = MonthZongYi
	IndexData["listNewVideos"] = listNewVideos
	DATA["INDEX_DATA"] = IndexData
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
