package config

import (
	"github.com/quarkcloudio/quark-smart/v2/pkg/env"
)

type RedisConfig struct {
	Host     string // 地址
	Password string // 密码
	Port     string // 端口
	Database int    // 数据库
}

// Redis配置信息
var Redis = &RedisConfig{

	// 地址
	Host: env.Get("REDIS_HOST", "").(string),

	// 密码
	Password: env.Get("REDIS_PASSWORD").(string),

	// 端口
	Port: env.Get("REDIS_PORT", "6379").(string),

	// 数据库
	Database: 0,
}
