package main

// import 
// 	"fmt"

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
	}
	
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "3")
// 	BTreeInsertData(root, "8")
// 	BTreeInsertData(root, "9")
// 	BTreeApplyInorder(root, fmt.Println)
// }
