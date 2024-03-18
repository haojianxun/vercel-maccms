package config

import (
	"goapi/pkg/config"
)

// 数据库连接参数
func init() {
	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{
			// 数据库连接信息
			"host":     config.Env("database.host", "127.0.0.1"),
			"port":     config.Env("database.port", "3306"),
			"database": config.Env("database.name", "goapi"),
			"username": config.Env("database.username", "root"),
			"password": config.Env("database.password", "root"),
			"prefix":   config.Env("database.prefix", "root"),
			"charset":  "utf8mb4",
			// 连接池配置
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 25),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 100),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 60),
		},
	})
}
