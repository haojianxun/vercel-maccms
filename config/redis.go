package config

import (
	"fmt"
	"goapi/pkg/config"
)

// redis 参数
func init() {
	fmt.Println("初始化redis配置")
	config.Add("redis", config.StrMap{
		"host":     config.Env("redis.host", "127.0.0.1"),
		"port":     config.Env("redis.port", 6379),
		"password": config.Env("redis.password", ""),
		"db":       config.Env("redis.db", 0),
	})
}
