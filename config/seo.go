package config

import (
	"goapi/pkg/config"
)

// redis 参数
func init() {
	config.Add("seo", config.StrMap{
		"bing.apikey":  config.Env("bing.apikey", ""),
		"baidu.apikey": config.Env("baidu.apikey", ""),
		"sogou.apikey": config.Env("sogou.apikey", ""),
	})
}
