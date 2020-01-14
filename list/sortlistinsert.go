package student

//import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	prev := l
	after := l
	node := l
	n := &NodeI{Data: data_ref}
	if prev.Data > data_ref {
		node = n
		node.Next = l
		l = node
		return l
	} else if prev.Next == nil {
		node = n
		l.Next = node 
		return l
	}	
	for prev.Next != nil {
		after = prev.Next
		node = n
		if after.Data > data_ref {
			
			prev.Next = node
			node.Next = after			
			return l
		} else if after.Next == nil {
			after.Next = node
			node.Next = nil
		}
		prev = after 
	}
	return l
}

/*func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
*/
func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

// func main() {

// 	var link *NodeI

// 	link = listPushBack(link, 1)
// 	link = listPushBack(link, 4)
// 	link = listPushBack(link, 9)

// 	PrintList(link)

// 	link = SortListInsert(link, 5)
// 	link = SortListInsert(link, 10)
// 	PrintList(link)
// }
