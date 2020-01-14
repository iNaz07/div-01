package main

// import 
//     "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

// func main() {
//      root := &TreeNode{Data: "6"}
//      BTreeInsertData(root, "9")
//      BTreeInsertData(root, "8")
// 	 BTreeInsertData(root, "1")
//      fmt.Println(root.Left.Data)
//      fmt.Println(root.Data)
//      fmt.Println(root.Right.Left.Data)
//      fmt.Println(root.Right.Data)

// }
