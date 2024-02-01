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
	key := "index.html"
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
	CacheCurrentlyTrending, _ := service.GetTable(key, "CurrentlyTrending", []models.MacVod{})
	if CacheCurrentlyTrending == nil {
		DB.Debug().Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
		service.SaveTable(key, "CurrentlyTrending", CurrentlyTrending)
	} else {
		CurrentlyTrending = *CacheCurrentlyTrending.(*[]models.MacVod)
	}

	// 今日更新
	DB.Debug().Where("vod_time >= ?", midnightUnix).Count(&TodayCount)
	IndexData["TodayCount"] = TodayCount

	// 首页电影
	CachelistDianYing, _ := service.GetTable(key, "listDianYing", []models.MacVod{})
	if CachelistDianYing == nil {
		DB.Debug().Where("type_id_1", 1).Order("vod_hits desc").Limit(14).Find(&listDianYing)
		service.SaveTable(key, "listDianYing", listDianYing)
	} else {
		listDianYing = *CachelistDianYing.(*[]models.MacVod)
	}

	// 电影月榜
	CacheMonthDianYing, _ := service.GetTable(key, "MonthDianYing", []models.MacVod{})
	if CacheMonthDianYing == nil {
		DB.Debug().Where("type_id_1", 1).Order("vod_hits_month desc").Limit(10).Find(&MonthDianYing)
		service.SaveTable(key, "MonthDianYing", MonthDianYing)
	} else {
		MonthDianYing = *CacheMonthDianYing.(*[]models.MacVod)
	}

	// 首页电视剧
	CachelistDianShiJu, _ := service.GetTable(key, "listDianShiJu", []models.MacVod{})
	if CachelistDianShiJu == nil {
		DB.Debug().Where("type_id_1", 2).Order("vod_hits desc").Limit(14).Find(&listDianShiJu)
		service.SaveTable(key, "listDianShiJu", listDianShiJu)
	} else {
		listDianShiJu = *CachelistDianShiJu.(*[]models.MacVod)
	}

	// 电视剧月榜
	CacheMonthDianShiJu, _ := service.GetTable(key, "MonthDianShiJu", []models.MacVod{})
	if CacheMonthDianShiJu == nil {
		DB.Debug().Where("type_id_1", 2).Order("vod_hits_month desc").Limit(10).Find(&MonthDianShiJu)
		service.SaveTable(key, "MonthDianShiJu", MonthDianShiJu)
	} else {
		MonthDianShiJu = *CacheMonthDianShiJu.(*[]models.MacVod)
	}

	// 首页动漫
	CachelistDongManu, _ := service.GetTable(key, "listDongMan", []models.MacVod{})
	if CachelistDongManu == nil {
		DB.Debug().Where("type_id", 4).Order("vod_hits desc").Limit(14).Find(&listDongMan)
		service.SaveTable(key, "listDongMan", listDongMan)
	} else {
		listDongMan = *CachelistDongManu.(*[]models.MacVod)
	}

	// 动漫月榜
	CacheMonthDongMan, _ := service.GetTable(key, "MonthDongMan", []models.MacVod{})
	if CacheMonthDongMan == nil {
		DB.Debug().Where("type_id", 4).Order("vod_hits_month desc").Limit(10).Find(&MonthDongMan)
		service.SaveTable(key, "MonthDongMan", MonthDongMan)
	} else {
		MonthDongMan = *CacheMonthDongMan.(*[]models.MacVod)
	}

	// 首页综艺
	CachelistZongYi, _ := service.GetTable(key, "listZongYi", []models.MacVod{})
	if CachelistZongYi == nil {
		DB.Debug().Where("type_id", 3).Order("vod_hits desc").Limit(14).Find(&listZongYi)
		service.SaveTable(key, "listZongYi", listZongYi)
	} else {
		listZongYi = *CachelistZongYi.(*[]models.MacVod)
	}

	// 综艺月榜
	CacheMonthZongYi, _ := service.GetTable(key, "MonthZongYi", []models.MacVod{})
	if CacheMonthZongYi == nil {
		DB.Debug().Where("type_id", 3).Order("vod_hits_month desc").Limit(10).Find(&MonthZongYi)
		service.SaveTable(key, "MonthZongYi", MonthZongYi)
	} else {
		MonthZongYi = *CacheMonthZongYi.(*[]models.MacVod)
	}

	// 最新影片
	DB.Debug().Order("vod_time desc").Limit(8).Find(&NewsAll)
	DB.Debug().Where("type_id_1", 1).Order("vod_time desc").Limit(8).Find(&NewsVideos)
	DB.Debug().Where("type_id_1", 2).Order("vod_time desc").Limit(8).Find(&NewsDianShiJu)
	DB.Debug().Where("type_id", 4).Order("vod_time desc").Limit(8).Find(&NewsDongMan)
	DB.Debug().Where("type_id", 3).Order("vod_time desc").Limit(8).Find(&NewsZongYi)

	IndexData["CurrentlyTrending"] = CurrentlyTrending
	IndexData["listDianYing"] = listDianYing
	IndexData["MonthDianYing"] = MonthDianYing
	IndexData["listDianShiJu"] = listDianShiJu
	IndexData["MonthDianShiJu"] = MonthDianShiJu
	IndexData["listDongMan"] = listDongMan
	IndexData["MonthDongMan"] = MonthDongMan
	IndexData["listZongYi"] = listZongYi
	IndexData["MonthZongYi"] = MonthZongYi
	IndexData["listNewVideos"] = map[string][]models.MacVod{
		"NewsAll":       NewsAll,
		"NewsVideos":    NewsVideos,
		"NewsDianShiJu": NewsDianShiJu,
		"NewsDongMan":   NewsDongMan,
		"NewsZongYi":    NewsZongYi,
	}
	DATA["INDEX_DATA"] = IndexData
	c.HTML(http.StatusOK, "index.html", DATA)
}

func (h *IndexController) Web(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
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
	data["title"] = "Search"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "search.html", data)
}
