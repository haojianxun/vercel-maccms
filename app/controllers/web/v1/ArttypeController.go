package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArttypeController struct {
	BaseController
}

func (h *ArttypeController) Zixun(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["page"] = "zixun"
	data["title"] = "资讯"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "arttype/zixun.html", data)
}
