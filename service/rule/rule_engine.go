package main

import (
	"fmt"
	"github.com/bilibili/gengine/engine"
)

// RiskControlServer 业务接口
type RiskControlServer struct {
	//gengine pool
	Pool        *engine.GenginePool
	poolSize    int
	maxIdlePool int
	rules       string
	//other params
}

//// Request request
//type Request struct {
//	Rid       int64
//	RuleNames []string
//	//other params
//}
//

// Result resp
type Result struct {
	MatchSuccess bool
	Rules        []string
	//other params
}

type Rule struct {
}

var RiskControl *RiskControlServer

func appendStr(slice []string, elems ...string) []string {
	return append(slice, elems...)
}

func InitRiskControl() {
	apis := make(map[string]interface{})
	apis["print"] = fmt.Println // 此处注入的应该是单个规则验证无关的通用函数，或全局变量
	apis["appendStr"] = appendStr
	RiskControl = NewRiskControlServer(10, 20, 1, rule1, apis)
}

func NewRiskControlServer(poolMinLen, poolMaxLen int64, em int, rulesStr string, apiOuter map[string]interface{}) *RiskControlServer {
	pool, e := engine.NewGenginePool(poolMinLen, poolMaxLen, em, rulesStr, apiOuter)
	if e != nil {
		panic(fmt.Sprintf("初始化gengine失败，err:%+v", e))
	}
	rcs := &RiskControlServer{Pool: pool}
	return rcs
}

//service

func (rcs *RiskControlServer) Serve(mr *MedicalRecord, cond Condition) (*Result, error) {
	rst := &Result{}
	//基于需要注入接口或数据,data这里最好仅注入与本次请求相关的结构体或数据，便于状态管理
	data := make(map[string]interface{})
	data["mr"] = mr
	data["cond"] = cond
	data["rst"] = rst
	//模块化业务逻辑,api
	e, _ := rcs.Pool.ExecuteConcurrent(data)
	if e != nil {
		println(fmt.Sprintf("pool execute rules error: %+v", e))
		return nil, e
	}
	return rst, nil
}

func main() {
	InitRiskControl()
	mr := &MedicalRecord{
		Diags:       []string{"R79.000x006", "XC02CAC0.12"},
		MedicalType: "住院",
		Items: []Item{
			{
				Code: "XA12CBL149",
			},
			{
				Code: "R79.000x006",
			},
		},
	}
	cond := Condition{
		MedicalType: "住院",
		ItemsCode:   []string{"R79.000x006"},
	}
	resp, err := RiskControl.Serve(mr, cond)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
