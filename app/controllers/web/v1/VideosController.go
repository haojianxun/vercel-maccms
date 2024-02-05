package v1

import (
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service"
	"goapi/pkg/mysql"
	"net/http"
)

type VideosController struct {
	BaseController
}

func (h *VideosController) Dianying(c *gin.Context) {
	table := "dianying.html"
	PageData := cmap.New().Items()
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	DATA := value.(gin.H)

	// 批量查询数据
	var (
		listMacType []models.MacType
		CurrentlyTrending, listDongZuoPian,
		listXiJuPian, listAiQingPian,
		listKeHuanPian, listKongBuPian,
		listJuQingPian, listZhanZhengPian,
		listDongHuaPian, listQiHuanPian,
		listJiLuPian []models.MacVod
	)
	MacVod := models.MacVodMgr(mysql.DB)
	MacType := models.MacTypeMgr(mysql.DB)

	// 电影详细分类
	CacheMacType, _ := service.GetTable(table, "listMacType", []models.MacType{})
	if CacheMacType == nil {
		MacType.Debug().Where("type_pid", 1).Where("type_status", 1).Order("type_sort desc").Find(&listMacType)
		service.SaveTable(table, "listMacType", listMacType)
	} else {
		listMacType = *CacheMacType.(*[]models.MacType)
	}

	// 正在热播
	CacheCurrentlyTrending, _ := service.GetTable(table, "CurrentlyTrending", []models.MacVod{})
	if CacheCurrentlyTrending == nil {
		MacVod.Debug().Where("type_id_1", 1).Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
		service.SaveTable(table, "CurrentlyTrending", CurrentlyTrending)
	} else {
		CurrentlyTrending = *CacheCurrentlyTrending.(*[]models.MacVod)
	}

	// 动作片
	CacheDongZuoPian, _ := service.GetTable(table, "listDongZuoPian", []models.MacVod{})
	if CacheDongZuoPian == nil {
		MacVod.Debug().Where("type_id", 6).Order("vod_hits desc").Limit(14).Find(&listDongZuoPian)
		service.SaveTable(table, "listDongZuoPian", listDongZuoPian)
	} else {
		listDongZuoPian = *CacheDongZuoPian.(*[]models.MacVod)
	}

	// 喜剧片
	CacheXiJu, _ := service.GetTable(table, "listXiJuPian", []models.MacVod{})
	if CacheXiJu == nil {
		MacVod.Debug().Where("type_id", 7).Order("vod_hits desc").Limit(14).Find(&listXiJuPian)
		service.SaveTable(table, "listXiJuPian", listXiJuPian)
	} else {
		listXiJuPian = *CacheXiJu.(*[]models.MacVod)
	}

	// 爱情片
	CacheAiQingPian, _ := service.GetTable(table, "listAiQingPian", []models.MacVod{})
	if CacheAiQingPian == nil {
		MacVod.Debug().Where("type_id", 8).Order("vod_hits desc").Limit(14).Find(&listAiQingPian)
		service.SaveTable(table, "listAiQingPian", listAiQingPian)
	} else {
		listAiQingPian = *CacheAiQingPian.(*[]models.MacVod)
	}

	// 科幻片
	CacheKeHuanPian, _ := service.GetTable(table, "listKeHuanPian", []models.MacVod{})
	if CacheKeHuanPian == nil {
		MacVod.Debug().Where("type_id", 9).Order("vod_hits desc").Limit(14).Find(&listKeHuanPian)
		service.SaveTable(table, "listKeHuanPian", listKeHuanPian)
	} else {
		listKeHuanPian = *CacheKeHuanPian.(*[]models.MacVod)
	}

	// 恐怖片
	CacheKongBuPian, _ := service.GetTable(table, "listKongBuPian", []models.MacVod{})
	if CacheKongBuPian == nil {
		MacVod.Debug().Where("type_id", 10).Order("vod_hits desc").Limit(14).Find(&listKongBuPian)
		service.SaveTable(table, "listKongBuPian", listKongBuPian)
	} else {
		listKongBuPian = *CacheKongBuPian.(*[]models.MacVod)
	}

	// 剧情片
	CacheJuQingPian, _ := service.GetTable(table, "listJuQingPian", []models.MacVod{})
	if CacheJuQingPian == nil {
		MacVod.Debug().Where("type_id", 11).Order("vod_hits desc").Limit(14).Find(&listJuQingPian)
		service.SaveTable(table, "listJuQingPian", listJuQingPian)
	} else {
		listJuQingPian = *CacheJuQingPian.(*[]models.MacVod)
	}

	// 战争片
	CacheZhanZhengPian, _ := service.GetTable(table, "listZhanZhengPian", []models.MacVod{})
	if CacheZhanZhengPian == nil {
		MacVod.Debug().Where("type_id", 12).Order("vod_hits desc").Limit(14).Find(&listZhanZhengPian)
		service.SaveTable(table, "listZhanZhengPian", listZhanZhengPian)
	} else {
		listZhanZhengPian = *CacheZhanZhengPian.(*[]models.MacVod)
	}

	// 动画片
	CacheDongHuaPian, _ := service.GetTable(table, "listDongHuaPian", []models.MacVod{})
	if CacheDongHuaPian == nil {
		MacVod.Debug().Where("type_id", 20).Order("vod_hits desc").Limit(14).Find(&listDongHuaPian)
		service.SaveTable(table, "listDongHuaPian", listDongHuaPian)
	} else {
		listDongHuaPian = *CacheDongHuaPian.(*[]models.MacVod)
	}

	// 奇幻片
	CacheQiHuanPian, _ := service.GetTable(table, "listQiHuanPian", []models.MacVod{})
	if CacheQiHuanPian == nil {
		MacVod.Debug().Where("type_id", 21).Order("vod_hits desc").Limit(14).Find(&listQiHuanPian)
		service.SaveTable(table, "listQiHuanPian", listQiHuanPian)
	} else {
		listQiHuanPian = *CacheQiHuanPian.(*[]models.MacVod)
	}

	// 记录片
	CacheJiLuPian, _ := service.GetTable(table, "listJiLuPian", []models.MacVod{})
	if CacheJiLuPian == nil {
		MacVod.Debug().Where("type_id", 22).Order("vod_hits desc").Limit(14).Find(&listJiLuPian)
		service.SaveTable(table, "listJiLuPian", listJiLuPian)
	} else {
		listJiLuPian = *CacheJiLuPian.(*[]models.MacVod)
	}

	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending
	PageData["listDongZuoPian"] = listDongZuoPian
	PageData["listXiJuPian"] = listXiJuPian
	PageData["listAiQingPian"] = listAiQingPian
	PageData["listKeHuanPian"] = listKeHuanPian
	PageData["listKongBuPian"] = listKongBuPian
	PageData["listJuQingPian"] = listJuQingPian
	PageData["listZhanZhengPian"] = listZhanZhengPian
	PageData["listDongHuaPian"] = listDongHuaPian
	PageData["listQiHuanPian"] = listQiHuanPian
	PageData["listJiLuPian"] = listJiLuPian
	DATA["PageData"] = PageData
	DATA["page"] = "dianying"
	c.HTML(http.StatusOK, "v/dianying.html", DATA)
}

func (h *VideosController) Dianshiju(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "dianshiju"
	data["title"] = "电视剧"
	c.HTML(http.StatusOK, "v/dianshiju.html", data)
}

func (h *VideosController) Dongman(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "dongman"
	data["title"] = "动漫"
	c.HTML(http.StatusOK, "v/dongman.html", data)
}

func (h *VideosController) Zongyi(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "zongyi"
	data["title"] = "综艺"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "v/zongyi.html", data)
}

func (h *VideosController) Play(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "play"
	data["title"] = "Play"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "play.html", data)
}

func (h *VideosController) Show(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "show"
	data["title"] = "show"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "show.html", data)
}

func (h *VideosController) PianKu(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "pianku"
	data["title"] = "PianKu"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "pianku.html", data)
}

func (h *VideosController) Top(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "top"
	data["title"] = "Top"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "top.html", data)
}
