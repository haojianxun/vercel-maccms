package config

import (
	"goapi/pkg/config"
)

// app 参数
func init() {
	config.Add("app", config.StrMap{

		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "goapi"),

		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"pprof": config.Env("APP_PPROF", false),

		// 应用服务端口
		"port": config.Env("APP_PORT", "80"),

		// 用以生成链接
		"url": config.Env("APP_URL", "http://127.0.0.1"),
	})
}
