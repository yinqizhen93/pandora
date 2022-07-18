package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bilibili/gengine/engine"
	"github.com/mitchellh/mapstructure"
	"plugin"
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

// Result resp
type Result struct {
	MatchSuccess bool
	Rules        []string
	//other params
}

var RiskControl *RiskControlServer

// todo 初始化一定要传入规则

func InitRiskControl() {
	apis := make(map[string]interface{})
	apis["print"] = fmt.Println // 此处注入的应该是单个规则验证无关的通用函数，或全局变量
	apis["appendStr"] = appendStr
	initRule := `rule "empty"  "" salience 0 begin end`
	RiskControl = NewRiskControlServer(10, 20, 1, initRule, apis)
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

var Plug *plugin.Plugin

func (rcs *RiskControlServer) Serve(entry map[string]interface{}, conds []map[string]interface{}, rule string) (*Result, error) {
	rst := &Result{}
	//基于需要注入接口或数据,data这里最好仅注入与本次请求相关的结构体或数据，便于状态管理
	data := make(map[string]interface{})
	plug, err := plugin.Open("data_struct.so")
	if err != nil {
		panic(err)
	}
	mr, err := plug.Lookup("MR")
	if err != nil {
		panic(err)
	}
	entryType, ok := entry["type"]
	if !ok {
		panic("key 'type' not found in entry map")
	}
	entryValue, ok := entry["value"]
	if !ok {
		panic("key 'value' not found in entry map")
	}
	if entryType.(string) == "MedicalRecord" {
		if err := mapstructure.Decode(entryValue, &mr); err != nil {
			panic(err)
		}
		data["mr"] = mr
	}
	for _, cond := range conds {

		v, err := convertType(cond)
		if err != nil {
			panic(err)
		}
		key, ok := cond["key"]
		if !ok {
			panic("key 'key' not found in cond map")
		}
		data[key.(string)] = v
	}
	data["rst"] = rst
	//模块化业务逻辑,api
	fmt.Println(data)
	if err := rcs.Pool.UpdatePooledRules(rule); err != nil {
		panic(err)
	}
	e, _ := rcs.Pool.Execute(data, true)
	if e != nil {
		println(fmt.Sprintf("pool execute rules error: %+v", e))
		return nil, e
	}
	return rst, nil
}

func convertType(cond map[string]interface{}) (interface{}, error) {
	typ, ok := cond["type"]
	if !ok {
		return nil, errors.New("key 'type' not found in cond map ")
	}
	val, ok := cond["value"]
	fmt.Printf("val type: %T\n", val)
	if !ok {
		return nil, errors.New("key 'value' not found in cond map ")
	}
	switch typ {
	case "int":
		var v int
		if err := json.Unmarshal([]byte(val.(string)), &v); err != nil {
			return nil, err
		}
		return v, nil
	case "float":
		var v float64
		if err := json.Unmarshal([]byte(val.(string)), &v); err != nil {
			return nil, err
		}
		return v, nil
	case "string":
		return val, nil
	case "intList":
		var v []int
		for _, e := range val.([]interface{}) {
			v = append(v, e.(int))
		}
		return v, nil
	case "floatList":
		var v []float64
		for _, e := range val.([]interface{}) {
			v = append(v, e.(float64))
		}
		return v, nil
	case "stringList":
		var v []string
		for _, e := range val.([]interface{}) {
			v = append(v, e.(string))
		}
		return v, nil
	}
	return nil, errors.New("unknown type")
}

func main() {
	//go watch()
	InitRiskControl()
	jsonEntry := `
	{
		"value": {
			"Diags": ["R79.000x006", "XC02CAC0.12"],
			"MedicalType": "住院",
			"Items": [
				{
					"Code": "XA12CBL149"
				},
				{
					"Code": "R79.000x006"
				}
			],
			"Age": 121
		},
		"type": "MedicalRecord"
	}`
	jsonCond := `
	[
		{
			"key": "medicalType",
			"value": "住院",
			"type": "string"
		},
		{
			"key": "itemsCode",  
			"value": ["R79.000x006"],
			"type": "stringList"
		}
	]`
	rule := `
		rule "MedicalRecordFor3"  "MedicalRecord ItemCode包含Condition Items的任何一个" salience 0
		begin
			print(3)
			if mr.MedicalType != medicalType {
				return
			} 
			mrItems := mr.Items
			forRange idx := mrItems {
				forRange i := itemsCode {
					elem := mrItems[idx]
					if elem.Code == itemsCode[i] {
						rst.MatchSuccess = true
						rst.Rules = appendStr(rst.Rules, @name)
					}
				}
			}
		end
   `
	entry := make(map[string]interface{})
	cond := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(jsonEntry), &entry); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(jsonCond), &cond); err != nil {
		panic(err)
	}
	resp, err := RiskControl.Serve(entry, cond, rule)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	select {}
}
