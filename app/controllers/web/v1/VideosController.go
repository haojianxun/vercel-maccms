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
		listMacType       []models.MacType
		CurrentlyTrending []models.MacVod
	)
	MacVod := models.MacVodMgr(mysql.DB)
	MacType := models.MacTypeMgr(mysql.DB)

	// 二级详细分类
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

	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		// 动作片
		CacheList, _ := service.GetTable(table, Name, []models.MacVod{})
		if CacheList == nil {
			MacVod.Debug().Where("type_id", TypeID).Order("vod_hits desc").Limit(14).Find(&BindList)
			service.SaveTable(table, Name, BindList)
		} else {
			BindList = *CacheList.(*[]models.MacVod)
		}
		PageData[Name] = BindList
	}
	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending
	DATA["PageData"] = PageData
	DATA["page"] = "dianying"
	c.HTML(http.StatusOK, "v/dianying.html", DATA)
}

func (h *VideosController) Dianshiju(c *gin.Context) {
	table := "dianshiju.html"
	PageData := cmap.New().Items()
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	DATA := value.(gin.H)
	// 批量查询数据
	var (
		listMacType       []models.MacType
		CurrentlyTrending []models.MacVod
	)
	MacVod := models.MacVodMgr(mysql.DB)
	MacType := models.MacTypeMgr(mysql.DB)
	// 二级详细分类
	CacheMacType, _ := service.GetTable(table, "listMacType", []models.MacType{})
	if CacheMacType == nil {
		MacType.Debug().Where("type_pid", 2).Where("type_status", 1).Order("type_sort asc").Find(&listMacType)
		service.SaveTable(table, "listMacType", listMacType)
	} else {
		listMacType = *CacheMacType.(*[]models.MacType)
	}

	// 正在热播
	CacheCurrentlyTrending, _ := service.GetTable(table, "CurrentlyTrending", []models.MacVod{})
	if CacheCurrentlyTrending == nil {
		MacVod.Debug().Where("type_id_1", 2).Where("vod_status", 1).Order("vod_hits desc").Limit(16).Find(&CurrentlyTrending)
		service.SaveTable(table, "CurrentlyTrending", CurrentlyTrending)
	} else {
		CurrentlyTrending = *CacheCurrentlyTrending.(*[]models.MacVod)
	}

	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending

	DATA["PageData"] = PageData
	DATA["page"] = "dianshiju"
	c.HTML(http.StatusOK, "v/dianshiju.html", DATA)
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
