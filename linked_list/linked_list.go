package linked_list

type Node struct {
	next *Node
	data any
}

func NewNode(d any) *Node {
	return &Node{
		next: nil,
		data: d,
	}
}

func (n *Node) AppendToTail(d any) {
	end := NewNode(d)
	for n.next != nil {
		n = n.next
	}
	n.next = end
}

func (n *Node) DeleteNode(d any) *Node {
	head := n
	if n.data == d {
		return head.next
	}
	for n != nil && n.next !=nil{
		if n.next.data == d {
			n.next = n.next.next
		}
		n = n.next
	}
	return head
}

func (n *Node) TraverseListIntoArray() []any {
	arr := []any{}
	for n != nil {
		arr = append(arr, n.data)
		n = n.next
	}
	return arr
}
