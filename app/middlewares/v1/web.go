package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Web 定义客户端中间件
func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := gin.H{}
		HTTP := "http://"
		data["__URL__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, c.Request.URL.String())
		data["__STATIC__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics")
		data["__CSS__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics/css")
		data["__JS__"] = fmt.Sprintf("%s%s%s", HTTP, c.Request.Host, "/statics/js")
		// 设置语言，默认英文
		SetLang(c, "en")
		path := c.FullPath()
		switch path {
		case "/app/v1/wechat/userSynchronize.json":
		case "/app/v1/config/setKlineToken.json":
		case "/app/v1/wechat/Send.json":
		case "/app/v1/wechat/init.json":
		case "/app/v1/seo/getGoogleIndex":
			// 继续往下面执行
			c.Set("data", data)
			c.Next()
			break
		default:
			c.Set("data", data)
			c.Next()
			break
		}
	}
}
