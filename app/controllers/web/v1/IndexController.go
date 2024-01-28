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

func (h *IndexController) Web(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "Web"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "web.html", data)
}

func (h *IndexController) App(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "App"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "label/app.html", data)
}

func (h *IndexController) Search(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "Search"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "search.html", data)
}
