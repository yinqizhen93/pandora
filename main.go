package main

import (
	"fmt"
	"github.com/spf13/viper"
	"pandora/db"
	"pandora/router"
	"pandora/service"
)

func main() {
	err := router.Router.Run(":5001")
	if err != nil {
		panic(err)
	}
}

func init() {
	InitConfig()
	service.InitLogger()
	db.InitDB()
	service.InitService()
	router.InitRouter()
}

func InitConfig() {
	viper.SetConfigFile("config/config.dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("获取配置文件失败")
		panic(err)
	}
	viper.WatchConfig() //监听配置文件变化
}
