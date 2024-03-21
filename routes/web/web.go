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
		webV1.GET("/robots.txt", GroupV1.IndexController.Robots)
		webV1.GET("/baidu_verify_codeva-irCSP8JvXx.html", GroupV1.IndexController.Baidu)
		webV1.GET("/sogousiteverification.txt", GroupV1.IndexController.SouGou)
		// 首页入口
		webV1.GET("/", GroupV1.IndexController.Index)
		webV1.GET("/google8fcf6226304ee6af.html", GroupV1.IndexController.Google)
		webV1.GET("/index.html", GroupV1.IndexController.Index)
		webV1.GET("/prestrain.html", GroupV1.IndexController.PresTrain)
		webV1.GET("/player/dplayer.html", GroupV1.IndexController.DPlayer)
		// 播放
		webV1.GET("/play-:params", GroupV1.VideosController.Play)
		// 搜索
		{
			webV1.GET("/search:search", GroupV1.IndexController.Search)
		}
		// 详情
		webV1.GET("/show-:params", GroupV1.VideosController.Show)
		// 排行榜
		webV1.GET("/top.html", GroupV1.VideosController.Top)
		// Web
		webV1.GET("/web.html", GroupV1.IndexController.Web)
		// About
		webV1.GET("/about.html", GroupV1.IndexController.About)
		// APP
		webV1.GET("/label/app.html", GroupV1.IndexController.App)
		// 电影以及二级分类
		v := webV1.Group("/v")
		{
			v.GET("/:params", GroupV1.VideosController.Category)
		}
		// 片库

		pianku := webV1.Group("/pianku")
		{
			pianku.GET("/:params", GroupV1.VideosController.PianKu)
		}

		arttype := webV1.Group("/arttype")
		{
			// 资讯
			arttype.GET("/zixun.html", GroupV1.ArttypeController.Zixun)
		}

		rss := webV1.Group("/rss")
		{
			// bing
			rss.GET("/bing.xml", GroupV1.SeoController.Bing)
		}
	}
}
