package config

import (
	"goapi/pkg/config"
)

// app 参数
func init() {
	config.Add("app", config.StrMap{

		// 应用名称，暂时没有使用到
		"name": config.Env("app.app_name", "goapi"),

		// 当前环境，用以区分多环境
		"env": config.Env("app.app_env", "production"),

		// 是否进入调试模式
		"pprof": config.Env("app.app_url", false),

		// 应用服务端口
		"port": config.Env("app.app_port", "80"),

		// 用以生成链接
		"url": config.Env("app.app_pprof", "http://127.0.0.1"),
	})
}
