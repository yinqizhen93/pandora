package main

type TreeNode struct {
	val   interface{}
	left  *TreeNode
	right *TreeNode
}

func (tn *TreeNode) Insert() {
	
}

func CreateTree(val interface{}) TreeNode {
	return TreeNode{
		val: val,
	}
}
