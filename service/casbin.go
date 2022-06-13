package service

import (
	"github.com/casbin/casbin/v2"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	//a, err := gormadapter.NewAdapterByDB(db.DB)
	//if err != nil {
	//	log.Printf("连接数据库错误: %v", err)
	//	return
	//}
	//e, err := casbin.NewEnforcer("config/rbac_casbin.conf", a)
	//if err != nil {
	//	log.Printf("初始化casbin错误: %v", err)
	//	panic(err)
	//	return
	//}
	////从DB加载策略
	//err = e.LoadPolicy()
	//if err != nil {
	//	panic(err)
	//}
	//Enforcer = e
	//fmt.Println(Enforcer.GetPolicy())
}
