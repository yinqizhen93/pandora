package main

import (
	"fmt"
)

func appendStr(slice []string, elems ...string) []string {
	return append(slice, elems...)
}

func toInt(val interface{}) int {
	v, ok := val.(int)
	if !ok {
		panic(fmt.Sprintf("%v can not convert to int", v))
	}
	return v
}

//func toDate(val interface{}) time.Time {
//	v, ok := val.(time.Time)
//	if !ok {
//		panic(fmt.Sprintf("%v can not convert to time", v))
//	}
//	return v
//}

func toIntArray(val interface{}) []int {
	v, ok := val.([]int)
	if !ok {
		panic(fmt.Sprintf("%v can not convert to int array", v))
	}
	return v
}

func toStringArray(val interface{}) []string {
	v, ok := val.([]string)
	if !ok {
		panic(fmt.Sprintf("%v can not convert to string array", v))
	}
	return v
}

//func toArrayItem(val interface{}) []Item {
//	v, ok := val.([]string)
//	if !ok {
//		panic(fmt.Sprintf("%v can not convert to string array", v))
//	}
//	return v
//}
