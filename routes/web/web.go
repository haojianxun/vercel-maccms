package web

import (
	"github.com/gin-gonic/gin"
	v1 "goapi/app/controllers/web/v1"
	middlewaresV1 "goapi/app/middlewares/v1"
)

// RegisterWebRoutes 注册 Web 路由
func RegisterWebRoutes(router *gin.RouterGroup) {
	router.Use(middlewaresV1.Web())
	GroupV1 := new(v1.GroupV1)
	webV1 := router.Group("/")
	{
		// 首页入口
		webV1.GET("/", GroupV1.IndexController.Index)
		webV1.GET("/index.html", GroupV1.IndexController.Index)
		// 片库
		webV1.GET("/pianku.html", GroupV1.VideosController.PianKu)
		// 播放
		webV1.GET("/play.html", GroupV1.VideosController.Play)
		// 搜索
		webV1.GET("/search.html", GroupV1.IndexController.Search)
		// 详情
		webV1.GET("/show-:params", GroupV1.VideosController.Show)
		// 排行榜
		webV1.GET("/top.html", GroupV1.VideosController.Top)
		// Web
		webV1.GET("/web.html", GroupV1.IndexController.Web)
		// APP
		webV1.GET("/label/app.html", GroupV1.IndexController.App)
		// 站点路由
		v := webV1.Group("/v")
		{
			// 电影
			v.GET("/dianying.html", GroupV1.VideosController.Dianying)
			// 电视剧
			v.GET("/dianshiju.html", GroupV1.VideosController.Dianshiju)
			// 动漫
			v.GET("/dongman.html", GroupV1.VideosController.Dongman)
			// 综艺
			v.GET("/zongyi.html", GroupV1.VideosController.Zongyi)
		}

		arttype := webV1.Group("/arttype")
		{
			// 资讯
			arttype.GET("/zixun.html", GroupV1.ArttypeController.Zixun)
		}
	}
}
