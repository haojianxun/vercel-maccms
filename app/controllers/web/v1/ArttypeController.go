package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArttypeController struct {
	BaseController
}

func (h *ArttypeController) Zixun(c *gin.Context) {
	c.HTML(http.StatusOK, "arttype/zixun.html", gin.H{
		"title": "资讯",
	})
}
