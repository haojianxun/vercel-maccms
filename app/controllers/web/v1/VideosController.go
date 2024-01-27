package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideosController struct {
	BaseController
}

func (h *VideosController) Dianying(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "电影"
	c.HTML(http.StatusOK, "v/dianying.html", data)
}

func (h *VideosController) Dianshiju(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "电视剧"
	c.HTML(http.StatusOK, "v/dongman.html", data)
}

func (h *VideosController) Dongman(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "动漫"
	c.HTML(http.StatusOK, "v/dongman.html", data)
}

func (h *VideosController) Zongyi(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "综艺"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "v/zongyi.html", data)
}

func (h *VideosController) Play(c *gin.Context) {
	value, exists := c.Get("data")
	if !exists {
		value = gin.H{}
	}
	data := value.(gin.H)
	data["title"] = "Play"
	data["list"] = gin.H{"asas": "asas"}
	c.HTML(http.StatusOK, "v/play.html", data)
}
