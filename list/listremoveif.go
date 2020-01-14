package student

//import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	prev := l.Head
	cur := l.Head

	for cur != nil && cur.Data == data_ref {
		l.Head = l.Head.Next
		cur = l.Head
	}
	for cur != nil {
		for cur != nil && cur.Data != data_ref {
			prev = cur
			cur = cur.Next
		}
		prev.Next = cur.Next
		cur = prev.Next
	}
}

/*func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}*/

func ListPushBack(l *List, data interface{}) *List {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return l
	}
	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	fmt.Println("----normal state----")
// 	ListPushBack(link2, 1)
// 	PrintList(link2)
// 	ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "There")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "How")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)
// 	PrintList(link)

// 	ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
// }
