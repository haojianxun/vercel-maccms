package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service"
	"goapi/pkg/maccms"
	"goapi/pkg/mysql"
	"net/http"
	"strings"
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
	startTime, _ := c.Get("startTime")
	fmt.Println("当前耗时", time.Now().UnixMilli()-startTime.(int64), "ms")
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
	table := "search.html"
	PageData := cmap.New().Items()
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	DATA := value.(gin.H)

	// 获取路由中的参数值
	params := strings.ReplaceAll(c.Param("params"), ".html", "")
	if len(params) > 6 {
		c.HTML(http.StatusOK, "404", nil)
		return
	}
	VodID := maccms.DecryptID(params)
	var detail models.MacVod
	err := models.MacVodMgr(mysql.DB).Where("vod_id", VodID).Find(&detail).Error
	if err != nil {
		c.HTML(http.StatusOK, "404", nil)
		return
	}
	var (
		Related, CurrentlyTrending []models.MacVod
		MacTypeDetail              models.MacType
	)
	err = models.MacTypeMgr(mysql.DB).Where("type_id", detail.TypeID1).Find(&MacTypeDetail).Error
	if err != nil {
		c.HTML(http.StatusOK, "404", nil)
		return
	}
	// 相关影片
	RelatedName := fmt.Sprintf("Related:%v", detail.TypeID)
	service.ListWhereMacVod(table, RelatedName, map[string]interface{}{
		"type_id":    detail.TypeID,
		"vod_status": 1,
	}, "vod_year desc,vod_hits desc", 16, &Related)

	// 正在热播
	ReBoName := fmt.Sprintf("CurrentlyTrending:%v", detail.TypeID)
	service.ListWhereMacVod(table, ReBoName, map[string]interface{}{
		"type_id":    detail.TypeID,
		"vod_status": 1,
	}, "vod_hits desc", 10, &CurrentlyTrending)

	PageData["detail"] = detail
	PageData["MacTypeDetail"] = MacTypeDetail
	PageData["Related"] = Related
	PageData["CurrentlyTrending"] = CurrentlyTrending
	DATA["VodID"] = VodID
	DATA["PageData"] = PageData
	DATA["page"] = MacTypeDetail.TypeEn
	c.HTML(http.StatusOK, "search.html", DATA)
}
