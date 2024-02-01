package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"goapi/pkg/validate"
	"strings"
	"time"
)

// 登录模块

type LoginController struct {
	BaseController
}

// 登录

func (h *LoginController) Login(c *gin.Context) {
	DB := models.MacAdminMgr(mysql.DB)
	var params requests.MacAdmin
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	option, err := DB.GetByOption(DB.WithAdminName(params.AdminName))
	if err != nil {
		panic(err)
		return
	}
	if !strings.EqualFold(option.AdminPwd, helpers.Md5(params.AdminPwd)) {
		echo.Error(c, "Failed", "密码错误")
		return
	}
	if option.AdminStatus != 1 {
		echo.Error(c, "Failed", "账号状态异常")
		return
	}
	toInt, _ := helpers.IpToInt(c.RemoteIP())
	option.AdminLoginTime = int(time.Now().Unix())
	option.AdminLastLoginTime = int(time.Now().Unix())
	option.AdminLoginIP = int(toInt)
	option.AdminLastLoginIP = int(toInt)
	option.AdminLoginNum = option.AdminLoginNum + 1
	err = DB.Where("admin_id = ?", option.AdminID).Updates(option).Error
	if err != nil {
		echo.Error(c, "Failed", "")
		return
	}
	marshal, _ := json.Marshal(option)
	_, _ = redis.Client.Add("token", string(marshal), 86400)
	echo.Success(c, option, "ok")
}

// 退出登录

func (h *LoginController) Logout(c *gin.Context) {
	token := c.Query("token")
	if len(token) > 0 {
		redis.Client.Delete(token)
	}
	echo.Success(c, nil, "ok")
}
