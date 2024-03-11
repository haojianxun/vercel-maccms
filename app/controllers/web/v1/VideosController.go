package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/service"
	"goapi/pkg/helpers"
	"goapi/pkg/maccms"
	"goapi/pkg/mysql"
	"goapi/pkg/page"
	"net/http"
	"strings"
)

type VideosController struct {
	BaseController
}

func (h *VideosController) PianKu(c *gin.Context) {
	// 获取参数
	strParams := strings.ReplaceAll(c.Param("params"), ".html", "")
	args := strings.Split(strParams, "-")
	PageName := args[0]
	if len(PageName) <= 0 {
		NoPage(c)
		return
	}
	DATA := GetDATA(c)
	PageData := cmap.New().Items()
	NavMenus := DATA["NavMenus"]
	fmt.Println(NavMenus)
	fmt.Println(NavMenus)

	// 默认为电影片库
	// 电视剧片库
	// 动漫片库
	// 综艺片库
	DATA["page"] = PageName
	DATA["title"] = "PianKu"
	DATA["PageData"] = PageData
	PageMs(c)
	c.HTML(http.StatusOK, "pianku.html", DATA)
}

func (h *VideosController) Category(c *gin.Context) {
	// 获取参数
	strParams := strings.ReplaceAll(c.Param("params"), ".html", "")
	args := strings.Split(strParams, "-")
	typeEn := args[0]
	if len(typeEn) <= 0 {
		NoPage(c)
		return
	}

	// 初始化数据
	DATA := GetDATA(c)
	PageData := cmap.New().Items()

	// 查询分类数据
	macType := models.MacType{}
	where := map[string]interface{}{
		"type_status": 1,
		"type_en":     typeEn,
	}
	if err := models.MacTypeMgr(mysql.DB).Where(where).First(&macType).Error; err != nil {
		NoPage(c)
		return
	}

	// 设置页面数据
	DATA["page"] = typeEn
	DATA["macType"] = macType

	// 处理顶级栏目下的数据
	if macType.TypePid == 0 {
		var listMacType []models.MacType
		service.ListMacType(typeEn, int(macType.TypeID), &listMacType)
		PageData["listMacType"] = listMacType

		// 获取正在热播的视频
		var CurrentlyTrending []models.MacVod
		service.ListWhereMacVod(typeEn, "CurrentlyTrending", map[string]interface{}{
			"type_id_1":  macType.TypeID,
			"vod_status": 1,
		}, "vod_hits desc", 16, &CurrentlyTrending)
		PageData["CurrentlyTrending"] = CurrentlyTrending

		// 查询每个子分类下的数据
		for _, item := range listMacType {
			var BindList []models.MacVod
			service.ListWhereMacVod(typeEn, item.TypeEn, map[string]interface{}{
				"type_id":    item.TypeID,
				"vod_status": 1,
			}, "vod_hits desc", 16, &BindList)
			PageData[item.TypeEn] = BindList
		}

		DATA["CategoryType"] = "top"
		DATA["NavName"] = macType.TypeEn

	} else { // 处理子栏目下的数据
		var listMacType []models.MacType
		service.ListMacType(typeEn, int(macType.TypePid), &listMacType)
		PageData["listMacType"] = listMacType

		// 获取子分类下的视频列表
		var pageList page.PageList
		var params requests.Search
		if len(args) >= 6 && len(args[6]) > 0 {
			params.PageNum = helpers.StringToInt64(args[6])
		}
		params.PageSize = 72
		whereSub := cmap.New().Items()
		whereSub["type_id"] = macType.TypeID
		whereSub["vod_status"] = 1

		models.MacVodMgr(mysql.DB).Where(whereSub).Count(&pageList.Total)
		pageList.CurrentPage = params.PageNum
		pageList.PageSize = params.PageSize
		page.InitPageList(&pageList)

		var listResult []models.MacVod
		models.MacVodMgr(mysql.DB).
			Where(whereSub).
			Offset(int(pageList.Offset)).
			Limit(int(pageList.PageSize)).
			Find(&listResult)

		pageList.List = listResult
		PageData["pageList"] = pageList
		PageData["PaginationHTML"] = PaginationHTML(int(pageList.CurrentPage), int(pageList.PageTotal), fmt.Sprintf("%s%s", "v/", macType.TypeEn),
			map[string]string{
				"movie":    "",
				"actor":    "",
				"area":     "",
				"director": "",
				"year":     "",
			},
		)

		DATA["CategoryType"] = "sub"
		DATA["NavName"] = maccms.TypeEn(DATA["NavMenus"], macType.TypePid)
	}

	// 设置页面数据并返回
	DATA["PageData"] = PageData
	c.HTML(http.StatusOK, "v/category.html", DATA)
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
	DATA["page"] = "show"
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
	DATA["page"] = "play"
	c.HTML(http.StatusOK, "play.html", DATA)
}

func (h *VideosController) Top(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "top"
	DATA["title"] = "Top"
	PageMs(c)
	c.HTML(http.StatusOK, "top.html", DATA)
}
