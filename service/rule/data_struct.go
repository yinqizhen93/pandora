package main

import "time"

type MedicalRecord struct {
	Diags []string
	// 主诊断
	MainDiag string
	// 次要诊断
	SecondDiags []string
	// 主手术
	MainOperation string
	// 次要手术
	SecondOperations []string
	// 出院时间
	DischargeAt time.Time
	// 住院时长
	HospitalDays int
	// 就诊类型
	MedicalType string
	// item 明细
	Items []Item
	// 参保类型
	InsuranceType string
	Age           int
}

type Item struct {
	// 类型
	ItemType string
	// Item name
	Name string
	// Item code
	Code string
	Nums float64
	// 数量
	Amount float64
	// 处方时间
	PrescriptionTime time.Time
}

// Condition 知识点涉及字段
type Condition struct {
	MedicalType      string
	ItemsCode        []string
	RelatedItemsCode []string
	DiagsCode        []string
	HospitalLevel    int
	Age              int
	InsuranceType    string
	Threshold        map[string]int // 各种阈值
}

var MR = MedicalRecord{}

var Cond = Condition{}
