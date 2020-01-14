package student

// import "fmt"


func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {
	if root != nil {
		for i := 0; bTreeApplyByLevel(root, f, i); i++ {
		
		}
	
	}
}
func bTreeApplyByLevel(root *TreeNode, f func (...interface{}) (int, error), level int) bool {
	if root == nil {
		return false
	}
	if level == 0 {
		f(root.Data)
		return true
	}
	left := bTreeApplyByLevel(root.Left, f, level - 1)
	right := bTreeApplyByLevel(root.Right, f, level - 1)
	
	return left || right
}

// func main() {
// 	root := &TreeNode{Data: "5"}
// 	BTreeInsertData(root, "3")
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "2")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "6")
// 	BTreeInsertData(root, "8")
// 	BTreeApplyByLevel(root, fmt.Println)
// }
