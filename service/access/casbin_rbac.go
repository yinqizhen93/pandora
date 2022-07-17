package access

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"log"
	"os"
	"pandora/ent"
	"path"
)

var policyFile = "rbac_casbin.conf"

var _ RBAC = CasbinRBAC{}

type CasbinRBAC struct {
	e *casbin.Enforcer
	a persist.Adapter
}

func NewCasbinRBAC(db *ent.Client) *CasbinRBAC {
	//dsn := getDsn(conf)
	//

	//if err != nil {
	//	log.Printf("连接数据库错误: %v", err)
	//	return nil
	//}
	a := NewEnta(db)
	e, err := casbin.NewEnforcer(getModelFile(), a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		panic(err)
		return nil
	}
	//从DB加载策略
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	cr := CasbinRBAC{
		e: e,
	}
	return &cr
	//fmt.Println(Enforcer.GetPolicy())
}

func (cr CasbinRBAC) HasAccess(user, url, method string) bool {
	ok, err := cr.e.Enforce(user, url, method)
	if err != nil {
		panic(err)
	}
	return ok
}

func getModelFile() string {
	var file string
	if fileFolder := os.Getenv("PANDORA_STATIC"); fileFolder != "" {
		file = path.Join(fileFolder, policyFile)
		return file
	}
	file = path.Join("configs", policyFile)
	return file
}

//func getDsn(conf config.Config) string {
//	host := conf.GetString("database.host")
//	port := conf.GetString("database.port")
//	user := conf.GetString("database.username")
//	passwd := conf.GetString("database.password")
//	database := conf.GetString("database.database")
//	//maxConnPool := viper.GetInt("database.maxConnPool")
//	//maxIdleConns := viper.GetInt("database.maxIdleConns")
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		user, passwd, host, port, database)
//	return dsn
//}
