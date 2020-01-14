package student

//import "fmt"

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	
	prev := n1
	cur := n1

	for n1 != nil {
		cur = prev.Next
		if cur.Next == nil {
			cur.Next = n2
			return SortList(n1)
		}
		prev = cur 
	}
	if n1 == nil {
		n1 = n2
	}
	return SortList(n1)
}

func SortList(l *NodeI) *NodeI {
	prev := l
	for prev != nil {
		cur := prev.Next
		for cur != nil {
			if prev.Data > cur.Data {
				prev.Data, cur.Data = cur.Data, prev.Data
			}
			cur = cur.Next
		}
		prev = prev.Next
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

/*func main() {
	var link *NodeI
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 12)
	link = listPushBack(link, 7)
	link = listPushBack(link, -9)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}
*/