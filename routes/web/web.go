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
		webV1.GET("/prestrain.html", GroupV1.IndexController.PresTrain)
		webV1.GET("/player/dplayer.html", GroupV1.IndexController.DPlayer)
		// 片库
		webV1.GET("/pianku.html", GroupV1.VideosController.PianKu)
		// 播放
		webV1.GET("/play-:params", GroupV1.VideosController.Play)
		// 搜索
		{
			// 搜索
			webV1.GET("/search.html", GroupV1.IndexController.Search)
			webV1.GET("/search:search", GroupV1.IndexController.Search)
			//// 电影名称搜索
			//// search-:电影名称----------:第几页---.html
			//webV1.GET("/search-:movieName----------:page---.html", GroupV1.IndexController.Search)
			//// search--:演员名称---------:第几页---.html
			//webV1.GET("/search--:actorName---------:page---.html", GroupV1.IndexController.Search)
			//// search---:大陆--------:第几页---.html
			//webV1.GET("/search---:areaName--------:page---.html", GroupV1.IndexController.Search)
			//// search------:导演-----:第几页---.html
			//webV1.GET("/search------:directorName-----:page---.html", GroupV1.IndexController.Search)
			//// search-----------:第几页---:年份.html
			//webV1.GET("/search-----------:page---:year.html", GroupV1.IndexController.Search)

			//-电影名称----------第几页---
			//--演员名称---------第几页---
			//---大陆--------第几页---
			//------导演-----第几页---
			//-----------第几页---年份
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
