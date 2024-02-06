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
	// 二级详细分类
	service.ListMacType(table, 1, &listMacType)
	// 正在热播
	service.CurrentlyTrending(table, 1, 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		// 动作片
		CacheList, _ := service.GetTable(table, Name, []models.MacVod{})
		if CacheList == nil {
			MacVod.Debug().Where("type_id", TypeID).Order("vod_hits desc").Limit(16).Find(&BindList)
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
	// 二级详细分类
	service.ListMacType(table, 2, &listMacType)
	// 正在热播
	service.CurrentlyTrending(table, 2, 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		// 动作片
		CacheList, _ := service.GetTable(table, Name, []models.MacVod{})
		if CacheList == nil {
			MacVod.Debug().Where("type_id", TypeID).Order("vod_hits desc").Limit(16).Find(&BindList)
			service.SaveTable(table, Name, BindList)
		} else {
			BindList = *CacheList.(*[]models.MacVod)
		}
		PageData[Name] = BindList
	}
	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending

	DATA["PageData"] = PageData
	DATA["page"] = "dianshiju"
	c.HTML(http.StatusOK, "v/dianshiju.html", DATA)
}

func (h *VideosController) Dongman(c *gin.Context) {
	table := "dongman.html"
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
	// 二级详细分类
	service.ListMacType(table, 4, &listMacType)
	// 正在热播
	service.CurrentlyTrending(table, 4, 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		// 动作片
		CacheList, _ := service.GetTable(table, Name, []models.MacVod{})
		if CacheList == nil {
			MacVod.Debug().Where("type_id", TypeID).Order("vod_hits desc").Limit(16).Find(&BindList)
			service.SaveTable(table, Name, BindList)
		} else {
			BindList = *CacheList.(*[]models.MacVod)
		}
		PageData[Name] = BindList
	}
	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending

	DATA["PageData"] = PageData
	DATA["page"] = "dongman"
	c.HTML(http.StatusOK, "v/dongman.html", DATA)
}

func (h *VideosController) Zongyi(c *gin.Context) {
	table := "zongyi.html"
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
	// 二级详细分类
	service.ListMacType(table, 3, &listMacType)
	// 正在热播
	service.CurrentlyTrending(table, 3, 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		// 动作片
		CacheList, _ := service.GetTable(table, Name, []models.MacVod{})
		if CacheList == nil {
			MacVod.Debug().Where("type_id", TypeID).Order("vod_hits desc").Limit(16).Find(&BindList)
			service.SaveTable(table, Name, BindList)
		} else {
			BindList = *CacheList.(*[]models.MacVod)
		}
		PageData[Name] = BindList
	}
	PageData["listMacType"] = listMacType
	PageData["CurrentlyTrending"] = CurrentlyTrending

	DATA["PageData"] = PageData
	DATA["page"] = "zongyi"
	c.HTML(http.StatusOK, "v/zongyi.html", DATA)
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
