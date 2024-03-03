package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Site(c *gin.Context) gin.H {
	data := gin.H{}
	HTTP := "https://"
	data["__URL__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, c.Request.URL.String())
	data["__STATIC__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics")
	data["__CSS__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics/css")
	data["__JS__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics/js")
	return data
}
