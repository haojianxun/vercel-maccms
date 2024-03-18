package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"strings"
)

func Site(c *gin.Context) gin.H {
	PageData := gin.H{}
	HttpStatic := "https://static.dyxs.site"
	HTTP := "https://"
	if strings.Contains(c.Request.Host, "localhost") || strings.Contains(c.Request.Host, "127.0.0.1") {
		HTTP = "http://"
		HttpStatic = fmt.Sprintf("%s%s/statics", HTTP, c.Request.Host)
	}
	PageData["__URL__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, c.Request.URL.String())
	PageData["__STATIC__"] = HttpStatic
	PageData["__CSS__"] = fmt.Sprintf("%s%s", HttpStatic, "/css")
	PageData["__JS__"] = fmt.Sprintf("%s%s", HttpStatic, "/js")

	var NavMenus []models.MacType
	ListMacType("NavMenus", 0, &NavMenus)
	PageData["NavMenus"] = NavMenus
	return PageData
}
