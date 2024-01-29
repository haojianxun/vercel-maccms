package v1

import (
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/mysql"
	"net/http"
	"time"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "index"
	// 获取当前时间
	now := time.Now()
	// 获取当天凌晨00:00的时间
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 获取当天凌晨00:00的时间戳
	midnightUnix := midnight.Unix()
	DB := models.MacVodMgr(mysql.DB)
	// 正在热播
	var CurrentlyTrending []models.MacVod
	DB.Debug().Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
	data["CurrentlyTrending"] = CurrentlyTrending
	// 今日更新
	var TodayCount int64
	DB.Debug().Where("vod_time >= ?", midnightUnix).Count(&TodayCount)
	data["TodayCount"] = TodayCount

	// 首页电影
	var listDianYing []models.MacVod
	DB.Debug().
		Where("type_id_1", 1).
		Order("vod_hits desc").
		Limit(14).
		Find(&listDianYing)
	data["listDianYing"] = listDianYing
	// 电影月榜
	var MonthDianYing []models.MacVod
	DB.Debug().
		Where("type_id_1", 1).
		Order("vod_hits_month desc").
		Limit(10).
		Find(&MonthDianYing)
	data["MonthDianYing"] = MonthDianYing

	// 首页电视剧
	var listDianShiJu []models.MacVod
	DB.Debug().
		Where("type_id_1", 2).
		Order("vod_hits desc").
		Limit(14).
		Find(&listDianShiJu)
	data["listDianShiJu"] = listDianShiJu
	// 电视剧月榜
	var MonthDianShiJu []models.MacVod
	DB.Debug().
		Where("type_id_1", 2).
		Order("vod_hits_month desc").
		Limit(10).
		Find(&MonthDianShiJu)
	data["MonthDianShiJu"] = MonthDianShiJu

	// 首页动漫
	var listDongMan []models.MacVod
	DB.Debug().
		Where("type_id", 4).
		Order("vod_hits desc").
		Limit(14).
		Find(&listDongMan)
	data["listDongMan"] = listDongMan
	// 动漫月榜
	var MonthDongMan []models.MacVod
	DB.Debug().
		Where("type_id", 4).
		Order("vod_hits_month desc").
		Limit(10).
		Find(&MonthDongMan)
	data["MonthDongMan"] = MonthDongMan

	// 首页综艺
	var listZongYi []models.MacVod
	DB.Debug().
		Where("type_id", 3).
		Order("vod_hits desc").
		Limit(14).
		Find(&listZongYi)
	data["listZongYi"] = listZongYi
	// 动漫月榜
	var MonthZongYi []models.MacVod
	DB.Debug().
		Where("type_id", 3).
		Order("vod_hits_month desc").
		Limit(10).
		Find(&MonthZongYi)
	data["MonthZongYi"] = MonthZongYi

	// 最新影片
	listNewVideos := cmap.New().Items()
	var NewsAll []models.MacVod
	DB.Debug().
		Order("vod_time desc").
		Limit(8).
		Find(&NewsAll)
	listNewVideos["NewsAll"] = NewsAll
	var NewsVideos []models.MacVod
	DB.Debug().
		Where("type_id_1", 1).
		Order("vod_time desc").
		Limit(8).
		Find(&NewsVideos)
	listNewVideos["NewsVideos"] = NewsVideos
	var NewsDianShiJu []models.MacVod
	DB.Debug().
		Where("type_id_1", 2).
		Order("vod_time desc").
		Limit(8).
		Find(&NewsDianShiJu)
	listNewVideos["NewsDianShiJu"] = NewsDianShiJu
	var NewsDongMan []models.MacVod
	DB.Debug().
		Where("type_id", 4).
		Order("vod_time desc").
		Limit(8).
		Find(&NewsDongMan)
	listNewVideos["NewsDongMan"] = NewsDongMan
	var NewsZongYi []models.MacVod
	DB.Debug().
		Where("type_id", 3).
		Order("vod_time desc").
		Limit(8).
		Find(&NewsZongYi)
	listNewVideos["NewsZongYi"] = NewsZongYi
	data["listNewVideos"] = listNewVideos
	c.HTML(http.StatusOK, "index.html", data)
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
