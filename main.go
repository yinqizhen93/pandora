package main

import (
	"fmt"
	"github.com/spf13/viper"
	"pandora/router"
	"pandora/service"
	"pandora/service/db"
	"pandora/service/logger"
)

func main() {
	err := router.Router.Run(":5001")
	if err != nil {
		panic(err)
	}
}

func init() {
	InitConfig() // todo 初始化顺序有依赖关系，如何解决？
	logger.InitLogger()
	db.InitDB()
	service.InitService()
	router.InitRouter()
	//validate.RegisterValidator()
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
