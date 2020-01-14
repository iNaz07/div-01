package main

// import "fmt"

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil && elem == root.Data {
		return root
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
		return BTreeSearchItem(root.Right, elem)	
	return nil	
}

// func main() {
// 	root := &TreeNode{Data: "7"}
// 	BTreeInsertData(root, "4")
// 	BTreeInsertData(root, "3")
// 	BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "9")
// 	BTreeInsertData(root, "8")
// 	selected := BTreeSearchItem(root, "3")
// 	fmt.Print("Item selected -> ")
// 	if selected != nil {
// 		fmt.Println(selected.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Parent of selected item -> ")
// 	if selected.Parent != nil {
// 		fmt.Println(selected.Parent.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Left child of selected item -> ")
// 	if selected.Left != nil {
// 		fmt.Println(selected.Left.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Right child of selected item -> ")
// 	if selected.Right != nil {
// 		fmt.Println(selected.Right.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}
// }
