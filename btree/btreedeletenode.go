package student

// import "fmt"



func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	} 
	if node.Right == nil && node.Left == nil {
		if node == node.Parent.Right {
			node.Parent.Right = nil
		} else {
			node.Parent.Left = nil
		}		 
	} else if node.Right != nil && node.Left == nil {
		if node == node.Parent.Right {
			node.Parent.Right = node.Right
		} else {
			node.Parent.Left = node.Right
		}
	} else if node.Left != nil && node.Right == nil {
		if node == node.Parent.Right {
			node.Parent.Right = node.Left
		} else {
			node.Parent.Left = node.Left
		}
	} else {
		temp := BMin(node.Right)
		node.Data = temp.Data
		BTreeDeleteNode(node.Right, temp) 
	}
	return root
}
func BMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
// func main() {
// 	root := &TreeNode{Data: "5"}
// 	BTreeInsertData(root, "3")
// 	BTreeInsertData(root, "4")
// 	BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "9")
// 	BTreeInsertData(root, "6")
// 	BTreeInsertData(root, "8")
	

// 	node := BTreeSearchItem(root, "9")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }