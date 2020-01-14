package main

// import "fmt"

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
if root != nil {
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
	}
}


// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyPostorder(root, fmt.Println)

// }
