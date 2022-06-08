package main

import "fmt"

type List struct {
	data *[]interface{}
}

type t struct {
 b int
}



func main() {
	//b := t{b:1}
	//a := make([]interface{}, 0)
	//fmt.Println(a)
	a := []int{1,2,3,45}
	b := a[1:3:3]
	fmt.Println(b)
	fmt.Println("b容量", cap(b))
	b[1] = 123
	b = append(b, 0)
	b = append(b, 0)
	b[0] = 88
	fmt.Println("容量", cap(b))
	fmt.Println(a)
	fmt.Println(b)

	//fmt.Println(a)
	//fmt.Println(reflect.TypeOf(a))
}