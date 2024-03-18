package config

import (
	"goapi/pkg/config"
)

// email 参数
func init() {
	config.Add("notify", config.StrMap{
		"email": config.Env("maccms.site.site_email", "mail@dyxs.site"),
	})
}
