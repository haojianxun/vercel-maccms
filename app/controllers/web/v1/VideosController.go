package v1

import (
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service"
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
	// 二级详细分类
	service.ListMacType(table, 1, &listMacType)
	// 正在热播
	service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
		"type_id_1":  1,
		"vod_status": 1,
	}, "vod_hits desc", 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
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
	// 二级详细分类
	service.ListMacType(table, 2, &listMacType)
	// 正在热播
	service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
		"type_id_1":  2,
		"vod_status": 1,
	}, "vod_hits desc", 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
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
	// 二级详细分类
	service.ListMacType(table, 4, &listMacType)
	// 正在热播
	service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
		"type_id_1":  4,
		"vod_status": 1,
	}, "vod_hits desc", 16, &CurrentlyTrending)

	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
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
	// 二级详细分类
	service.ListMacType(table, 3, &listMacType)
	// 正在热播
	service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
		"type_id_1":  3,
		"vod_status": 1,
	}, "vod_hits desc", 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
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
