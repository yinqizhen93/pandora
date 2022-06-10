package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigFile("config/config.dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("获取配置文件失败")
		panic(err)
	}
}
