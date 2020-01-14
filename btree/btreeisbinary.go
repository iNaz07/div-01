package main

import "fmt"

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
	return !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) 
			
		
		return false
	}
	return true	
}


func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(BTreeIsBinary(root))
}
