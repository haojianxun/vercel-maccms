package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArttypeController struct {
	BaseController
}

func (h *ArttypeController) Zixun(c *gin.Context) {
	DATA := GetDATA(c)
	DATA["page"] = "zixun"
	DATA["title"] = "资讯"
	DATA["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "arttype/zixun.html", DATA)
}
