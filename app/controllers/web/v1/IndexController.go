package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseController
}

func (h *IndexController) Index(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "index"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "index.html", data)
}

func (h *IndexController) Home(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "综艺"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "home.html", data)
}
