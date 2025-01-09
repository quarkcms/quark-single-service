package env

import (
	"fmt"

	"github.com/spf13/viper"
)

// 设置值
func Set(key string, value interface{}) {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.Set(key, value)
	// 写入配置文件
	err := viper.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error writing config file: %w", err))
	}
}

// 获取值
func Get(key ...string) interface{} {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		if len(key) == 2 {
			if viper.Get(key[0]) == nil {
				return key[1]
			}
		} else {
			return nil
		}
	}

	if len(key) == 2 {
		if viper.Get(key[0]) == nil {
			return key[1]
		}
	}

	return viper.Get(key[0])
}
