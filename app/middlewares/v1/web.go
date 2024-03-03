package v1

import (
	"github.com/gin-gonic/gin"
	"goapi/app/service"
)

// Web 定义客户端中间件
func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := service.Site(c)
		// 设置语言，默认英文
		SetLang(c, "zh")
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
