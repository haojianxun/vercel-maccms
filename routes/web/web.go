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
		webV1.GET("/", GroupV1.IndexController.Index)
		webV1.GET("/home", GroupV1.IndexController.Home)
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
			// Play
			v.GET("/play.html", GroupV1.VideosController.Play)
		}

		arttype := webV1.Group("/arttype")
		{
			// 资讯
			arttype.GET("/zixun.html", GroupV1.ArttypeController.Zixun)
		}
	}
}
