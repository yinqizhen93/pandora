package access

import (
	"fmt"
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
	fmt.Println(cr.e.GetPolicy())
	return &cr
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
