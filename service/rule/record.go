package main

import (
	"log"
	"time"
)

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

func (mr *MedicalRecord) Print(s interface{}) {
	log.Println(s)
}

func (mr *MedicalRecord) ContainsAnyItem(item []string) bool {
	for _, mrItem := range mr.Items {
		for _, i := range item {
			if mrItem.Code == i {
				return true
			}
		}
	}
	return false
}
