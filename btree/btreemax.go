package main

import "fmt"

func BTreeMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	// BTreeInsertData(root, "8")
	BTreeInsertData(root, "5")
	// BTreeInsertData(root, "9")
	// BTreeInsertData(root, "2")
	max := BTreeMax(root)
	fmt.Println(max.Data)
}
