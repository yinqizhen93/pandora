package main

import "fmt"

// BinarySearch 二分法查找
func BinarySearch(arr []int, left, right, target int) int {
	if right < left {
		// 没有匹配
		return -1
	}
	//mid := (right + left) / 2
	//不使用上面加法，避免出现和超过int能代表最大值
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearch(arr, mid+1, right, target)
	} else {
		return BinarySearch(arr, left, mid-1, target)
	}
}

func main() {
	arr := []int{1, 3, 4, 5, 7, 9, 20, 53, 66, 77, 99, 111, 322, 1000}
	left := 0
	right := len(arr) - 1
	a := BinarySearch(arr, left, right, 53)
	fmt.Printf("return value:%d\n", a)
}
