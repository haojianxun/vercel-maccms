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

func (h *IndexController) PresTrain(c *gin.Context) {
	value, _ := c.Get("data")
	DATA := value.(gin.H)
	c.HTML(http.StatusOK, "prestrain.html", DATA)
}

func (h *IndexController) DPlayer(c *gin.Context) {
	value, _ := c.Get("data")
	DATA := value.(gin.H)
	c.HTML(http.StatusOK, "dplayer.html", DATA)
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
		listMacType                                                                   []models.MacType
		CurrentlyTrending, listDianYing, MonthDianYing, listDianShiJu, MonthDianShiJu []models.MacVod
		listDongMan, MonthDongMan, listZongYi, MonthZongYi, NewsAll, NewsVideos       []models.MacVod
		NewsDianShiJu, NewsDongMan, NewsZongYi                                        []models.MacVod
	)
	DB := models.MacVodMgr(mysql.DB)
	// 查询所有分类
	service.ListMacType(table, -1, &listMacType)
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

func (h *IndexController) About(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "about"
	data["title"] = "About"
	data["list"] = gin.H{"a": "b"}
	c.HTML(http.StatusOK, "about.html", data)
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
