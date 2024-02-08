package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/service"
	"goapi/pkg/helpers"
	"goapi/pkg/maccms"
	"goapi/pkg/mysql"
	"net/http"
	"strings"
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
	}, "vod_year desc,vod_hits desc", 16, &CurrentlyTrending)
	// 根据分类遍历查询每个子类的下的数据，一般获取14条按照热度倒序排序
	for _, item := range listMacType {
		Name := item.TypeEn
		TypeID := item.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_year desc,vod_hits desc", 16, &BindList)
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

func (h *VideosController) Show(c *gin.Context) {
	table := "show.html"
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
	c.HTML(http.StatusOK, "show.html", DATA)
}

func (h *VideosController) Play(c *gin.Context) {
	table := "play.html"
	PageData := cmap.New().Items()
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	DATA := value.(gin.H)
	// 获取路由中的参数值
	params := strings.ReplaceAll(c.Param("params"), ".html", "")
	if len(params) < 10 {
		c.HTML(http.StatusOK, "404", nil)
		return
	}
	// 参数不正确
	if !strings.Contains(params, "-") {
		c.HTML(http.StatusOK, "404", nil)
		return
	}
	// TOCzRb-1-1
	array := strings.Split(params, "-")
	VodID := maccms.DecryptID(array[0]) // 视频ID
	Node := array[1]                    // 播放节点/播放线路
	Part := array[2]                    // 第几集/第几部分
	fmt.Println(Node)
	fmt.Println(Part)
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
	PageData["VodID"] = VodID
	PageData["Node"] = helpers.StringToInt(Node) - 1
	PageData["Part"] = helpers.StringToInt(Part) - 1
	DATA["PageData"] = PageData
	DATA["page"] = MacTypeDetail.TypeEn
	c.HTML(http.StatusOK, "play.html", DATA)
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
