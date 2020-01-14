package main

// import "fmt"

func BTreeLevelCount(root *TreeNode) int {
	
 	if root != nil {
	sizeLeft := BTreeLevelCount(root.Left) + 1
	sizeRight := BTreeLevelCount(root.Right) + 1
	
		if sizeLeft > sizeRight {
	return sizeLeft
} else {
	return sizeRight
	}
} 
return 0
}



// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "3")
// 	BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "6")
// 	fmt.Println(BTreeLevelCount(root))
// }
