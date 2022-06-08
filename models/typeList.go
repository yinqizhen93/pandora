package models

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next  *Node
}

type LinkList struct {
	Head *Node
	Pointer []interface{}
}

type LinkListMethod interface {
	append(v interface{})
	delete(i int)
	insert(i int, v interface{})
	len()
}

func (l *LinkList) append(v interface{}) {
	cur := l.Head
	for {
		if cur.Next == nil {
			cur.Next = CreateNode(v)
			break
		}
		cur = cur.Next
	}
}

func (l *LinkList) delete(i int) {
	fmt.Println("delete method")
}

func (l *LinkList) insert(i int, v interface{}) {
	fmt.Println("insert method")
}

func (l *LinkList) len() int {
	cur := l.Head
	i := 0
	for {
		if cur.Next == nil {
			return i
		}
		i++
		cur = cur.Next
	}
}

func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}

func InitLinkList() *LinkList {
	head := CreateNode(nil)
	return &LinkList{head, make([]interface{}, 0)}
}
//
//func PrintLinkList(l *LinkList) {
//	if l.Length == 0 {
//		fmt.Println("empty LinkList")
//		return
//	}
//	cur := l.Head.Next
//	for i:=1; i<=l.Length; i++ {
//		fmt.Println("value", cur.Data)
//		cur = cur.Next
//	}
//}

func main() {
	list := InitLinkList()

	list.append(1)
	list.append("aaa")
	fmt.Println(list.len())
}
