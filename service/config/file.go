package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path"
)

var configFile = "config.dev.yaml"

type FileConfig struct {
	file    string
	watcher fsnotify.Watcher
	config  *viper.Viper
}

var _ Config = &FileConfig{}

func NewFileConfig() *FileConfig {
	file := getConfigFile()
	fc := &FileConfig{
		file:   file,
		config: viper.New(),
	}
	fc.Load()
	go fc.Watch()
	return fc
}

// getConfigFile 获取配置文件， 配置文件所在文件夹configs需要与go可执行文件在同一个层级， go相对路径是相对可执行文件
func getConfigFile() (file string) {
	if fileFolder := os.Getenv("PANDORA_STATIC"); fileFolder != "" {
		file = path.Join(fileFolder, configFile)
		return
	}
	file = path.Join("configs", configFile)
	return
}

func (fc *FileConfig) Load() {
	fc.config.SetConfigFile(fc.file)
	fmt.Println(fc.file)
	err := fc.config.ReadInConfig()
	if err != nil {
		fmt.Printf("获取配置文件%s失败", fc.file)
		panic(err)
	}
}

func (fc *FileConfig) Get(key string) interface{} {
	return fc.config.Get(key)
}

func (fc *FileConfig) GetInt(key string) int {
	return fc.config.GetInt(key)
}

func (fc *FileConfig) GetString(key string) string {
	return fc.config.GetString(key)
}

func (fc *FileConfig) GetIntSlice(key string) []int {
	return fc.config.GetIntSlice(key)
}

func (fc *FileConfig) GetStringSlice(key string) []string {
	return fc.config.GetStringSlice(key)
}

func (fc *FileConfig) Watch() { // todo need a close method to stop watch in case of memory leak
	fc.config.WatchConfig() //监听配置文件变化
}
