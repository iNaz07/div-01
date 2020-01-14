package main

// import "fmt"


func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
	}

}

// func main() {
// 	root := &TreeNode{Data: "6"}
// 	BTreeInsertData(root, "4")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "8")
// 	BTreeApplyPreorder(root, fmt.Println)

// }
