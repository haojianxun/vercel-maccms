package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
)

func Site(c *gin.Context) gin.H {
	PageData := gin.H{}
	HTTP := "https://"
	PageData["__URL__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, c.Request.URL.String())
	HttpStatic := "https://static.dyxs.site"
	PageData["__STATIC__"] = HttpStatic
	PageData["__CSS__"] = fmt.Sprintf("%s%s", HttpStatic, "/css")
	PageData["__JS__"] = fmt.Sprintf("%s%s", HttpStatic, "/js")

	var NavMenus []models.MacType
	ListMacType("NavMenus", 0, &NavMenus)
	PageData["NavMenus"] = NavMenus
	return PageData
}
