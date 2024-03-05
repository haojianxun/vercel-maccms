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

func (h *VideosController) Category(c *gin.Context) {
	PageData := cmap.New().Items()
	DATA := GetDATA(c)
	// 批量查询数据
	var (
		macType           models.MacType
		listMacType       []models.MacType
		CurrentlyTrending []models.MacVod
	)
	typeEn := strings.ReplaceAll(c.Param("params"), ".html", "")
	if len(typeEn) <= 0 {
		NoPage(c)
		return
	}
	table := typeEn
	where := map[string]interface{}{}
	where["type_status"] = 1
	where["type_en"] = typeEn
	models.MacTypeMgr(mysql.DB).Debug().Where(where).Order("type_sort asc").Find(&macType)
	if macType.TypeID == 0 {
		NoPage(c)
		return
	}
	// 二级栏目详细分类获取
	if macType.TypePid == 0 {
		service.ListMacType(table, int(macType.TypeID), &listMacType)
		// 正在热播
		service.ListWhereMacVod(table, "CurrentlyTrending", map[string]interface{}{
			"type_id_1":  macType.TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &CurrentlyTrending)
		PageData["CurrentlyTrending"] = CurrentlyTrending
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
	} else {
		service.ListMacType(table, int(macType.TypePid), &listMacType)
		Name := macType.TypeEn
		TypeID := macType.TypeID
		var BindList []models.MacVod
		service.ListWhereMacVod(table, Name, map[string]interface{}{
			"type_id":    TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &BindList)
		PageData[Name] = BindList
	}
	PageData["listMacType"] = listMacType
	DATA["PageData"] = PageData
	DATA["page"] = typeEn
	c.HTML(http.StatusOK, "v/dianying.html", DATA)
}

func (h *VideosController) Dianshiju(c *gin.Context) {
	table := "dianshiju.html"
	PageData := cmap.New().Items()
	DATA := GetDATA(c)
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
	DATA := GetDATA(c)
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
	DATA := GetDATA(c)
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
	DATA := GetDATA(c)

	// 获取路由中的参数值
	params := strings.ReplaceAll(c.Param("params"), ".html", "")
	if len(params) > 6 {
		NoPage(c)
		return
	}
	VodID := maccms.DecryptID(params)
	var detail models.MacVod
	err := models.MacVodMgr(mysql.DB).Where("vod_id", VodID).Find(&detail).Error
	if err != nil {
		NoPage(c)
		return
	}
	var (
		Related, CurrentlyTrending []models.MacVod
		MacTypeDetail              models.MacType
	)
	err = models.MacTypeMgr(mysql.DB).Where("type_id", detail.TypeID1).Find(&MacTypeDetail).Error
	if err != nil {
		NoPage(c)
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
	DATA := GetDATA(c)
	// 获取路由中的参数值
	params := strings.ReplaceAll(c.Param("params"), ".html", "")
	if len(params) < 10 {
		NoPage(c)
		return
	}
	// 参数不正确
	if !strings.Contains(params, "-") {
		NoPage(c)
		return
	}
	// TOCzRb-1-1
	array := strings.Split(params, "-")
	VodID := maccms.DecryptID(array[0])   // 视频ID
	Node := helpers.StringToInt(array[1]) // 播放节点/播放线路
	Part := helpers.StringToInt(array[2]) // 第几集/第几部分
	var detail models.MacVod
	err := models.MacVodMgr(mysql.DB).Where("vod_id", VodID).Find(&detail).Error
	if err != nil {
		NoPage(c)
		return
	}
	var (
		Related, CurrentlyTrending []models.MacVod
		MacTypeDetail              models.MacType
	)
	err = models.MacTypeMgr(mysql.DB).Where("type_id", detail.TypeID1).Find(&MacTypeDetail).Error
	if err != nil {
		NoPage(c)
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
	PageData["Node"] = Node
	PageData["Part"] = Part
	PageData["Link"] = fmt.Sprintf("%v-%v-%v.html", array[0], Node, Part)
	PageData["LinkNext"] = fmt.Sprintf("%v-%v-%v.html", array[0], Node, Part+1)
	PageData["LinkPre"] = fmt.Sprintf("%v-%v-%v.html", array[0], Node, Part-1)
	// 播放地址
	VodPlayURLGroup := strings.Split(detail.VodPlayURL, "$$$")
	PartGroup := strings.Split(VodPlayURLGroup[Node-1], "#")
	var CurrentUrlInfo []string
	if len(PartGroup) < Part-1 {
		// 最大不能超出数组
		CurrentUrlInfo = strings.Split(PartGroup[len(PartGroup)-1], "$")
	} else {
		CurrentUrlInfo = strings.Split(PartGroup[Part-1], "$")
	}
	PageData["PartName"] = CurrentUrlInfo[0]
	PageData["NodeName"] = CurrentUrlInfo[0]
	PageData["Url"] = maccms.Base64encode(maccms.EncodeURL(CurrentUrlInfo[1]))
	if len(PartGroup) > Part {
		NextUrlInfo := strings.Split(PartGroup[Part], "$")
		PageData["UrlNext"] = maccms.Base64encode(maccms.EncodeURL(NextUrlInfo[1]))
	}
	DATA["PageData"] = PageData
	DATA["page"] = MacTypeDetail.TypeEn
	c.HTML(http.StatusOK, "play.html", DATA)
}

func (h *VideosController) PianKu(c *gin.Context) {
	DATA := GetDATA(c)

	// 默认为电影片库
	// 电视剧片库
	// 动漫片库
	// 综艺片库
	DATA["page"] = "pianku"
	DATA["title"] = "PianKu"
	PageMs(c)
	c.HTML(http.StatusOK, "pianku.html", DATA)
}

func (h *VideosController) Top(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "top"
	DATA["title"] = "Top"
	PageMs(c)
	c.HTML(http.StatusOK, "top.html", DATA)
}
