package service

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/ent-adapter"
	"github.com/spf13/viper"
	"log"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.username")
	passwd := viper.GetString("database.password")
	database := viper.GetString("database.database")
	//maxConnPool := viper.GetInt("database.maxConnPool")
	//maxIdleConns := viper.GetInt("database.maxIdleConns")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, database)
	a, err := entadapter.NewAdapter("mysql", dsn)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("config/rbac_casbin.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		panic(err)
		return
	}
	//从DB加载策略
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	Enforcer = e
	fmt.Println(Enforcer.GetPolicy())
}
