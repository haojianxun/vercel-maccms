package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"net/http"
	"time"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	key := "index.html"
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
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
	CurrentlyTrendingStr, err := redis.Client.Get(fmt.Sprintf("%s:%s", key, "CurrentlyTrending"))
	if err != nil && err.Error() != "redis: nil" {
		logger.Error(err)
		return
	}
	if len(CurrentlyTrendingStr) > 0 {
		err = json.Unmarshal([]byte(CurrentlyTrendingStr), &CurrentlyTrending)
		if err != nil {
			return
		}
	} else {
		DB.Debug().Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
		marshalStr, err := json.Marshal(CurrentlyTrending)
		if err != nil {
			logger.Error(err)
			return
		}
		_, err = redis.Client.Add(fmt.Sprintf("%s:%s", key, "CurrentlyTrending"), string(marshalStr), 0)
		if err != nil {
			logger.Error(err)
			return
		}
	}
	data["CurrentlyTrending"] = CurrentlyTrending
	// 今日更新
	DB.Debug().Where("vod_time >= ?", midnightUnix).Count(&TodayCount)
	data["TodayCount"] = TodayCount

	// 首页电影

	listDianYingStr, err := redis.Client.Get(fmt.Sprintf("%s:%s", key, "listDianYing"))
	if err != nil && err.Error() != "redis: nil" {
		logger.Error(err)
		return
	}
	if len(listDianYingStr) > 0 {
		err = json.Unmarshal([]byte(listDianYingStr), &listDianYing)
		if err != nil {
			return
		}
	} else {
		DB.Debug().Where("type_id_1", 1).Order("vod_hits desc").Limit(14).Find(&listDianYing)
		marshalStr, err := json.Marshal(listDianYing)
		if err != nil {
			logger.Error(err)
			return
		}
		_, err = redis.Client.Add(fmt.Sprintf("%s:%s", key, "listDianYing"), string(marshalStr), 0)
		if err != nil {
			logger.Error(err)
			return
		}
	}
	data["listDianYing"] = listDianYing

	// 电影月榜
	DB.Debug().Where("type_id_1", 1).Order("vod_hits_month desc").Limit(10).Find(&MonthDianYing)
	data["MonthDianYing"] = MonthDianYing

	// 首页电视剧
	DB.Debug().Where("type_id_1", 2).Order("vod_hits desc").Limit(14).Find(&listDianShiJu)
	data["listDianShiJu"] = listDianShiJu

	// 电视剧月榜
	DB.Debug().Where("type_id_1", 2).Order("vod_hits_month desc").Limit(10).Find(&MonthDianShiJu)
	data["MonthDianShiJu"] = MonthDianShiJu

	// 首页动漫
	DB.Debug().Where("type_id", 4).Order("vod_hits desc").Limit(14).Find(&listDongMan)
	data["listDongMan"] = listDongMan

	// 动漫月榜
	DB.Debug().Where("type_id", 4).Order("vod_hits_month desc").Limit(10).Find(&MonthDongMan)
	data["MonthDongMan"] = MonthDongMan

	// 首页综艺
	DB.Debug().Where("type_id", 3).Order("vod_hits desc").Limit(14).Find(&listZongYi)
	data["listZongYi"] = listZongYi

	// 综艺月榜
	DB.Debug().Where("type_id", 3).Order("vod_hits_month desc").Limit(10).Find(&MonthZongYi)
	data["MonthZongYi"] = MonthZongYi

	// 最新影片
	DB.Debug().Order("vod_time desc").Limit(8).Find(&NewsAll)
	DB.Debug().Where("type_id_1", 1).Order("vod_time desc").Limit(8).Find(&NewsVideos)
	DB.Debug().Where("type_id_1", 2).Order("vod_time desc").Limit(8).Find(&NewsDianShiJu)
	DB.Debug().Where("type_id", 4).Order("vod_time desc").Limit(8).Find(&NewsDongMan)
	DB.Debug().Where("type_id", 3).Order("vod_time desc").Limit(8).Find(&NewsZongYi)

	data["listNewVideos"] = map[string][]models.MacVod{
		"NewsAll":       NewsAll,
		"NewsVideos":    NewsVideos,
		"NewsDianShiJu": NewsDianShiJu,
		"NewsDongMan":   NewsDongMan,
		"NewsZongYi":    NewsZongYi,
	}

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
